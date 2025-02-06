package handlers

import (
	"context"

	"github.com/golang-documented-todo-api/app/repository"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/mock"
)

type taskServiceMock struct {
	mock.Mock
}

func (v *taskServiceMock) SelectAllTasksFromUser(ctx context.Context, userID pgtype.UUID) ([]repository.Todo, error) {
	return nil, nil
}

func (v *taskServiceMock) PostTask(ctx context.Context, arg repository.PostTaskParams) error {
	return nil
}
func (v *taskServiceMock) DeleteTaskByIDAndUserID(ctx context.Context, arg repository.DeleteTaskByIDAndUserIDParams) error {
	return nil
}

func (v *taskServiceMock) UpdateDoneAndTextFromTask(ctx context.Context, arg repository.UpdateDoneAndTextFromTaskParams) error {
	return nil
}

func (v *taskServiceMock) UpdateDoneFromTask(ctx context.Context, arg repository.UpdateDoneFromTaskParams) error {
	return nil
}

func (v *taskServiceMock) UpdateTextFromTask(ctx context.Context, arg repository.UpdateTextFromTaskParams) error {
	return nil
}
