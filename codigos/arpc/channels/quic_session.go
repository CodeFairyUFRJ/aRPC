// QUICSession e a implementacao de Session para o QUIC
type QUICSession struct {
	session quic.Session
}

// RemoteAddress retorna o endereco e a porta de uma QUICSession
func (qs *QUICSession) RemoteAddress() (string, int, error) {
	addrParts := strings.Split(qs.session.RemoteAddr().String(), ":")
	if len(addrParts) != 2 {
		return "", 0, fmt.Errorf("invalid remote address")
	}
	
	portInt64, err := strconv.ParseInt(addrParts[1], 10, 32)
	if err != nil {
		return "", 0, fmt.Errorf("invalid port number")
	}
	
	return addrParts[0], int(portInt64), nil
}

// AcceptStream recebe uma stream de dados de um client
func (qs *QUICSession) AcceptStream(ctx context.Context) (Stream, error) {
	stream, err := qs.session.AcceptStream(ctx)
	if err != nil {
		return nil, err
	}
	
	return &QUICStream{stream: stream}, nil
}

// OpenStream cria uma nova stream de dados com o servidor
func (qs *QUICSession) OpenStream(ctx context.Context) (Stream, error) {
	stream, err := qs.session.OpenStreamSync(ctx)
	if err != nil {
		return nil, err
	}
	
	return &QUICStream{stream: stream}, nil
}