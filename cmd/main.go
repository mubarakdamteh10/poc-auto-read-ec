package main

import (
	"fmt"
	"log"
	"os"
	"poc-auto-read-ec/pkg/process"
	"poc-auto-read-ec/pkg/sftp"
	"sync"
	"time"
)

func init() {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatalf("error getting hostname: %v", err)
	}
	fmt.Printf("logger initialized on host: %s\n", hostname)
}

func main() {
	fmt.Printf("Logger: Start Auto Read EDICustomerRequest Service at: %v\n", time.Now())

	sFTPService := sftp.NewSFTPService()
	defer sFTPService.CloseClient()

	if _, err := sFTPService.ConnectClient(); err != nil {
		log.Fatalf("cannot connect to SFTP server: %v", err)
	}

	autoReadECProcessService := process.NewProcessService()

	var wg sync.WaitGroup
	var mu sync.Mutex

	concurrencyLimit := 4
	semaphore := make(chan struct{}, concurrencyLimit)

	semaphore <- struct{}{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		defer func() { <-semaphore }()

		if err := autoReadECProcessService.ProcessAutoReadEC(); err != nil {
			mu.Lock()
			fmt.Printf("Error processing: %v\n", err)
			mu.Unlock()
			return
		}

		mu.Lock()
		fmt.Println("Successfully processed")
		mu.Unlock()
	}()

	wg.Wait()
	fmt.Printf("Auto Read EDICustomerRequest Service completed at: %v\n", time.Now())
}
