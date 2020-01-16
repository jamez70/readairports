package main

import (
	"fmt"
	"os"
	"strconv"
	"encoding/csv"
)

type Airport struct {
	ICAO string
	ApType string
	Description string
	LngRaw string
	Lng float64
	LatRaw string
	Lat float64
	ElevationRaw string
}

var airports[] Airport


func ReadCsv(filename string) ([][]string, error) {

    // Open CSV file
    f, err := os.Open(filename)
    if err != nil {
        return [][]string{}, err
    }
    defer f.Close()

    // Read File into a Variable
    reader := csv.NewReader(f)
    reader.Comma = '\t'
    reader.FieldsPerRecord=-1
    reader.LazyQuotes=true
    _,e2 := reader.Read()
    if e2 != nil {
    	return [][]string{}, e2
    }
    lines, err := reader.ReadAll()
    if err != nil {
        return [][]string{}, err
    }

    return lines, nil
}

func fixName(apname string, icao string) string {
	var st string
	if len(icao) > 1 {
		return icao
	}
	if apname[0] == '\'' {
		st = apname[1:]
	} else {
		st = apname
	}
	return st
}

func lnglatToFloat(lval string) float64 {
	
	sLen := len(lval)
	lastChar := lval[sLen-1]
	val,_ := strconv.ParseFloat(lval[:(sLen-1)],64)
	val = val / 3600.0
	if (lastChar == 'S') || (lastChar == 'W') {
		val = -val
	}
	return val
}

func readAirports(fname string) {
	lines, err := ReadCsv(fname)
	if err != nil {
		panic(err)
	}
	for _, line := range lines {
        data := Airport{
            ICAO: fixName(line[2],line[101]),
            ApType: line[1],
			Description: line[11],
			LngRaw: line[25],
			Lng: lnglatToFloat(line[25]),
			LatRaw: line[23],
			Lat: lnglatToFloat(line[23]),
		}	
		fmt.Printf("ICAO: %s %s %s %d\n",line[101],data.ICAO,line[103],len(line))
		airports = append(airports, data)
	}
}


func main() {
	readAirports("/var/www/airports.csv")
	ip := os.Getenv("REMOTE_ADDR")
	fmt.Printf("Content-type: text/html\n\n")
	fmt.Println("<html><body>")
	fmt.Println("<h3>Hello! from: "+ip+"</h3><pre>test123\n")
    // Loop through lines & turn into object

    for ap := range airports {
    	data := airports[ap]
		fmt.Printf("Airport: %s Desc '%s' Lat %f Lon %f\n",data.ICAO,data.Description,data.Lat,data.Lng)
    }
    
	fmt.Println("</pre></body></html>")
	os.Exit(0)
}
