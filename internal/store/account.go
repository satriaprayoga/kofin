package store

type Account struct {
	ID       int
	AccKode  string
	AccName  string
	Root     bool
	AccType  string
	AccGroup string
	ParentID int
}
