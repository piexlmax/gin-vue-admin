<template>
  <div class="upload">
    <div class="btn-list">
      <el-upload
        class="excel-btn"
        :action="`${path}/excel/importExcel`"
        :headers="{'x-token':token}"
        :on-success="loadExcel"
        :show-file-list="false"
      >
        <el-button size="small" type="primary" icon="el-icon-upload2">导入</el-button>
      </el-upload>
      <el-button class="excel-btn" size="small" type="primary" icon="el-icon-download" @click="handleExcelExport('ExcelExport.xlsx')">导出</el-button>
      <el-button class="excel-btn" size="small" type="success" icon="el-icon-download" @click="downloadExcelTemplate()">下载模板</el-button>
    </div>
    <el-table :data="tableData" border row-key="ID" stripe>
      <el-table-column label="ID" min-width="100" prop="ID" />
      <el-table-column label="路由Name" min-width="160" prop="name" />
      <el-table-column label="路由Path" min-width="160" prop="path" />
      <el-table-column label="是否隐藏" min-width="100" prop="hidden">
        <template slot-scope="scope">
          <span>{{ scope.row.hidden?"隐藏":"显示" }}</span>
        </template>
      </el-table-column>
      <el-table-column label="父节点" min-width="90" prop="parentId" />
      <el-table-column label="排序" min-width="70" prop="sort" />
      <el-table-column label="文件路径" min-width="360" prop="component" />
    </el-table>
  </div>
</template>

<script>
const path = process.env.VUE_APP_BASE_API
import { mapGetters } from 'vuex'
import infoList from '@/mixins/infoList'
import { exportExcel, loadExcelData, downloadTemplate } from '@/api/excel'
import { getMenuList } from '@/api/menu'
export default {
  name: 'Excel',
  mixins: [infoList],
  data() {
    return {
      listApi: getMenuList,
      path: path
    }
  },
  computed: {
    ...mapGetters('user', ['userInfo', 'token'])
  },
  created() {
    this.pageSize = 999
    this.getTableData()
  },
  methods: {
    handleExcelExport(fileName) {
      if (!fileName || typeof fileName !== 'string') {
        fileName = 'ExcelExport.xlsx'
      }
      exportExcel(this.tableData, fileName)
    },
    loadExcel() {
      this.listApi = loadExcelData
      this.getTableData()
    },
    downloadExcelTemplate() {
      downloadTemplate('ExcelTemplate.xlsx')
    }
  }
}
</script>

<style lang="scss" scoped>
.btn-list{
  display: flex;
  margin-bottom: 12px;
  .excel-btn+.excel-btn{
    margin-left: 12px;
  }
}
</style>
