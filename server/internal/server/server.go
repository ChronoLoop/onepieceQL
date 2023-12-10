package server

import (
	"net/http"
	"os"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/ikevinws/onepieceQL/internal/server/awsclient"
	"github.com/ikevinws/onepieceQL/internal/server/db"
	"github.com/ikevinws/onepieceQL/internal/server/models"
	"github.com/ikevinws/onepieceQL/internal/server/queries"
	"github.com/joho/godotenv"
)

func generatePlaygroundHTML() []byte {
	port := os.Getenv("PORT")
	html := `
	<!DOCTYPE html>
	<html lang="en">
	<body style="margin: 0; overflow-x: hidden; overflow-y: hidden">
	<div id="sandbox" style="height:100vh; width:100vw;"></div>
	<script src="https://embeddable-sandbox.cdn.apollographql.com/_latest/embeddable-sandbox.umd.production.min.js"></script>
	<script>
	new window.EmbeddedSandbox({
	  target: "#sandbox",
	  initialEndpoint: "http://localhost:` + port + `/graphql",
	});
	</script>
	</body>
	</html>
`
	return []byte(html)
}

func StartServer() {
	godotenv.Load("../.env")
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

	httpHandler := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})

	http.Handle("/graphql", httpHandler)
	http.Handle("/playground", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(generatePlaygroundHTML())
	}))
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
