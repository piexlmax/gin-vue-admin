<template>
  <div>
    <el-button type="primary" class="drawer-container" icon="el-icon-setting" @click="showSettingDrawer" />
    <el-drawer
      title="系统配置"
      :visible.sync="drawer"
      :direction="direction"
      :before-close="handleClose"
    >
      <div class="setting_body">
        <div class="setting_card">
          <div class="setting_title">侧边栏主题 (注：自定义请先配置背景色)</div>
          <div class="setting_content">
            <div class="theme-box">
              <div class="item" @click="changeMode('light')">
                <i v-if="mode === 'light'" class="el-icon-check check" />
                <img src="https://gw.alipayobjects.com/zos/antfincdn/NQ%24zoisaD2/jpRkZQMyYRryryPNtyIC.svg">
              </div>
              <div class="item" @click="changeMode('dark')">
                <i v-if="mode === 'dark'" class="el-icon-check check" />
                <img src="https://gw.alipayobjects.com/zos/antfincdn/XwFOFbLkSM/LCkqqYNmvBEbokSDscrm.svg">
              </div>
            </div>
            <div class="color-box">
              <div>
                <div class="setting_title">自定义背景色</div>
                <el-color-picker :value="sideMode" @change="changeMode" />
              </div>
              <div>
                <div class="setting_title">自定义基础色</div>
                <el-color-picker :value="baseColor" @change="changeBaseColor" />
              </div>
              <div>
                <div class="setting_title">活跃色</div>
                <el-color-picker :value="activeColor" @change="activeColorChange" />
              </div>
            </div>
          </div>
        </div>
      </div>
    </el-drawer>

  </div>
</template>

<script>
import { mapGetters } from 'vuex'
export default {
  data() {
    return {
      drawer: false,
      direction: 'rtl'
    }
  },
  computed: {
    ...mapGetters('user', ['sideMode', 'baseColor', 'activeColor', 'mode'])
  },
  methods: {
    handleClose() {
      this.drawer = false
    },
    showSettingDrawer() {
      this.drawer = true
    },
    changeMode(e) {
      if (e === null) {
        this.$store.dispatch('user/changeSideMode', 'dark')
        return
      }
      this.$store.dispatch('user/changeSideMode', e)
    },
    changeBaseColor(e) {
      if (e === null) {
        this.$store.dispatch('user/changeBaseColor', '#fff')
        return
      }
      this.$store.dispatch('user/changeBaseColor', e)
    },
    activeColorChange(e) {
      if (e === null) {
        this.$store.dispatch('user/changeActiveColor', '#1890ff')
        return
      }
      this.$store.dispatch('user/changeActiveColor', e)
    }
  }
}
</script>

<style lang="scss" scoped>
.drawer-container {
  position: fixed;
  right: 0;
  bottom: 15%;
  height: 40px;
  width: 40px;
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 999;
  color: #fff;
  border-radius: 4px 0 0 4px;
  cursor: pointer;
  -webkit-box-shadow: inset 0 0 6px rgb(0 ,0 ,0, 10%);
}
.setting_body{
  padding: 20px;
  .setting_card{
    margin-bottom: 20px;
  }
  .setting_content{
    margin-top: 20px;
    display: flex;
    flex-direction: column;
    >.theme-box{
     display: flex;
    }
    >.color-box{
      div{
        display: flex;
        flex-direction: column;
      }
    }
    .item{
      position: relative;
      display: flex;
      align-items: center;
      justify-content: center;
      .check{
        position: absolute;
        font-size: 20px;
        color: #00afff;
      }
      img{
        margin-right: 20px;
      }
    }
  }
}

</style>
