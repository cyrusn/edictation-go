import _ from 'lodash'

const mutations = {
  goto (state, route) {
    state.currentRoute = route
  },
  addAssessmentRecords (state, record) {
    state.assessment.records.push(record)
  },
  clearAssessmentRecords (state) {
    state.assessment.records = []
  },
  increment (state) {
    state.assessment.runningIndex ++
  },
  resetIndex (state) {
    state.assessment.runningIndex = 0
  },
  updateAssessmentOrders (state) {
    const {assessment} = state
    const {size} = assessment
    assessment.orders = _(size).range().shuffle().value()
  },
  updateVocab (state, vocab) {
    state.assessment.vocab = vocab
  },
  addAudioSrcToVocab (state, src) {
    state.assessment.vocab.src = src
  }
}

const userMutations = [{
  type: 'updateUserName',
  target: 'user.name'
}, {
  type: 'updateUserClazz',
  target: 'user.clazz'
}, {
  type: 'updateUserClazzNo',
  target: 'user.clazzNo'
}, {
  type: 'updateAssessmentName',
  target: 'assessment.name'
}, {
  type: 'updateAssessmentMode',
  target: 'assessment.mode'
}, {
  type: 'updateAssessmentSize',
  target: 'assessment.size'
}]

userMutations.forEach(({type, target}) => {
  const targets = target.split('.')
  mutations[type] = function (state, payload) {
    state[targets[0]][targets[1]] = payload
  }
})

export default mutations
