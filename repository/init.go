package repository

// Querries to setup the database

const initColumns = `CREATE TABLE IF NOT EXISTS 'states' (
							'name' 		VARCHAR(16),
							'stateID'  VARCHAR(16) PRIMARY KEY
							) ;`

const newTickets = `CREATE TABLE IF NOT EXISTS 'tickets' (
						'title'		VARCHAR(64),
						'desc'  	VARCHAR(256), 
						'deadline' 	VARCHAR(16),
						'priority'  INT,
						'id'    	VARCHAR(16) PRIMARY KEY,
						'stateID'  VARCHAR(16) 
						) ;`

const basicColumns = `INSERT INTO states VALUES (
							"ToDo", 
							"todo"	
						) ; 

						INSERT INTO states VALUES (
							"InProgress", 
							"inprogress"	
						) ;

						INSERT INTO states VALUES (
							"Done", 
							"done"	
						) ;
						`
