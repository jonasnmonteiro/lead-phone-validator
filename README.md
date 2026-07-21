# Phone Validator API 📞⚡

A high-performance microservice to validate, format, and normalize international phone numbers using Google's **LibPhoneNumber** (max metadata dataset). 

Specifically optimized for **chatbot integrations, CRM systems, and marketing automation bots** to guarantee correct format normalization before attempting outbound customer contact (e.g., initiating WhatsApp chats, SMS campaigns, or phone calls) without losing leads.

---

## 🚀 Key Features

* **Complete Global Validation**: Powered by `libphonenumber-js/max` supporting phone structures and types from all regions worldwide.
* **Full Phone Metadata**: Extends the validator to return offline carrier identification, geocoded geographic location, and timezone mappings.
* **Smart Calling Code Fallback**: Automatically recognizes and validates full international inputs missing the `+` prefix (e.g., converting `6581234567` automatically to Singapore `+65 8123 4567`).
* **WhatsApp Normalization Ready**: Returns standard WhatsApp format identifiers (`whatsapp`) and direct click-to-chat links (`whatsappUrl` e.g., `https://wa.me/number`).
* **Security & Footprint**: Light-weight, secure, and production-ready `Alpine` Docker container running as a non-root user.

---

## 🛠️ Getting Started

### Prerequisites

* [Node.js](https://nodejs.org/) (v22 or higher) OR [Docker](https://www.docker.com/)

### Running Locally

1. **Install dependencies:**
   ```bash
   npm install
   ```

2. **Start the server:**
   ```bash
   npm start
   ```
   *The server will start on port `3007`.*

---

### Running with Docker

1. **Build the Docker image:**
   ```bash
   docker build -t phone-validator-api .
   ```

2. **Run the container:**
   ```bash
   docker run -p 3007:3007 phone-validator-api
   ```

---

## 📖 API Reference

### Health Check

* **Endpoint**: `GET /`
* **Response**: `200 OK`
  ```json
  {
    "status": "ok",
    "message": "Phone Validator API online."
  }
  ```

---

### Validate Phone Number

* **Endpoint**: `POST /validate`
* **Headers**: `Content-Type: application/json`

#### Request Payload
| Field | Type | Required | Description |
| :--- | :--- | :--- | :--- |
| `phone` | `string` | **Yes** | The phone number to validate (can be local, international, formatted or raw). |
| `country` | `string` | No | ISO-2 country code hint (e.g., `BR`, `US`, `SG`) to help parse local numbers. |

**Example Request:**
```json
{
  "phone": "6581234567"
}
```

#### Response Payload (Valid Number)
```json
{
  "input": {
    "raw": "6581234567",
    "country": null
  },
  "parsed": {
    "isValid": true,
    "isPossible": true,
    "type": "MOBILE",
    "region": "SG"
  },
  "metadata": {
    "carrier": "Singtel",
    "location": "Singapore",
    "timezones": [
      "Asia/Singapore"
    ]
  },
  "formatted": {
    "e164": "+6581234567",
    "international": "+65 8123 4567",
    "national": "8123 4567",
    "rfc3966": "tel:+6581234567",
    "whatsapp": "6581234567",
    "whatsappUrl": "https://wa.me/6581234567"
  }
}
```

#### Response Payload (Invalid / Incomplete Number)
```json
{
  "input": {
    "raw": "1198765",
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

---

## 🤖 Bot Integration Example (cURL)

Verify numbers typed by users in chat widgets before storing in database or sending triggers:

```bash
curl -X POST "http://localhost:3007/validate" \
  -H "Content-Type: application/json" \
  -d '{
    "phone": "+55 (11) 98765-4321"
  }'
```
