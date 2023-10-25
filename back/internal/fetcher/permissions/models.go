package permissions

type PermissionGroup struct {
	Prefix string
	Users  []string
	Clock  ClockPermissions
	Order  OrderPermissions
}

type ClockPermissions struct {
	ShowClockMenu bool
	CreateClock   bool
	DeleteClock   bool
	UpdateClock   bool
}

type OrderPermissions struct {
	ShowOrderMenu bool
	DeleteOrder   bool
}
