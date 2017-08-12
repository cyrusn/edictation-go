<template id="home">
  <!-- TODO: use bootstrap 4.0 -->
<main-layout>
  <h1 class="text-center">
    S.K.H. Li Ping Secondary School<br><small>Online Vocabulary System</small>
  </h1>
  <hr>

  <form class="form-horizontal">
    <input-user-info />
    <input-assessment-info />
    <submit-button />
  </form>

</main-layout>
</template>

<script type="text/javascript">
import MainLayout from '../layouts/Main.vue'
import InputUserInfo from '../components/InputUserInfo.vue'
import InputAssessmentInfo from '../components/InputAssessmentInfo.vue'
import SubmitButton from '../components/SubmitButton.vue'

export default {
  components: {
    MainLayout, InputUserInfo, InputAssessmentInfo, SubmitButton
  },
  data () {
    return {
      isShowWarning: false
    }
  },
  mounted () {
    document.getElementById('name').focus()
  }
}
</script>
