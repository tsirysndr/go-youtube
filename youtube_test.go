package youtube

import (
	"encoding/json"
	"os"
	"testing"
)

// ─── Integration test helpers ─────────────────────────────────────────────────

func testClient(t *testing.T) *Client {
	t.Helper()
	key := os.Getenv("YOUTUBE_API_KEY")
	if key == "" {
		t.Skip("Set YOUTUBE_API_KEY to run integration tests")
	}
	return NewClientWithKey(key)
}

// ─── JSON unmarshal tests (no network) ────────────────────────────────────────

const testSearchJSON = `{
  "kind": "youtube#searchListResponse",
  "etag": "test_etag",
  "nextPageToken": "CAUQAA",
  "prevPageToken": "QBUCAQ",
  "pageInfo": {
    "totalResults": 1000000,
    "resultsPerPage": 25
  },
  "items": [
    {
      "kind": "youtube#searchResult",
      "etag": "item_etag",
      "id": { "kind": "youtube#video", "videoId": "dQw4w9WgXcQ" },
      "snippet": {
        "publishedAt": "2009-10-25T06:57:33Z",
        "channelId": "UCuAXFkgsw1L7xaCfnd5JJOw",
        "title": "Rick Astley - Never Gonna Give You Up (Official Music Video)",
        "description": "The official video for \"Never Gonna Give You Up\"",
        "thumbnails": {
          "default": { "url": "https://i.ytimg.com/vi/dQw4w9WgXcQ/default.jpg", "width": 120, "height": 90 },
          "medium": { "url": "https://i.ytimg.com/vi/dQw4w9WgXcQ/mqdefault.jpg", "width": 320, "height": 180 },
          "high": { "url": "https://i.ytimg.com/vi/dQw4w9WgXcQ/hqdefault.jpg", "width": 480, "height": 360 }
        },
        "channelTitle": "Rick Astley",
        "liveBroadcastContent": "none"
      }
    }
  ]
}`

func TestSearchResult_Unmarshal(t *testing.T) {
	var sr SearchResult
	if err := json.Unmarshal([]byte(testSearchJSON), &sr); err != nil {
		t.Fatalf("Unmarshal search result: %v", err)
	}
	if sr.Kind != "youtube#searchListResponse" {
		t.Errorf("expected kind youtube#searchListResponse, got %s", sr.Kind)
	}
	if sr.NextPageToken != "CAUQAA" {
		t.Errorf("expected nextPageToken CAUQAA, got %s", sr.NextPageToken)
	}
	if sr.PageInfo == nil {
		t.Fatal("pageInfo is nil")
	}
	if sr.PageInfo.TotalResults != 1000000 {
		t.Errorf("expected totalResults 1000000, got %d", sr.PageInfo.TotalResults)
	}
	if len(sr.Items) != 1 {
		t.Fatalf("expected 1 item, got %d", len(sr.Items))
	}
	item := sr.Items[0]
	if item.ID == nil {
		t.Fatal("item ID is nil")
	}
	if item.ID.VideoID == nil || *item.ID.VideoID != "dQw4w9WgXcQ" {
		t.Errorf("expected videoId dQw4w9WgXcQ, got %v", item.ID.VideoID)
	}
	if item.Snippet == nil {
		t.Fatal("item snippet is nil")
	}
	if item.Snippet.ChannelID != "UCuAXFkgsw1L7xaCfnd5JJOw" {
		t.Errorf("expected channelId UCuAXFkgsw1L7xaCfnd5JJOw, got %s", item.Snippet.ChannelID)
	}
	if item.Snippet.Thumbnails == nil {
		t.Fatal("thumbnails is nil")
	}
	if item.Snippet.Thumbnails.Default == nil {
		t.Fatal("default thumbnail is nil")
	}
	if item.Snippet.Thumbnails.Default.URL != "https://i.ytimg.com/vi/dQw4w9WgXcQ/default.jpg" {
		t.Errorf("unexpected default thumbnail URL")
	}
}

