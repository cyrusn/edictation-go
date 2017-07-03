var googleTTS = require('google-tts-api');
var express = require('express');
var morgan = require('morgan')
var app = express();

var logger = morgan('common')

app.use(logger)

app.get('/voice/:title', function(req, res){
  const title = req.params.title

  googleTTS(title, "en", 1)
  // speed normal = 1 (default), slow = 0.24
  .then(function (url) {
    // https://translate.google.com/translate_tts?...
    res.send(url);
  })
  .catch(function (err) {
    res.send(err.stack);
  });
});

const port = 4000
console.log(`Voice server start on http://localhost:${port}`)
app.listen(port);
