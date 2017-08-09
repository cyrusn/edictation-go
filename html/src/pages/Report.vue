<template>
  <main-layout>
    <h1>Report <small>{{assessment.name}}</small></h1>
    <hr>
    <h2>{{userInfo.name}} <small>{{userInfo.clazz}} ({{userInfo.clazzNo}})</small></h2>
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

  import userInfo from '../mixins/userInfo'
  import assessment from '../mixins/assessment'

  export default {
    mixins: [userInfo, assessment],
    components: {
      MainLayout,
      Badge
    },
    data () {
      return {
        userInfo,
        assessment,
        incorretVocabs: [],
        now: new Date()
      }
    },
    created () {
      const vm = this
      const name = assessment.name
      const records = assessment.report.records

      const incorrectVocabs = records.filter(obj => !obj.isCorrect)
      const incorrectVocabIndexes = incorrectVocabs.map(obj => obj.index)
      axios.post(`./api/assessment/${name}/report`, incorrectVocabIndexes)
        .then(response => {
          vm.incorretVocabs = _(response.data).map((vocab, n) => {
            vocab.answer = incorrectVocabs[n].answer
            vocab.index = incorrectVocabs[n].index
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
