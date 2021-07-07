<template>
  <div>
    <el-form :model="formData" label-position="right" label-width="80px">
    {{- range .Fields}}
      <el-form-item label="{{.FieldDesc}}:">
    {{ if eq .FieldType "bool" -}}
      <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.{{.FieldJson}}" clearable ></el-switch>
    {{ end -}}
    {{ if eq .FieldType "string" -}}
      <el-input v-model="formData.{{.FieldJson}}" clearable placeholder="请输入" />
    {{ end -}}
    {{ if eq .FieldType "int" -}}
    {{ if .DictType -}}
      <el-select v-model="formData.{{ .FieldJson }}" placeholder="请选择" clearable>
        <el-option v-for="(item,key) in {{ .DictType }}Options" :key="key" :label="item.label" :value="item.value"></el-option>
      </el-select>
    {{ else -}}
      <el-input v-model.number="formData.{{ .FieldJson }}" clearable placeholder="请输入"/>
    {{ end -}}
    {{ end -}}
    {{ if eq .FieldType "time.Time" }}
      <el-date-picker type="date" placeholder="选择日期" v-model="formData.{{ .FieldJson }}" clearable></el-date-picker>
    {{ end -}}
    {{ if eq .FieldType "float64" }}
      <el-input-number v-model="formData.{{ .FieldJson }}" :precision="2" clearable></el-input-number>
    {{ end -}}
      </el-form-item>
    {{ end -}}
      <el-form-item>
        <el-button size="mini" type="primary" @click="save">保存</el-button>
        <el-button size="mini" type="primary" @click="back">返回</el-button>
      </el-form-item>
    </el-form>
</div>
</template>

<script>
import {
  create{{.StructName}},
  update{{.StructName}},
  find{{.StructName}}
} from '@/api/{{.PackageName}}' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
export default {
  name: '{{.StructName}}',
  mixins: [infoList],
  data() {
    return {
      type: '',
      {{range .Fields}}
          {{- if .DictType }}
      {{ .DictType }}Options: [],
          {{ end -}}
      {{end -}}
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
  async created() {
    // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (this.$route.query.id) {
      const res = await find{{.StructName}}({ ID: this.$route.query.id })
      if (res.code === 0) {
        this.formData = res.data.re{{.Abbreviation}}
        this.type = 'update'
      }
    } else {
      this.type = 'create'
    }
    {{ range .Fields -}}
      {{- if .DictType }}
    await this.getDict("{{.DictType}}")
      {{ end -}}
    {{- end }}
  },
  methods: {
    async save() {
      let res
      switch (this.type) {
        case 'create':
          res = await create{{.StructName}}(this.formData)
          break
        case 'update':
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
      }
    },
    back() {
      this.$router.go(-1)
    }
  }
}
</script>

<style>
</style>