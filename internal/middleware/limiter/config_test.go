package limiter

import (
	"fmt"
	"testing"
)

func TestConfig_AddExcludePath(t *testing.T) {
	config := DefaultConfig
	fmt.Printf("%v\n", config)
	config.AddExcludePath("/demo", "test")
	fmt.Printf("%v\n", config)
	config.AddExcludePath("/demo2").AddExcludePath("/test2").AddExcludeIP("127.0.0.1")
	fmt.Printf("%v\n", config)
}

func TestConfig_AddExcludeIP(t *testing.T) {
	config := DefaultConfig
	fmt.Printf("%v\n", config)
	config.AddExcludeIP("127.0.0.1")
	fmt.Printf("%v\n", config)
	config.AddExcludePath("/demo2").AddExcludePath("/test2").AddExcludeIP("192.168.1.100")
	fmt.Printf("%v\n", config)
}
