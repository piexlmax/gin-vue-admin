<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="请求方法">
          <el-input v-model="searchInfo.method" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="请求路径">
          <el-input v-model="searchInfo.path" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="结果状态码">
          <el-input v-model="searchInfo.status" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" icon="el-icon-search" @click="onSubmit">查询</el-button>
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
    <el-table
      ref="multipleTable"
      :data="tableData"
      border
      stripe
      style="width: 100%"
      tooltip-effect="dark"
      @selection-change="handleSelectionChange"
    >
      <el-table-column type="selection" width="55" />
      <el-table-column label="操作人" width="140">
        <template slot-scope="scope">
          <div>{{ scope.row.user.userName }}({{ scope.row.user.nickName }})</div>
        </template>
      </el-table-column>
      <el-table-column label="日期" width="180">
        <template slot-scope="scope">{{ scope.row.CreatedAt|formatDate }}</template>
      </el-table-column>
      <el-table-column label="状态码" prop="status" width="120">
        <template slot-scope="scope">
          <div>
            <el-tag type="success">{{ scope.row.status }}</el-tag>
          </div>
        </template>
      </el-table-column>
      <el-table-column label="请求ip" prop="ip" width="120" />
      <el-table-column label="请求方法" prop="method" width="120" />
      <el-table-column label="请求路径" prop="path" width="240" />
      <el-table-column label="请求" prop="path" width="80">
        <template slot-scope="scope">
          <div>
            <el-popover v-if="scope.row.body" placement="top-start" trigger="hover">
              <div class="popover-box">
                <pre>{{ fmtBody(scope.row.body) }}</pre>
              </div>
              <i slot="reference" class="el-icon-view" />
            </el-popover>

            <span v-else>无</span>
          </div>
        </template>
      </el-table-column>
      <el-table-column label="响应" prop="path" width="80">
        <template slot-scope="scope">
          <div>
            <el-popover v-if="scope.row.resp" placement="top-start" trigger="hover">
              <div class="popover-box">
                <pre>{{ fmtBody(scope.row.resp) }}</pre>
              </div>
              <i slot="reference" class="el-icon-view" />
            </el-popover>
            <span v-else>无</span>
          </div>
        </template>
      </el-table-column>
      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-popover v-model="scope.row.visible" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
              <el-button size="mini" type="primary" @click="deleteSysOperationRecord(scope.row)">确定</el-button>
            </div>
            <el-button slot="reference" icon="el-icon-delete" size="mini" type="danger">删除</el-button>
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
  </div>
</template>

<script>
import {
  deleteSysOperationRecord,
  getSysOperationRecordList,
  deleteSysOperationRecordByIds
} from '@/api/sysOperationRecord' // 此处请自行替换地址
import { formatTimeToStr } from '@/utils/date'
import infoList from '@/mixins/infoList'

export default {
  name: 'SysOperationRecord',
  filters: {
    formatDate: function(time) {
      if (time !== null && time !== '') {
        var date = new Date(time)
        return formatTimeToStr(date, 'yyyy-MM-dd hh:mm:ss')
      } else {
        return ''
      }
    },
    formatBoolean: function(bool) {
      if (bool !== null) {
        return bool ? '是' : '否'
      } else {
        return ''
      }
    }
  },
  mixins: [infoList],
  data() {
    return {
      listApi: getSysOperationRecordList,
      dialogFormVisible: false,
      type: '',
      deleteVisible: false,
      multipleSelection: [],
      formData: {
        ip: null,
        method: null,
        path: null,
        status: null,
        latency: null,
        agent: null,
        error_message: null,
        user_id: null
      }
    }
  },
  created() {
    this.getTableData()
  },
  methods: {
    // 条件搜索前端看此方法
    onSubmit() {
      this.page = 1
      this.pageSize = 10
      this.getTableData()
    },
    handleSelectionChange(val) {
      this.multipleSelection = val
    },
    async onDelete() {
      const ids = []
      this.multipleSelection &&
        this.multipleSelection.map(item => {
          ids.push(item.ID)
        })
      const res = await deleteSysOperationRecordByIds({ ids })
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '删除成功'
        })
        if (this.tableData.length === ids.length && this.page > 1) {
          this.page--
        }
        this.deleteVisible = false
        this.getTableData()
      }
    },
    async deleteSysOperationRecord(row) {
      row.visible = false
      const res = await deleteSysOperationRecord({ ID: row.ID })
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '删除成功'
        })
        if (this.tableData.length === 1 && this.page > 1) {
          this.page--
        }
        this.getTableData()
      }
    },
    fmtBody(value) {
      try {
        return JSON.parse(value)
      } catch (err) {
        return value
      }
    }
  }
}
</script>

<style lang="scss">
.table-expand {
  padding-left: 60px;
  font-size: 0;
  label {
    width: 90px;
    color: #99a9bf;
    .el-form-item {
      margin-right: 0;
      margin-bottom: 0;
      width: 50%;
    }
  }
}
.popover-box {
  background: #112435;
  color: #f08047;
  height: 600px;
  width: 420px;
  overflow: auto;
}
.popover-box::-webkit-scrollbar {
  display: none; /* Chrome Safari */
}
</style>
