package generated

import (
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-tls-schema/v4/generated/data"
	"github.com/lonegunmanb/terraform-tls-schema/v4/generated/resource"
)

var Resources map[string]*tfjson.Schema
var DataSources map[string]*tfjson.Schema

func init() {
	resources := make(map[string]*tfjson.Schema, 0)
	dataSources := make(map[string]*tfjson.Schema, 0)  
	resources["tls_cert_request"] = resource.TlsCertRequestSchema()  
	resources["tls_locally_signed_cert"] = resource.TlsLocallySignedCertSchema()  
	resources["tls_private_key"] = resource.TlsPrivateKeySchema()  
	resources["tls_self_signed_cert"] = resource.TlsSelfSignedCertSchema()  
	dataSources["tls_certificate"] = data.TlsCertificateSchema()  
	dataSources["tls_public_key"] = data.TlsPublicKeySchema()  
	Resources = resources
	DataSources = dataSources
}