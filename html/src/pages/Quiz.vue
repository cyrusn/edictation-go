<template id="quiz">
<main-layout>
  <h1>Assessment <small>{{name}}</small></h1>
  <hr>
  <badge
    :mode='mode'
    :percentage='correctPercentage | toPercentage'
   />
   <div class="row">
      <progress-bar
        name='Progress'
         :classArray='["progress-bar"]'
         :percentage="progressPercentage | toPercentage"
       />
       <progress-bar
         v-if="mode != 'easy'"
         name='HP'
         :classArray='["progress-bar", hpBar]'
         :percentage="hp | toPercentage"
       />
   </div>

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

import assessment from '../mixins/assessment'

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
  mixins: [assessment],
  components: {
    MainLayout,
    ProgressBar,
    Badge
  },
  data () {
    return {
      hp: 100,
      // hpBar: 'progress-bar-default',
      mode: assessment.mode,
      name: assessment.name,
      size: '',
      records: [],
      tts_source: '',
      vocab: {},
      answer: '',
      progressPercentage: 0,
      correctPercentage: 0,
      runningIndex: 0,
      shuffledIndexes: []
    }
  },
  created () {
    const vm = this
    axios.get('./api/assessment/' + vm.name + '/size')
    .then(response => {
      const size = response.data
      vm.size = size
      vm.shuffledIndexes = _(size).range().shuffle().value()
      vm.getVocab()
    })
  },
  computed: {
    hpBar () {
      const vm = this
      switch (true) {
        case vm.hp < 0.5 && vm.hp > 0.2:
          return 'progress-bar-warning'
        case vm.hp <= 0.2:
          return 'progress-bar-danger'
        default:
          return 'progress-bar-success'
      }
    },
    percentage () {
      const vm = this
      return (vm.runningIndex + 1) / vm.size
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
      vm.records = []
      vm.shuffledIndexes = _.shuffle(vm.shuffledIndexes)
      vm.runningIndex = 0
      vm.progressPercentage = 0
      vm.hp = 100
      vm.correctPercentage = 0
    },
    getVocab () {
      const vm = this
      const index = vm.shuffledIndexes[vm.runningIndex]
      const name = vm.name

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
      const name = vm.name
      const index = vm.shuffledIndexes[vm.runningIndex]
      vm.tts_source = `./api/voice/assessment/${name}/index/${index}`
    },
    next () {
      const vm = this
      const name = vm.name
      const index = vm.shuffledIndexes[vm.runningIndex]
      const answer = vm.answer.toLowerCase()

      axios.get(`./api/check/assessment/${name}/index/${index}?ans=${answer}`)
        .then(response => {
          const isCorrect = response.data
          vm.records.push({index, isCorrect, answer})

          const noOfIncorrectAns = _(vm.records)
            .filter(obj => !obj.isCorrect)
            .value()
            .length

          vm.progressPercentage = ((vm.runningIndex + 1) / vm.size)
          vm.correctPercentage = (1 - (noOfIncorrectAns / vm.size))
          vm.hp = (1 - noOfIncorrectAns / (vm.size * acceptance(vm.mode)))

          if (vm.hp <= 0) {
            vm.restart()
            return
          }

          if (vm.runningIndex < vm.size - 1) {
            vm.runningIndex += 1
          } else {
            assessment.$emit('update:assessment-report', {
              records: vm.records,
              percentage: vm.correctPercentage
            })
            vm.$root.currentRoute = '/report'
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
