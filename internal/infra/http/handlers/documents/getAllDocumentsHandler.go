package documents

import (
	"github.com/google/wire"
	"github.com/luminosita/honeycomb/pkg/http/ctx"
	"github.com/luminosita/honeycomb/pkg/http/handlers"
	"github.com/luminosita/sample-bee/internal/interfaces/use-cases/documents"
)

var GetAllWireSet = wire.NewSet(NewGetAllDocumentsHandler,
	wire.Bind(new(handlers.Handler), new(*GetAllDocumentsHandler)))

type GetAllDocumentsHandler struct {
	cd documents.GetAllDocumenter
}

func NewGetAllDocumentsHandler(cd documents.GetAllDocumenter) *GetAllDocumentsHandler {
	return &GetAllDocumentsHandler{
		cd: cd,
	}
}

// GetAllDocuments godoc
// @Summary      Get All Documents
// @Description  dummy GET method
// @Tags         documents
// @Accept       json
// @Produce      json
// @Success      200  {object}  string
// @Failure      400  {object}  error
// @Failure      404  {object}  error
// @Failure      500  {object}  error
// @Router       /documents [get]
func (h *GetAllDocumentsHandler) Handle(ctx *ctx.Ctx) error {
	res, err := h.cd.Execute(&documents.GetAllDocumenterRequest{})

	if err != nil {
		return err
	}

	return ctx.SendResponse(res.Documents)
}
