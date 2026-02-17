package requests

type Author struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Rating struct  {
	Average float64 `json:"average"`
} 

type BookSummary struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Image    string `json:"image"`
	Authors   []Author `json:"authors"`
	Rating    *Rating 	`json:"rating"`
}

type Books struct {
	Available int    `json:"available"`
	BookSummary [][]BookSummary `json:"books"`
}

type Book struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Image    string `json:"image"`
	Authors   []Author `json:"authors"`
	PublishDate   float64 `json:"publish_date"`
	NumberOfPages float64 `json:"number_of_pages"`
	Description   string  `json:"description"`
	Rating    *Rating 	`json:"rating"`
}