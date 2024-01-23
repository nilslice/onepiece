package planinfra

import (
	"github.com/straw-hat-team/onepiece/go/onepiece"
	"github.com/straw-hat-team/onepiece/go/onepiece/eventsourcing"
	"unstable/plandomain/planactor"
	"unstable/plandomain/planproto"
)

var DispatchCommand = eventsourcing.NewDecider(
	planactor.Decider,
	streamID,
	marshalEvent,
	unmarshalEvent,
	eventTypeProvider,
)

func streamID(command *planproto.Command) (string, error) {
	switch c := command.Command.(type) {
	case *planproto.Command_CreatePlan:
		return createPlanStreamID(c.CreatePlan)
	case *planproto.Command_ArchivePlan:
		return archivePlanStreamID(c.ArchivePlan)
	case *planproto.Command_DrainPlan:
		return drainPlanStreamID(c.DrainPlan)
	case *planproto.Command_FailDrainPlan:
		return failPlanStreamID(c.FailDrainPlan)
	default:
		return "", onepiece.ErrUnknownCommand
	}
}
