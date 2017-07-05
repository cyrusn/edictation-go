<template>
  <main-layout>
    <h1>Report</h1>
    <hr>
    <h2>{{$root.name}} <small>{{$root.clazz}}{{$root.clazzNo}}</small></h2>

    <mode-and-level-badge></mode-and-level-badge>
    <h3>Statics</h3>

    <table class="table table-hover">
      <tbody>
        <tr>
          <th>Title</th>
          <th>Definition</th>
          <th>Your Anwser</th>
        </tr>
        <tr v-for="v in allIncorrectVocabs">
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
  import ModeAndLevelBadge from '../components/ModeAndLevelBadge.vue'
  import _ from 'lodash'
  import axios from 'axios'

  export default {
    components: {
      MainLayout,
      ModeAndLevelBadge
    },
    data () {
      return {
        allIncorrectVocabs: []
      }
    },
    computed: {
      resultPercentage () {
        const vm = this
        const root = vm.$root

        const noOfWrongAnswers = _(root.stat)
          .map(val => val)
          .filter(val => !val)
          .value()
          .length
        return noOfWrongAnswers / (root.noOfQuestions)
      }
    },
    mounted () {
      const vm = this
      const root = vm.$root

      const ids = _(root.stat)
        .omitBy(obj => {
          return obj.isCorrect
        })
        .map((val, key) => parseInt(key))
        .value()

      axios.post('./api/vocabs', ids)
        .then(response => {
          vm.allIncorrectVocabs = response.data.map(vocab => {
            const wrongAns = root.stat[vocab.id].answer
            vocab.answer = wrongAns
            return vocab
          })
        })
        .catch(err => {
          if (err) {
            console.log(err)
          }
        })
    }
  }
</script>
