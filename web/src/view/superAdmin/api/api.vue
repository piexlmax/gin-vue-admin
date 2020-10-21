<template>
  <div>
    <search-form @change="handleFormChange" @add="$refs.dialog.openDialog('addApi')" />

    <pagination-table ref="table" :dataSource="getTableData">
      <el-table slot-scope="data" :data="data.tableData" @sort-change="sortChange" border stripe>
        <el-table-column label="id" min-width="60" prop="ID" sortable="custom"></el-table-column>
        <el-table-column label="api路径" min-width="150" prop="path" sortable="custom"></el-table-column>
        <el-table-column label="api分组" min-width="150" prop="apiGroup" sortable="custom"></el-table-column>
        <el-table-column label="api简介" min-width="150" prop="description" sortable="custom"></el-table-column>
        <el-table-column label="请求" min-width="150" prop="method" sortable="custom">
          <template slot-scope="scope">
            <div>
              {{ scope.row.method }}
              <el-tag :key="scope.row.methodFiletr" :type="scope.row.method | tagTypeFiletr" effect="dark" size="mini">
                {{ scope.row.method | methodFiletr }}
              </el-tag>
            </div>
          </template>
        </el-table-column>
        <el-table-column fixed="right" label="操作" width="200">
          <template slot-scope="scope">
            <el-button
              @click="$refs.dialog.openDialog('edit', scope.row)"
              size="small"
              type="primary"
              icon="el-icon-edit"
              >编辑</el-button
            >
            <el-button @click="deleteApi(scope.row)" size="small" type="danger" icon="el-icon-delete">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </pagination-table>

    <Dialog ref="dialog" @success="$refs.table.getTableData()" />
  </div>
</template>

<script>
// 获取列表内容封装在mixins内部  getTableData方法 初始化已封装完成 条件搜索时候 请把条件安好后台定制的结构体字段 放到 this.searchInfo 中即可实现条件搜索

import { getApiList, deleteApi } from '@/api/api'
import { toSQLLine } from '@/utils/stringFun'

import PaginationTable from '@/components/PaginationTable'
import SearchForm from './components/SearchForm'
import Dialog from './components/Dialog'

import config from './components/config'

export default {
  name: 'Api',
  components: {
    PaginationTable,
    SearchForm,
    Dialog,
  },
  data() {
    return {
      methodOptions: config.methodOptions,
      searchInfo: {},
    }
  },
  filters: {
    methodFiletr(value) {
      const target = config.methodOptions.filter(item => item.value === value)[0]
      return target && `${target.label}`
    },
    tagTypeFiletr(value) {
      const target = config.methodOptions.filter(item => item.value === value)[0]
      return target && `${target.type}`
    },
  },
  methods: {
    // 通过此方法将搜索参数传入组件
    getTableData(params) {
      return getApiList({ ...params, ...this.searchInfo })
    },
    //  搜索条件变换后执行
    handleFormChange(searchInfo) {
      this.searchInfo = searchInfo
      this.$refs.table.getTableData()
    },
    // 排序
    sortChange({ prop, order }) {
      if (prop) {
        this.searchInfo.orderKey = toSQLLine(prop)
        this.searchInfo.desc = order == 'descending'
      }
      this.$refs.table.getTableData()
    },
    // 删除
    async deleteApi(row) {
      this.$confirm('此操作将永久删除所有角色下该菜单, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
      })
        .then(async () => {
          const res = await deleteApi(row)
          if (res.code == 0) {
            this.$message({
              type: 'success',
              message: '删除成功!',
            })
            this.$refs.table.getTableData()
          }
        })
        .catch(() => {
          this.$message({
            type: 'info',
            message: '已取消删除',
          })
        })
    },
  },
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
