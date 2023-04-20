package schema

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/iancoleman/strcase"
	"go/format"
	"strings"
	"text/template"
)

type Package string

const PackageResource = "resource"
const PackageData = "data"

const funcTemplate = `package {{ .Package }}  
  
import (  
"encoding/json"  
  
tfjson "github.com/hashicorp/terraform-json"  
)  
  
const {{.ResourceName}} = ` + "`{{.ResourceSchema}}`" + `
  
func {{.ResourceNameCamel}}Schema() *tfjson.Schema {  
	var result tfjson.Schema  
	_ = json.Unmarshal([]byte({{.ResourceName}}), &result)  
	return &result  
}  
`

const testTemplate = `package {{ .Package }}_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/{{ .RepoOwner }}/{{ .GoModule }}/v4/generated/{{ .Package }}"
	"github.com/stretchr/testify/assert"
)

func Test{{.ResourceNameCamel}}Schema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := {{ .Package }}.{{.ResourceNameCamel}}Schema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
`

type TemplateData struct {
	Package           string
	ResourceName      string
	ResourceSchema    string
	ResourceNameCamel string
	RepoOwner         string
	GoModule          string
}

func GenerateGoFileContent(blockName, schema string, pkg Package) (string, error) {
	return generateGoCode(blockName, funcTemplate, &schema, pkg)
}

func GenerateGoTestFileContent(blockName string, pkg Package) (string, error) {
	return generateGoCode(blockName, testTemplate, nil, pkg)
}

func generateGoCode(blockName, fileTemplate string, schema *string, pkg Package) (string, error) {
	parameter, err := buildTemplate(blockName, schema, pkg)
	if err != nil {
		return "", err
	}

	tmpl, err := template.New("goFileContent").Parse(fileTemplate)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, parameter); err != nil {
		return "", fmt.Errorf("error executing template: %w", err)
	}

	formattedSource, err := format.Source(buf.Bytes())
	if err != nil {
		return "", fmt.Errorf("error formatting generated Go file content: %w", err)
	}

	return string(formattedSource), nil
}

func buildTemplate(blockName string, schema *string, pkg Package) (*TemplateData, error) {
	var escapedSchema string
	if schema != nil {
		var schemaMap map[string]interface{}
		if err := json.Unmarshal([]byte(*schema), &schemaMap); err != nil {
			return nil, fmt.Errorf("error unmarshalling schema JSON: %s", err)
		}
		// Marshal the schema map back to a well-formatted JSON string
		formattedSchemaBytes, err := json.MarshalIndent(schemaMap, "", "  ")
		if err != nil {
			return nil, fmt.Errorf("error marshalling schema map: %w", err)
		}
		formattedSchema := string(formattedSchemaBytes)
		escapedSchema = strings.ReplaceAll(formattedSchema, "`", "` + \"`\" + `")
	}
	data := TemplateData{
		Package:           string(pkg),
		ResourceName:      strcase.ToLowerCamel(blockName),
		ResourceSchema:    escapedSchema,
		ResourceNameCamel: strcase.ToCamel(blockName),
		RepoOwner:         "lonegunmanb",
		GoModule:          "terraform-tls-schema",
	}
	return &data, nil
}
