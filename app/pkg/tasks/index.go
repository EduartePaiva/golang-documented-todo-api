package tasks

import (
	"context"

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

func PostTask(
	service db.TasksServices,
	data incomingPostData,
	userID pgtype.UUID,
	ctx context.Context,
) error {
	return service.PostTask(ctx, repository.PostTaskParams{
		ID:          data.ID,
		UserID:      userID,
		TodoText:    data.Text,
		Done:        data.Done,
		UpdatedAt:   pgtype.Timestamp{Time: data.UpdatedAt, Valid: true},
		TodoText_2:  data.Text,
		Done_2:      data.Done,
		UpdatedAt_2: pgtype.Timestamp{Time: data.UpdatedAt, Valid: true},
	})
}
