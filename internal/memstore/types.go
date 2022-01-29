package memstore

type Taco struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	ImageUrl    string  `json:"image_url"`
	Price       float32 `json:"price"`
}

type Address struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Address    string `json:"address"`
	City       string `json:"city"`
	State      string `json:"state"`
	PostalCode string `json:"postal_code"`
	Country    string `json:"country"`
	SaveInfo   bool   `json:"save_info"`
}

type Movie struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Genre       string `json:"genre"`
	ImgUrl      string `json:"imgUrl"`
	Description string `json:"description"`
	ReleaseDate int    `json:"releaseDate"`
	Length      string `json:"length"`
	Likes       int    `json:"likes"`
	Comments    int    `json:"comments"`
}

type User struct {
	ID       string `json:"id"`
	Username string `json:"name"`
	Password string `json:"password"`
}
