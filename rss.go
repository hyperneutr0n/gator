package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"
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

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
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

	rssFeeds.Channel.Title = html.UnescapeString(rssFeeds.Channel.Title)
	rssFeeds.Channel.Description = html.UnescapeString(rssFeeds.Channel.Description)

	for i, rssItem := range rssFeeds.Channel.Item {
		rssFeeds.Channel.Item[i].Title = html.UnescapeString(rssItem.Title)
		rssFeeds.Channel.Item[i].Description = html.UnescapeString(rssItem.Description)
	}

	return &rssFeeds, nil
}
