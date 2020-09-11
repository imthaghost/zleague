package tournament

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

// Create a map of teams and players
func Create(start, end time.Time) map[string]TeamBasic {
	// get current working directory
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	// file path
	filepath := path + "/test.csv"
	inputfile, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer inputfile.Close()
	// read file line by line
	lines, err := csv.NewReader(inputfile).ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// team map
	teamMap := map[string]TeamBasic{}
	// for each line in csv grab the division player id and teamname
	for _, line := range lines {
		div := line[0]
		team := line[1]
		player := strings.Replace(line[2], "#", "%23", -1)

		// teamname is present in map
		if _, ok := teamMap[team]; ok {
			t := teamMap[team]                        // reference map
			t.Teammates = append(t.Teammates, player) // append player to appopriate team
			teamMap[team] = t
		} else { // not present in map
			p := TeamBasic{Division: div, Teamname: team, Start: start, End: end} // create new struct reference
			p.Teammates = append(p.Teammates, player)                             // append teammate
			teamMap[team] = p                                                     // create key value pair
		}
	}
	fmt.Println("created map")
	return teamMap
}
