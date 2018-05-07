package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"os"
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
	// fmt.Printf("Longitude: %f\n", details.Longitude)
	// fmt.Printf("Latitude: %f\n", details.Latitude)

	for offset < detailsAddr {
		f, created := isFrame(file, int64(offset), types)
		if created {
			offset = int32(f.offset + int64(f.length) + 1 + 2)
			decryptFrame(f)
		} else {
			offset++
		}
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
		if idKey > 4096 {
			idKey = 0
		}

		f := Frame{offset: offset, typeID: id, length: length, payload: payload, bytekey: bytekey, key: Keys[idKey], encrypted: true}
		//fmt.Printf("> offset: %d, frame: %+v\n", offset, f)
		if offset > 80000 {
			os.Exit(-1)
		}
		// offset = 79938
		// if offset == 79938 {
		// 	decryptFrame(f)
		// }

		//offset == 758928 ||
		// if f.typeID == 7 {
		// 	//fmt.Printf("> offset: %d, offsetEnd: %d, id: %d, length: %d, end: %d, type: %s\n", offset, offset+2+int64(length), id, int(length), end, types[id])
		// 	fmt.Printf("> offset: %d, frame: %+v\n", offset, f)
		// 	decryptFrame(f)
		// }
		return f, true
	}
	return Frame{}, false
}

// byteKey: 96, key: 211,100,182,13,217,83,205,34, id: 096, dataOffset: 758931, dataLength: 55
func decryptFrame(f Frame) {
	// var tmpLong []byte
	// var decodecFloat float64
	// var decodecFloat2 float64
	// var decodecInt16 int16
	// var osd model.OSDRaw
	// CenterBatteryRaw
	// var raw AppGpsRaw
	// switch f.typeID {
	// case 14:
	// 	//var decoded [12]byte
	// 	if err := binary.Read(bytes.NewReader(f.payload), binary.LittleEndian, &raw); err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Printf("Raw: %+v\n", raw)
	// }

	//fmt.Printf("len: %d\n", binary.Size(decodecFloat))
	//var decoded byte
	//fmt.Printf("payload: %v\n", f.paylod)
	//tmpLong = f.payload[0:8]
	//fmt.Printf("tmpLong: %v\n", tmpLong)
	if f.typeID == 14 {
		//fmt.Printf("f: %+v\n", f)
		decryptedByte := decryptByteArray(f.payload, 0, len(f.payload), f.key)
		appGps := AppGps{}
		appGps.Longitude = Float64frombytes(decryptedByte[0:8])
		appGps.Latitude = Float64frombytes(decryptedByte[8:16])
		appGps.Accuracy = Float32frombytes(decryptedByte[16:20])
		fmt.Printf("appGps: %+v\n", appGps)
		//fmt.Printf("decrypt1: %d, decrypt2: %d, decrypt3: %d\n", decryptedByte[0:8], decryptedByte[8:16], decryptedByte[16:20])
		// bs := make([]byte, f.length)
		// for i, b := range f.payload {
		// 	fmt.Printf("i: %d, b: %d, key: %d\n", i, b, Keys[f.key][i])
		// 	//fmt.Printf("index: %d, byte: %b, key: %d, decodedByte: %d\n", i, b, Keys[f.key][i], b^byte(Keys[f.key][i]))

		// 	bs[i] = b ^ byte(Keys[f.key][i])
		// }
	}
	//tmpLong = decryptByteArray(f.payload, 0, 8, f.key)

	// -------- OSD TEST
	// if err := binary.Read(bytes.NewReader(tmpLong), binary.LittleEndian, &decodecFloat); err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("decoded: %v, float: %f\n", tmpLong, decodecFloat*180/math.Pi)

	// -------- OSD TEST
	// s := reflect.ValueOf(&osd).Elem()
	// for i := 0; i < s.NumField(); i++ {
	// 	field := s.Field(i)
	// 	switch field.Kind() {
	// 	case reflect.Float64:
	// 		buf := make([]byte, 8)
	// 		binary.LittleEndian.PutUint64(buf, math.Float64bits(field.Float()))
	// 		tmpLong2 := decryptByteArray(buf, 0, len(buf), f.key)
	// 		if err := binary.Read(bytes.NewReader(tmpLong2), binary.LittleEndian, &decodecFloat2); err != nil {
	// 			log.Fatal(err)
	// 		}
	// 		field.SetFloat(decodecFloat2 * 180 / math.Pi)
	// 	case reflect.Int16:
	// 		buf := make([]byte, 2)
	// 		binary.LittleEndian.PutUint16(buf, uint16(field.Int()))
	// 		tmpLong2 := decryptByteArray(buf, 0, len(buf), f.key)
	// 		if err := binary.Read(bytes.NewReader(tmpLong2), binary.LittleEndian, &decodecInt16); err != nil {
	// 			log.Fatal(err)
	// 		}
	// 		field.SetInt(int64(decodecInt16) / 10)
	// 	}
	// }

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
