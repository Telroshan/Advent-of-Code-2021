package main

import "fmt"

var data []uint8
var bitOffset int = 0
var totalVersionNumbers int = 0

func getSingleBitValue() int {
	bitGroup := data[bitOffset/8]
	indexInGroup := bitOffset % 8
	isolatedBit := bitGroup & (1 << (7 - indexInGroup))
	bitOffset++
	return int(isolatedBit >> (7 - indexInGroup))
}

func getMultipleBitsValue(length int) int {
	value := 0
	for i := 0; i < length; i++ {
		value <<= 1
		value |= getSingleBitValue()
	}
	return value
}

func getPacket() int {
	version := getMultipleBitsValue(3)
	totalVersionNumbers += version
	typeID := getMultipleBitsValue(3)
	result := 0
	if typeID == 4 {
		for {
			firstBit := getSingleBitValue()
			result <<= 4
			result += getMultipleBitsValue(4)
			if firstBit == 0 {
				return result
			}
		}
	} else {
		lengthType := getSingleBitValue()
		var innerValues []int
		if lengthType == 0 {
			subPacketsLength := getMultipleBitsValue(15)
			start := bitOffset
			for bitOffset < start+subPacketsLength {
				innerValues = append(innerValues, getPacket())
			}
		} else {
			subPacketsCount := getMultipleBitsValue(11)
			innerValues = make([]int, subPacketsCount)
			for i := 0; i < subPacketsCount; i++ {
				innerValues[i] = getPacket()
			}
		}
		switch typeID {
		case 0:
			for _, value := range innerValues {
				result += value
			}
		case 1:
			result = innerValues[0]
			if len(innerValues) > 1 {
				for i := 1; i < len(innerValues); i++ {
					result *= innerValues[i]
				}
			}
		case 2:
			result = innerValues[0]
			for i := 1; i < len(innerValues); i++ {
				if innerValues[i] < result {
					result = innerValues[i]
				}
			}
		case 3:
			result = innerValues[0]
			for i := 1; i < len(innerValues); i++ {
				if innerValues[i] > result {
					result = innerValues[i]
				}
			}
		case 5:
			if innerValues[0] > innerValues[1] {
				result = 1
			} else {
				result = 0
			}
		case 6:
			if innerValues[0] < innerValues[1] {
				result = 1
			} else {
				result = 0
			}
		case 7:
			if innerValues[0] == innerValues[1] {
				result = 1
			} else {
				result = 0
			}
		}
	}
	return result
}

func main() {
	input := "620D79802F60098803B10E20C3C1007A2EC4C84136F0600BCB8AD0066E200CC7D89D0C4401F87104E094FEA82B0726613C6B692400E14A305802D112239802125FB69FF0015095B9D4ADCEE5B6782005301762200628012E006B80162007B01060A0051801E200528014002A118016802003801E2006100460400C1A001AB3DED1A00063D0E25771189394253A6B2671908020394359B6799529E69600A6A6EB5C2D4C4D764F7F8263805531AA5FE8D3AE33BEC6AB148968D7BFEF2FBD204CA3980250A3C01591EF94E5FF6A2698027A0094599AA471F299EA4FBC9E47277149C35C88E4E3B30043B315B675B6B9FBCCEC0017991D690A5A412E011CA8BC08979FD665298B6445402F97089792D48CF589E00A56FFFDA3EF12CBD24FA200C9002190AE3AC293007A0A41784A600C42485F0E6089805D0CE517E3C493DC900180213D1C5F1988D6802D346F33C840A0804CB9FE1CE006E6000844528570A40010E86B09A32200107321A20164F66BAB5244929AD0FCBC65AF3B4893C9D7C46401A64BA4E00437232D6774D6DEA51CE4DA88041DF0042467DCD28B133BE73C733D8CD703EE005CADF7D15200F32C0129EC4E7EB4605D28A52F2C762BEA010C8B94239AAF3C5523CB271802F3CB12EAC0002FC6B8F2600ACBD15780337939531EAD32B5272A63D5A657880353B005A73744F97D3F4AE277A7DA8803C4989DDBA802459D82BCF7E5CC5ED6242013427A167FC00D500010F8F119A1A8803F0C62DC7D200CAA7E1BC40C7401794C766BB3C58A00845691ADEF875894400C0CFA7CD86CF8F98027600ACA12495BF6FFEF20691ADE96692013E27A3DE197802E00085C6E8F30600010882B18A25880352D6D5712AE97E194E4F71D279803000084C688A71F440188FB0FA2A8803D0AE31C1D200DE25F3AAC7F1BA35802B3BE6D9DF369802F1CB401393F2249F918800829A1B40088A54F25330B134950E0"

	data = make([]uint8, len(input) / 2)
	for i, char := range input {
		var value uint8
		if char >= '0' && char <= '9' {
			value = uint8(char - '0')
		} else {
			value = 10 + uint8(char - 'A')
		}
		if i%2 == 0 {
			data[i/2] = value << 4
		} else {
			data[i/2] |= value
		}
	}

	result := getPacket()

	fmt.Println("Total version numbers :", totalVersionNumbers)
	fmt.Println("Result :",result)
}