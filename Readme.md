# Streamblast API Client

Usage:

```
package main

import (
	"github.com/iamsalnikov/streamblast-api"
	"fmt"
	"os"
	"flag"
)

func main() {
	var dreamsContentID int64
	var episodeID string

	flag.Int64Var(&dreamsContentID, "dcid", 0, "Dreams content ID")
	flag.StringVar(&episodeID, "episode", "", "Episode ID")
	flag.Parse()

	client := streamblast_api.Client{
		BaseURI: "http://streamblast.cc",
		DreamsContentID: dreamsContentID,
	}

	links, err := client.GetLinks(episodeID)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for quantity, link := range links {
		fmt.Fprintf(os.Stdout, "%s : %s\n", quantity, link)
	}
}
```