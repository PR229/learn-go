package greetings

import( "fmt"
	"errors"
	"math/rand"
)

func Hellos(names []string) (map[string]string , error){
	messages := make(map[string]string)
	for _, name := range names{
	message, err := Hello(name)
	if err != nil{
		return nil, err
		}
	messages[name] = message
	} 
	return messages, nil
}


func Hello(name string) (string, error) {
	
	if name == ""{
		return "", errors.New("Empty name")
}
	message := fmt.Sprintf(randomFormat(), name)
	// message := fmt.Sprintf(randomFormat())
	return message, nil
}


func randomFormat() string {
	formats := []string{
		"Hey, %v. Welcome!",
		"Piacere %v!",
		"Hello hello hello %v",
	}
	return formats[rand.Intn(len(formats))]
}
