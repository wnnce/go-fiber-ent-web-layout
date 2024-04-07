package cache

import "github.com/google/wire"

var InjectSet = wire.NewSet(NewLoginUserCache)
