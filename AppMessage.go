package main

type AppMessage struct {
	Message string
}

func createAppMessage(decrypted []byte) AppMessage {
	a := AppMessage{}
	a.Message = string(decrypted)
	return a
}
