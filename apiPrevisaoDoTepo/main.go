package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

type WeatherResponse struct {
	Main struct {
		Temp     float64 `json:"temp"`
		Humidity int `json:"humidity"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
	Name string `json:"name"`
}

func main() {
	err := godotenv.Load()
	if err!=nil{
		fmt.Println("Erro ao carregar.env")
	}

	apiKey:=os.Getenv("OPENWEATHER_API_KEY")
	if apiKey==""{
		log.Fatal("API não configurada!!")
	}

	if len(os.Args)<2{
		fmt.Println("Uso:go run main.go [cidade desejada a pesquisa]")
	}

	cidade:=os.Args[1]

	urlBasico:="https://api.openweathermap.org/data/2.5/weather"

	params:=url.Values{}

	params.Add("q",cidade)
	params.Add("appid",apiKey)
	params.Add("units","metric")
	params.Add("lang","pt")

	urlGet:=urlBasico+"?"+params.Encode()

	resposta,err:=http.Get(urlGet)

	if err!=nil{
		fmt.Println("Probelma chamando a API")
		fmt.Println(resposta.Status)
		
	}
	defer resposta.Body.Close()

	var weather WeatherResponse
	if err := json.NewDecoder(resposta.Body).Decode(&weather); err != nil {
		log.Fatalf("Erro ao decodificar JSON: %v", err)
	}

	fmt.Printf("\nPrevisão do Tempo para %s:\n", weather.Name)
	fmt.Printf("Temperatura: %.1f°C\n", weather.Main.Temp)
	fmt.Printf("Umidade: %d%%\n", weather.Main.Humidity)
	if len(weather.Weather) > 0 {
		fmt.Printf("Condição: %s\n", weather.Weather[0].Description)
	}




}