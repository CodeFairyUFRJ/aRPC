package main

import (
    "github.com/almeida-raphael/arpc_examples/examples/greetings"
    "log"
)

func main(){
    rootCAPath := os.Getenv("CA_FILE")
    serverAddress := os.Getenv("SERVER_ADDRESS")

    rootCA, err := LoadCA(rootCAPath)
    if err != nil {
        log.Fatal(err)
    }

    tlsConfig := tls.Config{
        RootCAs:    rootCA,
        NextProtos: []string{"quic-arcp"},
    }

    aRPCController := controller.NewRPCController(
        channel.NewQUICChannel(
            serverAddress,
            7653,
            &tlsConfig,
            &quic.Config{
                MaxIncomingStreams: 100000,
            },
        ),
    )

    greetingsService := greetings.NewGreetings(&aRPCController)
    err := aRPCController.StartClient()
    if err != nil {
        log.Fatal(err)
    }
    
    hiResponse, err := greetingsService.SayHi(&requestData)

    // O tratamento de erros e resposta deve continuar abaixo
    
}