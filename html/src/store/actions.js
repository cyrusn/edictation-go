import axios from 'axios'

const actions = {
  getSize ({commit}, assessmentName) {
    return axios.get(`./api/assessment/${assessmentName}/size`)
    .then(response => {
      commit('updateAssessmentSize', response.data)
    })
  },
  getVocabFromAPI (context) {
    const {commit, state} = context
    const {name} = state.assessment
    const {index} = context.getters

    return axios.get(`./api/assessment/${name}/index/${index}`)
    .then(response => {
      commit('updateVocab', response.data)
      return {
        name, index
      }
    })
  },
  getAudioSrcFromAPI (context, payload) {
    const {commit} = context
    const {name, index} = payload
    commit('addAudioSrcToVocab', `./api/voice/assessment/${name}/index/${index}`)
  },
  getVocab (context) {
    const {dispatch} = context
    return dispatch('getVocabFromAPI')
      .then(payload => {
        dispatch('getAudioSrcFromAPI', payload)
      })
  },
  getOrders (context) {
    const {commit} = context
    return new Promise((resolve, reject) => {
      commit('updateAssessmentOrders')
      resolve()
    })
  },
  addRecords (context, answer) {
    const {commit, state} = context
    const {name} = state.assessment
    const {index} = context.getters

    return axios.get(`./api/check/assessment/${name}/index/${index}?ans=${answer}`)
      .then(response => {
        const isCorrect = response.data
        if (!isCorrect) {
          commit('addAssessmentRecords', {index, answer})
        }
      })
      .catch(console.error)
  }
}

export default actions
