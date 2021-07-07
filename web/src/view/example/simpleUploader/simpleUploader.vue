<template>
  <uploader
    :options="options"
    :file-status-text="statusText"
    :auto-start="false"
    class="uploader-example"
    @file-added="fileAdded"
    @file-progress="onFileProgress"
    @file-success="onFileSuccess"
    @file-error="onFileError"
  >
    <uploader-unsupport />
    <uploader-drop>
      <p>拖拽文件至此或点击</p>
      <uploader-btn>选择文件</uploader-btn>
    </uploader-drop>
    <uploader-list />
  </uploader>
</template>

<script>
var notUploadedChunks = [] // 已经上传过的文件chunkNumber数组
var isUploaded = false // 文件已经上传成功了
import { mapGetters } from 'vuex'
import { checkFileMd5, mergeFileMd5 } from '@/api/simpleUploader'
import SparkMD5 from 'spark-md5'
const path = process.env.VUE_APP_BASE_API
export default {
  name: 'SimpleUploader',
  data() {
    return {
      md5: ''
    }
  },
  computed: {
    ...mapGetters('user', ['userInfo', 'token']),
    statusText() {
      return {
        success: '成功了',
        error: '出错了',
        uploading: '上传中',
        paused: '暂停中',
        waiting: '等待中'
      }
    },
    options() {
      return {
        target: path + '/simpleUploader/upload',
        testChunks: false,
        simultaneousUploads: 5,
        chunkSize: 2 * 1024 * 1024,
        headers: {
          'x-token': this.token,
          'x-user-id': this.userInfo.ID
        },
        checkChunkUploadedByResponse(chunk) {
          if (isUploaded) {
            return true // return true 会忽略当前文件，不会再发送给后台
          } else {
            // 根据已经上传过的切片来进行忽略
            return (
              notUploadedChunks &&
                notUploadedChunks.some(
                  item => item.chunkNumber === chunk.offset + 1
                )
            )
          }
        }
      }
    }
  },
  methods: {

    // 上传单个文件
    fileAdded(file) {
      this.computeMD5(file) // 生成MD5
    },
    // 计算MD5值
    computeMD5(file) {
      var that = this
      isUploaded = false // 这个文件是否已经上传成功过
      notUploadedChunks = [] // 未成功的chunkNumber
      var fileReader = new FileReader()
      var md5 = ''

      file.pause()

      fileReader.readAsArrayBuffer(file.file)
      fileReader.onload = async function(e) {
        if (file.size !== e.target.result.byteLength) {
          this.error(
            'Browser reported success but could not read the file until the end.'
          )
          return false
        }
        md5 = SparkMD5.ArrayBuffer.hash(e.target.result, false)

        file.uniqueIdentifier = md5
        if (md5 !== '') {
          const res = await checkFileMd5({ md5: md5 })
          if (res.code === 0) {
            if (res.data.isDone) {
              // 上传成功过
              isUploaded = true
              that.$message({
                message: '该文件已经上传成功过了，秒传成功。',
                type: 'success'
              })

              file.cancel()
            } else {
              isUploaded = false
              notUploadedChunks = res.data.chunks
              if (notUploadedChunks.length) {
                file.resume()
              }
            }
          }
        }
      }
      fileReader.onerror = function() {
        this.error(
          'generater md5 时FileReader异步读取文件出错了，FileReader onerror was triggered, maybe the browser aborted due to high memory usage.'
        )
        return false
      }
    },
    // 上传进度
    onFileProgress() {},
    // 上传成功
    async onFileSuccess(rootFile, file) {
      await mergeFileMd5({ md5: file.uniqueIdentifier, fileName: file.name })
    },
    onFileError(rootFile, file, response) {
      this.$message({
        message: response,
        type: 'error'
      })
    }
  }
}
</script>

<style>
.uploader-example {
  width: 880px;
  padding: 15px;
  margin: 15px 15px 20px;
  font-size: 12px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.4);
}
.uploader-example .uploader-btn {
  margin-right: 4px;
}
.uploader-example .uploader-list {
  max-height: 440px;
  overflow: auto;
  overflow-x: hidden;
  overflow-y: auto;
}
</style>
