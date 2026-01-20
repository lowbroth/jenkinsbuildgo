package enum

const (
	ApiLoginSignNonceRdsKey = "api:login:sign_nonce:%s" //api:login:sign_nonce:{address}
	ApiLockLoginRdsKey      = "api:lock:login:%s"       //api:lock:login:{address}
)

const (
	TokenDataCacheRdsKey = "TOKEN:DATA:%s"    // TOKEN:DATA:<access_token> json string
	TokenSecretRdsKey    = "TOKEN:SECRET:%s"  // TOKEN:SECRET:<access_token> secret
	TokenAddressRdsKey   = "TOKEN:ADDRESS:%s" // TOKEN:ADDRESS:<address> access token
	TokenUidRdsKey       = "TOKEN:UID:%d"     // TOKEN:DATA:<uid> access token
)

type RedisUserInfo struct {
	MemberId uint32
	Address  string
}

type RedisUserData struct {
	UserInfo *RedisUserInfo
	Ip       string
	IsWeb    int
	Uuid     string
	OsType   string
	Time     int64
}
