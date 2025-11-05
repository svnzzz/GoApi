package initializers

import (
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
)

var Container *azcosmos.ContainerClient

func CreateConnection() error {
	endpoint, key, db, cont := LoadEnvVariables()
	credentials, err := azcosmos.NewKeyCredential(key)
	if err != nil {
		fmt.Printf("Missing credentials %v", err)
	}

	clientOptions := azcosmos.ClientOptions{
		EnableContentResponseOnWrite: true,
	}

	client, err := azcosmos.NewClientWithKey(endpoint, credentials, &clientOptions)
	if err != nil {
		fmt.Printf("Client error: %v", err)
	}

	database, err := client.NewDatabase(db)
	if err != nil {
		fmt.Printf("Database error: %v", err)
	}

	container, err := database.NewContainer(cont)
	if err != nil {
		fmt.Printf("Container error: %v", err)
	}

	Container = container
	return nil
}
