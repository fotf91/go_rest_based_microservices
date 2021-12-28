package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

type Time struct {
	CurrentTime string `json:"current_time"`
}

func getTime(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if len(vars) == 0 {
		timeutc := Time{CurrentTime: time.Now().UTC().String()}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(timeutc)
	} else {
		strings.Split(vars["tz"], ",")
		splitVars := strings.Split(vars["tz"], ",")
		timesList := []Time{}

		for _, tz := range splitVars {
			fmt.Println(tz)
			loc, err := time.LoadLocation(tz)
			if err != nil {
				http.Error(w, "invalid timezone", http.StatusNotFound)
				return
			} else {
				w.Header().Add("Content-Type", "application/json")
				timesList = append(timesList, Time{CurrentTime: time.Now().In(loc).String()})

				// json.NewEncoder(w).Encode({tz: time.Now().In(loc)})

				// json.NewEncoder(w).Encode(`{` + tz + `: ` + time.Now().In(loc).String() + `}`)

				// currentTz := `{` + tz + `: ` + time.Now().In(loc).String() + `}`

				// currentTzBytes := []byte(currentTz)
				// json.NewEncoder(w).Encode(currentTzBytes)

				// jsonParsed, _ := gabs.ParseJSON([]byte(`{` + tz + `: ` + time.Now().In(loc).String() + `}`))
				// respMap := make([]map[string]string, len(vars))
				// respMap[tz] = time.Now().In(loc).String()

				// jsonObj := gabs.New()
				// jsonObj.Set(10, "outter", "inner", "value")
				// jsonObj.SetP(20, "outter.inner.value2")
				// jsonObj.Set(30, "outter", "inner2", "value3")
				// fmt.Printf(jsonObj.String())
				// json.NewEncoder(w).Encode(jsonObj.String())

				// WORKING!!!
				// json.NewEncoder(w).Encode(map[string]string{tz: time.Now().In(loc).String()})
			}
		}
		json.NewEncoder(w).Encode(timesList)
	}

}
