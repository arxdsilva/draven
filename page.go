package draven

type Page struct {
	ID            string
	About         string
	AccessToken   string
	AdCampaign    AdCampaign
	Affiliation   string
	AppID         string
	AppLinks      AppLinks
	ArtistsWeLike string
	Attire        string
	Awards        string
	BandInterests string
	// BestPage      Page
	Bio          string
	Birthday     string
	BookingAgent string
	Built        string
	// Business
	CanCheckin      bool
	CanPost         bool
	Category        string
	CategoryList    []PageCategory
	Checkins        int32
	CompanyOverview string
	// ConnectedInstagramAccount
	ContactAddress MailingAddress
	// Context OpenGraphContext
	CountryPageLikes             int32
	Cover                        CoverPhoto
	CulinaryTeam                 string
	CurrentLocation              string
	Description                  string
	DescriptionHTML              string
	DirectedBy                   string
	DisplaySubtext               string
	DisplayedMessageResponseTime string
	Emails                       []string
	Engagement                   Engagement
	FanCount                     int32
	// to be continued...
	// https://developers.facebook.com/docs/graph-api/reference/page/
}

type AppLinks struct {
	ID string
	// Android          []AndroidAppLink
	// IOS              []IosAppLink
	// IPad             []IosAppLink
	// IPhone           []IosAppLink
	// Web              []WebAppLink
	// Windows          []WindowsAppLink
	// WindowsPhone     []WindowsAppLink
	// WindowsUniversal []WindowsAppLink
	// to be continued ...
	// https://developers.facebook.com/docs/graph-api/reference/app-links/
}
type AdCampaign struct {
	ID        string
	AccountID string
	// to be continued ...
	// https://developers.facebook.com/docs/marketing-api/reference/ad-campaign/
}

type MailingAddress struct {
	ID   string
	City string
	// CityPage   Page
	Contry     string
	PostalCode string
	Region     string
	Street1    string
	Street2    string
}

type PageCategory struct {
	ID               string
	APIEnum          string
	FBPageCategories []PageCategory
	Name             string
}

type Device struct {
	Hardware string
	OS       string
}

type Engagement struct {
	Count                     int32
	CountString               string
	CountStringWithLike       string
	CountStringWithoutLike    string
	SocialSentence            string
	SocialSentenceWithLike    string
	SocialSentenceWitouthLike string
}
