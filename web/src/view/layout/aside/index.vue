<template>
  <div>
    <el-scrollbar style="height:calc(100vh - 64px)">
      <transition :duration="{ enter: 800, leave: 100 }" mode="out-in" name="el-fade-in-linear">
        <el-menu
          :collapse="isCollapse"
          :collapse-transition="true"
          :default-active="active"
          :background-color="sideMode"
          :active-text-color="activeColor"
          :text-color="baseColor"
          class="el-menu-vertical"
          unique-opened
          @select="selectMenuItem"
        >
          <template v-for="item in asyncRouters[0].children">
            <aside-component v-if="!item.hidden" :key="item.name" :router-info="item" />
          </template>
        </el-menu>
      </transition>
    </el-scrollbar>
  </div>
</template>

<script>
import { mapGetters, mapMutations } from 'vuex'
import AsideComponent from '@/view/layout/aside/asideComponent'
export default {
  name: 'Aside',
  components: {
    AsideComponent
  },
  data() {
    return {
      active: '',
      isCollapse: false
    }
  },
  computed: {
    ...mapGetters('router', ['asyncRouters']),
    ...mapGetters('user', ['baseColor', 'activeColor', 'sideMode'])
  },
  watch: {
    $route() {
      this.active = this.$route.name
    }
  },
  created() {
    this.active = this.$route.name
    const screenWidth = document.body.clientWidth
    if (screenWidth < 1000) {
      this.isCollapse = !this.isCollapse
    }

    this.$bus.on('collapse', item => {
      this.isCollapse = item
    })
  },
  beforeDestroy() {
    this.$bus.off('collapse')
  },
  methods: {
    ...mapMutations('history', ['addHistory']),
    selectMenuItem(index, _, ele) {
      const query = {}
      const params = {}
      ele.route.parameters &&
      ele.route.parameters.map(item => {
        if (item.type === 'query') {
          query[item.key] = item.value
        } else {
          params[item.key] = item.value
        }
      })
      if (index === this.$route.name) return
      if (index.indexOf('http://') > -1 || index.indexOf('https://') > -1) {
        window.open(index)
      } else {
        this.$router.push({ name: index, query, params })
      }
    }
  }
}
</script>

<style lang="scss">
.el-submenu__title,.el-menu-item{
  i{
    color: inherit !important;
  }
}

.el-submenu__title:hover,.el-menu-item:hover{
  i{
    color: inherit !important;
  }
  span{
    color: inherit !important;
  }
}

.el-scrollbar {
  .el-scrollbar__view {
    height: 100%;
  }
}
.menu-info {
  .menu-contorl {
    line-height: 52px;
    font-size: 20px;
    display: table-cell;
    vertical-align: middle;
  }
}
</style>
