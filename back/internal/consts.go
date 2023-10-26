package consts

type ClockType int

const (
	One ClockType = 1
	Two ClockType = 2
)

func IntToClockType(num int) ClockType {
	if num == 1 {
		return One
	} else {
		return Two
	}
}

func ClockTypeToInt(num ClockType) int {
	if num == One {
		return 1
	} else {
		return 2
	}
}

const (
	ConfigName             = "config.yaml"
	DatabaseName           = "nixietech"
	CollectionOrderName    = "orders"
	CollectionClockName    = "clocks"
	CollectionSettingsName = "settings"
)
