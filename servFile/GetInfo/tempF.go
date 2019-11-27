func GetInfo() {

	client := &http.Client{}
	req, err := http.NewRequest(
		"GET", "https://api.patephone.com/client-api/books/13711", nil,
	)
	req.Header.Add("X-CLIENT-IDENTIFIER", "patephone_unlim_android")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("err %s", err)
		os.Exit(1)
	}
	fmt.Printf("content: %s\n", string(contents))
}

type Images struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}
type Role struct {
	Name string `json:"name"`
	Abbr string `json:"abbr"`
}

type Book struct {
	Authors          []Authors `json:"authors"`
	DateCreated      int       `json:"date_created"`
	DateUpdated      int       `json:"date_updated"`
	Description      string    `json:"description"`
	Duration         int       `json:"duration"`
	FileSize         int       `json:"file_size"`
	ID               int       `json:"id"`
	Images           []Images  `json:"images"`
	LastModification int       `json:"last_modification"`
	M4bFileSize      int       `json:"m4b_file_size"`
	Mp3PreviewURL    string    `json:"mp3_preview_url"`
	PreviewStreamURL string    `json:"preview_stream_url"`
	PriceTier        int       `json:"price_tier"`
	Publish          string    `json:"publish"`
	Rating           int       `json:"rating"`
	ShortDescription string    `json:"short_description"`
	Title            string    `json:"title"`
}

type Authors struct {
	ID         int    `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	MiddleName string `json:"middle_name"`
	Role       Role   `json:"role"`
}