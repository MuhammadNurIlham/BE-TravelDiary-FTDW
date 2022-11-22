package handlers

import (
	journeydto "be-journey/dto/journey"
	dto "be-journey/dto/result"
	"be-journey/models"
	"be-journey/repositories"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerJourney struct {
	JourneyRepository repositories.JourneyRepository
}

func HandlerJourney(JourneyRepository repositories.JourneyRepository) *handlerJourney {
	return &handlerJourney{JourneyRepository}
}

func (h *handlerJourney) FindJourneys(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	journeys, err := h.JourneyRepository.FindJourneys()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Create Embed Path File on Image property here ...
	for i, p := range journeys {
		journeys[i].Image = os.Getenv("PATH_FILE") + p.Image
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: journeys}
	json.NewEncoder(w).Encode(response)
}

// get journey by id journey
func (h *handlerJourney) GetJourney(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var journeys models.Journey
	journeys, err := h.JourneyRepository.GetJourney(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	// for i, p := range journey {
	// 	journey[i].Image = os.Getenv("PATH_FILE") + p.Image
	// }

	// delete this code if using cloudinary
	// journeys.Image = os.Getenv("PATH_FILE") + journeys.Image

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: journeys}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerJourney) CreateJourney(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// // get data user token
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	// // Get dataFile from midleware and store to filename variable here ...
	dataUpload := r.Context().Value("dataFile") // add this code
	filename := dataUpload.(string)             // add this code

	request := journeydto.JourneyRequest{
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
		Image:       path_file,
		UserID:      userId,
	}

	fmt.Println("ini dataku", filename, request)

	// request := new(journeydto.JourneyRequest)
	// if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	// 	json.NewEncoder(w).Encode(response)
	// 	return
	// }

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	//  ========== IF WANT TO DEPLOY TO HEROKU & CLOUDINARY, UNCOMMENT THIS CODE BELOW ============ \\
	// Declare Context Background, Cloud Name, API Key, API Secret
	var ctx = context.Background()

	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	//  ========== IF WANT TO DEPLOY TO HEROKU & CLOUDINARY, UNCOMMENT THIS CODE BELOW ============ \\
	// Add Cloudinary credentials
	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	//  ========== IF WANT TO DEPLOY TO HEROKU & CLOUDINARY, UNCOMMENT THIS CODE BELOW ============ \\
	// Upload file to Cloudinary
	resp, err := cld.Upload.Upload(ctx, filename, uploader.UploadParams{Folder: "thejourney"})

	CreatedAt := time.Now()

	journey := models.Journey{
		Title:       request.Title,
		Description: request.Description,
		// Image:       filename,
		Image:     resp.SecureURL,
		Books:     "false",
		CreatedAt: CreatedAt,
		UserID:    userId,
	}

	journey, err = h.JourneyRepository.CreateJourney(journey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	journey, _ = h.JourneyRepository.GetJourney(journey.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: journey}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerJourney) UpdateJourney(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	request := new(journeydto.JourneyUpdateRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	UpdatedAt := time.Now()

	journey, _ := h.JourneyRepository.GetJourney(int(id))

	journey.Title = request.Title
	journey.UserID = request.UserID
	journey.Image = request.Image
	journey.Description = request.Description
	journey.UpdatedAt = UpdatedAt

	if request.Title != "" {
		journey.Title = request.Title
	}

	if request.Description != "" {
		journey.Description = request.Description
	}

	// if filepath != "false" {
	// 	journey.Image = resp.SecureURL
	// }

	if request.Books != "" {
		journey.Books = request.Books
	}

	data, err := h.JourneyRepository.UpdateJourney(journey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseJourney(data)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerJourney) DeleteJourney(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	journey, err := h.JourneyRepository.GetJourney(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	deleteJourney, err := h.JourneyRepository.DeleteJourney(journey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: deleteJourney}
	json.NewEncoder(w).Encode(response)
}

func convertResponseJourney(u models.Journey) models.JourneyResponse {
	return models.JourneyResponse{
		ID:          u.ID,
		Title:       u.Title,
		Description: u.Description,
		Image:       u.Image,
		UserID:      u.UserID,
		User:        u.User,
		Books:       u.Books,
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
	}
}
