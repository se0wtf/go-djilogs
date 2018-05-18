package main

import (
	"math"
)

type Home struct {
	Latitude                   float64
	Longitude                  float64
	Height                     int32
	IocMode                    string
	IsIocEnabled               bool
	IsBeginnerMode             bool
	IsCompassCeleing           bool
	CompassCeleStatus          int16
	IsGoneHome                 bool
	GoneHomeStatus             string
	IsReachLimitHeight         bool
	IsReachLimitDistance       bool
	IsDynamicHomePointEnabled  bool
	AircraftHeadDirection      int16
	GoHomeMode                 int16
	IsHomeRecord               bool
	GoHomeHeight               int16
	CourseLockAngle            int16
	DataRecorderStatus         int
	DataRecorderRemainCapacity int
	DataRecorderRemainTime     int16
	CurDataRecorderFileIndex   int16
	IsFlycInSimulationMode     bool
	IsFlycInNavigationMode     bool
	IsWingBroken               bool
	MotorEscmState             string
	ForceLandingHeight         int
}

func createHome(decrypted []byte) Home {
	h := Home{}
	h.Longitude = Float64frombytes(decrypted[0:8]) * 180 / math.Pi
	h.Latitude = Float64frombytes(decrypted[8:16]) * 180 / math.Pi
	h.Height = Int32frombytes(decrypted[16:20])
	///TODO
	// if int((Int16frombytes(decrypted[20:22])&57344)>>13) >= len(IocMode) {
	// 	h.IocMode = "UNKNOWN"
	// } else {
	// 	h.IocMode = IocMode[(Int16frombytes(decrypted[20:22])&57344)>>13]
	// }
	// h.IsIocEnabled = ((Int16frombytes(decrypted[20:22]) & 57344) >> 13) != 0

	h.IsBeginnerMode = (Int16frombytes(decrypted[20:22]) >> 11 & 1) != 0
	h.IsCompassCeleing = ((Int16frombytes(decrypted[20:22]) & 1024) >> 10) != 0
	h.CompassCeleStatus = ((Int16frombytes(decrypted[20:22]) & 768) >> 8)
	h.IsGoneHome = ((Int16frombytes(decrypted[20:22]) & 128) >> 7) != 0

	if int((Int16frombytes(decrypted[20:22])&112)>>4) >= len(GoHomeStatus) {
		h.GoneHomeStatus = "UNKNOWN"
	} else {
		h.GoneHomeStatus = GoHomeStatus[(Int16frombytes(decrypted[20:22])&112)>>4]
	}

	h.IsReachLimitHeight = ((Int16frombytes(decrypted[20:22]) & 32) >> 5) != 0
	h.IsReachLimitDistance = ((Int16frombytes(decrypted[20:22]) & 16) >> 4) != 0
	h.IsDynamicHomePointEnabled = ((Int16frombytes(decrypted[20:22]) & 8) >> 3) != 0
	h.AircraftHeadDirection = ((Int16frombytes(decrypted[20:22]) & 4) >> 2)
	h.GoHomeMode = ((Int16frombytes(decrypted[20:22]) & 2) >> 1)
	h.IsHomeRecord = (Int16frombytes(decrypted[20:22]) & 1) != 0

	h.GoHomeHeight = Int16frombytes(decrypted[22:24])
	h.CourseLockAngle = Int16frombytes(decrypted[24:26])
	h.DataRecorderStatus = Intfrombytes(decrypted[26:27])
	h.DataRecorderRemainCapacity = Intfrombytes(decrypted[26:27])
	h.DataRecorderRemainTime = Int16frombytes(decrypted[28:30])
	h.CurDataRecorderFileIndex = Int16frombytes(decrypted[30:32])

	h.IsFlycInSimulationMode = (Intfrombytes(decrypted[32:33]) & 1) != 0
	h.IsFlycInNavigationMode = ((Intfrombytes(decrypted[32:33]) & 2) >> 1) != 0
	h.IsWingBroken = (Intfrombytes(decrypted[32:33]) & 4096) != 0
	//TODO
	//h.MotorEscmState = Int16frombytes(decrypted[30:32])
	h.ForceLandingHeight = Intfrombytes(decrypted[45:46])
	return h
}
