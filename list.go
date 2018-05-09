package main

var DroneType = [...]string{
	0:   "Unknown",
	1:   "Inspire",
	2:   "P3S",
	3:   "P3X",
	4:   "P3C",
	5:   "OpenFrame",
	6:   "ACEONE",
	7:   "WKM",
	8:   "NAZA",
	9:   "A2",
	10:  "A3",
	11:  "P4",
	13:  "Mavic",
	14:  "PM820",
	15:  "P34K",
	16:  "wm220",
	17:  "Orange2",
	18:  "Pomato",
	20:  "N3",
	23:  "PM820PRO",
	100: "None",
	255: "NoFlyc",
}

var RecoverDroneType = [...]string{
	0:   "Unknown",
	1:   "Inspire",
	2:   "P3S",
	3:   "P3X",
	4:   "P3C",
	5:   "OpenFrame",
	7:   "P4",
	13:  "Mavic",
	100: "None"}

var AppType = [...]string{
	0: "UNKNOWN",
	1: "IOS",
	2: "ANDROID"}

var FlycState = [...]string{
	0:   "MANUAL",
	1:   "ATTI",
	2:   "ATTI_COURSE_LOCK",
	3:   "ATTI_HOVER",
	4:   "HOVER",
	5:   "GSP_BLAKE",
	6:   "GPS_ATTI",
	7:   "GPS_COURSE_LOCK",
	8:   "GPS_HOME_LOCK",
	9:   "GPS_HOT_POINT",
	10:  "ASSISTED_TAKEOFF",
	11:  "AUTO_TAKEOFF",
	12:  "AUTO_LANDING",
	13:  "ATTI_LANDING",
	14:  "GPS_WAYPOINT",
	15:  "GO_HOME",
	16:  "CLICK_GO",
	17:  "JOYSTICK",
	18:  "GPS_ATTI_WRISTBAND",
	19:  "CINEMATIC",
	23:  "ATTI_LIMITED",
	24:  "DRAW",
	25:  "GPS_FOLLOW_ME",
	26:  "ACTIVE_TRACK",
	27:  "TAP_FLY",
	28:  "PANO",
	29:  "FARMING",
	30:  "FPV",
	31:  "GPS_SPORT",
	32:  "GPS_NOVICE",
	33:  "CONFIRM_LANDING",
	35:  "TERRAIN_FOLLOW",
	36:  "PALM_CONTROL",
	37:  "QUICK_SHOT",
	38:  "TRIPOD",
	39:  "TRACK_SPOTLIGHT",
	41:  "MOTORS_JUST_STARTED",
	43:  "GPS_GENTLE",
	255: "UNKNOWN"}

var FlightAction = [...]string{
	0:  "NONE",
	1:  "WARNING_POWER_GOHOME",
	2:  "WARNING_POWER_LANDING",
	3:  "SMART_POWER_GOHOME",
	4:  "SMART_POWER_LANDING",
	5:  "LOW_VOLTAGE_LANDING",
	6:  "LOW_VOLTAGE_GOHOME",
	7:  "SERIOUS_LOW_VOLTAGE_LANDING",
	8:  "RC_ONEKEY_GOHOME",
	9:  "RC_ASSISTANT_TAKEOFF",
	10: "RC_AUTO_TAKEOFF",
	11: "RC_AUTO_LANDING",
	12: "APP_AUTO_GOHOME",
	13: "APP_AUTO_LANDING",
	14: "APP_AUTO_TAKEOFF",
	15: "OUTOF_CONTROL_GOHOME",
	16: "API_AUTO_TAKEOFF",
	17: "API_AUTO_LANDING",
	18: "API_AUTO_GOHOME",
	19: "AVOID_GROUND_LANDING",
	20: "AIRPORT_AVOID_LANDING",
	21: "TOO_CLOSE_GOHOME_LANDING",
	22: "TOO_FAR_GOHOME_LANDING",
	23: "APP_WP_MISSION",
	24: "WP_AUTO_TAKEOFF",
	25: "GOHOME_AVOID",
	26: "GOHOME_FINISH",
	27: "VERT_LOW_LIMIT_LANDING",
	28: "BATTERY_FORCE_LANDING",
	29: "MC_PROTECT_GOHOME"}

var BatteryType = [...]string{
	0: "UNKNOWN",
	1: "NONSMART",
	2: "SMART"}

