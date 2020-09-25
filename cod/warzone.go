package cod

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/avast/retry-go"
)

// GetMoreWarzoneMatches returns 140 matches
// TODO: use this somewhere or delete u frickin pepegas
func GetMoreWarzoneMatches(username string) ([]MatchData, error) {
	var matches []MatchData

	// pagination ids
	reqs := []string{"null", "1598805852999", "1598214721999", "1598113460999", "1597612561999", "1597528982999", "1596998692999"}
	// base uri
	uri := "https://api.tracker.gg/api/v1/warzone/matches/atvi/%s?type=wz&next=%s"

	for _, val := range reqs {
		endpoint := fmt.Sprintf(uri, username, val)

		resp, err := http.Get(endpoint)
		if err != nil {
			return []MatchData{}, err
		}

		defer resp.Body.Close()

		if resp.StatusCode == 200 {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return []MatchData{}, err
			}

			var matchData MatchData
			err = json.Unmarshal(body, &matchData)
			if err != nil {
				return []MatchData{}, err
			}

			matches = append(matches, matchData)
		}
	}
	return matches, nil
}

// GetMatchData will return the data about a match for a given user
func GetMatchData(username string, client *http.Client) (MatchData, error) {
	var matchData MatchData
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
			// 500
			if s >= http.StatusInternalServerError {
				// Fully consume the body, which will also lead to us reading
				// the trailer headers after the body, if present.
				_, err = io.Copy(ioutil.Discard, resp.Body)
				if err != nil {
					return err
				}
				// close
				resp.Body.Close()
				// return custom error
				err := fmt.Errorf("Respone code: %d", s)
				return err
				// 200
			} else if s == http.StatusOK {
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
				_, err = io.Copy(ioutil.Discard, resp.Body)
				if err != nil {
					return err
				}
				// fully close
				resp.Body.Close()
				// no error
				return nil
				// 404
			} else if s == http.StatusNotFound {
				// Fully consume the body, which will also lead to us reading
				// the trailer headers after the body, if present.
				_, err = io.Copy(ioutil.Discard, resp.Body)
				if err != nil {
					return err
				}
				// close
				resp.Body.Close()
				// return custom error
				err := fmt.Errorf("NOT FOUND Respone code: %d", s)
				return err
			} else {
				err := fmt.Errorf("This was not handled: %d", s)
				return err
			}
		},
	)

	// this should never be called if everything goes well
	if retryErr != nil {
		return MatchData{}, retryErr
	}

	return matchData, nil
}

// GetStatData retrieves the stats of an individual player in warzone
func GetStatData(username string, client *http.Client) (StatData, error) {
	var statData StatData

	rawURL := "https://api.tracker.gg/api/v2/warzone/standard/profile/atvi/%s"
	url, err := url.Parse(fmt.Sprintf(rawURL, username))
	// invalid url (probably due to username)
	if err != nil {
		return StatData{}, err
	}
	// build http request if there is an error we issue a retry
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
			// 500
			if s >= http.StatusInternalServerError {
				// Fully consume the body, which will also lead to us reading
				// the trailer headers after the body, if present.
				_, err = io.Copy(ioutil.Discard, resp.Body)
				if err != nil {
					return err
				}
				// close
				resp.Body.Close()
				// return custom error
				err := fmt.Errorf("Respone code: %d", s)
				return err
				// resp - 200 OK
			} else if s == http.StatusOK {
				// read body
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					return err
				}
				// unmarshal json into stat data struct
				err = json.Unmarshal(body, &statData)
				if err != nil {
					return err
				}
				// Fully consume the body, which will also lead to us reading
				// the trailer headers after the body, if present.
				_, err = io.Copy(ioutil.Discard, resp.Body)
				if err != nil {
					return err
				}

				// manually close body
				resp.Body.Close()

				return nil
				// 404
			} else if s == http.StatusNotFound {
				// Fully consume the body, which will also lead to us reading
				// the trailer headers after the body, if present.
				_, err = io.Copy(ioutil.Discard, resp.Body)
				if err != nil {
					return err
				}

				// manually close body
				resp.Body.Close()

				return fmt.Errorf("NOT FOUND Respone code: %d", s)
			} else {
				return fmt.Errorf("This was not handled: %d", s)
			}
		},
	)

	// this should never be called
	if retryErr != nil {
		return StatData{}, retryErr
	}

	return statData, nil
}

// IsValid checks if a user with the username exists
func IsValid(user string) bool {
	url := strings.TrimSpace(fmt.Sprintf("https://api.tracker.gg/api/v2/warzone/standard/profile/atvi/%s", user))

	resp, err := http.Get(url)
	if err != nil {
		return false
	}

	if resp.StatusCode != http.StatusOK {
		return false
	}

	return true
}
