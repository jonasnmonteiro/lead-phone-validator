import express from "express";
import cors from "cors";
import { parsePhoneNumberFromString } from "libphonenumber-js/max";
import { carrier, geocoder, timezones } from "libphonenumber-geo-carrier";

const app = express();
app.use(cors());
app.use(express.json());

function cleanNumber(n) {
  return n.replace(/\D/g, "");
}

// Helper to format consistent responses
async function sendResponse(res, phone, countryInput, parsed, extra = {}) {
  const isValid = parsed ? parsed.isValid() : false;
  
  let metadata = null;
  if (parsed && isValid) {
    try {
      const [carrierName, locationName, tzList] = await Promise.all([
        carrier(parsed),
        geocoder(parsed),
        timezones(parsed)
      ]);
      metadata = {
        carrier: carrierName || null,
        location: locationName || null,
        timezones: tzList || []
      };
    } catch (e) {
      metadata = {
        carrier: null,
        location: null,
        timezones: []
      };
    }
  }

  return res.json({
    input: {
      raw: phone,
      country: countryInput || null
    },
    parsed: {
      isValid,
      isPossible: parsed ? parsed.isPossible() : false,
      type: (parsed && isValid) ? (parsed.getType() || null) : null,
      region: parsed ? (parsed.country || null) : null,
      ...extra
    },
    metadata,
    formatted: (parsed && isValid) ? {
      e164: parsed.format('E.164') || null,
      international: parsed.format('INTERNATIONAL') || null,
      national: parsed.format('NATIONAL') || null,
      rfc3966: parsed.format('RFC3966') || null,
      whatsapp: parsed.number ? parsed.number.replace("+", "") : null,
      whatsappUrl: parsed.number ? `https://wa.me/${parsed.number.replace("+", "")}` : null
    } : {}
  });
}

app.post("/validate", async (req, res) => {
  const { phone, country } = req.body;

  if (!phone) {
    return res.status(400).json({
      error: "'phone' field is required."
    });
  }

  const raw = cleanNumber(phone);

  try {
    // 0) If country is provided, attempt parsing directly with it first
    if (country) {
      const parsed = parsePhoneNumberFromString(phone, country.toUpperCase());
      if (parsed?.isValid()) {
        return await sendResponse(res, phone, country, parsed);
      }
    }

    // 0.5) If number doesn't start with "+" but prepending "+" makes it a valid international number, use it
    if (!phone.trim().startsWith("+")) {
      const parsed = parsePhoneNumberFromString("+" + raw);
      if (parsed?.isValid()) {
        return await sendResponse(res, phone, country, parsed);
      }
    }

    // 1) If the number starts with "+" → validate directly (international)
    if (phone.trim().startsWith("+")) {
      const parsed = parsePhoneNumberFromString(phone);
      if (parsed?.isValid()) {
        return await sendResponse(res, phone, country, parsed);
      }
      return await sendResponse(res, phone, country, null, { reason: "Invalid number." });
    }

    // 2) Numbers with 8–9 digits → Brazil without DDD → ask for DDD
    if (raw.length === 8 || raw.length === 9) {
      return await sendResponse(res, phone, country, null, {
        needsDDD: true,
        reason: "Brazilian number without area code."
      });
    }

    // 3) Numbers with 10 digits → potential USA/Canada
    if (raw.length === 10) {
      const parsed = parsePhoneNumberFromString(phone, "US");
      if (parsed?.isValid()) {
        return await sendResponse(res, phone, country, parsed);
      }

      return await sendResponse(res, phone, country, null, {
        needsCountry: true,
        reason: "10-digit number does not clearly correspond to a country."
      });
    }

    // 4) Number starting with "55" → can be Brazil complete
    if (raw.startsWith("55")) {
      const parsed = parsePhoneNumberFromString("+" + raw);
      if (parsed?.isValid()) {
        return await sendResponse(res, phone, country, parsed);
      }

      return await sendResponse(res, phone, country, null, {
        reason: "Number starting with 55, but invalid. Missing area code?"
      });
    }

    // 5) Numbers with 11 digits → potential BR (DDD + mobile)
    if (raw.length === 11) {
      const parsed = parsePhoneNumberFromString(raw, "BR");
      if (parsed?.isValid()) {
        return await sendResponse(res, phone, country, parsed);
      }

      return await sendResponse(res, phone, country, null, {
        needsDDD: true,
        reason: "Invalid 11-digit format for BR."
      });
    }

    // 6) International numbers (12+ digits)
    if (raw.length >= 12) {
      const parsed = parsePhoneNumberFromString("+" + raw);
      if (parsed?.isValid()) {
        return await sendResponse(res, phone, country, parsed);
      }

      return await sendResponse(res, phone, country, null, { reason: "Invalid international number." });
    }

    // 7) fallback → invalid number
    return await sendResponse(res, phone, country, null, { reason: "Invalid number." });

  } catch (error) {
    return res.status(500).json({
      error: "Internal error.",
      message: error.message
    });
  }
});

app.get("/", (req, res) => {
  res.json({
    status: "ok",
    message: "Phone Validator API online."
  });
});

const PORT = process.env.PORT || 3007;
app.listen(PORT, () => console.log(`Phone Validator online on port ${PORT}`));

