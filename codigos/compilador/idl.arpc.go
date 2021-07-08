package greetings

// Person e uma struct com os dados de uma pessoa
type Person struct {
	title uint64
	name  text
}

// SayHiRequest e a struct passada como entrada do procedimento SayHi
type SayHiRequest struct {
	person Person
}

// SayHiResponse e a struct passada como saida do procedimento SayHi
type SayHiResponse struct {
	response text
}

// SayHi declaracao do procedimento SayHi
type SayHi func(request *SayHiRequest) (*SayHiResponse, error)

