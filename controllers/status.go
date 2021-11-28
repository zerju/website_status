package controllers

import (
	"net/http"
	"sort"
	"time"

	"github.com/gin-gonic/gin"

	"website_status/models"
)

func RenderStatusPage(c *gin.Context) {
	stc := make(chan models.Site)
	var sites []models.Site = models.SiteList

	// check all websites
	for _, site := range sites {
		go checkLink(site, stc)
	}

	// wait for all websites to get the data back
	var updatedSites []models.Site
	for i := 0; i < len(sites); i++ {
		updatedSites = append(updatedSites, <-stc)
	}

	sortSitesById(&updatedSites)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Websites Statuses",
		"sites": updatedSites,
	})
}

func sortSitesById(sites *[]models.Site) {
	// sort by ID, ASC
	sort.Slice((*sites), func(p, q int) bool {
		return (*sites)[p].ID < (*sites)[q].ID
	})
}

func checkLink(s models.Site, stc chan models.Site) {
	_, err := http.Get(s.Link)
	s.LastChecked = time.Now().Format(time.UnixDate)
	if err != nil {
		s.IsUp = false
		stc <- s
		return
	}
	s.IsUp = true
	stc <- s
}