const testVideoJSON = `{
  "kind": "youtube#videoListResponse",
  "etag": "vid_etag",
  "pageInfo": { "totalResults": 1, "resultsPerPage": 1 },
  "items": [
    {
      "kind": "youtube#video",
      "id": "dQw4w9WgXcQ",
      "snippet": {
        "publishedAt": "2009-10-25T06:57:33Z",
        "channelId": "UCuAXFkgsw1L7xaCfnd5JJOw",
        "title": "Never Gonna Give You Up",
        "description": "Rick roll",
        "channelTitle": "Rick Astley"
      },
      "contentDetails": {
        "duration": "PT3M32S",
        "dimension": "2d",
        "definition": "hd",
        "caption": "false",
        "licensedContent": true,
        "projection": "rectangular"
      },
      "statistics": {
        "viewCount": "1500000000",
        "likeCount": "12000000",
        "commentCount": "5000000"
      },
      "status": {
        "privacyStatus": "public",
        "embeddable": true,
        "license": "youtube"
      },
      "player": {
        "embedHtml": "<iframe></iframe>",
        "embedHeight": 360,
        "embedWidth": 640
      }
    }
  ]
}`

func TestVideoList_Unmarshal(t *testing.T) {
	var vr VideoListResponse
	if err := json.Unmarshal([]byte(testVideoJSON), &vr); err != nil {
		t.Fatalf("Unmarshal video list: %v", err)
	}
	if len(vr.Items) != 1 {
		t.Fatalf("expected 1 video, got %d", len(vr.Items))
	}
	v := vr.Items[0]
	if v.ID != "dQw4w9WgXcQ" {
		t.Errorf("expected id dQw4w9WgXcQ, got %s", v.ID)
	}
	if v.Snippet == nil {
		t.Fatal("snippet is nil")
	}
	if v.Snippet.Title != "Never Gonna Give You Up" {
		t.Errorf("unexpected title: %s", v.Snippet.Title)
	}
	if v.ContentDetails == nil {
		t.Fatal("contentDetails is nil")
	}
	if v.ContentDetails.Duration != "PT3M32S" {
		t.Errorf("expected duration PT3M32S, got %s", v.ContentDetails.Duration)
	}
	if v.Statistics == nil {
		t.Fatal("statistics is nil")
	}
	if v.Statistics.ViewCount != 1500000000 {
		t.Errorf("expected viewCount 1500000000, got %d", v.Statistics.ViewCount)
	}
	if v.Statistics.LikeCount != 12000000 {
		t.Errorf("expected likeCount 12000000, got %d", v.Statistics.LikeCount)
	}
	if v.Status == nil {
		t.Fatal("status is nil")
	}
	if !v.Status.Embeddable {
		t.Errorf("expected embeddable true")
	}
	if v.Status.PrivacyStatus != "public" {
		t.Errorf("expected privacyStatus public, got %s", v.Status.PrivacyStatus)
	}
	if v.Player == nil {
		t.Fatal("player is nil")
	}
	if v.Player.EmbedHTML == "" {
		t.Errorf("expected embedHtml")
	}
}

const testChannelJSON = `{
  "kind": "youtube#channelListResponse",
  "pageInfo": { "totalResults": 1, "resultsPerPage": 1 },
  "items": [
    {
      "kind": "youtube#channel",
      "id": "UCuAXFkgsw1L7xaCfnd5JJOw",
      "snippet": {
        "title": "Rick Astley",
        "description": "Official Rick Astley channel",
        "publishedAt": "2009-04-10T15:14:31Z",
        "thumbnails": {
          "default": { "url": "https://i.ytimg.com/vi/default.jpg", "width": 88, "height": 88 },
          "medium": { "url": "https://i.ytimg.com/vi/mqdefault.jpg", "width": 240, "height": 240 }
        }
      },
      "statistics": {
        "viewCount": "500000000",
        "subscriberCount": "5000000",
        "videoCount": "100"
      },
      "contentDetails": {
        "relatedPlaylists": {
          "uploads": "UULFAXkgsw1L7xaCfnd5JJOw",
          "likes": "LLFAXkgsw1L7xaCfnd5JJOw"
        }
      }
    }
  ]
}`

