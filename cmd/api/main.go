package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/denisenkom/go-mssqldb"

	// See https://github.com/peopledatalabs/peopledatalabs-go
	pdl "github.com/peopledatalabs/peopledatalabs-go"
	pdlmodel "github.com/peopledatalabs/peopledatalabs-go/model"
)

func main() {
	// Set your API key
	apiKey := "YOUR API KEY"
	// Set API key as environmental variable
	// apiKey := os.Getenv("API_KEY")

	// Create a client, specifying your API key
	client := pdl.New(apiKey)

	// Create an array of parameters JSON objects
	params := pdlmodel.BulkEnrichPersonParams{
		Requests: []pdlmodel.BulkEnrichSinglePersonParams{
			{
				Params: pdlmodel.PersonParams{
					Profile: []string{"linkedin.com/in/seanthorne"},
				},
			},
			{
				Params: pdlmodel.PersonParams{
					Profile: []string{"linkedin.com/in/randrewn"},
				},
			},
		},
	}

	// Pass the parameters object to the Bulk Person Enrichment API
	responses, err := client.Person.BulkEnrich(context.Background(), params)

	// Iterate through the array of API responses
	for _, response := range responses {
		// Check for successful response
		if err == nil {
			// Convert the API response to JSON
			jsonResponse, jsonErr := json.Marshal(response.Data)
			if jsonErr == nil {
				var record map[string]interface{}
				json.Unmarshal(jsonResponse, &record)
				// Print selected fields
				fmt.Println(
					record["work_email"],
					record["full_name"],
					record["job_title"],
					record["job_company_name"])

				fmt.Println("Successfully enriched profile with PDL data.")

				// Save enrichment data to JSON file
				out, outErr := os.Create(fmt.Sprintf("my_pdl_enrichment.%s.jsonl", record["linkedin_username"]))
				defer out.Close()
				if outErr == nil {
					out.WriteString(string(jsonResponse) + "\n")
				}
				out.Sync()
			}
		} else {
			fmt.Println("Enrichment unsuccessful. See error and try again.")
			fmt.Println("error:", err)
		}
	}

}
