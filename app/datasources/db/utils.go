package db

import (
	"context"

	"github.com/golang-documented-todo-api/app/repository"
)

func GetOrCreateNewUserAndReturn(service UserServices, ctx context.Context, user repository.User) (repository.User, error) {
	currentUser, err := service.SelectUserFromProviderNameAndId(ctx, repository.SelectUserFromProviderNameAndIdParams{
		ProviderUserID: user.ProviderUserID,
		ProviderName:   user.ProviderName,
	})
	if err == nil {
		if user.AvatarUrl != currentUser.AvatarUrl {
			service.UpdateUserAvatarURL(ctx, repository.UpdateUserAvatarURLParams{
				AvatarUrl: user.AvatarUrl,
				ID:        user.ID,
			})
		}
		return currentUser, err
	}

	return currentUser, err
}
