<template lang="html">
<div class="progress">
  <div
  class="progress-bar"
  role="progressbar"
  aria-valuenow="60" aria-valuemin="0" aria-valuemax="100"
  :style='{
    width: percentage + "%"
  }'>
  <span class="sr-only">60% Complete</span>
  </div>
</div>
</template>

<script>
export default {
  props: {
    percentage: {
      type: Number,
      required: true
    }
  }
}
</script>

<style lang="css">
</style>
