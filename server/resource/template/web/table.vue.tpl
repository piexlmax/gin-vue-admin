<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
           {{- range .Fields}}  {{- if .FieldSearchType}} {{- if eq .FieldType "bool" }}
            <el-form-item label="{{.FieldDesc}}" prop="{{.FieldJson}}">
            <el-select v-model="searchInfo.{{.FieldJson}}" clear placeholder="请选择">
                <el-option
                    key="true"
                    label="是"
                    value="true">
                </el-option>
                <el-option
                    key="false"
                    label="否"
                    value="false">
                </el-option>
            </el-select>
            </el-form-item>
                  {{- else }}
        <el-form-item label="{{.FieldDesc}}">
          <el-input placeholder="搜索条件" v-model="searchInfo.{{.FieldJson}}" />
        </el-form-item> {{ end }} {{ end }}  {{ end }}
        <el-form-item>
          <el-button size="mini" type="primary" icon="el-icon-search" @click="onSubmit">查询</el-button>
          <el-button size="mini" type="primary" icon="el-icon-plus" @click="openDialog">新增</el-button>
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
      border
      stripe
      style="width: 100%"
      tooltip-effect="dark"
      :data="tableData"
      @selection-change="handleSelectionChange"
    >
      <el-table-column type="selection" width="55" />
      <el-table-column label="日期" width="180">
        <template slot-scope="scope">{{ "{{ scope.row.CreatedAt|formatDate }}" }}</template>
      </el-table-column>
      {{- range .Fields}}
      {{- if .DictType}}
      <el-table-column label="{{.FieldDesc}}" prop="{{.FieldJson}}" width="120">
        <template slot-scope="scope">
          {{"{{"}}filterDict(scope.row.{{.FieldJson}},"{{.DictType}}"){{"}}"}}
        </template>
      </el-table-column>
      {{- else if eq .FieldType "bool" }}
      <el-table-column label="{{.FieldDesc}}" prop="{{.FieldJson}}" width="120">
        <template slot-scope="scope">{{ "{{scope.row."}}{{.FieldJson}}{{"|formatBoolean}}" }}</template>
      </el-table-column> {{- else }}
      <el-table-column label="{{.FieldDesc}}" prop="{{.FieldJson}}" width="120" /> {{ end -}}
      {{ end -}}
      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="update{{.StructName}}(scope.row)">变更</el-button>
          <el-button type="danger" icon="el-icon-delete" size="mini" @click="deleteRow(scope.row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination
      layout="total, sizes, prev, pager, next, jumper"
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="[10, 30, 50, 100]"
      :style="{float:'right',padding:'20px'}"
      :total="total"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
    />
    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="弹窗操作">
      <el-form :model="formData" label-position="right" label-width="80px">
    {{- range .Fields}}
        <el-form-item label="{{.FieldDesc}}:">
      {{ if eq .FieldType "bool" }}
          <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.{{.FieldJson}}" clearable ></el-switch>
      {{ end -}}
      {{ if eq .FieldType "string" }}
          <el-input v-model="formData.{{.FieldJson}}" clearable placeholder="请输入" />
      {{ end -}}
      {{ if eq .FieldType "int" }}
      {{- if .DictType}}
          <el-select v-model="formData.{{ .FieldJson }}" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in {{ .DictType }}Options" :key="key" :label="item.label" :value="item.value" />
          </el-select>
      {{ else }}
          <el-input v-model.number="formData.{{ .FieldJson }}" clearable placeholder="请输入" />
      {{ end -}}
      {{ end -}}
      {{ if eq .FieldType "time.Time" }}
          <el-date-picker type="date" placeholder="选择日期" v-model="formData.{{ .FieldJson }}" clearable />
       {{ end -}}
       {{- if eq .FieldType "float64" }}
          <el-input-number v-model="formData.{{ .FieldJson }}" :precision="2" clearable />
       {{ end -}}
            </el-form-item>
       {{- end }}
     </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button type="primary" @click="enterDialog">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {
  create{{.StructName}},
  delete{{.StructName}},
  delete{{.StructName}}ByIds,
  update{{.StructName}},
  find{{.StructName}},
  get{{.StructName}}List
} from '@/api/{{.PackageName}}' //  此处请自行替换地址
import { formatTimeToStr } from '@/utils/date'
import infoList from '@/mixins/infoList'
export default {
  name: '{{.StructName}}',
  mixins: [infoList],
  data() {
    return {
      listApi: get{{ .StructName }}List,
      dialogFormVisible: false,
      type: '',
      deleteVisible: false,
      multipleSelection: [],
      {{ range .Fields}}
          {{- if .DictType }}
      {{ .DictType }}Options: [],
          {{ end -}}
      {{ end }}
      formData: {
        {{range .Fields}}
          {{- if eq .FieldType "bool" -}}
        {{.FieldJson}}: false,
          {{ end -}}
          {{- if eq .FieldType "string" -}}
        {{.FieldJson}}: '',
          {{ end -}}
          {{- if eq .FieldType "int" -}}
        {{.FieldJson}}: 0,
          {{ end -}}
          {{- if eq .FieldType "time.Time" -}}
        {{.FieldJson}}: new Date(),
          {{ end -}}
          {{- if eq .FieldType "float64" -}}
        {{.FieldJson}}: 0,
          {{ end -}}
        {{ end }}
      }
    }
  },
  filters: {
    formatDate: function(time) {
      if (time !== null && time !== '') {
        var date = new Date(time);
        return formatTimeToStr(date, 'yyyy-MM-dd hh:mm:ss');
      } else {
        return ''
      }
    },
    formatBoolean: function(bool) {
      if (bool != null) {
        return bool ? '是' : '否'
      } else {
        return ''
      }
    }
  },
  async created() {
    await this.getTableData()
    {{ range .Fields -}}
      {{- if .DictType }}
    await this.getDict('{{.DictType}}')
      {{ end -}}
    {{- end }}
  },
  methods: {
  // 条件搜索前端看此方法
    onSubmit() {
      this.page = 1
      this.pageSize = 10
      {{- range .Fields}} {{- if eq .FieldType "bool" }}
      if (this.searchInfo.{{.FieldJson}} === ""){
        this.searchInfo.{{.FieldJson}}=null
      } {{ end }} {{ end }}
      this.getTableData()
    },
    handleSelectionChange(val) {
      this.multipleSelection = val
    },
    deleteRow(row) {
      this.$confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.delete{{.StructName}}(row)
      })
    },
    async onDelete() {
      const ids = []
      if (this.multipleSelection.length === 0) {
        this.$message({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      this.multipleSelection &&
        this.multipleSelection.map(item => {
          ids.push(item.ID)
        })
      const res = await delete{{.StructName}}ByIds({ ids })
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
    async update{{.StructName}}(row) {
      const res = await find{{.StructName}}({ ID: row.ID })
      this.type = 'update'
      if (res.code === 0) {
        this.formData = res.data.re{{.Abbreviation}}
        this.dialogFormVisible = true
      }
    },
    closeDialog() {
      this.dialogFormVisible = false
      this.formData = {
        {{range .Fields}}
          {{- if eq .FieldType "bool" -}}
        {{.FieldJson}}: false,
          {{ end -}}
          {{- if eq .FieldType "string" -}}
        {{.FieldJson}}: '',
          {{ end -}}
          {{- if eq .FieldType "int" -}}
        {{.FieldJson}}: 0,
          {{ end -}}
          {{- if eq .FieldType "time.Time" -}}
        {{.FieldJson}}: new Date(),
          {{ end -}}
          {{- if eq .FieldType "float64" -}}
        {{.FieldJson}}: 0,
          {{ end -}}
        {{ end }}
      }
    },
    async delete{{.StructName}}(row) {
      const res = await delete{{.StructName}}({ ID: row.ID })
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '删除成功'
        })
        if (this.tableData.length === 1 && this.page > 1 ) {
          this.page--
        }
        this.getTableData()
      }
    },
    async enterDialog() {
      let res
      switch (this.type) {
        case "create":
          res = await create{{.StructName}}(this.formData)
          break
        case "update":
          res = await update{{.StructName}}(this.formData)
          break
        default:
          res = await create{{.StructName}}(this.formData)
          break
      }
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '创建/更改成功'
        })
        this.closeDialog()
        this.getTableData()
      }
    },
    openDialog() {
      this.type = 'create'
      this.dialogFormVisible = true
    }
  },
}
</script>

<style>
</style>
