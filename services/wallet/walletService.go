package wallet

import (
	"log"

	"github.com/dapr/go-sdk/client"
)

func StartWalletService() error {
	log.Println("Initialising Dapr Client ...")
	client, err := client.NewClient()
	if err != nil {
		return err
	}
	defer client.Close()
	log.Println("Dapr Client has been initialised.")

	// TODO Add wallet service implementaion
   

	return nil
}

