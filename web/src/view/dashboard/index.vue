<template>
  <div class="big">
    <el-row>
      <div class="card">
        <el-col :xs="24" :lg="16" :md="16">
          <div class="car-left">
            <el-row>
              <div>
                <el-col :xs="4" :md="3" :lg="3">
                  <span class="card-img">
                    <img :src="userInfo.headerImg" alt="">
                  </span>
                </el-col>
                <el-col :xs="20" :lg="12" :md="12">
                  <div class="text">
                    <h4>早安，管理员， 请开始您一天的工作吧！</h4>
                    <p class="tips-text">
                      <i class="el-icon-sunny el-icon" />
                      <span>今日晴，0℃ - 10℃，天气寒冷，注意添加衣物。</span>
                    </p>
                  </div>
                </el-col>
              </div>
            </el-row>
          </div>
        </el-col>
        <el-col :xs="24" :lg="8" :md="8">
          <div class="car-right">
            <el-row>
              <el-col :span="8">
                <div class="car-item">
                  <span class="flow"><i class="el-icon-s-grid" /></span>
                  <span>今日流量 </span>
                  <b>13260</b>
                </div>
              </el-col>
              <el-col :span="8">
                <div class="car-item">
                  <span class="user-number">
                    <i class="el-icon-s-custom" />
                  </span>
                  <span>总用户 </span>
                  <b>48286</b>
                </div>
              </el-col>
              <el-col :span="8">
                <div class="car-item">
                  <span class="feedback">
                    <i class="el-icon-star-on" />
                  </span>
                  <span>好评率 </span>
                  <b>98%</b>
                </div>
              </el-col>
            </el-row>
          </div>
        </el-col>
      </div>
    </el-row>
    <el-row>
      <el-card shadow="hover">
        <h2>
          使用教学：<a style="color:#409EFF" target="view_window" href="https://www.bilibili.com/video/BV1fV411y7dT/">https://www.bilibili.com/video/BV1fV411y7dT/</a>
        </h2>
        <br>
        <h2>
          工作流教学：<a style="color:#409EFF" target="view_window" href="https://www.bilibili.com/video/BV1Ka411F7Ji/">https://www.bilibili.com/video/BV1Ka411F7Ji/</a>
        </h2>
        <div />
      </el-card>
    </el-row>
    <div class="shadow">
      <el-row :gutter="20">
        <el-col
          v-for="(card, key) in toolCards"
          :key="key"
          :span="4"
          :xs="8"
          @click.native="toTarget(card.name)"
        >
          <el-card shadow="hover" class="grid-content">
            <i :class="card.icon" :style="{ color: card.color }" />
            <p>{{ card.label }}</p>
          </el-card>
        </el-col>
      </el-row>
    </div>
    <div class="bottom">
      <el-row :gutter="32">
        <el-col :xs="24" :sm="24" :lg="12">
          <div class="chart-player">
            <music-player />
          </div>
        </el-col>
        <el-col :xs="24" :sm="24" :lg="12">
          <div class="chart-player">
            <todo-list />
          </div>
        </el-col>
      </el-row>
    </div>
  </div>
</template>

<script>
import musicPlayer from './component/musicPlayer'
import TodoList from './component/todoList'
import { mapGetters } from 'vuex'
export default {
  name: 'Dashboard',
  components: {
    musicPlayer, // 音乐播放器
    TodoList // TodoList
    // RaddarChart, //雷达图
    // stackMap, //堆叠图
    // Sunburst, //旭日图
  },
  data() {
    return {
      toolCards: [
        {
          label: '用户管理',
          icon: 'el-icon el-icon-monitor',
          name: 'user',
          color: '#ff9c6e'
        },
        {
          label: '角色管理',
          icon: 'el-icon el-icon-setting',
          name: 'authority',
          color: '#69c0ff'
        },
        {
          label: '菜单管理',
          icon: 'el-icon el-icon-menu',
          name: 'menu',
          color: '#b37feb'
        },
        {
          label: '代码生成器',
          icon: 'el-icon el-icon-cpu',
          name: 'autoCode',
          color: '#ffd666'
        },
        {
          label: '表单生成器',
          icon: 'el-icon el-icon-document-checked',
          name: 'formCreate',
          color: '#ff85c0'
        },
        {
          label: '关于我们',
          icon: 'el-icon el-icon-user',
          name: 'about',
          color: '#5cdbd3'
        }
      ]
    }
  },
  computed: {
    ...mapGetters('user', ['userInfo'])
  },
  methods: {
    toTarget(name) {
      this.$router.push({ name })
    }
  }
}
</script>

<style lang="scss" scoped>
.big {
  margin: 100px 0 0 0;
  padding-top: 0;
  background-color: rgb(243, 243, 243);
  .top {
    width: 100%;
    height: 360px;
    margin-top: 20px;
    overflow: hidden;
    .chart-container {
      position: relative;
      width: 100%;
      height: 100%;
      padding: 20px;
      background-color: #fff;
    }
  }
  .mid {
    width: 100%;
    height: 380px;
    .chart-wrapper {
      height: 340px;
      background: #fff;
      padding: 16px 16px 0;
      margin-bottom: 32px;
    }
  }
  .bottom {
    width: 100%;
    height: 300px;
    // margin: 20px 0;
    .el-row {
      margin-right: 4px !important;
    }
    .chart-player {
      width: 100%;
      height: 270px;
      padding: 10px;
      background-color: #fff;
    }
  }
}
</style>
