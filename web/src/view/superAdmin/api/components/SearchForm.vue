<template>
  <div class="search-term">
    <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
      <el-form-item label="路径">
        <el-input placeholder="路径" v-model="searchInfo.path" @keyup.enter.native="emitSearch"></el-input>
      </el-form-item>
      <el-form-item label="描述">
        <el-input placeholder="描述" v-model="searchInfo.description" @keyup.enter.native="emitSearch"></el-input>
      </el-form-item>
      <el-form-item label="api组">
        <el-input placeholder="api组" v-model="searchInfo.apiGroup" @keyup.enter.native="emitSearch"></el-input>
      </el-form-item>
      <el-form-item label="请求">
        <el-select clearable placeholder="请选择" v-model="searchInfo.method" @change="emitSearch">
          <el-option
            :key="item.value"
            :label="`${item.label}(${item.value})`"
            :value="item.value"
            v-for="item in methodOptions"
          />
        </el-select>
      </el-form-item>
      <el-form-item>
        <el-button @click="onReset" type="primary">重置</el-button>
      </el-form-item>
      <el-form-item>
        <el-button @click="$emit('add')" type="primary">新增api</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import config from './config'

export default {
  data() {
    return {
      methodOptions: config.methodOptions,
      searchInfo: {},
    }
  },
  methods: {
    emitSearch() {
      this.$emit('change', this.searchInfo)
    },
    onReset() {
      this.searchInfo = {}
      this.emitSearch()
    },
  },
}
</script>
