package repository

// Querries to setup the database

const initColumns = `CREATE TABLE IF NOT EXISTS 'states' (
						'name' 		VARCHAR(16),
						'stateID'  	VARCHAR(16),
						'limit'   	INT
						) ;`

const newTickets = `CREATE TABLE IF NOT EXISTS 'tickets' (
						'title'		VARCHAR(64),
						'desc'  	VARCHAR(256), 
						'deadline' 	VARCHAR(16),
						'priority'  INT,
						'id'    	VARCHAR(16),
						'stateID'  	VARCHAR(16) 
						) ;`
