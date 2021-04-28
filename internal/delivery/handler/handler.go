package handler

import (
	"fmt"
	"net/http"
	"shortURL/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	ShortURL *service.ShortURL
}

func (h *Handler) UploadURL(context *gin.Context) {
	urls := context.PostForm("urls")
	fmt.Println(urls)

	if err := h.ShortURL.Upload(context, urls); err != nil {
		context.JSON(http.StatusNotFound, nil)
		return
	}

	//response := &service.ShortUrlResponse{
	//	Url_ID: ,
	//	//ShortUrl:
	//}

	context.JSON(http.StatusOK, nil)
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
