const googleTTS = require('google-tts-api');

const argv = require('yargs')
  .option('title', {
    alias: 't',
    default: "hello world",
    describe: 'choose a title',
    type: 'string'
  })
  .option('speed', {
    alias: 's',
    describe: 'choose a speed',
    default: 0.24,
    type: 'number'
  })
  .help()
  .argv

const title = argv.title
const speed = argv.speed

googleTTS(title, "en", speed)
.then(url => {
  process.stdout.write(url)
})
.catch(url => {
  process.stderr.write(url)
})
