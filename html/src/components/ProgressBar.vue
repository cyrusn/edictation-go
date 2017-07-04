<template lang="html">
<div class="progress">
  <div
  :class=[progressBar]
  role="progressbar"
  aria-valuenow="60" aria-valuemin="0" aria-valuemax="100"
  :style='{
    width: percentage + "%"
  }'>
  </div>
</div>
</template>

<script>
export default {
  props: {
    percentage: {
      type: Number,
      required: true
    },
    progressBar: {
      type: String,
      required: true
    }

  }
}
</script>

<style lang="css">
</style>
