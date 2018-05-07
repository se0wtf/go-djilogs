package main

var ConnStatus = [...]string{
	"NORMAL",
	"INVALID",
	"EXCEPTION",
	"OTHER"}

type CenterBatteryRaw struct {
	RelativeCapacity int8
	CurrentPV        int16
	CurrenctCapacity int16
	FullCapacity     int16
	Life             int8
	LoopNum          int16
	ErrorType        int32
	Current          int16
	Voltages         [12]byte
	SerialNo         int16
	ProductDate      int16
	Temperature      int16
	ConnStatus       int8
	TotalStudyCycle  int16
}
