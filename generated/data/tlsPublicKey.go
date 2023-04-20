package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const tlsPublicKey = `{
  "block": {
    "attributes": {
      "algorithm": {
        "computed": true,
        "description": "The name of the algorithm used by the given private key. Possible values are: ` + "`" + `RSA` + "`" + `, ` + "`" + `ECDSA` + "`" + `, ` + "`" + `ED25519` + "`" + `. ",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Unique identifier for this data source: hexadecimal representation of the SHA1 checksum of the data source.",
        "description_kind": "plain",
        "type": "string"
      },
      "private_key_openssh": {
        "description": "The private key (in  [OpenSSH PEM (RFC 4716)](https://datatracker.ietf.org/doc/html/rfc4716) format) to extract the public key from. This is _mutually exclusive_ with ` + "`" + `private_key_pem` + "`" + `. Currently-supported algorithms for keys are: ` + "`" + `RSA` + "`" + `, ` + "`" + `ECDSA` + "`" + `, ` + "`" + `ED25519` + "`" + `. ",
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "private_key_pem": {
        "description": "The private key (in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format) to extract the public key from. This is _mutually exclusive_ with ` + "`" + `private_key_openssh` + "`" + `. Currently-supported algorithms for keys are: ` + "`" + `RSA` + "`" + `, ` + "`" + `ECDSA` + "`" + `, ` + "`" + `ED25519` + "`" + `. ",
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "public_key_fingerprint_md5": {
        "computed": true,
        "description": "The fingerprint of the public key data in OpenSSH MD5 hash format, e.g. ` + "`" + `aa:bb:cc:...` + "`" + `. Only available if the selected private key format is compatible, as per the rules for ` + "`" + `public_key_openssh` + "`" + ` and [ECDSA P224 limitations](../../docs#limitations).",
        "description_kind": "plain",
        "type": "string"
      },
      "public_key_fingerprint_sha256": {
        "computed": true,
        "description": "The fingerprint of the public key data in OpenSSH SHA256 hash format, e.g. ` + "`" + `SHA256:...` + "`" + `. Only available if the selected private key format is compatible, as per the rules for ` + "`" + `public_key_openssh` + "`" + ` and [ECDSA P224 limitations](../../docs#limitations).",
        "description_kind": "plain",
        "type": "string"
      },
      "public_key_openssh": {
        "computed": true,
        "description": "The public key, in  [OpenSSH PEM (RFC 4716)](https://datatracker.ietf.org/doc/html/rfc4716) format. This is also known as ['Authorized Keys'](https://www.ssh.com/academy/ssh/authorized_keys/openssh#format-of-the-authorized-keys-file) format. This is not populated for ` + "`" + `ECDSA` + "`" + ` with curve ` + "`" + `P224` + "`" + `, as it is [not supported](../../docs#limitations). **NOTE**: the [underlying](https://pkg.go.dev/encoding/pem#Encode) [libraries](https://pkg.go.dev/golang.org/x/crypto/ssh#MarshalAuthorizedKey) that generate this value append a ` + "`" + `\\n` + "`" + ` at the end of the PEM. In case this disrupts your use case, we recommend using [` + "`" + `trimspace()` + "`" + `](https://www.terraform.io/language/functions/trimspace).",
        "description_kind": "plain",
        "type": "string"
      },
      "public_key_pem": {
        "computed": true,
        "description": "The public key, in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format. **NOTE**: the [underlying](https://pkg.go.dev/encoding/pem#Encode) [libraries](https://pkg.go.dev/golang.org/x/crypto/ssh#MarshalAuthorizedKey) that generate this value append a ` + "`" + `\\n` + "`" + ` at the end of the PEM. In case this disrupts your use case, we recommend using [` + "`" + `trimspace()` + "`" + `](https://www.terraform.io/language/functions/trimspace).",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Get a public key from a PEM-encoded private key.\n\nUse this data source to get the public key from a [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) or [OpenSSH PEM (RFC 4716)](https://datatracker.ietf.org/doc/html/rfc4716) formatted private key, for use in other resources.",
    "description_kind": "markdown"
  },
  "version": 0
}`

func TlsPublicKeySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(tlsPublicKey), &result)
	return &result
}
