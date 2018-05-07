package main

import (
	"fmt"
	"math"
)

// OSDRaw is the raw object from the log
type OSDRaw struct {
	Longitude         float64
	Latitude          float64
	Height            int16
	XSpeed            int16
	YSpeed            int16
	ZSpeed            int16
	Pitch             int16
	Rool              int16
	Yaw               int16
	RcState           int16
	AppCommand        int16
	Info              int32
	GpsNum            int16
	FlightAction      int16
	MotorFailedCause  int16
	NonGpsCause       int32
	WaypointLimitMode int32
	Battery           int32
	SwaveHeight       int16
	FlyTime           int32
	PlaceholderC      int16
	MotorRevolution   int16
	PlaceholderA      int16
	FlycVersion       int32
	DroneType         int32
	IMUinitFailReason int32
	MotorFailReason   int32
	PlaceholderB      int8
	SDKCtrlDevice     int32
}

// OSD is the humanified object
type OSD struct {
	Longitude            float64
	Latitude             float64
	Height               int16
	XSpeed               int16
	YSpeed               int16
	ZSpeed               int16
	Pitch                int16
	Rool                 int16
	Yaw                  int16
	RcState              bool
	FlycState            string
	AppCommand           int16
	IocWork              bool
	GroundOrSky          int32
	MotorUp              bool
	SwaveWork            bool
	GoHomeStatus         string
	ImuPreheated         bool
	VisionUsed           bool
	ModeChannel          int32
	VoltageWarning       int32
	CompassError         bool
	WaveError            bool
	GpsLevel             int32
	BatteryType          string
	AcceleratorOverRange bool
	Vibrating            bool
	BarometerDeadInAir   bool
	MotorBlocked         bool
	NotEnoughtForce      bool
	PropellerCatapult    bool
	GoHomeHeigthModified bool
	OutOfLimit           bool
	GpsNum               int16
	FlightAction         string
	MotorFailedCause     string
	NonGpsCause          string
	WaypointLimitMode    bool
	Battery              int32
	SwaveHeight          int16
	FlyTime              int32
	MotorRevolution      int16
	FlycVersion          int32
	DroneType            string
	IMUinitFailReason    string
	MotorFailReason      string
	SDKCtrlDevice        string
}

func (raw OSDRaw) from() OSD {

	osd := OSD{}

	osd.Longitude = raw.Longitude * 180 / math.Pi
	osd.Latitude = raw.Latitude * 180 / math.Pi
	osd.Height = raw.Height / 10
	osd.XSpeed = raw.XSpeed
	osd.YSpeed = raw.YSpeed
	osd.ZSpeed = raw.ZSpeed
	osd.Pitch = raw.Pitch
	osd.Rool = raw.Rool
	osd.Yaw = raw.Yaw
	osd.RcState = (raw.RcState & 128) == 0
	osd.FlycState = FlycState[raw.RcState&-129]
	osd.AppCommand = raw.AppCommand
	osd.IocWork = (raw.Info & 1) == 1
	osd.GroundOrSky = raw.Info >> 1 & 3
	osd.MotorUp = (raw.Info >> 3 & 1) == 1
	osd.SwaveWork = (raw.Info & 16) != 0
	osd.GoHomeStatus = GoHomeStatus[raw.Info>>5&7]
	osd.ImuPreheated = (raw.Info & 4096) != 0
	osd.VisionUsed = (raw.Info & 256) != 0
	//TODO
	osd.VoltageWarning = (raw.Info & 1536) >> 9
	osd.ModeChannel = (raw.Info & 24576) >> 13
	osd.CompassError = (raw.Info & 65536) != 0
	osd.WaveError = (raw.Info & 131075) != 0
	osd.GpsLevel = raw.Info >> 18 & 15
	osd.BatteryType = BatteryType[raw.Info>>22&3]
	osd.AcceleratorOverRange = (raw.Info >> 24 & 1) != 0
	osd.Vibrating = (raw.Info >> 25 & 1) != 0
	osd.BarometerDeadInAir = (raw.Info >> 26 & 1) != 0
	osd.MotorBlocked = (raw.Info >> 27 & 1) != 0
	osd.NotEnoughtForce = (raw.Info >> 28 & 1) != 0
	osd.PropellerCatapult = (raw.Info >> 29 & 1) != 0
	osd.GoHomeHeigthModified = (raw.Info >> 30 & 1) != 0
	osd.OutOfLimit = (raw.Info >> 31 & 1) != 0
	osd.GpsNum = raw.GpsNum
	osd.FlightAction = FlightAction[raw.FlightAction]
	osd.MotorFailedCause = MotorStartFailedCause[raw.MotorFailedCause&127]
	osd.NonGpsCause = NonGpsCause[raw.NonGpsCause&15]
	osd.WaypointLimitMode = (raw.WaypointLimitMode & 16) == 16
	osd.Battery = raw.Battery
	osd.SwaveHeight = raw.SwaveHeight
	osd.FlyTime = raw.FlyTime / 10
	osd.MotorRevolution = raw.MotorRevolution
	osd.FlycVersion = raw.FlycVersion
	osd.DroneType = DroneType[raw.DroneType]
	osd.IMUinitFailReason = ImuInitFailReason[raw.IMUinitFailReason]
	osd.MotorFailReason = MotorFailReason[raw.MotorFailReason]
	osd.SDKCtrlDevice = SdkControlDevice[raw.SDKCtrlDevice]

	fmt.Printf("OSD: %+v", osd)
	return osd
}
