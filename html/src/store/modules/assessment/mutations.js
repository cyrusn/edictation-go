import _ from 'lodash'

const mutations = {
  addRecords (state, record) {
    state.records.push(record)
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
    console.log(type, payload)
    state[target] = payload
  }
})

export default mutations
