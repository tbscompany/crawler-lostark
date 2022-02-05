package runner

import (
	"bytes"
	"crawler-lostark/pkg/config"
	"crawler-lostark/pkg/database"
	"crawler-lostark/pkg/logger"
	"crawler-lostark/pkg/models"
	"crawler-lostark/pkg/service"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gocolly/colly"
	"github.com/urfave/cli/v2"
)

type newsService struct {
	service service.DefaultNewsService
}

var allNews []*models.News

type Message struct {
	Content string `json:"content"`
}

func StartRunner(c *cli.Context) error {
	// Init logger
	logger.InitLogger(c.App.Version)

	// Init database
	database.InitDatabase()

	s := newsService{service.NewNewsService(models.NewNewsRepository())}

	for {
		scrape()

		for _, news := range allNews {
			n, err := s.service.GetNewsByTitle(news.Title)
			if err.Error != nil {
				logger.Logger.Error().Err(err.Error).Msg("Unable to get news by title")
			}
			if n.Title == "" {
				s.service.CreateNews(news)

				buf := new(bytes.Buffer)
				err := json.NewEncoder(buf).Encode(Message{Content: news.URL})
				if err != nil {
					logger.Logger.Error().Err(err).Msg("Unable to encode message")
				} else {
					resp, err := http.Post(config.DiscordWebhook, "application/json", buf)
					if err != nil {
						logger.Logger.Error().Err(err).Msg("Unable to send message")
					}
					if config.Debug {
						j, _ := json.Marshal(resp)
						fmt.Println(string(j))
					}
				}
			}
		}

		time.Sleep(time.Minute * 5)
	}
}

func scrape() {
	c := colly.NewCollector(
		colly.AllowedDomains("www.playlostark.com"),
	)

	c.OnHTML("div.ags-SlotModule--threePerRow", func(e *colly.HTMLElement) {
		news := new(models.News)
		news.URL = "https://www.playlostark.com" + e.ChildAttr("a.ags-SlotModule-spacer", "href")
		news.Date = e.ChildText("span.ags-SlotModule-contentContainer-date")
		news.Title = e.ChildText("span.ags-SlotModule-contentContainer-heading--blog")

		allNews = append(allNews, news)
	})

	c.Visit("https://www.playlostark.com/fr-fr/news")
}
