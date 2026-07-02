// Package youtube is a Go client library for the YouTube Data API v3.
package youtube

import (
	"context"
	"net/http"

	"github.com/dghubble/sling"
	"golang.org/x/oauth2"
)

// Client manages communication with the YouTube Data API v3.
type Client struct {
	base         *sling.Sling
	common       service

	Category     *CategoryService
	Channel      *ChannelService
	Comment      *CommentService
	Playlist     *PlaylistService
	PlaylistItem *PlaylistItemService
	Search       *SearchService
	Thumbnail    *ThumbnailService
	Video        *VideoService
}

type service struct {
	client *Client
}

// KeyParam holds the API key query parameter sent with every request.
type KeyParam struct {
	APIKey string `url:"key,omitempty"`
}

// NewClient creates a new YouTube API client authenticated with an OAuth2 access
// token and an API key. Pass an empty string for either if not needed.
func NewClient(accessToken, apiKey string) *Client {
	httpClient := http.DefaultClient
	if accessToken != "" {
		config := &oauth2.Config{}
		token := &oauth2.Token{AccessToken: accessToken}
		httpClient = config.Client(context.Background(), token)
	}
	return newWithHTTPClient(httpClient, apiKey)
}

// NewClientWithKey creates a client that uses only an API key (no OAuth).
// Suitable for public, read-only YouTube Data API operations.
func NewClientWithKey(apiKey string) *Client {
	return newWithHTTPClient(http.DefaultClient, apiKey)
}

func newWithHTTPClient(httpClient *http.Client, apiKey string) *Client {
	base := sling.New().Base("https://www.googleapis.com/youtube/v3/").Client(httpClient)
	if apiKey != "" {
		base = base.QueryStruct(&KeyParam{APIKey: apiKey})
	}
	c := &Client{}
	c.base = base
	c.common.client = c
	c.Category = (*CategoryService)(&c.common)
	c.Channel = (*ChannelService)(&c.common)
	c.Comment = (*CommentService)(&c.common)
	c.Playlist = (*PlaylistService)(&c.common)
	c.PlaylistItem = (*PlaylistItemService)(&c.common)
	c.Search = (*SearchService)(&c.common)
	c.Thumbnail = (*ThumbnailService)(&c.common)
	c.Video = (*VideoService)(&c.common)
	return c
}

// ─── Shared response types ───────────────────────────────────────────────────

// PageInfo holds paging metadata returned by most list endpoints.
type PageInfo struct {
	TotalResults   int `json:"totalResults,omitempty"`
	ResultsPerPage int `json:"resultsPerPage,omitempty"`
}

// ListResponse is the common envelope for paginated list endpoints.
type ListResponse struct {
	Kind          string    `json:"kind,omitempty"`
	Etag          string    `json:"etag,omitempty"`
	NextPageToken string    `json:"nextPageToken,omitempty"`
	PrevPageToken string    `json:"prevPageToken,omitempty"`
	PageInfo      *PageInfo `json:"pageInfo,omitempty"`
}

// ─── Shared resource snippets ────────────────────────────────────────────────

// ThumbnailSet holds the standard thumbnail resolutions returned by the API.
type ThumbnailSet struct {
	Default *Thumbnail `json:"default,omitempty"`
	Medium  *Thumbnail `json:"medium,omitempty"`
	High    *Thumbnail `json:"high,omitempty"`
	Standard *Thumbnail `json:"standard,omitempty"`
	Maxres   *Thumbnail `json:"maxres,omitempty"`
}

// Thumbnail represents a single thumbnail image.
type Thumbnail struct {
	URL    string `json:"url,omitempty"`
	Width  int    `json:"width,omitempty"`
	Height int    `json:"height,omitempty"`
}

// ResourceID identifies a YouTube resource (video, channel, playlist).
type ResourceID struct {
	Kind      string  `json:"kind,omitempty"`
	VideoID   *string `json:"videoId,omitempty"`
	ChannelID *string `json:"channelId,omitempty"`
	PlaylistID *string `json:"playlistId,omitempty"`
}

// Snippet is the common snippet object shared by most resources.
type Snippet struct {
	PublishedAt          string        `json:"publishedAt,omitempty"`
	ChannelID            string        `json:"channelId,omitempty"`
	ChannelTitle         string        `json:"channelTitle,omitempty"`
	Title                string        `json:"title,omitempty"`
	Description          string        `json:"description,omitempty"`
	Thumbnails           *ThumbnailSet `json:"thumbnails,omitempty"`
	LiveBroadcastContent string        `json:"liveBroadcastContent,omitempty"`
	DefaultLanguage      string        `json:"defaultLanguage,omitempty"`
}

// joinIDs joins video/channel/playlist IDs with a comma for the API query string.
func joinIDs(ids ...string) string {
	if len(ids) == 0 {
		return ""
	}
	if len(ids) == 1 {
		return ids[0]
	}
	var b []byte
	for i, id := range ids {
		if i > 0 {
			b = append(b, ',')
		}
		b = append(b, id...)
	}
	return string(b)
}
