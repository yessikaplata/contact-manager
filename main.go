package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Contact struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func saveContacts(contacts []Contact) error {
	file, err := os.Create("contacts.json")
	if err != nil {
		return err
	}

	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(contacts)
	if err != nil {
		return err
	}

	return nil
}

func loadContacts(contacts *[]Contact) error {
	file, err := os.Open("contacts.json")
	if err != nil {
		return err
	}

	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&contacts)
	if err != nil {
		return err
	}

	return nil
}

func addContact(reader *bufio.Reader, contacts *[]Contact) {
	var contact Contact
	fmt.Print("Nombre: ")
	contact.Name, _ = reader.ReadString('\n')
	fmt.Print("Correo Electrónico: ")
	contact.Email, _ = reader.ReadString('\n')
	fmt.Print("Teléfono: ")
	contact.Phone, _ = reader.ReadString('\n')
	*contacts = append(*contacts, contact)
	if err := saveContacts(*contacts); err != nil {
		fmt.Println("Error al guradar el contacto", err)
	}
}

func printContacts(contacts []Contact) {
	if len(contacts) == 0 {
		fmt.Println("No existen contactos.")
	} else {
		fmt.Println(">>>Lista de contactos:")
	}
	for index, contact := range contacts {
		fmt.Printf("%d. Nombre: %s Correo Electrónico:%s Teléfono: %s\n", index+1, contact.Name, contact.Email, contact.Phone)
	}
}

func main() {
	contacts := []Contact{}

	err := loadContacts(&contacts)
	if err != nil {
		fmt.Println("Error al cargar los contactos", err)
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("=== GESTOR DE CONTACTOS ===\n",
			"1. Agregar un contacto\n",
			"2. Mostrar todos los contactos\n",
			"3. Salir\n",
			"Elige una opción: ")

		var option int
		_, err = fmt.Scanln(&option)
		if err != nil {
			fmt.Println("Error al leer la opción", err)
			return
		}

		switch option {
		case 1:
			addContact(reader, &contacts)
		case 2:
			printContacts(contacts)
		case 3:
			return
		default:
			fmt.Println("Opción inválida")
		}
	}
}
