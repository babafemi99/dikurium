package graph

import (
	"context"
	"gorm.io/gorm"
	"log"
	"net/http"
	"test-dikurium/graph/model"
	"time"
)

const userLoaderKey = "userLoader"

func DataloaderMiddleware(db *gorm.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		userloader := UserLoader{
			maxBatch: 100,
			wait:     1 * time.Millisecond,
			fetch: func(id []string) ([]*model.User, []error) {
				log.Println("id .0: ", id[0])
				var users []*model.User
				var user model.User
				if err := db.First(&model.User{}, "email = ?", id[0]).Scan(&user).Error; err != nil {
					log.Println("error finding email", err)
					return nil, []error{err}
				}
				log.Println("user email is ", user.Email)
				users = append(users, &user)
				return users, nil
			},
		}
		ctx := context.WithValue(request.Context(), userLoaderKey, &userloader)
		next.ServeHTTP(writer, request.WithContext(ctx))
	})
}

func GetUserLoader(ctx context.Context) *UserLoader {
	loader := ctx.Value(userLoaderKey).(*UserLoader)
	return loader
}
