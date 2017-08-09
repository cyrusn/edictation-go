<template id="hello-world">
<main-layout>
  <h1 class="text-center">
    S.K.H. Li Ping Secondary School<br><small>eDictation System</small>
  </h1>
  <hr>

  <warning-message :show='isShowWarning' />

  <form class="form-horizontal" @keypress.enter.prevent="onSubmit">
    <input-user-info @validate='onValidate'></input-user-info>
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

import userInfo from '../mixins/userInfo'

export default {
  components: {
    MainLayout, InputUserInfo, InputAssessmentInfo, SubmitButton, WarningMessage
  },
  mixins: [userInfo],
  data () {
    return {
      isValid: false,
      isShowWarning: false,
      assessmentNames: []
    }
  },
  mounted () {
    document.getElementById('name').focus()
  },
  methods: {
    onSubmit: function (event) {
      const vm = this

      vm.isShowWarning = !vm.isValid
      if (vm.isValid) {
        vm.$root.currentRoute = '/quiz'
      }
    },
    onValidate (isValid) {
      this.isValid = isValid
    }
  }
}
</script>