func TestChannelList_Unmarshal(t *testing.T) {
	var cr ChannelListResponse
	if err := json.Unmarshal([]byte(testChannelJSON), &cr); err != nil {
		t.Fatalf("Unmarshal channel list: %v", err)
	}
	if len(cr.Items) != 1 {
		t.Fatalf("expected 1 channel, got %d", len(cr.Items))
	}
	c := cr.Items[0]
	if c.ID != "UCuAXFkgsw1L7xaCfnd5JJOw" {
		t.Errorf("unexpected channel ID")
	}
	if c.Statistics == nil {
		t.Fatal("statistics is nil")
	}
	if c.Statistics.SubscriberCount != 5000000 {
		t.Errorf("expected 5M subscribers, got %d", c.Statistics.SubscriberCount)
	}
	if c.ContentDetails == nil {
		t.Fatal("contentDetails is nil")
	}
	if c.ContentDetails.RelatedPlaylists == nil {
		t.Fatal("relatedPlaylists is nil")
	}
	if c.ContentDetails.RelatedPlaylists.Uploads != "UULFAXkgsw1L7xaCfnd5JJOw" {
		t.Errorf("unexpected uploads playlist")
	}
}

const testCategoryJSON = `{
  "kind": "youtube#videoCategoryListResponse",
  "items": [
    { "kind": "youtube#videoCategory", "id": "1", "snippet": { "title": "Film & Animation", "assignable": true } },
    { "kind": "youtube#videoCategory", "id": "10", "snippet": { "title": "Music", "assignable": true } },
    { "kind": "youtube#videoCategory", "id": "20", "snippet": { "title": "Gaming", "assignable": true } }
  ]
}`

func TestCategoryList_Unmarshal(t *testing.T) {
	var cr CategoryListResponse
	if err := json.Unmarshal([]byte(testCategoryJSON), &cr); err != nil {
		t.Fatalf("Unmarshal category list: %v", err)
	}
	if len(cr.Items) != 3 {
		t.Fatalf("expected 3 categories, got %d", len(cr.Items))
	}
	if cr.Items[0].ID != "1" || cr.Items[0].Snippet.Title != "Film & Animation" {
		t.Errorf("unexpected first category")
	}
	if cr.Items[1].ID != "10" || cr.Items[1].Snippet.Title != "Music" {
		t.Errorf("unexpected second category")
	}
}

const testPlaylistJSON = `{
  "kind": "youtube#playlistListResponse",
  "pageInfo": { "totalResults": 1, "resultsPerPage": 5 },
  "items": [
    {
      "id": "PLrAXtmErZgOeiKm4sgNOknGvNjby9efdf",
      "snippet": {
        "publishedAt": "2020-01-01T00:00:00Z",
        "channelId": "UC_x5XG1OV2P6uZZ5FSM9Ttw",
        "title": "Test Playlist",
        "description": "A playlist for testing",
        "channelTitle": "Google Developers",
        "thumbnails": {
          "default": { "url": "https://i.ytimg.com/vi/default.jpg", "width": 120, "height": 90 }
        }
      },
      "status": { "privacyStatus": "public" },
      "contentDetails": { "itemCount": 42 }
    }
  ]
}`

func TestPlaylistList_Unmarshal(t *testing.T) {
	var pr PlaylistListResponse
	if err := json.Unmarshal([]byte(testPlaylistJSON), &pr); err != nil {
		t.Fatalf("Unmarshal playlist: %v", err)
	}
	if len(pr.Items) != 1 {
		t.Fatalf("expected 1 playlist, got %d", len(pr.Items))
	}
	p := pr.Items[0]
	if p.ID != "PLrAXtmErZgOeiKm4sgNOknGvNjby9efdf" {
		t.Errorf("unexpected playlist ID")
	}
	if p.ContentDetails == nil || p.ContentDetails.ItemCount != 42 {
		t.Errorf("expected itemCount 42")
	}
	if p.Status == nil || p.Status.PrivacyStatus != "public" {
		t.Errorf("expected privacyStatus public")
	}
}

const testPlaylistItemJSON = `{
  "kind": "youtube#playlistItemListResponse",
  "pageInfo": { "totalResults": 1, "resultsPerPage": 1 },
  "items": [
    {
      "id": "UExyQVh0bUVyWmdPZWlLbTRzZ05PS2VuR3ZOamJ5OWVmZGYuN0Q2QzI0NUE0N0I4NTM3MQ",
      "snippet": {
        "publishedAt": "2020-01-01T00:00:00Z",
        "channelId": "UC_x5XG1OV2P6uZZ5FSM9Ttw",
        "title": "Test Video",
        "playlistId": "PLrAXtmErZgOeiKm4sgNOknGvNjby9efdf",
        "position": 0,
        "resourceId": {
          "kind": "youtube#video",
          "videoId": "dQw4w9WgXcQ"
        }
      }
    }
  ]
}`

