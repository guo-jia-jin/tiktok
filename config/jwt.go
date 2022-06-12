package config

type JWT struct {
	SigningKey  string `json:"signingKey" yaml:"signingKey"`   // jwt签名
	ExpiresTime int64  `json:"expiresTime" yaml:"expiresTime"` // 过期时间
	BufferTime  int64  `json:"bufferTime" yaml:"bufferTime"`   // 缓冲时间
	Issuer      string `json:"issuer" yaml:"issuer"`           // 签发者
}
