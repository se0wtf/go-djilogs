package main

type Flight struct {
	Uuid          string     `json:"uuid"`
	UploadDate    int64      `json:"uploadDate"`
	FlightDate    int64      `json:"flightDate"`
	TotalDistance float32    `json:"totalDistance"` // meters
	TotalTime     int32      `json:"totalTime"`     // ms
	MaxHeight     float32    `json:"maxHeight"`     //meters
	MaxHSpeed     float32    `json:"maxHSpeed"`     // m/s
	MaxVSpeed     float32    `json:"maxVSpeed"`     // m/s
	Gps           []GpsPoint `json:"gps"`
}

type GpsPoint struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
