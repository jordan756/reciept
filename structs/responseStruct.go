package structs

type PointsResponse struct {
	Points int `json:"points"`
}
type IdResponse struct {
	Id string `json:"id"`
}

type ErrorResponse struct {
	Error string `json:"Error Message"`
}
