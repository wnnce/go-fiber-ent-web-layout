package limiter

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestSha256KeyGenerate(t *testing.T) {
	keyGenerate := Sha256KeyGenerate()
	fmt.Printf("%p\n", keyGenerate)
	hash := sha256.New()
	for i := 0; i < 3; i++ {
		hash.Write([]byte("demo"))
		fmt.Println(hex.EncodeToString(hash.Sum(nil)))
		hash.Reset()
	}
}

func TestMd5KeyGenerate(t *testing.T) {
	keyGenerate := Md5KeyGenerate()
	fmt.Printf("%p\n", keyGenerate)
	hash := md5.New()
	for i := 0; i < 3; i++ {
		hash.Write([]byte("demo"))
		fmt.Println(hex.EncodeToString(hash.Sum(nil)))
		hash.Reset()
	}
}
