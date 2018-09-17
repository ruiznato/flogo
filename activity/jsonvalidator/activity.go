package jsonvalidator

import (
	"encoding/json"
	"fmt"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/xeipuuv/gojsonschema"
)

var log = logger.GetLogger("activity-json-validator")

type JSONValidatorActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new GPIOActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &JSONValidatorActivity{metadata: metadata}
}

// Metadata returns the activity's metadata
func (a *JSONValidatorActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements api.Activity.Eval - Invokes a REST Operation
func (a *JSONValidatorActivity) Eval(context activity.Context) (done bool, err error) {
	jsonDoc := context.GetInput("document").(map[string]interface{})
	jsonSchema := context.GetInput("schema").(map[string]interface{})

	docString, _ := json.Marshal(jsonDoc)
	schemaString, _ := json.Marshal(jsonSchema)

	doc := gojsonschema.NewStringLoader(fmt.Sprintf("%s", docString))
	schema := gojsonschema.NewStringLoader(fmt.Sprintf("%s", schemaString))

	result, err := gojsonschema.Validate(schema, doc)
	if err != nil {
		log.Errorf("ERROR: %v", err)
		return false, err
	}

	if !result.Valid() {

		var errors []string
		for _, err := range result.Errors() {
			errors = append(errors, fmt.Sprintf("%s", err))
		}
		context.SetOutput("error", errors)
	}
	context.SetOutput("valid", result.Valid())
	return true, nil
}
