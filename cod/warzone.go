package cod

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/zleague/backend/utils"
)

// GetMoreWarzoneMatches does
func GetMoreWarzoneMatches(username string) (utils.MatchData, error) {
	var matchData utils.MatchData
	var allMatches []utils.MatchData

	reqs := []string{"null", "1598805852999", "1598214721999", "1598113460999", "1597612561999", "1597528982999", "1596998692999"}
	// base uri
	uri := "https://api.tracker.gg/api/v1/warzone/matches/atvi/%s?type=wz&next=%s"

	for _, val := range reqs {
		endpoint := fmt.Sprintf(uri, username, val)
		resp, err := http.Get(endpoint)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		if resp.StatusCode == 200 {
			body, err := ioutil.ReadAll(resp.Body)

			err = json.Unmarshal(body, &matchData)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	return allMatches, nil
}

// GetWarzoneMatches retrieves a list of all the players previous warzone matches
func GetWarzoneMatches(username string) (utils.MatchData, error) {
	var matchData utils.MatchData

	resp, err := http.Get(fmt.Sprintf("https://api.tracker.gg/api/v1/warzone/matches/atvi/%s?type=wz&next=null", username))
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		body, err := ioutil.ReadAll(resp.Body)

		err = json.Unmarshal(body, &matchData)
		if err != nil {
			log.Fatal(err)
		}
		return matchData, nil
	}

	return matchData, fmt.Errorf("GetWarzoneMatches: status code %d: %s", resp.StatusCode, username)
}

// GetWarzoneStats retrieves the stats of an individual player in warzone
func GetWarzoneStats(username string) (utils.StatData, error) {
	var statData utils.StatData
	resp, err := http.Get(fmt.Sprintf("https://api.tracker.gg/api/v2/warzone/standard/profile/atvi/%s", username))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		body, err := ioutil.ReadAll(resp.Body)

		err = json.Unmarshal(body, &statData)
		if err != nil {
			log.Fatal(err)
		}
		return statData, nil
	}
	return statData, fmt.Errorf("GetWarzoneStats: status code %d: %s", resp.StatusCode, username)
}

// CheckUser checks if a user with the username exists
func CheckUser(user string) bool {
	resp, err := http.Get(fmt.Sprintf("https://api.tracker.gg/api/v2/warzone/standard/profile/atvi/%s", user))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.StatusCode)
	if resp.StatusCode == 200 {
		return true
	}
	return false
}
