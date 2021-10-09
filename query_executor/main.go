package query_executor

import (
	"log"
)

type Movie struct {
	MovieID int
	Title   string
	Genres  string
}

var data = []Movie{
	{
		MovieID: 1,
		Title:   "Toy Story (1995)",
		Genres:  "Adventure|Animation|Children|Comedy|Fantasy",
	},
	{
		MovieID: 2,
		Title:   "Jumanji (1995)",
		Genres:  "Adventure|Children|Fantasy",
	},
	{
		MovieID: 3,
		Title:   "Grumpier Old Men (1995)",
		Genres:  "Comedy|Romance",
	},
	{
		MovieID: 4,
		Title:   "Waiting to Exhale (1995)",
		Genres:  "Comedy|Drama|Romance",
	},
}

func main() {
	//scanner := Scan{}
	//scanner.Init(&data, 2)
	//for i := 0; i < 6; i ++ {
	//	m := scanner.Next()
	//	if m == nil {
	//		log.Println("No more data")
	//	} else {
	//		log.Printf("Next movie is: %v", m.Title)
	//	}
	//}

	lim := Limit{}
	lim.Init(3)
	for i := 0; i < 6; i++ {
		m := lim.Next()
		if m == nil {
			log.Println("No more data, limit met")
		} else {
			log.Printf("Next movie is: %v", m.Title)
		}
	}
}
