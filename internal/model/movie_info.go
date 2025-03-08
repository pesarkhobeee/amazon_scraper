package model

// MovieInformation
// sample data:
// {

// "title": "Um Jeden Preis",

// "release_year": 2013,

// "actors": ["Dennis Quaid","Zac Efron"],

// "poster": "http://ecx.images-
// amazon.com/images/I/51UZ8st2OdL._SX200_QL80_.jpg",

// "similar_ids":
// ["B00SWDQPOC","B00RBPBO1G","B00S2EMECI","B00M5GH53M","B00IH8BA3S",
// "B00M5JP1DA"]

// }
type MovieInformation struct {
	Title       string   `json:"title"`
	ReleaseYear int      `json:"release_year"`
	Actors      []string `json:"actors"`
	Poster      string   `json:"poster"`
	SimilarIds  []string `json:"similar_ids"`
}
