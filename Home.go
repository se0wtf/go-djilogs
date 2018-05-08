package main

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
	IsDynamicHomePointEnabled  bool
	AircraftHeadDirection      int16
	GoHomeMode                 int16
	IsHomeRecord               bool
	GoHomeHeight               int16
	CourseLockAngle            int16
	DataRecorderStatus         int
	DataRecorderRemainCapacity int16
	DataRecorderRemainTime     int16
	CurDataRecorderFileIndex   int16
	IsFlycInSimulationMode     bool
	IsFlycInNavigationMode     bool
	IsWingBroken               bool
	MotorEscmState             string
	ForceLandingHeight         int
}
