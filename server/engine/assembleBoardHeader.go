package engine

// AssembleBoardHeader returns a string with
// the board header template for index.html
func AssembleBoardHeader() string {
	return `

    <body>
        <h1 id="header">My kanban board</h1>
        <div class="main">
            <div class="board-header">
                <div>Todo</div>
                <div>InProgress</div>
                <div>Done</div>
            </div>
`
}
