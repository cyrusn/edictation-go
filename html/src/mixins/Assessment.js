import Vue from 'vue'

module.exports = new Vue({
  data () {
    return {
      name: '',
      mode: 'normal',
      size: 0,
      records: []
    }
  },
  created () {
    const vm = this
    const names = ['name', 'mode', 'records', 'size']

    names.forEach(name => {
      const eventName = 'update:assessment-' + name
      vm.$on(eventName, value => {
        console.log(eventName)
        vm[name] = value
      })
    })
  }
})
