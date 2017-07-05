const googleTTS = require('google-tts-api');
const express = require('express');
const morgan = require('morgan')
const app = express();
const logger = morgan('common')

const yargs = require('yargs')
const argv = require('yargs')
  .option('port', {
    alias: 'p',
    default: 5001,
    describe: 'choose a port',
    type: 'number'
  })
  .option('speed', {
    alias: 's',
    describe: 'choose a speed',
    default: 0.24,
    type: 'number'
  })
  .help()
  .argv

const port = argv.port
const speed = argv.speed

app.use(logger)

app.get('/voice/:title', function(req, res){
  const title = req.params.title

  googleTTS(title, "en", speed)
  // speed normal = 1 (default), slow = 0.24
  .then(function (url) {
    // https://translate.google.com/translate_tts?...
    res.send(url);
  })
  .catch(function (err) {
    res.send(err.stack);
  });
});

console.log(`Voice server start on http://localhost:${port}`)
console.log(`speed = ${speed}`)
app.listen(port);
