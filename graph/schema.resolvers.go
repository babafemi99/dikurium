package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.30

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"log"
	"test-dikurium/graph/model"
)

// SignUp is the resolver for the signUp field.
func (r *mutationResolver) SignUp(ctx context.Context, username string, email string, password string) (*model.SignUpResult, error) {
	// TODO create validate email and username function

	// hash password
	hashedPassword, err := r.CryptoService.HashPassword(password)
	if err != nil {
		r.Logger.Errorw("failed to hash password", "error", err)
		return nil, err
	}
	// save to DB
	user := model.User{
		Userid:   uuid.New().String(),
		Username: username,
		Email:    email,
		Password: hashedPassword,
	}
	err = r.Repo.PersistUser(user)
	if err != nil {
		r.Logger.Errorw("failed to create user", "error", err)
		return nil, err
	}

	r.Logger.Infow("User successfully created:", "username", user.Username, "email", user.Email)

	// generate Token
	token, err := r.TokenService.CreateToken(user)
	if err != nil {
		r.Logger.Errorw("failed to generate token", "error", err)
		return nil, err
	}
	// serve Response
	return &model.SignUpResult{
		Success:   true,
		Message:   "Successfully created user ",
		User:      &user,
		AuthToken: &token,
	}, nil
}

// SignIn is the resolver for the signIn field.
func (r *mutationResolver) SignIn(ctx context.Context, email string, password string) (*model.SignInResult, error) {
	// find by email
	user, err := r.Repo.GetByEmail(email)
	if err != nil {
		r.Logger.Errorw("failed to get user by email", "error", err)
		return nil, errors.New("user doesn't exists")
	}
	// check if passwords match
	err = r.CryptoService.ComparePassword(user.Password, password)
	if err != nil {
		r.Logger.Errorw("failed to compare password", "error", err)
		return nil, errors.New("invalid password")
	}

	// generate Token
	token, err := r.TokenService.CreateToken(*user)
	if err != nil {
		r.Logger.Errorw("failed to generate token", "error", err)
		return nil, err
	}
	r.Logger.Infow("User successfully signed in", "username", user.Username)
	// Serve Response
	return &model.SignInResult{
		Success:   true,
		Message:   "Successfully signed in",
		User:      user,
		AuthToken: &token,
	}, nil
}

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, text string) (*model.Todo, error) {
	// get userId from middleware
	email, ok := ctx.Value("email").(string)
	if !ok {
		r.Logger.Error("Failed to get user from context")
		return nil, errors.New("unauthenticated")
	}

	user, err := r.Repo.GetByEmail(email)
	if err != nil {
		return nil, err
	}

	// validate text inputs
	if text == "" {
		r.Logger.Error("Todo Input is empty")
		return nil, errors.New("todo cannot be empty")
	}

	// generate UUID for tasks
	Id := uuid.New().String()
	data := model.Todo{
		ID:    Id,
		Text:  text,
		Done:  false,
		Email: email,
		User: &model.User{
			Username: user.Username,
			Email:    "aa@aa.com",
		},
	}
	r.Repo.PersistTodo(data)

	r.Logger.Infow("created todo successfully", "user", user.Username)

	// serve response
	log.Println(user.Email)
	return &data, nil
}

// MarkComplete is the resolver for the markComplete field.
func (r *mutationResolver) MarkComplete(ctx context.Context, todoID string) (*model.Todo, error) {
	username, ok := ctx.Value("username").(string)
	if !ok {
		r.Logger.Error("Unable to get user from context")
		return nil, errors.New("unauthenticated")
	}
	todo, err := r.Repo.CompleteTodo(todoID)
	if err != nil {
		r.Logger.Errorw("Unable to complete todo", "error", err)
		return nil, err
	}
	r.Logger.Infow("Todo Marked completed", "username", username)
	return todo, nil
}

// DeleteTodo is the resolver for the deleteTodo field.
func (r *mutationResolver) DeleteTodo(ctx context.Context, id string) (bool, error) {
	//panic(fmt.Errorf("not implemented: DeleteTodo - deleteTodo"))
	username, ok := ctx.Value("username").(string)
	if !ok {
		r.Logger.Error("Unable to get user from context")
		return false, errors.New("unauthenticated")
	}

	res, err := r.Repo.DeleteTodo(id)
	if err != nil {
		r.Logger.Errorw("Unable to delete todo", "error", err)
		return false, err
	}

	r.Logger.Infow("Todo Marked completed", "username", username)
	return res, nil
}

// Me is the resolver for the me field.
func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	email, ok := ctx.Value("email").(string)
	if !ok {
		r.Logger.Error("Unable to get user from context")
		return nil, errors.New("unauthenticated")
	}
	//return r.Repo.GetByEmail(email)
	user, err := GetUserFromCtx(ctx, email)
	if err != nil {
		log.Println("here ohh 1")
		log.Println("error is", err)
		return nil, err
	}
	return user, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	todos, err := r.Repo.GetAllTodo(ctx)
	if err != nil {
		return nil, err
	}
	for _, todo := range todos {
		log.Println(todo.Email)
		data, err := GetUserFromCtx(ctx, todo.Email)
		if err != nil {
			log.Println("error is ", err)
			return nil, err
		}
		todo.User = data

	}
	return todos, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
