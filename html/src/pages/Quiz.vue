<template id="quiz">
<main-layout>
  <div class="panel panel-default">
    <code>{{$root.stat}}</code>
    <div class="panel-heading">
      <h3 class="panel-title">{{definition}}</h3>
      <h1>{{$root.noOfQuestions}}</h1>
      <code>
        {{questionIDs}}
      </code>
      <p>{{id}}</p>
      <p>{{tts_source}}</p>
    </div>

    <div class="panel-body">
      <div class="input-group">
        <input id="answer" type="text" class="form-control" v-model="answer">
        <div class="input-group-btn">
          <button class="btn btn-default" @click="next">
          <span class="glyphicon glyphicon-thumbs-up" aria-hidden="true"></span>
        </button>
          <button class="btn btn-default" @click="speak">
          <span class="glyphicon glyphicon-volume-up" aria-hidden="true"></span>
        </button>
        </div>
      </div>
    </div>

    <div class="panel-footer" role="alert">{{message}}</div>
  </div>
</main-layout>
</template>

<script>
import MainLayout from '../layouts/Main.vue'
import axios from "axios"
import _ from "lodash"

const defaultMessage = "Please type your answer!"

module.exports = {
  components: {
    MainLayout
  },
  data: function() {
    var vm = this
    return {
      index: 0,
      definition: "",
      tts_source: "",
      answer: "",
      message: defaultMessage
    }
  },
  computed: {
    id () {
      var vm = this
      return vm.questionIDs[vm.index]
    },
    questionIDs () {
      var vm = this
      var range = _.range(vm.$root.noOfQuestions)
      return _.shuffle(range)
    }
  },
  watch: {
    id () {
      var vm = this
      vm.fetchJSON()
      setTimeout(function() {
        this.speak()
      }.bind(vm), 500)
      vm.answer = ""
      document.getElementById("answer").focus()
    }
  },
  created: function() {
    var vm = this

    window.addEventListener("keydown", function(e) {
      if (e.keyCode == 13) {
        vm.next()
      }
    })
  },
  methods: {
    fetchJSON() {
      var vm = this
      var definition = ""
      axios.get("./api/" + vm.$root.level + "/" + vm.id)
        .then(function(response) {
          const definition = response.data.definition
          vm.definition = definition
        })
        .catch(function(err) {
          if (err) {
            console.log(err)
          }
        })
    },
    next() {
      var vm = this
      const data = new FormData();
      const id = vm.id
      data.append('answer', vm.answer)
      axios.post("./check/" + vm.$root.level + "/" + id, data)
        .then(function(response) {
          const isCorrect = response.data.result
          vm.$root.stat[id] = isCorrect

          if (vm.index < vm.$root.noOfQuestions - 1) {
            vm.index += 1
          } else {
            vm.$root.currentRoute = "/report"
          }

          if (isCorrect) {
            vm.message = defaultMessage
            vm.answer = ""
          } else {
            vm.message = "Wrong ans, please try again!"
          }
        })
        .catch(function(err) {
          if (err) {
            console.log(err)
          }
        })
    },
    speak() {
      var vm = this
      const level = vm.$root.level
      const id = vm.id

      vm.tts_source = `/mp3/${level}/${id}`
      var tts = new Audio(vm.tts_source)
      tts.play()
    }
  }
}
</script>
