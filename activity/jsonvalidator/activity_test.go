package jsonvalidator

import (
	"io/ioutil"
	"testing"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

func TestJsonSchema(t *testing.T) {
	a := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	var jsonSchema = `
	{
		"$schema": "http://json-schema.org/draft-04/schema#",
		"type": "object",
		"properties": {
			"data": {"type": "string"}
		},
		"required": ["data"]
	}
	`

	var jsonDoc = `
	{
		"data": "this is data"
	}
	`

	tc.SetInput("document", jsonDoc)
	tc.SetInput("schema", jsonSchema)

	a.Eval(tc)
	val := tc.GetOutput("valid")

	if val.(bool) != true {
		t.Errorf("JSON should be valid")
	}

}
