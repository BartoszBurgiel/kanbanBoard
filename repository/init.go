package repository

// Query to setup the database
const initState = `CREATE TABLE IF NOT EXISTS 'tasks' (
						'title' 	VARCHAR(64),
						'desc'  	VARCHAR(256), 
						'state' 	VARCHAR(64),
						'deadline' 	VARCHAR(16),
						'priority'  INT,
						'id'    	VARCHAR(16) 
						) ;`
