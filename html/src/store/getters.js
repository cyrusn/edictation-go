const _ = require('lodash')

const requiredString = (name) => {
  return (val) => {
    if (val === undefined || val === null || val.length === 0) {
      return {
        message: `"${name}" is required`,
        isValid: false
      }
    }
    return { isValid: true }
  }
}

const validators = {
  name (name) {
    if (name === undefined || name === null || name.length === 0) {
      return {
        message: `"Name" is required.`,
        isValid: false
      }
    }

    if (name.length < 3) {
      return {
        message: `Name must be at least 3 characters long.`,
        isValid: false
      }
    }
    return { isValid: true }
  },
  clazz: requiredString('Class'),
  clazzNo (n) {
    if (n === undefined || n === null || n === '') {
      return {
        message: `"Class No." is required.`,
        isValid: false
      }
    }

    const clazzNo = Number(n)

    if (typeof clazzNo === 'string') {
      return {
        message: `"Class No." must be an integer.`,
        isValid: false
      }
    }

    if (!Number.isInteger(clazzNo)) {
      return {
        message: `"Class No." must be an integer.`,
        isValid: false
      }
    }

    if (clazzNo < 0) {
      return {
        message: `"Class No." must be larger than or equal to 1.`,
        isValid: false
      }
    }

    if (clazzNo > 40) {
      return {
        message: `"Class No." must be less than or equal to 40.`,
        isValid: false
      }
    }
    return { isValid: true }
  },
  assessmentName: requiredString('Assessment'),
  mode: requiredString('Mode')
}

function validate (obj) {
  const messages = []
  _.forIn(obj, (val, key) => {
    const {isValid, message} = validators[key](val)
    if (!isValid) {
      messages.push(message)
    }
  })

  return messages
}

const getters = {
  validateMessages (state) {
    const {name, clazz, clazzNo} = state.user
    const {mode} = state.assessment
    const assessmentName = state.assessment.name

    return validate({
      name,
      clazz,
      clazzNo,
      assessmentName,
      mode
    })
  }
}

export default getters
