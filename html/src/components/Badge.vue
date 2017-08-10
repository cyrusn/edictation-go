<template lang="html">
  <div class="row lead" >
    <div class="col-sm-4">
      Mode <span class="badge">{{mode | capitalize}}</span>
    </div>
    <div class="col-sm-8">
      Percentage <span class="badge">{{numerator}} / {{denominator}}</span>
    </div>
  </div>
</template>

<script>
export default {
  props: ['mode', 'percentage', 'numerator', 'denominator']
}

</script>
