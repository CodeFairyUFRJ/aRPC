// QUICRPC implementacao do canal de RPC para QUIC
type QUICRPC struct {
	address   string
	port      int
	tlsConfig *tls.Config
	config    *quic.Config
}

// Listen escuta as conexoes de entrada no QUICRPC
func (q *QUICRPC) Listen() (Listener, error) {
	listener, err := quic.ListenAddr(
	    fmt.Sprintf("%s:%d", q.address, q.port), 
	    q.tlsConfig, 
	    q.config,
	)
	if err != nil {
		return nil, err
	}
	return &QUICListener{listener: listener}, nil
}

// Connect conecta num servidor com QUICRPC
func (q *QUICRPC) Connect() (Session, error) {
	session, err := quic.DialAddr(
	    fmt.Sprintf("%s:%d", q.address, q.port), 
	    q.tlsConfig, 
	    q.config,
    )
	if err != nil {
		return nil, err
	}
	return &QUICSession{session: session}, nil
}
