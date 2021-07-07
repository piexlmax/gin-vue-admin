<template>
  <div class="hello">
    <el-divider content-position="left">大文件上传</el-divider>
    <form id="fromCont" method="post">
      <div class="fileUpload" @click="inputChange">
        选择文件
        <input v-show="false" id="file" ref="Input" multiple="multiple" type="file" @change="choseFile">
      </div>
    </form>
    <el-button :disabled="limitFileSize" type="primary" size="mini" class="uploadBtn" @click="getFile">上传文件</el-button>
    <div class="el-upload__tip">请上传不超过5MB的文件</div>
    <div class="list">
      <transition name="list" tag="p">
        <div v-if="file" class="list-item">
          <i class="el-icon-document" />
          <span>{{ file.name }}</span>
          <span class="percentage">{{ percentage }}%</span>
          <el-progress :show-text="false" :text-inside="false" :stroke-width="2" :percentage="percentage" />
        </div>
      </transition>
    </div>
    <!-- <span
      v-if="this.file"
    >{{Math.floor(((this.formDataList.length-this.waitNum)/this.formDataList.length)*100)}}%</span> -->
    <div class="tips">此版本为先行体验功能测试版，样式美化和性能优化正在进行中，上传切片文件和合成的完整文件分别再QMPlusserver目录的breakpointDir文件夹和fileDir文件夹</div>
  </div>
</template>

<script>
import SparkMD5 from 'spark-md5'
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
      formDataList: [],
      waitUpLoad: [],
      waitNum: 0,
      limitFileSize: false,
      percentage: 0,
      percentageFlage: true,
      customColor: '#409eff'
    }
  },
  methods: {
    // 选中文件的函数
    async choseFile(e) {
      const fileR = new FileReader() // 创建一个reader用来读取文件流
      const file = e.target.files[0] // 获取当前文件
      const maxSize = 5 * 1024 * 1024
      this.file = file // file 丢全局方便后面用 可以改进为func传参形式
      this.percentage = 0
      if (file.size < maxSize) {
        fileR.readAsArrayBuffer(file) // 把文件读成ArrayBuffer  主要为了保持跟后端的流一致
        fileR.onload = async e => {
          // 读成arrayBuffer的回调 e 为方法自带参数 相当于 dom的e 流存在e.target.result 中
          const blob = e.target.result
          const spark = new SparkMD5.ArrayBuffer() // 创建md5制造工具 （md5用于检测文件一致性 这里不懂就打电话问我）
          spark.append(blob) // 文件流丢进工具
          this.fileMd5 = spark.end() // 工具结束 产生一个a 总文件的md5
          const FileSliceCap = 1 * 1024 * 1024 // 分片字节数
          let start = 0 // 定义分片开始切的地方
          let end = 0 // 每片结束切的地方a
          let i = 0 // 第几片
          this.formDataList = [] // 分片存储的一个池子 丢全局
          while (end < file.size) {
            // 当结尾数字大于文件总size的时候 结束切片
            start = i * FileSliceCap // 计算每片开始位置
            end = (i + 1) * FileSliceCap // 计算每片结束位置
            var fileSlice = this.file.slice(start, end) // 开始切  file.slice 为 h5方法 对文件切片 参数为 起止字节数
            const formData = new window.FormData() // 创建FormData用于存储传给后端的信息
            formData.append('fileMd5', this.fileMd5) // 存储总文件的Md5 让后端知道自己是谁的切片
            formData.append('file', fileSlice) // 当前的切片
            formData.append('chunkNumber', i) // 当前是第几片
            formData.append('fileName', this.file.name) // 当前文件的文件名 用于后端文件切片的命名  formData.appen 为 formData对象添加参数的方法
            this.formDataList.push({ key: i, formData }) // 把当前切片信息 自己是第几片 存入我们方才准备好的池子
            i++
          }
          const params = {
            fileName: this.file.name,
            fileMd5: this.fileMd5,
            chunkTotal: this.formDataList.length
          }
          const res = await findFile(params)
          // 全部切完以后 发一个请求给后端 拉当前文件后台存储的切片信息 用于检测有多少上传成功的切片
          const finishList = res.data.file.ExaFileChunk // 上传成功的切片
          const IsFinish = res.data.file.IsFinish // 是否是同文件不同命 （文件md5相同 文件名不同 则默认是同一个文件但是不同文件名 此时后台数据库只需要拷贝一下数据库文件即可 不需要上传文件 即秒传功能）
          if (!IsFinish) {
            // 当是断点续传时候
            this.waitUpLoad = this.formDataList.filter(all => {
              return !(
                finishList &&
              finishList.some(fi => fi.FileChunkNumber === all.key)
              ) // 找出需要上传的切片
            })
          } else {
            this.waitUpLoad = [] // 秒传则没有需要上传的切片
          }
          this.waitNum = this.waitUpLoad.length // 记录长度用于百分比展示
        }
      } else {
        this.limitFileSize = true
        this.$message('请上传小于5M文件')
      }
    },
    getFile() {
      // 确定按钮
      if (this.file === null) {
        this.$message('请先上传文件')
        return
      }
      this.percentage = Math.floor(((this.formDataList.length - this.waitNum) / this.formDataList.length) * 100)
      if (this.percentage === 100) {
        this.percentageFlage = false
      }
      this.sliceFile() // 上传切片
    },
    sliceFile() {
      this.waitUpLoad &&
        this.waitUpLoad.map(item => {
          // 需要上传的切片
          item.formData.append('chunkTotal', this.formDataList.length) // 切片总数携带给后台 总有用的
          const fileR = new FileReader() // 功能同上
          const file = item.formData.get('file')
          fileR.readAsArrayBuffer(file)
          fileR.onload = e => {
            const spark = new SparkMD5.ArrayBuffer()
            spark.append(e.target.result)
            item.formData.append('chunkMd5', spark.end()) // 获取当前切片md5 后端用于验证切片完整性
            this.upLoadFileSlice(item)
          }
        })
    },
    async upLoadFileSlice(item) {
      // 切片上传
      await axios.post(process.env.VUE_APP_BASE_API + '/fileUploadAndDownload/breakpointContinue', item.formData)
      this.waitNum-- // 百分数增加
      if (this.waitNum === 0) {
        // 切片传完以后 合成文件
        const params = {
          fileName: this.file.name,
          fileMd5: this.fileMd5
        }
        const res = await breakpointContinueFinish(params)
        if (res.success) {
          // 合成文件过后 删除缓存切片
          const params = {
            fileName: this.file.name,
            fileMd5: this.fileMd5,
            filePath: res.data.filePath
          }
          await removeChunk(params)
        }
      }
    },
    inputChange() {
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
