# README #

This is a microservice written in go providing weather data from forecast.io
To use this service you must create a developer account on forecast.io to get an API Key.

### How do I get set up? ###

* Clone the repo
* Set the GOPATH to the cloned repo directory
* From bitbucket-pipelines run the commands under the "script" node
* Register a [forecast.io account|https://developer.forecast.io/register]
* Copy the API Key from the lower part of [the page|https://developer.forecast.io/]
* Run the weather service using the coordinates from Loxone Configuration and the API Key: 

```
#!bash
LoxoneGoWeather -forecast-apikey $YOURAPIKEY -long $LOXONE_LONGITUDE lat $LOXONE_LATITUDE{code}

```

* Afterwards you should be able to send requests for http://localhost:8080/weather 

Within Loxone you can use 1 Virtual HTTP Input to query all the data with a single http query and parse it afterwards
into seperate variables

### Example Output ###

```
#!csv
time:June 22 at 9:39am CEST
weather:Clear
temperature:19.62°C
humidity:63%
wind:0.38 m/s
winddirection:ENE
coverage:18%
visibility:0 kilometers
pressure:1025.2 mbar

```

### Stuff todo ###
* Caching data if the internet connection goes down
* Querying of real forecasts e.g. 1h ahead of time 

### Who do I talk to? ###

* Repo owner or admin
* Other community or team contact