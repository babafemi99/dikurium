package Repository

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"test-dikurium/graph/model"
)

type GormRepo struct {
	Client   *gorm.DB
	Database string
	User     string
	Password string
	Host     string
	Port     string
}

func NewGormRepo(user, password, host, port, database string) *GormRepo {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, database)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}

	err = db.AutoMigrate(&model.User{})
	if err != nil {
		panic(fmt.Sprintf("failed to setup database user: %v", err))
	}

	err = db.AutoMigrate(&model.Todo{})
	if err != nil {
		panic(fmt.Sprintf("failed to setup database todo: %v", err))
	}

	return &GormRepo{Client: db}
}

func (p GormRepo) PersistUser(user model.User) error {
	// check if user exists
	_, err := p.GetByEmail(user.Email)
	if err == nil {
		return errors.New("user already exists")
	}
	return p.Client.Create(user).Error
}

func (p GormRepo) GetByEmail(email string) (*model.User, error) {
	var User model.User
	if err := p.Client.First(&model.User{}, "email = ?", email).
		Scan(&User).Error; err != nil {
		return nil, err
	}
	return &User, nil
}

func (p GormRepo) CompleteTodo(id string) (*model.Todo, error) {
	// check if it exists
	err := p.Client.First(&model.Todo{}, "id = ? AND done = ?", id, false).Error
	if err != nil {
		return nil, err
	}

	// update
	err = p.Client.Model(&model.Todo{}).Where("id = ?, id").Update("done", true).Error
	if err != nil {
		return nil, err
	}

	// return new version
	var todo model.Todo
	if err = p.Client.First(&model.Todo{}).Where("id = ?, id").
		Scan(&todo).Error; err != nil {
		return nil, err
	}
	return &todo, nil

}

func (p GormRepo) DeleteTodo(id string) (bool, error) {
	if err := p.Client.Delete(&model.Todo{}, id).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (p GormRepo) PersistTodo(todo model.Todo) error {
	return p.Client.Create(&todo).Error
}

func (p GormRepo) GetAllTodo(ctx context.Context) ([]*model.Todo, error) {
	var todos []*model.Todo
	err := p.Client.Find(&todos).Error
	if err != nil {
		return nil, err
	}

	return todos, nil
}
