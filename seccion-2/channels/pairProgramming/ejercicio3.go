package main
import (
    "fmt"
    "time"
)
type Message struct {
    typ     string
    content string
}
func main() {
    chanel := make(chan Message)
    for {
        var mensaje Message
        fmt.Println("(Para salir escriba salir)")
        fmt.Println("Ingrese el tipo de mensaje: ")
        fmt.Scan(&mensaje.typ)
        if mensaje.typ == "salir" {
            break
        }
        fmt.Println("Ingrese el contenido del mensaje: ")
        fmt.Scan(&mensaje.content)
        if mensaje.typ == "email" {
            go setEmail(mensaje, chanel)
            time.Sleep(time.Second)
            confirmMessage(chanel)
        }
        if mensaje.typ == "SMS" {
            go setSms(mensaje, chanel)
            time.Sleep(time.Second)
            confirmMessage(chanel)
        }
        if mensaje.typ == "push" {
            go setPush(mensaje, chanel)
            time.Sleep(time.Second)
            confirmMessage(chanel)
        }
    }
    close(chanel)
}
func setEmail(message Message, chanel chan Message) {
    chanel <- message
}
func confirmMessage(c chan Message) {
    mensaje := <-c
    fmt.Println("Mensaje recibido", mensaje)
}
func setSms(message Message, chanel chan Message) {
    chanel <- message
}
func setPush(message Message, chanel chan Message) {
    chanel <- message
}