package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error ){
	if name == ""{
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf("Hi, %v. Welcome!", name)  //one liner var
	return message, nil
}

// var message string       other way to declare
// message = ftm.Sprintf("Hi, %v. Welcome!", name)

// func Hey(name string){
// 	message2 := ftm.Sprintf("Hey there, %v", name)
// 	return message2
// }