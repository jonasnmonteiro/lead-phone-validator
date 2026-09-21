package telecom

import (
	"regexp"
	"strings"

	"leadphone-validator/internal/models"
)

var digitRegex = regexp.MustCompile(`\D`)

func SanitizeDigits(n string) string {
	return digitRegex.ReplaceAllString(n, "")
}

func FixNinthDigit(input string) models.NinthDigitResponse {
	raw := SanitizeDigits(input)

	if len(raw) == 8 {
		firstDigit := raw[0]
		if firstDigit >= '6' && firstDigit <= '9' {
			fixed := "9" + raw
			return models.NinthDigitResponse{
				Input:      input,
				Normalized: fixed,
				Changed:    true,
				IsMobile:   true,
				IsLandline: false,
				IsValid:    true,
				Message:    "Prepend ninth digit 9 to legacy 8-digit mobile number.",
			}
		}
		if firstDigit >= '2' && firstDigit <= '5' {
			return models.NinthDigitResponse{
				Input:      input,
				Normalized: raw,
				Changed:    false,
				IsMobile:   false,
				IsLandline: true,
				IsValid:    true,
				Message:    "Standard 8-digit landline number. No 9th digit required.",
			}
		}
	}

	if len(raw) == 10 {
		ddd := raw[0:2]
		number := raw[2:]
		firstDigit := number[0]

		if firstDigit >= '6' && firstDigit <= '9' {
			fixed := ddd + "9" + number
			return models.NinthDigitResponse{
				Input:      input,
				Normalized: fixed,
				Changed:    true,
				IsMobile:   true,
				IsLandline: false,
				IsValid:    true,
				Message:    "Inserted ninth digit 9 after area code.",
			}
		}
		if firstDigit >= '2' && firstDigit <= '5' {
			return models.NinthDigitResponse{
				Input:      input,
				Normalized: raw,
				Changed:    false,
				IsMobile:   false,
				IsLandline: true,
				IsValid:    true,
				Message:    "Standard 10-digit landline number. No 9th digit required.",
			}
		}
	}

	if len(raw) == 12 && strings.HasPrefix(raw, "55") {
		ddd := raw[2:4]
		number := raw[4:]
		firstDigit := number[0]

		if firstDigit >= '6' && firstDigit <= '9' {
			fixed := "55" + ddd + "9" + number
			return models.NinthDigitResponse{
				Input:      input,
				Normalized: fixed,
				Changed:    true,
				IsMobile:   true,
				IsLandline: false,
				IsValid:    true,
				Message:    "Inserted ninth digit 9 into full Brazilian international number.",
			}
		}
		if firstDigit >= '2' && firstDigit <= '5' {
			return models.NinthDigitResponse{
				Input:      input,
				Normalized: raw,
				Changed:    false,
				IsMobile:   false,
				IsLandline: true,
				IsValid:    true,
				Message:    "Standard 12-digit international landline number.",
			}
		}
	}

	if len(raw) == 11 {
		thirdDigit := raw[2]
		if thirdDigit == '9' {
			return models.NinthDigitResponse{
				Input:      input,
				Normalized: raw,
				Changed:    false,
				IsMobile:   true,
				IsLandline: false,
				IsValid:    true,
				Message:    "Already compliant 11-digit mobile format.",
			}
		}
	}

	if len(raw) == 13 && strings.HasPrefix(raw, "55") {
		fifthDigit := raw[4]
		if fifthDigit == '9' {
			return models.NinthDigitResponse{
				Input:      input,
				Normalized: raw,
				Changed:    false,
				IsMobile:   true,
				IsLandline: false,
				IsValid:    true,
				Message:    "Already compliant 13-digit international mobile format.",
			}
		}
	}

	return models.NinthDigitResponse{
		Input:      input,
		Normalized: raw,
		Changed:    false,
		IsMobile:   false,
		IsLandline: false,
		IsValid:    false,
		Message:    "Unrecognized digit length or structure for ninth digit conversion.",
	}
}
