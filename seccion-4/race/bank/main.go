package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Estructura BankAccount que guarda el saldo de la cuenta
type BankAccount struct {
	Balance int
	mutex   sync.Mutex
}

// Método que incrementa el saldo de la cuenta con la cantidad dada
func (account *BankAccount) Deposit(amount int) {
	account.mutex.Lock()
	account.Balance += amount
	fmt.Printf("Depósito: %d. Balance actual: %d\n", amount, account.Balance) // imprime la cantidad depositada y el saldo actual
	account.mutex.Unlock()
}

// Método que decrementa el saldo de la cuenta con la cantidad dada
func (account *BankAccount) Withdraw(amount int) {
	account.mutex.Lock()
	if account.Balance >= amount {
		account.Balance -= amount
		fmt.Printf("Retiro: %d. Balance actual: %d\n", amount, account.Balance) // imprime la cantidad retirada y el saldo actua
	} else {
		fmt.Printf("Retiro: %d. El retiro ha sido DENEGADO\n", amount)

	}
	account.mutex.Unlock()
}

// Función que realiza una transacción (depósito o retiro) en la cuenta
func processTransaction(account *BankAccount, transactionType string, amount int, w *sync.WaitGroup) {
	switch transactionType {
	case "deposit":
		account.Deposit(amount)
	case "withdraw":
		account.Withdraw(amount)
	}
	defer w.Done()
}

func main() {
	var w sync.WaitGroup
	rand.Seed(time.Now().UnixNano())                     // semilla para generar números aleatorios
	account := &BankAccount{Balance: 1000}               // inicializa una cuenta con un saldo de 1000
	fmt.Printf("Balance inicial: %d\n", account.Balance) // imprime el saldo inicial

	for i := 0; i < 5; i++ { // loop para generar 5 transacciones
		w.Add(1)
		transactionType := "" // inicializa el tipo de transacción
		if i%2 == 0 {         // si el índice es par, la transacción es un depósito
			transactionType = "deposit"
		} else { // si el índice es impar, la transacción es un retiro
			transactionType = "withdraw"
		}
		amount := rand.Intn(500) + 1                                // genera una cantidad aleatoria entre 1 y 500 para la transacción
		go processTransaction(account, transactionType, amount, &w) // inicia una goroutine para procesar la transacción
	}

	w.Wait()
	fmt.Printf("Balance final: %d\n", account.Balance) // imprime el saldo final
}
