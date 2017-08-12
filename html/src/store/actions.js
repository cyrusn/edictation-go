const actions = {
  onStart ({commit, dispatch}) {
    dispatch('assessment/getSize').then(() => {
      commit('assessment/shuffleOrders')
    }).then(() => {
      dispatch('updateQuizPanel')
    })
  },
  onNext ({commit, state, dispatch, getters}) {
    dispatch('vocab/checkAnswer').then(payload => {
      commit('assessment/addRecords', payload)
    }).then(() => {
      commit('vocab/increment')
    }).then(() => {
      switch (getters['assessment/nextMove']) {
        case 'RESTART_ASSESSMENT':
          dispatch('restart')
          break
        case 'UPDATE_QUIZ_PANEL':
          dispatch('updateQuizPanel')
          break
        case 'GO_TO_REPORT':
          commit('router/goto', '/report')
          break
        default:
          console.error('Do know the decision')
      }
    })
  },
  restart ({commit, dispatch, state}) {
    commit('assessment/clearRecords')
    commit('vocab/resetRunningIndex')
    commit('assessment/shuffleOrders')

    dispatch('vocab/getDefinitionAndPartOfSpeech').then(payload => {
      dispatch('vocab/getAudioSrc', payload)
    })
  },
  updateQuizPanel ({commit, dispatch}) {
    commit('vocab/updateAnswer', '')

    dispatch('vocab/getDefinitionAndPartOfSpeech').then(payload => {
      dispatch('vocab/getAudioSrc', payload)
    })
  }
}

export default actions
