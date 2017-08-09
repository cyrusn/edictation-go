<template>
  <div>
    <div class="container">
      <slot></slot>
    </div>
  </div>
</template>

<script>
import VLink from '../components/VLink.vue'

export default {
  components: { VLink }
}

</script>

<style scoped>
.container {
  max-width: 600px;
  margin: 0 auto;
  padding-top:12px;
}
</style>
