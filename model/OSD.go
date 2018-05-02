package model

import "math"

type OSD struct {
	Longitude         float64
	Latitude          float64
	Height            int16
	XSpeed            int16
	YSpeed            int16
	ZSpeed            int16
	Pitch             int16
	Rool              int16
	Yaw               int16
	RcState           int8
	AppCommand        int8
	Info              int32
	GpsNum            int8
	FlightAction      int8
	MotorFailedCause  int8
	NonGpsCause       int8
	WaypointLimitMode int8
	Battery           int8
	SwaveHeight       int8
	FlyTime           int16
	MotorRevolution   int8
	PlaceholderA      int16
	FlycVersion       int8
	DroneType         int8
	IMUinitFailReason int8
	MotorFailReason   int8
	PlaceholderB      int8
	SDKCtrlDevice     int8
}

func (osd OSD) modifyFields() {
	osd.Longitude = osd.Longitude * 180 / math.Pi
	osd.Latitude = osd.Latitude * 180 / math.Pi
	osd.Height = osd.Height / 10
}
