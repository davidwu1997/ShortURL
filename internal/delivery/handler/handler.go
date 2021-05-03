package handler

import (
	"net/http"
	"shortURL/config"
	"shortURL/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	ShortURL *service.ShortURL
}

type UploadRequest struct {
	Url string `json:"url"`
}

func (h *Handler) UploadURL(context *gin.Context) {
	request := &UploadRequest{}
	if err := context.BindJSON(request); err != nil {
		context.String(http.StatusNotFound, "url can not receive")
		return
	}

	urlID, err := h.ShortURL.Upload(context, request.Url)
	if err != nil {
		context.JSON(http.StatusNotFound, nil)
		return
	}

	response := &service.ShortUrlResponse{
		Url_ID:   urlID,
		ShortUrl: config.Get().ShortURL.BasePath + urlID,
	}

	context.JSON(http.StatusOK, response)
}

func (h *Handler) DeleteURL(context *gin.Context) {
	url, ok := context.Params.Get("url")
	if !ok {
		context.JSON(http.StatusNotFound, nil)
		return
	}

	if err := h.ShortURL.Delete(context, url); err != nil {
		context.JSON(http.StatusNotFound, nil)
		return
	}

	context.JSON(http.StatusOK, nil)
}

func (h *Handler) RedirectURL(context *gin.Context) {
	url, ok := context.Params.Get("url")
	if !ok {
		context.JSON(http.StatusNotFound, nil)
		return
	}

	originalUrl, err := h.ShortURL.Redirect(context, url)
	if err != nil {
		context.JSON(http.StatusNotFound, nil)
		return
	}
	context.Redirect(302, originalUrl)
}
