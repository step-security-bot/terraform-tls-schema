package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const tlsCertificate = `{
  "block": {
    "attributes": {
      "certificates": {
        "computed": true,
        "description": "The certificates protecting the site, with the root of the chain first.",
        "description_kind": "markdown",
        "type": [
          "list",
          [
            "object",
            {
              "cert_pem": "string",
              "is_ca": "bool",
              "issuer": "string",
              "not_after": "string",
              "not_before": "string",
              "public_key_algorithm": "string",
              "serial_number": "string",
              "sha1_fingerprint": "string",
              "signature_algorithm": "string",
              "subject": "string",
              "version": "number"
            }
          ]
        ]
      },
      "content": {
        "description": "The content of the certificate in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.",
        "description_kind": "markdown",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Unique identifier of this data source: hashing of the certificates in the chain.",
        "description_kind": "markdown",
        "type": "string"
      },
      "url": {
        "description": "URL of the endpoint to get the certificates from. Accepted schemes are: ` + "`" + `https` + "`" + `, ` + "`" + `tls` + "`" + `. For scheme ` + "`" + `https://` + "`" + ` it will use the HTTP protocol and apply the ` + "`" + `proxy` + "`" + ` configuration of the provider, if set. For scheme ` + "`" + `tls://` + "`" + ` it will instead use a secure TCP socket.",
        "description_kind": "markdown",
        "optional": true,
        "type": "string"
      },
      "verify_chain": {
        "description": "Whether to verify the certificate chain while parsing it or not (default: ` + "`" + `true` + "`" + `).",
        "description_kind": "markdown",
        "optional": true,
        "type": "bool"
      }
    },
    "description": "Get information about the TLS certificates securing a host.\n\nUse this data source to get information, such as SHA1 fingerprint or serial number, about the TLS certificates that protects a URL.",
    "description_kind": "markdown"
  },
  "version": 0
}`

func TlsCertificateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(tlsCertificate), &result)
	return &result
}
