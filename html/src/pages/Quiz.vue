<template id="quiz">
<main-layout>
  <h1>Quiz</h1>
  {{$root.stat}}
  <hr>
  <mode-and-level-badge></mode-and-level-badge>
  <progress-bar :percentage="percentage"></progress-bar>

  <div class="panel panel-success">
    <div class="panel-heading">
      <h1 class="panel-title">{{definition}} {{partOfSpeech}}</h1>
    </div>

    <div class="panel-body">
      <div class="input-group" @keypress.enter="next">
        <input id="answer" type="text" class="form-control" v-model="answer">
        <div class="input-group-btn">
          <button class="btn btn-default" @click.prevent="next">
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
import ModeAndLevelBadge from '../components/ModeAndLevelBadge.vue'
import routes from '../routes.js'
import axios from 'axios'
import _ from 'lodash'
import config from '../config'
const without = config.without

const defaultMessage = 'Please type your answer!'

export default {
  components: {
    MainLayout,
    ProgressBar,
    ModeAndLevelBadge
  },
  data () {
    return {
      index: 0,
      definition: '',
      partOfSpeech: '',
      tts_source: '',
      answer: '',
      message: defaultMessage
    }
  },
  computed: {
    id () {
      const vm = this
      return vm.questionIDs[vm.index]
    },
    percentage () {
      const vm = this
      return (vm.index + 1) * 100 / vm.$root.noOfQuestions
    },
    questionIDs () {
      const vm = this
      const root = vm.$root

      const range = _.range(root.noOfQuestions)
      return _(range).map(i => i + 1).shuffle().value()
    },
    tts () {
      const vm = this
      return new Audio(vm.tts_source)
    }
  },
  watch: {
    id () {
      const vm = this
      const root = vm.$root

      // fetch infomation for new question
      vm.fetchJSON()

      vm.tts_source = `${without}/voice/${root.level}/${vm.id}`
      // reset answer
      vm.answer = ''
      document.getElementById('answer').focus()
    },
    tts_source () {
      const vm = this
      setTimeout(vm.speak, 500)
    }
  },
  methods: {
    fetchJSON () {
      const vm = this

      axios.get('.' + without + '/api/' + vm.$root.level + '/' + vm.id)
        .then(function (response) {
          const definition = response.data.definition
          const partOfSpeech = response.data.partOfSpeech

          vm.definition = definition
          vm.partOfSpeech = partOfSpeech
        })
        .catch(err => {
          if (err) {
            console.log(err)
          }
        })
    },
    next () {
      const vm = this
      const root = vm.$root

      const data = new FormData()
      const id = vm.id
      data.append('answer', vm.answer)
      axios.post('.' + without + '/check/' + root.level + '/' + id, data)
        .then(response => {
          const correct = response.data.result
          root.stat[id] = correct

          if (vm.index < root.noOfQuestions - 1) {
            vm.index += 1
          } else {
            root.currentRoute = '/report'
            window.history.pushState(
              null,
              routes['/report'],
              without + '/report'
            )
          }
        })
        .catch(err => {
          if (err) {
            console.log(err)
          }
        })
    },
    speak () {
      const vm = this
      vm.tts.play()
    }
  }
}
</script>
