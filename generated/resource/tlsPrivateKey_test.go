package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-tls-schema/v4/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestTlsPrivateKeySchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.TlsPrivateKeySchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
