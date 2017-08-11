import axios from 'axios'

const actions = {
  getSize ({commit}, name) {
    return axios.get(`./api/assessment/${name}/size`)
      .then(response => {
        commit('updateSize', response.data)
      })
  }
}

export default actions
