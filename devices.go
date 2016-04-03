package pt3

const (
	ChannelTypeSatellite = 0
	ChannelTypeGround    = 1
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
	isdbtConvertTable["151"] = &isdbtChannel{setFreq: 0, channelType: ChannelTypeSatellite, addFreq: 0}
	isdbtConvertTable["161"] = &isdbtChannel{setFreq: 0, channelType: ChannelTypeSatellite, addFreq: 1}
	isdbtConvertTable["191"] = &isdbtChannel{setFreq: 1, channelType: ChannelTypeSatellite, addFreq: 0}
	isdbtConvertTable["171"] = &isdbtChannel{setFreq: 1, channelType: ChannelTypeSatellite, addFreq: 1}
	isdbtConvertTable["192"] = &isdbtChannel{setFreq: 2, channelType: ChannelTypeSatellite, addFreq: 0}
	isdbtConvertTable["193"] = &isdbtChannel{setFreq: 2, channelType: ChannelTypeSatellite, addFreq: 1}
	isdbtConvertTable["201"] = &isdbtChannel{setFreq: 3, channelType: ChannelTypeSatellite, addFreq: 0}
	isdbtConvertTable["202"] = &isdbtChannel{setFreq: 3, channelType: ChannelTypeSatellite, addFreq: 0}
	isdbtConvertTable["236"] = &isdbtChannel{setFreq: 3, channelType: ChannelTypeSatellite, addFreq: 1}
	isdbtConvertTable["256"] = &isdbtChannel{setFreq: 3, channelType: ChannelTypeSatellite, addFreq: 2}
	isdbtConvertTable["211"] = &isdbtChannel{setFreq: 4, channelType: ChannelTypeSatellite, addFreq: 0}
	isdbtConvertTable["200"] = &isdbtChannel{setFreq: 4, channelType: ChannelTypeSatellite, addFreq: 1}
	isdbtConvertTable["222"] = &isdbtChannel{setFreq: 4, channelType: ChannelTypeSatellite, addFreq: 2}
	isdbtConvertTable["238"] = &isdbtChannel{setFreq: 5, channelType: ChannelTypeSatellite, addFreq: 0}
	isdbtConvertTable["241"] = &isdbtChannel{setFreq: 5, channelType: ChannelTypeSatellite, addFreq: 1}
	isdbtConvertTable["231"] = &isdbtChannel{setFreq: 5, channelType: ChannelTypeSatellite, addFreq: 2}
	isdbtConvertTable["232"] = &isdbtChannel{setFreq: 5, channelType: ChannelTypeSatellite, addFreq: 2}
	isdbtConvertTable["233"] = &isdbtChannel{setFreq: 5, channelType: ChannelTypeSatellite, addFreq: 2}
	isdbtConvertTable["531"] = &isdbtChannel{setFreq: 5, channelType: ChannelTypeSatellite, addFreq: 2}
	isdbtConvertTable["141"] = &isdbtChannel{setFreq: 6, channelType: ChannelTypeSatellite, addFreq: 0}
	isdbtConvertTable["181"] = &isdbtChannel{setFreq: 6, channelType: ChannelTypeSatellite, addFreq: 1}
	isdbtConvertTable["101"] = &isdbtChannel{setFreq: 7, channelType: ChannelTypeSatellite, addFreq: 0}
	isdbtConvertTable["102"] = &isdbtChannel{setFreq: 7, channelType: ChannelTypeSatellite, addFreq: 0}
	isdbtConvertTable["103"] = &isdbtChannel{setFreq: 7, channelType: ChannelTypeSatellite, addFreq: 1}
	isdbtConvertTable["910"] = &isdbtChannel{setFreq: 7, channelType: ChannelTypeSatellite, addFreq: 1}
	isdbtConvertTable["291"] = &isdbtChannel{setFreq: 8, channelType: ChannelTypeSatellite, addFreq: 2}
	isdbtConvertTable["292"] = &isdbtChannel{setFreq: 8, channelType: ChannelTypeSatellite, addFreq: 2}
	isdbtConvertTable["294"] = &isdbtChannel{setFreq: 8, channelType: ChannelTypeSatellite, addFreq: 1}
	isdbtConvertTable["295"] = &isdbtChannel{setFreq: 8, channelType: ChannelTypeSatellite, addFreq: 1}
	isdbtConvertTable["296"] = &isdbtChannel{setFreq: 8, channelType: ChannelTypeSatellite, addFreq: 1}
	isdbtConvertTable["297"] = &isdbtChannel{setFreq: 8, channelType: ChannelTypeSatellite, addFreq: 1}
	isdbtConvertTable["298"] = &isdbtChannel{setFreq: 8, channelType: ChannelTypeSatellite, addFreq: 2}
	isdbtConvertTable["234"] = &isdbtChannel{setFreq: 9, channelType: ChannelTypeSatellite, addFreq: 0}
	isdbtConvertTable["242"] = &isdbtChannel{setFreq: 9, channelType: ChannelTypeSatellite, addFreq: 1}
	isdbtConvertTable["243"] = &isdbtChannel{setFreq: 9, channelType: ChannelTypeSatellite, addFreq: 2}
	isdbtConvertTable["252"] = &isdbtChannel{setFreq: 10, channelType: ChannelTypeSatellite, addFreq: 0}
	isdbtConvertTable["244"] = &isdbtChannel{setFreq: 10, channelType: ChannelTypeSatellite, addFreq: 1}
	isdbtConvertTable["245"] = &isdbtChannel{setFreq: 10, channelType: ChannelTypeSatellite, addFreq: 2}
	isdbtConvertTable["251"] = &isdbtChannel{setFreq: 11, channelType: ChannelTypeSatellite, addFreq: 0}
	isdbtConvertTable["255"] = &isdbtChannel{setFreq: 11, channelType: ChannelTypeSatellite, addFreq: 1}
	isdbtConvertTable["258"] = &isdbtChannel{setFreq: 11, channelType: ChannelTypeSatellite, addFreq: 2}
	isdbtConvertTable["CS2"] = &isdbtChannel{setFreq: 12, channelType: ChannelTypeSatellite, addFreq: 0}
	isdbtConvertTable["CS4"] = &isdbtChannel{setFreq: 13, channelType: ChannelTypeSatellite, addFreq: 0}
	isdbtConvertTable["CS6"] = &isdbtChannel{setFreq: 14, channelType: ChannelTypeSatellite, addFreq: 0}
	isdbtConvertTable["CS8"] = &isdbtChannel{setFreq: 15, channelType: ChannelTypeSatellite, addFreq: 0}
	isdbtConvertTable["CS10"] = &isdbtChannel{setFreq: 16, channelType: ChannelTypeSatellite, addFreq: 0}
	isdbtConvertTable["CS12"] = &isdbtChannel{setFreq: 17, channelType: ChannelTypeSatellite, addFreq: 0}
	isdbtConvertTable["CS14"] = &isdbtChannel{setFreq: 18, channelType: ChannelTypeSatellite, addFreq: 0}
	isdbtConvertTable["CS16"] = &isdbtChannel{setFreq: 19, channelType: ChannelTypeSatellite, addFreq: 0}
	isdbtConvertTable["CS18"] = &isdbtChannel{setFreq: 20, channelType: ChannelTypeSatellite, addFreq: 0}
	isdbtConvertTable["CS20"] = &isdbtChannel{setFreq: 21, channelType: ChannelTypeSatellite, addFreq: 0}
	isdbtConvertTable["CS22"] = &isdbtChannel{setFreq: 22, channelType: ChannelTypeSatellite, addFreq: 0}
	isdbtConvertTable["CS24"] = &isdbtChannel{setFreq: 23, channelType: ChannelTypeSatellite, addFreq: 0}
	isdbtConvertTable["1"] = &isdbtChannel{setFreq: 0, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["2"] = &isdbtChannel{setFreq: 1, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["3"] = &isdbtChannel{setFreq: 2, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["C13"] = &isdbtChannel{setFreq: 3, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["C14"] = &isdbtChannel{setFreq: 4, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["C15"] = &isdbtChannel{setFreq: 5, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["C16"] = &isdbtChannel{setFreq: 6, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["C17"] = &isdbtChannel{setFreq: 7, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["C18"] = &isdbtChannel{setFreq: 8, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["C19"] = &isdbtChannel{setFreq: 9, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["C20"] = &isdbtChannel{setFreq: 10, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["C21"] = &isdbtChannel{setFreq: 11, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["C22"] = &isdbtChannel{setFreq: 12, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["4"] = &isdbtChannel{setFreq: 13, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["5"] = &isdbtChannel{setFreq: 14, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["6"] = &isdbtChannel{setFreq: 15, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["7"] = &isdbtChannel{setFreq: 16, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["8"] = &isdbtChannel{setFreq: 17, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["9"] = &isdbtChannel{setFreq: 18, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["10"] = &isdbtChannel{setFreq: 19, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["11"] = &isdbtChannel{setFreq: 20, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["12"] = &isdbtChannel{setFreq: 21, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["C23"] = &isdbtChannel{setFreq: 22, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["C24"] = &isdbtChannel{setFreq: 23, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["C25"] = &isdbtChannel{setFreq: 24, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["C26"] = &isdbtChannel{setFreq: 25, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["C27"] = &isdbtChannel{setFreq: 26, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["C28"] = &isdbtChannel{setFreq: 27, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["C29"] = &isdbtChannel{setFreq: 28, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["C30"] = &isdbtChannel{setFreq: 29, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["C31"] = &isdbtChannel{setFreq: 30, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["C32"] = &isdbtChannel{setFreq: 31, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["C33"] = &isdbtChannel{setFreq: 32, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["C34"] = &isdbtChannel{setFreq: 33, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["C35"] = &isdbtChannel{setFreq: 34, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["C36"] = &isdbtChannel{setFreq: 35, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["C37"] = &isdbtChannel{setFreq: 36, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["C38"] = &isdbtChannel{setFreq: 37, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["C39"] = &isdbtChannel{setFreq: 38, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["C40"] = &isdbtChannel{setFreq: 39, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["C41"] = &isdbtChannel{setFreq: 40, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["C42"] = &isdbtChannel{setFreq: 41, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["C43"] = &isdbtChannel{setFreq: 42, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["C44"] = &isdbtChannel{setFreq: 43, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["C45"] = &isdbtChannel{setFreq: 44, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["C46"] = &isdbtChannel{setFreq: 45, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["C47"] = &isdbtChannel{setFreq: 46, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["C48"] = &isdbtChannel{setFreq: 47, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["C49"] = &isdbtChannel{setFreq: 48, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["C50"] = &isdbtChannel{setFreq: 49, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["C51"] = &isdbtChannel{setFreq: 50, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["C52"] = &isdbtChannel{setFreq: 51, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["C53"] = &isdbtChannel{setFreq: 52, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["C54"] = &isdbtChannel{setFreq: 53, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["C55"] = &isdbtChannel{setFreq: 54, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["C56"] = &isdbtChannel{setFreq: 55, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["C57"] = &isdbtChannel{setFreq: 56, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["C58"] = &isdbtChannel{setFreq: 57, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["C59"] = &isdbtChannel{setFreq: 58, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["C60"] = &isdbtChannel{setFreq: 59, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["C61"] = &isdbtChannel{setFreq: 60, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["C62"] = &isdbtChannel{setFreq: 61, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["C63"] = &isdbtChannel{setFreq: 62, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["13"] = &isdbtChannel{setFreq: 63, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["14"] = &isdbtChannel{setFreq: 64, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["15"] = &isdbtChannel{setFreq: 65, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["16"] = &isdbtChannel{setFreq: 66, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["17"] = &isdbtChannel{setFreq: 67, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["18"] = &isdbtChannel{setFreq: 68, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["19"] = &isdbtChannel{setFreq: 69, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["20"] = &isdbtChannel{setFreq: 70, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["21"] = &isdbtChannel{setFreq: 71, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["22"] = &isdbtChannel{setFreq: 72, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["23"] = &isdbtChannel{setFreq: 73, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["24"] = &isdbtChannel{setFreq: 74, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["25"] = &isdbtChannel{setFreq: 75, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["26"] = &isdbtChannel{setFreq: 76, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["27"] = &isdbtChannel{setFreq: 77, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["28"] = &isdbtChannel{setFreq: 78, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["29"] = &isdbtChannel{setFreq: 79, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["30"] = &isdbtChannel{setFreq: 80, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["31"] = &isdbtChannel{setFreq: 81, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["32"] = &isdbtChannel{setFreq: 82, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["33"] = &isdbtChannel{setFreq: 83, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["34"] = &isdbtChannel{setFreq: 84, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["35"] = &isdbtChannel{setFreq: 85, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["36"] = &isdbtChannel{setFreq: 86, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["37"] = &isdbtChannel{setFreq: 87, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["38"] = &isdbtChannel{setFreq: 88, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["39"] = &isdbtChannel{setFreq: 89, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["40"] = &isdbtChannel{setFreq: 90, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["41"] = &isdbtChannel{setFreq: 91, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["42"] = &isdbtChannel{setFreq: 92, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["43"] = &isdbtChannel{setFreq: 93, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["44"] = &isdbtChannel{setFreq: 94, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["45"] = &isdbtChannel{setFreq: 95, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["46"] = &isdbtChannel{setFreq: 96, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["47"] = &isdbtChannel{setFreq: 97, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["48"] = &isdbtChannel{setFreq: 98, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["49"] = &isdbtChannel{setFreq: 99, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["50"] = &isdbtChannel{setFreq: 100, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["51"] = &isdbtChannel{setFreq: 101, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["52"] = &isdbtChannel{setFreq: 102, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["53"] = &isdbtChannel{setFreq: 103, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["54"] = &isdbtChannel{setFreq: 104, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["55"] = &isdbtChannel{setFreq: 105, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["56"] = &isdbtChannel{setFreq: 106, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["57"] = &isdbtChannel{setFreq: 107, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["58"] = &isdbtChannel{setFreq: 108, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["59"] = &isdbtChannel{setFreq: 109, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["60"] = &isdbtChannel{setFreq: 110, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["61"] = &isdbtChannel{setFreq: 111, channelType: ChannelTypeGround, addFreq: 0}
	isdbtConvertTable["62"] = &isdbtChannel{setFreq: 112, channelType: ChannelTypeGround, addFreq: 0}
}
