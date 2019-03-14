package model

import "encoding/json"

// Profile model
type Profile struct {
	URL string
	ID  string

	Name       string
	Gender     string
	Age        int
	Height     int
	Weight     int
	Income     string
	Marriage   string
	Education  string
	Occupation string
	Hokou      string
	Xinzuo     string
	House      string
	Car        string
}

// FromJSONObj form profile model
func FromJSONObj(o interface{}) (Profile, error) {
	var profile Profile
	s, err := json.Marshal(o)
	if err != nil {
		return profile, err
	}
	err = json.Unmarshal(s, &profile)
	return profile, err
}
