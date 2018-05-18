package main

import "fmt"

type Recover struct {
	DroneType       string
	AppType         string
	AppVersion      string
	AircraftSN      string
	AircraftName    string
	ActiveTimestamp string
	CameraSN        string
	RcSN            string
	BatterySN       string
}

func createRecover(decrypted []byte) Recover {
	r := Recover{}

	if Intfrombytes(decrypted[0:1]) >= len(RecoverDroneType) {
		return r
	}

	r.DroneType = RecoverDroneType[Intfrombytes(decrypted[0:1])]

	if Intfrombytes(decrypted[1:2]) >= len(AppType) {
		r.AppType = "UNKNOWN"
	} else {
		r.AppType = AppType[Intfrombytes(decrypted[1:2])]
	}

	r.AppVersion = fmt.Sprintf("%d.%d.%d\n", Intfrombytes(decrypted[2:3]), Intfrombytes(decrypted[3:4]), Intfrombytes(decrypted[4:5]))
	r.AircraftSN = string(decrypted[5:15])

	r.AircraftName = string(decrypted[15:39])
	r.ActiveTimestamp = string(Int32frombytes(decrypted[47:55]))
	r.CameraSN = string(decrypted[55:65])
	r.RcSN = string(decrypted[65:75])
	r.BatterySN = string(decrypted[75:85])
	return r
}
