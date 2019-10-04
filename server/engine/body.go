package engine

const body = `<!DOCTYPE html>
<html>
	<head>
		<title>My kanban board</title>
		<link rel="stylesheet" href="/css/style.css" />
		<meta content="width=device-width, initial-scale=1" name="viewport" />
	</head>

<body>
	<h1 id="header">My kanban board</h1>

	<div class="main">

	<div class="board-body">
		<div class="board-body-column">

		<div class="board-header">Todo</div>
		{{range .ToDo}}

			<div class="ticket">
				<div class="ticket-header">{{.Title}}</div>
				<div class="ticket-desc">{{.Description}}</div>
				<form action="/" method="POST">
					<input type="hidden" id="custId" name="ticketID" value="{{.ID}}">
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
					<input type="hidden" id="custId" name="ticketID" value="{{.ID}}">
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
						<input type="hidden" id="custId" name="ticketID" value="{{.ID}}">
						<input class="ticket-button" type="submit" name='{{.ID}}' value="X" />
					</form>

			</div>

		{{end}}

		</div>

	<!-- close board-body -->
	</div>

<!-- close main -->
</div>
</body>
</html>`
