package uuid

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/google/uuid"
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

func TestUUID(t *testing.T) {
	a := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	a.Eval(tc)
	val := tc.GetOutput("output")

	uuid, err := uuid.Parse(fmt.Sprintf("%v", val))
	if err != nil {
		t.Error("UUID cannot be parsed")
	}

	if v := uuid.Version(); v != 4 {
		t.Errorf("Random UUID of version %s", v)
	}

	fmt.Printf("UUID Generated: %s", uuid.String())
}
