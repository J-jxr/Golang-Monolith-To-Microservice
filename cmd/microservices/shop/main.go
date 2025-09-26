package main
import (
	"log"
	"os"
	"fmt"
	"net/http"	
)

func main() {
	log.Println("Starting the shop microservice...")

	ctx := cmd.Context()

	r := createShopMicroservice()

	server := &http.Server{
		Addr:    os.Getenv("SHOP_PRODUCT_SERVICE_BIND_ADDR"),
		Handler: r,
	}

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed  {
			panic(err)
		}
	}()

	<-ctx.Done()

	log.Println("closing the shop microservice ...")

	if err := server.Close(); err != nil {
		panic(err)
	}
}

func createShopMicroservice() (router *chi.Mux) {


	
}
