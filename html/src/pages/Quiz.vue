<template id="quiz">
<main-layout>
  <h1>Assessment <small>{{$root.assessmentName}}</small></h1>
  <hr>
  <badge></badge>

  <progress-bar :percentage="percentage * 100" progressBar='progress-bar'></progress-bar>
  <progress-bar v-if="$root.mode != 'easy'" :percentage="tolerantPercentage * 100" progressBar='progress-bar progress-bar-danger'></progress-bar>

  <div class="panel panel-success">
    <div class="panel-heading">
      <h1 class="panel-title">{{vocab.definition}} {{vocab.partOfSpeech}}</h1>
    </div>

    <div class="panel-body">
      <div class="input-group" @keypress.enter="next">
        <input
          placeholder="Please type your answer!"
          id="answer" type="text" class="form-control" v-model="answer">
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
/* global Audio */
import MainLayout from '../layouts/Main.vue'
import ProgressBar from '../components/ProgressBar.vue'
import Badge from '../components/Badge.vue'
import axios from 'axios'
import _ from 'lodash'

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
    Badge
  },
  data () {
    return {
      tts_source: '',
      vocab: {},
      answer: '',
      tolerantPercentage: 0,
      correctPercentage: 0,
      runningIndex: 0,
      indexes: []
    }
  },
  created () {
    const vm = this
    axios.get('./api/assessment/' + vm.$root.assessmentName + '/size')
    .then(response => {
      const size = response.data
      vm.$root.assessmentSize = size
      vm.indexes = _(size).range().shuffle().value()
      vm.getVocab()
    })
  },
  computed: {
    percentage () {
      const vm = this
      return (vm.runningIndex + 1) / vm.$root.assessmentSize
    },
    tts () {
      const vm = this
      return new Audio(vm.tts_source)
    }
  },
  watch: {
    runningIndex () {
      const vm = this
      // fetch infomation for new question
      vm.getVocab()
      vm.getVoice()

      // reset answer
      vm.answer = ''
      document.getElementById('answer').focus()
    },
    tts_source () {
      const vm = this
      setTimeout(vm.speak, 200)
    }
  },
  methods: {
    restart () {
      const vm = this
      vm.$root.assessmentRecords = []
      vm.indexes = _.shuffle(vm.indexes)
      vm.runningIndex = 0
      vm.tolerantPercentage = 0
      vm.correctPercentage = 0
    },
    getVocab () {
      const vm = this
      const index = vm.indexes[vm.runningIndex]
      const name = vm.$root.assessmentName

      axios.get(`./api/assessment/${name}/index/${index}`)
        .then(function (response) {
          vm.vocab = response.data
        })
        .catch(err => {
          console.log(err)
        })
    },
    getVoice () {
      const vm = this
      const root = vm.$root
      const name = root.assessmentName
      const index = vm.indexes[vm.runningIndex]
      vm.tts_source = `./api/voice/assessment/${name}/index/${index}`
    },
    next () {
      const vm = this
      const root = vm.$root

      const name = root.assessmentName
      const index = vm.indexes[vm.runningIndex]
      const answer = vm.answer.toLowerCase()

      axios.get(`./api/check/assessment/${name}/index/${index}?ans=${answer}`)
        .then(response => {
          const isCorrect = response.data
          root.assessmentRecords.push({index, isCorrect, answer})

          const noOfIncorrectAns = _(root.assessmentRecords)
            .filter(obj => !obj.isCorrect)
            .value()
            .length

          // updated $root.correctPercentage
          root.correctPercentage = (1 - noOfIncorrectAns / root.assessmentSize).toFixed(2) * 100

          // updated tolerantPercentage for ProgressBar
          vm.tolerantPercentage = noOfIncorrectAns / (root.assessmentSize * acceptance(root.mode))

          if (vm.tolerantPercentage > 1) {
            vm.restart()
            return
          }

          if (vm.runningIndex < root.assessmentSize - 1) {
            vm.runningIndex += 1
          } else {
            root.currentRoute = '/report'
          }
        })
        .catch(err => {
          console.log(err)
        })
    },
    speak () {
      const vm = this
      vm.tts.play()
    }
  }
}
</script>
