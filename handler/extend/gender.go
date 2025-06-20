package extend

import (
	"encoding/json"
	"fmt"
	"github.com/baiv84/personio/model"
	"io"
	"log"
	"net/http"
	"net/url"
)

type PersonGender struct {
	Count       int32
	Name        string
	Gender      string
	Probability float32
}

func ExtractGender(person *model.Person) {
	var gender string
	wrappedName := url.QueryEscape(person.FirstName)
	genderUrl := fmt.Sprintf("https://api.genderize.io/?name=%s", wrappedName)

	genderResp, err := http.Get(genderUrl)
	if err != nil {
		log.Fatal("error download name")
	}
	defer genderResp.Body.Close()

	body, _ := io.ReadAll(genderResp.Body)
	personGender := PersonGender{}
	json.Unmarshal(body, &personGender)

	if personGender.Gender == "male" && personGender.Probability < 0.5 {
		gender = "female"
	} else if personGender.Gender == "female" && personGender.Probability < 0.5 {
		gender = "male"
	} else {
		gender = personGender.Gender
	}

	person.Gender = gender

}
