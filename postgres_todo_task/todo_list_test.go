package todo_list

import (
  "testing"
  "time"
)


func TestAddToDatabase(t *testing.T) {
  Initialize()
  var want uint64 = 0
  todo := &Todo{Task: "Run", Finished: false, CreatedAt: time.Now()}
  if got := AddToDatabase(todo); got.Id == want {
    t.Errorf("Todo List Should Contain an ID value")
  }
}

func TestUpdateDatabase(t *testing.T) {
  Initialize()
  todo := &Todo{Task: "Run", Finished: false, CreatedAt: time.Now()}
  todo = AddToDatabase(todo)
  if got := todo.UpdateDatabase("Jump", true, time.Now()); !got {
    t.Errorf("Could not update")
  }
}

func TestGetFromDatabase(t *testing.T) {
  Initialize()
  todo := GetFromDatabase(1)
  if todo == nil {
    t.Errorf("Could not get from database")
  }
}

func TestDeleteFromDatabase(t *testing.T) {
  Initialize()
  if deleted := DeleteFromDatabase(65); deleted < 1 {
    t.Errorf("Could not delete")
  }
}
