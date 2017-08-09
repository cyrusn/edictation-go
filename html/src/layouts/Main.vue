<template>
  <div>
    <div class="container">
      <slot></slot>
    </div>
  </div>
</template>

<script>
  export default {}
</script>

<style scoped>
.container {
  max-width: 600px;
  margin: 0 auto;
  padding-top:12px;
}
</style>
