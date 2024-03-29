package tournament

import (
	"encoding/csv"
	"io"
	"log"
	"strings"
	"zleague/api/models"
)

// CreateTeams a map of teams and players
func CreateTeams(csvData io.Reader) map[string]models.TeamBasic {
	// read file line by line
	lines, err := csv.NewReader(csvData).ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	// team map
	teamMap := map[string]models.TeamBasic{}
	// for each line in csv grab the division player id and teamname
	for _, line := range lines {
		div := line[0]
		team := line[1]
		player := strings.Replace(line[2], "#", "%23", -1)

		// teamname is present in map
		if _, ok := teamMap[team]; ok {
			t := teamMap[team]                        // reference map
			t.Teammates = append(t.Teammates, player) // append player to appropriate team
			teamMap[team] = t
		} else { // not present in map
			p := models.TeamBasic{Division: div, Teamname: team} // create new struct reference
			p.Teammates = append(p.Teammates, player)                                    // append teammate
			teamMap[team] = p                                                            // create key value pair
		}
	}

	return teamMap
}
