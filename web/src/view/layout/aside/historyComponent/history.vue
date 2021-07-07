<template>
  <div class="router-history">
    <el-tabs
      v-model="activeValue"
      :closable="!(historys.length===1&&this.$route.name===defaultRouter)"
      type="card"
      @contextmenu.prevent.native="openContextMenu($event)"
      @tab-click="changeTab"
      @tab-remove="removeTab"
    >
      <el-tab-pane
        v-for="item in historys"
        :key="name(item)"
        :label="item.meta.title"
        :name="name(item)"
        :tab="item"
        class="gva-tab"
      >
        <span slot="label" :style="{color: activeValue===name(item)?activeColor:'#333'}"><i class="dot" :style="{backgroundColor:activeValue===name(item)?activeColor:'#ddd'}" /> {{ item.meta.title }}</span>
      </el-tab-pane>
    </el-tabs>

    <!--自定义右键菜单html代码-->
    <ul v-show="contextMenuVisible" :style="{left:left+'px',top:top+'px'}" class="contextmenu">
      <li @click="closeAll">关闭所有</li>
      <li @click="closeLeft">关闭左侧</li>
      <li @click="closeRight">关闭右侧</li>
      <li @click="closeOther">关闭其他</li>
    </ul>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
const getFmtString = (item) => {
  return item.name +
      JSON.stringify(item.query) +
      JSON.stringify(item.params)
}
export default {
  name: 'HistoryComponent',
  data() {
    return {
      historys: [],
      activeValue: '',
      contextMenuVisible: false,
      left: 0,
      top: 0,
      isCollapse: false,
      isMobile: false,
      rightActive: ''
    }
  },
  computed: {
    ...mapGetters('user', ['userInfo', 'activeColor']),
    defaultRouter() {
      return this.userInfo.authority.defaultRouter
    }
  },
  watch: {
    contextMenuVisible() {
      if (this.contextMenuVisible) {
        document.body.addEventListener('click', () => {
          this.contextMenuVisible = false
        })
      } else {
        document.body.removeEventListener('click', () => {
          this.contextMenuVisible = false
        })
      }
    },
    $route(to, now) {
      this.historys = this.historys.filter(item => !item.meta.closeTab)
      this.setTab(to)
      sessionStorage.setItem('historys', JSON.stringify(this.historys))
      this.activeValue = window.sessionStorage.getItem('activeValue')
      if (now && to && now.name === to.name) {
        this.$bus.$emit('reload')
      }
    }
  },
  created() {
    this.$bus.on('mobile', isMobile => {
      this.isMobile = isMobile
    })
    this.$bus.on('collapse', isCollapse => {
      this.isCollapse = isCollapse
    })
    const initHistorys = [
      {
        name: this.defaultRouter,
        meta: {
          title: '首页'
        },
        query: {},
        params: {}
      }
    ]
    this.historys =
      JSON.parse(sessionStorage.getItem('historys')) || initHistorys
    if (!window.sessionStorage.getItem('activeValue')) {
      this.activeValue = getFmtString(this.$route)
    } else {
      this.activeValue = window.sessionStorage.getItem('activeValue')
    }
    this.setTab(this.$route)
  },

  beforeDestroy() {
    this.$bus.off('collapse')
    this.$bus.off('mobile')
  },
  methods: {
    name(item) {
      return item.name + JSON.stringify(item.query) + JSON.stringify(item.params)
    },
    openContextMenu(e) {
      if (this.historys.length === 1 && this.$route.name === this.defaultRouter) {
        return false
      }
      if (e.srcElement.id) {
        this.contextMenuVisible = true
        let width
        if (this.isCollapse) {
          width = 54
        } else {
          width = 220
        }
        if (this.isMobile) {
          width = 0
        }
        this.left = e.clientX - width
        this.top = e.clientY + 10
        this.rightActive = e.srcElement.id.split('-')[1]
      }
    },
    closeAll() {
      this.historys = [
        {
          name: this.defaultRouter,
          meta: {
            title: '首页'
          },
          query: {},
          params: {}
        }
      ]
      this.$router.push({ name: this.defaultRouter })
      this.contextMenuVisible = false
      sessionStorage.setItem('historys', JSON.stringify(this.historys))
    },
    closeLeft() {
      let right
      const rightIndex = this.historys.findIndex(item => {
        if (getFmtString(item) === this.rightActive) {
          right = item
        }
        return (
          getFmtString(item) === this.rightActive
        )
      })
      const activeIndex = this.historys.findIndex(
        item => getFmtString(item) === this.activeValue
      )
      this.historys.splice(0, rightIndex)
      if (rightIndex > activeIndex) {
        this.$router.push(right)
      }
      sessionStorage.setItem('historys', JSON.stringify(this.historys))
    },
    closeRight() {
      let right
      const leftIndex = this.historys.findIndex(item => {
        if (getFmtString(item) === this.rightActive) {
          right = item
        }
        return (getFmtString(item) === this.rightActive)
      })
      const activeIndex = this.historys.findIndex(
        item => getFmtString(item) === this.activeValue
      )
      this.historys.splice(leftIndex + 1, this.historys.length)
      if (leftIndex < activeIndex) {
        this.$router.push(right)
      }
      sessionStorage.setItem('historys', JSON.stringify(this.historys))
    },
    closeOther() {
      let right
      this.historys = this.historys.filter(item => {
        if (getFmtString(item) === this.rightActive
        ) {
          right = item
        }
        return (getFmtString(item) === this.rightActive
        )
      })
      this.$router.push(right)
      sessionStorage.setItem('historys', JSON.stringify(this.historys))
    },
    isSame(route1, route2) {
      if (route1.name !== route2.name) {
        return false
      }
      for (const key in route1.query) {
        if (route1.query[key] !== route2.query[key]) {
          return false
        }
      }
      for (const key in route1.params) {
        if (route1.params[key] !== route2.params[key]) {
          return false
        }
      }
      return true
    },
    setTab(route) {
      if (!this.historys.some(item => this.isSame(item, route))) {
        const obj = {}
        obj.name = route.name
        obj.meta = route.meta
        obj.query = route.query
        obj.params = route.params
        this.historys.push(obj)
      }
      window.sessionStorage.setItem(
        'activeValue',
        getFmtString(this.$route)
      )
    },
    changeTab(component) {
      const tab = component.$attrs.tab
      this.$router.push({
        name: tab.name,
        query: tab.query,
        params: tab.params
      })
    },
    removeTab(tab) {
      const index = this.historys.findIndex(
        item => getFmtString(item) === tab
      )
      if (
        getFmtString(this.$route) === tab
      ) {
        if (this.historys.length === 1) {
          this.$router.push({ name: this.defaultRouter })
        } else {
          if (index < this.historys.length - 1) {
            this.$router.push({
              name: this.historys[index + 1].name,
              query: this.historys[index + 1].query,
              params: this.historys[index + 1].params
            })
          } else {
            this.$router.push({
              name: this.historys[index - 1].name,
              query: this.historys[index - 1].query,
              params: this.historys[index - 1].params
            })
          }
        }
      }
      this.historys.splice(index, 1)
    }
  }
}
</script>

<style lang="scss">
.contextmenu {
  width: 100px;
  margin: 0;
  border: 1px solid #ccc;
  background: #fff;
  z-index: 3000;
  position: absolute;
  list-style-type: none;
  padding: 5px 0;
  border-radius: 4px;
  font-size: 14px;
  color: #333;
  box-shadow: 2px 2px 3px 0 rgba(0, 0, 0, 0.2);
}
.el-tabs__item .el-icon-close{
  color: initial !important;
}
.el-tabs__item .dot {
  content: "";
  width: 9px;
  height: 9px;
  margin-right: 8px;
  display: inline-block;
  border-radius: 50%;
  transition: background-color .2s;
}

.contextmenu li {
  margin: 0;
  padding: 7px 16px;
}
.contextmenu li:hover {
  background: #f2f2f2;
  cursor: pointer;
}
</style>
