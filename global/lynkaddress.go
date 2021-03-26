package globals

import "strings"

type LynkData struct {
	Address   string
	GSTN      string
	StateCode string
	StateName string
}

var GetLynkAddressDetails = func(storeName string) (*LynkData, error) {
	if strings.Contains(strings.ToLower(storeName), Hyderabad) {
		return &LynkData{
			Address:   "8-3-976/87, PLOT NO 87/87A, SHALIVAHAN NAGAR, SRINAGAR COLONY, Hyderabad, Telangana, 500073",
			GSTN:      TelenganaGSTN,
			StateCode: TelenganaStateCode,
			StateName: "Telengana",
		}, nil
	} else if strings.Contains(strings.ToLower(storeName), Chennai) {
		return &LynkData{
			Address:   "108A, VM STREET ,94 DR RADHAKRISHNAN SALAI, MYLAPORE, Chennai, Tamil Nadu, 600004",
			GSTN:      TamilNaduGSTN,
			StateCode: TamilNaduStateCode,
			StateName: "Tamilnadu",
		}, nil
	} else if strings.Contains(strings.ToLower(storeName), Mumbai) {
		return &LynkData{
			Address:   "THIRD FLOOR 310 ASCOT CENTER, SAHAR, ANDHERI EAST MUMBAI MAHARASHTRA, 400099",
			GSTN:      MumbaiGSTN,
			StateCode: MaharasthtraStateCode,
			StateName: "Maharashtra",
		}, nil
	} else if strings.Contains(strings.ToLower(storeName), Pune) {
		return &LynkData{
			Address:   "THIRD FLOOR 310 ASCOT CENTER, SAHAR, ANDHERI EAST MUMBAI MAHARASHTRA, 400099",
			GSTN:      MumbaiGSTN,
			StateCode: MaharasthtraStateCode,
			StateName: "Maharashtra",
		}, nil
	} else if strings.Contains(strings.ToLower(storeName), Bangalore) {
		return &LynkData{
			Address:   "3rd Floor, 1/58, Sycon Polaris, 8Th Main, RMV Extension, Sadashivanagar, Bangalore, Bengaluru (Bangalore) Urban, Karnataka, 560080",
			GSTN:      BangaloreGSTN,
			StateCode: BangaloreStateCode,
			StateName: "Karnataka",
		}, nil
	} else {
		return &LynkData{
			Address:   "108A, VM STREET ,94 DR RADHAKRISHNAN SALAI, MYLAPORE, Chennai, Tamil Nadu, 600004",
			GSTN:      TamilNaduGSTN,
			StateCode: TamilNaduStateCode,
			StateName: "Tamilnadu",
		}, nil
	}
}
