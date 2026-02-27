package model_interface

import (
	"ticket-booking/models/domain/ticket"
)

type ITicketRepository interface {
	IBaseRepository[ticket.Ticket, ticket.Ticket]
}
