package main 

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"net/url"
    
    "./print"

	"github.com/Sirupsen/logrus"
)

const (
	forecastAPIURI = "https://api.forecast.io/forecast"
)

var (
	forecastAPIKey string
	port     string

    latitude float64
    longitude float64
    units   string
)


// JSONResponse is a map[string]string
// response from the web server
type JSONResponse map[string]string

// String returns the string representation of the
// JSONResponse object
func (j JSONResponse) String() string {
	str, err := json.MarshalIndent(j, "", "  ")
	if err != nil {
		return fmt.Sprintf(`{
  "error": "%v"
}`, err)
	}

	return string(str)
}

// forecastHandler takes a forecast.Request object
// and passes it to the forecast.io API
func forecastHandler(w http.ResponseWriter, r *http.Request) {

//    xclude = []string{"daily", "minutely"}
// , "exclude: {string(xclude)}
    data := url.Values{"units": {units} }


	// request the forecast.io API
	url := fmt.Sprintf("%s/%s/%g,%g?%s", forecastAPIURI, forecastAPIKey, latitude, longitude, data.Encode())
	resp, err := http.Get(url)
	if err != nil {
		writeError(w, fmt.Sprintf("request to %s failed: %v", url, err))
		return
	}
	defer resp.Body.Close()
/*
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		writeError(w, fmt.Sprintf("reading response body from %s failed: %v", url, err))
		return
	}
    */
    var  forecast print.Forecast 
    dec := json.NewDecoder(resp.Body)
	if err = dec.Decode(&forecast); err != nil {
		return//forecast, fmt.Errorf("Decoding forecast response failed: %v", err)
	}

	if forecast.Error != "" {
		return// forecast, fmt.Errorf("Forecast API response error: %s", forecast.Error)
	}


	// write the response from the API to our client
	w.WriteHeader(resp.StatusCode)

    print.PrintCurrent(w, forecast, true)

//	if _, err := w.Write(body); err != nil {
//		writeError(w, fmt.Sprintf("writing response from %s failed: %v", url, err))
//		return
//	}

    

	return
}


// failHandler returns not a valid endpoint
func failHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, JSONResponse{
		"error": fmt.Sprintf("Not a valid endpoint: %s", r.URL.Path),
	})
	return
}

// writeError sends an error back to the requester
// and also logrus. the error
func writeError(w http.ResponseWriter, msg string) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, JSONResponse{
		"error": msg,
	})
	logrus.Printf("writing error: %s", msg)
	return
}

func init() {
	flag.StringVar(&forecastAPIKey, "forecast-apikey", "", "Key for forecast.io API")
	flag.StringVar(&port, "p", "8080", "port for server to run on")
    flag.Float64Var(&latitude, "lat", 0., "latutide for forecast.io")
    flag.Float64Var(&longitude, "long", 0., "longitude for forecast.io")
    flag.StringVar(&units, "u", "auto", "Units ( si, auto )")
	flag.Parse()

	if forecastAPIKey == "" {
		logrus.Fatalf("You need to pass a forecast.io API Key")
	}
    if longitude == 0 {
        logrus.Fatalf("You need to pass a longitude")
    }
    if latitude == 0  {
        logrus.Fatalf("You need to pass latitude")
    }
}

func main() {
	// create mux server
	mux := http.NewServeMux()

	mux.HandleFunc("/weather", forecastHandler) // forecast handler
	mux.HandleFunc("/", failHandler)             // everything else fail handler

	// set up the server
	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
	logrus.Infof("Starting server on port %q", port)
	logrus.Fatal(server.ListenAndServe())
}