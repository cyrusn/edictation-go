import _ from 'lodash'

const mutations = {
  addRecords (state, {isCorrect, vocabIndex, answer}) {
    if (!isCorrect) {
      state.records.push({
        vocabIndex, answer
      })
    }
  },
  clearRecords (state) {
    state.records = []
  },
  shuffleOrders (state) {
    const {size} = state
    state.orders = _(size).range().shuffle().value()
  }
}

const mutationTypes = [{
  type: 'updateName',
  target: 'name'
}, {
  type: 'updateMode',
  target: 'mode'
}, {
  type: 'updateSize',
  target: 'size'
}]

mutationTypes.forEach(({type, target}) => {
  mutations[type] = function (state, payload) {
    state[target] = payload
  }
})

export default mutations
