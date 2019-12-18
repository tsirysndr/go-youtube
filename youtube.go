package youtube

import (
	"github.com/dghubble/sling"
)

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

func NewClient(accessToken string) *Client {
	c := &Client{}
	base := sling.New().Base("https://www.googleapis.com/youtube/v3/")
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
