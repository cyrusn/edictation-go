import axios from 'axios'

const actions = {
  getDefinitionAndPartOfSpeech ({commit, rootState, getters}) {
    const {name} = rootState.assessment
    const {index} = getters

    return axios.get(`./api/assessment/${name}/index/${index}`)
      .then(response => {
        const {definition, partOfSpeech} = response.data
        commit('updateDefinition', definition)
        commit('updatePartOfSpeech', partOfSpeech)
        return {
          name, index
        }
      })
      .catch(console.error)
  },
  getAudioSrcFromAPI ({commit}, payload) {
    const {name, index} = payload
    const audioSource = `./api/voice/assessment/${name}/index/${index}`
    commit('updateAudioSource', audioSource)
    return {name, index}
  },
  checkAnswer ({commit, state, rootState, getters}) {
    const {answer} = state.answer
    const {name} = rootState.assessment
    const {index} = getters

    return axios.get(`./api/check/assessment/${name}/index/${index}?ans=${answer}`)
      .then(response => {
        const isCorrect = response.data
        commit('updateIsCorrect', isCorrect)
        return {
          index, answer
        }
      })
      .catch(console.error)
  }
}

export default actions
