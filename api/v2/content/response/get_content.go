package response

import "hexa/business/content"

type GetContents struct {
	Contents []*GetContentbyIDResponse `json:"contents"`
}

func NewGetContents(contentList []content.Content) *GetContents {
	contentResponse := make([]*GetContentbyIDResponse, 0)

	for _, value := range contentList {
		contentResponse = append(contentResponse, NewGetByIDResponse(value))
	}

	return &GetContents{
		Contents: contentResponse,
	}

}
