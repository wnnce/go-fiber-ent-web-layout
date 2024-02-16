package usercase

type User struct {
	Username string   `json:"username,omitempty" validate:"required,lte=10,gte=5"`  // 用户名
	Password string   `json:"password,omitempty" validate:"required,lte=8,gte=5"`   // 密码
	Scopes   []string `json:"scopes,omitempty" validate:"omitempty,required,lte=1"` // 权限列表
}

type IUserService interface {
	Login(user *User) (string, error)
}
