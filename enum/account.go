package enum

const (
	AccountBalanceStateNormal    = 1
	AccountBalanceStateLiquidate = 2
	AccountBalanceStateTakeover  = 3
	AccountBalanceStateADL       = 4
)

const (
	ApiLockDepositRdsKey = "api:lock:deposit:%d" //api:lock:deposit:{uid}
)
