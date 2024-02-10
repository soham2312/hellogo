package main

import ("fmt"
"log"
"os"
"github.com/joho/godotenv"
"github.com/go-chi/chi"
"github.com/go-chi/cors"
"net/http"
"github.com/lib/pq"
"github.com/soham2312/hellogo/internal/database"
)

type apiConfig struct {
	DB *database.Queries
}

func main()  {
	godotenv.Load(".env")
	portString := os.Getenv("PORT")
	if portString == "" {
		fmt.Println("PORT is not set")
	}
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		fmt.Println("DB_URL is not set")
	}

	conn,err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Can't connect to the databse",err)
	}

	queries,err := database.New(conn)
	if err != nil {
		log.Fatal("Can't create connection to databse",err)
	}

	apiCfg := apiConfig{
		DB: queries,
	}

	router :=chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"htttps://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	}))

	v1 := chi.NewRouter()

	v1.Get("/healthz", handlerReadiness)
	v1.Get("/err",handleeErr)
	v1.Post("/users",apiCfg.handlerCreateUser)

	router.Mount("/v1", v1)

	srv := &http.Server{
		Handler: router,
		Addr: ":" + portString,
	}
	log.Printf("Server is running on port %v", portString)
	err := srv.ListenAndServe()
	if err != nil{
		log.Fatal(err)
	}
}