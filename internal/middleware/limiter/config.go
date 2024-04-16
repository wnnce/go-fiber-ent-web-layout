package limiter

import "time"

// Config 限流中间件配置
type Config struct {
	KeyGenerate     KeyGenerate       // 键生成策略
	CallbackHandler CallbackHandler   // 请求被限流后的处理策略
	ExcludePaths    []string          // 排除的路径
	ExcludeIPs      []string          // 排除的请求IP
	Sliding         SlidingConfig     // 滑动时间窗口限流器配置
	TokenBucket     TokenBucketConfig // 令牌桶限流器配置
	CustomLimiters  []Limiter         // 用户自定义的限流器
}

// AddExcludePath 添加需要排除的请求路径 链式调用
func (c *Config) AddExcludePath(paths ...string) *Config {
	c.ExcludePaths = append(c.ExcludePaths, paths...)
	return c
}

// AddExcludeIP 添加需要排除的请求IP 链式调用
func (c *Config) AddExcludeIP(ips ...string) *Config {
	c.ExcludeIPs = append(c.ExcludeIPs, ips...)
	return c
}

// RegisterLimiter 注册用户自定义的限流器 链式调用
func (c *Config) RegisterLimiter(limiters ...Limiter) *Config {
	c.CustomLimiters = append(c.CustomLimiters, limiters...)
	return c
}

// SlidingConfig 滑动时间窗口限流器配置
type SlidingConfig struct {
	Enable    bool          `json:"enable" yaml:"enable"`
	Name      string        `json:"name" yaml:"name"`            // 自定义限流器的名称
	WindowNum int           `json:"windowNum" yaml:"window-num"` // 时间窗口的数量 滑动窗口的统计时间 = 窗口数量 * 更新间隔
	Interval  time.Duration `json:"interval" yaml:"interval"`
	Threshold int64         `json:"threshold" yaml:"threshold"` // 客户端在时间端内的最大请求数量
}

// TokenBucketConfig 令牌桶限流器配置
type TokenBucketConfig struct {
	Enable          bool          `json:"enable" yaml:"enable"`
	Name            string        `json:"name" yaml:"name"`                        // 自定义的限流器名称
	MaxNum          int64         `json:"maxNum" yaml:"max-num"`                   // 令牌桶中令牌的最大数量
	DefaultAvail    int64         `json:"defaultAvail" yaml:"default-avail"`       // 令牌还未释放时的默认初始化数量
	ReleaseInterval time.Duration `json:"releaseInterval" yaml:"release-interval"` // 每次令牌释放的间隔时间
	ReleaseNum      int           `json:"releaseNum" yaml:"release-num"`           // 每次释放的令牌数量
}

var DefaultConfig = Config{
	KeyGenerate:     Md5KeyGenerate(),
	CallbackHandler: DefaultCallbackHandler,
	Sliding:         DefaultSlidingConfig,
	TokenBucket:     DefaultTokenBucketConfig,
}

func configDefault(cfg Config) Config {
	if cfg.KeyGenerate == nil {
		cfg.KeyGenerate = DefaultConfig.KeyGenerate
	}
	if cfg.CallbackHandler == nil {
		cfg.CallbackHandler = DefaultConfig.CallbackHandler
	}
	return cfg
}

var DefaultSlidingConfig = SlidingConfig{
	Enable:    true,
	Name:      "sliding-window-limiter",
	WindowNum: 10,
	Interval:  6 * time.Second,
	Threshold: 300,
}

func slidingConfigDefault(cfg SlidingConfig) SlidingConfig {
	if cfg.Name == "" {
		cfg.Name = DefaultSlidingConfig.Name
	}
	if cfg.WindowNum <= 1 {
		cfg.WindowNum = DefaultSlidingConfig.WindowNum
	}
	if cfg.Interval == 0 {
		cfg.Interval = DefaultSlidingConfig.Interval
	}
	if cfg.Threshold <= 0 {
		cfg.Threshold = DefaultSlidingConfig.Threshold
	}
	return cfg
}

var DefaultTokenBucketConfig = TokenBucketConfig{
	Enable:          true,
	Name:            "token-bucket-limiter",
	MaxNum:          1000,
	DefaultAvail:    50,
	ReleaseInterval: 5 * time.Second,
	ReleaseNum:      100,
}

func tokenBucketConfigDefault(cfg TokenBucketConfig) TokenBucketConfig {
	if cfg.Name == "" {
		cfg.Name = DefaultTokenBucketConfig.Name
	}
	if cfg.MaxNum <= 0 {
		cfg.MaxNum = DefaultTokenBucketConfig.MaxNum
	}
	if cfg.DefaultAvail <= 0 {
		cfg.DefaultAvail = DefaultTokenBucketConfig.DefaultAvail
	}
	if cfg.ReleaseInterval == 0 {
		cfg.ReleaseInterval = DefaultTokenBucketConfig.ReleaseInterval
	}
	if cfg.ReleaseNum <= 0 {
		cfg.ReleaseNum = DefaultTokenBucketConfig.ReleaseNum
	}
	return cfg
}
