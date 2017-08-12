<template>
  <main-layout>
    <h1>Report <small>{{assessmentName}}</small></h1>
    <hr>
    <h2>{{name}} <small>{{clazz}} ({{clazzNo}})</small></h2>
    <badge
      :mode='mode'
      :numerator='size - mistake' :denominator='size'
    />

    <h3>{{now}}</h3>
    <table class="table table-hover">
      <tbody>
        <tr>
          <th>No</th>
          <th>Title</th>
          <th>Definition</th>
          <th>Your Anwser</th>
        </tr>
        <tr v-for="v in orderedIncorretVocabs">
          <td>{{v.vocabIndex + 1}}</td>
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
        incorretVocabs: []
      }
    },
    computed: {
      ...mapState('assessment', {
        assessmentName: state => {
          console.log(state)
          return state.name
        }
      }),
      ...mapState('assessment', ['mode', 'size', 'records']),
      ...mapState('user', ['name', 'clazz', 'clazzNo']),
      ...mapGetters('assessment', [
        'mistake'
      ]),
      orderedIncorretVocabs () {
        return _.orderBy(this.incorretVocabs, 'vocabIndex')
      },
      now () {
        const now = new Date()
        return now.toDateString()
      }
    },
    mounted () {
      this.updateIncorretVocabs()
    },
    methods: {
      updateIncorretVocabs () {
        const {assessmentName, records, incorretVocabs} = this

        const incorrectVocabIndexes = records.map(obj => obj.vocabIndex)
        axios.post(`./api/assessment/${assessmentName}/report`, incorrectVocabIndexes)
          .then(response => {
            _(response.data).forEach((vocab, n) => {
              const {answer, vocabIndex} = records[n]
              vocab.answer = answer
              vocab.vocabIndex = vocabIndex
              incorretVocabs.push(vocab)
            })
          })
          .catch(console.error)
      }
    }
  }
</script>
