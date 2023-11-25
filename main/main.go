package main

import (
	"fmt"
	// "strconv"
	// "strings"

	"github.com/ashutsh/stringSearch/data"
	"github.com/ashutsh/stringSearch/search"
	// "github.com/ashutsh/stringSearch/utils"
)

func main() {
	result := search.CharacterPresence("cig", data.StringSlice)

	// maxlen := 0
	// for _, r := range data.StringSlice {
	// 	if len(r) > maxlen {
	// 		maxlen = len(r)
	// 	}
	// }


	// // Mainly used for debugging while adding features.
	// rawresult := search.CharacterPresenceRAW("word", data.StringSlice)		
	// for i, r := range rawresult {
	// 	ii, _ := strconv.Atoi(strings.Split(r, ":")[0])
	// 	if i > 9 { 
	// 		fmt.Printf("%d. %s %s =>  %s      %s \n", i, data.StringSlice[ii], strings.Repeat(" ", maxlen-len(data.StringSlice[ii])), r, search.ExctractBetweenDeliminators('|', ']', r))
	// 	} else {
	// 		fmt.Printf("%d.  %s %s =>  %s      %s \n", i, data.StringSlice[ii], strings.Repeat(" ", maxlen-len(data.StringSlice[ii])), r, search.ExctractBetweenDeliminators('|', ']', r))
	// 	}
	// }

	// // Just iterating over the complete data.StringSlice to know the data
	// for i, r := range data.StringSlice {
	// 	if i+1 > 9 {
	// 		fmt.Printf("%d. %s\n", i+1, r)
	// 	} else {
	// 		fmt.Printf("%d.  %s\n", i+1, r)
	// 	}
	// }

	
	for i, r := range result {
		if i+1 > 9 {
			fmt.Printf("%d. %s\n", i+1, r)
		} else {
			fmt.Printf("%d.  %s\n", i+1, r)
		}
	}
	// for _, r := range result {
	// 	intExtract := utils.ConvertStringToIntSlice(search.ExctractBetweenDeliminators('|', ']', r))
	// 	fmt.Println(intExtract, "=>", utils.UniqueIntSlc(intExtract))
	// }
}