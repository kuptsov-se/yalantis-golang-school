package main

import (
	encrypting "app/encrypting"
	translation "app/translation"
	"fmt"
)

func main() {
	inputString := "ANY! raNDoM,striNG foooorr encodE; decode AND TranSlate to PiG?Latin"
	fmt.Println(inputString)
	fmt.Scanln(&inputString)
	encrypter := encrypting.GetEncrypter()
	encruptedString := encrypter.Encrypt(inputString)
	fmt.Println(encruptedString)
	decriptedString := encrypter.Decrypt(encruptedString)
	fmt.Println(decriptedString)
	translater := translation.GetTranslator()
	translatedString := translater.Translate(inputString)
	fmt.Println(translatedString)
}
