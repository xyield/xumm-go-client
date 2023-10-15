package main

import (
	"log"
	"net/http"

	"github.com/xyield/xumm-go-client/utils"
	anyjson "github.com/xyield/xumm-go-client/utils/json"
	"github.com/xyield/xumm-go-client/xumm"
	"github.com/xyield/xumm-go-client/xumm/client"
	"github.com/xyield/xumm-go-client/xumm/models"
)

func main() {

	// Set API key and Secret with environment variables XUMM_API_KEY and XUMM_API_SECRET or manually by adding an argument.
	cfg, err := xumm.NewConfig()
	if err != nil {
		log.Panicln(err)
	}

	// Initialise new client with xumm config
	client := client.New(cfg)

	// Test connectivity with XUMM api
	pong, err := client.Meta.Ping()
	if err != nil {
		log.Panicln("Failed to connect to xumm api")
	}

	utils.PrettyPrintJson(pong)

	http.HandleFunc("/send-payment", func(w http.ResponseWriter, r *http.Request) {
		cp, err := client.Payload.PostPayload(models.XummPostPayload{
			TxJson: anyjson.AnyJson{
				"TransactionType": "Payment",
				"Account":         "rQNrSWi3t6ojNFof8gE3Wq8Pwz88QUr6Hx",
				"Amount":          "1",
				"Destination":     "rwietsevLFg8XSmG3bEZzFein1g8RBqWDZ",
				"Fee":             "12",
			},
		})
		if err != nil {
			log.Println("Failed to create payload", err)
			w.WriteHeader(400)
			return
		}
		log.Println("Sign request sent...")
		xp, err := client.Payload.Subscribe(cp.UUID)
		if err != nil {
			log.Println(err)
			w.WriteHeader(400)
			return
		}

		utils.PrettyPrintJson(xp)

		// Do something with signed/rejected payload
	})

	log.Println("Starting server...")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Panic(err)
	}
}
