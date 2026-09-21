package telecom

import (
	"net/url"
	"strings"

	"leadphone-validator/internal/models"
	"leadphone-validator/internal/validator"
)

func BuildWhatsAppLink(req models.WhatsAppLinkRequest) models.WhatsAppLinkResponse {
	valResp := validator.Validate(models.ValidateRequest{
		Phone: req.Phone,
	})

	digits := ""
	if valResp.Parsed.IsValid && valResp.Formatted.WhatsApp != nil {
		digits = *valResp.Formatted.WhatsApp
	} else {
		digits = SanitizeDigits(req.Phone)
	}

	if digits == "" {
		return models.WhatsAppLinkResponse{
			Phone:   req.Phone,
			E164:    "",
			Message: req.Message,
			URL:     "",
			IsValid: false,
		}
	}

	e164Formatted := "+" + digits
	if valResp.Formatted.E164 != nil {
		e164Formatted = *valResp.Formatted.E164
	}

	waURL := "https://wa.me/" + digits
	if strings.TrimSpace(req.Message) != "" {
		encodedMsg := url.QueryEscape(strings.TrimSpace(req.Message))
		waURL = waURL + "?text=" + encodedMsg
	}

	return models.WhatsAppLinkResponse{
		Phone:   req.Phone,
		E164:    e164Formatted,
		Message: req.Message,
		URL:     waURL,
		IsValid: valResp.Parsed.IsValid,
	}
}
