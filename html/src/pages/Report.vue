<template>
  <main-layout>
    <h1>Report <small>{{assessment.name}}</small></h1>
    <hr>
    <h2>{{user.name}} <small>{{user.clazz}} ({{user.clazzNo}})</small></h2>
    <badge
      :mode='assessment.mode'
      :numerator='assessment.size - mistake' :denominator='assessment.size'
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

  import {mapState, mapGetters} from 'vuex'

  export default {
    components: {
      MainLayout,
      Badge
    },
    data () {
      return {
        incorretVocabs: [],
        now: new Date()
      }
    },
    computed: {
      ...mapState(['assessment', 'user']),
      ...mapGetters(['mistake'])
    },
    created () {
      const {assessment, incorretVocabs} = this
      const {name, records} = assessment

      const incorrectVocabIndexes = records.map(obj => obj.index)
      axios.post(`./api/assessment/${name}/report`, incorrectVocabIndexes)
        .then(response => {
          _(response.data).forEach((vocab, n) => {
            const {answer, index} = records[n]
            vocab.answer = answer
            vocab.index = index
            incorretVocabs.push(vocab)
          })
        })
        .catch(console.error)
    }
  }
</script>
