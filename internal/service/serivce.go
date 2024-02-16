package service

import "github.com/google/wire"

var InjectSet = wire.NewSet(NewExampleService, NewUserService)
