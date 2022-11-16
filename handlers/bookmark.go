package handlers

import (
	bookmarkdto "be-journey/dto/bookmark"
	dto "be-journey/dto/result"
	"be-journey/models"
	"be-journey/repositories"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerBookmark struct {
	BookmarkRepository repositories.BookmarkRepository
}

func HandlerBookmark(BookmarkRepository repositories.BookmarkRepository) *handlerBookmark {
	return &handlerBookmark{BookmarkRepository}
}

func (h *handlerBookmark) FindBookmarks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	bookmarks, err := h.BookmarkRepository.FindBookmarks()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: bookmarks}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerBookmark) GetBookmark(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var bookmarks models.Bookmark

	bookmarks, err := h.BookmarkRepository.GetBookmark(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: bookmarks}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerBookmark) CreateBookmark(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	var request bookmarkdto.BookmarkRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	CreatedAt := time.Now()

	bookmark := models.Bookmark{
		UserID:    userId,
		JourneyID: request.JourneyID,
		CreatedAt: CreatedAt,
	}

	data, err := h.BookmarkRepository.CreateBookmark(bookmark)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	bookmark, _ = h.BookmarkRepository.GetBookmark(data.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseBookmark(bookmark)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerBookmark) DeleteBookmark(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	bookmark, err := h.BookmarkRepository.GetBookmark(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.BookmarkRepository.DeleteBookmark(bookmark)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseBookmark(data)}
	json.NewEncoder(w).Encode(response)
}

func convertResponseBookmark(u models.Bookmark) models.Bookmark {
	return models.Bookmark{
		ID:        u.ID,
		UserID:    u.UserID,
		User:      u.User,
		JourneyID: u.JourneyID,
		Journey:   u.Journey,
		CreatedAt: u.CreatedAt,
	}
}

func (h *handlerBookmark) GetBookmarks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var bookmarks models.Bookmark

	bookmarks, err := h.BookmarkRepository.GetBookmarks(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: bookmarks}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerBookmark) DeleteBookmarks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	bookmark, err := h.BookmarkRepository.GetBookmarks(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.BookmarkRepository.DeleteBookmark(bookmark)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseBookmark(data)}
	json.NewEncoder(w).Encode(response)
}
