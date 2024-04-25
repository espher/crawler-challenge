package models

type Vehicle struct {
	CreatedAt        int64 `gorm:"<-:create"`
	LastSeenAt       int64
	Source           string `gorm:"primaryKey" json:"source"`
	VIN              string `gorm:"primaryKey" json:"vin"`
	Year             string `json:"year"`
	Make             string `json:"make"`
	Model            string `json:"model"`
	Trim             string `json:"trim"`
	BodyStyle        string `json:"bodyStyle"`
	Mileage          string `json:"mileage"`
	OfferPrice       string `json:"offerPrice"`
	Certified        bool   `json:"certified"`
	ExteriorColor    string `json:"exteriorColor"`
	InteriorColor    string `json:"interiorColor"`
	Transmission     string `json:"transmission"`
	Engine           string `json:"engine"`
	DetailPage       string `json:"detailPage"`
	Description      string `json:"description"`
	StockNumber      string `json:"stockNumber"`
	OneOwner         bool   `json:"oneOwner"`
	DealerZipCode    string `json:"dealerZipCode"`
	DealerName       string `json:"dealerName"`
	Condition        string `json:"type"`
	InstalledOptions string `json:"installed_options"`
}
