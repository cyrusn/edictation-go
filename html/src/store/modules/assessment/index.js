import mutations from './mutations'
import actions from './actions'

const state = {
  name: '',
  mode: '',
  size: 0,
  orders: [],
  records: []
}

const getters = {
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
