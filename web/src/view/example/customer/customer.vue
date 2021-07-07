<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item>
          <el-button size="mini" type="primary" icon="el-icon-search" @click="openDialog">新增</el-button>
        </el-form-item>
      </el-form>
    </div>
    <el-table
      ref="multipleTable"
      :data="tableData"
      border
      stripe
      style="width: 100%"
      tooltip-effect="dark"
    >
      <el-table-column type="selection" width="55" />
      <el-table-column label="接入日期" width="180">
        <template slot-scope="scope">{{ scope.row.CreatedAt|formatDate }}</template>
      </el-table-column>
      <el-table-column label="姓名" prop="customerName" width="120" />
      <el-table-column label="电话" prop="customerPhoneData" width="120" />
      <el-table-column label="接入人ID" prop="sysUserId" width="120" />
      <el-table-column label="按钮组" min-width="160">
        <template slot-scope="scope">
          <el-button size="small" type="text" @click="updateCustomer(scope.row)">变更</el-button>
          <el-popover v-model="scope.row.visible" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
              <el-button type="primary" size="mini" @click="deleteCustomer(scope.row)">确定</el-button>
            </div>
            <el-button slot="reference" type="danger" icon="el-icon-delete" size="mini">删除</el-button>
          </el-popover>
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

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="客户">
      <el-form :inline="true" :model="form" label-width="80px">
        <el-form-item label="客户名">
          <el-input v-model="form.customerName" autocomplete="off" />
        </el-form-item>
        <el-form-item label="客户电话">
          <el-input v-model="form.customerPhoneData" autocomplete="off" />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button type="primary" @click="enterDialog">确 定</el-button>
      </div>
    </el-dialog>
    <div class="tips">在资源权限中将此角色的资源权限清空 或者不包含创建者的角色 即可屏蔽此客户资源的显示</div>
  </div>
</template>

<script>
import {
  createExaCustomer,
  updateExaCustomer,
  deleteExaCustomer,
  getExaCustomer,
  getExaCustomerList
} from '@/api/customer'
import { formatTimeToStr } from '@/utils/date'
import infoList from '@/mixins/infoList'

export default {
  name: 'Customer',
  filters: {
    formatDate: function(time) {
      if (time !== null && time !== '') {
        var date = new Date(time)
        return formatTimeToStr(date, 'yyyy-MM-dd hh:mm:ss')
      } else {
        return ''
      }
    }
  },
  mixins: [infoList],
  data() {
    return {
      listApi: getExaCustomerList,
      dialogFormVisible: false,
      type: '',
      form: {
        customerName: '',
        customerPhoneData: ''
      }
    }
  },
  created() {
    this.getTableData()
  },
  methods: {
    async updateCustomer(row) {
      const res = await getExaCustomer({ ID: row.ID })
      this.type = 'update'
      if (res.code === 0) {
        this.form = res.data.customer
        this.dialogFormVisible = true
      }
    },
    closeDialog() {
      this.dialogFormVisible = false
      this.form = {
        customerName: '',
        customerPhoneData: ''
      }
    },
    async deleteCustomer(row) {
      row.visible = false
      const res = await deleteExaCustomer({ ID: row.ID })
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: ''
        })
        if (this.tableData.length === 1 && this.page > 1) {
          this.page--
        }
        this.getTableData()
      }
    },
    async enterDialog() {
      let res
      switch (this.type) {
        case 'create':
          res = await createExaCustomer(this.form)
          break
        case 'update':
          res = await updateExaCustomer(this.form)
          break
        default:
          res = await createExaCustomer(this.form)
          break
      }

      if (res.code === 0) {
        this.closeDialog()
        this.getTableData()
      }
    },
    openDialog() {
      this.type = 'create'
      this.dialogFormVisible = true
    }
  }
}
</script>

<style>
</style>
