<template>
  <div>
    <!-- 从数据库直接获取字段 -->
    <el-collapse v-model="activeNames">
      <el-collapse-item name="1">
        <template slot="title">
          <div :style="{fontSize:'16px',paddingLeft:'20px'}">
            点这里从现有数据库创建代码
            <i class="header-icon el-icon-thumb" />
          </div>
        </template>
        <el-form ref="getTableForm" :inline="true" :model="dbform" label-width="120px">
          <el-form-item label="数据库名" prop="structName">
            <el-select v-model="dbform.dbName" filterable placeholder="请选择数据库" @change="getTable">
              <el-option
                v-for="item in dbOptions"
                :key="item.database"
                :label="item.database"
                :value="item.database"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="表名" prop="structName">
            <el-select
              v-model="dbform.tableName"
              :disabled="!dbform.dbName"
              filterable
              placeholder="请选择表"
            >
              <el-option
                v-for="item in tableOptions"
                :key="item.tableName"
                :label="item.tableName"
                :value="item.tableName"
              />
            </el-select>
          </el-form-item>
          <el-form-item>
            <el-button size="mini" type="primary" @click="getColumn">使用此表创建</el-button>
          </el-form-item>
        </el-form>
      </el-collapse-item>
    </el-collapse>

    <el-divider />
    <!-- 初始版本自动化代码工具 -->
    <el-form ref="autoCodeForm" :rules="rules" :model="form" label-width="120px" :inline="true">
      <el-form-item label="Struct名称" prop="structName">
        <el-input v-model="form.structName" placeholder="首字母自动转换大写" />
      </el-form-item>
      <el-form-item label="tableName" prop="tableName">
        <el-input v-model="form.tableName" placeholder="指定表名（非必填）" />
      </el-form-item>
      <el-form-item label="Struct简称" prop="abbreviation">
        <el-input v-model="form.abbreviation" placeholder="简称会作为入参对象名和路由group" />
      </el-form-item>
      <el-form-item label="Struct中文名称" prop="description">
        <el-input v-model="form.description" placeholder="中文描述作为自动api描述" />
      </el-form-item>
      <el-form-item label="文件名称" prop="packageName">
        <el-input v-model="form.packageName" placeholder="生成文件的默认名称" />
      </el-form-item>
      <el-form-item label="自动创建api">
        <el-checkbox v-model="form.autoCreateApiToSql" />
      </el-form-item>
      <el-form-item label="自动移动文件">
        <el-checkbox v-model="form.autoMoveFile" />
      </el-form-item>
    </el-form>
    <!-- 组件列表 -->
    <div class="button-box clearflex">
      <el-button size="mini" type="primary" @click="editAndAddField()">新增Field</el-button>
    </div>
    <el-table :data="form.fields" border stripe>
      <el-table-column type="index" label="序列" width="100" />
      <el-table-column prop="fieldName" label="Field名" />
      <el-table-column prop="fieldDesc" label="中文名" />
      <el-table-column prop="fieldJson" label="FieldJson" />
      <el-table-column prop="fieldType" label="Field数据类型" width="130" />
      <el-table-column prop="dataType" label="数据库字段类型" width="130" />
      <el-table-column prop="dataTypeLong" label="数据库字段长度" width="130" />
      <el-table-column prop="columnName" label="数据库字段" width="130" />
      <el-table-column prop="comment" label="数据库字段描述" width="130" />
      <el-table-column prop="fieldSearchType" label="搜索条件" width="130" />
      <el-table-column prop="dictType" label="字典" width="130" />
      <el-table-column label="操作" width="300">
        <template slot-scope="scope">
          <el-button
            size="mini"
            type="primary"
            icon="el-icon-edit"
            @click="editAndAddField(scope.row)"
          >编辑</el-button>
          <el-button
            size="mini"
            type="text"
            :disabled="scope.$index === 0"
            @click="moveUpField(scope.$index)"
          >上移</el-button>
          <el-button
            size="mini"
            type="text"
            :disabled="(scope.$index + 1) === form.fields.length"
            @click="moveDownField(scope.$index)"
          >下移</el-button>
          <el-popover v-model="scope.row.visible" placement="top">
            <p>确定删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
              <el-button type="primary" size="mini" @click="deleteField(scope.$index)">确定</el-button>
            </div>
            <el-button slot="reference" size="mini" type="danger" icon="el-icon-delete">删除</el-button>
          </el-popover>
        </template>
      </el-table-column>
    </el-table>
    <el-tag type="danger">id , created_at , updated_at , deleted_at 会自动生成请勿重复创建</el-tag>
    <!-- 组件列表 -->
    <div class="button-box clearflex">
      <el-button size="mini" type="primary" @click="enterForm(true)">预览代码</el-button>
      <el-button size="mini" type="primary" @click="enterForm(false)">生成代码</el-button>
    </div>
    <!-- 组件弹窗 -->
    <el-dialog title="组件内容" :visible.sync="dialogFlag">
      <FieldDialog v-if="dialogFlag" ref="fieldDialog" :dialog-middle="dialogMiddle" />
      <div slot="footer" class="dialog-footer">
        <el-button size="mini" @click="closeDialog">取 消</el-button>
        <el-button size="mini" type="primary" @click="enterDialog">确 定</el-button>
      </div>
    </el-dialog>

    <el-dialog :visible.sync="previewFlag">
      <PreviewCodeDialg v-if="previewFlag" :preview-code="preViewCode" />
      <div slot="footer" class="dialog-footer">
        <el-button type="primary" @click="previewFlag = false">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
