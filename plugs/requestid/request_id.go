package requestid

import "github.com/AlexanderChen1989/plug"

const (
	// DefaultHTTPHeader default http header for request id
	DefaultHTTPHeader = "x-request-id"
)

type requestIDPlug struct {
	next plug.Plugger

	httpHeader string
}

func newPlug(header string) *requestIDPlug {
	return &requestIDPlug{httpHeader: header}
}

// New create a new request id Plugger
func New() plug.Plugger {
	return NewWithHeader(DefaultHTTPHeader)
}

// NewWithHeader create a new request id Plugger with customized http header
func NewWithHeader(header string) plug.Plugger {
	return newPlug(header)
}

func (p *requestIDPlug) Plug(next plug.Plugger) plug.Plugger {
	p.next = next
	return p
}

func validRequestID(rid string) bool {
	size := len(rid)
	return size >= 20 && size <= 200
}

func genRequestID() string {
	return randString(32)
}

func (p *requestIDPlug) getRequestID(conn plug.Conn) string {
	rid := conn.Request.Header.Get(p.httpHeader)

	if validRequestID(rid) {
		return rid
	}

	return genRequestID()
}

func (p *requestIDPlug) setRequestID(conn plug.Conn, rid string) plug.Conn {
	conn.ResponseWriter.Header().Add(p.httpHeader, rid)
	return conn
}

func (p *requestIDPlug) HandleConn(conn plug.Conn) {
	rid := p.getRequestID(conn)
	conn = p.setRequestID(conn, rid)
	p.next.HandleConn(conn)
}
