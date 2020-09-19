package cod

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/avast/retry-go"
)

// GetMoreWarzoneMatches returns 140 matches
// TODO: use this somewhere or delete u frickin pepegas
func GetMoreWarzoneMatches(username string) ([]MatchData, error) {
	var allMatches []MatchData

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
			var matchData MatchData
			err = json.Unmarshal(body, &matchData)
			if err != nil {
				log.Fatal(err)
			}
			allMatches = append(allMatches, matchData)
		}
	}
	return allMatches, nil
}

// GetMatchData will return the data about a match for a given user
func GetMatchData(username string, client *http.Client) (MatchData, error) {
	var matchData MatchData
	var Code int
	// url
	rawURL := "https://api.tracker.gg/api/v1/warzone/matches/atvi/%s?type=wz&next=null"
	// api url
	url, err := url.Parse(fmt.Sprintf(rawURL, username))
	// invalid url
	if err != nil {
		fmt.Println("Unable to parse url")
	}
	// build http request
	retryErr := retry.Do(
		func() error {
			req, err := http.NewRequest(
				http.MethodGet,
				url.String(),
				nil,
			)
			if err != nil {
				return err
			}
			// set a normal/non-hackerman user agent
			req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36")
			// just to see if their server is checking remote ip ;)
			req.Header.Set("X-Remote-IP", "127.0.0.1")
			// default client
			resp, err := client.Do(req)
			if err != nil {
				return err
			}
			// read response status code
			s := resp.StatusCode
			// resp - 500
			if s >= 500 {
				// assign current status code
				Code = s
				// Fully consume the body, which will also lead to us reading
				// the trailer headers after the body, if present.
				io.Copy(ioutil.Discard, resp.Body)
				// close
				resp.Body.Close()
				// return custom error
				err := fmt.Errorf("Respone code: %d", s)
				return err
				// resp - 200 OK
			} else if s == http.StatusOK {
				// assign current status code
				Code = s
				// read body
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					return err
				}
				// unmarshal json into match data struct
				err = json.Unmarshal(body, &matchData)
				if err != nil {
					return err
				}
				// Fully consume the body, which will also lead to us reading
				// the trailer headers after the body, if present.
				io.Copy(ioutil.Discard, resp.Body)
				// fully close
				resp.Body.Close()
				// no error
				return nil
				// resp - 404 - NOT FOUND
			} else if s == 404 {
				// assign current status code
				Code = s
				// Fully consume the body, which will also lead to us reading
				// the trailer headers after the body, if present.
				io.Copy(ioutil.Discard, resp.Body)
				// close
				resp.Body.Close()
				// return custom error
				err := fmt.Errorf("NOT FOUND Respone code: %d", s)
				return err
			} else {
				Code = s
				err := fmt.Errorf("This was not handled: %d", s)
				return err
			}
		},
	)
	// hope fully never gets called
	if retryErr != nil {
		fmt.Println(retryErr)
	}
	return matchData, fmt.Errorf("GetMatchData: status code %d: %s", Code, username)
}

// GetWarzoneStats retrieves the stats of an individual player in warzone
func GetWarzoneStats(username string) (StatData, error) {
	var statData StatData
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

// IsValid checks if a user with the username exists
func IsValid(user string) bool {
	resp, err := http.Get(fmt.Sprintf("https://api.tracker.gg/api/v2/warzone/standard/profile/atvi/%s", user))
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode == 200 {
		return true
	}
	return false
}