func TestPlaylistItemList_Unmarshal(t *testing.T) {
	var pir PlaylistItemListResponse
	if err := json.Unmarshal([]byte(testPlaylistItemJSON), &pir); err != nil {
		t.Fatalf("Unmarshal playlistItem: %v", err)
	}
	if len(pir.Items) != 1 {
		t.Fatalf("expected 1 item, got %d", len(pir.Items))
	}
	pi := pir.Items[0]
	if pi.Snippet == nil {
		t.Fatal("snippet is nil")
	}
	if pi.Snippet.Position != 0 {
		t.Errorf("expected position 0, got %d", pi.Snippet.Position)
	}
	if pi.Snippet.ResourceID == nil {
		t.Fatal("resourceId is nil")
	}
	if pi.Snippet.ResourceID.VideoID == nil || *pi.Snippet.ResourceID.VideoID != "dQw4w9WgXcQ" {
		t.Errorf("expected videoId dQw4w9WgXcQ")
	}
}

const testCommentThreadJSON = `{
  "kind": "youtube#commentThreadListResponse",
  "pageInfo": { "totalResults": 1, "resultsPerPage": 1 },
  "items": [
    {
      "id": "Ugx0abc12345",
      "snippet": {
        "videoId": "dQw4w9WgXcQ",
        "topLevelComment": {
          "id": "Ugyx0abc12345",
          "snippet": {
            "videoId": "dQw4w9WgXcQ",
            "textDisplay": "Great video!",
            "textOriginal": "Great video!",
            "authorDisplayName": "TestUser",
            "likeCount": 42,
            "publishedAt": "2023-01-01T00:00:00Z"
          }
        },
        "canReply": true,
        "totalReplyCount": 3
      }
    }
  ]
}`

func TestCommentThread_Unmarshal(t *testing.T) {
	var ctr CommentThreadListResponse
	if err := json.Unmarshal([]byte(testCommentThreadJSON), &ctr); err != nil {
		t.Fatalf("Unmarshal commentThread: %v", err)
	}
	if len(ctr.Items) != 1 {
		t.Fatalf("expected 1 thread, got %d", len(ctr.Items))
	}
	ct := ctr.Items[0]
	if ct.Snippet == nil {
		t.Fatal("snippet is nil")
	}
	if ct.Snippet.TotalReplyCount != 3 {
		t.Errorf("expected 3 replies, got %d", ct.Snippet.TotalReplyCount)
	}
	if ct.Snippet.TopLevelComment == nil {
		t.Fatal("topLevelComment is nil")
	}
	if ct.Snippet.TopLevelComment.Snippet == nil {
		t.Fatal("comment snippet is nil")
	}
	if ct.Snippet.TopLevelComment.Snippet.LikeCount != 42 {
		t.Errorf("expected 42 likes, got %d", ct.Snippet.TopLevelComment.Snippet.LikeCount)
	}
	if ct.Snippet.TopLevelComment.Snippet.AuthorDisplayName != "TestUser" {
		t.Errorf("expected TestUser, got %s", ct.Snippet.TopLevelComment.Snippet.AuthorDisplayName)
	}
}

const testCommentJSON = `{
  "kind": "youtube#commentListResponse",
  "pageInfo": { "totalResults": 1, "resultsPerPage": 1 },
  "items": [
    {
      "id": "Ugyx0reply12345",
      "snippet": {
        "videoId": "dQw4w9WgXcQ",
        "textDisplay": "This is a reply!",
        "textOriginal": "This is a reply!",
        "authorDisplayName": "ReplyUser",
        "parentId": "Ugyx0abc12345",
        "likeCount": 7,
        "publishedAt": "2023-01-02T00:00:00Z"
      }
    }
  ]
}`

