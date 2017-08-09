<template>
  <main-layout>
    <h1>Report <small>{{assessmentName}}</small></h1>
    <hr>
    <h2>{{name}} <small>{{clazz}} ({{clazzNo}})</small></h2>
    <badge
      :mode='assessment.mode'
      :percentage='assessment.report.percentage | toPercentage'
    />
    <h3>{{now.toDateString()}}</h3>
    <table class="table table-hover">
      <tbody>
        <tr>
          <th>No</th>
          <th>Title</th>
          <th>Definition</th>
          <th>Your Anwser</th>
        </tr>
        <tr v-for="v in incorretVocabs">
          <td>{{v.index + 1}}</td>
          <td>{{v.title}} ({{v.partOfSpeech}})</td>
          <td>{{v.definition}}</td>
          <td class="text text-danger">{{v.answer}}</td>
        </tr>
      </tbody>
    </table>
  </main-layout>
</template>

<script>
  import MainLayout from '../layouts/Main.vue'
  import Badge from '../components/Badge.vue'
  import _ from 'lodash'
  import axios from 'axios'

  import UserInfo from '../mixins/UserInfo'
  import Assessment from '../mixins/Assessment'

  export default {
    mixins: [UserInfo, Assessment],
    components: {
      MainLayout,
      Badge
    },
    data () {
      return {
        name: UserInfo.name,
        clazz: UserInfo.clazz,
        clazzNo: UserInfo.clazzNo,
        assessmentName: Assessment.name,
        records: Assessment.report.records,
        incorretVocabs: [],
        now: new Date()
      }
    },
    created () {
      const vm = this
      const name = vm.assessmentName
      const records = vm.records

      const incorrectVocabIndexes = records.map(obj => obj.index)
      axios.post(`./api/assessment/${name}/report`, incorrectVocabIndexes)
        .then(response => {
          vm.incorretVocabs = _(response.data).map((vocab, n) => {
            vocab.answer = records[n].answer
            vocab.index = records[n].index
            return vocab
          })
          .orderBy('index')
          .value()
        })
        .catch(err => {
          console.log(err)
        })
    }
  }
</script>
