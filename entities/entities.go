package entities

import (
	"encoding/json"
	"io"
	"net/http"
)

type OutData interface {
	GetData() (CepOutput, error)
}
type ApiCepOutput struct {
	Code       string `json:"code"`
	State      string `json:"state"`
	City       string `json:"city"`
	District   string `json:"district"`
	Address    string `json:"address"`
	Status     int    `json:"status"`
	Ok         bool   `json:"ok"`
	StatusText string `json:"statusText"`
}

func (o *ApiCepOutput) GetData(cep string) (CepOutput, error) {
	url := "https://cdn.apicep.com/file/apicep/" + cep + ".json"
	body, err := getDataFromHttp(url)

	err = json.Unmarshal(body, &o)
	if err != nil {
		return CepOutput{}, err
	}
	output := CepOutput{
		CEP:      o.Code,
		Address:  o.Address,
		District: o.District,
		City:     o.City,
		State:    o.State,
	}

	return output, nil
}

type ViaCepOutput struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func (o *ViaCepOutput) GetData(cep string) (CepOutput, error) {
	url := "https://viacep.com.br/ws/" + cep + "/json"
	body, err := getDataFromHttp(url)

	err = json.Unmarshal(body, &o)
	if err != nil {
		return CepOutput{}, err
	}
	output := CepOutput{
		CEP:      o.Cep,
		Address:  o.Logradouro,
		District: o.Bairro,
		City:     o.Localidade,
		State:    o.Uf,
	}
	return output, nil
}

type CepOutput struct {
	CEP      string
	Address  string
	District string
	City     string
	State    string
}

func getDataFromHttp(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
