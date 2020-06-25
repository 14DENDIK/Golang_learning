package todo_list

import (
  "fmt"
  "time"
  "log"
  "database/sql"
	_ "github.com/lib/pq"
)

const (
  host     = "localhost"
	port     = 5432
	user     = "sardor"
	password = "1"
	dbname   = "gopsql"
)

var (
  db *sql.DB
  err error
  psqlInfo string = fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
)

type Todo struct {
  Id uint64
  Task string
  Finished bool
  CreatedAt time.Time
}

func Initialize() {
    db, err = sql.Open("postgres", psqlInfo)
    if err != nil {
      log.Fatal(err)
    }
}

func AddToDatabase(t *Todo) *Todo {
  err = db.QueryRow("INSERT INTO todos(task, finished, created_at) VALUES($1, $2, $3) returning id;", t.Task, t.Finished, t.CreatedAt).Scan(&t.Id)
  if err != nil {
    log.Fatal(err)
  }
  return t
}

func GetFromDatabase(id int64) *Todo {
  var t Todo
  err = db.QueryRow("SELECT * FROM todos WHERE id=$1", id).Scan(&t.Id, &t.Task, &t.Finished, &t.CreatedAt)
  if err != nil {
    log.Fatal(err)
  }
  return &t
}

func DeleteFromDatabase(id int64) int64 {
  stmt, err := db.Prepare("DELETE FROM todos WHERE id=$1;")
  if err != nil {
    log.Fatal(err)
  }
  res, err := stmt.Exec(id)
  if err != nil {
    log.Fatal(err)
  }
  affect, err := res.RowsAffected()
  if err != nil {
    log.Fatal(err)
  }
  return affect
}

func (t *Todo) UpdateDatabase(task string, finished bool, createdAt time.Time) bool {
  stmt, err := db.Prepare("UPDATE todos SET task=$1, finished=$2, created_at=$3 WHERE id=$4;")
  if err != nil {
    log.Fatal(err)
  }
  res, err := stmt.Exec(task, finished, createdAt, t.Id)
  if err != nil {
    log.Fatal(err)
  }
  affect, err := res.RowsAffected()
  if err != nil {
    log.Fatal(err)
  }
  if affect == 1 {
    t.Task = task
    t.Finished = finished
    t.CreatedAt = createdAt
    return true
  }
  return false
}

// func main() {
//   Initialize()
//   deleted := DeleteFromDatabase(61)
//   if deleted < 1 {
//     panic("Could not delete")
//   }
//   fmt.Println("Deleted: ", deleted)
//   // fmt.Println(t)
//   // if updated := t.UpdateDatabase("Work", true, time.Now()); !updated {
//   //   panic("Could not update")
//   // }
//   // t := GetFromDatabase(8)
//
// }
