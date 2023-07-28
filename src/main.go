package main

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Main function for the action
func Main(obj map[string]interface{}) map[string]interface{} {
	name, ok := obj["name"].(string)
	if !ok {
		name = "world"
	}
	fmt.Println("Love, " + name + "!")
	log.Debug().Str("name", name).Msg("Love")
	msg := make(map[string]interface{})
	msg["body"] = "Love, " + name + "!"
	msg["statusCode"] = 200
	msg["headers"] = map[string]string{"Content-Type": "text/plain"}
	return msg
}
