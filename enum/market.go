package enum

const (
	RdsKeyMarketInstrument    = "market:instrument:%s"     //market:instrument:BTC-240330-10000-Call
	RdsKeyPubMarketInstrument = "market:pub:instrument:%s" //market:pub:instrument:BTC-240330-10000-Call

	RdsKeyMarketUnderlying    = "market:underlying:%s"     //market:underlying:BTC
	RdsKeyPubMarketUnderlying = "market:pub:underlying:%s" //market:pub:underlying:BTC

	RdsKeyMarketUnderlyingIndex    = "market:underlying_index:%s"     //market:underlying_index:BTC-240330
	RdsKeyPubMarketUnderlyingIndex = "market:pub:underlying_index:%s" //market:pub:underlying_index:BTC-240330
)

const (
	ApiCacheUnderlyingRdsKey            = "api:cache:underlying_cache"          //api:cache:underlying
	ApiCacheUnderlyingIndexRdsKey       = "api:cache:underlying_index:%s"       //api:cache:underlying_index:{underlying}
	ApiCacheInstrumentsRdsKey           = "api:cache:instruments:%s"            //api:cache:instruments:{underlying_index}
	ApiCacheInstrumentRdsKey            = "api:cache:instrument:%s"             //api:cache:instrument:{instrumentID}
	ApiCacheUnderlyingInstrumentsRdsKey = "api:cache:underlying_instruments:%s" //api:cache:underlying_instruments:{underlying_index}
)

const (
	LiteOrderCountRdsKey = "market:lite_order_count:%s:%s" //market:lite_order_count:{underlying}:{date}
)
