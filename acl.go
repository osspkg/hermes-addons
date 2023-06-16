package hermesaddons

type ACLGetter interface {
	Setup() []ACLModel
}

type ACLModel struct {
	ID      uint
	Title   string
	FormIDs []uint
}
