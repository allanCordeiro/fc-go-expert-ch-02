package main

import (
	"fmt"
	"multi-threading/entities"
	"os"
	"regexp"
	"strings"
	"time"
)

func main() {
	//default cep case user don't add cep in command args.
	cep := formatCEP("01211100")
	cepByArgs := os.Args[1:]
	if len(cepByArgs) > 0 {
		cep = formatCEP(cepByArgs[0])
	}
	if !isCepValid(cep) {
		panic(fmt.Sprintf("Entrada de CEP inválida. %s", cep))
	}
	apicepCh := make(chan entities.CepOutput)
	viacepCh := make(chan entities.CepOutput)

	go func() {
		apicep := entities.ApiCepOutput{}
		output, err := apicep.GetData(cep)
		if err != nil {
			panic(err)
		}
		apicepCh <- output
	}()

	go func() {
		viacep := entities.ViaCepOutput{}
		output, err := viacep.GetData(cep)
		if err != nil {
			panic(err)
		}
		viacepCh <- output
	}()

	select {
	case msg := <-apicepCh:
		fmt.Printf("Dado recebido de APICEP: %s", msg)
	case msg := <-viacepCh:
		fmt.Printf("Dado recebido de VIACEP: %s", msg)
	case <-time.After(time.Second * 3):
		print("Nenhuma das APIS respondeu dentro de 3 segundos.")
	}
}

// APICEP for some reason cannot recognize CEP without properly mask format #####-###
func formatCEP(cep string) string {
	if strings.ContainsAny(cep, "-") {
		return fmt.Sprintf("%08s", cep)
	}
	newCep := cep[:5] + "-" + cep[5:]
	return fmt.Sprintf("%08s", newCep)
}

func isCepValid(cep string) bool {
	re := regexp.MustCompile("^\\d{5}-\\d{3}$")
	return re.MatchString(cep)
}
