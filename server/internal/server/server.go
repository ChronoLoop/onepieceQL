package server

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/ikevinws/onepieceQL/internal/server/awsclient"
	"github.com/ikevinws/onepieceQL/internal/server/db"
	"github.com/ikevinws/onepieceQL/internal/server/models"
	"github.com/ikevinws/onepieceQL/internal/server/queries"
	"github.com/ikevinws/onepieceQL/pkg/utils"
	"github.com/joho/godotenv"
)

const baseURL = "https://onepieceql.up.railway.app"

func generateHTML(initialEndpoint string) []byte {
	html := `
            <!DOCTYPE html>
            <html lang="en">
            <body style="margin: 0; overflow-x: hidden; overflow-y: hidden">
            <div id="sandbox" style="height:100vh; width:100vw;"></div>
            <script src="https://embeddable-sandbox.cdn.apollographql.com/_latest/embeddable-sandbox.umd.production.min.js"></script>
            <script>
            new window.EmbeddedSandbox({
              target: "#sandbox",
              initialEndpoint: "` + initialEndpoint + `",
            });
            </script>
            </body>
            </html>
        `
	return []byte(html)
}

func generatePlaygroundHTML() []byte {
	serverEnv := os.Getenv("SERVER_ENV")
	if serverEnv == "development" {
		port := os.Getenv("PORT")
		return generateHTML("http://localhost:" + port + "/graphql")
	}
	return generateHTML(baseURL + "/graphql")
}

func StartServer() {
	exPath := utils.GetExePath()
	godotenv.Load(filepath.Join(exPath, "../.env"))
	db.Initialize()

	awsclient.ConfigS3()
	models.SeedDatabase()

	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:   "RootQuery",
			Fields: queries.GetRootFields(),
		}),
	}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		panic(err)
	}

	httpGraphQLHandler := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})

	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{baseURL},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
	}))

	r.Use(middleware.Logger)
	r.Use(httprate.LimitAll(250, time.Minute))

	r.Handle("/graphql", httpGraphQLHandler)
	r.Handle("/playground", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(generatePlaygroundHTML())
	}))

	// // Serve index.html for all other routes
	r.Get("/*", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		filePath := r.URL.Path
		if strings.HasPrefix(filePath, "/assets/") {
			w.Header().Set("Cache-Control", "public, max-age=31536000")
			staticDir := http.Dir(filepath.Join(exPath, "../client/dist"))
			http.FileServer(staticDir).ServeHTTP(w, r)
		} else {
			w.Header().Set("Cache-Control", "no-cache")
			http.ServeFile(w, r, filepath.Join(exPath, "../client/dist/index.html"))
		}
	}))

	http.ListenAndServe(":"+os.Getenv("PORT"), r)
}
