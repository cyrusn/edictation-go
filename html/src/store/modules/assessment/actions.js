import axios from 'axios'

const actions = {
  getSize ({commit, state}) {
    const {name} = state
    return axios.get(`./api/assessment/${name}/size`)
      .then(response => {
        commit('updateSize', response.data)
      })
  }
}

export default actions
