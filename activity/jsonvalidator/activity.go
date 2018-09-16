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

	log.Infof("DOC: %v", jsonDoc)
	log.Infof("SCHEMA: %v", jsonSchema)
	doc := gojsonschema.NewStringLoader(jsonDoc)
	schema := gojsonschema.NewStringLoader(jsonSchema)

	log.Info("Loaded doc and schema")
	result, err := gojsonschema.Validate(schema, doc)
	if err != nil {
		log.Infof("ERROR: %v", err)
		return false, err
	}

	log.Infof("Doc validated: %v", result.Valid())
	context.SetOutput("valid", result.Valid())
	return true, nil
}
