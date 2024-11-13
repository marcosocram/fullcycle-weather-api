package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var ErrCEPNotFound = errors.New("can not find zipcode")

type ViaCEPResponse struct {
	Localidade string `json:"localidade"`
}

func GetCidadeFromCEP(cep string) (string, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep), nil)
	if err != nil {
		fmt.Println("1: ", err)
		return "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("2: ", err)
		return "", err
	}
	defer resp.Body.Close()

	//resp, err := http.Get(fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep))
	//if err != nil {
	//	return "", err
	//}
	//defer resp.Body.Close()

	var data ViaCEPResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil || data.Localidade == "" {
		return "", ErrCEPNotFound
	}

	return data.Localidade, nil
}
