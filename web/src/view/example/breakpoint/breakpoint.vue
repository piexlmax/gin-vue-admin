<template>
  <div class="hello">
    <el-divider content-position="left">大文件上传</el-divider>
    <form id="fromCont" method="post">
      <div class="fileUpload" @click="inputChange">
        选择文件
        <input v-show="false" id="file" ref="Input" multiple="multiple" type="file" @change="choseFile">
      </div>
    </form>
    <el-button :disabled="limitFileSize" type="primary" size="mini" class="uploadBtn" @click="getFile">上传文件（超大文件点击此按钮会卡顿，等待即可）</el-button>
    <el-progress v-if="file" :text-inside="true" :show-text="true" :format="function (p) {
      return p !== 100 ? p + '% 正在处理视频切片，请稍等...' : p + '% 视频切片处理完成，请点击上传'
    }" :stroke-width="26" :percentage="localProcessPercentage" status="success"></el-progress>
    <div class="el-upload__tip">请上传不超过 10GB 的文件</div>
    <div class="list">
      <transition name="list" tag="p">
        <div v-if="uploading" class="list-item">
          <i class="el-icon-document" />
          <span>{{ file.name }}</span>
          <span class="percentage">{{ percentage }}%</span>
          <el-progress :show-text="false" :text-inside="false" :stroke-width="2" :percentage="percentage" />
          <div class="tips">进度条长时间卡住，重新点击上传即可</div>
        </div>
      </transition>
    </div>
    <div class="tips">此版本为先行体验功能测试版，样式美化和性能优化正在进行中</div>
    <div class="tips">上传切片文件和合成的完整文件分别在server目录的breakpointDir文件夹和fileDir文件夹</div>
    <div class="tips">超大文件在后台合并时比较占用资源</div>
  </div>
</template>

