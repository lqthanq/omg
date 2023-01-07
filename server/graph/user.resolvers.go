package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/lqthanq/omg-simpler/app"
	"github.com/lqthanq/omg-simpler/graph/model"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	var userExists *model.User
	app.DB.Where("email = ?", strings.ToLower(input.Email)).Take(&userExists)

	if userExists != nil {
		return nil, errors.New("email already exists")
	}

	hasPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
	if err != nil {
		return nil, err
	}

	var user = model.User{
		Email:    input.Email,
		Username: input.Username,
		Password: string(hasPassword),
	}

	if input.Age != nil {
		user.Age = input.Age
	}

	if input.Gender != nil {
		user.Gender = input.Gender
	}

	if input.Phone != nil {
		user.Phone = input.Phone
	}

	if input.Address != nil {
		user.Address = input.Address
	}

	if err := app.DB.Save(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUser) (*model.User, error) {
	panic(fmt.Errorf("not implemented: UpdateUser - updateUser"))
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id int) (*bool, error) {
	panic(fmt.Errorf("not implemented: DeleteUser - deleteUser"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, filter model.UserFilter) (*model.UserConnection, error) {
	panic(fmt.Errorf("not implemented: Users - users"))
}

// UserByID is the resolver for the userByID field.
func (r *queryResolver) UserByID(ctx context.Context, id int) (*model.User, error) {
	panic(fmt.Errorf("not implemented: UserByID - userByID"))
}