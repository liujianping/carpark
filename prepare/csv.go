package prepare

import (
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/gocarina/gocsv"
	"github.com/x-mod/errors"
)

//CsvItem for csv item
type CsvItem struct {
	CarParkNo           string  `csv:"car_park_no"`
	Address             string  `csv:"address"`
	Xcoord              float64 `csv:"x_coord"`
	Ycoord              float64 `csv:"y_coord"`
	CarParkType         string  `csv:"car_park_type"`
	TypeOfParkingSystem string  `csv:"type_of_parking_system"`
	ShortTermParking    string  `csv:"short_term_parking"`
	FreeParking         string  `csv:"free_parking"`
	NightParking        string  `csv:"night_parking"`
	CarParkDecks        int32   `csv:"car_park_decks"`
	GantryHeight        float64 `csv:"gantry_height"`
	CarParkBasement     string  `csv:"car_park_basement"`
}

//Parse csv io.reader
func Parse(rd io.Reader) ([]*CsvItem, error) {
	items := []*CsvItem{}
	if err := gocsv.Unmarshal(rd, &items); err != nil {
		return nil, errors.Annotate(err, "csv unmarshal")
	}
	return items, nil
}

func parkingPeriod(parktime string) (time.Duration, time.Duration) {
	switch strings.ToUpper(parktime) {
	case "NO":
		return time.Duration(0), time.Duration(0)
	case "WHOLE DAY":
		return time.Duration(0), time.Second * 24 * 60 * 60
	}
	ft := strings.Split(strings.ToUpper(parktime), "-")
	if len(ft) != 2 {
		return time.Duration(0), time.Duration(0)
	}
	return parseDuration(ft[0]), parseDuration(ft[1])
}

//s format: 7AM-10.30PM
func parseDuration(s string) time.Duration {
	pm := false
	if strings.HasSuffix(s, "PM") {
		pm = true
	}
	ns := strings.TrimSuffix(s, "AM")
	ns = strings.TrimSuffix(ns, "PM")
	ms := strings.Split(ns, ".")
	if len(ms) == 1 {
		h, err := strconv.Atoi(ms[0])
		if err != nil {
			panic(err)
		}
		if pm {
			h = h + 12
		}
		d, _ := time.ParseDuration(fmt.Sprintf("%dh", h))
		return d
	}
	if len(ms) == 2 {
		h, err := strconv.Atoi(ms[0])
		if err != nil {
			panic(err)
		}
		if pm {
			h = h + 12
		}
		m, err := strconv.Atoi(ms[1])
		if err != nil {
			panic(err)
		}
		d, _ := time.ParseDuration(fmt.Sprintf("%dh%dm", h, m))
		return d
	}
	return time.Duration(0)
}
