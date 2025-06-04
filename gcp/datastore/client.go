package datastore

import (
	"context"
	"os"

	"cloud.google.com/go/compute/metadata"
	"cloud.google.com/go/datastore"
	"google.golang.org/api/option"
)

func Must(ctx context.Context, databaseID string) *datastore.Client {
	projectID, err := metadata.ProjectIDWithContext(ctx)
	if err != nil {
		panic(err)
	}

	var opts []option.ClientOption

	if os.Getenv("DATASTORE_EMULATOR_HOST") == "" {
		if credentialsFile := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"); credentialsFile != "" {
			opts = append(opts, option.WithCredentialsFile(credentialsFile))
		}
	}

	var client *datastore.Client

	if databaseID == "" {
		client, err = datastore.NewClient(ctx, projectID, opts...)
	} else {
		client, err = datastore.NewClientWithDatabase(ctx, projectID, databaseID, opts...)
	}

	if err != nil {
		panic(err)
	}

	return client
}
