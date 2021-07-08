// QUICStream is a channel stream implementation for QUIC
type QUICStream struct {
	stream quic.Stream
}

// Close fecha a QUICStream
func (s *QUICStream) Close() error {
	return s.stream.Close()
}

// Write escreve os dados em uma QUICStream
func (s *QUICStream) Write(buf []byte) (int, error) {
	return s.stream.Write(buf)
}

// Read le os dados de uma QUICStream
func (s *QUICStream) Read(buf []byte) (int, error) {
	return s.stream.Read(buf)
}