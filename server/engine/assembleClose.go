package engine

// AssembleClose returns a string with
// the close template for index.html
func AssembleClose() string {
	return `

	<!-- close board-body -->
	</div>
	
	<!-- close main -->
	</div>
	</body>
	</html>

`
}
