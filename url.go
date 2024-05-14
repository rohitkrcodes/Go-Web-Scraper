package main

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

type URLFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Language    string    `xml:"language"`
		Item        []URLItem `xml:"item"`
	} `xml:"channel"`
}

type URLItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubdate"`
}

func urlToFeed(url string) (URLFeed, error) {
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := httpClient.Get(url)
	if err != nil {
		return URLFeed{}, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return URLFeed{}, err
	}

	urlFeed := URLFeed{}
	err = xml.Unmarshal(data, &urlFeed)
	if err != nil {
		return URLFeed{}, err
	}

	return urlFeed, nil
}
