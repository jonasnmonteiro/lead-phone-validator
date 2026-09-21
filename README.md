# LeadPhone Validator

A high-performance phone number validation, normalization, and telecom intelligence microservice written in Go.

LeadPhone Validator parses, sanitizes, and normalizes international telephone numbers using Google libphonenumber metadata. It is built for marketing automations, CRM pipelines, lead capture forms, and chatbot routing to prevent lead loss and format numbers before outbound communication.

## Core Capabilities

* Sub-millisecond phone parsing and validity checking for all global territories.
* Offline carrier name identification, geocoding location resolution, and timezone extraction.
* Intelligent dialing code recovery for international numbers missing the leading plus prefix.
* Area code detection and validation for domestic and international formats.
* Out-of-the-box WhatsApp normalization providing raw digits and direct click to chat links.
* High-throughput concurrent batch processing via native Go worker routines.
* Minimal resource footprint compiling to a standalone static binary in Alpine Linux.

## Requirements

* Go 1.22 or higher (for local builds)
* Docker (optional, for containerized deployments)

## Quick Start

### Running Locally

```bash
make run
```

Or using the Go toolchain directly:

```bash
go run .
```

The server starts by default on port 3007. To specify a custom port:

```bash
PORT=8080 go run .
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

Or manually:

```bash
docker build -t leadphone-validator .
docker run -p 3007:3007 leadphone-validator
```

## API Documentation

### 1. Health Check

* Method: `GET`
* Paths: `/`, `/health`

Example Response:

```json
{
  "status": "ok",
  "service": "LeadPhone Validator",
  "version": "1.0.0"
}
```

### 2. Single Phone Validation

* Method: `POST`
* Paths: `/validate`, `/v1/validate`
* Headers: `Content-Type: application/json`

#### Request Payload

| Field | Type | Required | Description |
| :--- | :--- | :--- | :--- |
| `phone` | `string` | Yes | Phone number in local or international format. |
| `country` | `string` | No | ISO-2 country code hint (e.g. BR, US, SG). |

#### Request Example

```json
{
  "phone": "+55 (11) 98765-4321"
}
```

#### Valid Number Response

```json
{
  "input": {
    "raw": "+55 (11) 98765-4321",
    "country": null
  },
  "parsed": {
    "isValid": true,
    "isPossible": true,
    "type": "MOBILE",
    "region": "BR",
    "countryCode": 55,
    "nationalNumber": 11987654321
  },
  "metadata": {
    "carrier": "TIM",
    "location": "São Paulo",
    "timezones": [
      "America/Sao_Paulo"
    ]
  },
  "formatted": {
    "e164": "+5511987654321",
    "international": "+55 11 98765-4321",
    "national": "(11) 98765-4321",
    "rfc3966": "tel:+55-11-98765-4321",
    "whatsapp": "5511987654321",
    "whatsappUrl": "https://wa.me/5511987654321"
  }
}
```

#### Missing Area Code Response

```json
{
  "input": {
    "raw": "987654321",
    "country": null
  },
  "parsed": {
    "isValid": false,
    "isPossible": false,
    "type": null,
    "region": null,
    "needsDDD": true,
    "reason": "Brazilian number without area code."
  },
  "metadata": null,
  "formatted": {}
}
```

### 3. Batch Validation

* Method: `POST`
* Paths: `/validate/batch`, `/v1/validate/batch`
* Headers: `Content-Type: application/json`

#### Request Example

```json
{
  "phones": [
    { "phone": "6581234567" },
    { "phone": "+55 11 98765-4321" },
    { "phone": "987654321" }
  ]
}
```

#### Response Example

```json
{
  "total": 3,
  "valid": 2,
  "invalid": 1,
  "results": [
    {
      "input": { "raw": "6581234567", "country": null },
      "parsed": {
        "isValid": true,
        "isPossible": true,
        "type": "MOBILE",
        "region": "SG",
        "countryCode": 65,
        "nationalNumber": 81234567
      },
      "metadata": {
        "carrier": "SingTel",
        "location": "Singapore",
        "timezones": ["Asia/Singapore"]
      },
      "formatted": {
        "e164": "+6581234567",
        "international": "+65 8123 4567",
        "national": "8123 4567",
        "rfc3966": "tel:+65-8123-4567",
        "whatsapp": "6581234567",
        "whatsappUrl": "https://wa.me/6581234567"
      }
    },
    {
      "input": { "raw": "+55 11 98765-4321", "country": null },
      "parsed": {
        "isValid": true,
        "isPossible": true,
        "type": "MOBILE",
        "region": "BR",
        "countryCode": 55,
        "nationalNumber": 11987654321
      },
      "metadata": {
        "carrier": "TIM",
        "location": "São Paulo",
        "timezones": ["America/Sao_Paulo"]
      },
      "formatted": {
        "e164": "+5511987654321",
        "international": "+55 11 98765-4321",
        "national": "(11) 98765-4321",
        "rfc3966": "tel:+55-11-98765-4321",
        "whatsapp": "5511987654321",
        "whatsappUrl": "https://wa.me/5511987654321"
      }
    },
    {
      "input": { "raw": "987654321", "country": null },
      "parsed": {
        "isValid": false,
        "isPossible": false,
        "type": null,
        "region": null,
        "needsDDD": true,
        "reason": "Brazilian number without area code."
      },
      "metadata": null,
      "formatted": {}
    }
  ]
}
```

## Integration Examples

### cURL

```bash
curl -X POST "http://localhost:3007/validate" \
  -H "Content-Type: application/json" \
  -d '{"phone": "+5511987654321"}'
```

### Python

```python
import requests

response = requests.post(
    "http://localhost:3007/validate",
    json={"phone": "+5511987654321"}
)
data = response.json()
if data["parsed"]["isValid"]:
    print("WhatsApp Link:", data["formatted"]["whatsappUrl"])
```

### Node.js / TypeScript

```typescript
const res = await fetch("http://localhost:3007/validate", {
  method: "POST",
  headers: { "Content-Type": "application/json" },
  body: JSON.stringify({ phone: "+5511987654321" }),
});
const data = await res.json();
console.log(data);
```
