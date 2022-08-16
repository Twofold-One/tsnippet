package main

import (
	"net/http"

	"github.com/Twofold-One/tsnippet/ui"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	// custom NotFound function
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.NotFound(w)
	})

	fileServer := http.FileServer(http.FS(ui.Files))

	router.Handler(http.MethodGet, "/static/*filepath", fileServer)

	// Code for not embeded static files fileServer
	//
	// fileServer := http.FileServer(neuteredFileSystem{http.Dir("./ui/static/")})
	// router.Handler(http.MethodGet, "/static", http.NotFoundHandler())
	// router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))

	dynamic := alice.New(app.sessionManager.LoadAndSave, noSurf, app.authenticate)

	router.Handler(http.MethodGet, "/", dynamic.ThenFunc(app.home))
	router.Handler(http.MethodGet, "/snippet/view/:id", dynamic.ThenFunc(app.snippetView))
	router.Handler(http.MethodGet, "/user/login", dynamic.ThenFunc(app.userLogin))
	router.Handler(http.MethodGet, "/user/signup", dynamic.ThenFunc(app.userSignup))
	router.Handler(http.MethodPost, "/user/signup", dynamic.ThenFunc(app.userSignupPost))
	router.Handler(http.MethodPost, "/user/login", dynamic.ThenFunc(app.userLoginPost))

	protected := dynamic.Append(app.requireAuthentication)

	router.Handler(http.MethodGet, "/snippet/create", protected.ThenFunc(app.snippetCreate))
	router.Handler(http.MethodPost, "/snippet/create", protected.ThenFunc(app.snippetCreatePost))
	router.Handler(http.MethodPost, "/user/logout", protected.ThenFunc(app.userLogoutPost))

	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	return standard.Then(router)
}

// Code for not embeded static files fileServer
//
// neuteredFileSystem with restricted "/static" path.
// type neuteredFileSystem struct {
// 	fs http.FileSystem
// }

// func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
// 	f, err := nfs.fs.Open(path)
// 	if err != nil {
// 		return nil, err
// 	}

// 	s, err := f.Stat()
// 	if err != nil {
// 		return nil, err
// 	}
// 	if s.IsDir() {
// 		index := filepath.Join(path, "index.html")
// 		if _, err := nfs.fs.Open(index); err != nil {
// 			closeErr := f.Close()
// 			if closeErr != nil {
// 				return nil, closeErr
// 			}

// 			return nil, err
// 		}
// 	}
// 	return f, nil
// }
