package limiter

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"github.com/gofiber/fiber/v3"
	"hash"
	"sync"
)

type KeyGenerate func(fiber.Ctx) string

type HashFunc func() hash.Hash

// Sha256KeyGenerate 通过sha256生成请求key
func Sha256KeyGenerate() KeyGenerate {
	return HashKeyGenerate(sha256.New)
}

// Md5KeyGenerate 通过md5生成请求key
func Md5KeyGenerate() KeyGenerate {
	return HashKeyGenerate(md5.New)
}

func HashKeyGenerate(hashFunc HashFunc) KeyGenerate {
	hashPool := &sync.Pool{
		New: func() interface{} {
			return hashFunc()
		},
	}
	return func(ctx fiber.Ctx) string {
		hasher := hashPool.Get().(hash.Hash)
		defer func() {
			hasher.Reset()
			hashPool.Put(hasher)
		}()
		originKey := ctx.IP() + "-" + ctx.Get(fiber.HeaderUserAgent, "")
		hasher.Write([]byte(originKey))
		return hex.EncodeToString(hasher.Sum(nil))
	}
}
