<template>
  <a
    v-bind:href="href"
    v-bind:class="{ active: isActive }"
    v-on:click.prevent="go"
  >
    <slot></slot>
  </a>
</template>

<script>
  export default {
    props: {
      href: {
        type: String,
        required: true
      }
    },
    computed: {
      isActive () {
        return this.href === this.$root.currentRoute
      }
    },
    methods: {
      go (event) {
        this.$root.currentRoute = this.href
      }
    }
  }
</script>

<style scoped>
  .active {
    color: cornflowerblue;
  }
</style>
