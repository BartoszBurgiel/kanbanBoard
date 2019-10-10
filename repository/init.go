package repository

// Querries to setup the database

const initColumns = `CREATE TABLE IF NOT EXISTS 'columns' (
							'name' 		VARCHAR(16),
							'columnID'  VARCHAR(16) PRIMARY KEY
							) ;`

const newTickets = `CREATE TABLE IF NOT EXISTS 'tickets' (
						'title'		VARCHAR(64),
						'desc'  	VARCHAR(256), 
						'deadline' 	VARCHAR(16),
						'priority'  INT,
						'id'    	VARCHAR(16) PRIMARY KEY,
						'columnID'  VARCHAR(16) 
						) ;`
