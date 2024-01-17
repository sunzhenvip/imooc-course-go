package forms

type PassWordLoginForm struct {
	Mobile   string `form:"mobile" json:"mobile" binding:"required,mobile"` // 手机号码自定义规范
	PassWord string `form:"password" json:"password" binding:"required,min=3,max=20"`
}
