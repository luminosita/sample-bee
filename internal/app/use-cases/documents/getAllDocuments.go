package documents

import (
	"github.com/google/wire"
	"github.com/luminosita/sample-bee/internal/interfaces/respositories/documents"
	documents2 "github.com/luminosita/sample-bee/internal/interfaces/use-cases/documents"
)

var GetAllWireSet = wire.NewSet(NewGetAllDocuments,
	wire.Bind(new(documents2.GetAllDocumenter), new(*GetAllDocuments)))

type GetAllDocuments struct {
	repo documents.GetAllDocumentsRepositorer
}

func NewGetAllDocuments(r documents.GetAllDocumentsRepositorer) *GetAllDocuments {
	return &GetAllDocuments{
		repo: r,
	}
}

func (d *GetAllDocuments) Execute(
	_ *documents2.GetAllDocumenterRequest) (*documents2.GetAllDocumenterResponse, error) {
	data := &documents.GetAllDocumentsRepositorerRequest{}

	res, err := d.repo.GetAllDocuments(data)

	if err != nil {
		return nil, err
	}

	return &documents2.GetAllDocumenterResponse{
		Documents: res.Documents,
	}, nil
}
