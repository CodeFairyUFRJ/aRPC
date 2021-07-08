package greetings

// Codigo gerado pelo aRPC; NAO EDITAR!

import "github.com/almeida-raphael/arpc/controller"

type GreetingsServer interface {
	SayHi(SayHiRequest *SayHiRequest)(*SayHiResponse, error)
}

func bindSayHi(server GreetingsServer)(
	func(msg []byte)([]byte, error),
) {
	return func(msg []byte)([]byte, error){
		SayHiRequest := SayHiRequest{}
		err := SayHiRequest.UnmarshalBinary(msg)
		if err != nil {
			return nil, err
		}

		response, err := server.SayHi(&SayHiRequest)
		if err != nil {
			return nil, err
		}

		responseBytes, err := response.MarshalBinary()
		if err != nil {
			return nil, err
		}

		return responseBytes, nil
	}
}

func RegisterGreetingsServer(controller controller.RPC, srv GreetingsServer){
	controller.RegisterService(
		4148486943,
		map[uint16]func(message []byte)([]byte, error){
			0: bindSayHi(srv),
		},
	)
}