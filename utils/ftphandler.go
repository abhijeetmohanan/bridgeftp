package utils

import (
	"bytes"
	"io"
	"log"
	"time"

	"github.com/jlaffaye/ftp"
)

func FtpClientHandler(source map[string]string, destination map[string]string) {
	// Make connections
	source_ftp, err := ftp.Dial(source["host"], ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Source Connected Successfully")
	}
	destination_ftp, err := ftp.Dial(destination["host"], ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Destination Connected Successfully")
	}

	// Login in using creds
	err = source_ftp.Login(source["username"], source["password"])
	if err != nil {
		log.Fatal(err)
	}
	err = destination_ftp.Login(destination["username"], destination["password"])
	if err != nil {
		log.Fatal(err)
	}

	// start Reading from Source
	reader_ftp, err := source_ftp.Retr(source["path"])
	if err != nil {
		log.Fatal(err)
	}

	// initalize a buffer of size 64Kb
	p := make([]byte, 64*1024)

	for {
		n, err := reader_ftp.Read(p)

		if err == io.EOF {
			break
		}

		wdata := bytes.NewBufferString(string(p[:n]))
		err = destination_ftp.Append(destination["path"], wdata)
		if err != nil {
			log.Fatal(err)
		}

	}
	// Close Source FTP
	defer reader_ftp.Close()
}
