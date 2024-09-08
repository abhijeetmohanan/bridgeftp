package utils

import (
	"bytes"
	"io"
	"log"
	"time"

	"github.com/jlaffaye/ftp"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func FtpClientHandler(source map[string]string, destination map[string]string, bytesize int) {
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
	chuckbuffer := make([]byte, bytesize*1024)

	for {
		n, err := reader_ftp.Read(chuckbuffer)

		if err == io.EOF {
			break
		}

		wdata := bytes.NewBufferString(string(chuckbuffer[:n]))
		err = destination_ftp.Append(destination["path"], wdata)
		if err != nil {
			log.Fatal(err)
		}

	}
	// Close Source FTP
	defer reader_ftp.Close()
}

func connectToSFTP(user, password, host string) (*sftp.Client, error) {
	log.Printf("starting clinet connection for  %s", host)
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn, err := ssh.Dial("tcp", host, config)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
		return nil, err
	}

	client, err := sftp.NewClient(conn)
	if err != nil {
		log.Fatalf("failed to create SFTP client: %v", err)
		return nil, err
	}
	log.Printf("SFTP Client initalized %s", host)

	return client, nil
}

func SftpClientHandler(source map[string]string, destination map[string]string, bytesize int) {
	// Establish connections to both SFTP servers

	srcClient, err := connectToSFTP(source["username"], source["password"], source["host"])
	if err != nil {
		log.Fatalf("Failed to connect to source SFTP: %v", err)
	}
	defer srcClient.Close()

	dstClient, err := connectToSFTP(destination["username"], destination["password"], destination["host"])
	if err != nil {
		log.Fatalf("Failed to connect to destination SFTP: %v", err)
	}
	defer dstClient.Close()

	// Open the source file
	srcFile, err := srcClient.Open(source["path"])
	if err != nil {
		log.Fatalf("Failed to open source file: %v", err)
	}
	defer srcFile.Close()

	// Create the destination file
	dstFile, err := dstClient.Create(destination["path"])
	if err != nil {
		log.Fatalf("Failed to create destination file: %v", err)
	}
	defer dstFile.Close()

	log.Println("starting stream")

	buffer := make([]byte, bytesize*1024)

	// Stream the file from source to destination
	bytesCopied, err := io.CopyBuffer(dstFile, srcFile, buffer)
	if err != nil {
		log.Fatalf("Failed to copy file: %v", err)
	}

	log.Printf("Successfully copied %d bytes ", bytesCopied)
}
