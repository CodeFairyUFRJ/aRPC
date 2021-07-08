// QUICListener e a implementacao do Listener para QUIC
type QUICListener struct {
	listener quic.Listener
}

// Close fecha o QUICListener
func (l *QUICListener) Close() error {
	return l.listener.Close()
}

// Accept recebe uma conexao para um dado QUICListener
func (l *QUICListener) Accept(ctx context.Context) (Session, error) {
	session, err := l.listener.Accept(ctx)
	if err != nil {
		return nil, err
	}
	
	return &QUICSession{session: session}, nil
}