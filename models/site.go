package models

type Site struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Link        string `json:"link"`
	LastChecked string `json:"lastChecked"`
	IsUp        bool   `json:"isUp"`
}

var SiteList = []Site{
	{ID: 1, Name: "Google", Link: "https://google.com", IsUp: false},
	{ID: 2, Name: "Amazon", Link: "https://amazon.com", IsUp: false},
	{ID: 3, Name: "Stack Overflow", Link: "https://stackoverflow.com", IsUp: false},
	{ID: 4, Name: "Facebook", Link: "https://facebook.com", IsUp: false},
	{ID: 5, Name: "Instagram", Link: "https://instagram.com", IsUp: false},
	{ID: 6, Name: "Kitsu", Link: "https://kitsu.io", IsUp: false},
	{ID: 7, Name: "Youtube", Link: "https://youtube.com", IsUp: false},
	{ID: 8, Name: "Reddit", Link: "https://reddit.com", IsUp: false},
	{ID: 9, Name: "Coingecko", Link: "https://coingecko.com", IsUp: false},
	{ID: 10, Name: "Linkedin", Link: "https://linkedin.com", IsUp: false},
}
