<template id="quiz">
<main-layout>
  <h1>Quiz</h1>
  <hr>
  <mode-and-level-badge></mode-and-level-badge>
  <progress-bar :percentage="percentage * 100" progressBar='progress-bar'></progress-bar>
  <progress-bar v-if="$root.mode != 'easy'" :percentage="wrongPercentage * 100" progressBar='progress-bar progress-bar-danger'></progress-bar>
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
/*global Audio, FormData*/
import MainLayout from '../layouts/Main.vue'
import ProgressBar from '../components/ProgressBar.vue'
import ModeAndLevelBadge from '../components/ModeAndLevelBadge.vue'
import axios from 'axios'
import _ from 'lodash'

const defaultMessage = 'Please type your answer!'

function acceptance (mode) {
  switch (mode) {
    case 'normal':
      return 0.5
    case 'hard':
      return 0.25
    default:
      return 1
  }
}

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
      wrongPercentage: 0,
      message: defaultMessage
    }
  },
  computed: {
    id () {
      const vm = this
      return vm.$root.questionIDs[vm.index]
    },
    percentage () {
      const vm = this
      return (vm.index + 1) / vm.$root.noOfQuestions
    },
    tts () {
      const vm = this
      return new Audio(vm.tts_source)
    }
  },
  watch: {
    id () {
      const vm = this
      // fetch infomation for new question
      vm.fetchAPI()

      vm.tts_source = `./voice/id/${vm.id}`
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
    fetchAPI () {
      const vm = this

      axios.get('./api/id/' + vm.id)
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
      const id = vm.id
      const answer = vm.answer.toLowerCase()

      axios.post('./check/id/' + id, {answer})
        .then(response => {
          const isCorrect = response.data
          root.stat[id] = { isCorrect, answer }

          const noOfWrongAnswers = _(root.stat)
            .map(val => val.isCorrect)
            .filter(val => !val)
            .value()
            .length

          vm.wrongPercentage = noOfWrongAnswers / (root.noOfQuestions * acceptance(root.mode))

          if (vm.wrongPercentage > 1) {
            window.location.pathname = '/'
          } else if (vm.index < root.noOfQuestions - 1) {
            vm.index += 1
          } else {
            root.currentRoute = '/report'
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
