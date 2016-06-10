package print

import (
	"fmt"
	"math"
	"time"
	"net/http"
)

// UnitMeasures are the location specific terms for weather data.
type UnitMeasures struct {
	Degrees       string
	Speed         string
	Length        string
	Precipitation string
}

var (
	// UnitFormats describe each regions UnitMeasures.
	UnitFormats = map[string]UnitMeasures{
		"us": {
			Degrees:       "°F",
			Speed:         "mph",
			Length:        "miles",
			Precipitation: "in/hr",
		},
		"si": {
			Degrees:       "°C",
			Speed:         "m/s",
			Length:        "kilometers",
			Precipitation: "mm/h",
		},
		"ca": {
			Degrees:       "°C",
			Speed:         "km/h",
			Length:        "kilometers",
			Precipitation: "mm/h",
		},
		// deprecated, use "uk2" in stead
		"uk": {
			Degrees:       "°C",
			Speed:         "mph",
			Length:        "kilometers",
			Precipitation: "mm/h",
		},
		"uk2": {
			Degrees:       "°C",
			Speed:         "mph",
			Length:        "miles",
			Precipitation: "mm/h",
		},
	}
	// Directions contain all the combinations of N,S,E,W
	Directions = []string{
		"N", "NNE", "NE", "ENE", "E", "ESE", "SE", "SSE", "S", "SSW", "SW", "WSW", "W", "WNW", "NW", "NNW",
	}
)

func epochFormat(seconds int64) string {
	epochTime := time.Unix(0, seconds*int64(time.Second))
	return epochTime.Format("January 2 at 3:04pm MST")
}

func epochFormatDate(seconds int64) string {
	epochTime := time.Unix(0, seconds*int64(time.Second))
	return epochTime.Format("January 2 (Monday)")
}

func epochFormatTime(seconds int64) string {
	epochTime := time.Unix(0, seconds*int64(time.Second))
	return epochTime.Format("3:04pm MST")
}

func getBearingDetails(degrees float64) string {
	index := int(math.Mod((degrees+11.25)/22.5, 16))
	return Directions[index]
}

func printCommon(w http.ResponseWriter, weather Weather, unitsFormat UnitMeasures) error {
	if weather.Humidity > 0 {
		humidity := fmt.Sprintf("%v%s", weather.Humidity*100, "%")
		fmt.Fprintf(w,"humidity:%s\n", humidity)
	}

	if weather.PrecipIntensity > 0 {
		precInt := fmt.Sprintf("%v %s", weather.PrecipIntensity, unitsFormat.Precipitation)
		fmt.Fprintf(w,"precipitationintensity:%s,%s\n", weather.PrecipType, precInt)
	}

	if weather.PrecipProbability > 0 {
		prec := fmt.Sprintf("%v%s", weather.PrecipProbability*100, "%")
		fmt.Fprintf(w,"precipitationprobability:%s\n", prec)
	}

	if weather.NearestStormDistance > 0 {
		dist := fmt.Sprintf("%v %s %v", weather.NearestStormDistance, unitsFormat.Length, getBearingDetails(weather.NearestStormBearing))
		fmt.Fprintf(w,"neareststrom:%s\n", dist)
	}

	if weather.WindSpeed > 0 {
		wind := fmt.Sprintf("%v %s", weather.WindSpeed, unitsFormat.Speed)
		fmt.Fprintf(w,"wind:%s\n", wind)
		fmt.Fprintf(w,"winddirection:%s\n",getBearingDetails(weather.WindBearing))
	}

	if weather.CloudCover > 0 {
		cloudCover := fmt.Sprintf("%v%s", weather.CloudCover*100, "%")
		fmt.Fprintf(w,"coverage:%s\n", cloudCover)
	}

	if weather.Visibility < 10 {
		visibility := fmt.Sprintf("%v %s", weather.Visibility, unitsFormat.Length)
		fmt.Fprintf(w,"visibility:%s\n", visibility)
	}

	if weather.Pressure > 0 {
		pressure := fmt.Sprintf("%v %s", weather.Pressure, "mbar")
		fmt.Fprintf(w,"pressure:%s\n", pressure)
	}

	return nil
}

// PrintCurrent pretty prints the current forecast data.
func PrintCurrent(w http.ResponseWriter, forecast Forecast, ignoreAlerts bool) error {
	unitsFormat := UnitFormats[forecast.Flags.Units]

	/* icon, err := getIcon(forecast.Currently.Icon)
	if err != nil {
		return err
	}

	fmt.Println(icon)
*/
	//location := colorstring.Color(fmt.Sprintf("[green]%s in %s", geolocation.City, geolocation.Region))
	// var location ="your home"

	fmt.Fprintf(w, "time:%s\n",epochFormat(forecast.Currently.Time))
	fmt.Fprintf(w, "weather:%s\n", forecast.Currently.Summary)

	temp := fmt.Sprintf("%v%s", forecast.Currently.Temperature, unitsFormat.Degrees)
	feelslike := fmt.Sprintf("%v%s", forecast.Currently.ApparentTemperature, unitsFormat.Degrees)
	if temp == feelslike {
		fmt.Fprintf(w,"temperature:%s\n", temp)
	} else {
		fmt.Fprintf(w,"feeledtemperature:%s\n", feelslike)
	}

	if !ignoreAlerts {
		for _, alert := range forecast.Alerts {
			if alert.Title != "" {
				fmt.Fprintf(w,"alert:%s\n",alert.Title)
			}
//			if alert.Description != "" {
//				fmt.Fprintf(w,"%s\n", alert.Description)
//			}
//			fmt.Println(w,"\t\t\t" + "Created: "+epochFormat(alert.Time))
//			fmt.Println(w,"\t\t\t" + "Expires: "+epochFormat(alert.Expires) + "\n")
		}
	}

	return printCommon(w, forecast.Currently, unitsFormat)
}

/*// PrintDaily pretty prints the daily forecast data.
func PrintDaily(w http.ResponseWriter, forecast Forecast, days int) error {
	unitsFormat := UnitFormats[forecast.Flags.Units]

	fmt.Println(colorstring.Color("\n" + fmt.Sprintf("%v Day Forecast", days)))

	// Ignore the current day as it's printed before
	for index, daily := range forecast.Daily.Data[1:] {
		// only do the amount of days they request
		if index == days {
			break
		}

		fmt.Println(colorstring.Color("\n[magenta]" + epochFormatDate(daily.Time)))

		tempMax := colorstring.Color(fmt.Sprintf("[blue]%v%s", daily.TemperatureMax, unitsFormat.Degrees))
		tempMin := colorstring.Color(fmt.Sprintf("[blue]%v%s", daily.TemperatureMin, unitsFormat.Degrees))
		feelsLikeMax := colorstring.Color(fmt.Sprintf("[cyan]%v%s", daily.ApparentTemperatureMax, unitsFormat.Degrees))
		feelsLikeMin := colorstring.Color(fmt.Sprintf("[cyan]%v%s", daily.ApparentTemperatureMin, unitsFormat.Degrees))
		fmt.Fprintf(w,"The temperature high is %s, feels like %s around %s, and low is %s, feels like %s around %s\n\n", tempMax, feelsLikeMax, epochFormatTime(daily.TemperatureMaxTime), tempMin, feelsLikeMin, epochFormatTime(daily.TemperatureMinTime))

		printCommon(w,daily, unitsFormat)
	}

	return nil
}
*/