// этот модуль является точкой входа для остальных модулей, подключает rest обработчики и является местом
// регистрации новых пользователей

package forum

import (
	"fmt"
	"log"
	"net/http"

	_ "forumless/docs"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"

	"forumless/app/models"
	"forumless/app/repo"
	"forumless/app/user"
)

type Forum struct {
	Host string
	Name string
	Port string

	repo repo.Forumer
	user *user.User
}

func New(host, name, port string, repo repo.Forumer, user *user.User) *Forum {
	return &Forum{host, name, port, repo, user}
}

// Run ...
// @title Swagger forumless API
// @version 1.0
// @description This is a headless forum
// @termsOfService http://swagger.io/terms/
// @contact.name Ivan
// @contact.url http://t.me/fess932
// @contact.email fess932@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost
// @BasePath /v2
func (f Forum) Run() {
	r := chi.NewRouter()

	// swagger
	{
		r.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL(fmt.Sprintf("http://localhost:%s/swagger/doc.json", f.Port))))
	}

	// forum
	{

		r.Get("/", f.MainHandler)
		r.Post("/post", f.CreatePostHandler)
	}

	// user
	{
		r.Post("/user", f.user.CreateUserHandler)
	}

	host := fmt.Sprintf(":%s", f.Port)

	log.Printf("forum %s started at http://0.0.0.0:%s", f.Name, f.Port)
	log.Fatal(http.ListenAndServe(host, r))
}

func (f Forum) CreatePost(u models.User, p models.Post) error {
	return f.repo.CreatePost(u, p)
}

func (f Forum) CreateComment(u models.User, p models.Post, c models.Comment) error {
	return f.repo.CreateComment(u, p, c)
}
