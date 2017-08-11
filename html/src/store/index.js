import Vue from 'vue'
import Vuex from 'vuex'

import user from './modules/user'
import assessment from './modules/assessment'
import vocab from './modules/vocab'
import router from './modules/router'

Vue.use(Vuex)

const store = new Vuex.Store({
  modules: {
    user, assessment, vocab, router
  }
})

export default store
