<template>
  <main-layout>
    <h1>Report <small>{{now.toDateString()}}</small></h1>
    <hr>
    <h2>{{$root.name}} <small>{{$root.clazz}} ({{$root.clazzNo}})</small></h2>
    <badge></badge>
    <h3>{{$root.assessmentName}} result</h3>
    <table class="table table-hover">
      <tbody>
        <tr>
          <th>No</th>
          <th>Title</th>
          <th>Definition</th>
          <th>Your Anwser</th>
        </tr>
        <tr v-for="v in incorretVocab">
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

  export default {
    components: {
      MainLayout,
      Badge
    },
    data () {
      return {
        incorretVocab: [],
        now: new Date()
      }
    },
    created () {
      const vm = this
      const root = vm.$root
      const name = root.assessmentName
      const records = root.assessmentRecords

      const incorrectVocabs = records.filter(obj => !obj.isCorrect)
      const incorrectVocabIndexes = incorrectVocabs.map(obj => obj.index)

      axios.post(`./api/assessment/${name}/report`, incorrectVocabIndexes)
        .then(response => {
          vm.incorretVocab = _(response.data).map((vocab, n) => {
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
