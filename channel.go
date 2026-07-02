package youtube

// ChannelService provides access to the channels collection.
type ChannelService service

// Channel is a single YouTube channel resource.
type Channel struct {
	Kind           string              `json:"kind,omitempty"`
	Etag           string              `json:"etag,omitempty"`
	ID             string              `json:"id,omitempty"`
	Snippet        *Snippet            `json:"snippet,omitempty"`
	Statistics     *ChannelStatistics  `json:"statistics,omitempty"`
	ContentDetails *ChannelContentDetails `json:"contentDetails,omitempty"`
	Status         *ChannelStatus      `json:"status,omitempty"`
}

// ChannelStatistics holds channel-level analytics.
type ChannelStatistics struct {
	ViewCount             int64 `json:"viewCount,omitempty,string"`
	SubscriberCount       int64 `json:"subscriberCount,omitempty,string"`
	HiddenSubscriberCount bool  `json:"hiddenSubscriberCount,omitempty"`
	VideoCount            int64 `json:"videoCount,omitempty,string"`
}

// ChannelContentDetails holds channel content-related metadata.
type ChannelContentDetails struct {
	RelatedPlaylists *struct {
		Uploads       string `json:"uploads,omitempty"`
		Likes         string `json:"likes,omitempty"`
		WatchHistory  string `json:"watchHistory,omitempty"`
		WatchLater    string `json:"watchLater,omitempty"`
	} `json:"relatedPlaylists,omitempty"`
}

// ChannelStatus holds channel status information.
type ChannelStatus struct {
	PrivacyStatus string `json:"privacyStatus,omitempty"`
	IsLinked      bool   `json:"isLinked,omitempty"`
	LongUploadsStatus string `json:"longUploadsStatus,omitempty"`
	MadeForKids    bool   `json:"madeForKids,omitempty"`
}

// ChannelListResponse is the response from channels.list.
type ChannelListResponse struct {
	ListResponse
	Items []Channel `json:"items,omitempty"`
}

// ChannelListParams are the parameters for channels.list.
type ChannelListParams struct {
	Part        string `url:"part,omitempty"`
	ID          string `url:"id,omitempty"`
	ForUsername string `url:"forUsername,omitempty"`
	Mine        bool   `url:"mine,omitempty"`
	ManagedByMe bool   `url:"managedByMe,omitempty"`
	CategoryID  string `url:"categoryId,omitempty"`
	MaxResults  int    `url:"maxResults,omitempty"`
	PageToken   string `url:"pageToken,omitempty"`
}

// List returns channels by one or more channel IDs.
func (s *ChannelService) List(part string, ids ...string) (*ChannelListResponse, error) {
	return s.list(&ChannelListParams{
		Part: part,
		ID:   joinIDs(ids...),
	})
}

// ListByUsername looks up a channel by its legacy username.
func (s *ChannelService) ListByUsername(part, username string) (*ChannelListResponse, error) {
	return s.list(&ChannelListParams{
		Part:        part,
		ForUsername: username,
	})
}

// ListMine returns the authenticated user's channel.
func (s *ChannelService) ListMine(part string) (*ChannelListResponse, error) {
	return s.list(&ChannelListParams{
		Part: part,
		Mine: true,
	})
}

func (s *ChannelService) list(params *ChannelListParams) (*ChannelListResponse, error) {
	if params.Part == "" {
		params.Part = "snippet"
	}
	res := new(ChannelListResponse)
	_, err := s.client.base.New().Get("channels").QueryStruct(params).Receive(res, nil)
	if err != nil {
		return nil, err
	}
	return res, nil
}
