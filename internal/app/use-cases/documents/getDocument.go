package documents

import (
	"github.com/google/wire"
	"github.com/luminosita/sample-bee/internal/interfaces/respositories/documents"
	documents2 "github.com/luminosita/sample-bee/internal/interfaces/use-cases/documents"
)

var GetWireSet = wire.NewSet(NewGetDocument,
	wire.Bind(new(documents2.GetDocumenter), new(*GetDocument)))

type GetDocument struct {
	repo documents.GetDocumentRepositorer
}

func NewGetDocument(r documents.GetDocumentRepositorer) *GetDocument {
	return &GetDocument{
		repo: r,
	}
}

func (d *GetDocument) Execute(
	documentData *documents2.GetDocumenterRequest) (*documents2.GetDocumenterResponse, error) {
	data := &documents.GetDocumentRepositorerRequest{
		DocumentID: documentData.DocumentId,
	}

	res, err := d.repo.GetDocument(data)

	if err != nil {
		return nil, err
	}

	return &documents2.GetDocumenterResponse{
		Document: res.Document,
	}, nil
}
