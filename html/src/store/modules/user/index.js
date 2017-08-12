import mutations from './mutations'

const state = {
  name: '',
  clazz: '',
  clazzNo: ''
}

const namespaced = true
export default {
  namespaced, state, mutations
}
