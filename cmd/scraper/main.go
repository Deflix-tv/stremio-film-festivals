package main

import (
	"flag"
	"strings"
)

var (
	dataDir = flag.String("dataDir", ".", "Location of the data directory. This is where the CSV files will be written.")
)

func main() {
	flag.Parse()

	// Clean input
	if strings.HasSuffix(*dataDir, "/") {
		*dataDir = strings.TrimRight(*dataDir, "/")
	}

	imdbClient := newIMDbClient()
	wikiClient := newWikiClient()
	wikiClient.scrapeAcademyAwardWinners(*dataDir + "/academy-awards-winners.csv")
	imdbClient.scrapePalmeDorWinners(*dataDir + "/palme-dor-winners.csv")
	imdbClient.scrapeGoldenLionWinners(*dataDir + "/golden-lion-winners.csv")
	imdbClient.scrapeGoldenBearWinners(*dataDir + "/golden-bear-winners.csv")
}
