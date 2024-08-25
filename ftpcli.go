package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/jlaffaye/ftp"
)

func main() {
	c1, err := ftp.Dial("ftp_server_1:21", ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	c2, err := ftp.Dial("ftp_server_2:21", ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection Made")

	err = c1.Login("user1", "pass1")
	if err != nil {
		log.Fatal(err)
	}
	err = c2.Login("user2", "pass2")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("logged in")

	reader_ftp, err := c1.Retr("dumpdata.none")
	if err != nil {
		log.Fatal(err)
	}
	defer reader_ftp.Close()

	p := make([]byte, 4)

	for {
		n, err := reader_ftp.Read(p)

		if err == io.EOF {
			break
		}
		wdata := bytes.NewBufferString(string(p[:n]))
		err = c2.Append("dumpdata.none", wdata)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(p[:n]))
	}
}
