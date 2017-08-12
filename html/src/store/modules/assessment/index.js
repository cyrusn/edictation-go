import mutations from './mutations'
import actions from './actions'

const state = {
  name: undefined,
  mode: undefined,
  size: undefined,
  orders: [],
  records: []
}

const getters = {
  nextMove: (state, getters, rootState) => {
    const {size} = state
    const {hp} = getters
    const {runningIndex} = rootState.vocab

    switch (true) {
      case hp < 0:
        return 'RESTART_ASSESSMENT'
      case runningIndex < size:
        return 'UPDATE_QUIZ_PANEL'
      default:
        return 'GO_TO_REPORT'
    }
  },
  mistake: state => {
    return state.records.length
  },
  hp: (state, getters) => {
    const {mistake} = getters
    const {mode, size} = state

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
  contextualColor (state, getters) {
    const {hp} = getters
    switch (true) {
      case hp < 0.4 && hp > 0.1:
        return 'warning'
      case hp <= 0.1:
        return 'danger'
      default:
        return 'success'
    }
  },
  progressPercentage: (state, getters, rootState) => {
    const {size} = state
    const {runningIndex} = rootState.vocab
    return runningIndex / size
  }
}

const namespaced = true
export default {
  namespaced, state, getters, mutations, actions
}
