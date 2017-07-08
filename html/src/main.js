import Vue from 'vue'
import routes from './routes'
import _ from 'lodash'

/* eslint-disable no-new */
new Vue({
  el: '#app',
  data () {
    return {
      currentRoute: '/',
      name: '',
      clazz: '',
      clazzNo: '',
      mode: 'normal',
      level: 1,
      questionIDs: [],
      stat: {},
      correctPercentage: 0
    }
  },
  computed: {
    ViewComponent () {
      return routes[this.currentRoute]
    },
    noOfQuestions () {
      return this.questionIDs.length
    }
  },
  render (h) {
    return h(this.ViewComponent)
  }
})
