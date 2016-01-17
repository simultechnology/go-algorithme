package main

import (
	"fmt"
)

type STATION struct {
	name        string
	transitions []*STATION
}

func (station *STATION) AddStation(transition *STATION) {
	// 既に登録されている場合はリターンする
	for _, s := range station.transitions {
		if s == transition {
			return
		}
	}
	//fmt.Printf("%s : %p\n", station.name, station)
	station.transitions = append(station.transitions, transition)
}

func (station *STATION) Print() {
	fmt.Printf("%s : ", station.name)
	for _, s := range station.transitions {
		fmt.Printf("->%s ", s.name)
	}
	fmt.Printf("\n")
}

func main() {
	//stations := make([]STATION, 6)
	var stations []*STATION

	kamakura := &STATION{name: "鎌倉"}
	stations = append(stations, kamakura)
	fujisawa := &STATION{name: "藤沢"}
	stations = append(stations, fujisawa)
	yokohama := &STATION{name: "横浜"}
	stations = append(stations, yokohama)
	yokosuka := &STATION{name: "横須賀"}
	stations = append(stations, yokosuka)
	chigasaki := &STATION{name: "茅ヶ崎"}
	stations = append(stations, chigasaki)
	tokyo := &STATION{name: "東京"}
	stations = append(stations, tokyo)

	//fmt.Printf("kamakura %p\n", kamakura)
	kamakura.AddStation(yokosuka)
	kamakura.AddStation(fujisawa)
	kamakura.AddStation(yokohama)

	fujisawa.AddStation(kamakura)
	fujisawa.AddStation(chigasaki)
	fujisawa.AddStation(yokosuka)

	yokohama.AddStation(fujisawa)
	yokohama.AddStation(kamakura)
	yokohama.AddStation(tokyo)

	yokosuka.AddStation(kamakura)
	chigasaki.AddStation(fujisawa)
	tokyo.AddStation(yokohama)

	for _, s := range stations {
		s.Print()
	}
}
