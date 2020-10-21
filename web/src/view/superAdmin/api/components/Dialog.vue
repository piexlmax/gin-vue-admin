<template>
  <div>
    <el-dialog :before-close="closeDialog" :title="dialogTitle" :visible.sync="dialogFormVisible">
      <el-form :inline="true" :model="form" :rules="rules" label-width="80px" ref="apiForm">
        <el-form-item label="路径" prop="path">
          <el-input autocomplete="off" v-model="form.path"></el-input>
        </el-form-item>
        <el-form-item label="请求" prop="method">
          <el-select placeholder="请选择" v-model="form.method">
            <el-option
              :key="item.value"
              :label="`${item.label}(${item.value})`"
              :value="item.value"
              v-for="item in methodOptions"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="api分组" prop="apiGroup">
          <el-input autocomplete="off" v-model="form.apiGroup"></el-input>
        </el-form-item>
        <el-form-item label="api简介" prop="description">
          <el-input autocomplete="off" v-model="form.description"></el-input>
        </el-form-item>
      </el-form>
      <div class="warning">新增Api需要在角色管理内配置权限才可使用</div>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { getApiById, createApi, updateApi } from '@/api/api'
import config from './config'

export default {
  data() {
    return {
      methodOptions: config.methodOptions,
      dialogFormVisible: false,
      dialogTitle: '新增Api',
      type: '',
      form: {
        path: '',
        apiGroup: '',
        method: '',
        description: '',
      },
      rules: {
        path: [{ required: true, message: '请输入api路径', trigger: 'blur' }],
        apiGroup: [{ required: true, message: '请输入组名称', trigger: 'blur' }],
        method: [{ required: true, message: '请选择请求方式', trigger: 'blur' }],
        description: [{ required: true, message: '请输入api介绍', trigger: 'blur' }],
      },
    }
  },
  methods: {
    closeDialog() {
      this.initForm()
      this.dialogFormVisible = false
    },
    async openDialog(type, row) {
      let res
      switch (type) {
        case 'addApi':
          this.dialogTitlethis = '新增Api'
          break
        case 'edit':
          res = await getApiById({ id: row.ID })
          this.form = res.data.api
          this.dialogTitlethis = '编辑Api'
          break
        default:
          break
      }
      this.type = type
      this.dialogFormVisible = true
    },
    initForm() {
      this.$refs.apiForm.resetFields()
      this.form = {
        path: '',
        apiGroup: '',
        method: '',
        description: '',
      }
    },
    async enterDialog() {
      this.$refs.apiForm.validate(async valid => {
        if (valid) {
          switch (this.type) {
            case 'addApi':
              {
                const res = await createApi(this.form)
                if (res.code == 0) {
                  this.$message({
                    type: 'success',
                    message: '添加成功',
                    showClose: true,
                  })
                }
                this.$emit('success')
                this.closeDialog()
              }

              break
            case 'edit':
              {
                const res = await updateApi(this.form)
                if (res.code == 0) {
                  this.$message({
                    type: 'success',
                    message: '编辑成功',
                    showClose: true,
                  })
                }
                this.$emit('success')
                this.closeDialog()
              }
              break
            default:
              {
                this.$message({
                  type: 'error',
                  message: '未知操作',
                  showClose: true,
                })
              }
              break
          }
        }
      })
    },
  },
}
</script>
