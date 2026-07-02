package youtube

// PlaylistItemService provides access to the playlistItems collection.
type PlaylistItemService service

// PlaylistItem is a single item within a playlist.
type PlaylistItem struct {
	Kind    string            `json:"kind,omitempty"`
	Etag    string            `json:"etag,omitempty"`
	ID      string            `json:"id,omitempty"`
	Snippet *PlaylistItemSnippet `json:"snippet,omitempty"`
	Status  *PlaylistItemStatus  `json:"status,omitempty"`
}

// PlaylistItemSnippet holds playlist item metadata.
type PlaylistItemSnippet struct {
	PublishedAt       string        `json:"publishedAt,omitempty"`
	ChannelID         string        `json:"channelId,omitempty"`
	ChannelTitle      string        `json:"channelTitle,omitempty"`
	Title             string        `json:"title,omitempty"`
	Description       string        `json:"description,omitempty"`
	Thumbnails        *ThumbnailSet `json:"thumbnails,omitempty"`
	PlaylistID        string        `json:"playlistId,omitempty"`
	Position          int           `json:"position,omitempty"`
	ResourceID        *ResourceID   `json:"resourceId,omitempty"`
	VideoOwnerChannelTitle string    `json:"videoOwnerChannelTitle,omitempty"`
	VideoOwnerChannelID   string    `json:"videoOwnerChannelId,omitempty"`
}

// PlaylistItemStatus holds the privacy status of a playlist item.
type PlaylistItemStatus struct {
	PrivacyStatus string `json:"privacyStatus,omitempty"`
}

// PlaylistItemListResponse is the response from playlistItems.list.
type PlaylistItemListResponse struct {
	ListResponse
	Items []PlaylistItem `json:"items,omitempty"`
}

// PlaylistItemListParams are the parameters for playlistItems.list.
type PlaylistItemListParams struct {
	Part       string `url:"part,omitempty"`
	PlaylistID string `url:"playlistId,omitempty"`
	ID         string `url:"id,omitempty"`
	MaxResults int    `url:"maxResults,omitempty"`
	PageToken  string `url:"pageToken,omitempty"`
}

// List returns items in a playlist.
func (s *PlaylistItemService) List(part, playlistID string, maxResults int) (*PlaylistItemListResponse, error) {
	if maxResults == 0 {
		maxResults = 50
	}
	return s.list(&PlaylistItemListParams{
		Part:       part,
		PlaylistID: playlistID,
		MaxResults: maxResults,
	})
}

func (s *PlaylistItemService) list(params *PlaylistItemListParams) (*PlaylistItemListResponse, error) {
	if params.Part == "" {
		params.Part = "snippet"
	}
	res := new(PlaylistItemListResponse)
	_, err := s.client.base.New().Get("playlistItems").QueryStruct(params).Receive(res, nil)
	if err != nil {
		return nil, err
	}
	return res, nil
}
