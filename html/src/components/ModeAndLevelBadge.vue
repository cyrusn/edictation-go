<template lang="html">
  <div class="row lead" >
    <div class="col-sm-4">
      Level <span class="badge">{{$root.level}}</span>
    </div>
    <div class="col-sm-4">
      Mode <span class="badge">{{$root.mode}}</span>
    </div>
    <div class="col-sm-4">
      Percentage <span class="badge">{{$root.correctPercentage}}%</span>
    </div>
  </div>
</template>
