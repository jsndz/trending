package techcrunch

import (
	"encoding/xml"
	"time"

	"github.com/jsndz/trending/internal/feed"
	"github.com/jsndz/trending/internal/model"
	"github.com/jsndz/trending/pkg/util"
)

type TechCrunch struct {
	URL string
}

func NewTechCrunch() *TechCrunch {
	return &TechCrunch{
		URL: "https://techcrunch.com/feed/",
	}
}

func (t *TechCrunch) GetSource() model.Source {
	return model.SourceTechCrunch
}

func (t *TechCrunch) Parser() (*[]model.Article, error) {
	data, err := feed.Fetch(t.URL)
	if err != nil {
		return nil, err
	}
	var rss RSS
	err = xml.Unmarshal(data, &rss)
	if err != nil {
		return nil, err
	}
	var articles []model.Article
	for _, item := range rss.Channel.Items {
		pubTime, _ := time.Parse(time.RFC1123Z, item.PublishedAt)
		url, _ := util.NormalizeURL(item.Link)
		articles = append(articles, model.Article{
			Title:       item.Title,
			PublishedAt: pubTime,
			Link:        url,
			Author:      item.Author,
			Description: item.Description,
			Source:      model.SourceTechCrunch,
		})
	}
	return &articles, nil
}
