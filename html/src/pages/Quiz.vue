<template id="quiz">
<main-layout>
  <progress-bar :percentage="percentage"></progress-bar>
  <div class="panel panel-default">

    <div class="panel-heading">
      <h3 class="panel-title">{{definition}}</h3>
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
import ProgressBar from '../components/ProgressBar.vue'
import axios from "axios"
import _ from "lodash"

const defaultMessage = "Please type your answer!"

module.exports = {
  components: {
    MainLayout,
    ProgressBar
  },
  data() {
    const vm = this
    return {
      index: 0,
      definition: "",
      tts_source: "",
      answer: "",
      message: defaultMessage
    }
  },
  computed: {
    percentage() {
      const vm = this
      return (vm.index + 1) * 100 / vm.$root.noOfQuestions
    },
    id() {
      const vm = this
      return vm.questionIDs[vm.index]
    },
    questionIDs() {
      const vm = this
      const root = vm.$root
      const range = _.range(root.noOfQuestions)

      const result =  _(range).map(i=>i+1).shuffle().value()
      console.log(result);
      return result
    }
  },
  watch: {
    id() {
      const vm = this
      // fetch infomation for new question
      vm.fetchJSON()

      // speak the vocab after 0.5s
      vm.speak()
      
      // reset answer
      vm.answer = ""
      document.getElementById("answer").focus()
    }
  },
  created() {
    const vm = this

    window.addEventListener("keydown", function(e) {
      if (e.keyCode == 13) {
        vm.next()
      }
    })
  },
  methods: {
    fetchJSON() {
      const vm = this
      const definition = ""
      axios.get("./api/" + vm.$root.level + "/" + vm.id)
        .then(function(response) {
          const definition = response.data.definition
          vm.definition = definition
        })
        .catch(err => {
          if (err) {
            console.log(err)
          }
        })
    },
    next() {
      const vm = this
      const root = vm.$root

      const data = new FormData();
      const id = vm.id
      data.append('answer', vm.answer)
      axios.post("./check/" + root.level + "/" + id, data)
        .then(response => {
          const correct = response.data.result
          root.stat[id] = correct

          if (vm.index < root.noOfQuestions - 1) {
            vm.index += 1
          } else {
            root.currentRoute = "/report"
          }

        })
        .catch(err => {
          if (err) {
            console.log(err)
          }
        })
    },
    speak() {
      const vm = this
      const root = vm.$root
      const id = vm.id
      const level = vm.$root.level

      vm.tts_source = `/mp3/${level}/${id}`
      const tts = new Audio(vm.tts_source)
      tts.play()
    }
  }
}
</script>
