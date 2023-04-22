package main

import (
	"fmt"
	"time"
)

type Message struct {
	Type    string
	Content string
}

func main() {
	channelEmail := make(chan Message)
	channelPush := make(chan Message)
	channelSms := make(chan Message)
	stopEmail := false
	stopPush := false
	stopSms := false

	for {
		var message Message
		var typeMessage string
		time.Sleep(time.Second)

		fmt.Println("Ingrese el tipo de mensaje (email, sms, push) o alguna de estas opciones (salir, pausar, reanudar) ")
		fmt.Scanln(&typeMessage)

		if typeMessage == "salir" {
			break
		}

		if typeMessage == "pausar" || typeMessage == "reanudar" {
			optionsSecondary(&stopEmail, &stopPush, &stopSms, typeMessage)
			//fmt.Println(stopEmail)
		} else {
			message.Type = typeMessage
			fmt.Println("Ingrese el contenido del mensaje: ")
			fmt.Scanln(&message.Content)
			sendMessage(message, channelEmail, channelPush, channelSms, &stopEmail, &stopSms, &stopPush)
			//time.Sleep(time.Second)
		}
	}
	close(channelEmail)
	close(channelPush)
	close(channelSms)
}
func menuBase() {
	fmt.Println("1. Email")
	fmt.Println("2. SMS")
	fmt.Println("3. Push")
	fmt.Print("Elija una opción: ")
}

func sendMessage(message Message, channelEmail, channelPush, channelSms chan Message, stopEmail, stopSms, stopPush *bool) {
	if message.Type == "email" {
		go setMessageToChannel(message, channelEmail)
		time.Sleep(time.Second)
		go getMessageFromChannel(channelEmail, stopEmail)

	}
	if message.Type == "sms" {
		go setMessageToChannel(message, channelSms)
		time.Sleep(time.Second)
		go getMessageFromChannel(channelSms, stopSms)
	}
	if message.Type == "push" {
		go setMessageToChannel(message, channelPush)
		time.Sleep(time.Second)
		go getMessageFromChannel(channelPush, stopPush)
	}
}

func optionsSecondary(stopEmail, stopPush, stopSms *bool, option string) {
	menuBase()
	var typeMessage string
	fmt.Scanln(&typeMessage)
	if option == "reanudar" {
		if typeMessage == "1" {
			*stopEmail = false
		}
		if typeMessage == "2" {
			*stopSms = false
		}
		if typeMessage == "3" {
			*stopPush = false
		}
	} else if option == "pausar" {
		if typeMessage == "1" {
			*stopEmail = true
		}
		if typeMessage == "2" {
			*stopSms = true
		}
		if typeMessage == "3" {
			*stopPush = true
		}
	}

}

func setMessageToChannel(message Message, chanel chan Message) {
	fmt.Println(message.Type, " recibido")
	chanel <- message
}
func getMessageFromChannel(c chan Message, pausa *bool) {
	for !*pausa {
		select {
		case mensaje := <-c:
			fmt.Println("Mensaje procesado", mensaje, *pausa)
		}
	}
}
