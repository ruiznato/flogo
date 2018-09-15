package uuid

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/google/uuid"
)

var log = logger.GetLogger("activity-uuid-generator")

type UUIDActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new GPIOActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &UUIDActivity{metadata: metadata}
}

// Metadata returns the activity's metadata
func (a *UUIDActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements api.Activity.Eval - Invokes a REST Operation
func (a *UUIDActivity) Eval(context activity.Context) (done bool, err error) {
	u := uuid.New()
	logger.Debug("Generated uuid: %s", u.String())
	context.SetOutput("output", u.String())
	return true, nil
}
