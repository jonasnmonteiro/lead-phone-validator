package telecom

import (
	"strings"
	"testing"

	"leadphone-validator/internal/models"
)

func TestLookupDDD(t *testing.T) {
	info, found := LookupDDD("11")
	if !found {
		t.Fatalf("expected DDD 11 to be found")
	}
	if info.State != "SP" {
		t.Fatalf("expected state SP, got %s", info.State)
	}
}

func TestSearchDDDByCity(t *testing.T) {
	results := SearchDDD("Florianopolis")
	if len(results) == 0 {
		t.Fatalf("expected search for Florianopolis to return results")
	}
	if results[0].Code != "48" {
		t.Fatalf("expected DDD 48, got %s", results[0].Code)
	}
}

func TestLookupDDI(t *testing.T) {
	info, found := LookupDDI("55")
	if !found {
		t.Fatalf("expected DDI 55 to be found")
	}
	if info.ISO2 != "BR" {
		t.Fatalf("expected ISO2 BR, got %s", info.ISO2)
	}
}

func TestSearchDDIByCountry(t *testing.T) {
	results := SearchDDI("Singapore")
	if len(results) == 0 {
		t.Fatalf("expected search for Singapore to return results")
	}
	if results[0].CallingCode != "65" {
		t.Fatalf("expected DDI 65, got %s", results[0].CallingCode)
	}
}

func TestFixNinthDigitMobile(t *testing.T) {
	resp := FixNinthDigit("1187654321")
	if !resp.Changed {
		t.Fatalf("expected mobile missing 9th digit to be changed")
	}
	if resp.Normalized != "11987654321" {
		t.Fatalf("expected 11987654321, got %s", resp.Normalized)
	}
}

func TestFixNinthDigitLandline(t *testing.T) {
	resp := FixNinthDigit("1134567890")
	if resp.Changed {
		t.Fatalf("expected landline to remain unchanged")
	}
	if !resp.IsLandline {
		t.Fatalf("expected IsLandline to be true")
	}
}

func TestCheckSuspiciousPatternsIdentical(t *testing.T) {
	resp := CheckSuspiciousPatterns("11111111111")
	if !resp.IsSuspicious {
		t.Fatalf("expected all identical digits to be suspicious")
	}
	if resp.Score < 80 {
		t.Fatalf("expected score >= 80, got %d", resp.Score)
	}
}

func TestBuildWhatsAppLinkWithMessage(t *testing.T) {
	resp := BuildWhatsAppLink(models.WhatsAppLinkRequest{
		Phone:   "+55 11 98765-4321",
		Message: "Hello, schedule demo",
	})
	if !strings.Contains(resp.URL, "https://wa.me/5511987654321") {
		t.Fatalf("expected URL to contain wa.me/5511987654321, got %s", resp.URL)
	}
	if !strings.Contains(resp.URL, "text=Hello%2C+schedule+demo") && !strings.Contains(resp.URL, "text=Hello%2C%20schedule%20demo") {
		t.Fatalf("expected URL to have encoded message, got %s", resp.URL)
	}
}
