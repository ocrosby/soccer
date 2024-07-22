package handlers

import (
	"encoding/json"
	"github.com/ocrosby/soccer/internal/tds-coaching-change-service/models"
	"github.com/ocrosby/soccer/pkg/common"
	"github.com/ocrosby/soccer/pkg/tds"
	"net/http"
)

type ChangesHandler struct {
	Url string
}

func NewChangesHandler() *ChangesHandler {
	return &ChangesHandler{
		Url: "https://www.topdrawersoccer.com/college-soccer-articles/tracking-division-i-coaching-changes_aid52742",
	}
}

// Read handles the GET /changes endpoint
// It contains an optional query parameter gender that can be used to filter the results
// Example: GET /changes, retrieves all changes
// Example: GET /changes?gender=male, retrieves only changes for male coaches
// Example: GET /changes?gender=female, retrieves only changes for female coaches

// swagger:route GET /changes changes listChanges
// Retrieves coaching changes.
// Optional query parameter "gender" can be used to filter the results by gender.
// responses:
//   200: changesResponse
//   400: badRequestResponse

// swagger:parameters listChanges
type changesParamsWrapper struct {
	// in:query
	// Gender to filter the coaching changes. Can be "male", "female", or omitted for all changes.
	Gender string `json:"gender"`
}

// swagger:response changesResponse
type changesResponseWrapper struct {
	// in:body
	Body []models.Change
}

// swagger:response badRequestResponse
type badRequestResponseWrapper struct {
	// in:body
	Body struct {
		Error string `json:"error"`
	}
}

func (h *ChangesHandler) Read(w http.ResponseWriter, r *http.Request) {
	var (
		err     error
		changes []models.Change
		gender  common.Gender
	)

	// Set the Content-Type header
	w.Header().Set("Content-Type", "application/json")

	// Parse the query parameters.
	queryParams := r.URL.Query()

	genderParam := queryParams.Get("gender") // Access the gender query parameter
	if gender, err = common.StringToGender(genderParam); err != nil {
		// An error happened while converting the specified gender.
		RespondToBadRequest(w, err)
		return // Stop further processing
	}

	if gender == common.Male {
		// Return Male Coaching Changes
		if changes, err = tds.GetMaleCoachingChanges(h.Url); err != nil {
			RespondToBadRequest(w, err)
			return // Stop further processing
		}
	} else if gender == common.Female {
		// Return Female Coaching Changes
		if changes, err = tds.GetFemaleCoachingChanges(h.Url); err != nil {
			RespondToBadRequest(w, err)
			return // Stop further processing
		}
	} else {
		// Return All Coaching Changes
		if changes, err = tds.GetAllCoachingChanges(h.Url); err != nil {
			RespondToBadRequest(w, err)
			return // Stop further processing
		}
	}

	// Send the response
	_ = json.NewEncoder(w).Encode(changes)

	return
}

func RespondToBadRequest(w http.ResponseWriter, err error) {
	// Set the Content-Type header
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusBadRequest) // Set HTTP status code to 400
	errMsg := map[string]string{
		"error": err.Error(),
	}

	_ = json.NewEncoder(w).Encode(errMsg)
}
