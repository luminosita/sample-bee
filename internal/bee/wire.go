//go:build wireinject
// +build wireinject

package bee

import (
	"context"
	"github.com/google/wire"
	documents2 "github.com/luminosita/sample-bee/internal/app/use-cases/documents"
	"github.com/luminosita/sample-bee/internal/infra/db/mongodb/repositories"
	"github.com/luminosita/sample-bee/internal/infra/http/handlers/documents"
)

func MakeGetDocumentHandler(ctx context.Context) *documents.GetDocumentHandler {
	wire.Build(documents.GetWireSet, documents2.GetWireSet, repositories.GetWireSet)

	return nil
}

func MakeGetAllDocumentsHandler(ctx context.Context) *documents.GetAllDocumentsHandler {
	wire.Build(documents.GetAllWireSet, documents2.GetAllWireSet, repositories.GetAllWireSet)

	return nil
}

func MakeCreateDocumentHandler(ctx context.Context) *documents.CreateDocumentHandler {
	wire.Build(documents.CreateWireSet, documents2.CreateWireSet, repositories.CreateWireSet)

	return nil
}
