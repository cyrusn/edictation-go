import axios from 'axios'

const actions = {
  getDefinitionAndPartOfSpeech ({commit, rootState, getters}) {
    const {name} = rootState.assessment
    const {vocabIndex} = getters

    return axios.get(`./api/assessment/${name}/index/${vocabIndex}`)
      .then(response => {
        const {definition, partOfSpeech} = response.data
        commit('updateDefinition', definition)
        commit('updatePartOfSpeech', partOfSpeech)
        return {
          name, vocabIndex
        }
      })
      .catch(console.error)
  },
  getAudioSrc ({commit}, {name, vocabIndex}) {
    const audioSource = `./api/voice/assessment/${name}/index/${vocabIndex}`
    commit('updateAudioSource', audioSource)
    return {name, vocabIndex}
  },
  checkAnswer ({commit, state, rootState, getters}) {
    const {answer} = state
    const {name} = rootState.assessment
    const {vocabIndex} = getters

    return axios.get(`./api/check/assessment/${name}/index/${vocabIndex}?ans=${answer}`)
      .then(response => {
        const isCorrect = response.data
        commit('updateIsCorrect', isCorrect)
        return {
          isCorrect, vocabIndex, answer
        }
      })
      .catch(console.error)
  }
}

export default actions
