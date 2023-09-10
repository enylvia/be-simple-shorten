package handler

import (
	"encoding/json"
	"fmt"
	"go-shorten-link/service"
	"go-shorten-link/utils"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type RedisShortenLinkHandler interface {
	CreateShortenLink(w http.ResponseWriter, r *http.Request)
	ResolveShortenLink(w http.ResponseWriter, r *http.Request)
	ListLink(w http.ResponseWriter, r *http.Request)
}

type RedisShortenLinkHandlerImplement struct {
	redisSvc service.ServiceRedis
}

func NewRedisShortenLinkHandler(redisSvc service.ServiceRedis) RedisShortenLinkHandler {
	return &RedisShortenLinkHandlerImplement{redisSvc: redisSvc}
}

func (h *RedisShortenLinkHandlerImplement) CreateShortenLink(w http.ResponseWriter, r *http.Request) {
	request, err := io.ReadAll(r.Body)
	if err != nil {
		utils.ReturnJSON(w, utils.InternalServerErrorResponse(err.Error()))
		return
	}
	defer r.Body.Close()

	var requestData struct {
		OriginalUrl string `json:"original_url"`
	}

	// marshal request to request data
	err = json.Unmarshal(request, &requestData)
	if err != nil {
		utils.ReturnJSON(w, utils.BadRequestResponse(err.Error()))
		return
	}

	shortLink, err := h.redisSvc.SetShortenLink(requestData.OriginalUrl)
	if err != nil {
		utils.ReturnJSON(w, utils.BadRequestResponse(err.Error()))
		return
	}

	protoc := "http"
	if r.TLS != nil {
		protoc = "https"
	}

	baseUrl := fmt.Sprintf("%s://%s/%s", protoc, r.Host, shortLink)

	var baseStruct struct {
		ShortUrl string `json:"short_url"`
	}

	baseStruct.ShortUrl = baseUrl

	utils.ReturnJSON(w, utils.CreatedResponse(baseStruct))

}
func (h *RedisShortenLinkHandlerImplement) ResolveShortenLink(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	getOriginalUrl, err := h.redisSvc.GetShortenLink(key)
	if err != nil {
		utils.ReturnJSON(w, utils.BadRequestResponse(err.Error()))
		return
	}

	http.Redirect(w, r, getOriginalUrl, http.StatusMovedPermanently)
}

func (h *RedisShortenLinkHandlerImplement) ListLink(w http.ResponseWriter, r *http.Request) {
	response, err := h.redisSvc.ListShortenLink()
	if err != nil {
		utils.ReturnJSON(w, utils.BadRequestResponse(err.Error()))
		return
	}
	utils.ReturnJSON(w, utils.SuccessResponse(response))
}
