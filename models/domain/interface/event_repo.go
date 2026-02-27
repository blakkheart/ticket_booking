package model_interface

import (
	"ticket-booking/models/domain/event"
)

type IEventRepository interface {
	IBaseRepository[event.Event, event.Event]
}
