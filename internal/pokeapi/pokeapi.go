package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func MakePokeAPIRequest[T any](url string, structPointer *T) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	err = bodyToStruct(body, structPointer)
	if err != nil {
		return fmt.Errorf("error while turning body to struct")
	}

	return nil
}

func bodyToStruct[T any](body []byte, structPointer *T) error {
	if err := json.Unmarshal(body, structPointer); err != nil {
		return err
	}
	return nil
}
