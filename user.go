package draven

import "time"

type User struct {
	About                       string
	ID                          string
	Address                     Location
	AgeRange                    AgeRange
	Birthday                    string
	CanReviewMeasurementRequest bool
	Context                     UserContext
	CoverPhoto                  CoverPhoto
	Currency                    Currency
	Devices                     []Device
	Education                   []EducationExperience
	Email                       string
	EmployeeNumber              string
	FavoriteAthletes            []Experience
	FavoriteTeams               []Experience
	FirstName                   string
	Gender                      string
	Hometown                    Page
	InspirationalPeople         []Experience
	// InstallType Enum
	Installed        bool
	InterestedIn     []string
	IsPaymentEnabled bool
	IsSharedLogin    bool
	IsVerified       bool
	// Labels []PageLabel
	Languages []Experience
	// LastAdReferral
	LastName   string
	Link       string
	Locale     string
	Location   Page
	MeetingFor []string
	MiddleName string
	Name       string
	NameFormat string
	// PaymentPricePoints PaymentPricePoints
	Political          string
	ProfilePic         string
	PublicKey          string
	Quotes             string
	RelationshipStatus string
	Religion           string
	// SecuritySettings SecuritySettings
	// SharedLoginUpgradeRequiredBy
	ShortName string
	// SignificantOther User
	Sports           []Experience
	TestGroup        int32
	ThirdPartyID     string
	TimeZone         float32
	TokenForBusiness string
	UpdateTime       time.Time
	Verified         bool
	// VideoUploadLimits VideoUploadLimits
	ViewerCanSendGift bool
	Website           string
	// Work []WorkExperience
	// https://developers.facebook.com/docs/graph-api/reference/user/
}

type CoverPhoto struct {
	ID      string
	OffsetX float32
	OffsetY float32
	Source  string
}
type Currency struct {
	CurrencyOffset     int32
	USDExchange        float32
	USDExchangeInverse float32
	UserCurrency       string
}

type UserContext struct {
	ID string
}

type AgeRange struct {
	Max int32
	Min int32
}

type Location struct {
	City        string
	CityID      int32
	Country     string
	CountryCode string
	Latitude    float32
	// LocatedIn ???
	Longitude float32
	Name      string
	Region    string
	RegionID  int32
	State     string
	Street    string
	ZIP       string
}
