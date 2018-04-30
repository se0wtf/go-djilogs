package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"os"
)

var (
	test        []byte
	detailsAddr int32
	encrypted   int8
	offset      int32 = 12
)

type Details struct {
	SubStreet       [20]byte
	Street          [20]byte
	City            [20]byte
	Area            [20]byte
	Favorite        int8
	New             int8
	NeedsUpload     int8
	RecordLineCount int32
	_               int32
	UpdateTime      int64
	Longitude       float64
	Latitude        float64
	TotalDistance   float32
	TotalTime       int32
	MaxHeight       float32
	MaxHSpeed       float32
	MaxVSpeed       float32
	Photonum        int32
	VideoTime       int32
	_               [124]byte
	AircraftSn      [24]byte
	_               int8
	AircraftName    [10]byte
	_               [8]byte
	ActiveTimestamp float64
	CameraSn        [10]byte
	RcSn            [10]byte
	BatterySn       [10]byte
	AppType         int8
	AppVersionA     int8
	AppVersionB     int8
	AppVersionC     int8
}

type GPSPoint struct {
	Longitude float64
	Latitude  float64
	Accuracy  int32
}

type Message struct {
}

type CenterBattery struct {
	RelativeCapacity int8
	CurrentPV        int16
	CurrentCapacity  int16
	FullCapacity     int16
	Life             int8
	LoopNum          int16
	ErrorType        int64
	Current          int16
	Voltages         [12]byte
	Sn               int16
	ProductDate      int16
	Temperature      int16
	ConnStatus       int8
	TotalStudyCycle  int16
	LastStudyCycle   int16
}

type OSDRaw struct {
	Longitude         float64
	Latitude          float64
	Heigh             int16
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
	_                 int16
	FlycVersion       int8
	DroneType         int8
	IMUinitFailReason int8
	MotorFailReason   int8
	_                 int8
	SDKCtrlDevice     int8
}

func main() {
	path := "test.txt"

	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Error while opening file", err)
	}
	defer file.Close()
	fmt.Printf("%s opened\n", path)

	details := Details{}

	//details addr
	err = binary.Read(bytes.NewBuffer(readNextBytes(file, 0, 4)), binary.LittleEndian, &detailsAddr)
	if err != nil {
		log.Fatal("binary.Read failed", err)
	}
	fmt.Printf("The file is %d bytes long\n", fi.Size())

	//encrypted
	err = binary.Read(bytes.NewBuffer(readNextBytes(file, 10, 1)), binary.LittleEndian, &encrypted)
	if err != nil {
		log.Fatal("binary.Read failed", err)
	}

	//detail
	err = binary.Read(bytes.NewBuffer(readNextBytes(file, int64(detailsAddr), 452)), binary.LittleEndian, &details)
	if err != nil {
		log.Fatal("binary.Read failed", err)
	}

	fi, err := file.Stat()
	if err != nil {
		// Could not obtain stat, handle error
	}

	//1433907
	fmt.Printf("Longitude: %f\n", details.Longitude)
	fmt.Printf("Latitude: %f\n", details.Latitude)

	//	for offset < detailsAddr {
	//		isFrame(file, int64(offset))
	//		offset++
	//	}
}

func readNextBytes(file *os.File, offset int64, lenght int) []byte {
	bytes := make([]byte, lenght)
	_, err := file.ReadAt(bytes, offset)
	if err != nil {
		log.Fatal(err)
	}

	return bytes
}

func isFrame(file *os.File, offset int64) {
	tId := make([]byte, 1)
	_, err := file.ReadAt(tId, offset+1)
	if err != nil {
		log.Fatal(err)
	}

	length := make([]byte, 1)
	_, err2 := file.ReadAt(length, offset+2)
	if err2 != nil {
		log.Fatal(err2)
	}

	end := make([]byte, 1)
	_, err3 := file.ReadAt(end, offset+int64(binary.LittleEndian.Uint64(length)))
	if err3 != nil {
		log.Fatal(err3)
	}
	fmt.Printf("id: %d, length: %d, end: %x", tId, length, end)
}

func extractFrame(file *os.File, offset int64) {

}
