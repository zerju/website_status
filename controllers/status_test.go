package controllers

import (
	"math/rand"
	"net/http"
	"reflect"
	"testing"
	"time"
	"website_status/models"
)

var testSites = []models.Site{
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

func TestCheckLink(t *testing.T) {
	var sites []models.Site = testSites
	var firstSite models.Site = sites[0]
	var stc = make(chan models.Site)

	resp, err := http.Get(firstSite.Link)
	go checkLink(firstSite, stc)

	var isUp = resp.StatusCode < 400 || err != nil
	var checkedIsUp = (<-stc).IsUp

	if checkedIsUp != isUp {
		t.Errorf("Expected the checked link isUp value to be %v, but instead got %v", isUp, checkedIsUp)
	}
}

func TestSortSitesById(t *testing.T) {
	var sites []models.Site = testSites
	// one slice should be left sorted for later comparison
	var sortedSites []models.Site = testSites

	// shuffle slice randomly (it's already ordered otherwise)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(sites), func(i, j int) {
		sites[i], sites[j] = sites[j], sites[i]
	})

	sortSitesById(&sites)

	var bothSorted = reflect.DeepEqual(sortedSites, sites)
	if !bothSorted {
		t.Errorf("The already sorted and newly sorted sites are not equal, meaning that one of them is not sorted correctly")
	}
}
