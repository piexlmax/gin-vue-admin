import { login } from '@/api/user'
import { jsonInBlacklist } from '@/api/jwt'
import router from '@/router/index'
import { setUserInfo } from '@/api/user'
import { Message } from 'element-ui'

export const user = {
  namespaced: true,
  state: {
    userInfo: {
      uuid: '',
      nickName: '',
      headerImg: '',
      authority: '',
      sideMode: 'dark',
      activeColor: '#1890ff',
      baseColor: '#fff'
    },
    token: ''
  },
  mutations: {
    setUserInfo(state, userInfo) {
      // 这里的 `state` 对象是模块的局部状态
      state.userInfo = userInfo
    },
    setToken(state, token) {
      // 这里的 `state` 对象是模块的局部状态
      state.token = token
    },
    NeedInit(state) {
      state.userInfo = {}
      state.token = ''
      sessionStorage.clear()
      router.push({ name: 'Init', replace: true })
    },
    LoginOut(state) {
      state.userInfo = {}
      state.token = ''
      sessionStorage.clear()
      router.push({ name: 'Login', replace: true })
      window.location.reload()
    },
    ResetUserInfo(state, userInfo = {}) {
      state.userInfo = { ...state.userInfo,
        ...userInfo
      }
    },
    ChangeActiveColor: async(state, val) => {
      state.userInfo.activeColor = val
    },
    ChangeSideMode: async(state, val) => {
      state.userInfo.sideMode = val
    },
    ChangeBaseColor: (state, val) => {
      state.userInfo.baseColor = val
    }
  },
  actions: {
    async LoginIn({ commit, dispatch, rootGetters, getters }, loginInfo) {
      const res = await login(loginInfo)
      if (res.code === 0) {
        commit('setUserInfo', res.data.user)
        commit('setToken', res.data.token)
        await dispatch('router/SetAsyncRouter', {}, { root: true })
        const asyncRouters = rootGetters['router/asyncRouters']
        router.addRoutes(asyncRouters)
        // const redirect = router.history.current.query.redirect
        // console.log(redirect)
        // if (redirect) {
        //     router.push({ path: redirect })
        // } else {
        router.push({ name: getters['userInfo'].authority.defaultRouter })
        // }
        return true
      }
    },
    async LoginOut({ commit }) {
      const res = await jsonInBlacklist()
      if (res.code === 0) {
        commit('LoginOut')
      }
    },
    async changeActiveColor({ commit, state }, data) {
      const res = await setUserInfo({ activeColor: data, ID: state.userInfo.ID })
      if (res.code === 0) {
        commit('ChangeActiveColor', data)
        Message({
          type: 'success',
          message: '设置成功'
        })
      }
    },
    async changeSideMode({ commit, state }, data) {
      const res = await setUserInfo({ sideMode: data, ID: state.userInfo.ID })
      if (res.code === 0) {
        commit('ChangeSideMode', data)
        Message({
          type: 'success',
          message: '设置成功'
        })
      }
    },
    async changeBaseColor({ commit, state }, data) {
      const res = await setUserInfo({ baseColor: data, ID: state.userInfo.ID })
      if (res.code === 0) {
        commit('ChangeBaseColor', data)
        Message({
          type: 'success',
          message: '设置成功'
        })
      }
    }
  },
  getters: {
    userInfo(state) {
      return state.userInfo
    },
    token(state) {
      return state.token
    },
    mode(state) {
      return state.userInfo.sideMode
    },
    sideMode(state) {
      if (state.userInfo.sideMode === 'dark') {
        return '#191a23'
      } else if (state.userInfo.sideMode === 'light') {
        return '#fff'
      } else {
        return state.userInfo.sideMode
      }
    },
    baseColor(state) {
      if (state.userInfo.sideMode === 'dark') {
        return '#fff'
      } else if (state.userInfo.sideMode === 'light') {
        return '#191a23'
      } else {
        return state.userInfo.baseColor
      }
    },
    activeColor(state) {
      if (state.userInfo.sideMode === 'dark' || state.userInfo.sideMode === 'light') {
        return '#1890ff'
      }
      return state.userInfo.activeColor
    }

  }
}
