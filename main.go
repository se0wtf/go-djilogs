package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"os"

	"github.com/satori/go.uuid"
)

var (
	test        []byte
	detailsAddr int32
	encrypted   int8
	offset      int32 = 12
	types       []string
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

type Frame struct {
	offset    int64
	typeID    uint8
	length    uint8
	bytekey   uint8
	key       [8]byte
	payload   []byte
	encrypted bool
}

func main() {
	//path := "testair.txt"
	path := "testv3.txt"

	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Error while opening file", err)
	}
	defer file.Close()
	fmt.Printf("%s opened\n", path)

	details := Details{}
	fl := Flight{}
	fl.Uuid = uuid.Must(uuid.NewV4()).String()

	// init types
	types := make([]string, 256)
	types[1] = "OSD"
	types[2] = "HOME"
	types[3] = "GIMBAL"
	types[4] = "RC"
	types[5] = "CUSTOM"
	types[6] = "DEFORM"
	types[7] = "CENTER_BATTERY"
	types[8] = "SMART_BATTERY"
	types[9] = "APP_TIP"
	types[10] = "APP_WARN"
	types[11] = "RC_GPS"
	types[12] = "RC_DEBUG"
	types[13] = "RECOVER"
	types[14] = "APP_GPS"
	types[15] = "FIRMWARE"
	types[16] = "OFDM_DEBUG"
	types[17] = "VISION_GROUP"
	types[18] = "VISION_WARN"
	types[19] = "MC_PARAM"
	types[20] = "APP_OPERATION"
	types[254] = "OTHER"
	types[255] = "END"

	//details addr
	err = binary.Read(bytes.NewBuffer(readNextBytes(file, 0, 4)), binary.LittleEndian, &detailsAddr)
	if err != nil {
		log.Fatal("binary.Read failed", err)
	}

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

	fl.FlightDate = details.UpdateTime
	fl.MaxHeight = details.MaxHeight
	fl.MaxHSpeed = details.MaxHSpeed
	fl.MaxVSpeed = details.MaxVSpeed
	fl.TotalDistance = details.TotalDistance
	fl.TotalTime = details.TotalTime

	for offset < detailsAddr {
		f, created := isFrame(file, int64(offset), types)
		if created {
			offset = int32(f.offset + int64(f.length) + 1 + 2)
			decryptFrame(f, &fl)
		} else if isImage(file, int64(offset)) {
			// we check if its an image
			fmt.Printf("WE GOT AN IMAGE §§§\n")
			offset++
		} else {
			offset++
		}
	}

	fmt.Printf("Flight: %+v\n", fl)
}

func readNextBytes(file *os.File, offset int64, length int) []byte {
	bytes := make([]byte, length)
	_, err := file.ReadAt(bytes, offset)
	if err != nil {
		log.Fatal(err)
	}

	return bytes
}

func isImage(file *os.File, offset int64) bool {
	var header uint32
	err := binary.Read(bytes.NewBuffer(readNextBytes(file, offset, 4)), binary.LittleEndian, &header)
	if err != nil {
		log.Fatal("binary.Read failed", err)
	}
	// JFIF header 0xFF 0xD8 0xFF 0xE0
	return header == 3774863615
}

func isFrame(file *os.File, offset int64, types []string) (Frame, bool) {
	var id uint8
	var length uint8
	var end uint8
	var bytekey uint8
	//var payload []byte
	err := binary.Read(bytes.NewBuffer(readNextBytes(file, offset, 1)), binary.LittleEndian, &id)
	if err != nil {
		log.Fatal("binary.Read failed", err)
	}

	err = binary.Read(bytes.NewBuffer(readNextBytes(file, offset+1, 1)), binary.LittleEndian, &length)
	if err != nil {
		log.Fatal("binary.Read failed", err)
	}

	err = binary.Read(bytes.NewBuffer(readNextBytes(file, offset+2+int64(length), 1)), binary.LittleEndian, &end)
	if err != nil {
		log.Fatal("binary.Read failed", err)
	}

	if id != 0 && end == 0xFF {
		payload := make([]byte, length)
		err = binary.Read(bytes.NewBuffer(readNextBytes(file, offset+3, int(length))), binary.LittleEndian, &payload)
		if err != nil {
			log.Fatal("binary.Read failed", err)
		}

		err = binary.Read(bytes.NewBuffer(readNextBytes(file, offset+2, 1)), binary.LittleEndian, &bytekey)
		if err != nil {
			log.Fatal("binary.Read failed", err)
		}

		idKey := (int(id-1) * 256) + int(bytekey)
		if idKey > 4095 {
			return Frame{}, false
		}

		f := Frame{offset: offset, typeID: id, length: length, payload: payload, bytekey: bytekey, key: Keys[idKey], encrypted: true}
		//fmt.Printf("frame: %+v\n", f)
		return f, true
	}
	return Frame{}, false
}

func decryptFrame(f Frame, fl *Flight) *Flight {
	decryptedByte := decryptByteArray(f.payload, 0, len(f.payload), f.key)

	switch f.typeID {
	case 1:
		if len(decryptedByte) >= 53 {
			createOSD(decryptedByte)
		}
	case 2:
		if len(decryptedByte) >= 46 {
			createHome(decryptedByte)
		}
	case 3:
		createGimbal(decryptedByte)
	case 10:
		createAppMessage(decryptedByte)
	case 13:
		if len(decryptedByte) >= 85 {
			recover := createRecover(decryptedByte)
			if (recover != Recover{}) {
				fmt.Printf("Recover: %+v\n", recover)
			}
		}
	case 14:
		if len(decryptedByte) >= 20 {
			appGps := createAppGps(decryptedByte)
			fl.Gps = append(fl.Gps, GpsPoint{Latitude: appGps.Latitude, Longitude: appGps.Longitude})
		}

	}

	return fl
}

func decryptByteArray(payload []byte, offset int, length int, key [8]byte) []byte {
	decrypted := make([]byte, length)
	for i, b := range payload[offset:length] {
		decrypted[i] = b ^ byte(key[i%8])
	}
	return decrypted
}

func Float64frombytes(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	float := math.Float64frombits(bits)
	return float
}

func Float32frombytes(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)
	float := math.Float32frombits(bits)
	return float
}

func Intfrombytes(bytes []byte) int {
	return int(bytes[0])
}

func Int16frombytes(bytes []byte) int16 {
	bits := binary.LittleEndian.Uint16(bytes)
	return int16(bits)
}

func Int32frombytes(bytes []byte) int32 {
	bits := binary.LittleEndian.Uint32(bytes)
	return int32(bits)
}
