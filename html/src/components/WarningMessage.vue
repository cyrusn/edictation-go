<template lang="html">
  <div v-if="show" class="alert alert-danger">
    <div v-for='message in messages'>
      <strong>Warning!</strong> &nbsp{{message}}
    </div>
  </div>
</template>

<script>
export default {
  props: ['messages', 'show']
}
</script>

<style lang="css">
</style>
