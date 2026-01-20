package enum

const (
	OrderStatusToFilled   int8 = 1
	OrderStatusFullFilled int8 = 2
	OrderStatusCanceled   int8 = -1
)

const (
	OrderRejectNo      int8 = 0
	OrderRejectYes     int8 = 1
	OrderRejectExpired int8 = 2
	OrderRejectSettled int8 = 3
)

const (
	OrderDirectionBuy  int8 = 1
	OrderDirectionSell int8 = 2
)

const (
	OrderTypeMarket int8 = 1
	OrderTypeLimit  int8 = 2
)

const (
	OrderTimeInForceMarket int8 = -1
	OrderTimeInForceGtc    int8 = 1
	OrderTimeInForceIoc    int8 = 2
	OrderTimeInForceFok    int8 = 3
	OrderTimeInForceGtd    int8 = 4
)

const (
	OrderMarketLob int8 = 0
	OrderMarketRfq int8 = 1
)

const (
	OrderIsSynNo  int8 = 0
	OrderIsSynYes int8 = 1
)
