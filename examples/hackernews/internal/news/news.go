package news

import (
	"fmt"
	"github.com/maddalax/htmgo/framework/h"
	"hackernews/internal/batch"
	"hackernews/internal/httpjson"
	"hackernews/internal/timeformat"
	"log/slog"
	"strconv"
	"time"
)

const baseUrl = "https://hacker-news.firebaseio.com/v0/"

func url(path string, qs *h.Qs) string {
	return baseUrl + path + ".json?" + qs.ToString()
}

type Category struct {
	Name string
	Path string
}

var Categories = []Category{
	{"Top Stories", "topstories"},
	{"Best Stories", "beststories"},
	{"New Stories", "newstories"},
}

type Comment struct {
	By      string    `json:"by"`
	Text    string    `json:"text"`
	TimeRaw int64     `json:"time"`
	Time    time.Time `json:"-"`
	Type    string    `json:"type"`
	Kids    []int     `json:"kids"`
	Parent  int       `json:"parent"`
	Id      int       `json:"id"`
}

type Story struct {
	Id          int    `json:"id"`
	By          string `json:"by"`
	Text        string `json:"text"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	Descendents int    `json:"descendants"`
	Score       int    `json:"score"`
	Url         string
	TimeRaw     int64     `json:"time"`
	Time        time.Time `json:"-"`
	// comment ids
	Kids []int
}

type GetTopStoriesRequest struct {
	Limit int
	Page  int
}

func MustItemId(ctx *h.RequestContext) int {
	raw := h.GetQueryParam(ctx, "item")
	parsed, err := strconv.ParseInt(raw, 10, 64)
	if err != nil {
		return 0
	}
	return int(parsed)
}

func GetStories(category string, page int, limit int) []Story {
	top, err := httpjson.Get[[]int](url(category, h.NewQs()))
	if err != nil {
		slog.Error("failed to load top stories", slog.String("err", err.Error()))
		return make([]Story, 0)
	}
	ids := *top
	start := page * limit
	end := start + limit

	if start > len(ids) {
		return make([]Story, 0)
	}

	if end > len(ids) {
		end = len(ids)
	}

	return batch.ParallelProcess[int, Story](
		ids[start:end],
		50,
		func(id int) Story {
			story, err := GetStory(id)
			if err != nil {
				slog.Error("failed to load story", slog.Int("id", id), slog.String("err", err.Error()))
				return Story{}
			}
			return *story
		},
	)
}

func GetTopStories(page int, limit int) []Story {
	return GetStories("topstories", page, limit)
}

func GetBestStories(page int, limit int) []Story {
	return GetStories("beststories", page, limit)
}

func GetNewStories(page int, limit int) []Story {
	return GetStories("newstories", page, limit)
}

func GetComments(ids []int) []Comment {
	return batch.ParallelProcess(
		ids,
		50,
		func(id int) Comment {
			comment, err := GetComment(id)
			if err != nil {
				slog.Error("failed to load comment", slog.Int("id", id), slog.String("err", err.Error()))
				return Comment{}
			}
			return *comment
		},
	)
}

func GetComment(id int) (*Comment, error) {
	c, err := httpjson.Get[Comment](url(fmt.Sprintf("item/%d", id), h.NewQs()))
	if err != nil {
		return nil, err
	}
	c.Time = timeformat.ParseUnix(c.TimeRaw)
	return c, nil
}

func GetStory(id int) (*Story, error) {
	s, err := httpjson.Get[Story](url(fmt.Sprintf("item/%d", id), h.NewQs()))
	if err != nil {
		return nil, err
	}
	s.Time = timeformat.ParseUnix(s.TimeRaw)
	return s, nil
}
