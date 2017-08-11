const mutations = {}
const mutationTypes = [{
  type: 'updateName',
  target: 'name'
}, {
  type: 'updateClazz',
  target: 'clazz'
}, {
  type: 'updateClazzNo',
  target: 'clazzNo'
}]

mutationTypes.forEach(({type, target}) => {
  mutations[type] = function (state, payload) {
    state[target] = payload
  }
})

export default mutations
