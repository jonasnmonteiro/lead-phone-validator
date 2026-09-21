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
