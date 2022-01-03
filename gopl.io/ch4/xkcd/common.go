package xkcd

const queryURL = "https://xkcd.com/"

type commicSearchResult struct {
	Transcript string
	Img        string
	Title      string
	Num        int
}
