package draven

import "time"

// Video https://developers.facebook.com/docs/graph-api/reference/video/
type Video struct {
	AdBreaks               []int32
	BackdatedTime          time.Time
	ContentTags            []string
	CreatedTime            time.Time
	ID                     string
	CustomLabels           []string
	Description            string
	EmbedHTML              string
	Embeddable             bool
	Event                  Event
	Format                 VideoFormat
	Icon                   string
	IsCrossPostVideo       bool
	IsCrossPostingElegible bool
	IsInstagramElegible    bool
	IsReferenceOnly        bool
	Lenght                 float32
	PermalinkURL           string
	Picture                string
	Privacy                Privacy
	Published              bool
	ScheduledPublishTime   time.Time
	Source                 string
	Status                 VideoStatus
	Title                  string
	UniversalVideoID       string
	UpdatedTime            time.Time
	// BackdatedTimeGranularity time.Time
	// ContentCategory int32
	// From                   User?
	// LiveStatus int
	// MonetizationStatus int
	// Place Place
}

// VideoStatus https://developers.facebook.com/docs/graph-api/reference/video-status/
type VideoStatus struct {
	ProcessingProgress int32
	VideoStatus        int
}

// VideoFormat https://developers.facebook.com/docs/graph-api/reference/video-format/
type VideoFormat struct {
	EmbedHTML string
	Filter    string
	Height    int32
	Picture   string
	Width     int32
}

// Privacy https://developers.facebook.com/docs/graph-api/reference/privacy/
type Privacy struct {
	Description string
	Networks    string
	// Allow
	// Deny
	// Friends int
	// Value int
}
