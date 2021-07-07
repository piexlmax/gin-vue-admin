<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="路径">
          <el-input v-model="searchInfo.path" placeholder="路径" />
        </el-form-item>
        <el-form-item label="描述">
          <el-input v-model="searchInfo.description" placeholder="描述" />
        </el-form-item>
        <el-form-item label="api组">
          <el-input v-model="searchInfo.apiGroup" placeholder="api组" />
        </el-form-item>
        <el-form-item label="请求">
          <el-select v-model="searchInfo.method" clearable placeholder="请选择">
            <el-option
              v-for="item in methodOptions"
              :key="item.value"
              :label="`${item.label}(${item.value})`"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" icon="el-icon-search" @click="onSubmit">查询</el-button>
          <el-button size="mini" type="primary" icon="el-icon-plus" @click="openDialog('addApi')">新增</el-button>
          <el-popover v-model="deleteVisible" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="deleteVisible = false">取消</el-button>
              <el-button size="mini" type="primary" @click="onDelete">确定</el-button>
            </div>
            <el-button slot="reference" icon="el-icon-delete" size="mini" type="danger" style="margin-left: 10px;">批量删除</el-button>
          </el-popover>
        </el-form-item>
      </el-form>
    </div>
    <el-table :data="tableData" border stripe @sort-change="sortChange" @selection-change="handleSelectionChange">
      <el-table-column
        type="selection"
        width="55"
      />
      <el-table-column label="id" min-width="60" prop="ID" sortable="custom" />
      <el-table-column label="api路径" min-width="150" prop="path" sortable="custom" />
      <el-table-column label="api分组" min-width="150" prop="apiGroup" sortable="custom" />
      <el-table-column label="api简介" min-width="150" prop="description" sortable="custom" />
      <el-table-column label="请求" min-width="150" prop="method" sortable="custom">
        <template slot-scope="scope">
          <div>
            {{ scope.row.method }}
            <el-tag
              :key="scope.row.methodFiletr"
              :type="scope.row.method|tagTypeFiletr"
              effect="dark"
              size="mini"
            >{{ scope.row.method|methodFiletr }}</el-tag>
            <!-- {{scope.row.method|methodFiletr}} -->
          </div>
        </template>
      </el-table-column>

      <el-table-column fixed="right" label="操作" width="200">
        <template slot-scope="scope">
          <el-button size="small" type="primary" icon="el-icon-edit" @click="editApi(scope.row)">编辑</el-button>
          <el-button
            size="small"
            type="danger"
            icon="el-icon-delete"
            @click="deleteApi(scope.row)"
          >删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="[10, 30, 50, 100]"
      :style="{float:'right',padding:'20px'}"
      :total="total"
      layout="total, sizes, prev, pager, next, jumper"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
    />

    <el-dialog :before-close="closeDialog" :title="dialogTitle" :visible.sync="dialogFormVisible">
      <el-form ref="apiForm" :inline="true" :model="form" :rules="rules" label-width="80px">
        <el-form-item label="路径" prop="path">
          <el-input v-model="form.path" autocomplete="off" />
        </el-form-item>
        <el-form-item label="请求" prop="method">
          <el-select v-model="form.method" placeholder="请选择">
            <el-option
              v-for="item in methodOptions"
              :key="item.value"
              :label="`${item.label}(${item.value})`"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="api分组" prop="apiGroup">
          <el-input v-model="form.apiGroup" autocomplete="off" />
        </el-form-item>
        <el-form-item label="api简介" prop="description">
          <el-input v-model="form.description" autocomplete="off" />
        </el-form-item>
      </el-form>
      <div class="warning">新增Api需要在角色管理内配置权限才可使用</div>
      <div slot="footer" class="dialog-footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button type="primary" @click="enterDialog">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
// 获取列表内容封装在mixins内部  getTableData方法 初始化已封装完成 条件搜索时候 请把条件安好后台定制的结构体字段 放到 this.searchInfo 中即可实现条件搜索

import {
  getApiById,
  getApiList,
  createApi,
  updateApi,
  deleteApi,
  deleteApisByIds
} from '@/api/api'
import infoList from '@/mixins/infoList'
import { toSQLLine } from '@/utils/stringFun'
const methodOptions = [
  {
    value: 'POST',
    label: '创建',
    type: 'success'
  },
  {
    value: 'GET',
    label: '查看',
    type: ''
  },
  {
    value: 'PUT',
    label: '更新',
    type: 'warning'
  },
  {
    value: 'DELETE',
    label: '删除',
    type: 'danger'
  }
]

