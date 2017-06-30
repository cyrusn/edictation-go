<template>
  <main-layout>
    <h1>Report</h1>
    <p>Authors: TSF, CyrusN</p>
    <h2>{{$root.name}}</h2>
    <code>
      {{$root.stat}}
    </code>
  </main-layout>
</template>

<script>
  import MainLayout from '../layouts/Main.vue'

  export default {
    components: {
      MainLayout
    }
  }
</script>
