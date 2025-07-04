package repository

import (
	"context"
	"thanhldt060802/infrastructure"
	"thanhldt060802/internal/model"
)

type userRepository struct {
}

type UserRepository interface {
	GetViewById(ctx context.Context, id string) (*model.UserView, error)

	GetById(ctx context.Context, id string) (*model.User, error)
	GetByUsername(ctx context.Context, username string) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	Create(ctx context.Context, newUser *model.User) error
	Update(ctx context.Context, updatedUser *model.User) error
	DeleteById(ctx context.Context, id string) error

	// Elasticsearch integrattion (init data for elasticsearch-service)
	GetAllViews(ctx context.Context) ([]*model.UserView, error)
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (userRepository *userRepository) GetViewById(ctx context.Context, id string) (*model.UserView, error) {
	user := new(model.UserView)

	query := infrastructure.PostgresDB.NewSelect().Model(user).Where("_user.id = ?", id)

	if err := query.Scan(ctx); err != nil {
		return nil, err
	}

	return user, nil
}

func (userRepository *userRepository) GetById(ctx context.Context, id string) (*model.User, error) {
	user := new(model.User)

	query := infrastructure.PostgresDB.NewSelect().Model(user).Where("id = ?", id)

	if err := query.Scan(ctx); err != nil {
		return nil, err
	}

	return user, nil
}

func (userRepository *userRepository) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	user := new(model.User)

	query := infrastructure.PostgresDB.NewSelect().Model(user).Where("username = ?", username)

	if err := query.Scan(ctx); err != nil {
		return nil, err
	}

	return user, nil
}

func (userRepository *userRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	user := new(model.User)

	query := infrastructure.PostgresDB.NewSelect().Model(user).Where("email = ?", email)

	if err := query.Scan(ctx); err != nil {
		return nil, err
	}

	return user, nil
}

func (userRepository *userRepository) Create(ctx context.Context, newUser *model.User) error {
	_, err := infrastructure.PostgresDB.NewInsert().Model(newUser).Returning("*").Exec(ctx)
	return err
}

func (userRepository *userRepository) Update(ctx context.Context, updatedUser *model.User) error {
	_, err := infrastructure.PostgresDB.NewUpdate().Model(updatedUser).Where("id = ?", updatedUser.Id).Exec(ctx)
	return err
}

func (userRepository *userRepository) DeleteById(ctx context.Context, id string) error {
	_, err := infrastructure.PostgresDB.NewDelete().Model(&model.User{}).Where("id = ?", id).Exec(ctx)
	return err
}

func (userRepository *userRepository) GetAllViews(ctx context.Context) ([]*model.UserView, error) {
	var users []*model.UserView

	query := infrastructure.PostgresDB.NewSelect().Model(&users)

	if err := query.Scan(ctx); err != nil {
		return nil, err
	}

	return users, nil
}
