package api

import (
	"net/http"

	"github.com/carlosclavijo/loginsolid/internal/handlers"
)

func router() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /users", handlers.Repo.PostUser)
	mux.HandleFunc("POST /login", handlers.Repo.LoginUser)
	mux.HandleFunc("OPTIONS /users", handlers.Repo.ChangeEncryption)
	mux.HandleFunc("OPTIONS /login", handlers.Repo.ChangeLogging)

	return mux
}
