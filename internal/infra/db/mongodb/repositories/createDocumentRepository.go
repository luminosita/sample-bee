package repositories

import (
	"context"
	"github.com/google/wire"
	"github.com/luminosita/honeycomb/pkg/infra/db/mongodb"
	"github.com/luminosita/sample-bee/internal/interfaces/respositories/documents"
)

var CreateWireSet = wire.NewSet(NewCreateDocumentRepository,
	wire.Bind(new(documents.CreateDocumentRepositorer), new(*CreateDocumentRepository)))

type CreateDocumentRepository struct {
	ctx context.Context
}

func NewCreateDocumentRepository(ctx context.Context) *CreateDocumentRepository {
	return &CreateDocumentRepository{
		ctx: ctx,
	}
}

func (r *CreateDocumentRepository) CreateDocument(
	docData *documents.CreateDocumentRepositorerRequest) (*documents.CreateDocumentRepositorerResponse, error) {
	col := mongodb.GetDbCollection(r.ctx, DOCUMENTS)
	_, err := col.InsertOne(r.ctx, docData) //, createdAt: new Date());
	if err != nil {
		return nil, err
	}

	return nil, nil
}
