<template>
  <div v-loading="loading">
    <slot :tableData="list" />
    <el-pagination
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="[10, 30, 50, 100]"
      :style="{ float: 'right', padding: '20px' }"
      :total="total"
      @current-change="changePage"
      @size-change="changeSize"
      layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>
  </div>
</template>

<script>
export default {
  props: {
    dataSource: {
      type: Function,
      default: () => {
        return null
      },
    },
    defaultPageSize: {
      type: Number,
      default: () => {
        return 10
      },
    },
  },
  data() {
    return {
      loading: false,
      list: [],
      total: 0,
      page: 1,
      pageSize: 10,
    }
  },
  mounted() {
    this.pageSize = this.defaultPageSize
    this.getTableData()
  },
  methods: {
    async getTableData(newParams = {}) {
      if (this.dataSource.call) {
        this.loading = true
        const res = await this.dataSource({
          page: this.page,
          pageSize: this.pageSize,
          ...newParams,
        })

        this.loading = false
        this.list = res.data.list
        this.total = res.data.total
      }
    },
    changePage(index) {
      this.page = index
      this.getTableData()
    },
    changeSize(limit) {
      this.pageSize = limit
      this.getTableData()
    },
  },
}
</script>
