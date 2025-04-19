package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"path"
	"strconv"
)

func commandMap(config *config, input []string) error {
	for i := 0; i < 20; i++ {
		id := path.Base(config.next)
		val, ok := cache.Get(id)
		if ok {
			fmt.Println(string(val))
		} else {
			res, err := http.Get(config.next)
			if err != nil {
				return err
			}
			body, err := io.ReadAll(res.Body)
			res.Body.Close()
			if res.StatusCode > 299 {
				msg := fmt.Sprintf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
				return errors.New(msg)
			}
			if err != nil {
				return (err)
			}

			location := map[string]string{}
			json.Unmarshal(body, &location)

			fmt.Println(location["name"])
			cache.Add(id, []byte(location["name"]))
		}

		config.previous = config.next
		id = path.Base(config.previous)
		nextId, _ := strconv.Atoi(id)
		nextId += 1
		config.next = fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%d", nextId)
	}
	return nil
}
