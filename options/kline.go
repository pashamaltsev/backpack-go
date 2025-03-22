package options

type KLineInterval string

const (
	KLineInterval1m  KLineInterval = "1m"
	KLineInterval3m  KLineInterval = "3m"
	KLineInterval5m  KLineInterval = "5m"
	KLineInterval15m KLineInterval = "15m"
	KLineInterval30m KLineInterval = "30m"
	KLineInterval1h  KLineInterval = "1h"
	KLineInterval2h  KLineInterval = "2h"
	KLineInterval4h  KLineInterval = "4h"
	KLineInterval6h  KLineInterval = "6h"
	KLineInterval8h  KLineInterval = "8h"
	KLineInterval12h KLineInterval = "12h"
	KLineInterval1d  KLineInterval = "1d"
	KLineInterval3d  KLineInterval = "3d"
	KLineInterval1w  KLineInterval = "1w"
	KLineInterval1M  KLineInterval = "1M"
)
