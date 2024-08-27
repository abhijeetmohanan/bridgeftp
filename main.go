package main

import (
	"flag"
	"log"

	"github.com/abhijeetmohanan/ftpcli/utils"
)

func main() {
	log.Println("Starting Connection")
	// Define flags
	source_ftp := flag.String("src", "", "Source ftp endpoint Example: ftp://user:pass@host/filepath")
	destination_ftp := flag.String("dest", "", "Destination ftp endpoint Example: ftp://user:pass@host/filepath")

	flag.Parse()

	utils.NullChecker("source", *source_ftp)
	utils.NullChecker("destination", *destination_ftp)

	source_map_kv := utils.ParseInput("source", *source_ftp)
	destination_map_kv := utils.ParseInput("destination", *destination_ftp)

	if utils.SchemeValidator(source_map_kv["scheme"]) && utils.SchemeValidator(destination_map_kv["scheme"]) {
		log.Println("Scheme Validated Both source and Destination are ftp endpoints")

		// validate ftp parameters : panic on failure
		utils.FtpParamsValidator(source_map_kv)
		utils.FtpParamsValidator(destination_map_kv)

		// Create Coonection and start streaming

		utils.FtpClientHandler(source_map_kv, destination_map_kv)

	} else {
		log.Panicln("Only FTP endooints are supported as of now")
	}
}
