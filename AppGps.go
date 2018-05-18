package main

// AppGps AppGps
type AppGps struct {
	Latitude  float64
	Longitude float64
	Accuracy  float32
}

func createAppGps(decrypted []byte) AppGps {
	a := AppGps{}
	a.Longitude = Float64frombytes(decrypted[0:8])
	a.Latitude = Float64frombytes(decrypted[8:16])
	a.Accuracy = Float32frombytes(decrypted[16:20])
	return a
}
