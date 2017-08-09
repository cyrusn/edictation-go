<template lang="html">
<div class="form-group">
  <label class="col-sm-2 control-label">{{name}}</label>
  <div class="col-sm-10">
    <div class="progress ">
      <div
      role="progressbar"
      :aria-valuenow='percentage'
      aria-valuemin="0"
      aria-valuemax="100"
      :class='classArray'
      :style='{width: percentage + "%"}'
      />
    </div>
  </div>
</div>
</template>

<script>
export default {
  props: {
    name: '',
    percentage: {
      type: Number,
      required: true
    },
    progressBar: {
      type: String
    },
    classArray: {
      type: Array
    }
  }
}
</script>

<style lang="css">
</style>
