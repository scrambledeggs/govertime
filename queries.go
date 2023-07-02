package main

const CreateTableQuery string = `
  CREATE TABLE IF NOT EXISTS overtimes (
    id INTEGER NOT NULL PRIMARY KEY,
    name TEXT NOT NULL,
    time_in TEXT NOT NULL,
    time_out TEXT NOT NULL,
    hours_ot INTEGER NOT NULL,
    reason TEXT NOT NULL
  );
`
const InsertOvertimeQuery string = `
INSERT INTO overtimes VALUES(NULL, ?, ?, ?, ?, ?)
`

const ViewMonthOvertimeQuery string = `SELECT * from overtimes WHERE time_in BETWEEN DATE('now','start of month') AND DATE('now', 'start of month', '+1 month', '-1 day') ORDER BY time_in ASC`

const ViewMonthGetDatThirtyBroOvertimeQuery string = `SELECT * from overtimes WHERE time_in BETWEEN DATE('now','start of month', '-2 day') AND DATE('now', 'start of month', '+1 month', '-1 day') ORDER BY time_in ASC`
