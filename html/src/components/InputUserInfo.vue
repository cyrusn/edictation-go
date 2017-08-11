<template lang="html">
<div>
  <div class="form-group">
    <label class="col-sm-2 control-label">Name</label>
    <div class="col-sm-10">
      <input
        class="form-control"
        placeholder="Name"
        @input='onUpdateUserName'
      >
    </div>
  </div>

  <div class="form-group">
    <label class="col-sm-2 control-label">Class</label>
    <div class=" col-sm-10">
      <select
        class="form-control"
        :value="clazz"
        @change="onUpdateUserClazz"
      >
        <option value="" disabled selected>Select your class</option>
        <option v-for="clazz in ['1A','1B','1C','1D','2A','2B','2C','2D','3A','3B','3C','3D','4A','4B','4C','4D','4E','5A','5B','5C','5D','5E','6A','6B','6C','6D','6E']">
          {{clazz}}
        </option>
      </select>
    </div>
  </div>

  <div class="form-group">
    <label class="col-sm-2 control-label">Class No.</label>
    <div class="col-sm-10">
      <input
        placeholder="Class No."
        :value="clazzNo"
        class="form-control"
        @input='onUpdateUserClazzNo'
      >
    </div>
  </div>
</div>
</template>

<script>
import {mapState, mapMutations} from 'vuex'

export default {
  computed: {
    ...mapState(['user', 'clazz', 'clazzNo'])
  },
  methods: {
    ...mapMutations([
      'updateUserName',
      'updateUserClazz',
      'updateUserClazzNo'
    ]),
    onUpdateUserName (e) {
      this.updateUserName(e.target.value)
    },
    onUpdateUserClazz (e) {
      this.updateUserClazz(e.target.value)
    },
    onUpdateUserClazzNo (e) {
      this.updateUserClazzNo(e.target.value)
    }
  }
}
</script>
