<template lang="html">
  <div class="form-group">
    <div class="col-sm-offset-2 col-sm-10">
      <button class="btn btn-success" @click.prevent="submit">Start</button>
    </div>
  </div>

</template>

<script>
export default {
  props: ['submit']
}
</script>

<style lang="css">
</style>
