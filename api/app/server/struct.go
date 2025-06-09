package server

type server struct {
	Port int `json:"port"`
}

// Response structures
type ListResponse struct {
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	UUID        string `json:"uuid"`
}

type ItemResponse struct {
	Name string `json:"name"`
	UUID string `json:"uuid"`
}

type DataResponse struct {
	Data interface{} `json:"data"`
}

type CreateListRequest struct {
	Name        string `json:"name"`
	DisplayName string `json:"display_name,omitempty"`
}
