package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/matkinhig/go-topup/config"
	"github.com/matkinhig/go-topup/router"
	_ "github.com/mattn/go-oci8"
)

func Run() {
	config.Load()
	fmt.Printf(" \n\t Listening [::]:%d \n", config.PORT)
	listen(config.PORT)
}

func listen(port int) {
	r := router.New()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
