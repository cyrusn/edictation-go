const state = {
  currentRoute: '/'
}

const mutations = {
  goto (state, route) {
    state.currentRoute = route
  }
}

const namespaced = true
export default {
  namespaced, state, mutations
}
