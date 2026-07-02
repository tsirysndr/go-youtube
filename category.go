package youtube

// CategoryService provides access to the videoCategories collection.
type CategoryService service

// Category is a video category resource.
type Category struct {
	Kind    string          `json:"kind,omitempty"`
	Etag    string          `json:"etag,omitempty"`
	ID      string          `json:"id,omitempty"`
	Snippet *CategorySnippet `json:"snippet,omitempty"`
}

// CategorySnippet holds category metadata.
type CategorySnippet struct {
	Title       string `json:"title,omitempty"`
	Assignable  bool   `json:"assignable,omitempty"`
	ChannelID   string `json:"channelId,omitempty"`
}

// CategoryListResponse is the response from videoCategories.list.
type CategoryListResponse struct {
	ListResponse
	Items []Category `json:"items,omitempty"`
}

// CategoryListParams are the parameters for videoCategories.list.
type CategoryListParams struct {
	Part       string `url:"part,omitempty"`
	ID         string `url:"id,omitempty"`
	RegionCode string `url:"regionCode,omitempty"`
}

// List returns video categories for a region.
func (s *CategoryService) List(regionCode string) (*CategoryListResponse, error) {
	return s.list(&CategoryListParams{
		Part:       "snippet",
		RegionCode: regionCode,
	})
}

func (s *CategoryService) list(params *CategoryListParams) (*CategoryListResponse, error) {
	if params.Part == "" {
		params.Part = "snippet"
	}
	res := new(CategoryListResponse)
	_, err := s.client.base.New().Get("videoCategories").QueryStruct(params).Receive(res, nil)
	if err != nil {
		return nil, err
	}
	return res, nil
}
