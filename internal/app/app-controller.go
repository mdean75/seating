package app

import (
	"seating/internal/handlers/attendeeadapter"
	"seating/internal/handlers/eventadapter"
	"seating/internal/handlers/groupadapter"
	"seating/internal/handlers/industryadapter"
)

type Controller struct {
	GroupHandler    *groupadapter.HTTPHandler
	EventHandler    *eventadapter.HTTPHandler
	AttendeeHandler *attendeeadapter.HTTPHandler
	Industryhandler *industryadapter.HTTPHandler
}

func NewController(group *groupadapter.HTTPHandler, event *eventadapter.HTTPHandler, attendee *attendeeadapter.HTTPHandler, industry *industryadapter.HTTPHandler) *Controller {
	return &Controller{
		GroupHandler:    group,
		EventHandler:    event,
		AttendeeHandler: attendee,
		Industryhandler: industry,
	}
}
