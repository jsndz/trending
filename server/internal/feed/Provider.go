package feed

import "github.com/jsndz/trending/internal/model"

type Provider interface {
	GetSource() model.Source
	Parser() (*[]model.Article, error)
}
