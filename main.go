package main

import (
	"log"

	convoy "github.com/frain-dev/convoy-go"
)

const (
	someValidWebhookUrl = "https://eotozt3z49nhv9t.m.pipedream.net"
	aValidConvoyApiKey  = "---------------INSERT YOUR API KEY HERE----------------"
)

func main() {
	// Initalise client
	convoyClient := convoy.New(convoy.Options{
		APIKey: aValidConvoyApiKey,
	})

	// Create application
	app, err := convoyClient.Applications.Create(&convoy.CreateApplicationRequest{
		Name:         "Zflash",
		SupportEmail: "support@fluffycloud.comz",
	}, nil)

	if err != nil {
		log.Fatal("failed to create app  \n", err)
	}

	// Create endpoint
	endpoint, err := convoyClient.Endpoints.Create(app.UID, &convoy.CreateEndpointRequest{
		URL:         someValidWebhookUrl,
		Secret:      "this is a secret i want",
		Description: "Flashbot reconciliation endpoint",
	}, nil)

	if err != nil {
		log.Fatal("failed to create app endpoint \n", err)
	}

	//updateEndpoint 🚨 the issue
	updatedEndpoint, err := convoyClient.Endpoints.Update(app.UID, endpoint.UID, &convoy.CreateEndpointRequest{
		URL:         someValidWebhookUrl,
		Secret:      "i'm a changed secret", // 👀 I've updated my secret here
		Description: "Flashbot Reconfigaration endpoint",
	}, nil)

	if err != nil {
		log.Fatal("failed to update app endpoint \n", err)
	}

	foundEndpoint, err := convoyClient.Endpoints.Find(app.UID, endpoint.UID, nil)
	if err != nil {
		log.Fatal("failed to find app endpoint \n", err)
	}

	log.Printf("Inital  Secret for endpoint 🪪 (%s) is 🔐(%s) ", endpoint.UID, endpoint.Secret)
	log.Printf("Updated Secret for endpoint 🪪 (%s) is 🔐(%s) ", updatedEndpoint.UID, updatedEndpoint.Secret)
	log.Printf("Endpoint from foundEndpoint 🪪 (%s) is 🔐(%s) ", foundEndpoint.UID, foundEndpoint.Secret)
}
