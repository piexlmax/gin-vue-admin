package sensitive_word

import (
	"gin-vue-admin/global"
	"sync"
	"unicode/utf8"

	"go.uber.org/zap"
)

type Library struct {
	root *trieNode
	mu   *sync.RWMutex
}

func (this *Library) insertWord(word []rune) {
	currNode := this.root
	for _, c := range word {
		if cildNode, exist := currNode.children[c]; !exist {
			cildNode = newtrieNode()
			currNode.children[c] = cildNode
			currNode = cildNode
		} else {
			currNode = cildNode
		}
	}

	currNode.isEndOfWord = true
}

func (this *Library) startsWith(prefix []rune) bool {
	currNode := this.root
	for _, c := range prefix {
		if cildNode, exist := currNode.children[c]; !exist {
			return false
		} else {
			currNode = cildNode
		}
	}

	return true
}

func (this *Library) searcWord(word []rune) bool {
	currNode := this.root
	for _, c := range word {
		if cildNode, exist := currNode.children[c]; !exist {
			return false
		} else {
			currNode = cildNode
		}
	}

	return currNode.isEndOfWord
}

func (this *Library) searcSentence(sentence string) (matchIndexList []*matchIndex) {
	start, end := 0, 1
	sentenceRuneList := []rune(sentence)

	startsWith := false
	for end <= len(sentenceRuneList) {
		if this.startsWith(sentenceRuneList[start:end]) {
			startsWith = true
			end += 1
		} else {
			flag := false
			if startsWith == true {
				for index := end - 1; index > start; index-- {
					if this.searcWord(sentenceRuneList[start:index]) {
						flag = true
						matchIndexList = append(matchIndexList, newMatchIndex(start, index-1))
						break
					}
				}
			}
			if flag {
				start = end - 1
			} else {
				start = start + 1
			}
			end = start + 1
			startsWith = false
		}
	}

	if startsWith {
		for index := end - 1; index > start; index-- {
			if this.searcWord(sentenceRuneList[start:index]) {
				matchIndexList = append(matchIndexList, newMatchIndex(start, index-1))
				break
			}
		}
	}
	return
}

func (this *Library) HandleWord(sentence string, replaceCh rune) (string, bool) {
	this.mu.RLock()
	defer this.mu.RUnlock()
	matchIndexList := this.searcSentence(sentence)
	var check bool

	if len(matchIndexList) == 0 {
		return sentence, check
	}

	sentenceList := []rune(sentence)
	if len(matchIndexList) > 0 {
		check = true
	} else {
		check = false
	}
	for _, matchIndexObj := range matchIndexList {
		for index := matchIndexObj.start; index <= matchIndexObj.end; index++ {
			sentenceList[index] = replaceCh
		}
	}

	return string(sentenceList), check
}

func (this *Library) Contains(k string) bool {
	this.mu.RLock()
	defer this.mu.RUnlock()
	match := this.searcSentence(k)
	if len(match) == 1 {
		if item := match[0]; item.end-item.start == utf8.RuneCountInString(k)-1 {
			return true
		}
	}
	return false
}

func (l *Library) Build() *Library {
	// TODO read from redis
	list, err := LoadMySQLWords()
	if err != nil {
		global.GVA_LOG.Error("query sensitive words fail", zap.Any("error", err.Error()))
		return nil
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	l.root = newtrieNode()
	for _, v := range list {
		wordRuneList := []rune(v.Word)
		if len(wordRuneList) > 0 {
			l.insertWord(wordRuneList)
		}
	}
	return l
}

type matchIndex struct {
	start int
	end   int
}

func newMatchIndex(start, end int) *matchIndex {
	return &matchIndex{
		start: start,
		end:   end,
	}
}
