package plug

type emptyPlugger struct{}

// NewEmpty create a new Empty Plugger
func NewEmpty() Plugger {
	return emptyPlugger{}
}

// Plug implements Plugger.Plug
func (empty emptyPlugger) Plug(Plugger) Plugger {
	return empty
}

// HandleConn implements Plugger.HandleConn
func (empty emptyPlugger) HandleConn(Conn) {
}
