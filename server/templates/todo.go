package templates

// Ticket represents a ticket on the kanban board
type Ticket struct {
	Title       string
	Description string
}

// TicketList will be passed over to the HTML template
type TicketList []Ticket
