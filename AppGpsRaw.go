package main

type AppGpsRaw struct {
	payload []byte
}

type AppGps struct {
	Latitude  float64
	Longitude float64
	Accuracy  float32
}
