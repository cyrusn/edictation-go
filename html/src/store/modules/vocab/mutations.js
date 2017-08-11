const mutations = {
  increment (state) {
    state.assessment.runningIndex ++
  },
  resetIndex (state) {
    state.assessment.runningIndex = 0
  },
  updateDefinition (state, definition) {
    state.assessment.definition = definition
  },
  updatePartOfSpeech (state, partOfSpeech) {
    state.assessment.partOfSpeech = partOfSpeech
  },
  updateAudioSource (state, src) {
    state.assessment.audioSource = src
  },
  updateAnswer (state, answer) {
    state.assessment.answer = answer
  },
  updateIsCorrect (state, isCorrect) {
    state.assessment.isCorrect = isCorrect
  }
}

export default mutations
