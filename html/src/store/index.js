import Vue from 'vue'
import Vuex from 'vuex'

import mutations from './mutations'
import actions from './actions'

Vue.use(Vuex)

const store = new Vuex.Store({
  state: {
    user: {
      name: '',
      clazz: '',
      clazzNo: 0
    },
    currentRoute: '/',
    assessment: {
      name: '',
      mode: '',
      runningIndex: 0,
      size: 0,
      vocab: {},
      orders: [],
      records: []
    }
  },
  getters: {
    isValid: (state) => {
      const {assessment, user} = state
      const {name, clazz, clazzNo} = user
      const {mode} = assessment
      const assessmentName = assessment.name
      const isValid = name !== '' && clazz !== '' && clazzNo !== '' && assessmentName !== '' && mode !== ''
      return isValid
    },
    mistake: state => {
      return state.assessment.records.length
    },
    index: state => {
      const {runningIndex, orders} = state.assessment
      return orders[runningIndex]
    },
    hp: (state, getters) => {
      const {mistake} = getters
      const {assessment} = state
      const {mode, size} = assessment

      var acceptanceRatio
      switch (mode) {
        case 'normal':
          acceptanceRatio = 0.5
          break
        case 'hard':
          acceptanceRatio = 0.25
          break
        default:
          acceptanceRatio = 1
      }

      const acceptance = size * acceptanceRatio
      return 1 - mistake / acceptance
    },
    progressPercentage: state => {
      // const {} = getters
      const {assessment} = state
      const {runningIndex, size} = assessment
      return runningIndex / size
    }
  },
  mutations,
  actions
})

export default store
