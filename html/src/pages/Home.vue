<template id="home">
  <!-- TODO: use bootstrap 4.0 -->
<main-layout>
  <h1 class="text-center">
    S.K.H. Li Ping Secondary School<br><small>Online Vocabulary System</small>
  </h1>
  <hr>
  <warning-message :show='isShowWarning' />

  <form class="form-horizontal" @keypress.enter.prevent="onSubmit">
    <input-user-info />
    <input-assessment-info />
    <submit-button :submit='onSubmit' />
  </form>

</main-layout>
</template>

<script type="text/javascript">
import MainLayout from '../layouts/Main.vue'
import InputUserInfo from '../components/InputUserInfo.vue'
import InputAssessmentInfo from '../components/InputAssessmentInfo.vue'
import SubmitButton from '../components/SubmitButton.vue'
import WarningMessage from '../components/WarningMessage.vue'
import {mapState, mapGetters, mapMutations} from 'vuex'

export default {
  components: {
    MainLayout, InputUserInfo, InputAssessmentInfo, SubmitButton, WarningMessage
  },
  data () {
    return {
      isShowWarning: false,
      assessmentNames: []
    }
  },
  mounted () {
    document.getElementById('name').focus()
  },
  computed: {
    ...mapState(['user', 'assessment']),
    ...mapGetters(['isValid'])
  },
  methods: {
    ...mapMutations(['goto']),
    onSubmit: function (event) {
      // TODO: only user and assessment on form submit,
      // and seperate validation from store, validate values before submit.
      const {isValid, goto} = this
      this.isShowWarning = !isValid
      if (isValid) {
        goto('/quiz')
      }
    }
  }
}
</script>
