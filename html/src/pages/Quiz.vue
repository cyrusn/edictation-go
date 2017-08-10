<template id="quiz">
<main-layout>
  <player ref='player' />
  <h1>Assessment <small>{{name}}</small></h1>
  <hr>
  <badge :mode='mode' :numerator='size - mistake' :denominator='size' />
  <div class="row">
    <progress-bar name='Progress' :percentage="progressPercentage | toPercentage" />
    <progress-bar v-if="mode != 'easy'" name='HP' :contextualColor='contextualColor' :percentage="hp | toPercentage" />
  </div>

  <quiz-panel :vocab='vocab' :next='next' @updateAnswer='updateAnswer' :speak='speak' />
</main-layout>
</template>

<script>
import axios from 'axios'
import _ from 'lodash'

import MainLayout from '../layouts/Main.vue'
import ProgressBar from '../components/ProgressBar.vue'
import Badge from '../components/Badge.vue'
import Player from '../components/Player.vue'
import QuizPanel from '../components/QuizPanel.vue'

import Assessment from '../mixins/Assessment'
import Router from '../mixins/Router'

export default {
  mixins: [Assessment],
  components: {
    MainLayout,
    ProgressBar,
    Badge,
    Player,
    QuizPanel
  },
  data () {
    return {
      hp: 100,
      size: 0,
      mode: Assessment.mode,
      name: Assessment.name,
      records: [],
      vocab: {},
      mistake: '',
      runningIndex: 0,
      shuffledIndexes: []
    }
  },
  mounted () {
    const vm = this
    axios.get('./api/assessment/' + vm.name + '/size')
      .then(response => {
        const size = response.data
        vm.size = size
        Assessment.$emit('update:assessment-size', size)
        vm.shuffledIndexes = _(size).range().shuffle().value()
        vm.getVocab()
      })
  },
  computed: {
    index () {
      const vm = this
      return vm.shuffledIndexes[vm.runningIndex]
    },
    contextualColor () {
      const vm = this
      switch (true) {
        case vm.hp < 0.5 && vm.hp > 0.2:
          return 'warning'
        case vm.hp <= 0.2:
          return 'danger'
        default:
          return 'success'
      }
    },
    progressPercentage () {
      const vm = this
      return vm.runningIndex / vm.size
    }
  },
  methods: {
    updateAnswer (answer) {
      this.answer = answer.toLowerCase()
    },
    restart () {
      const vm = this
      vm.shuffledIndexes = _.shuffle(vm.shuffledIndexes)
      vm.records = []
      vm.runningIndex = 0
      vm.progressPercentage = 0
      vm.mistake = 0
      vm.hp = 100
    },
    getVocab () {
      const vm = this
      const {index, name} = vm

      axios.get(`./api/assessment/${name}/index/${index}`)
        .then(function (response) {
          vm.vocab = response.data
          const player = vm.$refs.player
          player.$emit('updateAudioSource', {name, index})
        })
        .catch(console.error)
    },
    speak () {
      const player = this.$refs.player
      player.play()
    },
    next () {
      const vm = this
      const {name, index, answer} = vm

      axios.get(`./api/check/assessment/${name}/index/${index}?ans=${answer}`)
        .then(response => {
          const isCorrect = response.data
          vm.updateRecords(isCorrect, index, answer)
          vm.updateProgressBar()
          vm.decideNextStep()
        })
        .catch(console.error)
    },
    updateRecords (isCorrect, index, answer) {
      if (!isCorrect) {
        this.records.push({index, answer})
      }
    },
    updateProgressBar () {
      const vm = this

      // where mistake indicate the number of incorrect ans
      const mistake = vm.records.length
      vm.mistake = mistake

      var acceptanceRatio
      switch (vm.mode) {
        case 'normal':
          acceptanceRatio = 0.5
          break
        case 'hard':
          acceptanceRatio = 0.25
          break
        default:
          acceptanceRatio = 1
      }

      const acceptance = vm.size * acceptanceRatio
      vm.hp = 1 - mistake / acceptance
    },
    updateReport () {
      const {records} = this
      Assessment.$emit('update:assessment-records', records)
    },
    decideNextStep () {
      const vm = this
      const isWrongToMach = vm.hp < 0
      const isLastQuestion = vm.runningIndex < vm.size - 1

      switch (true) {
        case isWrongToMach:
          vm.restart()
          break
        case isLastQuestion:
          vm.runningIndex += 1
          vm.getVocab()
          break
        default:
          vm.updateReport()
          Router.$emit('update:route', '/report')
      }
    }
  }
}
</script>
