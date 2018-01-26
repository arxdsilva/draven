package draven

import "time"

// Event https://developers.facebook.com/docs/graph-api/reference/event/
type Event struct {
	ID                   string
	AttendingCount       int32
	CanGuestsInvite      bool
	CanViewerPost        bool
	Category             string
	Cover                CoverPhoto
	DeclinedCount        int32
	Description          string
	EndTime              string
	GuestListEnabled     bool
	InterestedCount      int32
	IsCanceled           bool
	IsDraft              bool
	IsPageOwned          bool
	IsViewerAdmin        bool
	MaybeCount           int32
	Name                 string
	NoReplyCount         int32
	ScheduledPublishTime string
	StartTime            string
	TicketURI            string
	TicketingPrivacyURI  string
	TicketingTermsURI    string
	UpdatedTime          time.Time
	// EventTimes ChildEvent
	// Owner
	// ParentGroup Group
	// Place Place
	// Timezone int
	// Type int
}
