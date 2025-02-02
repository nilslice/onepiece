package planinfra

import (
	"encoding/json"
	"github.com/straw-hat-team/onepiece/go/onepiece"
	"github.com/straw-hat-team/onepiece/go/onepiece/eventsourcing"
	"github.com/straw-hat-team/onepiece/go/onepiece/protobuf"
	"unstable/plandomain/commands/archiveplan"
	"unstable/plandomain/commands/createplan"
	"unstable/plandomain/commands/drainplan"
	"unstable/plandomain/commands/faildrainplan"
	"unstable/plandomain/planproto"
)

var DispatchCreatePlan = eventsourcing.NewDecider(
	createplan.Decider,
	createPlanStreamID,
	marshalEvent,
	unmarshalEvent,
	eventTypeProvider,
)

var DispatchArchivePlan = eventsourcing.NewDecider(
	archiveplan.Decider,
	archivePlanStreamID,
	marshalEvent,
	unmarshalEvent,
	eventTypeProvider,
)

var DispatchDrainPlan = eventsourcing.NewDecider(
	drainplan.Decider,
	drainPlanStreamID,
	marshalEvent,
	unmarshalEvent,
	eventTypeProvider,
)

var DispatchFailDrainPlan = eventsourcing.NewDecider(
	faildrainplan.Decider,
	failPlanStreamID,
	marshalEvent,
	unmarshalEvent,
	eventTypeProvider,
)

func failPlanStreamID(command *planproto.FailDrainPlan) (string, error) {
	return protobuf.StreamID(command, command.PlanId), nil
}

func drainPlanStreamID(command *planproto.DrainPlan) (string, error) {
	return protobuf.StreamID(command, command.PlanId), nil
}

func archivePlanStreamID(command *planproto.ArchivePlan) (string, error) {
	return protobuf.StreamID(command, command.PlanId), nil
}

func createPlanStreamID(command *planproto.CreatePlan) (string, error) {
	return protobuf.StreamID(command, command.PlanId), nil
}

func unmarshalEvent(eventType string, data []byte) (*planproto.Event, error) {
	event := &planproto.Event{}
	err := json.Unmarshal(data, event)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func marshalEvent(event *planproto.Event) (eventsourcing.ContentType, []byte, error) {
	bytes, err := json.Marshal(event)
	return eventsourcing.ContentTypeJson, bytes, err
}

func eventTypeProvider(event *planproto.Event) (string, error) {
	switch e := event.Event.(type) {
	case *planproto.Event_PlanCreated:
		return protobuf.MessageFullName(e.PlanCreated).String(), nil
	case *planproto.Event_PlanArchived:
		return protobuf.MessageFullName(e.PlanArchived).String(), nil
	case *planproto.Event_PlanDrained:
		return protobuf.MessageFullName(e.PlanDrained).String(), nil
	case *planproto.Event_PlanDrainFailed:
		return protobuf.MessageFullName(e.PlanDrainFailed).String(), nil
	default:
		return "", onepiece.ErrUnknownEvent
	}
}
