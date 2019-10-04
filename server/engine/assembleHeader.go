package engine

// AssembleHeader returns a string with
// the header template for index.html
func AssembleHeader() string {
	return `

	<!DOCTYPE html>
	<html>
		<head>
			<link rel="stylesheet" type="text/css" href="/css/style.css">
			<title>My own kanban boarr</title>
		</head>
	
`
}
