package secrets

import (
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"context"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
	"log"
)

var client *secretmanager.Client = nil
var ctx = context.Background()

// GCP project in which to store secrets in Secret Manager.
var projectID = "splatstats-312616"

func initializeSecrets() {
	// Create the client.
	var err error
	client, err = secretmanager.NewClient(ctx)
	if err != nil {
		log.Fatalf("failed to setup client: %v", err)
	}
}

func GetSecret(secretName string) string {
	if client == nil {
		initializeSecrets()
	}
	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: "projects/" + projectID + "/secrets/" + secretName + "/versions/latest",
	}
	resp, err := client.AccessSecretVersion(ctx, req)
	if err != nil {
		panic(err)
	}
	return resp.Payload.String()
}
