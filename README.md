# eDictation Spec

## How to use?
- build javascript for html
``` sh
cd html
npm install
npm run build
# or npm run watchify for development
```

- run server app
``` sh
go run main.go
```

## Use case

### Student
- student should enter his / her name, class and class number for print the report.

- the webpage should show the Chinese definition and pronounce the english word one by one in random order

- student can select `number of questions`, `level` and `mode`
  - `level`
    - *level 1* to *level 6* for form 1 to form 6 students
  - `mode`
    - *easy mode*:
      - test won't restart even there are too many mistakes.
    - *normal mode*:
      - test will restart if the incorrect percentage is larger then 50%
    - *hard mode*:
      - test will restart if the incorrect percentage is larger then 25%

- student should enter the english vocabulary, then
  - if the answer is correct, then next question will be shown.
  - else, it will show the correct vocabulary and then next question will be shown.

- when the test finish, student can print out the report that should have following contents:
  - student's name, class and class number
  - test level, mode
  - date of test
  - percentage
  - if `hard mode` is selected, a list will be included in the report to show students the vocabulary items that students failed to input correctly.

### admin
- admin can update the question bank in `JSON` format, if developer have sufficient time, webpage interface should be used to update the question bank, of course, login system are also required.

# reference
- [Writing Web Applications - The Go Programming Language](https://golang.org/doc/articles/wiki/)
- [Load & Parse Json File in Golang](https://www.chazzuka.com/2015/03/load-parse-json-file-golang/)
- [What Is JSON? — Squarespace Developers](https://developers.squarespace.com/what-is-json/)
- [template - The Go Programming Language](https://golang.org/pkg/html/template/)
- [http - What exactly is RESTful programming? - Stack Overflow](https://stackoverflow.com/questions/671118/what-exactly-is-restful-programming)
- [Vue.js](https://vuejs.org/)
- [gorilla/mux: A powerful URL router and dispatcher for golang.](https://github.com/gorilla/mux)
- [Text to Speech | TTS SDK | Speech Recognition (ASR)](https://www.ispeech.org/)
- [google-tts-api](https://www.npmjs.com/package/google-tts-api)
