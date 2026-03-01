package middleware

import (
	"log"
	"net/http"

	"ticket-booking/internal/auth"

	"github.com/casbin/casbin/v2"
)

func Authorizer(e *casbin.Enforcer, jwt *auth.JWTManager) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if r.URL.Path == "/api/register" {
				next.ServeHTTP(w, r)
				return
			}

			token, error := auth.GetTokenFromPayload(r)
			if error != nil {
				log.Fatal("middleware auth: ", error)
			}
			claims, err := jwt.Parse(token)

			// account, err := server.GetAccountFromToken(r)
			// account := user.Account{Role: user.Admin}
			// var err interface{} = nil // TODO solve that

			// if err != nil {
			// 	role = user.Anonymous
			// } else {
			// 	role = account.Role
			// }
			log.Printf("Enforcing: role=%s, path=%s, method=%s", claims.Role, r.URL.Path, r.Method)

			res, err := e.Enforce(string(claims.Role), r.URL.Path, r.Method)

			if err != nil {
				log.Fatal("middleware auth: ", err)
			}

			log.Println("res: ", res)
			if !res {
				w.WriteHeader(http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
