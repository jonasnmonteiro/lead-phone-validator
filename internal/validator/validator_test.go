package validator

import (
	"testing"

	"leadphone-validator/internal/models"
)

func TestValidateBrazilianWithoutDDD(t *testing.T) {
	resp := Validate(models.ValidateRequest{
		Phone: "987654321",
	})

	if resp.Parsed.IsValid {
		t.Fatalf("expected invalid for 9-digit number without DDD")
	}

	if resp.Parsed.NeedsDDD == nil || !*resp.Parsed.NeedsDDD {
		t.Fatalf("expected needsDDD to be true")
	}
}

func TestValidateBrazilianWithDDD(t *testing.T) {
	resp := Validate(models.ValidateRequest{
		Phone: "11987654321",
	})

	if !resp.Parsed.IsValid {
		t.Fatalf("expected valid for 11-digit Brazilian number")
	}

	if resp.Parsed.Region == nil || *resp.Parsed.Region != "BR" {
		t.Fatalf("expected region BR, got %v", resp.Parsed.Region)
	}

	if resp.Formatted.E164 == nil || *resp.Formatted.E164 != "+5511987654321" {
		t.Fatalf("expected E164 +5511987654321, got %v", resp.Formatted.E164)
	}

	if resp.Formatted.WhatsAppURL == nil || *resp.Formatted.WhatsAppURL != "https://wa.me/5511987654321" {
		t.Fatalf("expected WhatsApp URL https://wa.me/5511987654321, got %v", resp.Formatted.WhatsAppURL)
	}
}

func TestValidateSingaporeWithoutPlus(t *testing.T) {
	resp := Validate(models.ValidateRequest{
		Phone: "6581234567",
	})

	if !resp.Parsed.IsValid {
		t.Fatalf("expected valid for Singapore number without plus")
	}

	if resp.Parsed.Region == nil || *resp.Parsed.Region != "SG" {
		t.Fatalf("expected region SG, got %v", resp.Parsed.Region)
	}
}

func TestValidateUSPhone(t *testing.T) {
	resp := Validate(models.ValidateRequest{
		Phone: "2025550125",
	})

	if !resp.Parsed.IsValid {
		t.Fatalf("expected valid for US phone")
	}

	if resp.Parsed.Region == nil || *resp.Parsed.Region != "US" {
		t.Fatalf("expected region US, got %v", resp.Parsed.Region)
	}
}

func TestValidateBatch(t *testing.T) {
	requests := []models.ValidateRequest{
		{Phone: "11987654321"},
		{Phone: "6581234567"},
		{Phone: "987654321"},
	}

	batchResp := ValidateBatch(requests)

	if batchResp.Total != 3 {
		t.Fatalf("expected total 3, got %d", batchResp.Total)
	}

	if batchResp.Valid != 2 {
		t.Fatalf("expected valid 2, got %d", batchResp.Valid)
	}

	if batchResp.Invalid != 1 {
		t.Fatalf("expected invalid 1, got %d", batchResp.Invalid)
	}
}
