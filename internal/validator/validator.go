package validator

import (
	"regexp"
	"strings"

	"github.com/nyaruka/phonenumbers"
	"leadphone-validator/internal/models"
)

var nonDigitRegex = regexp.MustCompile(`\D`)

func CleanNumber(n string) string {
	return nonDigitRegex.ReplaceAllString(n, "")
}

func NumberTypeName(t phonenumbers.PhoneNumberType) string {
	switch t {
	case phonenumbers.FIXED_LINE:
		return "FIXED_LINE"
	case phonenumbers.MOBILE:
		return "MOBILE"
	case phonenumbers.FIXED_LINE_OR_MOBILE:
		return "FIXED_LINE_OR_MOBILE"
	case phonenumbers.TOLL_FREE:
		return "TOLL_FREE"
	case phonenumbers.PREMIUM_RATE:
		return "PREMIUM_RATE"
	case phonenumbers.SHARED_COST:
		return "SHARED_COST"
	case phonenumbers.VOIP:
		return "VOIP"
	case phonenumbers.PERSONAL_NUMBER:
		return "PERSONAL_NUMBER"
	case phonenumbers.PAGER:
		return "PAGER"
	case phonenumbers.UAN:
		return "UAN"
	case phonenumbers.VOICEMAIL:
		return "VOICEMAIL"
	default:
		return "UNKNOWN"
	}
}

func buildValidResponse(rawPhone string, countryInput string, num *phonenumbers.PhoneNumber) models.ValidationResponse {
	countryCode := num.GetCountryCode()
	nationalNumber := num.GetNationalNumber()
	region := phonenumbers.GetRegionCodeForNumber(num)
	numType := NumberTypeName(phonenumbers.GetNumberType(num))
	isValid := phonenumbers.IsValidNumber(num)
	isPossible := phonenumbers.IsPossibleNumber(num)

	e164 := phonenumbers.Format(num, phonenumbers.E164)
	international := phonenumbers.Format(num, phonenumbers.INTERNATIONAL)
	national := phonenumbers.Format(num, phonenumbers.NATIONAL)
	rfc3966 := phonenumbers.Format(num, phonenumbers.RFC3966)
	waNum := strings.TrimPrefix(e164, "+")
	waURL := "https://wa.me/" + waNum

	var carrierName *string
	if c, err := phonenumbers.GetCarrierForNumber(num, "en"); err == nil && c != "" {
		carrierName = &c
	}

	var locationName *string
	if loc, err := phonenumbers.GetGeocodingForNumber(num, "en"); err == nil && loc != "" {
		locationName = &loc
	}

	timezones, _ := phonenumbers.GetTimezonesForNumber(num)
	if timezones == nil {
		timezones = []string{}
	}

	var countryPtr *string
	if countryInput != "" {
		countryPtr = &countryInput
	}

	return models.ValidationResponse{
		Input: models.InputData{
			Raw:     rawPhone,
			Country: countryPtr,
		},
		Parsed: models.ParsedData{
			IsValid:        isValid,
			IsPossible:     isPossible,
			Type:           &numType,
			Region:         &region,
			CountryCode:    &countryCode,
			NationalNumber: &nationalNumber,
		},
		Metadata: &models.MetadataData{
			Carrier:   carrierName,
			Location:  locationName,
			Timezones: timezones,
		},
		Formatted: models.FormattedData{
			E164:          &e164,
			International: &international,
			National:      &national,
			RFC3966:       &rfc3966,
			WhatsApp:      &waNum,
			WhatsAppURL:   &waURL,
		},
	}
}

func buildInvalidResponse(rawPhone string, countryInput string, reason string, needsDDD bool, needsCountry bool) models.ValidationResponse {
	var countryPtr *string
	if countryInput != "" {
		countryPtr = &countryInput
	}

	var reasonPtr *string
	if reason != "" {
		reasonPtr = &reason
	}

	var dddPtr *bool
	if needsDDD {
		dddPtr = &needsDDD
	}

	var cntryPtr *bool
	if needsCountry {
		cntryPtr = &needsCountry
	}

	return models.ValidationResponse{
		Input: models.InputData{
			Raw:     rawPhone,
			Country: countryPtr,
		},
		Parsed: models.ParsedData{
			IsValid:      false,
			IsPossible:   false,
			Type:         nil,
			Region:       nil,
			NeedsDDD:     dddPtr,
			NeedsCountry: cntryPtr,
			Reason:       reasonPtr,
		},
		Metadata:  nil,
		Formatted: models.FormattedData{},
	}
}