var MotorStartFailedCause = [...]string{
	0:   "None",
	1:   "CompassError",
	2:   "AssistantProtected",
	3:   "DeviceLocked",
	4:   "DistanceLimit",
	5:   "IMUNeedCalibration",
	6:   "IMUSNError",
	7:   "IMUWarning",
	8:   "CompassCalibrating",
	9:   "AttiError",
	10:  "NoviceProtected",
	11:  "BatteryCellError",
	12:  "BatteryCommuniteError",
	13:  "SeriouLowVoltage",
	14:  "SeriouLowPower",
	15:  "LowVoltage",
	16:  "TempureVolLow",
	17:  "SmartLowToLand",
	18:  "BatteryNotReady",
	19:  "SimulatorMode",
	20:  "PackMode",
	21:  "AttitudeAbNormal",
	22:  "UnActive",
	23:  "FlyForbiddenError",
	24:  "BiasError",
	25:  "EscError",
	26:  "ImuInitError",
	27:  "SystemUpgrade",
	28:  "SimulatorStarted",
	29:  "ImuingError",
	30:  "AttiAngleOver",
	31:  "GyroscopeError",
	32:  "AcceletorError",
	33:  "CompassFailed",
	34:  "BarometerError",
	35:  "BarometerNegative",
	36:  "CompassBig",
	37:  "GyroscopeBiasBig",
	38:  "AcceletorBiasBig",
	39:  "CompassNoiseBig",
	40:  "BarometerNoiseBig",
	41:  "InvalidSn",
	44:  "FLASH_OPERATING",
	45:  "GPS_DISCONNECT",
	47:  "SDCardException",
	61:  "IMUNoconnection",
	62:  "RCCalibration",
	63:  "RCCalibrationException",
	64:  "RCCalibrationUnfinished",
	65:  "RCCalibrationException2",
	66:  "RCCalibrationException3",
	67:  "AircraftTypeMismatch",
	68:  "FoundUnfinishedModule",
	70:  "CYRO_ABNORMAL",
	71:  "BARO_ABNORMAL",
	72:  "COMPASS_ABNORMAL",
	73:  "GPS_ABNORMAL",
	74:  "NS_ABNORMAL",
	75:  "TOPOLOGY_ABNORMAL",
	76:  "RC_NEED_CALI",
	77:  "INVALID_FLOAT",
	78:  "M600_BAT_TOO_LITTLE",
	79:  "M600_BAT_AUTH_ERR",
	80:  "M600_BAT_COMM_ERR",
	81:  "M600_BAT_DIF_VOLT_LARGE_1",
	82:  "M600_BAT_DIF_VOLT_LARGE_2",
	83:  "INVALID_VERSION",
	84:  "GIMBAL_GYRO_ABNORMAL",
	85:  "GIMBAL_ESC_PITCH_NON_DATA",
	86:  "GIMBAL_ESC_ROLL_NON_DATA",
	87:  "GIMBAL_ESC_YAW_NON_DATA",
	88:  "GIMBAL_FIRM_IS_UPDATING",
	89:  "GIMBAL_DISORDER",
	90:  "GIMBAL_PITCH_SHOCK",
	91:  "GIMBAL_ROLL_SHOCK",
	92:  "GIMBAL_YAW_SHOCK",
	93:  "IMUcCalibrationFinished",
	101: "BatVersionError",
	102: "RTK_BAD_SIGNAL",
	103: "RTK_DEVIATION_ERROR",
	112: "ESC_CALIBRATING",
	113: "GPS_SIGN_INVALID",
	114: "GIMBAL_IS_CALIBRATING",
	115: "LOCK_BY_APP",
	116: "START_FLY_HEIGHT_ERROR",
	117: "ESC_VERSION_NOT_MATCH",
	118: "IMU_ORI_NOT_MATCH",
	119: "STOP_BY_APP",
	120: "COMPASS_IMU_ORI_NOT_MATCH",
	256: "OTHER"}

var NonGpsCause = [...]string{
	0: "ALREADY",
	1: "FORBIN",
	2: "GPSNUM_NONENOUGH",
	3: "GPS_HDOP_LARGE",
	4: "GPS_POSITION_NONMATCH",
	5: "SPEED_ERROR_LARGE",
	6: "YAW_ERROR_LARGE",
	7: "COMPASS_ERROR_LARGE",
	8: "UNKNOWN"}

var ImuInitFailReason = [...]string{
	0:   "MONITOR_ERROR",
	1:   "COLLECTING_DATA",
	2:   "GYRO_DEAD",
	3:   "ACCE_DEAD",
	4:   "COMPASS_DEAD",
	5:   "BAROMETER_DEAD",
	6:   "BAROMETER_NEGATIVE",
	7:   "COMPASS_MOD_TOO_LARGE",
	8:   "GYRO_BIAS_TOO_LARGE",
	9:   "ACCE_BIAS_TOO_LARGE",
	10:  "COMPASS_NOISE_TOO_LARGE",
	11:  "BAROMETER_NOISE_TOO_LARGE",
	12:  "WAITING_MC_STATIONARY",
	13:  "ACCE_MOVE_TOO_LARGE",
	14:  "MC_HEADER_MOVED",
	15:  "MC_VIBRATED",
	16:  "NONE",
	255: "TOTO"}

var MotorFailReason = [...]string{
	94:  "TAKEOFF_EXCEPTION",
	95:  "ESC_STALL_NEAR_GROUND",
	96:  "ESC_UNBALANCE_ON_GRD",
	97:  "ESC_PART_EMPTY_ON_GRD",
	98:  "ENGINE_START_FAILED",
	99:  "AUTO_TAKEOFF_LANCH_FAILED",
	100: "ROLL_OVER_ON_GRD",
	128: "OTHER"}

var SdkControlDevice = [...]string{
	0:   "RC",
	1:   "APP",
	2:   "ONBOARD_DEVICE",
	3:   "CAMERA",
	128: "OTHER"}

var GoHomeStatus = [...]string{
	0: "STANDBY",
	1: "PREASCENDING",
	2: "ALIGN",
	3: "ASCENDING",
	4: "CRUISE",
	7: "OTHER"}

var IocMode = [...]string{
	1:   "CourseLock",
	2:   "HomeLock",
	3:   "HotspotSurround",
	100: "OTHER"}

var MotorEscmState = [...]string{
	0:   "NON_SMART",
	1:   "DISCONNECT",
	2:   "SIGNAL_ERROR",
	3:   "RESISTANCE_ERROR",
	4:   "BLOCK",
	5:   "NON_BALANCE",
	6:   "ESCM_ERROR",
	7:   "PROPELLER_OFF",
	8:   "MOTOR_IDLE",
	9:   "MOTOR_UP",
	10:  "MOTOR_OFF",
	11:  "NON_CONNECT",
	100: "OTHER"}
