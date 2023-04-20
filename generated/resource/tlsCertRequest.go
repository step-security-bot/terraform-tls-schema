package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const tlsCertRequest = `{
  "block": {
    "attributes": {
      "cert_request_pem": {
        "computed": true,
        "description": "The certificate request data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format. **NOTE**: the [underlying](https://pkg.go.dev/encoding/pem#Encode) [libraries](https://pkg.go.dev/golang.org/x/crypto/ssh#MarshalAuthorizedKey) that generate this value append a ` + "`" + `\\n` + "`" + ` at the end of the PEM. In case this disrupts your use case, we recommend using [` + "`" + `trimspace()` + "`" + `](https://www.terraform.io/language/functions/trimspace).",
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
      "id": {
        "computed": true,
        "description": "Unique identifier for this resource: hexadecimal representation of the SHA1 checksum of the resource.",
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
      "uris": {
        "description": "List of URIs for which a certificate is being requested (i.e. certificate subjects).",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
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
    "description": "Creates a Certificate Signing Request (CSR) in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.\n\nPEM is the typical format used to request a certificate from a Certificate Authority (CA).\n\nThis resource is intended to be used in conjunction with a Terraform provider for a particular certificate authority in order to provision a new certificate.",
    "description_kind": "markdown"
  },
  "version": 0
}`

func TlsCertRequestSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(tlsCertRequest), &result)
	return &result
}
