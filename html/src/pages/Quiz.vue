<template id="quiz">
<main-layout>
  <player ref='player' />
  <h1>Assessment <small>{{assessment.name}}</small></h1>
  <hr>
  <badge :mode='assessment.mode' :numerator='assessment.size - mistake' :denominator='assessment.size' />
  <div class="row">
    <progress-bar name='Progress' :percentage="progressPercentage | toPercentage" />
    <progress-bar v-if="assessment.mode != 'easy'" name='HP' :contextualColor='contextualColor' :percentage="hp | toPercentage" />
  </div>

  <quiz-panel :index='assessment.runningIndex + 1' :vocab='assessment.vocab' :next='next' @updateAnswer='updateAnswer' :speak='speak' />
</main-layout>
</template>

<script>
import MainLayout from '../layouts/Main.vue'
import ProgressBar from '../components/ProgressBar.vue'
import Badge from '../components/Badge.vue'
import Player from '../components/Player.vue'
import QuizPanel from '../components/QuizPanel.vue'

import {mapState, mapMutations, mapActions, mapGetters} from 'vuex'

export default {
  components: {
    MainLayout, ProgressBar, Badge, Player, QuizPanel},
  data () {
    return {
      answer: ''
    }
  },
  computed: {
    ...mapState(['assessment']),
    ...mapGetters(['mistake', 'index', 'hp', 'progressPercentage']),
    contextualColor () {
      const {hp} = this
      switch (true) {
        case hp < 0.4 && hp > 0.2:
          return 'warning'
        case hp <= 0.1:
          return 'danger'
        default:
          return 'success'
      }
    }
  },
  methods: {
    ...mapMutations([
      'goto', 'increment', 'resetIndex', 'clearAssessmentRecords'
    ]),
    ...mapActions([
      'addRecords', 'getOrders', 'getSize', 'getVocab'
    ]),
    updateAnswer (answer) {
      this.answer = answer.toLowerCase()
    },
    restart () {
      const {
        speak, getVocab, getOrders, resetIndex, clearAssessmentRecords
      } = this

      resetIndex()
      clearAssessmentRecords()

      getOrders()
      .then(getVocab)
      .then(speak)
    },
    speak () {
      const player = this.$refs.player
      player.play()
    },
    next (answer) {
      const {
        addRecords, decideNextStep, increment
      } = this

      addRecords(answer)
        .then(increment)
        .then(decideNextStep)
    },
    decideNextStep () {
      const {
        hp, assessment, restart, getVocab, goto, speak
      } = this
      const isWrongToMuch = hp < 0
      const {size, runningIndex} = assessment
      switch (true) {
        case isWrongToMuch:
          restart()
          break
        case runningIndex < size:
          getVocab()
          .then(speak)
          break
        default:
          goto('/report')
      }
    }
  },
  mounted () {
    const {
      assessment, getSize, getOrders, getVocab, speak
    } = this
    const {name} = assessment

    getSize(name)
      .then(getOrders)
      .then(getVocab)
      .then(speak)
  }
}
</script>
