import Vue from 'vue'
import _ from 'lodash'

module.exports = new Vue({
  data () {
    return {
      name: '',
      clazz: '',
      clazzNo: 0
    }
  },
  created () {
    var vm = this
    const eventName = 'update:UserInfo'
    vm.$on(eventName, obj => {
      console.log(eventName)
      _.forIn(obj, (val, key) => {
        vm[key] = val
      })
    })
  }
})