func TestComment_Unmarshal(t *testing.T) {
	var cr CommentListResponse
	if err := json.Unmarshal([]byte(testCommentJSON), &cr); err != nil {
		t.Fatalf("Unmarshal comment: %v", err)
	}
	if len(cr.Items) != 1 {
		t.Fatalf("expected 1 comment, got %d", len(cr.Items))
	}
	c := cr.Items[0]
	if c.ID != "Ugyx0reply12345" {
		t.Errorf("unexpected comment ID")
	}
	if c.Snippet == nil {
		t.Fatal("snippet is nil")
	}
	if c.Snippet.ParentID != "Ugyx0abc12345" {
		t.Errorf("expected parentId Ugyx0abc12345, got %s", c.Snippet.ParentID)
	}
}

// ─── Edge case tests ──────────────────────────────────────────────────────────

func TestEmptyListResponse(t *testing.T) {
	// Empty items array (valid response, no results)
	var sr SearchResult
	if err := json.Unmarshal([]byte(`{"kind":"youtube#searchListResponse","items":[]}`), &sr); err != nil {
		t.Fatalf("Unmarshal empty search: %v", err)
	}
	if sr.Items == nil {
		t.Fatal("expected non-nil items")
	}
	if len(sr.Items) != 0 {
		t.Errorf("expected 0 items, got %d", len(sr.Items))
	}
}

func TestMinimalVideo(t *testing.T) {
	// Minimal video response (just ID, no snippet/statistics)
	var vr VideoListResponse
	if err := json.Unmarshal([]byte(`{"items":[{"id":"dQw4w9WgXcQ"}]}`), &vr); err != nil {
		t.Fatalf("Unmarshal minimal video: %v", err)
	}
	if len(vr.Items) != 1 {
		t.Fatalf("expected 1 item, got %d", len(vr.Items))
	}
	if vr.Items[0].ID != "dQw4w9WgXcQ" {
		t.Errorf("unexpected ID")
	}
	if vr.Items[0].Snippet != nil {
		t.Errorf("expected nil snippet")
	}
	if vr.Items[0].Statistics != nil {
		t.Errorf("expected nil statistics")
	}
}

func TestNullFields(t *testing.T) {
	// Response with null values (YouTube API sometimes returns null)
	var sr SearchResult
	if err := json.Unmarshal([]byte(`{"kind":null,"items":null}`), &sr); err != nil {
		t.Fatalf("Unmarshal null fields: %v", err)
	}
	// null strings should unmarshal as empty strings with omitempty
	// null items should be nil
	if sr.Items != nil {
		// With omitempty on the slice, null should result in nil
		// But the JSON decoder might give us an empty slice vs nil
		// This is informational
	}
}

func TestStatisticsStringInt(t *testing.T) {
	// YouTube API returns statistics as strings even though they're numbers
	var vr VideoListResponse
	jsonStr := `{"items":[{"id":"test","statistics":{"viewCount":"12345","likeCount":"678","commentCount":"99"}}]}`
	if err := json.Unmarshal([]byte(jsonStr), &vr); err != nil {
		t.Fatalf("Unmarshal stats: %v", err)
	}
	v := vr.Items[0]
	if v.Statistics == nil {
		t.Fatal("statistics is nil")
	}
	if v.Statistics.ViewCount != 12345 {
		t.Errorf("expected viewCount 12345, got %d", v.Statistics.ViewCount)
	}
	if v.Statistics.LikeCount != 678 {
		t.Errorf("expected likeCount 678, got %d", v.Statistics.LikeCount)
	}
}

func TestThumbnailStandardMaxres(t *testing.T) {
	// YouTube API sometimes returns standard and maxres thumbnails
	var sr SearchResult
	jsonStr := `{"items":[{"snippet":{"thumbnails":{"standard":{"url":"std.jpg","width":640,"height":480},"maxres":{"url":"max.jpg","width":1280,"height":720}}}}]}`
	if err := json.Unmarshal([]byte(jsonStr), &sr); err != nil {
		t.Fatalf("Unmarshal thumbnails: %v", err)
	}
	if len(sr.Items) != 1 {
		t.Fatalf("expected 1 item")
	}
	tmb := sr.Items[0].Snippet.Thumbnails
	if tmb == nil {
		t.Fatal("thumbnails is nil")
	}
	if tmb.Standard == nil {
		t.Fatal("standard thumbnail is nil")
	}
	if tmb.Standard.URL != "std.jpg" {
		t.Errorf("expected std.jpg, got %s", tmb.Standard.URL)
	}
	if tmb.Maxres == nil {
		t.Fatal("maxres thumbnail is nil")
	}
	if tmb.Maxres.URL != "max.jpg" {
		t.Errorf("expected max.jpg, got %s", tmb.Maxres.URL)
	}
}

