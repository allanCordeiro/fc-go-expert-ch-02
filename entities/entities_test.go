package entities

import (
	"testing"
)

func TestApiCep(t *testing.T) {
	t.Run("Given a valid CEP, when call APICEP api, then should return a valid data", func(t *testing.T) {
		apicep := ApiCepOutput{}
		expectedData := CepOutput{
			CEP:      "04266-000",
			Address:  "Rua Padre Marchetti",
			District: "Ipiranga",
			City:     "São Paulo",
			State:    "SP",
		}
		expectedCep := "04266-000"

		output, err := apicep.GetData(expectedCep)
		if err != nil {
			t.Errorf("unexpected error, %s", err.Error())
		}

		if expectedData.CEP != output.CEP {
			t.Errorf("CEP Inválido. Esperado %s | recebido %s", expectedData.CEP, output.CEP)
		}
		if expectedData.Address != output.Address {
			t.Errorf("Address Inválido. Esperado %s | recebido %s", expectedData.Address, output.Address)
		}
		if expectedData.District != output.District {
			t.Errorf("District Inválido. Esperado %s | recebido %s", expectedData.District, output.District)
		}
		if expectedData.City != output.City {
			t.Errorf("City Inválido. Esperado %s | recebido %s", expectedData.City, output.City)
		}
		if expectedData.State != output.State {
			t.Errorf("State Inválido. Esperado %s | recebido %s", expectedData.State, output.State)
		}
	})
}

func TestViaCep(t *testing.T) {
	t.Run("Given a valid CEP, when call VIACEP api, then should return a valid data", func(t *testing.T) {
		apicep := ViaCepOutput{}
		expectedData := CepOutput{
			CEP:      "04266-000",
			Address:  "Rua Padre Marchetti",
			District: "Ipiranga",
			City:     "São Paulo",
			State:    "SP",
		}
		expectedCep := "04266-000"

		output, err := apicep.GetData(expectedCep)
		if err != nil {
			t.Errorf("unexpected error, %s", err.Error())
		}

		if expectedData.CEP != output.CEP {
			t.Errorf("CEP Inválido. Esperado %s | recebido %s", expectedData.CEP, output.CEP)
		}
		if expectedData.Address != output.Address {
			t.Errorf("Address Inválido. Esperado %s | recebido %s", expectedData.Address, output.Address)
		}
		if expectedData.District != output.District {
			t.Errorf("District Inválido. Esperado %s | recebido %s", expectedData.District, output.District)
		}
		if expectedData.City != output.City {
			t.Errorf("City Inválido. Esperado %s | recebido %s", expectedData.City, output.City)
		}
		if expectedData.State != output.State {
			t.Errorf("State Inválido. Esperado %s | recebido %s", expectedData.State, output.State)
		}
	})
}
