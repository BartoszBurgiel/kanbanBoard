package repository

// Query to setup the database
const initState = `CREATE TABLE IF NOT EXISTS 'columns' (
							'name' 		VARCHAR(16),
							'columnID'  VARCHAR(16)
							) ;`

// Query create a new ticket
const newColumn = `CREATE TABLE IF NOT EXISTS 'ticket' (
						'title'		VARCHAR(64),
						'desc'  	VARCHAR(256), 
						'state' 	VARCHAR(64),
						'deadline' 	VARCHAR(16),
						'priority'  INT,
						'id'    	VARCHAR(16),
						'columnID'  VARCHAR(16)
						) ;`
