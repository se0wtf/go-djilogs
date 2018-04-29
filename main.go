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

func main() {
	path := "/home/seo/testv3.txt"

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
	fmt.Printf("The file is %d bytes long\n", fi.Size())
	fmt.Printf("%d\n", detailsAddr)
	fmt.Printf("%d\n", encrypted)
	fmt.Printf("%+v\n", details)
	fmt.Printf("Substreet: %s\n", details.SubStreet)
	fmt.Printf("Street: %s\n", details.Street)
	fmt.Printf("City: %s\n", details.City)
	fmt.Printf("Area: %s\n", details.Area)
	fmt.Printf("Longitude: %f\n", details.Longitude)
	fmt.Printf("Latitude: %f\n", details.Latitude)
	fmt.Printf("CameraSN: %s\n", details.CameraSn)
}

func readNextBytes(file *os.File, offset int64, lenght int) []byte {
	bytes := make([]byte, lenght)

	_, err := file.ReadAt(bytes, offset)
	if err != nil {
		log.Fatal(err)
	}

	return bytes
}
