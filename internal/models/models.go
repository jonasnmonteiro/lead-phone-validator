package models

type ValidateRequest struct {
	Phone   string `json:"phone"`
	Country string `json:"country,omitempty"`
	Raw     string `json:"raw,omitempty"`
}

type BatchValidateRequest struct {
	Phones []ValidateRequest `json:"phones"`
}

type InputData struct {
	Raw     string  `json:"raw"`
	Country *string `json:"country"`
}

type ParsedData struct {
	IsValid        bool    `json:"isValid"`
	IsPossible     bool    `json:"isPossible"`
	Type           *string `json:"type"`
	Region         *string `json:"region"`
	CountryCode    *int32  `json:"countryCode,omitempty"`
	NationalNumber *uint64 `json:"nationalNumber,omitempty"`
	NeedsDDD       *bool   `json:"needsDDD,omitempty"`
	NeedsCountry   *bool   `json:"needsCountry,omitempty"`
	Reason         *string `json:"reason,omitempty"`
}

type MetadataData struct {
	Carrier   *string  `json:"carrier"`
	Location  *string  `json:"location"`
	Timezones []string `json:"timezones"`
}

type FormattedData struct {
	E164          *string `json:"e164,omitempty"`
	International *string `json:"international,omitempty"`
	National      *string `json:"national,omitempty"`
	RFC3966       *string `json:"rfc3966,omitempty"`
	WhatsApp      *string `json:"whatsapp,omitempty"`
	WhatsAppURL   *string `json:"whatsappUrl,omitempty"`
}

type ValidationResponse struct {
	Input     InputData     `json:"input"`
	Parsed    ParsedData    `json:"parsed"`
	Metadata  *MetadataData `json:"metadata"`
	Formatted FormattedData `json:"formatted"`
}

type BatchValidationResponse struct {
	Total   int                  `json:"total"`
	Valid   int                  `json:"valid"`
	Invalid int                  `json:"invalid"`
	Results []ValidationResponse `json:"results"`
}

type HealthResponse struct {
	Status  string `json:"status"`
	Service string `json:"service"`
	Version string `json:"version"`
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
}

type DDDInfo struct {
	Code        string   `json:"code"`
	State       string   `json:"state"`
	StateName   string   `json:"stateName"`
	Region      string   `json:"region"`
	MajorCities []string `json:"majorCities"`
	Timezone    string   `json:"timezone"`
}

type DDDResponse struct {
	Query   string    `json:"query"`
	Found   bool      `json:"found"`
	Results []DDDInfo `json:"results"`
}

type DDIInfo struct {
	CallingCode      string   `json:"callingCode"`
	CountryName      string   `json:"countryName"`
	ISO2             string   `json:"iso2"`
	ISO3             string   `json:"iso3"`
	IDDPrefix        string   `json:"iddPrefix"`
	EmergencyNumbers []string `json:"emergencyNumbers"`
	TimezoneHint     string   `json:"timezoneHint"`
}

type DDIResponse struct {
	Query   string    `json:"query"`
	Found   bool      `json:"found"`
	Results []DDIInfo `json:"results"`
}

type NinthDigitRequest struct {
	Phone string `json:"phone"`
}

type NinthDigitResponse struct {
	Input      string `json:"input"`
	Normalized string `json:"normalized"`
	Changed    bool   `json:"changed"`
	IsMobile   bool   `json:"isMobile"`
	IsLandline bool   `json:"isLandline"`
	IsValid    bool   `json:"isValid"`
	Message    string `json:"message"`
}

type WhatsAppLinkRequest struct {
	Phone   string `json:"phone"`
	Message string `json:"message,omitempty"`
}

type WhatsAppLinkResponse struct {
	Phone       string `json:"phone"`
	E164        string `json:"e164"`
	Message     string `json:"message,omitempty"`
	URL         string `json:"url"`
	IsValid     bool   `json:"isValid"`
}

type PatternCheckRequest struct {
	Phone string `json:"phone"`
}

type PatternCheckResponse struct {
	Phone       string   `json:"phone"`
	IsSuspicious bool    `json:"isSuspicious"`
	Flags       []string `json:"flags"`
	Score       int      `json:"score"`
	Description string   `json:"description"`
}
