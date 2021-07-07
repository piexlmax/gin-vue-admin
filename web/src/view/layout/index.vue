<template>
  <el-container class="layout-cont">
    <el-container :class="[isSider?'openside':'hideside',isMobile ? 'mobile': '']">
      <el-row :class="[isShadowBg?'shadowBg':'']" @click.native="changeShadow()" />
      <el-aside class="main-cont main-left">
        <div class="tilte" :style="{background: backgroundColor}">
          <img alt class="logoimg" :src="$GIN_VUE_ADMIN.appLogo">
          <h2 v-if="isSider" class="tit-text" :style="{color:textColor}">{{ $GIN_VUE_ADMIN.appName }}</h2>
        </div>
        <Aside class="aside" />
      </el-aside>
      <!-- 分块滑动功能 -->
      <el-main class="main-cont main-right">
        <transition :duration="{ enter: 800, leave: 100 }" mode="out-in" name="el-fade-in-linear">
          <div
            :style="{width: `calc(100% - ${isMobile?'0px':isCollapse?'54px':'220px'})`}"
            class="topfix"
          >
            <el-row>
              <!-- :xs="8" :sm="6" :md="4" :lg="3" :xl="1" -->
              <el-header class="header-cont">
                <el-col :xs="2" :lg="1" :md="1" :sm="1" :xl="1">
                  <div class="menu-total" @click="totalCollapse">
                    <i v-if="isCollapse" class="el-icon-s-unfold" />
                    <i v-else class="el-icon-s-fold" />
                  </div>
                </el-col>
                <el-col :xs="10" :lg="14" :md="14" :sm="9" :xl="14">
                  <el-breadcrumb class="breadcrumb" separator-class="el-icon-arrow-right">
                    <el-breadcrumb-item
                      v-for="item in matched.slice(1,matched.length)"
                      :key="item.path"
                    >{{ item.meta.title }}</el-breadcrumb-item>
                  </el-breadcrumb>
                </el-col>
                <el-col :xs="12" :lg="9" :md="9" :sm="14" :xl="9">
                  <div class="fl-right right-box">
                    <Search />
                    <Screenfull class="screenfull" :style="{cursor:'pointer'}" />
                    <el-dropdown>
                      <span class="header-avatar" style="cursor: pointer">
                        <CustomPic />
                        <span style="margin-left: 5px">{{ userInfo.nickName }}</span>
                        <i class="el-icon-arrow-down" />
                      </span>
                      <el-dropdown-menu slot="dropdown" class="dropdown-group">
                        <el-dropdown-item>
                          <span>
                            更多信息
                            <el-badge is-dot />
                          </span>
                        </el-dropdown-item>
                        <el-dropdown-item icon="el-icon-s-custom" @click.native="toPerson">个人信息</el-dropdown-item>
                        <el-dropdown-item icon="el-icon-table-lamp" @click.native="LoginOut">登 出</el-dropdown-item>
                      </el-dropdown-menu>
                    </el-dropdown>
                  </div>
                </el-col>

              </el-header>
            </el-row>
            <!-- 当前面包屑用路由自动生成可根据需求修改 -->
            <!--
            :to="{ path: item.path }" 暂时注释不用-->
            <HistoryComponent />
          </div>
        </transition>
        <transition mode="out-in" name="el-fade-in-linear">
          <keep-alive>
            <router-view v-if="$route.meta.keepAlive && reloadFlag" v-loading="loadingFlag" element-loading-text="正在加载中" class="admin-box" />
          </keep-alive>
        </transition>
        <transition mode="out-in" name="el-fade-in-linear">
          <router-view v-if="!$route.meta.keepAlive && reloadFlag" v-loading="loadingFlag" element-loading-text="正在加载中" class="admin-box" />
        </transition>
        <BottomInfo />
        <setting />
      </el-main>
    </el-container>

  </el-container>
</template>

<script>
import Aside from '@/view/layout/aside'
import HistoryComponent from '@/view/layout/aside/historyComponent/history'
import Screenfull from '@/view/layout/screenfull'
import Search from '@/view/layout/search/search'
import BottomInfo from '@/view/layout/bottomInfo/bottomInfo'
import { mapGetters, mapActions } from 'vuex'
import CustomPic from '@/components/customPic'
import Setting from './setting'
export default {
  name: 'Layout',
  components: {
    Aside,
    HistoryComponent,
    Screenfull,
    Search,
    BottomInfo,
    CustomPic,
    Setting
  },
  data() {
    return {
      show: false,
      isCollapse: false,
      isSider: true,
      isMobile: false,
      isShadowBg: false,
      loadingFlag: false,
      reloadFlag: true,
      value: ''
    }
  },
  computed: {
    ...mapGetters('user', ['userInfo', 'sideMode', 'baseColor']),
    textColor() {
      if (this.$store.getters['user/sideMode'] === 'dark') {
        return '#fff'
      } else if (this.$store.getters['user/sideMode'] === 'light') {
        return '#191a23'
      } else {
        return this.baseColor
      }
    },
    backgroundColor() {
      if (this.sideMode === 'dark') {
        return '#191a23'
      } else if (this.sideMode === 'light') {
        return '#fff'
      } else {
        return this.sideMode
      }
    },
    title() {
      return this.$route.meta.title || '当前页面'
    },
    matched() {
      return this.$route.matched
    }
  },
  mounted() {
    const screenWidth = document.body.clientWidth
    if (screenWidth < 1000) {
      this.isMobile = true
      this.isSider = false
      this.isCollapse = true
    } else if (screenWidth >= 1000 && screenWidth < 1200) {
      this.isMobile = false
      this.isSider = false
      this.isCollapse = true
    } else {
      this.isMobile = false
      this.isSider = true
      this.isCollapse = false
    }
    this.$bus.emit('collapse', this.isCollapse)
    this.$bus.emit('mobile', this.isMobile)
    this.$bus.on('reload', this.reload)
    this.$bus.on('showLoading', () => {
      this.loadingFlag = true
    })
    this.$bus.on('closeLoading', () => {
      this.loadingFlag = false
    })
    window.onresize = () => {
      return (() => {
        const screenWidth = document.body.clientWidth
        if (screenWidth < 1000) {
          this.isMobile = true
          this.isSider = false
          this.isCollapse = true
        } else if (screenWidth >= 1000 && screenWidth < 1200) {
          this.isMobile = false
          this.isSider = false
          this.isCollapse = true
        } else {
          this.isMobile = false
          this.isSider = true
          this.isCollapse = false
        }
        this.$bus.emit('collapse', this.isCollapse)
        this.$bus.emit('mobile', this.isMobile)
      })()
    }
  },
  methods: {
    ...mapActions('user', ['LoginOut']),
    reload() {
      this.reloadFlag = false
      this.$nextTick(() => {
        this.reloadFlag = true
      })
    },
    totalCollapse() {
      this.isCollapse = !this.isCollapse
      this.isSider = !this.isCollapse
      this.isShadowBg = !this.isCollapse
      this.$bus.emit('collapse', this.isCollapse)
    },
    toPerson() {
      this.$router.push({ name: 'person' })
    },
    changeShadow() {
      this.isShadowBg = !this.isShadowBg
      this.isSider = !!this.isCollapse
      this.totalCollapse()
    }
  }
}
</script>

<style lang="scss">
@import '@/style/mobile.scss';

.dark{
  background-color: #191a23 !important;
  color: #fff !important;
}
.light{
  background-color: #fff !important;
  color: #000 !important;
}
</style>
