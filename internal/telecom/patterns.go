package telecom

import (
	"strconv"
	"strings"

	"leadphone-validator/internal/models"
)

func CheckSuspiciousPatterns(phone string) models.PatternCheckResponse {
	raw := SanitizeDigits(phone)
	if len(raw) < 7 {
		return models.PatternCheckResponse{
			Phone:        phone,
			IsSuspicious: true,
			Flags:        []string{"TOO_SHORT"},
			Score:        95,
			Description:  "Number contains too few digits to be a legitimate subscriber line.",
		}
	}

	var flags []string
	score := 0

	allSame := true
	for i := 1; i < len(raw); i++ {
		if raw[i] != raw[0] {
			allSame = false
			break
		}
	}
	if allSame {
		flags = append(flags, "ALL_IDENTICAL_DIGITS")
		score += 90
	}

	isAscending := true
	for i := 1; i < len(raw); i++ {
		if int(raw[i])-int(raw[i-1]) != 1 {
			isAscending = false
			break
		}
	}
	if isAscending {
		flags = append(flags, "SEQUENTIAL_ASCENDING_DIGITS")
		score += 85
	}

	isDescending := true
	for i := 1; i < len(raw); i++ {
		if int(raw[i-1])-int(raw[i]) != 1 {
			isDescending = false
			break
		}
	}
	if isDescending {
		flags = append(flags, "SEQUENTIAL_DESCENDING_DIGITS")
		score += 85
	}

	if len(raw) >= 10 && strings.Contains(raw, "55501") {
		idx := strings.Index(raw, "55501")
		if idx+7 <= len(raw) {
			suffix := raw[idx+5 : idx+7]
			if val, err := strconv.Atoi(suffix); err == nil && val >= 0 && val <= 99 {
				flags = append(flags, "FICTITIOUS_NANPA_555_RANGE")
				score += 95
			}
		}
	}

	if len(raw) >= 8 {
		repeatedSuffix := true
		lastFour := raw[len(raw)-4:]
		for i := 1; i < len(lastFour); i++ {
			if lastFour[i] != lastFour[0] {
				repeatedSuffix = false
				break
			}
		}
		if repeatedSuffix && (lastFour == "0000" || lastFour == "9999" || lastFour == "1111") {
			flags = append(flags, "REPEATED_TRAIL_DIGITS")
			score += 30
		}
	}

	if (len(raw) == 10 || len(raw) == 11) && !allSame && !isAscending {
		ddd := raw[0:2]
		if _, exists := LookupDDD(ddd); !exists {
			flags = append(flags, "NON_EXISTENT_BRAZIL_DDD")
			score += 45
		}
	}

	if score > 100 {
		score = 100
	}

	isSuspicious := score >= 50

	var desc string
	if isSuspicious {
		desc = "High probability of dummy, test, or invalid sequence pattern."
	} else if score > 0 {
		desc = "Minor pattern flags detected. Verify line activity."
	} else {
		desc = "Clean digit distribution. No synthetic patterns detected."
	}

	return models.PatternCheckResponse{
		Phone:        phone,
		IsSuspicious: isSuspicious,
		Flags:        flags,
		Score:        score,
		Description:  desc,
	}
}
