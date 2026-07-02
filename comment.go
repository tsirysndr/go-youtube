package youtube

// CommentService provides access to the commentThreads and comments collections.
type CommentService service

// ── CommentThread ─────────────────────────────────────────────────────────────

// CommentThread is a top-level comment thread on a video.
type CommentThread struct {
	Kind    string                  `json:"kind,omitempty"`
	Etag    string                  `json:"etag,omitempty"`
	ID      string                  `json:"id,omitempty"`
	Snippet *CommentThreadSnippet   `json:"snippet,omitempty"`
	Replies *CommentThreadReplies   `json:"replies,omitempty"`
}

// CommentThreadSnippet holds comment thread metadata.
type CommentThreadSnippet struct {
	VideoID         string   `json:"videoId,omitempty"`
	TopLevelComment *Comment `json:"topLevelComment,omitempty"`
	CanReply        bool     `json:"canReply,omitempty"`
	TotalReplyCount int      `json:"totalReplyCount,omitempty"`
	IsPublic        bool     `json:"isPublic,omitempty"`
}

// CommentThreadReplies holds paginated replies to a thread.
type CommentThreadReplies struct {
	Comments []Comment `json:"comments,omitempty"`
}

// CommentThreadListResponse is the response from commentThreads.list.
type CommentThreadListResponse struct {
	ListResponse
	Items []CommentThread `json:"items,omitempty"`
}

// CommentThreadListParams are the parameters for commentThreads.list.
type CommentThreadListParams struct {
	Part            string `url:"part,omitempty"`
	VideoID         string `url:"videoId,omitempty"`
	ChannelID       string `url:"channelId,omitempty"`
	AllThreadsRelatedToChannelID string `url:"allThreadsRelatedToChannelId,omitempty"`
	ID              string `url:"id,omitempty"`
	MaxResults      int    `url:"maxResults,omitempty"`
	PageToken       string `url:"pageToken,omitempty"`
	ModerationStatus string `url:"moderationStatus,omitempty"`
	Order           string `url:"order,omitempty"`
	SearchTerms     string `url:"searchTerms,omitempty"`
	TextFormat      string `url:"textFormat,omitempty"`
}

// ThreadsList returns comment threads for a video.
func (s *CommentService) ThreadsList(part, videoID string, maxResults int) (*CommentThreadListResponse, error) {
	if maxResults == 0 {
		maxResults = 20
	}
	return s.threadsList(&CommentThreadListParams{
		Part:       part,
		VideoID:    videoID,
		MaxResults: maxResults,
	})
}

func (s *CommentService) threadsList(params *CommentThreadListParams) (*CommentThreadListResponse, error) {
	if params.Part == "" {
		params.Part = "snippet"
	}
	res := new(CommentThreadListResponse)
	_, err := s.client.base.New().Get("commentThreads").QueryStruct(params).Receive(res, nil)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// ── Comment ───────────────────────────────────────────────────────────────────

// Comment is a single comment resource.
type Comment struct {
	Kind    string          `json:"kind,omitempty"`
	Etag    string          `json:"etag,omitempty"`
	ID      string          `json:"id,omitempty"`
	Snippet *CommentSnippet `json:"snippet,omitempty"`
}

// CommentSnippet holds comment body and metadata.
type CommentSnippet struct {
	VideoID        string `json:"videoId,omitempty"`
	ChannelID      string `json:"channelId,omitempty"`
	CommentThreadID string `json:"commentThreadId,omitempty"`
	TextDisplay    string `json:"textDisplay,omitempty"`
	TextOriginal   string `json:"textOriginal,omitempty"`
	AuthorDisplayName string `json:"authorDisplayName,omitempty"`
	AuthorChannelURL  string `json:"authorChannelUrl,omitempty"`
	AuthorProfileImageURL string `json:"authorProfileImageUrl,omitempty"`
	AuthorChannelID *struct {
		Value string `json:"value,omitempty"`
	} `json:"authorChannelId,omitempty"`
	CanRate    bool   `json:"canRate,omitempty"`
	ViewerRating string `json:"viewerRating,omitempty"`
	LikeCount  int    `json:"likeCount,omitempty"`
	PublishedAt string `json:"publishedAt,omitempty"`
	UpdatedAt   string `json:"updatedAt,omitempty"`
	ParentID    string `json:"parentId,omitempty"`
	ModerationStatus string `json:"moderationStatus,omitempty"`
}

// CommentListResponse is the response from comments.list.
type CommentListResponse struct {
	ListResponse
	Items []Comment `json:"items,omitempty"`
}

// CommentListParams are the parameters for comments.list.
type CommentListParams struct {
	Part       string `url:"part,omitempty"`
	ID         string `url:"id,omitempty"`
	ParentID   string `url:"parentId,omitempty"`
	MaxResults int    `url:"maxResults,omitempty"`
	PageToken  string `url:"pageToken,omitempty"`
	TextFormat string `url:"textFormat,omitempty"`
}

// CommentsList returns replies to a comment thread (by parent comment ID).
func (s *CommentService) CommentsList(part, parentID string, maxResults int) (*CommentListResponse, error) {
	if maxResults == 0 {
		maxResults = 20
	}
	return s.commentsList(&CommentListParams{
		Part:     part,
		ParentID: parentID,
		MaxResults: maxResults,
	})
}

func (s *CommentService) commentsList(params *CommentListParams) (*CommentListResponse, error) {
	if params.Part == "" {
		params.Part = "snippet"
	}
	res := new(CommentListResponse)
	_, err := s.client.base.New().Get("comments").QueryStruct(params).Receive(res, nil)
	if err != nil {
		return nil, err
	}
	return res, nil
}
