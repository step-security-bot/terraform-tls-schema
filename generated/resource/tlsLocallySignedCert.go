package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const tlsLocallySignedCert = `{
  "block": {
    "attributes": {
      "allowed_uses": {
        "description": "List of key usages allowed for the issued certificate. Values are defined in [RFC 5280](https://datatracker.ietf.org/doc/html/rfc5280) and combine flags defined by both [Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.3) and [Extended Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.12). Accepted values: ` + "`" + `any_extended` + "`" + `, ` + "`" + `cert_signing` + "`" + `, ` + "`" + `client_auth` + "`" + `, ` + "`" + `code_signing` + "`" + `, ` + "`" + `content_commitment` + "`" + `, ` + "`" + `crl_signing` + "`" + `, ` + "`" + `data_encipherment` + "`" + `, ` + "`" + `decipher_only` + "`" + `, ` + "`" + `digital_signature` + "`" + `, ` + "`" + `email_protection` + "`" + `, ` + "`" + `encipher_only` + "`" + `, ` + "`" + `ipsec_end_system` + "`" + `, ` + "`" + `ipsec_tunnel` + "`" + `, ` + "`" + `ipsec_user` + "`" + `, ` + "`" + `key_agreement` + "`" + `, ` + "`" + `key_encipherment` + "`" + `, ` + "`" + `microsoft_commercial_code_signing` + "`" + `, ` + "`" + `microsoft_kernel_code_signing` + "`" + `, ` + "`" + `microsoft_server_gated_crypto` + "`" + `, ` + "`" + `netscape_server_gated_crypto` + "`" + `, ` + "`" + `ocsp_signing` + "`" + `, ` + "`" + `server_auth` + "`" + `, ` + "`" + `timestamping` + "`" + `.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "ca_cert_pem": {
        "description": "Certificate data of the Certificate Authority (CA) in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ca_key_algorithm": {
        "computed": true,
        "description": "Name of the algorithm used when generating the private key provided in ` + "`" + `ca_private_key_pem` + "`" + `. ",
        "description_kind": "plain",
        "type": "string"
      },
      "ca_private_key_pem": {
        "description": "Private key of the Certificate Authority (CA) used to sign the certificate, in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.",
        "description_kind": "plain",
        "required": true,
        "sensitive": true,
        "type": "string"
      },
      "cert_pem": {
        "computed": true,
        "description": "Certificate data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format. **NOTE**: the [underlying](https://pkg.go.dev/encoding/pem#Encode) [libraries](https://pkg.go.dev/golang.org/x/crypto/ssh#MarshalAuthorizedKey) that generate this value append a ` + "`" + `\\n` + "`" + ` at the end of the PEM. In case this disrupts your use case, we recommend using [` + "`" + `trimspace()` + "`" + `](https://www.terraform.io/language/functions/trimspace).",
        "description_kind": "plain",
        "type": "string"
      },
      "cert_request_pem": {
        "description": "Certificate request data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "early_renewal_hours": {
        "computed": true,
        "description": "The resource will consider the certificate to have expired the given number of hours before its actual expiry time. This can be useful to deploy an updated certificate in advance of the expiration of the current certificate. However, the old certificate remains valid until its true expiration time, since this resource does not (and cannot) support certificate revocation. Also, this advance update can only be performed should the Terraform configuration be applied during the early renewal period. (default: ` + "`" + `0` + "`" + `)",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "id": {
        "computed": true,
        "description": "Unique identifier for this resource: the certificate serial number.",
        "description_kind": "plain",
        "type": "string"
      },
      "is_ca_certificate": {
        "computed": true,
        "description": "Is the generated certificate representing a Certificate Authority (CA) (default: ` + "`" + `false` + "`" + `).",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "ready_for_renewal": {
        "computed": true,
        "description": "Is the certificate either expired (i.e. beyond the ` + "`" + `validity_period_hours` + "`" + `) or ready for an early renewal (i.e. within the ` + "`" + `early_renewal_hours` + "`" + `)?",
        "description_kind": "plain",
        "type": "bool"
      },
      "set_subject_key_id": {
        "computed": true,
        "description": "Should the generated certificate include a [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: ` + "`" + `false` + "`" + `).",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "validity_end_time": {
        "computed": true,
        "description": "The time until which the certificate is invalid, expressed as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.",
        "description_kind": "plain",
        "type": "string"
      },
      "validity_period_hours": {
        "description": "Number of hours, after initial issuing, that the certificate will remain valid for.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "validity_start_time": {
        "computed": true,
        "description": "The time after which the certificate is valid, expressed as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Creates a TLS certificate in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format using a Certificate Signing Request (CSR) and signs it with a provided (local) Certificate Authority (CA).",
    "description_kind": "markdown"
  },
  "version": 0
}`

func TlsLocallySignedCertSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(tlsLocallySignedCert), &result)
	return &result
}
