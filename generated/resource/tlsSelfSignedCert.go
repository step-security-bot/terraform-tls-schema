package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const tlsSelfSignedCert = `{
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
      "cert_pem": {
        "computed": true,
        "description": "Certificate data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format. **NOTE**: the [underlying](https://pkg.go.dev/encoding/pem#Encode) [libraries](https://pkg.go.dev/golang.org/x/crypto/ssh#MarshalAuthorizedKey) that generate this value append a ` + "`" + `\\n` + "`" + ` at the end of the PEM. In case this disrupts your use case, we recommend using [` + "`" + `trimspace()` + "`" + `](https://www.terraform.io/language/functions/trimspace).",
        "description_kind": "plain",
        "type": "string"
      },
      "dns_names": {
        "description": "List of DNS names for which a certificate is being requested (i.e. certificate subjects).",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
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
      "ip_addresses": {
        "description": "List of IP addresses for which a certificate is being requested (i.e. certificate subjects).",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "is_ca_certificate": {
        "computed": true,
        "description": "Is the generated certificate representing a Certificate Authority (CA) (default: ` + "`" + `false` + "`" + `).",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "key_algorithm": {
        "computed": true,
        "description": "Name of the algorithm used when generating the private key provided in ` + "`" + `private_key_pem` + "`" + `. ",
        "description_kind": "plain",
        "type": "string"
      },
      "private_key_pem": {
        "description": "Private key in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format, that the certificate will belong to. This can be read from a separate file using the [` + "`" + `file` + "`" + `](https://www.terraform.io/language/functions/file) interpolation function. Only an irreversible secure hash of the private key will be stored in the Terraform state.",
        "description_kind": "plain",
        "required": true,
        "sensitive": true,
        "type": "string"
      },
      "ready_for_renewal": {
        "computed": true,
        "description": "Is the certificate either expired (i.e. beyond the ` + "`" + `validity_period_hours` + "`" + `) or ready for an early renewal (i.e. within the ` + "`" + `early_renewal_hours` + "`" + `)?",
        "description_kind": "plain",
        "type": "bool"
      },
      "set_authority_key_id": {
        "computed": true,
        "description": "Should the generated certificate include an [authority key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.1): for self-signed certificates this is the same value as the [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: ` + "`" + `false` + "`" + `).",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "set_subject_key_id": {
        "computed": true,
        "description": "Should the generated certificate include a [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: ` + "`" + `false` + "`" + `).",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "uris": {
        "description": "List of URIs for which a certificate is being requested (i.e. certificate subjects).",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
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
    "block_types": {
      "subject": {
        "block": {
          "attributes": {
            "common_name": {
              "description": "Distinguished name: ` + "`" + `CN` + "`" + `",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "country": {
              "description": "Distinguished name: ` + "`" + `C` + "`" + `",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "locality": {
              "description": "Distinguished name: ` + "`" + `L` + "`" + `",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "organization": {
              "description": "Distinguished name: ` + "`" + `O` + "`" + `",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "organizational_unit": {
              "description": "Distinguished name: ` + "`" + `OU` + "`" + `",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "postal_code": {
              "description": "Distinguished name: ` + "`" + `PC` + "`" + `",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "province": {
              "description": "Distinguished name: ` + "`" + `ST` + "`" + `",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "serial_number": {
              "description": "Distinguished name: ` + "`" + `SERIALNUMBER` + "`" + `",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "street_address": {
              "description": "Distinguished name: ` + "`" + `STREET` + "`" + `",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "description": "The subject for which a certificate is being requested. The acceptable arguments are all optional and their naming is based upon [Issuer Distinguished Names (RFC5280)](https://tools.ietf.org/html/rfc5280#section-4.1.2.4) section.",
          "description_kind": "markdown"
        },
        "max_items": 1,
        "nesting_mode": "list"
      }
    },
    "description": "Creates a **self-signed** TLS certificate in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.",
    "description_kind": "markdown"
  },
  "version": 0
}`

func TlsSelfSignedCertSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(tlsSelfSignedCert), &result)
	return &result
}
