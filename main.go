package main

import (
	"flag"
	"log"

	"github.com/abhijeetmohanan/bridgeftp/utils"
)

var version string = "testing"

func main() {
	log.Printf("BridgeFTp Version = %s", version)
	// Define flags
	source_ftp := flag.String("src", "", "Source ftp endpoint Example: ftp://user:pass@host/filepath")
	destination_ftp := flag.String("dest", "", "Destination ftp endpoint Example: ftp://user:pass@host/filepath")
	bytesize := flag.Int("bs", 64, "Defines the chuck byte size")

	flag.Parse()

	utils.NullChecker("source", *source_ftp)
	utils.NullChecker("destination", *destination_ftp)

	bsvalue := *bytesize

	log.Printf("Streaming Byte Size is %d Kb", bsvalue)

	source_map_kv, err := utils.ParseInput("source", *source_ftp)
	if err != nil {
		log.Fatalf("Failed to parse input %v", err)
	}
	destination_map_kv, err := utils.ParseInput("destination", *destination_ftp)
	if err != nil {
		log.Fatalf("Failed to parse input %v", err)
	}

	if source_map_kv["scheme"] == "ftp" && destination_map_kv["scheme"] == "ftp" {
		log.Println("Source and Destination are ftp endpoints")

		// validate ftp parameters : panic on failure
		utils.FtpParamsValidator(source_map_kv)
		utils.FtpParamsValidator(destination_map_kv)

		// Create Connection and start streaming

		utils.FtpClientHandler(source_map_kv, destination_map_kv, bsvalue)
	}
	if source_map_kv["scheme"] == "sftp" && destination_map_kv["scheme"] == "sftp" {
		log.Println("Source and Destination are ftp endpoints")

		// validate ftp parameters : panic on failure
		utils.FtpParamsValidator(source_map_kv)
		utils.FtpParamsValidator(destination_map_kv)

		// Create Coonection and start streaming

		utils.SftpClientHandler(source_map_kv, destination_map_kv, bsvalue)
	}
}
