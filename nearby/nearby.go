package nearby

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)
type RequestBody struct {
	Business string `json:"business"`
	Latitude string `json:"latitude"`
	Longitude string `json:"longitude"`
	Lang string `json:"lang"`
	Region string `json:"region"`
}
func HandleNearbyPlaces (w http.ResponseWriter, r *http.Request) {
	var requestBody RequestBody 
	
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		fmt.Println(err.Error())
	}

	
	url := fmt.Sprintf("https://local-business-data.p.rapidapi.com/search-nearby?query=%s&lat=%s&lng=%s&limit=20&language=%s&region=%s",
		requestBody.Business, requestBody.Latitude, requestBody.Longitude, requestBody.Lang, requestBody.Region)
	fmt.Println(url)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-key", "f0a326144cmshcffe5c82b870ae7p130023jsn0736c64bc5f1")
	req.Header.Add("x-rapidapi-host", "local-business-data.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body) 
	if err != nil {
		http.Error(w, "Problem with reading the data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.StatusCode)
	w.Write(body)
}	
