package youtube

// VideoService provides access to the videos collection.
type VideoService service

// Video is a single video resource.
type Video struct {
	Kind           string         `json:"kind,omitempty"`
	Etag           string         `json:"etag,omitempty"`
	ID             string         `json:"id,omitempty"`
	Snippet        *Snippet       `json:"snippet,omitempty"`
	ContentDetails *VideoContentDetails `json:"contentDetails,omitempty"`
	Statistics     *VideoStatistics     `json:"statistics,omitempty"`
	Status         *VideoStatus         `json:"status,omitempty"`
	Player         *VideoPlayer         `json:"player,omitempty"`
}

// VideoContentDetails holds video-level content metadata.
type VideoContentDetails struct {
	Duration        string `json:"duration,omitempty"`
	Dimension       string `json:"dimension,omitempty"`
	Definition      string `json:"definition,omitempty"`
	Caption         string `json:"caption,omitempty"`
	LicensedContent bool   `json:"licensedContent,omitempty"`
	RegionRestriction *struct {
		Allowed []string `json:"allowed,omitempty"`
		Blocked []string `json:"blocked,omitempty"`
	} `json:"regionRestriction,omitempty"`
	ContentRating *struct {
		YtRating string `json:"ytRating,omitempty"`
	} `json:"contentRating,omitempty"`
	Projection string `json:"projection,omitempty"`
}

// VideoStatistics holds video view/like/dislike/comment counts.
type VideoStatistics struct {
	ViewCount     int64 `json:"viewCount,omitempty,string"`
	LikeCount     int64 `json:"likeCount,omitempty,string"`
	DislikeCount  int64 `json:"dislikeCount,omitempty,string"`
	FavoriteCount int64 `json:"favoriteCount,omitempty,string"`
	CommentCount  int64 `json:"commentCount,omitempty,string"`
}

// VideoStatus holds video publishing/live status.
type VideoStatus struct {
	UploadStatus string `json:"uploadStatus,omitempty"`
	FailureReason string `json:"failureReason,omitempty"`
	RejectionReason string `json:"rejectionReason,omitempty"`
	PrivacyStatus  string `json:"privacyStatus,omitempty"`
	PubsubAutoFire bool   `json:"pubsubAutoFire,omitempty,string"`
	License        string `json:"license,omitempty"`
	Embeddable     bool   `json:"embeddable,omitempty"`
	MadeForKids    bool   `json:"madeForKids,omitempty"`
}

// VideoPlayer holds the embed HTML for a video.
type VideoPlayer struct {
	EmbedHTML string `json:"embedHtml,omitempty"`
	EmbedHeight int  `json:"embedHeight,omitempty"`
	EmbedWidth  int  `json:"embedWidth,omitempty"`
}

// VideoListResponse is the response from videos.list.
type VideoListResponse struct {
	ListResponse
	Items []Video `json:"items,omitempty"`
}

// VideoListParams are the parameters for videos.list.
type VideoListParams struct {
	Part     string `url:"part,omitempty"`
	ID       string `url:"id,omitempty"`
	Chart    string `url:"chart,omitempty"`
	RegionCode string `url:"regionCode,omitempty"`
	MaxResults int  `url:"maxResults,omitempty"`
	PageToken string `url:"pageToken,omitempty"`
	VideoCategoryID string `url:"videoCategoryId,omitempty"`
	MyRating string `url:"myRating,omitempty"`
	MaxHeight int `url:"maxHeight,omitempty"`
	MaxWidth  int `url:"maxWidth,omitempty"`
}

// List returns videos by one or more video IDs.
func (s *VideoService) List(part string, ids ...string) (*VideoListResponse, error) {
	return s.list(&VideoListParams{
		Part: part,
		ID:   joinIDs(ids...),
	})
}

// ListByChart returns videos from a chart (e.g., "mostPopular").
func (s *VideoService) ListByChart(part, chart, regionCode string, maxResults int) (*VideoListResponse, error) {
	if maxResults == 0 {
		maxResults = 25
	}
	return s.list(&VideoListParams{
		Part:       part,
		Chart:      chart,
		RegionCode: regionCode,
		MaxResults: maxResults,
	})
}

// ListByRating returns videos the authenticated user has rated.
func (s *VideoService) ListByRating(part, rating string) (*VideoListResponse, error) {
	return s.list(&VideoListParams{
		Part:     part,
		MyRating: rating,
	})
}

func (s *VideoService) list(params *VideoListParams) (*VideoListResponse, error) {
	if params.Part == "" {
		params.Part = "snippet"
	}
	res := new(VideoListResponse)
	_, err := s.client.base.New().Get("videos").QueryStruct(params).Receive(res, nil)
	if err != nil {
		return nil, err
	}
	return res, nil
}
