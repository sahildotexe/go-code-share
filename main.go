package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		colorCode := "\033[33m"
		resetCode := "\033[0m"
		statement := "Usage: goshare <file_path>"
		fmt.Println("You must specify a file path.")
		fmt.Printf("%s%s%s\n", colorCode, statement, resetCode)
		os.Exit(1)
	}

	filePath := os.Args[1]

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Printf("File not found: %s\n", filePath)
		os.Exit(1)
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	tempDir, err := os.MkdirTemp("", "file-sharing-tool")
	if err != nil {
		log.Fatal(err)
	}

	tempFilePath := filepath.Join(tempDir, "shared_file.txt")
	err = os.WriteFile(tempFilePath, content, 0644)
	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command("ngrok", "http", "8080")
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()

	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
		}
	}()

	go func() {
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
		}
	}()
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Second * 2)

	ngrokURL, err := getNgrokURL()
	if err != nil {
		log.Fatal(err)
	}
	colorCode := "\033[32m"
	resetCode := "\033[0m"

	fmt.Printf("Share the following link:\n%s%s%s\n", colorCode, ngrokURL, resetCode)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, tempFilePath)
	})

	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatal(err)
		}
	}()

	fmt.Println("Press Ctrl+C to stop the server.")
	select {}
}

func getNgrokURL() (string, error) {
	resp, err := http.Get("http://127.0.0.1:4040/api/tunnels")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var tunnels struct {
		Tunnels []struct {
			PublicURL string `json:"public_url"`
		} `json:"tunnels"`
	}

	if err := json.Unmarshal(body, &tunnels); err != nil {
		return "", err
	}

	if len(tunnels.Tunnels) == 0 {
		return "", errors.New("ngrok URL not found")
	}

	return tunnels.Tunnels[0].PublicURL, nil
}
