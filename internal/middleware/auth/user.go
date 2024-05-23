package auth

import (
	"go-fiber-ent-web-layout/internal/tools/pool"
	"sync"
	"time"
)

// LoginUser 登录用户接口
type LoginUser interface {
	GetUserId() uint64
	GetUserName() string
	GetRoles() []string
	GetPermissions() []string
}

type LoginUserCache interface {
	// AddUser 添加管理端登录用户
	AddUser(userId uint64, user LoginUser, expire time.Duration)
	// ResetExpire 重置Token的过期时间
	ResetExpire(userId uint64, expire time.Duration)
	// RemoveUser 删除管理端登录用户
	RemoveUser(userId uint64)
	// GetUser 获取管理端登录用户
	GetUser(userId uint64) LoginUser
}

func NewLoginUserCache() LoginUserCache {
	return &inMemoryLoginUserCache{
		cache: make(map[uint64]*cacheNode),
		nodePool: &sync.Pool{
			New: func() any {
				return new(cacheNode)
			},
		},
	}
}

type cacheNode struct {
	expireTime int64
	value      LoginUser
}

func (cn *cacheNode) Reset() {
	cn.expireTime = 0
	cn.value = nil
}

type inMemoryLoginUserCache struct {
	cache    map[uint64]*cacheNode
	mutex    sync.RWMutex
	nodePool *sync.Pool
}

func (mc *inMemoryLoginUserCache) AddUser(userId uint64, user LoginUser, expire time.Duration) {
	if user == nil {
		return
	}
	mc.mutex.Lock()
	node := mc.nodePool.Get().(*cacheNode)
	node.expireTime = time.Now().UnixMilli() + expire.Milliseconds()
	node.value = user
	mc.cache[userId] = node
	mc.mutex.Unlock()
}

func (mc *inMemoryLoginUserCache) ResetExpire(userId uint64, expire time.Duration) {
	mc.mutex.Lock()
	if node, ok := mc.cache[userId]; ok {
		node.expireTime = time.Now().UnixMilli() + expire.Milliseconds()
	}
	mc.mutex.Unlock()
}

func (mc *inMemoryLoginUserCache) RemoveUser(userId uint64) {
	mc.mutex.Lock()
	if node, ok := mc.cache[userId]; ok {
		node.Reset()
		mc.nodePool.Put(node)
		delete(mc.cache, userId)
	}
	mc.mutex.Unlock()
}

func (mc *inMemoryLoginUserCache) GetUser(userId uint64) LoginUser {
	mc.mutex.RLock()
	defer mc.mutex.RUnlock()
	node, ok := mc.cache[userId]
	if !ok {
		return nil
	}
	if node.expireTime <= time.Now().UnixMilli() {
		// 异步删除
		pool.Go(func() {
			mc.RemoveUser(userId)
		})
		return nil
	}
	return node.value
}

var (
	defaultLoginUserCache LoginUserCache
	// LoginUserCacheExpireTime 管理端登录用户的过期时间
	LoginUserCacheExpireTime = 30 * time.Minute
)

func init() {
	defaultLoginUserCache = NewLoginUserCache()
}

// AddLoginUser 添加管理端登录用户
// token 请求中的token参数
// user 管理端登录用户
// expire 过期时间
func AddLoginUser(userId uint64, user LoginUser, expire time.Duration) {
	defaultLoginUserCache.AddUser(userId, user, expire)
}

// ResetLoginUserExpire 重置管理端登录用户的过期时间
// token 请求携带的token
// expire 新的过期时间
func ResetLoginUserExpire(userId uint64, expire time.Duration) {
	defaultLoginUserCache.ResetExpire(userId, expire)
}

// RemoveLoginUser 删除管理端登录用户
func RemoveLoginUser(userId uint64) {
	defaultLoginUserCache.RemoveUser(userId)
}

// GetLoginUser 获取管理端登录用户
func GetLoginUser(userId uint64) LoginUser {
	return defaultLoginUserCache.GetUser(userId)
}
