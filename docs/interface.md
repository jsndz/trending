Interface:

Interface in general means a place where 2 system meet.
Two systems are consumer and inteface
Like user meets the computer using cli as interface



Consider here for example, 
Computer and microphones are system
If you want these to meet you need something that convert the audio input to digital input.
For this to work the audio device need to be of the exact type of the interface 
Meaning the audio device should agree to the contract

In terms of programming with the above example,
The db should conform with the contract of interface to communicate with the app

Like in golang to conform to the interface they should implement the function with same params

Application wont care about implementation of methods it only cares that it can implement it

Example in code:


This is the interface:
type Provider interface {
   GetSource() model.Source
   Parser() (*[]model.Article, error)
}



This is one of the implementations:

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



For the consumer side the consumer struct needs to have the interface(contract)


Here feed service will be using any one of the sources

type FeedService struct {
   ArticlesRepo *repository.ArticlesRepository
   Providers    feed.Provider
}
func NewFeedService(articleRepo *repository.ArticlesRepository, provider feed.Provider) *FeedService {
   return &FeedService{
       ArticlesRepo: articleRepo,
       Provider:     provider,
   }
}





You can provide any of the provider struct that implements Provider interface


Here you can clearly see the FeedService does not care about the how of provider 
Fully abstracted