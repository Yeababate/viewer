package main
import (
	"context"
	//"fmt"
	"html/template"
	"log"
	"net/http"
	//"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var tpl *template.Template

func main(){
	var err error
	//use WaitGroup to wait for all goroutines fetching API data to complete before starting server
        var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		CarsData, err = getCars()
		if err != nil {
			log.Fatal(err)
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
        	Manufacturers, err = getManufacturers()
		if err != nil {
			log.Fatal(err)
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
        	Categories, err = getCategories()
		if err != nil {
			log.Fatal(err)
		}
	}()
	//wait until above functions complete before continuing
	wg.Wait()

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	tpl,_ = tpl.ParseGlob("templates/*html")

	http.HandleFunc("/", homeHandler) 
	http.HandleFunc("GET /car/{id}", carDetailHandler)
	http.HandleFunc("GET /filter", filterHandler)

	//ONLY FOR TESTING
	http.HandleFunc("GET /test", testHandler)
	
	// http.HandleFunc("GET /models/", allmodelsHandler)
	// http.HandleFunc("GET /models/{id}", modelsHandler)
	// http.HandleFunc("GET /manufacturers", allManufacturersHandler)
	// http.HandleFunc("GET /manufacturers/{id}", manufacturersHandler)
	// http.HandleFunc("GET /categories", allCategoriesHandler)
	// http.HandleFunc("GET /categories/{id}", categoriesHandler)

	server := &http.Server{
		Addr: ":8080",
		Handler: nil,
	}

	go func () {
		log.Println("Server is running")
		if err = server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	//signal.NotifyContext returns a copy of the parent context marked done when one of the listed signals arrives
	//signal.NotifyContext returns (context.Context, context.CancelFunc)
	shutdown, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	//wait for shutdown.Done() signal
	<-shutdown.Done()
	log.Println("Shutdown signal received")

	//create 30-second delay to finish ongoing tasks
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	//force-quit if more than 30 seconds pass
	if err = server.Shutdown(ctx); err != nil {
		log.Fatalf("Forced shutdown: %s", err)
	}
	
	//if no force quit, server shuts down once all ongoing tasks are finished
	log.Println("Server exited gracefully")

}
