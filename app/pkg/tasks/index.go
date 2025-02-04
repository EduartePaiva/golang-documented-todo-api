package tasks

import (
	"context"
	"fmt"
	"sync"

	"github.com/golang-documented-todo-api/app/datasources/db"
	"github.com/golang-documented-todo-api/app/repository"
	"github.com/jackc/pgx/v5/pgtype"
)

type todosWithoutUserID struct {
	ID        pgtype.UUID      `json:"id"`
	TodoText  string           `json:"text"`
	Done      pgtype.Bool      `json:"done"`
	CreatedAt pgtype.Timestamp `json:"updatedAt"`
	UpdatedAt pgtype.Timestamp `json:"createdAt"`
}

// this function will fetch the database and prepare the data to send to users
func GetTasksForUser(service db.TasksServices, userID pgtype.UUID, ctx context.Context) ([]todosWithoutUserID, error) {
	todos, err := service.SelectAllTasksFromUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	sanitizedTodos := make([]todosWithoutUserID, 0, (len(todos)))
	for i := 0; i < len(todos); i++ {
		sanitizedTodos = append(sanitizedTodos, todosWithoutUserID{
			ID:        todos[i].ID,
			TodoText:  todos[i].TodoText,
			Done:      todos[i].Done,
			CreatedAt: todos[i].CreatedAt,
			UpdatedAt: todos[i].UpdatedAt,
		})
	}
	return sanitizedTodos, nil
}

func postTask(
	service db.TasksServices,
	ctx context.Context,
	wg *sync.WaitGroup,
	task repository.PostTaskParams,
) {
	defer wg.Done()
	err := service.PostTask(ctx, task)
	if err != nil {
		fmt.Printf("error while inserting value: %+v\n", task)
		fmt.Printf("error message: %v\n", err)
	}
}

func PostTasks(
	service db.TasksServices,
	tasks []incomingPostData,
	userID pgtype.UUID,
	ctx context.Context,
) {
	var wg sync.WaitGroup
	for _, task := range tasks {
		wg.Add(1)
		go postTask(service, ctx, &wg, repository.PostTaskParams{
			ID:          task.ID,
			UserID:      userID,
			TodoText:    task.Text,
			Done:        task.Done,
			UpdatedAt:   pgtype.Timestamp{Time: task.UpdatedAt, Valid: true},
			TodoText_2:  task.Text,
			Done_2:      task.Done,
			UpdatedAt_2: pgtype.Timestamp{Time: task.UpdatedAt, Valid: true},
		})
	}
	wg.Wait()
}
