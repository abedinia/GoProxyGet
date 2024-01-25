package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"golang.org/x/net/proxy"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Usage: [SOCKS5 proxy address] [URL to download]")
		os.Exit(1)
	}

	proxyAddr := os.Args[1]
	targetURL := os.Args[2]

	dialer, err := proxy.SOCKS5("tcp", proxyAddr, nil, proxy.Direct)
	if err != nil {
		panic(err)
	}

	transport := &http.Transport{
		DialContext: func(ctx context.Context, network, addr string) (conn net.Conn, err error) {
			return dialer.Dial(network, addr)
		},
	}

	client := &http.Client{Transport: transport}

	request, err := http.NewRequest("GET", targetURL, nil)
	if err != nil {
		panic(err)
	}

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	contentLength := response.ContentLength
	fmt.Printf("File size: %.2f MB\n", float64(contentLength)/1024.0/1024.0)

	parsedURL, err := url.Parse(targetURL)
	if err != nil {
		panic(err)
	}
	filename := filepath.Base(parsedURL.Path)

	outFile, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	drawProgressBar := func(downloaded, total int64) {
		const barLength = 40
		percent := float64(downloaded) / float64(total) * 100
		filled := int(float64(barLength) * percent / 100)
		bar := strings.Repeat("█", filled) + strings.Repeat("-", barLength-filled)
		fmt.Printf("\r\033[32m[%s] %.2f%%\033[0m", bar, percent)
	}

	buffer := make([]byte, 32*1024)
	var downloaded int64
	startTime := time.Now()
	for {
		read, err := response.Body.Read(buffer)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if read == 0 {
			break
		}
		_, writeErr := outFile.Write(buffer[:read])
		if writeErr != nil {
			panic(writeErr)
		}
		downloaded += int64(read)
		elapsed := time.Since(startTime)
		speed := float64(downloaded) / elapsed.Seconds() / 1024 / 1024
		fmt.Printf(" - %.2f MB/s", speed)
		drawProgressBar(downloaded, contentLength)
	}
	fmt.Println("\n\033[36mFile downloaded successfully\033[0m")
}