export default {
  name: 'Api',
  filters: {
    methodFiletr(value) {
      const target = methodOptions.filter(item => item.value === value)[0]
      // return target && `${target.label}(${target.value})`
      return target && `${target.label}`
    },
    tagTypeFiletr(value) {
      const target = methodOptions.filter(item => item.value === value)[0]
      return target && `${target.type}`
    }
  },
  mixins: [infoList],
  data() {
    return {
      deleteVisible: false,
      listApi: getApiList,
      dialogFormVisible: false,
      dialogTitle: '新增Api',
      apis: [],
      form: {
        path: '',
        apiGroup: '',
        method: '',
        description: ''
      },
      methodOptions: methodOptions,
      type: '',
      rules: {
        path: [{ required: true, message: '请输入api路径', trigger: 'blur' }],
        apiGroup: [
          { required: true, message: '请输入组名称', trigger: 'blur' }
        ],
        method: [
          { required: true, message: '请选择请求方式', trigger: 'blur' }
        ],
        description: [
          { required: true, message: '请输入api介绍', trigger: 'blur' }
        ]
      }
    }
  },
  created() {
    this.getTableData()
  },
  methods: {
    //  选中api
    handleSelectionChange(val) {
      this.apis = val
    },
    async onDelete() {
      const ids = this.apis.map(item => item.ID)
      const res = await deleteApisByIds({ ids })
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: res.msg
        })
        if (this.tableData.length === ids.length && this.page > 1) {
          this.page--
        }
        this.deleteVisible = false
        this.getTableData()
      }
    },
    // 排序
    sortChange({ prop, order }) {
      if (prop) {
        this.searchInfo.orderKey = toSQLLine(prop)
        this.searchInfo.desc = order === 'descending'
      }
      this.getTableData()
    },
    // 条件搜索前端看此方法
    onSubmit() {
      this.page = 1
      this.pageSize = 10
      this.getTableData()
    },
    initForm() {
      this.$refs.apiForm.resetFields()
      this.form = {
        path: '',
        apiGroup: '',
        method: '',
        description: ''
      }
    },
    closeDialog() {
      this.initForm()
      this.dialogFormVisible = false
    },
    openDialog(type) {
      switch (type) {
        case 'addApi':
          this.dialogTitle = '新增Api'
          break
        case 'edit':
          this.dialogTitle = '编辑Api'
          break
        default:
          break
      }
      this.type = type
      this.dialogFormVisible = true
    },
    async editApi(row) {
      const res = await getApiById({ id: row.ID })
      this.form = res.data.api
      this.openDialog('edit')
    },
    async deleteApi(row) {
      this.$confirm('此操作将永久删除所有角色下该api, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(async() => {
          const res = await deleteApi(row)
          if (res.code === 0) {
            this.$message({
              type: 'success',
              message: '删除成功!'
            })
            if (this.tableData.length === 1 && this.page > 1) {
              this.page--
            }
            this.getTableData()
          }
        })
        .catch(() => {
          this.$message({
            type: 'info',
            message: '已取消删除'
          })
        })
    },
    async enterDialog() {
      this.$refs.apiForm.validate(async valid => {
        if (valid) {
          switch (this.type) {
            case 'addApi':
              {
                const res = await createApi(this.form)
                if (res.code === 0) {
                  this.$message({
                    type: 'success',
                    message: '添加成功',
                    showClose: true
                  })
                }
                this.getTableData()
                this.closeDialog()
              }

              break
            case 'edit':
              {
                const res = await updateApi(this.form)
                if (res.code === 0) {
                  this.$message({
                    type: 'success',
                    message: '编辑成功',
                    showClose: true
                  })
                }
                this.getTableData()
                this.closeDialog()
              }
              break
            default:
              // eslint-disable-next-line no-lone-blocks
              {
                this.$message({
                  type: 'error',
                  message: '未知操作',
                  showClose: true
                })
              }
              break
          }
        }
      })
    }
  }
}
</script>

<style scoped lang="scss">
.button-box {
  padding: 10px 20px;
  .el-button {
    float: right;
  }
}
.el-tag--mini {
  margin-left: 5px;
}
.warning {
  color: #dc143c;
}
</style>
