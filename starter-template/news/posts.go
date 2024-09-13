package news

import (
	"fmt"
	"github.com/maddalax/mhtml/framework/util/httpjson"
	"sync"
)

type Post struct {
	By          string `json:"by"`
	Descendants int    `json:"descendants"`
	Id          int    `json:"id"`
	Kids        []int  `json:"kids"`
	Score       int    `json:"score"`
	Time        int    `json:"time"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	Url         string `json:"url"`
}

func List() ([]Post, error) {
	responseIds, err := httpjson.Get[[]int64]("https://hacker-news.firebaseio.com/v0/topstories.json")
	responseIds = responseIds[0:50]
	if err != nil {
		return []Post{}, err
	}

	var wg sync.WaitGroup
	posts := make([]Post, len(responseIds))

	for index, id := range responseIds {
		wg.Add(1)
		id := id
		index := index
		go func() {
			defer wg.Done()
			url := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json", id)
			post, err := httpjson.Get[Post](url)
			if err != nil {
				println(err.Error())
			}
			posts[index] = post
		}()
	}

	wg.Wait()

	return posts, nil
}

func Get(id string) (Post, error) {
	url := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%s.json", id)
	post, err := httpjson.Get[Post](url)
	if err != nil {
		return Post{}, err
	}
	return post, nil
}
