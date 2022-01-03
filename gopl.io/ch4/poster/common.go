package poster

const queryURL = "https://www.omdbapi.com/?apikey=b166cb0b&t="

type Movie struct {
	Title                    string
	Year, Released, Runtime  string
	Rated, Genre             string
	Director, Writer, Actors string
	Plot                     string
	Language, Country        string
	Awards                   string
	Poster                   string
}
