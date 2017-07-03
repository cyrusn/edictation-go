<template>
  <div>
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css" integrity="sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u" crossorigin="anonymous">
    <div class="container">
      <!-- <v-link href="/">Home</v-link> -->
      <slot></slot>
    </div>
  </div>
</template>

<script>
import VLink from '../components/VLink.vue'

export default {
  components: {
    VLink
  }
}
</script>

<style scoped>
.container {
  max-width: 600px;
  margin: 0 auto;
  padding-top:12px;
}
</style>
