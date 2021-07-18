// +build go1.16

package main

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"os"
)

//go:embed resource config.yaml
var resource embed.FS

func restoreFolder(dirPath string) {
	fis, err := resource.ReadDir(dirPath)
	if err != nil {
		fmt.Println("[embed restore resource file] err:", err)
		return
	}
	if _, err := os.Stat(dirPath); err != nil && os.IsNotExist(err) {
		os.MkdirAll(dirPath, os.ModePerm)
	}

	for _, fi := range fis {
		if !fi.IsDir() {
			restoreFile(dirPath, fi)
		} else {
			newDirPath := ""
			if dirPath != "." {
				newDirPath = dirPath + "/" + fi.Name()
			} else {
				newDirPath = fi.Name()
			}
			restoreFolder(newDirPath)
		}
	}
}

func restoreFile(dirPath string, f fs.DirEntry) {
	filePath := ""
	if dirPath == "." {
		filePath = f.Name()
	} else {
		filePath = dirPath + string(os.PathSeparator) + f.Name()
	}

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		srcFile, err2 := resource.Open(filePath)
		if err2 != nil {
			fmt.Println("[embed restore resource file] err:", err2)
			return
		}

		dstFile, err2 := os.Create(filePath)
		if err2 != nil {
			fmt.Println("[embed restore resource file] err:", err2)
		}
		if _, err2 = io.Copy(dstFile, srcFile); err2 != nil {
			fmt.Printf("[embed restore resource file] write file failed: %s\n", dstFile)
		}
		srcFile.Close()
		dstFile.Close()
	} else {
		fmt.Printf("[embed restore resource file] file exist, skip: %s\n", filePath)
	}
}

func init() {
	restoreFolder(".")
}
