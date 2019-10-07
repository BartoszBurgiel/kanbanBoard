CREATE TABLE 'todotasks' (
    'title' VARCHAR(64), 
    'desc'  VARCHAR(256), 
    'ID'    VARCHAR(64)
)
;

CREATE TABLE 'inprogresstasks' (
    'title' VARCHAR(64), 
    'desc'  VARCHAR(256), 
    'ID'    VARCHAR(64)
)
;

CREATE TABLE 'donetasks' (
    'title' VARCHAR(64), 
    'desc'  VARCHAR(256), 
    'ID'    VARCHAR(64)
)
;

INSERT INTO todotasks 
VALUES (
    'first', 
    'desc', 
    '1'
)
;

INSERT INTO inprogresstasks 
VALUES (
    'sec', 
    'desc', 
    '2'
)
;

INSERT INTO donetasks 
VALUES (
    'third', 
    'desc', 
    '3'
)
;