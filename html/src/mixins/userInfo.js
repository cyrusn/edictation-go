import Vue from 'vue'

module.exports = new Vue({
  data () {
    return {
      name: '',
      clazz: '',
      clazzNo: 0
    }
  },
  created () {
    const vm = this

    const names = ['name', 'clazz', 'clazzNo']

    names.forEach(name => {
      const eventName = 'update:' + name
      vm.$on(eventName, value => {
        console.log(eventName)
        vm[name] = value
      })
    })
  }
})
