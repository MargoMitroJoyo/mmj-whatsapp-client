package whatsapp

import (
	"context"
	"fmt"
	"os"

	"github.com/mdp/qrterminal/v3"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types"
	waLog "go.mau.fi/whatsmeow/util/log"
	"google.golang.org/protobuf/proto"

	_ "modernc.org/sqlite"
)

// WhatsAppClient is a wrapper around the whatsmeow.Client
type WhatsAppClient struct {
	Client *whatsmeow.Client
}

// Initializes a new WhatsApp client with SQLite store.
func NewClient() (*WhatsAppClient, error) {
	dbLog := waLog.Stdout("Database", "INFO", true)
	container, err := sqlstore.New(context.Background(), "sqlite", "file:.store/session.db?_pragma=journal_mode(WAL)&_pragma=busy_timeout(5000)&_pragma=foreign_keys(ON)", dbLog)
	if err != nil {
		return nil, err
	}

	deviceStore, err := container.GetFirstDevice(context.Background())
	if err != nil {
		return nil, err
	}

	clientLog := waLog.Stdout("Client", "INFO", true)
	client := whatsmeow.NewClient(deviceStore, clientLog)

	return &WhatsAppClient{Client: client}, nil
}

// Sends a message to a specified WhatsApp number (should be in E.164 format).
func (wac *WhatsAppClient) SendMessage(to string, message string) error {
	targetJID := types.NewJID(to, "s.whatsapp.net")

	msg := &waE2E.Message{
		Conversation: proto.String(message),
	}

	_, err := wac.Client.SendMessage(context.Background(), targetJID, msg)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}
	return nil
}

// Start initializes the WhatsApp client.
func (wac *WhatsAppClient) Start() error {
	if wac.Client.Store.ID == nil {
		qrChan, _ := wac.Client.GetQRChannel(context.Background())
		err := wac.Client.Connect()
		if err != nil {
			return err
		}

		for evt := range qrChan {
			if evt.Event == "code" {
				fmt.Println("Scan the QR code to log in:")
				qrterminal.GenerateWithConfig(evt.Code, qrterminal.Config{
					Level:      qrterminal.L,
					Writer:     os.Stdout,
					HalfBlocks: true,
				})
			} else {
				fmt.Println("Login event:", evt.Event)
			}
		}
	} else {
		err := wac.Client.Connect()
		if err != nil {
			return err
		}
	}
	return nil
}

// Stop disconnects the WhatsApp client.
func (wac *WhatsAppClient) Stop() {
	wac.Client.Disconnect()
}
