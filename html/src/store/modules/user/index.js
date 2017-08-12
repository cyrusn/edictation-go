import mutations from './mutations'

const state = {
  name: undefined,
  clazz: undefined,
  clazzNo: undefined
}

const namespaced = true
export default {
  namespaced, state, mutations
}
