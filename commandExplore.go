package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func commandExplore(config *config, input []string) error {
	if len(input) == 0 {
		return errors.New("Can not explore without a location")
	}
	location := input[0]
	fmt.Printf("Exploring %s...\n", location)

	val, ok := cache.Get(location)
	if ok {
		var locationArea LocationArea
		err := json.Unmarshal(val, &locationArea)
		if err != nil {
			msg := fmt.Sprintf("Unmarshal error : %s", err)
			return errors.New(msg)
		}

		fmt.Println("Found Pokemon:")
		for _, val := range locationArea.PokemonEncounters {
			fmt.Printf("- %s\n", val.Pokemon.Name)
		}
	} else {
		api_location := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", location)
		res, err := http.Get(api_location)
		if err != nil {
			return err
		}
		defer res.Body.Close()
		if res.StatusCode > 299 {
			msg := fmt.Sprintf("This location doesn't exist!")
			return errors.New(msg)
		}

		var locationArea LocationArea
		decoder := json.NewDecoder(res.Body)
		if err = decoder.Decode(&locationArea); err != nil {
			msg := fmt.Sprintf("Decoder error : %s", err)
			return errors.New(msg)
		}

		fmt.Println("Found Pokemon:")
		for _, val := range locationArea.PokemonEncounters {
			fmt.Printf("- %s\n", val.Pokemon.Name)
		}
		data, err := json.Marshal(locationArea)
		if err != nil {
			return fmt.Errorf("error with Marshal: %s", err)
		}
		cache.Add(location, []byte(data))
	}
	return nil
}

// used json to go for this
type LocationArea struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int           `json:"chance"`
				ConditionValues []interface{} `json:"condition_values"`
				MaxLevel        int           `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}
