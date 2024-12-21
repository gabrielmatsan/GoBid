package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (api *Api) BindRoutes() {
	api.Router.Use(middleware.RequestID, middleware.Recoverer, middleware.Logger,
		api.Sessions.LoadAndSave)

	// 	csrfMiddleware := csrf.Protect([]byte(os.Getenv("GOBID_CSRF_KEY")),
	// 	csrf.Secure(false), // DEV ONLY
	// 	csrf.ErrorHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 		log.Printf("CSRF Token Received: %s", r.Header.Get("X-CSRF-Token"))
	// 		log.Printf("Cookies Received: %v", r.Cookies())
	// 		http.Error(w, "Forbidden - CSRF token invalid", http.StatusForbidden)
	// 	})),
	// )
	// api.Router.Use(csrfMiddleware)

	// !!!o de cima serve para debugar o csrf token!!!

	// csrfMiddleware := csrf.Protect([]byte(os.Getenv("GOBID_CSRF_KEY")),
	// 	csrf.Secure(false), // DEV ONLY
	// )
	// api.Router.Use(csrfMiddleware)

	api.Router.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {

			r.Get("/csrf-token", api.HandlerGetCsrfToken)

			r.Route("/users", func(r chi.Router) {
				r.Group(func(r chi.Router) {
					r.Use(api.AuthMiddleware)
					r.Post("/logout", api.handleLogOutUser)
				})
				r.Post("/signup", api.handleSignUpUser)
				r.Post("/login", api.handleLoginUser)
				//r.With(api.AuthMiddleware).Post("/logout", api.handleLogOutUser)
			})

			r.Route("/products", func(r chi.Router) {
				r.Group(func(r chi.Router) {
					r.Use(api.AuthMiddleware)
					r.Post("/", api.handleCreateProduct)
				})
			})
		})
	})
}