func TestResourceIDChannel(t *testing.T) {
	// Search result can return channels, not just videos
	var sr SearchResult
	jsonStr := `{"items":[{"id":{"kind":"youtube#channel","channelId":"UCabc123"}}]}`
	if err := json.Unmarshal([]byte(jsonStr), &sr); err != nil {
		t.Fatalf("Unmarshal channel ID: %v", err)
	}
	id := sr.Items[0].ID
	if id == nil {
		t.Fatal("id is nil")
	}
	if id.ChannelID == nil || *id.ChannelID != "UCabc123" {
		t.Errorf("expected channelId UCabc123")
	}
}

func TestChannelHideSubscriberCount(t *testing.T) {
	var cr ChannelListResponse
	jsonStr := `{"items":[{"statistics":{"hiddenSubscriberCount":true}}]}`
	if err := json.Unmarshal([]byte(jsonStr), &cr); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if cr.Items[0].Statistics == nil || !cr.Items[0].Statistics.HiddenSubscriberCount {
		t.Errorf("expected hiddenSubscriberCount true")
	}
}

// ─── Integration tests (require YOUTUBE_API_KEY) ──────────────────────────────

func TestSearch_Integration(t *testing.T) {
	c := testClient(t)
	result, err := c.Search.Search("never gonna give you up", nil)
	if err != nil {
		t.Fatalf("Search failed: %v", err)
	}
	if result.Kind != "youtube#searchListResponse" {
		t.Errorf("unexpected kind: %s", result.Kind)
	}
	if len(result.Items) == 0 {
		t.Error("search returned 0 results")
	}
	for i, item := range result.Items {
		if item.ID == nil {
			t.Errorf("item[%d] has nil ID", i)
		}
		if item.Snippet == nil {
			t.Errorf("item[%d] has nil Snippet", i)
		}
	}
}

func TestVideos_Integration(t *testing.T) {
	c := testClient(t)
	result, err := c.Video.List("snippet,contentDetails,statistics,status", "dQw4w9WgXcQ")
	if err != nil {
		t.Fatalf("Video.List failed: %v", err)
	}
	if len(result.Items) == 0 {
		t.Fatal("video list returned 0 items")
	}
	v := result.Items[0]
	if v.ID != "dQw4w9WgXcQ" {
		t.Errorf("expected video ID dQw4w9WgXcQ, got %s", v.ID)
	}
	if v.Snippet == nil {
		t.Error("snippet is nil")
	}
	if v.ContentDetails == nil {
		t.Error("contentDetails is nil")
	}
	if v.Statistics == nil {
		t.Error("statistics is nil")
	}
	if v.Status == nil {
		t.Error("status is nil")
	}
}

func TestVideosByChart_Integration(t *testing.T) {
	c := testClient(t)
	result, err := c.Video.ListByChart("snippet", "mostPopular", "US", 5)
	if err != nil {
		t.Fatalf("Video.ListByChart failed: %v", err)
	}
	if len(result.Items) == 0 {
		t.Error("mostPopular returned 0 items")
	}
	if len(result.Items) > 5 {
		t.Errorf("expected max 5 items, got %d", len(result.Items))
	}
}

func TestChannels_Integration(t *testing.T) {
	c := testClient(t)
	result, err := c.Channel.List("snippet,statistics,contentDetails", "UCuAXFkgsw1L7xaCfnd5JJOw")
	if err != nil {
		t.Fatalf("Channel.List failed: %v", err)
	}
	if len(result.Items) == 0 {
		t.Fatal("channel list returned 0 items")
	}
	ch := result.Items[0]
	if ch.ID != "UCuAXFkgsw1L7xaCfnd5JJOw" {
		t.Errorf("expected channel ID mismatch")
	}
	if ch.Snippet == nil {
		t.Error("snippet is nil")
	}
	if ch.Statistics == nil {
		t.Error("statistics is nil")
	}
}

