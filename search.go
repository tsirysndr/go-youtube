package youtube

type SearchService service

type SearchResult struct {
	Kind          string `json:"kind,omitempty"`
	Etag          string `json:"etag,omitempty"`
	NextPageToken string `json:"nextPageToken,omitempty"`
	RegionCode    string `json:"regionCode,omitempty"`
	PageInfo      *struct {
		TotalResults   int `json:"totalResults,omitempty"`
		ResultsPerPage int `json:"resultsPerPage,omitempty"`
	} `json:"pageInfo,omitempty"`
	Items []Item `json:"items,omitempty"`
}

type Item struct {
	Kind string `json:"kind,omitempty"`
	Etag string `json:"etag,omitempty"`
	ID   *struct {
		Kind      string  `json:"kind,omitempty"`
		ChannelID *string `json:"channelId,omitempty"`
		VideoID   *string `json:"videoId,omitempty"`
	} `json:"id,omitempty"`
	Snippet Snippet `json:"snippet,omitempty"`
}

type Snippet struct {
	PublishedAt string     `json:"publishedAt,omitempty"`
	ChannelID   string     `json:"channelId,omitempty"`
	Title       string     `json:"title,omitempty"`
	Description string     `json:"description,omitempty"`
	Thumbnails  Thumbnails `json:"thumbnails,omitempty"`
}

type Thumbnails struct {
	Default *struct {
		URL    string `json:"url,omitempty"`
		Width  int    `json:"width,omitempty"`
		Height int    `json:"height,omitempty"`
	} `json:"default,omitempty"`
	Medium *struct {
		URL    string `json:"url,omitempty"`
		Width  int    `json:"width,omitempty"`
		Height int    `json:"height,omitempty"`
	} `json:"medium,omitempty"`
	High *struct {
		URL    string `json:"url,omitempty"`
		Width  int    `json:"width,omitempty"`
		Height int    `json:"height,omitempty"`
	} `json:"high,omitempty"`
}

type SearchParams struct {
	Part       string `url:"part,omitempty"`
	MaxResults int    `url:"maxResults,omitempty"`
	Q          string `url:"q,omitempty"`
}

func (s *SearchService) Search(q string) (*SearchResult, error) {
	var err error
	params := &SearchParams{
		"snippet",
		25,
		q,
	}
	result := new(SearchResult)
	s.client.base.Get("search").QueryStruct(params).Receive(result, err)
	return result, err
}
