package tasks

import (
	"database/sql"
	"time"
)

const fileName = "./tasks.db"
const (
	schemaSql = `CREATE TABLE IF NOT EXISTS tasks (
           id INTEGER PRIMARY KEY AUTOINCREMENT,
           word VARCHAR(100) NOT NULL,
           translate VARCHAR(100) NOT NULL,
           frequency INTEGER(3) DEFAULT 0,
           is_active BOOLEAN DEFAULT true,
           time TIMESTAMP 
    );
    CREATE INDEX IF NOT EXISTS tasks_word ON tasks(word); 
    CREATE INDEX IF NOT EXISTS tasks_frequency ON tasks(frequency);
    CREATE INDEX IF NOT EXISTS tasks_isActive ON tasks(isActive); 
   `

	schemaInsert = `INSERT INTO tasks (word, translate, frequency, is_active, time)
                    VALUES(?,?,?,?,?)`
	schemaUpdate = ``
	schemaSelect = `SELECT * FROM tasks ORDER BY frequency DESC WHERE is_active = ?`
	schemaDelete = ``
)

//Task is structure for word/translate
type Task struct {
	Id        int
	Word      string
	Translate string
	Frequency int
	IsActive  bool
	Time      time.Time
}

//DB is a database of stock tasks
type DB struct {
	sql    *sql.DB
	stmt   *sql.Stmt
	buffer []Task
}

// NewDb constructs a Tasks value for managing stock tasks in a
// SQLite database. This API is not thread safe.
func NewDb() (*DB, error) {
	//_ = os.Remove(fileName)

	sqlDB, err := sql.Open("sqlite3", fileName)
	if err != nil {
		return nil, err
	}

	if _, err = sqlDB.Exec(schemaSql); err != nil {
		return nil, err
	}

	db := DB{
		sql:    sqlDB,
		buffer: make([]Task, 0, 1024),
	}

	return &db, nil
}

// Create stores a task into the buffer. Once the buffer is full, the
// tasks are flushed to the database.
func (db *DB) Create(task Task) error {
	defer db.sql.Close()

	stmt, err := db.sql.Prepare(schemaInsert)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = db.stmt.Exec(task)
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) Show(frequency int) ([]Task, error) {

	//get chicle for all elements
	stack := append(db.buffer, Task{})

	return stack, nil
}
