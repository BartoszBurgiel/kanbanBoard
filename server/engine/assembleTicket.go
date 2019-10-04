package engine

// AssembleTicket returns a string with
// the ticket template for index.html
func AssembleTicket() string {
	return `

	<div class="board-body-column">
	
		{{range .ToDo}}    
		
			<div class="ticket">
				<div class="ticket-header">{{.Title}}</div>
				<div class="ticket-desc">{{.Description}}</div>                        
			</div>
	
		{{end}}
	
	</div>
	
`
}
