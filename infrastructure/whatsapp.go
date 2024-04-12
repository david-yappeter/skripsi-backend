package infrastructure

import (
	"context"
	"fmt"
	"myapp/global"

	_ "github.com/mattn/go-sqlite3"
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
)

type WhatsappManager interface {
	IsLoggedIn(ctx context.Context) (isLoggedIn bool)
	LoginQr(ctx context.Context) (qrLogin chan (string), err error)
	SendMessage(ctx context.Context, to types.JID, message *waProto.Message) (err error)
	Disconnect()
}

type whatsappManager struct {
	sqlstoreContainer *sqlstore.Container
	client            *whatsmeow.Client
}

func eventHandler(evt interface{}) {
	switch v := evt.(type) {
	case *events.Message:
		_ = v
		// fmt.Println("Received a message!", v.Message.GetConversation())
	}
}

func NewWhatsappManager(whatappConfig global.WhatsappConfig) WhatsappManager {
	// import sqlite3
	container, err := sqlstore.New("sqlite3", fmt.Sprintf("file:%s?_foreign_keys=on", whatappConfig.SqlStoreFilePath), nil)
	if err != nil {
		panic(err)
	}
	// If you want multiple sessions, remember their JIDs and use .GetDevice(jid) or .GetAllDevices() instead.
	deviceStore, err := container.GetFirstDevice()
	if err != nil {
		panic(err)
	}
	// clientLog := waLog.Stdout("Client", "DEBUG", true)
	client := whatsmeow.NewClient(deviceStore, nil)
	client.AddEventHandler(eventHandler)

	return &whatsappManager{
		sqlstoreContainer: container,
		client:            client,
	}
}

func (i *whatsappManager) IsLoggedIn(ctx context.Context) (isLoggedIn bool) {
	return i.client.IsLoggedIn()
}

func (i *whatsappManager) LoginQr(ctx context.Context) (chan (string), error) {
	// Create a channel to receive the QR code string
	qrStringChan := make(chan string)

	// Run the QR code retrieval process in a goroutine
	go func() {
		defer func() {
			close(qrStringChan)
		}()

		// Assuming i.client is your WhatsMeow client
		if i.client.Store.ID == nil {
			// ctxWithCancel, cancel := context.WithTimeout(ctx, time.Minute)
			// defer cancel()

			qrChan, _ := i.client.GetQRChannel(ctx)

			i.client.Connect()

			for evt := range qrChan {
				if evt.Event == "code" {
					// Assuming evt.Code contains the QR code string
					qrStringChan <- evt.Code
				} else {
					fmt.Println("Login event:", evt.Event)
				}
			}
		}

		qrStringChan <- "" // Return an empty string if QR code retrieval fails
	}()

	return qrStringChan, nil
	// Wait for the QR code string or timeout
	// select {
	// case qrLogin := <-qrStringChan:
	// 	return qrLogin, nil
	// 	// case <-ctxWithCancel.Done():
	// 	// 	return "", ctxWithCancel.Err() // Return error if context timeout occurs
	// }

	// var qrLogin = ""
	// if i.client.Store.ID == nil {
	// 	qrChan, err := i.client.GetQRChannel(ctx)
	// 	if err != nil {
	// 		return qrLogin, err
	// 	}

	// 	err = i.client.Connect()
	// 	if err != nil {
	// 		return qrLogin, err
	// 	}

	// 	for evt := range qrChan {
	// 		if evt.Event == "code" {
	// 			// Assuming evt.Code contains the QR code string
	// 			qrLogin = evt.Code
	// 			break
	// 		}
	// 	}

	// 	return qrLogin, nil
	// }

	// return qrLogin, nil
}

func (i *whatsappManager) SendMessage(ctx context.Context, to types.JID, message *waProto.Message) error {
	_, err := i.client.SendMessage(ctx, to, message)
	if err != nil {
		return err
	}

	return nil
}

func (i *whatsappManager) Disconnect() {
	i.client.Disconnect()
}
