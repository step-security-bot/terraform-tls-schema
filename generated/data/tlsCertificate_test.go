package data_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-tls-schema/v4/generated/data"
	"github.com/stretchr/testify/assert"
)

func TestTlsCertificateSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := data.TlsCertificateSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
