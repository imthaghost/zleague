package utils

// MatchError stores the JSON response of an error
type MatchError struct {
	Errors []struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"errors"`
}
