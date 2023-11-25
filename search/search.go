package search

import (
	// "unicode"
	"fmt"
	"strings"
	"strconv"

	"github.com/ashutsh/stringSearch/utils"
)


func CharacterPresence(substr string, dataSlc []string) []string {
	rawresult := CharacterPresenceRAW(substr, dataSlc)
	result := make([]string, 0)
	for _, r := range rawresult {
		i, _ := strconv.Atoi(strings.Split(r, ":")[0])
		result = append(result, dataSlc[i])
	}

	return result
}

func CharacterPresenceRAW(substr string, dataSlc []string) []string {
	result := make([]string, 0)
	dataSlcSub := make([]string, 0)
	dataSlcSub = append(dataSlcSub, dataSlc...)
	for dsi, dsr := range dataSlcSub {
		resultStr := ""
		dsr = strings.ToLower(dsr)

		for ri, rr := range dsr {
			for si, sr := range substr {
				if rr == sr {
					if resultStr == "" {
						resultStr = fmt.Sprintf("%d: ",dsi)
					}
					resultStr += fmt.Sprintf("[%c|%d]_%d, ",sr,si,ri)
				}
			}
		}
		if resultStr != "" {
			extract := ExctractBetweenDeliminators('|', ']', resultStr)
			isPresent, isConsequtive := IsAllPresent(extract, substr)
			if isPresent && isConsequtive {
				// resultStr += fmt.Sprint(isConsequtive)
				result = append(result, resultStr)
			}
		}
	}

	return result
}


func ExctractBetweenDeliminators(del1, del2 rune, str string) []string {
	result := make([]string, 0)
	resultStr := ""
	isdel := false
	for _, r := range str {
		if r == del1 {
			isdel = true
		} else if r == del2 {
			isdel = false
			result = append(result, resultStr)
			resultStr = ""
		}
		if r != del1 && isdel {
			resultStr += string(r) 
		}
	}

	return result
}

func IsAllPresent(dataStrSlc []string, refStr string) (bool, bool) {
	dataIntSlc := utils.ConvertStringToIntSlice(dataStrSlc)
	uniq := utils.UniqueIntSlc(dataIntSlc)
	
	return len(uniq) == len(refStr), IsConsequtive(dataIntSlc, uniq)
}

func IsConsequtive(data, uniq []int) bool {
	// uniq := utils.UniqueIntSlc(data)
	isStart := false
	tracker := 0
	for _, r := range data {
		if r == 0 {
			isStart = true
		}
		if isStart && r == tracker + 1  {
			tracker++
		}
	}
	return tracker +1 == len(uniq)
}