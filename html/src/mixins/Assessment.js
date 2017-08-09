import Vue from 'vue'

module.exports = new Vue({
  data () {
    return {
      name: '',
      mode: 'normal',
      report: {
        percentage: 0,
        records: []
      }
    }
  },
  created () {
    const vm = this
    const names = ['name', 'mode', 'report']

    names.forEach(name => {
      const eventName = 'update:assessment-' + name
      vm.$on(eventName, value => {
        console.log(eventName)
        vm[name] = value
      })
    })
  }
})
