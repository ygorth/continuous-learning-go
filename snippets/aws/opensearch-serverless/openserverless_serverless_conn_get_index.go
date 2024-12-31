package main

import (
        "os"
        "context"
        "log"

        "github.com/aws/aws-sdk-go-v2/config"
        opensearch "github.com/opensearch-project/opensearch-go/v2"
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

        res, err := client.Indices.Get([]string{"movies"})
        if err != nil {
           log.Printf("error occurred: [%s]", err.Error())
        }
        log.Printf("response: [%+v]", res)
}
