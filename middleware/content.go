package middleware

import (
	"net/http"

	"api.default.indicoinnovation.pt/config/constants"
	"api.default.indicoinnovation.pt/entity"
	"api.default.indicoinnovation.pt/pkg/helpers"
	"github.com/gofiber/fiber/v2"
)

func ValidateContentType() func(context *fiber.Ctx) error {
	return func(context *fiber.Ctx) error {
		if context.GetReqHeaders()["Accept"] == "" || !helpers.Contains(constants.AllowedContentTypes, context.GetReqHeaders()["Accept"]) {
			return helpers.CreateResponse(context, &entity.ErrorResponse{
				Message:    "Content Not Accepted",
				StatusCode: http.StatusNotAcceptable,
			}, http.StatusNotAcceptable)
		}

		return context.Next()
	}
}
