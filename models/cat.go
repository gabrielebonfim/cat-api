package models

type Cat struct {
	ID     string  `json:"id"`
	Width  int     `json:"width"`
	Height int     `json:"height"`
	Url    string  `json:"url"`
	Breeds []Breed `json:"breeds"`
}

type Breed struct {
	Weight       Weight `json:"weight"`
	Id           string `json:"id"`
	Name         string `json:"name"`
	Temperament  string `json:"temperament"`
	Origin       string `json:"origin"`
	CountryCodes string `json:"countryCodes"`
	LifeSpan     string `json:"lifeSpan"`
	WikipediaUrl string `json:"wikipediaUrl"`
}

type Weight struct {
	Imperial string `json:"imperial"`
	Metric   string `json:"metric"`
}
