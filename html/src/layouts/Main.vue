<template>
<div class="container" style="padding-top:12px">
  <!-- <v-link href="/">Home</v-link> -->
  <!-- <ul>
      <li>
        <v-link href="/quiz">Quiz</v-link>
        <v-link href="/about">About</v-link>
      </li>
    </ul> -->
  <slot></slot>
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
  max-width: 800px;
  margin: 0 auto;
}
</style>
