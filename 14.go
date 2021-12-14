package main

import "fmt"

func pairToIndex(s string) int {
	firstChar := s[0]
	secondChar := s[1]
	return int(firstChar-'A')*26 + int(secondChar-'A')
}

func main() {
	inputRows := []string{
		"KKOSPHCNOCHHHSPOBKVF",
		"",
		"NV -> S",
		"OK -> K",
		"SO -> N",
		"FN -> F",
		"NB -> K",
		"BV -> K",
		"PN -> V",
		"KC -> C",
		"HF -> N",
		"CK -> S",
		"VP -> H",
		"SK -> C",
		"NO -> F",
		"PB -> O",
		"PF -> P",
		"VC -> C",
		"OB -> S",
		"VF -> F",
		"BP -> P",
		"HO -> O",
		"FF -> S",
		"NF -> B",
		"KK -> C",
		"OC -> P",
		"OV -> B",
		"NK -> B",
		"KO -> C",
		"OH -> F",
		"CV -> F",
		"CH -> K",
		"SC -> O",
		"BN -> B",
		"HS -> O",
		"VK -> V",
		"PV -> S",
		"BO -> F",
		"OO -> S",
		"KB -> N",
		"NS -> S",
		"BF -> N",
		"SH -> F",
		"SB -> S",
		"PP -> F",
		"KN -> H",
		"BB -> C",
		"SS -> V",
		"HP -> O",
		"PK -> P",
		"HK -> O",
		"FH -> O",
		"BC -> N",
		"FK -> K",
		"HN -> P",
		"CC -> V",
		"FO -> F",
		"FP -> C",
		"VO -> N",
		"SF -> B",
		"HC -> O",
		"NN -> K",
		"FC -> C",
		"CS -> O",
		"FV -> P",
		"HV -> V",
		"PO -> H",
		"BH -> F",
		"OF -> P",
		"PC -> V",
		"CN -> O",
		"HB -> N",
		"CF -> P",
		"HH -> K",
		"VH -> H",
		"OP -> F",
		"BK -> S",
		"SP -> V",
		"BS -> V",
		"VB -> C",
		"NH -> H",
		"SN -> K",
		"KH -> F",
		"OS -> N",
		"NP -> P",
		"VN -> V",
		"KV -> F",
		"KP -> B",
		"VS -> F",
		"NC -> F",
		"ON -> S",
		"FB -> C",
		"SV -> O",
		"PS -> K",
		"KF -> H",
		"CP -> H",
		"FS -> V",
		"VV -> H",
		"CB -> P",
		"PH -> N",
		"CO -> N",
		"KS -> K",
	}

	connectionsOutputs := make([]byte, 26*26)
	connectionsCounts := make([]int, 26*26)
	elementCounts := make([]int, 26)

	for i := 2; i < len(inputRows); i++ {
		row := inputRows[i]
		pair := row[:2]
		resultElement := row[len(row)-1]
		connectionsOutputs[pairToIndex(pair)] = resultElement
	}

	polymer := inputRows[0]
	for _, char := range polymer {
		elementCounts[int(char-'A')] += 1
	}

	for i := 0; i < len(polymer)-1; i++ {
		pair := polymer[i : i+2]
		connectionsCounts[pairToIndex(pair)] += 1
	}

	for step := 1; step <= 40; step++ {
		newConnectionsCounts := make([]int, len(connectionsCounts))
		for connectionIndex, count := range connectionsCounts {
			if count == 0 {
				continue
			}
			firstChar := byte(connectionIndex / 26) + 'A'
			secondChar := byte(connectionIndex % 26) + 'A'
			charToInsert := connectionsOutputs[connectionIndex]

			elementCounts[charToInsert - 'A'] += count

			newLeftPair := string(firstChar) + string(charToInsert)
			newRightPair := string(charToInsert) + string(secondChar)
			newConnectionsCounts[pairToIndex(newLeftPair)] += count
			newConnectionsCounts[pairToIndex(newRightPair)] += count
		}
		connectionsCounts = newConnectionsCounts
	}

	max := -1
	min := -1
	for _, count := range elementCounts {
		if count > 0 && (max < 0 || count > max) {
			max = count
		}
		if count > 0 && (min < 0 || count < min) {
			min = count
		}
	}

	fmt.Println(max, "-", min, "=", max-min)
}
