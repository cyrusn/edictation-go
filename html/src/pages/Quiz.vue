<template id="quiz">
<main-layout>
  <div>
    <audio :src="playerSrc" autoplay ref='player'>
    </audio>
  </div>

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

function acceptanceRatio (mode) {
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
      mode: assessment.mode,
      name: assessment.name,
      size: '',
      records: [],
      vocab: {},
      answer: '',
      playerSrc: '',
      progressPercentage: 0,
      correctPercentage: 0,
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
      vm.shuffledIndexes = _(size).range().shuffle().value()
      vm.getVocab()
      vm.$on('load:player', () => {
        console.log('load:player')
        vm.$refs.player.load()
      })
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
    }
  },
  watch: {
    runningIndex () {
      const vm = this
      // fetch infomation for new question
      vm.getVocab()
      // reset answer
      vm.answer = ''
      document.getElementById('answer').focus()
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
          vm.updatePlayer(name, index)
          vm.$emit('load:player')
        })
        .catch(err => {
          console.log(err)
        })
    },
    updatePlayer (name, index) {
      const vm = this
      vm.playerSrc = `./api/voice/assessment/${name}/index/${index}`
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

          // where mistake indicate the number of incorrect ans
          const mistake = _(vm.records)
            .filter(obj => !obj.isCorrect)
            .value()
            .length

          const acceptance = vm.size * acceptanceRatio(vm.mode)

          vm.correctPercentage = 1 - (mistake / vm.size)
          vm.hp = 1 - mistake / acceptance

          if (vm.hp <= 0) {
            vm.restart()
            return
          }

          if (vm.runningIndex < vm.size - 1) {
            vm.runningIndex += 1
            vm.progressPercentage = vm.runningIndex / vm.size
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
      vm.$refs.player.play()
    }
  }
}
</script>
