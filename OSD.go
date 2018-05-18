package main

import (
	"math"
)

// OSD is the humanified object
type OSD struct {
	Longitude              float64
	Latitude               float64
	Height                 int16
	XSpeed                 int16
	YSpeed                 int16
	ZSpeed                 int16
	Pitch                  int16
	Rool                   int16
	Yaw                    int16
	RcState                bool
	FlycState              string
	AppCommand             int
	IocWork                bool
	GroundOrSky            int32
	IsMotorUp              bool
	IsSwaveWork            bool
	GoHomeStatus           string
	IsImuPreheated         bool
	IsVisionUsed           bool
	VoltageWarning         int32
	ModeChannel            int32
	IsCompassError         bool
	IsWaveError            bool
	GpsLevel               int32
	BatteryType            string
	IsAcceleratorOverRange bool
	IsVibrating            bool
	IsBarometerDeadInAir   bool
	IsMotorBlocked         bool
	IsNotEnoughtForce      bool
	IsPropellerCatapult    bool
	IsGoHomeHeigthModified bool
	IsOutOfLimit           bool
	GpsNum                 int
	FlightAction           string
	MotorFailedCause       string
	NonGpsCause            string
	IsWaypointLimitMode    bool
	Battery                int
	SwaveHeight            int
	FlyTime                int16
	MotorRevolution        int
	FlycVersion            int
	DroneType              string
	IMUinitFailReason      string
	MotorFailReason        string
	SDKCtrlDevice          string
}

func createOSD(decrypted []byte) OSD {

	osd := OSD{}

	osd.Longitude = Float64frombytes(decrypted[0:8]) * 180 / math.Pi
	osd.Latitude = Float64frombytes(decrypted[8:16]) * 180 / math.Pi
	osd.Height = Int16frombytes(decrypted[16:18])
	osd.XSpeed = Int16frombytes(decrypted[18:20])
	osd.YSpeed = Int16frombytes(decrypted[20:22])
	osd.ZSpeed = Int16frombytes(decrypted[22:24])
	osd.Pitch = Int16frombytes(decrypted[24:26])
	osd.Rool = Int16frombytes(decrypted[26:28])
	osd.Yaw = Int16frombytes(decrypted[28:30])
	osd.RcState = (Intfrombytes(decrypted[30:31]) & 128) == 0

	if Intfrombytes(decrypted[30:31])&-129 > len(FlycState) {
		osd.FlycState = "UNKNOWN"
	} else {
		osd.FlycState = FlycState[Intfrombytes(decrypted[30:31])&-129]
	}

	osd.AppCommand = Intfrombytes(decrypted[31:32])
	osd.IocWork = (Int32frombytes(decrypted[32:36]) & 1) == 1
	osd.GroundOrSky = Int32frombytes(decrypted[32:36]) >> 1 & 3
	osd.IsMotorUp = (Int32frombytes(decrypted[32:36]) >> 3 & 1) == 1
	osd.IsSwaveWork = (Int32frombytes(decrypted[32:36]) & 16) != 0

	if int(Int32frombytes(decrypted[32:36])>>5&7) >= len(GoHomeStatus) {
		osd.GoHomeStatus = "UNKNOWN"
	} else {
		osd.GoHomeStatus = GoHomeStatus[Int32frombytes(decrypted[32:36])>>5&7]
	}

	osd.IsImuPreheated = (Int32frombytes(decrypted[32:36]) & 4096) != 0
	osd.IsVisionUsed = (Int32frombytes(decrypted[32:36]) & 256) != 0
	osd.VoltageWarning = (Int32frombytes(decrypted[32:36]) & 1536) >> 9
	osd.ModeChannel = (Int32frombytes(decrypted[32:36]) & 24576) >> 13
	osd.IsCompassError = (Int32frombytes(decrypted[32:36]) & 65536) != 0
	osd.IsWaveError = (Int32frombytes(decrypted[32:36]) & 131072) != 0
	osd.GpsLevel = Int32frombytes(decrypted[32:36]) >> 18 & 15

	if int(Int32frombytes(decrypted[32:36])>>22&3) >= len(BatteryType) {
		osd.BatteryType = "UNKNOWN"
	} else {
		osd.BatteryType = BatteryType[Int32frombytes(decrypted[32:36])>>22&3]
	}

	osd.IsAcceleratorOverRange = (Int32frombytes(decrypted[32:36]) >> 24 & 1) != 0
	osd.IsVibrating = (Int32frombytes(decrypted[32:36]) >> 25 & 1) != 0
	osd.IsBarometerDeadInAir = (Int32frombytes(decrypted[32:36]) >> 26 & 1) != 0
	osd.IsMotorBlocked = (Int32frombytes(decrypted[32:36]) >> 27 & 1) != 0
	osd.IsNotEnoughtForce = (Int32frombytes(decrypted[32:36]) >> 28 & 1) != 0
	osd.IsPropellerCatapult = (Int32frombytes(decrypted[32:36]) >> 29 & 1) != 0
	osd.IsGoHomeHeigthModified = (Int32frombytes(decrypted[32:36]) >> 30 & 1) != 0
	osd.IsOutOfLimit = (Int32frombytes(decrypted[32:36]) >> 31 & 1) != 0
	osd.GpsNum = Intfrombytes(decrypted[36:37])

	if Intfrombytes(decrypted[37:38]) >= len(FlightAction) {
		osd.FlightAction = "UNKNOWN"
	} else {
		osd.FlightAction = FlightAction[Intfrombytes(decrypted[37:38])]
	}

	if Intfrombytes(decrypted[38:39])&127 >= len(MotorStartFailedCause) {
		osd.MotorFailedCause = "UNKNOWN"
	} else {
		osd.MotorFailedCause = MotorStartFailedCause[Intfrombytes(decrypted[38:39])&127]
	}

	if Intfrombytes(decrypted[39:40])&15 >= len(NonGpsCause) {
		osd.NonGpsCause = "UNKNOWN"
	} else {
		osd.NonGpsCause = NonGpsCause[Intfrombytes(decrypted[39:40])&15]
	}

	osd.IsWaypointLimitMode = (Intfrombytes(decrypted[39:40]) & 16) == 16
	osd.Battery = Intfrombytes(decrypted[40:41])
	osd.SwaveHeight = Intfrombytes(decrypted[41:42])
	osd.FlyTime = Int16frombytes(decrypted[42:44]) / 10
	osd.MotorRevolution = Intfrombytes(decrypted[44:45])
	osd.FlycVersion = Intfrombytes(decrypted[40:41])

	if Intfrombytes(decrypted[48:49]) >= len(DroneType) {
		osd.DroneType = "UNKNOWN"
	} else {
		osd.DroneType = DroneType[Intfrombytes(decrypted[48:49])]
	}

	if Intfrombytes(decrypted[49:50]) >= len(ImuInitFailReason) {
		osd.IMUinitFailReason = "UNKNOWN"
	} else {
		osd.IMUinitFailReason = ImuInitFailReason[Intfrombytes(decrypted[49:50])]
	}

	if Intfrombytes(decrypted[50:51]) >= len(MotorFailReason) {
		osd.MotorFailReason = "UNKNOWN"
	} else {
		osd.MotorFailReason = MotorFailReason[Intfrombytes(decrypted[50:51])]
	}

	if Intfrombytes(decrypted[52:53]) >= len(SdkControlDevice) {
		osd.SDKCtrlDevice = "UNKNOWN"
	} else {
		osd.SDKCtrlDevice = SdkControlDevice[Intfrombytes(decrypted[52:53])]
	}

	return osd
}
