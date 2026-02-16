package youtube

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/madr/backend/pkg/logger"
)

const (
	YouTubeAPIKey    = "AIzaSyAnXmKQn5nka20et5qkyptOySfxmB6h5BY"
	YouTubeChannelID = "UCPTIg2Lw3V81c74awdSBqnQ"
	YouTubeAPIURL    = "https://www.googleapis.com/youtube/v3/search"
)

// Video represents a YouTube video
type Video struct {
	VideoID      string    `json:"video_id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	PublishedAt  time.Time `json:"published_at"`
	ThumbnailURL string    `json:"thumbnail_url"`
	ChannelTitle string    `json:"channel_title"`
}

// YouTubeSearchResponse represents the YouTube API search response
type YouTubeSearchResponse struct {
	Items []struct {
		ID struct {
			VideoID string `json:"videoId"`
		} `json:"id"`
		Snippet struct {
			Title        string `json:"title"`
			Description  string `json:"description"`
			PublishedAt string `json:"publishedAt"`
			Thumbnails  struct {
				High struct {
					URL string `json:"url"`
				} `json:"high"`
			} `json:"thumbnails"`
			ChannelTitle string `json:"channelTitle"`
		} `json:"snippet"`
	} `json:"items"`
}

// Service handles YouTube API interactions
type Service interface {
	GetRecentVideos(days int) ([]Video, error)
}

type service struct{}

// NewService creates a new YouTube service
func NewService() Service {
	return &service{}
}

// GetRecentVideos fetches videos uploaded in the last N days
func (s *service) GetRecentVideos(days int) ([]Video, error) {
	// Calculate publishedAfter date (30 days ago)
	publishedAfter := time.Now().AddDate(0, 0, -days).Format(time.RFC3339)

	// Build request URL with proper encoding
	reqURL, err := url.Parse(YouTubeAPIURL)
	if err != nil {
		return nil, fmt.Errorf("invalid API URL: %w", err)
	}

	params := url.Values{}
	params.Add("part", "snippet")
	params.Add("channelId", YouTubeChannelID)
	params.Add("order", "date")
	params.Add("maxResults", "50")
	params.Add("type", "video")
	params.Add("key", YouTubeAPIKey)
	params.Add("publishedAfter", publishedAfter)
	reqURL.RawQuery = params.Encode()

	fullURL := reqURL.String()
	logger.Debug().Str("url", fullURL).Msg("Fetching YouTube videos")

	// Make HTTP request
	resp, err := http.Get(fullURL)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to fetch YouTube videos")
		return nil, fmt.Errorf("failed to fetch YouTube videos: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		logger.Error().
			Int("status", resp.StatusCode).
			Str("body", string(body)).
			Msg("YouTube API returned error")
		
		// Try to parse error response from YouTube API
		var errorResp struct {
			Error struct {
				Code    int    `json:"code"`
				Message string `json:"message"`
				Errors  []struct {
					Domain  string `json:"domain"`
					Reason  string `json:"reason"`
					Message string `json:"message"`
				} `json:"errors"`
			} `json:"error"`
		}
		if err := json.Unmarshal(body, &errorResp); err == nil && errorResp.Error.Message != "" {
			return nil, fmt.Errorf("YouTube API error (%d): %s", errorResp.Error.Code, errorResp.Error.Message)
		}
		
		return nil, fmt.Errorf("YouTube API error: status %d, body: %s", resp.StatusCode, string(body))
	}

	// Parse response
	var apiResp YouTubeSearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		logger.Error().Err(err).Msg("Failed to parse YouTube API response")
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	// Convert to Video structs
	videos := make([]Video, 0, len(apiResp.Items))
	for _, item := range apiResp.Items {
		publishedAt, err := time.Parse(time.RFC3339, item.Snippet.PublishedAt)
		if err != nil {
			logger.Warn().Err(err).Str("date", item.Snippet.PublishedAt).Msg("Failed to parse published date")
			continue
		}

		videos = append(videos, Video{
			VideoID:      item.ID.VideoID,
			Title:        item.Snippet.Title,
			Description:  item.Snippet.Description,
			PublishedAt:  publishedAt,
			ThumbnailURL: item.Snippet.Thumbnails.High.URL,
			ChannelTitle: item.Snippet.ChannelTitle,
		})
	}

	logger.Info().Int("count", len(videos)).Msg("Successfully fetched YouTube videos")
	return videos, nil
}
