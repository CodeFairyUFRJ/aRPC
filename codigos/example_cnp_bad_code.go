//  Monta uma nova mensagem vazia, alocando uma struct do Cap'n Proto.
msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
if err != nil {
	panic(err)
}

// Cria uma nova struct de Book.
book, err := books.NewRootBook(seg)
if err != nil {
	panic(err)
}
book.SetTitle("Ready Player One")
book.SetPageCount(389)

// Escreve a mensagem para a saida padrao.
err = capnp.NewEncoder(os.Stdout).Encode(msg)
if err != nil {
	panic(err)
}