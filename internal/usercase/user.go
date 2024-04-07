package usercase

type User struct {
	UserId   uint64   `json:"userId"`
	Username string   `json:"username,omitempty" validate:"required,lte=10,gte=5"`  // 用户名
	Password string   `json:"password,omitempty" validate:"required,lte=8,gte=5"`   // 密码
	Scopes   []string `json:"scopes,omitempty" validate:"omitempty,required,lte=1"` // 权限列表
}

func (u *User) GetUserId() uint64 {
	return u.UserId
}
func (u *User) GetUserName() string {
	return u.Username
}
func (u *User) GetRoles() []string {
	return make([]string, 0)
}
func (u *User) GetPermissions() []string {
	return u.Scopes
}

type IUserService interface {
	Login(user *User) (string, error)
}
