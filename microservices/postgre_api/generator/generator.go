package generator

import (
	"encoding/json"
	"net/http"
)

var myClient = &http.Client{}

var names_male = []string{
	"Михаил",
	"Олег",
	"Вячеслав",
	"Денис",
	"Леонид",
	"Матвей",
	"Евгений",
}

var names_female = []string{
	"Ольга",
	"Татьяна",
	"Евгения",
	"Екатерина",
	"Елена",
	"Александра",
	"Мария",
}

func generate(count int) {
	for i := 0; i < count; i++ {
		if err := getJson("localhost:4042")
	}
}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
