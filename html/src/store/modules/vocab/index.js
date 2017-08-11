import mutations from './mutations'
import actions from './actions'

const state = {
  runningIndex: 0,
  definition: '',
  partOfSpeech: '',
  answer: '',
  audioSource: '',
  isCorrect: false
}

const getters = {
  index: (state, getters, rootState) => {
    const {orders} = rootState.assessment
    const {runningIndex} = state
    return orders[runningIndex]
  }
}

const namespaced = true
export default {
  namespaced, state, getters, mutations, actions
}