<script>
import SparkMD5 from 'spark-md5'
import { mapGetters } from "vuex";
import axios from 'axios'
import {
  findFile,
  breakpointContinueFinish,
  removeChunk
} from '@/api/breakpoint'
export default {
  name: 'HelloWorld',
  data() {
    return {
      file: null,
      fileMd5: '',
      successFile: '',
      failedCount: 0,
      parts: [],
      waitUpLoad: [],
      waitNum: 0,
      limitFileSize: false,
      uploading: false,
      localProcessPercentage: 0,
      percentage: 0,
      customColor: '#409eff',
      sliceProcessing: false,
    }
  },
  computed: {
    ...mapGetters('user', ['userInfo', 'token']),
  },
  methods: {
    init() {
      this.sliceProcessing = true;
      this.uploading = false;
      this.file = null;
      this.parts = [];
      this.percentage = 0;
      this.failedCount = 0;
      this.localProcessPercentage = 0;
      this.waitNum = 0;
      this.waitUpLoad = [];
      this.sliceProcessing = false;
    },
    // 选中文件的函数
    async choseFile(e) {
      this.init();

      let file = e.target.files[0];                          // shortcut to file
      let fileSize = file.size;                          // total size of file
      let sliceSize = 1 << 23;                             // slice size, here 8 mb
      let count = Math.ceil(fileSize / sliceSize);       // calc number of parts
      let index = 0;                                     // part/slice index

      this.file = file;
      const spark = new SparkMD5.ArrayBuffer()

      let _slice = () => {                               // async "loop"
        let start = index * sliceSize;                   // calc start pos. of current slice
        let end = Math.min(start + sliceSize, fileSize); // calc end pos. of current slice
        let slice = file.slice(start, end);              // slice the file blob to new blob
        let fr = new FileReader();                       // read blob part as ArrayBuffer
        fr.onload = async e => {
          const blob = e.target.result
          spark.append(blob)
          this.parts.push({ key: index + 1, blob })
          index++;                                       // next slice
          this.localProcessPercentage = parseFloat(((index / count) * 100).toFixed(2))

          if (index < count) _slice();                   // more slices? keep calm slice on
          else {
            this.sliceProcessing = false;
            this.fileMd5 = spark.end()
          }                              // we're done, use some callback
        };
        // todo handle errors
        fr.readAsArrayBuffer(slice);                     // convertt blob->ArrayBuffer
      }
      _slice()
    },
    success() {
      this.successFile = `${this.successFile}<div>${this.fileMd5}</div>`
    },
    async getFile() {
      // 确定按钮
      if (this.file === null) {
        this.$message('请先选择上传文件')
        return
      }
      if (this.sliceProcessing) {
        this.$message('正在处理视频分片，请稍等')
        return
      }

      const params = {
        fileName: this.file.name,
        fileMd5: this.fileMd5,
        chunkTotal: this.parts.length
      }
      const res = await findFile(params)

      // 全部切完以后 发一个请求给后端 拉当前文件后台存储的切片信息 用于检测有多少上传成功的切片
      const finishList = res.data.file.ExaFileChunk // 上传成功的切片
      const IsFinish = res.data.file.IsFinish // 是否是同文件不同命 （文件md5相同 文件名不同 则默认是同一个文件但是不同文件名 此时后台数据库只需要拷贝一下数据库文件即可 不需要上传文件 即秒传功能）
      if (!IsFinish) {
        // 当是断点续传时候
        this.waitUpLoad = this.parts.filter(all => {
          return !(
              finishList &&
              finishList.some(fi => fi.FileChunkNumber === all.key)
          ) // 找出需要上传的切片
        })
      } else {
        this.$message.warning('文件已经存在于服务器中！')
        this.success()
        this.init()
        this.$refs.Input.value = '';
        return
      }
      this.waitNum = this.waitUpLoad.length // 记录长度用于百分比展示

      this.percentage = Math.floor(((this.parts.length - this.waitNum) / this.parts.length) * 100)
      this.sliceFile() // 上传切片
    },
    sliceFile() {
      this.uploading = true
      this.failedCount = 0
      this.waitUpLoad &&
      this.waitUpLoad.map(({ blob, key }) => {
        const formData = new window.FormData() // 创建FormData用于存储传给后端的信息
        // 需要上传的切片
        formData.append('chunkTotal', this.parts.length) // 切片总数携带给后台 总有用的
        formData.append('file', new Blob([blob], {type: this.file.type})) // 当前的切片
        formData.append('chunkNumber', key) // 当前是第几片
        formData.append('fileMd5', this.fileMd5)
        formData.append('fileName', this.file.name)
        const spark = new SparkMD5.ArrayBuffer()
        spark.append(blob)
        formData.append('chunkMd5', spark.end()) // 获取当前切片md5 后端用于验证切片完整性
        this.upLoadFileSlice(formData)
      })
    },
    upLoadFileSlice(formData) {
      // 切片上传
      axios.post(process.env.VUE_APP_BASE_API + '/fileUploadAndDownload/breakpointContinue',
          formData, {
            headers: {
              'x-token': this.token,
              'Content-Type': 'multipart/form-data'
            }
          }).then(
          async ({ data: { code } }) => {
            console.log('code ', code)
            if (code !== 0) {
              this.failedCount++
            }
            this.waitNum-- // 百分数增加
            console.log('waitNumb', this.waitNum)
            this.percentage = Math.floor(((this.parts.length - this.waitNum) / this.parts.length) * 100)
            if (this.waitNum === 0) {
              if (this.failedCount > 0) {
                this.$message.error('部分切片上传失败，请重新点击上传');
                return
              }
              // 切片传完以后 合成文件
              const params = {
                fileName: this.file.name,
                fileMd5: this.fileMd5
              }

              const res = await breakpointContinueFinish(params)
              if (res.code === 0) {
                this.$refs.Input.value = '';
                this.$message.success('上传文件完成')
                // 合成文件过后 删除缓存切片
                const params = {
                  fileName: this.file.name,
                  fileMd5: this.fileMd5,
                  filePath: res.data.filePath
                }
                this.success()
                await removeChunk(params) // TODO 对象存储需要分片上传时，注释此行即可
              }

              this.init();
            }
          }
      )

    },
    inputChange() {
      if (this.sliceProcessing) {
        this.$message('正在处理视频分片，请稍等，放弃请刷新浏览器')
        return
      }

      this.$refs.Input.dispatchEvent(new MouseEvent('click'))
    }
  }
}
</script>

<style lang='scss' scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
#fromCont{
  display: inline-block;
}
.fileUpload{
    padding: 3px 10px;
    font-size: 12px;
    height: 20px;
    line-height: 20px;
    position: relative;
    cursor: pointer;
    color: #000;
    border: 1px solid #c1c1c1;
    border-radius: 4px;
    overflow: hidden;
    display: inline-block;
    input{
      position: absolute;
      font-size: 100px;
      right: 0;
      top: 0;
      opacity: 0;
      cursor: pointer;
    }
}
 .fileName{
    display: inline-block;
    vertical-align: top;
    margin: 6px 15px 0 15px;
  }
  .uploadBtn{
    position: relative;
    top: -10px;
    margin-left: 15px;
  }
  .tips{
    margin-top: 30px;
    font-size: 14px;
    font-weight: 400;
    color: #606266;
  }
  .el-divider{
    margin: 0 0 30px 0;
  }

 .list{
   margin-top:15px;
 }
 .list-item {
  display: block;
  margin-right: 10px;
  color: #606266;
  line-height: 25px;
  margin-bottom: 5px;
  width: 40%;
   .percentage{
          float: right;
        }
}
.list-enter-active, .list-leave-active {
  transition: all 1s;
}
.list-enter, .list-leave-to
/* .list-leave-active for below version 2.1.8 */ {
  opacity: 0;
  transform: translateY(-30px);
}
</style>
