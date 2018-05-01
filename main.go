package main

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"log"
	"math"
	"os"
	"reflect"
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
	Useless           int16
	FlycVersion       int8
	DroneType         int8
	IMUinitFailReason int8
	MotorFailReason   int8
	Useless2          int8
	SDKCtrlDevice     int8
}

type Frame struct {
	offset    int64
	typeID    uint8
	length    uint8
	key       uint8
	payload   []byte
	encrypted bool
}

func main() {
	path := "testv3.txt"

	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Error while opening file", err)
	}
	defer file.Close()
	fmt.Printf("%s opened\n", path)

	details := Details{}

	// init types
	types := make([]string, 256)
	//types = append(types, "OSD", "HOME", "GIMBAL", "RC", "CUSTOM", "DEFORM", "CENTER_BATTERY", "SMART_BATTERY", "APP_TIP", "APP_WARN", "RC_GPS", "RC_DEBUG", "RECOVER", "APP_GPS", "FIRMWARE", "OFDM_DEBUG", "VISION_GROUP", "VISION_WARN", "MC_PARAM", "APP_OPERATION")
	//types = []string{"OSD", "HOME", "GIMBAL", "RC", "CUSTOM", "DEFORM", "CENTER_BATTERY", "SMART_BATTERY", "APP_TIP", "APP_WARN", "RC_GPS", "RC_DEBUG", "RECOVER", "APP_GPS", "FIRMWARE", "OFDM_DEBUG", "VISION_GROUP", "VISION_WARN", "MC_PARAM", "APP_OPERATION"}
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

	//1433907
	fmt.Printf("Longitude: %f\n", details.Longitude)
	fmt.Printf("Latitude: %f\n", details.Latitude)

	for offset < detailsAddr {
		isFrame(file, int64(offset), types)
		offset++
	}
}

func readNextBytes(file *os.File, offset int64, length int) []byte {
	bytes := make([]byte, length)
	_, err := file.ReadAt(bytes, offset)
	if err != nil {
		log.Fatal(err)
	}

	return bytes
}

func isFrame(file *os.File, offset int64, types []string) Frame {
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

	if id != 0 && end == 255 && types[id] != "" {
		payload := make([]byte, length)
		err = binary.Read(bytes.NewBuffer(readNextBytes(file, offset+3, int(length))), binary.LittleEndian, &payload)
		if err != nil {
			log.Fatal("binary.Read failed", err)
		}

		err = binary.Read(bytes.NewBuffer(readNextBytes(file, offset+2, 1)), binary.LittleEndian, &bytekey)
		if err != nil {
			log.Fatal("binary.Read failed", err)
		}

		f := Frame{offset: offset, typeID: id, length: length, payload: payload, key: bytekey, encrypted: true}
		if offset == 758928 {
			//fmt.Printf("> offset: %d, offsetEnd: %d, id: %d, length: %d, end: %d, type: %s\n", offset, offset+2+int64(length), id, int(length), end, types[id])
			fmt.Printf("> offset: %d, frame: %+v\n", offset, f)
			decryptFrame(f)
		}
		return f
	}
	return Frame{}
}

// byteKey: 96, key: 211,100,182,13,217,83,205,34, id: 096, dataOffset: 758931, dataLength: 55
func decryptFrame(f Frame) {
	var tmpLong []byte
	var decodecFloat float64
	var decodecFloat2 float64
	var osd OSDRaw
	//fmt.Printf("len: %d\n", binary.Size(decodecFloat))
	//var decoded byte
	//fmt.Printf("payload: %v\n", f.paylod)
	//tmpLong = f.payload[0:8]
	//fmt.Printf("tmpLong: %v\n", tmpLong)
	// bs := make([]byte, 8)
	// for i, b := range tmpLong {
	// 	//fmt.Printf("index: %d, byte: %b, key: %d, decodedByte: %d\n", i, b, Keys[f.key][i], b^byte(Keys[f.key][i]))
	// 	bs[i] = b ^ byte(Keys[f.key][i])
	// }
	tmpLong = decryptByteArray(f.payload, 0, 8, f.key)
	if err := binary.Read(bytes.NewReader(tmpLong), binary.LittleEndian, &decodecFloat); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("decoded: %v, float: %f\n", tmpLong, decodecFloat*180/math.Pi)
	//var long float64

	if err := binary.Read(bytes.NewReader(f.payload), binary.LittleEndian, &osd); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("OSD: %+v\n", osd)

	s := reflect.ValueOf(&osd).Elem()
	//v.NumField()
	//typeOfT := s.Type()

	for i := 0; i < s.NumField(); i++ {
		field := s.Field(i)
		//fmt.Printf("size: %d \n", binary.Size(field.Type()))
		switch field.Interface().(type) {
		case float64:
			//fmt.Printf("float: %f\n", field.Interface())
			// var buf [8]byte
			// binary.LittleEndian.PutUint64(buf[:], math.Float64bits(field))
			var buf bytes.Buffer
			enc := gob.NewEncoder(&buf)
			enc.Encode(field.Interface())
			//fmt.Printf("bytes: %v\n", buf.Bytes())
			tmpLong2 := decryptByteArray(buf.Bytes(), 4, 12, f.key)
			if err := binary.Read(bytes.NewReader(tmpLong2), binary.LittleEndian, &decodecFloat2); err != nil {
				log.Fatal(err)
			}
			fmt.Printf("decoded: %v, float: %f\n", tmpLong2, decodecFloat2*180/math.Pi)
		}
		// fmt.Printf("%d: %s %s = %v\n", i,
		// 	typeOfT.Field(i).Name, field.Type(), field.Interface())
	}
}

func decryptByteArray(payload []byte, offset int, length int, keyID uint8) []byte {
	fmt.Printf("\ndecryptByteArray :: payload: %v, offset: %d, length: %d, keyID: %d\n", payload[offset:length], offset, length, keyID)
	decrypted := make([]byte, length)
	for i, b := range payload[offset:length] {
		//fmt.Printf("i: %d, b: %b\n", i, b)
		decrypted[i] = b ^ byte(Keys[keyID][i])
	}
	return decrypted
}
