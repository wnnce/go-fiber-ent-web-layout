package limiter

import (
	"strings"
	"sync"
)

// IPMatch ip匹配接口
type IPMatch interface {
	Match(pattern, ip string) bool
}

func NewIPMatch() IPMatch {
	return &iPv4Match{}
}

// IPv4Match ipv4匹配
type iPv4Match struct {
	patternSplitCache sync.Map
}

func (m *iPv4Match) Match(pattern, ip string) bool {
	return m.doMatch(pattern, ip)
}

func (m *iPv4Match) doMatch(pattern, ip string) bool {
	pattSlice := make([]string, 0)
	if v, ok := m.patternSplitCache.Load(pattern); ok {
		pattSlice = v.([]string)
	}
	if len(pattSlice) == 0 {
		pattSlice = strings.Split(pattern, ".")
		m.patternSplitCache.Store(pattern, pattSlice)
	}
	ipSlice := strings.Split(ip, ".")
	if len(ipSlice) != len(pattSlice) {
		return false
	}
	for i, v := range pattSlice {
		if "*" == v {
			continue
		}
		if pattSlice[i] != ipSlice[i] {
			return false
		}
	}
	return true
}
