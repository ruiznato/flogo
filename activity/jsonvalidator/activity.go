package jsonvalidator

import (
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
	jsonDoc, _ := context.GetInput("document").(string)
	jsonSchema, _ := context.GetInput("schema").(string)

	doc := gojsonschema.NewStringLoader(jsonDoc)
	schema := gojsonschema.NewStringLoader(jsonSchema)

	result, err := gojsonschema.Validate(schema, doc)
	if err != nil {
		return false, err
	}

	context.SetOutput("valid", result.Valid())
	return true, nil
}
