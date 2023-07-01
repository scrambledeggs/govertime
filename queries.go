package main

const CreateTableQuery string = `
  CREATE TABLE IF NOT EXISTS overtime (
    id INTEGER NOT NULL PRIMARY KEY,
    name TEXT NOT NULL,
    time_in TEXT NOT NULL,
    time_out TEXT NOT NULL,
    hours_ot INTEGER NOT NULL,
    reason TEXT NOT NULL
  );
`
const InsertOvertimeQuery string = `
INSERT INTO overtime VALUES(NULL, ?, ?, ?, ?, ?)
`
