package documents

import (
	"encoding/json"
	"github.com/luminosita/honeycomb/pkg/http/ctx"
	"github.com/luminosita/sample-bee/internal/domain/entities"
	"github.com/luminosita/sample-bee/internal/interfaces/use-cases/documents"
)

type CreateDocumentHandler struct {
	cd documents.CreateDocumenter
}

func NewCreateDocumentHandler(cd documents.CreateDocumenter) *CreateDocumentHandler {
	return &CreateDocumentHandler{
		cd: cd,
	}
}

// CreateDocument godoc
// @Summary      Create Document
// @Description  dummy POST method
// @Tags         documents
// @Accept       json
// @Produce      json
// @Success      200  {object}  string
// @Failure      400  {object}  error
// @Failure      404  {object}  error
// @Failure      500  {object}  error
// @Router       /documents [post]
func (h *CreateDocumentHandler) Handle(ctx *ctx.Ctx) error {
	doc := &entities.Document{}

	err := json.Unmarshal(ctx.Body, doc)
	if err != nil {
		return err
	}

	res, err := h.cd.Execute(&documents.CreateDocumenterRequest{
		Document: doc,
	})

	if err != nil {
		return err
	}

	return ctx.SendString(res.DocumentId)
}
