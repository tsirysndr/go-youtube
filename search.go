package youtube

type SearchService service

type SearchResult struct {
	Kind          string `json:"kind"`
	Etag          string `json:"etag"`
	NextPageToken string `json:"nextPageToken"`
	RegionCode    string `json:"regionCode"`
	PageInfo      *struct {
		TotalResults   int `json:"totalResults"`
		ResultsPerPage int `json:"resultsPerPage"`
	} `json:"pageInfo"`
	Items []Item `json:"items"`
}

type Item struct {
	Kind string `json:"kind"`
	Etag string `json:"etag"`
	ID   *struct {
		Kind      string  `json:"kind"`
		ChannelID *string `json:"channelId"`
		VideoID   *string `json:"videoId"`
	} `json:"id"`
	Snippet Snippet `json:"snippet"`
}

type Snippet struct {
	PublishedAt string     `json:"publishedAt"`
	ChannelID   string     `json:"channelId"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Thumbnails  Thumbnails `json:"thumbnails"`
}

type Thumbnails struct {
	Default *struct {
		URL    string `json:"url"`
		Width  int    `json:"width"`
		Height int    `json:"height"`
	} `json:"default"`
	Medium *struct {
		URL    string `json:"url"`
		Width  int    `json:"width"`
		Height int    `json:"height"`
	} `json:"medium"`
	High *struct {
		URL    string `json:"url"`
		Width  int    `json:"width"`
		Height int    `json:"height"`
	} `json:"high"`
}

type SearchParams struct {
	Part       string `url:"part"`
	MaxResults int    `url:"maxResults"`
	Q          string `url:"q"`
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
