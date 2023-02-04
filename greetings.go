package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error ){
	if name == ""{
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)  //one liner var
	return message, nil
}

func Hellos(names []string) (map[string]string, error){
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

func init(){
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string{
	formats := []string {
		"Hi, %v. Welcome!",
		"Great to see ya %v!",
		"I see you once agian %v!",
	}

	return formats[rand.Intn(len(formats))]
}

// var message string       other way to declare
// message = ftm.Sprintf("Hi, %v. Welcome!", name)

// func Hey(name string){
// 	message2 := ftm.Sprintf("Hey there, %v", name)
// 	return message2
// }