func Validate(req models.ValidateRequest) models.ValidationResponse {
	phone := strings.TrimSpace(req.Phone)
	country := strings.TrimSpace(req.Country)
	raw := CleanNumber(phone)

	if country != "" {
		if parsed, err := phonenumbers.Parse(phone, strings.ToUpper(country)); err == nil {
			if phonenumbers.IsValidNumber(parsed) {
				return buildValidResponse(phone, country, parsed)
			}
		}
	}

	if !strings.HasPrefix(phone, "+") && len(raw) >= 10 {
		if parsed, err := phonenumbers.Parse("+"+raw, ""); err == nil {
			if phonenumbers.IsValidNumber(parsed) {
				return buildValidResponse(phone, country, parsed)
			}
		}
	}

	if strings.HasPrefix(phone, "+") {
		if parsed, err := phonenumbers.Parse(phone, ""); err == nil {
			if phonenumbers.IsValidNumber(parsed) {
				return buildValidResponse(phone, country, parsed)
			}
		}
		return buildInvalidResponse(phone, country, "Invalid number.", false, false)
	}

	if len(raw) == 8 || len(raw) == 9 {
		return buildInvalidResponse(phone, country, "Brazilian number without area code.", true, false)
	}

	if len(raw) == 10 {
		if parsed, err := phonenumbers.Parse(phone, "US"); err == nil {
			if phonenumbers.IsValidNumber(parsed) {
				return buildValidResponse(phone, country, parsed)
			}
		}
		return buildInvalidResponse(phone, country, "10-digit number does not clearly correspond to a country.", false, true)
	}

	if strings.HasPrefix(raw, "55") {
		if parsed, err := phonenumbers.Parse("+"+raw, "BR"); err == nil {
			if phonenumbers.IsValidNumber(parsed) {
				return buildValidResponse(phone, country, parsed)
			}
		}
		return buildInvalidResponse(phone, country, "Number starting with 55, but invalid. Missing area code?", false, false)
	}

	if len(raw) == 11 {
		if parsed, err := phonenumbers.Parse(raw, "BR"); err == nil {
			if phonenumbers.IsValidNumber(parsed) {
				return buildValidResponse(phone, country, parsed)
			}
		}
		return buildInvalidResponse(phone, country, "Invalid 11-digit format for BR.", true, false)
	}

	if len(raw) >= 12 {
		if parsed, err := phonenumbers.Parse("+"+raw, ""); err == nil {
			if phonenumbers.IsValidNumber(parsed) {
				return buildValidResponse(phone, country, parsed)
			}
		}
		return buildInvalidResponse(phone, country, "Invalid international number.", false, false)
	}

	return buildInvalidResponse(phone, country, "Invalid number.", false, false)
}

func ValidateBatch(requests []models.ValidateRequest) models.BatchValidationResponse {
	n := len(requests)
	results := make([]models.ValidationResponse, n)

	if n == 0 {
		return models.BatchValidationResponse{
			Total:   0,
			Valid:   0,
			Invalid: 0,
			Results: results,
		}
	}

	type indexedResult struct {
		index int
		resp  models.ValidationResponse
	}

	resChan := make(chan indexedResult, n)
	for i, req := range requests {
		go func(idx int, r models.ValidateRequest) {
			resChan <- indexedResult{
				index: idx,
				resp:  Validate(r),
			}
		}(i, req)
	}

	validCount := 0
	invalidCount := 0

	for i := 0; i < n; i++ {
		res := <-resChan
		results[res.index] = res.resp
		if res.resp.Parsed.IsValid {
			validCount++
		} else {
			invalidCount++
		}
	}

	return models.BatchValidationResponse{
		Total:   n,
		Valid:   validCount,
		Invalid: invalidCount,
		Results: results,
	}
}
