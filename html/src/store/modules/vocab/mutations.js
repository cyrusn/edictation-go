const mutations = {
  increment (state) {
    state.runningIndex ++
  },
  resetRunningIndex (state) {
    state.runningIndex = 0
  },
  updateDefinition (state, definition) {
    state.definition = definition
  },
  updatePartOfSpeech (state, partOfSpeech) {
    state.partOfSpeech = partOfSpeech
  },
  updateAudioSource (state, src) {
    state.audioSource = src
  },
  updateAnswer (state, answer) {
    state.answer = answer.toLowerCase()
  },
  updateIsCorrect (state, isCorrect) {
    state.isCorrect = isCorrect
  }
}

export default mutations