const fieldTemplate = {
  fieldName: '',
  fieldDesc: '',
  fieldType: '',
  dataType: '',
  fieldJson: '',
  columnName: '',
  dataTypeLong: '',
  comment: '',
  fieldSearchType: '',
  dictType: ''
}

import FieldDialog from '@/view/systemTools/autoCode/component/fieldDialog.vue'
import PreviewCodeDialg from '@/view/systemTools/autoCode/component/previewCodeDialg.vue'
import { toUpperCase, toHump, toSQLLine } from '@/utils/stringFun'
import { createTemp, getDB, getTable, getColumn, preview } from '@/api/autoCode'
import { getDict } from '@/utils/dictionary'

export default {
  name: 'AutoCode',
  components: {
    FieldDialog,
    PreviewCodeDialg
  },
  data() {
    return {
      activeNames: [''],
      preViewCode: {},
      dbform: {
        dbName: '',
        tableName: ''
      },
      dbOptions: [],
      tableOptions: [],
      addFlag: '',
      fdMap: {},
      form: {
        structName: '',
        tableName: '',
        packageName: '',
        abbreviation: '',
        description: '',
        autoCreateApiToSql: false,
        autoMoveFile: false,
        fields: []
      },
      rules: {
        structName: [
          { required: true, message: '请输入结构体名称', trigger: 'blur' }
        ],
        abbreviation: [
          { required: true, message: '请输入结构体简称', trigger: 'blur' }
        ],
        description: [
          { required: true, message: '请输入结构体描述', trigger: 'blur' }
        ],
        packageName: [
          {
            required: true,
            message: '文件名称：sys_xxxx_xxxx',
            trigger: 'blur'
          }
        ]
      },
      dialogMiddle: {},
      bk: {},
      dialogFlag: false,
      previewFlag: false
    }
  },
  created() {
    this.getDb()
    this.setFdMap()
  },
  methods: {
    editAndAddField(item) {
      this.dialogFlag = true
      if (item) {
        this.addFlag = 'edit'
        this.bk = JSON.parse(JSON.stringify(item))
        this.dialogMiddle = item
      } else {
        this.addFlag = 'add'
        this.dialogMiddle = JSON.parse(JSON.stringify(fieldTemplate))
      }
    },
    moveUpField(index) {
      if (index === 0) {
        return
      }
      const oldUpField = this.form.fields[index - 1]
      this.form.fields.splice(index - 1, 1)
      this.form.fields.splice(index, 0, oldUpField)
    },
    moveDownField(index) {
      const fCount = this.form.fields.length
      if (index === fCount - 1) {
        return
      }
      const oldDownField = this.form.fields[index + 1]
      this.form.fields.splice(index + 1, 1)
      this.form.fields.splice(index, 0, oldDownField)
    },
    enterDialog() {
      this.$refs.fieldDialog.$refs.fieldDialogFrom.validate(valid => {
        if (valid) {
          this.dialogMiddle.fieldName = toUpperCase(
            this.dialogMiddle.fieldName
          )
          if (this.addFlag === 'add') {
            this.form.fields.push(this.dialogMiddle)
          }
          this.dialogFlag = false
        } else {
          return false
        }
      })
    },
    closeDialog() {
      if (this.addFlag === 'edit') {
        this.dialogMiddle = this.bk
      }
      this.dialogFlag = false
    },
    deleteField(index) {
      this.form.fields.splice(index, 1)
    },
    async enterForm(isPreview) {
      if (this.form.fields.length <= 0) {
        this.$message({
          type: 'error',
          message: '请填写至少一个field'
        })
        return false
      }
      if (
        this.form.fields.some(item => item.fieldName === this.form.structName)
      ) {
        this.$message({
          type: 'error',
          message: '存在与结构体同名的字段'
        })
        return false
      }
      this.$refs.autoCodeForm.validate(async valid => {
        if (valid) {
          this.form.structName = toUpperCase(this.form.structName)
          if (this.form.structName === this.form.abbreviation) {
            this.$message({
              type: 'error',
              message: 'structName和struct简称不能相同'
            })
            return false
          }
          this.form.humpPackageName = toSQLLine(this.form.packageName)
          debugger
          if (isPreview) {
            const data = await preview(this.form)
            this.preViewCode = data.data.autoCode
            this.previewFlag = true
          } else {
            const data = await createTemp(this.form)
            if (data.headers?.success === 'false') {
              return
            } else {
              this.$message({
                type: 'success',
                message: '自动化代码创建成功，正在下载'
              })
            }
            const blob = new Blob([data])
            const fileName = 'ginvueadmin.zip'
            if ('download' in document.createElement('a')) {
              // 不是IE浏览器
              const url = window.URL.createObjectURL(blob)
              const link = document.createElement('a')
              link.style.display = 'none'
              link.href = url
              link.setAttribute('download', fileName)
              document.body.appendChild(link)
              link.click()
              document.body.removeChild(link) // 下载完成移除元素
              window.URL.revokeObjectURL(url) // 释放掉blob对象
            } else {
              // IE 10+
              window.navigator.msSaveBlob(blob, fileName)
            }
          }
        } else {
          return false
        }
      })
    },
    async getDb() {
      const res = await getDB()
      if (res.code === 0) {
        this.dbOptions = res.data.dbs
      }
    },
    async getTable() {
      const res = await getTable({ dbName: this.dbform.dbName })
      if (res.code === 0) {
        this.tableOptions = res.data.tables
      }
      this.dbform.tableName = ''
    },
    async getColumn() {
      const gormModelList = ['id', 'created_at', 'updated_at', 'deleted_at']
      const res = await getColumn(this.dbform)
      if (res.code === 0) {
        const tbHump = toHump(this.dbform.tableName)
        this.form.structName = toUpperCase(tbHump)
        this.form.tableName = this.dbform.tableName
        this.form.packageName = tbHump
        this.form.abbreviation = tbHump
        this.form.description = tbHump + '表'
        this.form.autoCreateApiToSql = true
        this.form.fields = []
        res.data.columns &&
          res.data.columns.map(item => {
            if (!gormModelList.some(gormfd => gormfd === item.columnName)) {
              const fbHump = toHump(item.columnName)
              this.form.fields.push({
                fieldName: toUpperCase(fbHump),
                fieldDesc: item.columnComment || fbHump + '字段',
                fieldType: this.fdMap[item.dataType],
                dataType: item.dataType,
                fieldJson: fbHump,
                dataTypeLong: item.dataTypeLong,
                columnName: item.columnName,
                comment: item.columnComment,
                fieldSearchType: '',
                dictType: ''
              })
            }
          })
      }
    },
    async setFdMap() {
      const fdTypes = ['string', 'int', 'bool', 'float64', 'time.Time']
      fdTypes.map(async fdtype => {
        const res = await getDict(fdtype)
        res && res.map(item => {
          this.fdMap[item.label] = fdtype
        })
      })
    }
  }
}
</script>

<style scoped lang="scss">
.button-box {
  padding: 10px 20px;
  .el-button {
    margin-right: 20px;
    float: right;
  }
}
</style>
