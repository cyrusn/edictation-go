<template lang="html">
<div>

  <div class="form-group">
    <label class="col-sm-2 control-label">Name</label>
    <div class="col-sm-10">
      <input
        id="name"
        class="form-control"
        placeholder="Name"
        v-model='name'
        @input='updateName(name)'
      >
    </div>
  </div>

  <div class="form-group">
    <label class="col-sm-2 control-label">Class</label>
    <div class=" col-sm-10">
      <select
        class="form-control"
        v-model="clazz"
        @change="updateClazz(clazz)"
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
        id="name"
        v-model="clazzNo"
        class="form-control"
        @input='updateClazzNo(clazzNo)'
      >
    </div>
  </div>
</div>
</template>

<script>
import userInfo from '../mixins/userInfo'

export default {
  mixins: [userInfo],
  data () {
    return {
      name: '',
      clazz: '',
      clazzNo: ''
    }
  },
  methods: {
    updateName (name) {
      userInfo.$emit('update:name', name)
      this.validate()
    },
    updateClazz (clazz) {
      userInfo.$emit('update:clazz', clazz)
      this.validate()
    },
    updateClazzNo (clazzNo) {
      userInfo.$emit('update:clazzNo', clazzNo)
      this.validate()
    },
    validate () {
      const vm = this
      const isValid = vm.name !== '' && vm.clazz !== '' && vm.clazzNo !== ''
      this.$emit('validate', isValid)
    }
  }
}
</script>
