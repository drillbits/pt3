package pt3

const (
	channelTypeSatellite = 0
	channelTypeGround    = 1
)

var bsDevices = []string{
	"/dev/pt3video1",
	"/dev/pt3video0",
	"/dev/pt3video5",
	"/dev/pt3video4",
}

var isdbtDevices = []string{
	"/dev/pt3video2",
	"/dev/pt3video3",
	"/dev/pt3video6",
	"/dev/pt3video7",
}

type isdbtChannel struct {
	setFreq     int
	channelType int
	addFreq     int
}

var isdbtConvertTable map[string]*isdbtChannel

func init() {
	isdbtConvertTable = make(map[string]*isdbtChannel, 0)
	isdbtConvertTable["151"] = &isdbtChannel{setFreq: 0, channelType: channelTypeSatellite, addFreq: 0}
	isdbtConvertTable["161"] = &isdbtChannel{setFreq: 0, channelType: channelTypeSatellite, addFreq: 1}
	isdbtConvertTable["191"] = &isdbtChannel{setFreq: 1, channelType: channelTypeSatellite, addFreq: 0}
	isdbtConvertTable["171"] = &isdbtChannel{setFreq: 1, channelType: channelTypeSatellite, addFreq: 1}
	isdbtConvertTable["192"] = &isdbtChannel{setFreq: 2, channelType: channelTypeSatellite, addFreq: 0}
	isdbtConvertTable["193"] = &isdbtChannel{setFreq: 2, channelType: channelTypeSatellite, addFreq: 1}
	isdbtConvertTable["201"] = &isdbtChannel{setFreq: 3, channelType: channelTypeSatellite, addFreq: 0}
	isdbtConvertTable["202"] = &isdbtChannel{setFreq: 3, channelType: channelTypeSatellite, addFreq: 0}
	isdbtConvertTable["236"] = &isdbtChannel{setFreq: 3, channelType: channelTypeSatellite, addFreq: 1}
	isdbtConvertTable["256"] = &isdbtChannel{setFreq: 3, channelType: channelTypeSatellite, addFreq: 2}
	isdbtConvertTable["211"] = &isdbtChannel{setFreq: 4, channelType: channelTypeSatellite, addFreq: 0}
	isdbtConvertTable["200"] = &isdbtChannel{setFreq: 4, channelType: channelTypeSatellite, addFreq: 1}
	isdbtConvertTable["222"] = &isdbtChannel{setFreq: 4, channelType: channelTypeSatellite, addFreq: 2}
	isdbtConvertTable["238"] = &isdbtChannel{setFreq: 5, channelType: channelTypeSatellite, addFreq: 0}
	isdbtConvertTable["241"] = &isdbtChannel{setFreq: 5, channelType: channelTypeSatellite, addFreq: 1}
	isdbtConvertTable["231"] = &isdbtChannel{setFreq: 5, channelType: channelTypeSatellite, addFreq: 2}
	isdbtConvertTable["232"] = &isdbtChannel{setFreq: 5, channelType: channelTypeSatellite, addFreq: 2}
	isdbtConvertTable["233"] = &isdbtChannel{setFreq: 5, channelType: channelTypeSatellite, addFreq: 2}
	isdbtConvertTable["531"] = &isdbtChannel{setFreq: 5, channelType: channelTypeSatellite, addFreq: 2}
	isdbtConvertTable["141"] = &isdbtChannel{setFreq: 6, channelType: channelTypeSatellite, addFreq: 0}
	isdbtConvertTable["181"] = &isdbtChannel{setFreq: 6, channelType: channelTypeSatellite, addFreq: 1}
	isdbtConvertTable["101"] = &isdbtChannel{setFreq: 7, channelType: channelTypeSatellite, addFreq: 0}
	isdbtConvertTable["102"] = &isdbtChannel{setFreq: 7, channelType: channelTypeSatellite, addFreq: 0}
	isdbtConvertTable["103"] = &isdbtChannel{setFreq: 7, channelType: channelTypeSatellite, addFreq: 1}
	isdbtConvertTable["910"] = &isdbtChannel{setFreq: 7, channelType: channelTypeSatellite, addFreq: 1}
	isdbtConvertTable["291"] = &isdbtChannel{setFreq: 8, channelType: channelTypeSatellite, addFreq: 2}
	isdbtConvertTable["292"] = &isdbtChannel{setFreq: 8, channelType: channelTypeSatellite, addFreq: 2}
	isdbtConvertTable["294"] = &isdbtChannel{setFreq: 8, channelType: channelTypeSatellite, addFreq: 1}
	isdbtConvertTable["295"] = &isdbtChannel{setFreq: 8, channelType: channelTypeSatellite, addFreq: 1}
	isdbtConvertTable["296"] = &isdbtChannel{setFreq: 8, channelType: channelTypeSatellite, addFreq: 1}
	isdbtConvertTable["297"] = &isdbtChannel{setFreq: 8, channelType: channelTypeSatellite, addFreq: 1}
	isdbtConvertTable["298"] = &isdbtChannel{setFreq: 8, channelType: channelTypeSatellite, addFreq: 2}
	isdbtConvertTable["234"] = &isdbtChannel{setFreq: 9, channelType: channelTypeSatellite, addFreq: 0}
	isdbtConvertTable["242"] = &isdbtChannel{setFreq: 9, channelType: channelTypeSatellite, addFreq: 1}
	isdbtConvertTable["243"] = &isdbtChannel{setFreq: 9, channelType: channelTypeSatellite, addFreq: 2}
	isdbtConvertTable["252"] = &isdbtChannel{setFreq: 10, channelType: channelTypeSatellite, addFreq: 0}
	isdbtConvertTable["244"] = &isdbtChannel{setFreq: 10, channelType: channelTypeSatellite, addFreq: 1}
	isdbtConvertTable["245"] = &isdbtChannel{setFreq: 10, channelType: channelTypeSatellite, addFreq: 2}
	isdbtConvertTable["251"] = &isdbtChannel{setFreq: 11, channelType: channelTypeSatellite, addFreq: 0}
	isdbtConvertTable["255"] = &isdbtChannel{setFreq: 11, channelType: channelTypeSatellite, addFreq: 1}
	isdbtConvertTable["258"] = &isdbtChannel{setFreq: 11, channelType: channelTypeSatellite, addFreq: 2}
	isdbtConvertTable["CS2"] = &isdbtChannel{setFreq: 12, channelType: channelTypeSatellite, addFreq: 0}
	isdbtConvertTable["CS4"] = &isdbtChannel{setFreq: 13, channelType: channelTypeSatellite, addFreq: 0}
	isdbtConvertTable["CS6"] = &isdbtChannel{setFreq: 14, channelType: channelTypeSatellite, addFreq: 0}
	isdbtConvertTable["CS8"] = &isdbtChannel{setFreq: 15, channelType: channelTypeSatellite, addFreq: 0}
	isdbtConvertTable["CS10"] = &isdbtChannel{setFreq: 16, channelType: channelTypeSatellite, addFreq: 0}
	isdbtConvertTable["CS12"] = &isdbtChannel{setFreq: 17, channelType: channelTypeSatellite, addFreq: 0}
	isdbtConvertTable["CS14"] = &isdbtChannel{setFreq: 18, channelType: channelTypeSatellite, addFreq: 0}
	isdbtConvertTable["CS16"] = &isdbtChannel{setFreq: 19, channelType: channelTypeSatellite, addFreq: 0}
	isdbtConvertTable["CS18"] = &isdbtChannel{setFreq: 20, channelType: channelTypeSatellite, addFreq: 0}
	isdbtConvertTable["CS20"] = &isdbtChannel{setFreq: 21, channelType: channelTypeSatellite, addFreq: 0}
	isdbtConvertTable["CS22"] = &isdbtChannel{setFreq: 22, channelType: channelTypeSatellite, addFreq: 0}
	isdbtConvertTable["CS24"] = &isdbtChannel{setFreq: 23, channelType: channelTypeSatellite, addFreq: 0}
	isdbtConvertTable["1"] = &isdbtChannel{setFreq: 0, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["2"] = &isdbtChannel{setFreq: 1, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["3"] = &isdbtChannel{setFreq: 2, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["C13"] = &isdbtChannel{setFreq: 3, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["C14"] = &isdbtChannel{setFreq: 4, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["C15"] = &isdbtChannel{setFreq: 5, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["C16"] = &isdbtChannel{setFreq: 6, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["C17"] = &isdbtChannel{setFreq: 7, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["C18"] = &isdbtChannel{setFreq: 8, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["C19"] = &isdbtChannel{setFreq: 9, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["C20"] = &isdbtChannel{setFreq: 10, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["C21"] = &isdbtChannel{setFreq: 11, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["C22"] = &isdbtChannel{setFreq: 12, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["4"] = &isdbtChannel{setFreq: 13, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["5"] = &isdbtChannel{setFreq: 14, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["6"] = &isdbtChannel{setFreq: 15, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["7"] = &isdbtChannel{setFreq: 16, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["8"] = &isdbtChannel{setFreq: 17, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["9"] = &isdbtChannel{setFreq: 18, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["10"] = &isdbtChannel{setFreq: 19, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["11"] = &isdbtChannel{setFreq: 20, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["12"] = &isdbtChannel{setFreq: 21, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["C23"] = &isdbtChannel{setFreq: 22, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["C24"] = &isdbtChannel{setFreq: 23, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["C25"] = &isdbtChannel{setFreq: 24, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["C26"] = &isdbtChannel{setFreq: 25, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["C27"] = &isdbtChannel{setFreq: 26, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["C28"] = &isdbtChannel{setFreq: 27, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["C29"] = &isdbtChannel{setFreq: 28, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["C30"] = &isdbtChannel{setFreq: 29, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["C31"] = &isdbtChannel{setFreq: 30, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["C32"] = &isdbtChannel{setFreq: 31, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["C33"] = &isdbtChannel{setFreq: 32, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["C34"] = &isdbtChannel{setFreq: 33, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["C35"] = &isdbtChannel{setFreq: 34, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["C36"] = &isdbtChannel{setFreq: 35, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["C37"] = &isdbtChannel{setFreq: 36, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["C38"] = &isdbtChannel{setFreq: 37, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["C39"] = &isdbtChannel{setFreq: 38, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["C40"] = &isdbtChannel{setFreq: 39, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["C41"] = &isdbtChannel{setFreq: 40, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["C42"] = &isdbtChannel{setFreq: 41, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["C43"] = &isdbtChannel{setFreq: 42, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["C44"] = &isdbtChannel{setFreq: 43, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["C45"] = &isdbtChannel{setFreq: 44, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["C46"] = &isdbtChannel{setFreq: 45, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["C47"] = &isdbtChannel{setFreq: 46, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["C48"] = &isdbtChannel{setFreq: 47, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["C49"] = &isdbtChannel{setFreq: 48, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["C50"] = &isdbtChannel{setFreq: 49, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["C51"] = &isdbtChannel{setFreq: 50, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["C52"] = &isdbtChannel{setFreq: 51, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["C53"] = &isdbtChannel{setFreq: 52, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["C54"] = &isdbtChannel{setFreq: 53, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["C55"] = &isdbtChannel{setFreq: 54, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["C56"] = &isdbtChannel{setFreq: 55, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["C57"] = &isdbtChannel{setFreq: 56, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["C58"] = &isdbtChannel{setFreq: 57, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["C59"] = &isdbtChannel{setFreq: 58, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["C60"] = &isdbtChannel{setFreq: 59, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["C61"] = &isdbtChannel{setFreq: 60, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["C62"] = &isdbtChannel{setFreq: 61, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["C63"] = &isdbtChannel{setFreq: 62, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["13"] = &isdbtChannel{setFreq: 63, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["14"] = &isdbtChannel{setFreq: 64, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["15"] = &isdbtChannel{setFreq: 65, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["16"] = &isdbtChannel{setFreq: 66, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["17"] = &isdbtChannel{setFreq: 67, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["18"] = &isdbtChannel{setFreq: 68, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["19"] = &isdbtChannel{setFreq: 69, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["20"] = &isdbtChannel{setFreq: 70, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["21"] = &isdbtChannel{setFreq: 71, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["22"] = &isdbtChannel{setFreq: 72, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["23"] = &isdbtChannel{setFreq: 73, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["24"] = &isdbtChannel{setFreq: 74, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["25"] = &isdbtChannel{setFreq: 75, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["26"] = &isdbtChannel{setFreq: 76, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["27"] = &isdbtChannel{setFreq: 77, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["28"] = &isdbtChannel{setFreq: 78, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["29"] = &isdbtChannel{setFreq: 79, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["30"] = &isdbtChannel{setFreq: 80, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["31"] = &isdbtChannel{setFreq: 81, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["32"] = &isdbtChannel{setFreq: 82, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["33"] = &isdbtChannel{setFreq: 83, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["34"] = &isdbtChannel{setFreq: 84, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["35"] = &isdbtChannel{setFreq: 85, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["36"] = &isdbtChannel{setFreq: 86, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["37"] = &isdbtChannel{setFreq: 87, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["38"] = &isdbtChannel{setFreq: 88, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["39"] = &isdbtChannel{setFreq: 89, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["40"] = &isdbtChannel{setFreq: 90, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["41"] = &isdbtChannel{setFreq: 91, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["42"] = &isdbtChannel{setFreq: 92, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["43"] = &isdbtChannel{setFreq: 93, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["44"] = &isdbtChannel{setFreq: 94, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["45"] = &isdbtChannel{setFreq: 95, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["46"] = &isdbtChannel{setFreq: 96, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["47"] = &isdbtChannel{setFreq: 97, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["48"] = &isdbtChannel{setFreq: 98, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["49"] = &isdbtChannel{setFreq: 99, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["50"] = &isdbtChannel{setFreq: 100, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["51"] = &isdbtChannel{setFreq: 101, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["52"] = &isdbtChannel{setFreq: 102, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["53"] = &isdbtChannel{setFreq: 103, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["54"] = &isdbtChannel{setFreq: 104, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["55"] = &isdbtChannel{setFreq: 105, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["56"] = &isdbtChannel{setFreq: 106, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["57"] = &isdbtChannel{setFreq: 107, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["58"] = &isdbtChannel{setFreq: 108, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["59"] = &isdbtChannel{setFreq: 109, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["60"] = &isdbtChannel{setFreq: 110, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["61"] = &isdbtChannel{setFreq: 111, channelType: channelTypeGround, addFreq: 0}
	isdbtConvertTable["62"] = &isdbtChannel{setFreq: 112, channelType: channelTypeGround, addFreq: 0}
}