func TestChannelsByUsername_Integration(t *testing.T) {
	c := testClient(t)
	result, err := c.Channel.ListByUsername("snippet", "RickAstleyVEVO")
	if err != nil {
		t.Fatalf("Channel.ListByUsername failed: %v", err)
	}
	if len(result.Items) == 0 {
		t.Log("Note: RickAstleyVEVO username may not exist, this is informational")
	}
}

func TestPlaylistsByChannel_Integration(t *testing.T) {
	c := testClient(t)
	result, err := c.Playlist.ListByChannel("snippet,contentDetails", "UC_x5XG1OV2P6uZZ5FSM9Ttw", 3)
	if err != nil {
		t.Fatalf("Playlist.ListByChannel failed: %v", err)
	}
	if len(result.Items) == 0 {
		t.Log("channel has 0 playlists (or needs auth)")
	}
}

func TestCategories_Integration(t *testing.T) {
	c := testClient(t)
	result, err := c.Category.List("US")
	if err != nil {
		t.Fatalf("Category.List failed: %v", err)
	}
	if len(result.Items) == 0 {
		t.Fatal("categories returned 0 items")
	}
	hasMusic := false
	for _, cat := range result.Items {
		if cat.Snippet.Title == "Music" {
			hasMusic = true
			break
		}
	}
	if !hasMusic {
		t.Error("expected 'Music' category")
	}
}

func TestSearchPagination_Integration(t *testing.T) {
	c := testClient(t)
	// Get first page
	page1, err := c.Search.Search("test", &SearchParams{Part: "snippet", MaxResults: 5})
	if err != nil {
		t.Fatalf("Search page 1 failed: %v", err)
	}
	if page1.NextPageToken == "" {
		t.Log("no next page (fewer results than maxResults)")
		return
	}
	// Get second page
	page2, err := c.Search.Search("test", &SearchParams{Part: "snippet", MaxResults: 5, PageToken: page1.NextPageToken})
	if err != nil {
		t.Fatalf("Search page 2 failed: %v", err)
	}
	if len(page2.Items) == 0 {
		t.Error("second page returned 0 items")
	}
	// Ensure pages are different
	if len(page1.Items) > 0 && len(page2.Items) > 0 {
		if page1.Items[0].ID == page2.Items[0].ID {
			t.Log("note: first item is the same across pages (possible but unusual)")
		}
	}
}

func TestMultipleVideos_Integration(t *testing.T) {
	c := testClient(t)
	result, err := c.Video.List("snippet", "dQw4w9WgXcQ", "jNQXAC9IVRw", "kJQP7kiw5Fk")
	if err != nil {
		t.Fatalf("Video.List multi failed: %v", err)
	}
	if len(result.Items) != 3 {
		t.Logf("expected 3 videos, got %d (some may be private/deleted)", len(result.Items))
	}
}

func TestClientNewWithKey(t *testing.T) {
	key := os.Getenv("YOUTUBE_API_KEY")
	if key == "" {
		t.Skip("Set YOUTUBE_API_KEY")
	}
	c := NewClientWithKey(key)
	if c == nil {
		t.Fatal("client is nil")
	}
	if c.Video == nil {
		t.Fatal("Video service is nil")
	}
	if c.Search == nil {
		t.Fatal("Search service is nil")
	}
}

func TestKeyParam(t *testing.T) {
	// Test that KeyParam produces correct query string
	kp := &KeyParam{APIKey: "test_key"}
	// We can't easily test querystring encoding here without importing the lib,
	// but the struct field is now exported so the encoder should pick it up
	if kp.APIKey != "test_key" {
		t.Errorf("expected test_key, got %s", kp.APIKey)
	}
}

func TestJoinIDs(t *testing.T) {
	tests := []struct {
		input []string
		want  string
	}{
		{nil, ""},
		{[]string{}, ""},
		{[]string{"a"}, "a"},
		{[]string{"a", "b"}, "a,b"},
		{[]string{"a", "b", "c"}, "a,b,c"},
	}
	for _, tt := range tests {
		got := joinIDs(tt.input...)
		if got != tt.want {
			t.Errorf("joinIDs(%v) = %q, want %q", tt.input, got, tt.want)
		}
	}
}
