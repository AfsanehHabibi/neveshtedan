package auth

import (
	"context"
	"net/http"

	"github.com/AfsanehHabibi/neveshtedan/pkg/jwt"
	"github.com/AfsanehHabibi/neveshtedan/pkg/repository"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func Middleware(userRepo repository.UserRepository) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")

			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			tokenStr := header
			userId, err := jwt.ParseToken(tokenStr)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}

			user, err := userRepo.GetById(r.Context(), userId)
			if err != nil || user == nil {
				next.ServeHTTP(w, r)
				return
			}
			ctx := AddUserToContext(r.Context(), userId)

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func GetUseFromContext(ctx context.Context) *int {
	raw, _ := ctx.Value(userCtxKey).(*int)
	return raw
}

func AddUserToContext(ctx context.Context, userId int) context.Context {
	return context.WithValue(ctx, userCtxKey, &userId)
}
