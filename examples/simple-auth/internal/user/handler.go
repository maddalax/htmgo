package user

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/service"
	"simpleauth/internal/db"
)

type CreateUserRequest struct {
	Email    string
	Password string
}

type LoginUserRequest struct {
	Email    string
	Password string
}

type CreatedUser struct {
	Id    string
	Email string
}

func Create(ctx *h.RequestContext, request CreateUserRequest) (int64, error) {
	if len(request.Password) < 6 {
		return 0, errors.New("password must be at least 6 characters long")
	}

	queries := service.Get[db.Queries](ctx.ServiceLocator())

	hashedPassword, err := HashPassword(request.Password)

	if err != nil {
		return 0, errors.New("something went wrong")
	}

	id, err := queries.CreateUser(context.Background(), db.CreateUserParams{
		Email:    request.Email,
		Password: hashedPassword,
	})

	if err != nil {

		if err.Error() == "UNIQUE constraint failed: user.email" {
			return 0, errors.New("email already exists")
		}

		return 0, err
	}

	return id, nil
}

func Login(ctx *h.RequestContext, request LoginUserRequest) (int64, error) {

	queries := service.Get[db.Queries](ctx.ServiceLocator())

	user, err := queries.GetUserByEmail(context.Background(), request.Email)

	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		return 0, errors.New("email or password is incorrect")
	}

	if !PasswordMatches(request.Password, user.Password) {
		return 0, errors.New("email or password is incorrect")
	}

	session, err := CreateSession(ctx, user.ID)

	if err != nil {
		return 0, errors.New("something went wrong")
	}

	WriteSessionCookie(ctx, session)

	return user.ID, nil
}

func ParseMeta(meta any) map[string]interface{} {
	if meta == nil {
		return map[string]interface{}{}
	}
	if m, ok := meta.(string); ok {
		var dest map[string]interface{}
		json.Unmarshal([]byte(m), &dest)
		return dest
	}
	return meta.(map[string]interface{})
}

func GetMetaKey(meta map[string]interface{}, key string) string {
	if val, ok := meta[key]; ok {
		return val.(string)
	}
	return ""
}

func SetMeta(ctx *h.RequestContext, userId int64, meta map[string]interface{}) error {
	queries := service.Get[db.Queries](ctx.ServiceLocator())
	serialized, _ := json.Marshal(meta)
	fmt.Printf("serialized: %s\n", string(serialized))
	err := queries.UpdateUserMetadata(context.Background(), db.UpdateUserMetadataParams{
		JsonPatch: serialized,
		ID:        userId,
	})
	if err != nil {
		return err
	}
	return nil
}
