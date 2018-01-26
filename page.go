package draven

import "time"

// fazer um metodo do request
// para o cara receber raw do
// body da requisicao ao inves
// dessas structs

// Page is a facebook Page
// Page https://developers.facebook.com/docs/graph-api/reference/page/
type Page struct {
	ID                              string
	About                           string
	AccessToken                     string
	AdCampaign                      AdCampaign
	Affiliation                     string
	AppID                           string
	AppLinks                        AppLinks
	ArtistsWeLike                   string
	Attire                          string
	Awards                          string
	BandInterests                   string
	BestPage                        *Page
	Bio                             string
	Birthday                        string
	BookingAgent                    string
	Built                           string
	CanCheckin                      bool
	CanPost                         bool
	Category                        string
	CategoryList                    []PageCategory
	Checkins                        int32
	CompanyOverview                 string
	ContactAddress                  MailingAddress
	CountryPageLikes                int32
	Cover                           CoverPhoto
	CulinaryTeam                    string
	CurrentLocation                 string
	Description                     string
	DescriptionHTML                 string
	DirectedBy                      string
	DisplaySubtext                  string
	DisplayedMessageResponseTime    string
	Emails                          []string
	Engagement                      Engagement
	FanCount                        int32
	FeaturedVideo                   Video
	Features                        string
	FoodStyles                      []string
	Founded                         string
	GeneralInfo                     string
	GeneralManager                  string
	Genre                           string
	GlobalBrandPageName             string
	GlobalBrandRootID               string
	HasAddedApp                     bool
	HasWhatsappNumber               bool
	Hometown                        string
	Hours                           map[string]string
	Impressum                       string
	Influences                      string
	IsAlwaysOpen                    bool
	IsChain                         bool
	IsCommunityPage                 bool
	IsEligibleForBrandedContent     bool
	IsOwned                         bool
	IsPermanentlyClosed             bool
	IsPublished                     bool
	IsUnclaimed                     bool
	IsWebhooksSubscribed            bool
	LeadgenHasCRMIntegration        bool
	LeadgenHasFatPingCRMIntegration bool
	LeadgenTosAcceptanceTime        time.Time
	LeadgenTosAccepted              bool
	LeadgenTosAcceptingUser         User
	Link                            string
	Location                        Location
	Members                         string
	MerchantID                      string
	MerchantReviewStatus            int
	MessengerAdsQuickRepliesType    int
	Mission                         string
	MPG                             string
	Name                            string
	NameWithLocationDescriptor      string
	Network                         string
	NewLikeCount                    int32
	OfferEligible                   bool
	OverallStarRating               float32
	PageToken                       string
	ParentPage                      *Page
	// MessengerAdsDefaultPageWelcome
	// LeadgenFormPreviewDetails LeadgenFormPreviewDetails
	// InstantArticlesReviewStatus int
	// InstagramBusinessAccount ShadowIGUser
	// Business
	// ConnectedInstagramAccount
	// Context OpenGraphContext
}

// AppLinks https://developers.facebook.com/docs/graph-api/reference/app-links/
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
}

// AdCampaign https://developers.facebook.com/docs/marketing-api/reference/ad-campaign/
type AdCampaign struct {
	ID        string
	AccountID string
	// to be continued ...
}

type MailingAddress struct {
	ID         string
	City       string
	Contry     string
	PostalCode string
	Region     string
	Street1    string
	Street2    string
	CityPage   *Page
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
