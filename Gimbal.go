package main

var GimbalMode = [...]string{
	"YawNoFollow",
	"FPV",
	"YawFollow",
	"OTHER"}

type Gimbal struct {
	Pitch                 int16
	Roll                  int16
	Yaw                   int16
	RollAdjust            int
	YawAngle              int16
	JoystickVerDirection  int
	JoystickHorDirection  int
	IsAutoCalibration     bool
	AutoCalibrationResult bool
	IsPitchInLimit        bool
	IsRollInLimit         bool
	IsYawInLimit          bool
	IsStuck               bool
	Mode                  string
	SubMode               int
	Version               int
	IsDoubleClick         bool
	IsTripleClick         bool
	IsSingleClick         bool
}

func createGimbal(decrypted []byte) Gimbal {
	//fmt.Printf("decrypted: %d\n", decrypted)
	g := Gimbal{}
	g.Pitch = Int16frombytes(decrypted[0:2])
	g.Roll = Int16frombytes(decrypted[2:4])
	g.Yaw = Int16frombytes(decrypted[4:6])
	g.RollAdjust = Intfrombytes(decrypted[7:8]) //7
	g.YawAngle = Int16frombytes(decrypted[8:10])
	g.JoystickVerDirection = Intfrombytes(decrypted[8:9]) & 3
	g.JoystickHorDirection = (Intfrombytes(decrypted[8:9]) >> 2) & 3
	g.IsAutoCalibration = Intfrombytes(decrypted[10:11]) != 0
	g.AutoCalibrationResult = (Intfrombytes(decrypted[10:11]) & 16) != 0
	g.IsPitchInLimit = (Intfrombytes(decrypted[10:11]) & 1) != 0
	g.IsRollInLimit = (Intfrombytes(decrypted[10:11]) & 2) != 0
	g.IsYawInLimit = (Intfrombytes(decrypted[10:11]) & 4) != 0
	g.IsStuck = (Intfrombytes(decrypted[10:11]) & 64) != 0
	g.Mode = GimbalMode[Intfrombytes(decrypted[6:7])>>6]
	g.SubMode = (Intfrombytes(decrypted[6:7]) >> 5) & 1
	g.Version = Intfrombytes(decrypted[11:12]) & 15
	g.IsDoubleClick = (Intfrombytes(decrypted[11:12]) & 32) != 0
	g.IsTripleClick = (Intfrombytes(decrypted[11:12]) & 64) != 0
	g.IsSingleClick = (Intfrombytes(decrypted[11:12]) & 128) != 0

	return g
}
