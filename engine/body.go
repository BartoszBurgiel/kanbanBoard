package engine

const body = `<!DOCTYPE html>
<html>
	<head>
	<title>My kanban board</title>
		<link rel="stylesheet" type="text/css" href="./style.css" />
		<meta content="width=device-width, initial-scale=1" name="viewport" />
	</head>

<body>
	<div class="header">My kanban board</div>

	<div class="main">

	<div class="board-body">
		<div class="board-body-column">

		<div class="board-header">Todo</div>
		{{range .ToDo}}

			<div class="ticket">
				<div class="ticket-header">{{.Title}}</div>
				<div class="ticket-desc">{{.Description}}</div>
				<form action="/" method="POST">
					<input type="hidden" name="ticketID" value="{{.ID}}">
					<input type="hidden" name="state" value="ToDo">
					<input class="ticket-button" type="submit" name='{{.ID}}' value="&rarr;" />
				</form>

			</div>

		{{end}}

		</div>
		<div class="board-body-column">
		<div class="board-header">InProgress</div>

		{{range .InProgress}}

			<div class="ticket">
				<div class="ticket-header">{{.Title}}</div>
				<div class="ticket-desc">{{.Description}}</div>
				<form action="/" method="POST">
					<input type="hidden" name="ticketID" value="{{.ID}}">
					<input type="hidden" name="state" value="InProgress">
					<input class="ticket-button" type="submit" name='{{.ID}}' value="&rarr;" />
				</form>

			</div>

		{{end}}

		</div>
		<div class="board-body-column">
		<div class="board-header">Done</div>

		{{range .Done}}

			<div class="ticket">
					<div class="ticket-header">{{.Title}}</div>
					<div class="ticket-desc">{{.Description}}</div>
					<form action="/" method="post">
						<input type="hidden" name="ticketID" value="{{.ID}}">
						<input type="hidden" name="state" value="Done">
						<input class="ticket-button" type="submit" name='{{.ID}}' value="X" />
					</form>

			</div>

		{{end}}

		</div>
	</div>
	<div class="new-ticket">
                <div class="new-ticket-head">New ticket...</div>
                <div class="new-ticket-body">
                    <form action="/" method="POST">
                        <div class="new-ticket-title">
                            <input type="text" name="newTitle" placeholder="Title..." required />
                        </div>

                        <div class="new-ticket-description">
                            <textarea name="newDescription" placeholder="Description..." required/></textarea>
                        </div>
                        
                        <input class="new-ticket-button" type="submit" name="submitNewTicket" value="Add" />
                    </form>
                </div>
            </div>
		</div>
	</body>
</html>`
