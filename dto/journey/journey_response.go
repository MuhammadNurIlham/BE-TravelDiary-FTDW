package journeydto

type JourneyResponse struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
	Image string `json:"image"`
	User  int    `json:"user"`
}
