package youtube

import (
	"github.com/dghubble/sling"
	"golang.org/x/oauth2"
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

type KeyParam struct {
	key string `url:"key"`
}

func NewClient(accessToken, apikey string) *Client {
	config := &oauth2.Config{}
	token := &oauth2.Token{AccessToken: accessToken}
	httpClient := config.Client(oauth2.NoContext, token)
	params := &KeyParam{apikey}
	base := sling.New().Base("https://www.googleapis.com/youtube/v3/").Client(httpClient)
	c := &Client{}
	c.base = base.QueryStruct(params)
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
