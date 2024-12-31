package main

import (
        "os"
        "context"
        "log"
        "strings"

        "github.com/aws/aws-sdk-go-v2/config"
        opensearch "github.com/opensearch-project/opensearch-go/v2"
        opensearchapi "github.com/opensearch-project/opensearch-go/v2/opensearchapi"
        requestsigner "github.com/opensearch-project/opensearch-go/v2/signer/awsv2"
)

func main() {
        ctx := context.Background()

        var endpoint = os.Getenv("AWS_OPENSEARCH_ENDPOINT")


        awsCfg, err := config.LoadDefaultConfig(ctx,
                config.WithRegion(os.Getenv("AWS_REGION")),
        )
        if err != nil {
                log.Fatalf("failed to load aws configuraiton: %v", err) // Do not log.fatal in a production ready app.
        }

        // Create an AWS request Signer and load AWS configuration using default config folder or env vars.
        signer, err := requestsigner.NewSignerWithService(awsCfg, "aoss") // Use "aoss" for Amazon OpenSearch Serverless
        if err != nil {
                log.Fatalf("failed to create signer: %v", err)
        }

        // Create an opensearch client and use the request-signer.
        client, err := opensearch.NewClient(opensearch.Config{
                Addresses: []string{endpoint},
                Signer:    signer,
        })
        if err != nil {
                log.Fatalf("failed to create new opensearch client: %v", err)
        }

        indexName := "go-test-index-cloudops-ygor"

        // Define index mapping.
        mapping := strings.NewReader(`{
         "settings": {
           "index": {
                "number_of_shards": 4
                }
              }
         }`)

        // Create an index with non-default settings.
        createIndex := opensearchapi.IndicesCreateRequest{
                Index: indexName,
                Body:  mapping,
        }

        createIndexResponse, err := createIndex.Do(ctx, client)
        if err != nil {
                log.Fatalf("failed to create index: %v", err)
        }
        log.Printf("created index: %#v", createIndexResponse)

        // Delete previously created index.
        deleteIndex := opensearchapi.IndicesDeleteRequest{
                Index: []string{indexName},
        }

        deleteIndexResponse, err := deleteIndex.Do(ctx, client)
        if err != nil {
                log.Fatalf("failed to delete index: %v", err)
        }
        log.Printf("deleted index: %#v", deleteIndexResponse)
}
