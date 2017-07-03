<template>
  <main-layout>
    <h1>Report</h1>
    <hr>
    <h2>{{$root.name}} <small>{{$root.clazz}}{{$root.clazzNo}}</small></h2>

    <mode-and-level-badge></mode-and-level-badge>
    <h3>Statics</h3>
    <code>
      {{JSON.stringify($root.stat, null, "  ")}}
    </code>
  </main-layout>
</template>

<script>
  import MainLayout from '../layouts/Main.vue'
  import ModeAndLevelBadge from '../components/ModeAndLevelBadge.vue'

  export default {
    components: {
      MainLayout,
      ModeAndLevelBadge
    }
  }
</script>
