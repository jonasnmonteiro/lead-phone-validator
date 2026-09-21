# LeadPhone Validator

A high-performance phone number validation, telecom intelligence, and lead normalization microservice written in Go.

LeadPhone Validator parses, sanitizes, and normalizes international telephone numbers using Google libphonenumber metadata. It provides Brazilian DDD area code intelligence, global DDI calling code directories, 9th digit normalization, synthetic pattern detection, and WhatsApp click-to-chat generation for marketing automation bots, CRM funnels, and customer communication workflows.

## Core Capabilities

* Sub-millisecond phone parsing and validity checking for all global territories.
* Complete Brazilian DDD intelligence covering all 67 official area codes with municipal reverse search.
* Global DDI calling code directory with ISO codes, IDD exit prefixes, and emergency numbers.
* Brazilian 9th digit auto-fixer converting legacy 8-digit mobiles while preserving landlines.
* Suspicious pattern and dummy lead detector flagging repeated and sequential digits.
* WhatsApp link builder supporting custom URL-encoded pre-filled greeting templates.
* Offline carrier name identification, geocoding location resolution, and timezone extraction.
* High-throughput concurrent batch processing via native Go worker routines.
* Minimal resource footprint compiling to a standalone static binary in Alpine Linux.

## Requirements

* Go 1.22 or higher (for local builds)
* Node.js 20 or higher (for Astro frontend builds)
* Docker (optional, for containerized deployments)

## Quick Start

### Running Locally

```bash
make run
```

The server starts by default on port 3007.

To run the interactive Astro frontend with hot-reload in development:

```bash
make web-dev
```

### Running Tests

```bash
make test
```

### Building Binary

```bash
make build
```

Compiled binary will be created at `bin/leadphone-validator`.

### Running with Docker

```bash
make docker-build
make docker-run
```

## API Documentation

### 1. Health Check

* Method: `GET`
* Paths: `/`, `/health`, `/v1/health`

```json
{
  "status": "ok",
  "service": "LeadPhone Validator",
  "version": "1.1.0"
}
```

### 2. Single Phone Validation

* Method: `POST`
* Paths: `/validate`, `/v1/validate`

Request Payload:

```json
{
  "phone": "+55 (11) 98765-4321"
}
```

### 3. Batch Lead Validation

* Method: `POST`
* Paths: `/validate/batch`, `/v1/validate/batch`

Request Payload:

```json
{
  "phones": [
    { "phone": "6581234567" },
    { "phone": "+55 11 98765-4321" },
    { "phone": "987654321" }
  ]
}
```

### 4. Brazilian DDD Area Code Lookup

* Method: `GET`
* Paths: `/v1/ddd?q=11` or `/v1/ddd?q=curitiba`

Example Response:

```json
{
  "query": "11",
  "found": true,
  "results": [
    {
      "code": "11",
      "state": "SP",
      "stateName": "São Paulo",
      "region": "Sudeste",
      "majorCities": ["São Paulo", "Guarulhos", "Santo André", "Osasco"],
      "timezone": "America/Sao_Paulo"
    }
  ]
}
```

### 5. International DDI Directory

* Method: `GET`
* Paths: `/v1/ddi?q=55` or `/v1/ddi?q=singapore`

Example Response:

```json
{
  "query": "55",
  "found": true,
  "results": [
    {
      "callingCode": "55",
      "countryName": "Brazil",
      "iso2": "BR",
      "iso3": "BRA",
      "iddPrefix": "00",
      "emergencyNumbers": ["190", "192", "193"],
      "timezoneHint": "America/Sao_Paulo"
    }
  ]
}
```

### 6. Brazilian 9th Digit Normalizer

* Method: `POST`
* Paths: `/v1/sanitize/ninth-digit`

Request Payload:

```json
{
  "phone": "1187654321"
}
```

Response:

```json
{
  "input": "1187654321",
  "normalized": "11987654321",
  "changed": true,
  "isMobile": true,
  "isLandline": false,
  "isValid": true,
  "message": "Inserted ninth digit 9 after area code."
}
```

### 7. WhatsApp Direct Link Builder

* Method: `POST`
* Paths: `/v1/whatsapp/link`

Request Payload:

```json
{
  "phone": "+55 11 98765-4321",
  "message": "Hello, scheduling a product demo"
}
```

Response:

```json
{
  "phone": "+55 11 98765-4321",
  "e164": "+5511987654321",
  "message": "Hello, scheduling a product demo",
  "url": "https://wa.me/5511987654321?text=Hello%2C+scheduling+a+product+demo",
  "isValid": true
}
```

### 8. Suspicious Pattern Detector

* Method: `POST`
* Paths: `/v1/check/patterns`

Request Payload:

```json
{
  "phone": "11999999999"
}
```

Response:

```json
{
  "phone": "11999999999",
  "isSuspicious": false,
  "flags": ["REPEATED_TRAIL_DIGITS"],
  "score": 30,
  "description": "Minor pattern flags detected. Verify line activity."
}
```
