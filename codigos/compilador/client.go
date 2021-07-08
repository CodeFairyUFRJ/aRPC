package greetings

// Codigo gerado pelo aRPC; NAO EDITAR!

import (
	"context"
	"github.com/almeida-raphael/arpc/controller"
)

type Greetings struct {
	controller *controller.RPC
}

func NewGreetings(controller *controller.RPC) Greetings {
	return Greetings{
		controller: controller,
	}
}

func (greetings *Greetings)SayHi(SayHiRequest *SayHiRequest, ctx ...context.Context)(*SayHiResponse, error){
	if ctx == nil || len(ctx) == 0 {
		ctx = []context.Context{context.Background()}
	}
	
	response := SayHiResponse{}
	
	err := greetings.controller.SendRPCCall(
	    ctx[0], 
	    4148486943, 
	    0, 
	    SayHiRequest, 
	    &response,
	)
	
	if err != nil {
		return nil, err
	}
	
	return &response, nil
}