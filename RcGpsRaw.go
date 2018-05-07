package main

type RcGpsRaw struct {
	PlaceholderA [7]byte
	Latitude     int64
	Longitude    int64
	XSpeed       int64
	YSpeed       int64
	GpsNum       int8
	GpsStatus    int16
}

type RcGps struct {
	Latitude  int64
	Longitude int64
	XSpeed    int64
	YSpeed    int64
	GpsNum    int8
	GpsStatus bool
}

func (raw RcGpsRaw) from() RcGps {
	msg := RcGps{}
	msg.Latitude = raw.Latitude
	msg.Longitude = raw.Longitude
	msg.XSpeed = raw.XSpeed / 1000
	msg.YSpeed = raw.YSpeed / 1000
	msg.GpsNum = raw.GpsNum
	msg.GpsStatus = raw.GpsStatus == 1

	return msg
}
