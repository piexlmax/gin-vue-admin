<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="字典名（中）">
          <el-input v-model="searchInfo.name" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="字典名（英）">
          <el-input v-model="searchInfo.type" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-select v-model="searchInfo.status" clear placeholder="请选择">
            <el-option key="true" label="是" value="true" />
            <el-option key="false" label="否" value="false" />
          </el-select>
        </el-form-item>
        <el-form-item label="描述">
          <el-input v-model="searchInfo.desc" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" icon="el-icon-search" @click="onSubmit">查询</el-button>
          <el-button size="mini" type="primary" icon="el-icon-plus" @click="openDialog">新增</el-button>
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
      <el-table-column label="日期" width="180">
        <template slot-scope="scope">{{ scope.row.CreatedAt|formatDate }}</template>
      </el-table-column>

      <el-table-column label="字典名（中）" prop="name" width="120" />

      <el-table-column label="字典名（英）" prop="type" width="120" />

      <el-table-column label="状态" prop="status" width="120">
        <template slot-scope="scope">{{ scope.row.status|formatBoolean }}</template>
      </el-table-column>

      <el-table-column label="描述" prop="desc" width="280" />

      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button size="mini" type="success" @click="toDetile(scope.row)">详情</el-button>
          <el-button size="mini" type="primary" @click="updateSysDictionary(scope.row)">变更</el-button>
          <el-popover v-model="scope.row.visible" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
              <el-button type="primary" size="mini" @click="deleteSysDictionary(scope.row)">确定</el-button>
            </div>
            <el-button slot="reference" type="danger" icon="el-icon-delete" size="mini" style="margin-left:10px">删除</el-button>
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

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="弹窗操作">
      <el-form ref="elForm" :model="formData" :rules="rules" size="medium" label-width="110px">
        <el-form-item label="字典名（中）" prop="name">
          <el-input
            v-model="formData.name"
            placeholder="请输入字典名（中）"
            clearable
            :style="{width: '100%'}"
          />
        </el-form-item>
        <el-form-item label="字典名（英）" prop="type">
          <el-input
            v-model="formData.type"
            placeholder="请输入字典名（英）"
            clearable
            :style="{width: '100%'}"
          />
        </el-form-item>
        <el-form-item label="状态" prop="status" required>
          <el-switch v-model="formData.status" active-text="开启" inactive-text="停用" />
        </el-form-item>
        <el-form-item label="描述" prop="desc">
          <el-input v-model="formData.desc" placeholder="请输入描述" clearable :style="{width: '100%'}" />
        </el-form-item>
      </el-form>

      <div slot="footer" class="dialog-footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button type="primary" @click="enterDialog">确 定</el-button>
      </div>
    </el-dialog>

    <div style="margin-top:40px;color:red">获取字典且缓存方法已在前端utils/dictionary 已经封装完成 不必自己书写 使用方法查看文件内注释</div>
  </div>
</template>

<script>
import {
  createSysDictionary,
  deleteSysDictionary,
  updateSysDictionary,
  findSysDictionary,
  getSysDictionaryList
} from '@/api/sysDictionary' //  此处请自行替换地址
import { formatTimeToStr } from '@/utils/date'
import infoList from '@/mixins/infoList'
export default {
  name: 'SysDictionary',
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
      listApi: getSysDictionaryList,
      dialogFormVisible: false,
      type: '',
      formData: {
        name: null,
        type: null,
        status: true,
        desc: null
      },
      rules: {
        name: [
          {
            required: true,
            message: '请输入字典名（中）',
            trigger: 'blur'
          }
        ],
        type: [
          {
            required: true,
            message: '请输入字典名（英）',
            trigger: 'blur'
          }
        ],
        desc: [
          {
            required: true,
            message: '请输入描述',
            trigger: 'blur'
          }
        ]
      }
    }
  },
  async created() {
    this.getTableData()
  },
  methods: {
    toDetile(row) {
      this.$router.push({
        name: 'dictionaryDetail',
        params: {
          id: row.ID
        }
      })
    },
    // 条件搜索前端看此方法
    onSubmit() {
      this.page = 1
      this.pageSize = 10
      if (this.searchInfo.status === '') {
        this.searchInfo.status = null
      }
      this.getTableData()
    },
    async updateSysDictionary(row) {
      const res = await findSysDictionary({ ID: row.ID })
      this.type = 'update'
      if (res.code === 0) {
        this.formData = res.data.resysDictionary
        this.dialogFormVisible = true
      }
    },
    closeDialog() {
      this.dialogFormVisible = false
      this.formData = {
        name: null,
        type: null,
        status: true,
        desc: null
      }
    },
    async deleteSysDictionary(row) {
      row.visible = false
      const res = await deleteSysDictionary({ ID: row.ID })
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
    async enterDialog() {
      this.$refs['elForm'].validate(async valid => {
        if (!valid) return
        let res
        switch (this.type) {
          case 'create':
            res = await createSysDictionary(this.formData)
            break
          case 'update':
            res = await updateSysDictionary(this.formData)
            break
          default:
            res = await createSysDictionary(this.formData)
            break
        }
        if (res.code === 0) {
          this.closeDialog()
          this.getTableData()
        }
      })
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
