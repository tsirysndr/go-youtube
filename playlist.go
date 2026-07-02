package youtube

// PlaylistService provides access to the playlists collection.
type PlaylistService service

// Playlist is a single playlist resource.
type Playlist struct {
	Kind           string              `json:"kind,omitempty"`
	Etag           string              `json:"etag,omitempty"`
	ID             string              `json:"id,omitempty"`
	Snippet        *Snippet            `json:"snippet,omitempty"`
	Status         *PlaylistStatus     `json:"status,omitempty"`
	ContentDetails *PlaylistContentDetails `json:"contentDetails,omitempty"`
	Player         *PlaylistPlayer     `json:"player,omitempty"`
}

// PlaylistStatus holds the privacy status of a playlist.
type PlaylistStatus struct {
	PrivacyStatus string `json:"privacyStatus,omitempty"`
}

// PlaylistContentDetails holds item count for a playlist.
type PlaylistContentDetails struct {
	ItemCount int `json:"itemCount,omitempty"`
}

// PlaylistPlayer holds the embed player for a playlist.
type PlaylistPlayer struct {
	EmbedHTML string `json:"embedHtml,omitempty"`
}

// PlaylistListResponse is the response from playlists.list.
type PlaylistListResponse struct {
	ListResponse
	Items []Playlist `json:"items,omitempty"`
}

// PlaylistListParams are the parameters for playlists.list.
type PlaylistListParams struct {
	Part       string `url:"part,omitempty"`
	ID         string `url:"id,omitempty"`
	ChannelID  string `url:"channelId,omitempty"`
	Mine       bool   `url:"mine,omitempty"`
	MaxResults int    `url:"maxResults,omitempty"`
	PageToken  string `url:"pageToken,omitempty"`
}

// List returns playlists by one or more playlist IDs.
func (s *PlaylistService) List(part string, ids ...string) (*PlaylistListResponse, error) {
	return s.list(&PlaylistListParams{
		Part: part,
		ID:   joinIDs(ids...),
	})
}

// ListByChannel returns playlists owned by the given channel.
func (s *PlaylistService) ListByChannel(part, channelID string, maxResults int) (*PlaylistListResponse, error) {
	if maxResults == 0 {
		maxResults = 50
	}
	return s.list(&PlaylistListParams{
		Part:       part,
		ChannelID:  channelID,
		MaxResults: maxResults,
	})
}

// ListMine returns the authenticated user's playlists.
func (s *PlaylistService) ListMine(part string, maxResults int) (*PlaylistListResponse, error) {
	if maxResults == 0 {
		maxResults = 50
	}
	return s.list(&PlaylistListParams{
		Part:       part,
		Mine:       true,
		MaxResults: maxResults,
	})
}

func (s *PlaylistService) list(params *PlaylistListParams) (*PlaylistListResponse, error) {
	if params.Part == "" {
		params.Part = "snippet"
	}
	res := new(PlaylistListResponse)
	_, err := s.client.base.New().Get("playlists").QueryStruct(params).Receive(res, nil)
	if err != nil {
		return nil, err
	}
	return res, nil
}
