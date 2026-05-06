package techcrunch

import (
	"encoding/xml"
	"time"

	"github.com/jsndz/trending/internal/model"
	"github.com/jsndz/trending/pkg/util"
)

func Parser(data []byte) (*[]model.Article, error) {
	var rss RSS
	err := xml.Unmarshal(data, &rss)
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
