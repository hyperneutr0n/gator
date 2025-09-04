package main

import (
	"context"
	"encoding/xml"
	// "errors"
	"fmt"
	"html"
	"io"
	"net/http"
)

type RSSFeed struct {
	Channel struct {
		Title		string		`xml:"title"`
		Link		string		`xml:"link"`
		Description	string		`xml:"description"`
		Item		[]RSSItem	`xml:"item"`
	}	`xml:"channel"`
}

type RSSItem struct {
	Title		string	`xml:"title"`
	Link		string	`xml:"link"`
	Description	string	`xml:"description"`
	PubDate		string	`xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "", feedURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create new http request: %w", err)
	}

	req.Header.Set("User-Agent", "gator")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed during http request: %w", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %w", err)
	}

	var rssFeeds RSSFeed
	if err := xml.Unmarshal(body, &rssFeeds); err != nil {
		return nil, fmt.Errorf("failed unmarshalling rss feed: %w", err)
	}

	return &rssFeeds, nil
}

func handlerAgg(s *state, cmd command) error {
	// if len(cmd.Args) < 1 {
	// 	return errors.New("agg expects feed url as an argument")
	// }

	rssFeeds, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("failed fetching feeds: %w", err)
	}

	rssFeeds.Channel.Title = html.UnescapeString(rssFeeds.Channel.Title)
	rssFeeds.Channel.Description = html.UnescapeString(rssFeeds.Channel.Description)

	for _, rssItem := range rssFeeds.Channel.Item {
		rssItem.Title = html.UnescapeString(rssItem.Title)
		rssItem.Description = html.UnescapeString(rssItem.Description)
	}

	fmt.Println(rssFeeds)
	return nil
}