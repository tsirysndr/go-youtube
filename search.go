package youtube

// SearchService provides access to the search collection.
type SearchService service

// SearchResult is the response envelope for search.list.
type SearchResult struct {
	ListResponse

	Items []SearchItem `json:"items,omitempty"`
}

// SearchItem is a single search result item.
type SearchItem struct {
	Kind    string      `json:"kind,omitempty"`
	Etag    string      `json:"etag,omitempty"`
	ID      *ResourceID `json:"id,omitempty"`
	Snippet *Snippet    `json:"snippet,omitempty"`
	// Note: the Snippet type is defined in youtube.go
}

// SearchParams are the parameters for search.list.
type SearchParams struct {
	Part            string `url:"part,omitempty"`
	Q               string `url:"q,omitempty"`
	MaxResults      int    `url:"maxResults,omitempty"`
	Order           string `url:"order,omitempty"`
	PageToken       string `url:"pageToken,omitempty"`
	RegionCode      string `url:"regionCode,omitempty"`
	RelevanceLanguage string `url:"relevanceLanguage,omitempty"`
	SafeSearch      string `url:"safeSearch,omitempty"`
	Type            string `url:"type,omitempty"`
	VideoCaption    string `url:"videoCaption,omitempty"`
	VideoCategoryID string `url:"videoCategoryId,omitempty"`
	VideoDuration   string `url:"videoDuration,omitempty"`
	VideoEmbeddable string `url:"videoEmbeddable,omitempty"`
	VideoLicense    string `url:"videoLicense,omitempty"`
	VideoSyndicated string `url:"videoSyndicated,omitempty"`
	VideoType       string `url:"videoType,omitempty"`
	ChannelID       string `url:"channelId,omitempty"`
	ChannelType     string `url:"channelType,omitempty"`
	Location        string `url:"location,omitempty"`
	LocationRadius  string `url:"locationRadius,omitempty"`
	PublishedAfter  string `url:"publishedAfter,omitempty"`
	PublishedBefore string `url:"publishedBefore,omitempty"`
}

// Search performs a YouTube search. Pass an empty SearchParams{} for defaults
// (part="snippet", maxResults=25).
func (s *SearchService) Search(q string, params *SearchParams) (*SearchResult, error) {
	if params == nil {
		params = &SearchParams{Part: "snippet", MaxResults: 25}
	}
	params.Q = q
	if params.Part == "" {
		params.Part = "snippet"
	}
	if params.MaxResults == 0 {
		params.MaxResults = 25
	}

	res := new(SearchResult)
	_, err := s.client.base.New().Get("search").QueryStruct(params).Receive(res, nil)
	if err != nil {
		return nil, err
	}
	return res, nil
}
