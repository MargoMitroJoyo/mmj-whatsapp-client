# WhatsApp Client API

[![Go Reference](https://pkg.go.dev/badge/go.mau.fi/whatsmeow.svg)](https://pkg.go.dev/go.mau.fi/whatsmeow)

This is a go package that is used as a wrapper over whatsmeow. This package aims to simplify sending WhatsApp messages programmatically using go.

## Available Features

1. Login using QRCode printed on terminal ✅
2. Save login session using SQLite ✅
3. Send Messages ✅

## Usage

This is the minimum example usage using whatsapp client API.

```go
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/tegaraditya/mmj-whatsapp-client/pkg/whatsapp"
)

func main() {
	client, err := whatsapp.NewClient()
	if err != nil {
		panic(fmt.Sprintf("Failed to create WhatsApp client: %v", err))
	}

	err = client.Start()
	if err != nil {
		panic(fmt.Sprintf("Failed to start WhatsApp client: %v", err))
	}

	err = client.SendMessage("6285212337564", "Hello from Go!")
	if err != nil {
		panic(fmt.Sprintf("Failed to send message: %v", err))
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	client.Stop()
}

```
