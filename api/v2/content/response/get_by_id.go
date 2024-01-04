package response

import "hexa/business/content"

type GetContentbyIDResponse struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Tipe  string `json:"tipe"`
}

func NewGetByIDResponse(content content.Content) *GetContentbyIDResponse {
	var contentResponse GetContentbyIDResponse
	contentResponse.ID = content.ID
	contentResponse.Title = content.Name
	contentResponse.Tipe = content.Description

	return &contentResponse
}
