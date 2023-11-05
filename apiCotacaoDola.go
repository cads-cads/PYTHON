
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ExchangeRate struct {
	Rates map[string]float64 `json:"rates"`
}

func main() {
	// Substitua 'YOUR_API_KEY' pela sua chave de API real
	apiKey := "YOUR_API_KEY"

	url := fmt.Sprintf("https://api.apilayer.com/exchangerates_data/latest?base=USD&access_key=%s", apiKey)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Erro ao fazer a solicitação HTTP:", err)
		return
	}
	defer resp.Body.Close()

	var exchangeRate ExchangeRate
	err = json.NewDecoder(resp.Body).Decode(&exchangeRate)
	if err != nil {
		fmt.Println("Erro ao decodificar a resposta JSON:", err)
		return
	}

	// Substitua "BRL" pela moeda desejada, se necessário
	cotacaoDolar := exchangeRate.Rates["BRL"]
	fmt.Printf("A cotação atual do dólar para BRL é: %.2f\n", cotacaoDolar)
}
