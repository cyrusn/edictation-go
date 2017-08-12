<template id="quiz">
<main-layout>
  <h1>Assessment <small>{{name}}</small></h1>
  <hr>
  <badge
    :mode='mode'
    :numerator='size - mistake'
    :denominator='size'
  />

  <div class="row">
    <progress-bar
      name='Progress'
      :percentage="progressPercentage | toPercentage"
    />
    <progress-bar
      v-if="mode != 'easy'"
      name='HP'
      :contextualColor='contextualColor'
      :percentage="hp | toPercentage"
    />
  </div>
  <quiz-panel />
</main-layout>
</template>

<script>
import MainLayout from '../layouts/Main.vue'
import Badge from '../components/Badge.vue'
import ProgressBar from '../components/ProgressBar.vue'
import QuizPanel from '../components/QuizPanel.vue'

import {mapState, mapGetters} from 'vuex'

export default {
  components: {
    MainLayout, ProgressBar, Badge, QuizPanel
  },
  computed: {
    ...mapState('assessment', [
      'mode', 'size'
    ]),
    ...mapGetters('assessment', [
      'contextualColor',
      'hp',
      'mistake',
      'progressPercentage'
    ])
  }
}
</script>
