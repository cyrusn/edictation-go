<template lang="html">
  <div v-if="show" class="alert alert-warning">
    <strong>Warning!</strong> Invalid user information
  </div>
</template>

<script>
export default {
  props: ['show']
}
</script>

<style lang="css">
</style>
