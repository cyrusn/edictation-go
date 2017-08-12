const Joi = require('joi')

export default {
  isValid (state) {
    const {name, clazz, clazzNo} = state.user
    const {mode} = state.assessment
    const assessmentName = state.assessment.name

    const schema = Joi.object().keys({
      name: Joi.string().min(3).required().label('Name'),
      clazz: Joi.string().required().label('Class'),
      clazzNo: Joi.number().integer().min(1).max(40).required().label('Class No.'),
      assessmentName: Joi.string().required().label('Assessment'),
      mode: Joi.string().required().label('Mode')
    })

    return Joi.validate({
      name,
      clazz,
      clazzNo,
      assessmentName,
      mode
    }, schema, {
      abortEarly: false
    }).error
  }
}
