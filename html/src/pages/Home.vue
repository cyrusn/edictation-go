<template id="hello-world">
<main-layout>
  <h1 class="text-center">
    S.K.H. Li Ping Secondary School<br><small>eDictation System</small>
  </h1>
  <hr>
<!-- {{assessment.name}} -->
  <div v-if="showWarning" class="alert alert-danger" role="alert">Invalid user information</div>

  <form class="form-horizontal" @keypress.enter.prevent="submit">

    <input-user-info @validate='onValidate'></input-user-info>
    <input-assessment-info />

    <div class="form-group">
      <div class="col-sm-offset-2 col-sm-10">
        <button class="btn btn-success" @click.prevent="submit">Next</button>
      </div>
    </div>
  </form>

</main-layout>
</template>

<script type="text/javascript">
import MainLayout from '../layouts/Main.vue'
import InputUserInfo from '../components/InputUserInfo.vue'
import InputAssessmentInfo from '../components/InputAssessmentInfo.vue'
import userInfo from '../mixins/userInfo'

export default {
  mixins: [userInfo],
  data () {
    return {
      // assessment,
      isValid: false,
      showWarning: false,
      assessmentNames: []
    }
  },
  mounted () {
    document.getElementById('name').focus()
  },
  methods: {
    submit: function (event) {
      const vm = this
      const root = vm.$root

      vm.showWarning = !vm.isValid
      if (vm.isValid) {
        root.currentRoute = '/quiz'
      }
    },
    onValidate (isValid) {
      this.isValid = isValid
    }
  },
  components: {
    MainLayout, InputUserInfo, InputAssessmentInfo
  }
}
</script>
