// Package whatsapp provides a simple abstraction over the WhatsMeow library
// to interact with WhatsApp Web. It supports features like persistent session
// handling, QR-based login, sending messages, and managing connection states.
//
// This package is particularly useful for building bots or automating tasks
// over WhatsApp using Go.
//
// Example usage:
//
//	client, err := whatsapp.NewClient()
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	err = client.Start()
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	err = client.SendMessage("123456789@s.whatsapp.net", "Hello from Go!")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	client.Stop()
package whatsapp
