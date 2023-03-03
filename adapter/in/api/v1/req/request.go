package req


// Id   uint  `binding:"required" form:"user" json:"user" xml:"user"  time_format:"2006-01-02"` `form:"colors[]"`

// if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
// 	v.RegisterValidation("bookabledate", bookableDate)
// }

// binding:"required,bookabledate"

// var bookableDate validator.Func = func(fl validator.FieldLevel) bool {
// 	date, ok := fl.Field().Interface().(time.Time)
// 	if ok {
// 		today := time.Now()
// 		if today.After(date) {
// 			return false
// 		}
// 	}
// 	return true
// }