package telecom

import (
	"strings"

	"leadphone-validator/internal/models"
)

var ddiDatabase = []models.DDIInfo{
	{CallingCode: "1", CountryName: "United States / Canada", ISO2: "US", ISO3: "USA", IDDPrefix: "011", EmergencyNumbers: []string{"911"}, TimezoneHint: "America/New_York"},
	{CallingCode: "44", CountryName: "United Kingdom", ISO2: "GB", ISO3: "GBR", IDDPrefix: "00", EmergencyNumbers: []string{"999", "112"}, TimezoneHint: "Europe/London"},
	{CallingCode: "55", CountryName: "Brazil", ISO2: "BR", ISO3: "BRA", IDDPrefix: "00", EmergencyNumbers: []string{"190", "192", "193"}, TimezoneHint: "America/Sao_Paulo"},
	{CallingCode: "351", CountryName: "Portugal", ISO2: "PT", ISO3: "PRT", IDDPrefix: "00", EmergencyNumbers: []string{"112"}, TimezoneHint: "Europe/Lisbon"},
	{CallingCode: "34", CountryName: "Spain", ISO2: "ES", ISO3: "ESP", IDDPrefix: "00", EmergencyNumbers: []string{"112"}, TimezoneHint: "Europe/Madrid"},
	{CallingCode: "33", CountryName: "France", ISO2: "FR", ISO3: "FRA", IDDPrefix: "00", EmergencyNumbers: []string{"112", "15", "17", "18"}, TimezoneHint: "Europe/Paris"},
	{CallingCode: "49", CountryName: "Germany", ISO2: "DE", ISO3: "DEU", IDDPrefix: "00", EmergencyNumbers: []string{"110", "112"}, TimezoneHint: "Europe/Berlin"},
	{CallingCode: "39", CountryName: "Italy", ISO2: "IT", ISO3: "ITA", IDDPrefix: "00", EmergencyNumbers: []string{"112", "113", "118"}, TimezoneHint: "Europe/Rome"},
	{CallingCode: "65", CountryName: "Singapore", ISO2: "SG", ISO3: "SGP", IDDPrefix: "001", EmergencyNumbers: []string{"999", "995"}, TimezoneHint: "Asia/Singapore"},
	{CallingCode: "81", CountryName: "Japan", ISO2: "JP", ISO3: "JPN", IDDPrefix: "010", EmergencyNumbers: []string{"110", "119"}, TimezoneHint: "Asia/Tokyo"},
	{CallingCode: "86", CountryName: "China", ISO2: "CN", ISO3: "CHN", IDDPrefix: "00", EmergencyNumbers: []string{"110", "120", "119"}, TimezoneHint: "Asia/Shanghai"},
	{CallingCode: "91", CountryName: "India", ISO2: "IN", ISO3: "IND", IDDPrefix: "00", EmergencyNumbers: []string{"112"}, TimezoneHint: "Asia/Kolkata"},
	{CallingCode: "61", CountryName: "Australia", ISO2: "AU", ISO3: "AUS", IDDPrefix: "0011", EmergencyNumbers: []string{"000"}, TimezoneHint: "Australia/Sydney"},
	{CallingCode: "52", CountryName: "Mexico", ISO2: "MX", ISO3: "MEX", IDDPrefix: "00", EmergencyNumbers: []string{"911"}, TimezoneHint: "America/Mexico_City"},
	{CallingCode: "54", CountryName: "Argentina", ISO2: "AR", ISO3: "ARG", IDDPrefix: "00", EmergencyNumbers: []string{"911", "101", "107"}, TimezoneHint: "America/Argentina/Buenos_Aires"},
	{CallingCode: "56", CountryName: "Chile", ISO2: "CL", ISO3: "CHL", IDDPrefix: "1230", EmergencyNumbers: []string{"133", "131", "132"}, TimezoneHint: "America/Santiago"},
	{CallingCode: "57", CountryName: "Colombia", ISO2: "CO", ISO3: "COL", IDDPrefix: "00", EmergencyNumbers: []string{"123"}, TimezoneHint: "America/Bogota"},
	{CallingCode: "51", CountryName: "Peru", ISO2: "PE", ISO3: "PER", IDDPrefix: "00", EmergencyNumbers: []string{"105", "116"}, TimezoneHint: "America/Lima"},
	{CallingCode: "598", CountryName: "Uruguay", ISO2: "UY", ISO3: "URY", IDDPrefix: "00", EmergencyNumbers: []string{"911"}, TimezoneHint: "America/Montevideo"},
	{CallingCode: "595", CountryName: "Paraguay", ISO2: "PY", ISO3: "PRY", IDDPrefix: "00", EmergencyNumbers: []string{"911"}, TimezoneHint: "America/Asuncion"},
	{CallingCode: "31", CountryName: "Netherlands", ISO2: "NL", ISO3: "NLD", IDDPrefix: "00", EmergencyNumbers: []string{"112"}, TimezoneHint: "Europe/Amsterdam"},
	{CallingCode: "41", CountryName: "Switzerland", ISO2: "CH", ISO3: "CHE", IDDPrefix: "00", EmergencyNumbers: []string{"112", "117", "118", "144"}, TimezoneHint: "Europe/Zurich"},
	{CallingCode: "46", CountryName: "Sweden", ISO2: "SE", ISO3: "SWE", IDDPrefix: "00", EmergencyNumbers: []string{"112"}, TimezoneHint: "Europe/Stockholm"},
	{CallingCode: "47", CountryName: "Norway", ISO2: "NO", ISO3: "NOR", IDDPrefix: "00", EmergencyNumbers: []string{"112", "110", "113"}, TimezoneHint: "Europe/Oslo"},
	{CallingCode: "45", CountryName: "Denmark", ISO2: "DK", ISO3: "DNK", IDDPrefix: "00", EmergencyNumbers: []string{"112"}, TimezoneHint: "Europe/Copenhagen"},
	{CallingCode: "353", CountryName: "Ireland", ISO2: "IE", ISO3: "IRL", IDDPrefix: "00", EmergencyNumbers: []string{"999", "112"}, TimezoneHint: "Europe/Dublin"},
	{CallingCode: "971", CountryName: "United Arab Emirates", ISO2: "AE", ISO3: "ARE", IDDPrefix: "00", EmergencyNumbers: []string{"999", "998"}, TimezoneHint: "Asia/Dubai"},
	{CallingCode: "966", CountryName: "Saudi Arabia", ISO2: "SA", ISO3: "SAU", IDDPrefix: "00", EmergencyNumbers: []string{"999", "997"}, TimezoneHint: "Asia/Riyadh"},
	{CallingCode: "27", CountryName: "South Africa", ISO2: "ZA", ISO3: "ZAF", IDDPrefix: "00", EmergencyNumbers: []string{"10111", "10177"}, TimezoneHint: "Africa/Johannesburg"},
	{CallingCode: "82", CountryName: "South Korea", ISO2: "KR", ISO3: "KOR", IDDPrefix: "001", EmergencyNumbers: []string{"112", "119"}, TimezoneHint: "Asia/Seoul"},
}

func LookupDDI(code string) (models.DDIInfo, bool) {
	cleanCode := strings.TrimPrefix(strings.TrimSpace(code), "+")
	for _, item := range ddiDatabase {
		if item.CallingCode == cleanCode {
			return item, true
		}
	}
	return models.DDIInfo{}, false
}

func SearchDDI(query string) []models.DDIInfo {
	q := strings.ToLower(strings.TrimSpace(strings.TrimPrefix(query, "+")))
	if q == "" {
		return ddiDatabase
	}

	var results []models.DDIInfo
	for _, item := range ddiDatabase {
		if item.CallingCode == q || strings.ToLower(item.ISO2) == q || strings.ToLower(item.ISO3) == q || strings.Contains(strings.ToLower(item.CountryName), q) {
			results = append(results, item)
		}
	}

	return results
}
