package main

import (
	"flag"
	"log"
	"strconv"

	"github.com/abhijeetmohanan/bridgeftp/utils"
)

func main() {
	log.Println("Starting Connection")
	// Define flags
	source_ftp := flag.String("src", "", "Source ftp endpoint Example: ftp://user:pass@host/filepath")
	destination_ftp := flag.String("dest", "", "Destination ftp endpoint Example: ftp://user:pass@host/filepath")
	bytesize := flag.String("bs", "", "Defines the chuck byte size")

	flag.Parse()

	utils.NullChecker("source", *source_ftp)
	utils.NullChecker("destination", *destination_ftp)

	bsvalue := 64
	var err error

	// validate byte size
	if *bytesize == "" {
		bsvalue, err = strconv.Atoi("64")
		if err != nil {
			log.Panicln("Invalid byte size ", *bytesize)
		}
	} else {
		bsvalue, err = strconv.Atoi(*bytesize)
		if err != nil {
			log.Panicln("Failed :: Invalid default byte size", *bytesize)
		}
	}

	log.Printf("Streaming Byte Size is %d Kb", bsvalue)

	source_map_kv := utils.ParseInput("source", *source_ftp)
	destination_map_kv := utils.ParseInput("destination", *destination_ftp)

	if source_map_kv["scheme"] == "ftp" && destination_map_kv["scheme"] == "ftp" {
		log.Println("Source and Destination are ftp endpoints")

		// validate ftp parameters : panic on failure
		utils.FtpParamsValidator(source_map_kv)
		utils.FtpParamsValidator(destination_map_kv)

		// Create Coonection and start streaming

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
