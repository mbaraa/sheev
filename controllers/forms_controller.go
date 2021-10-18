package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/mbaraa/sheev/data"
	"github.com/mbaraa/sheev/models"
	"github.com/mbaraa/sheev/utils/formgen"
)

// FormsController well it's the forms controller :]
// ok fine it's responsible for the endpoint /forms/
type FormsController struct {
	endPoints  map[string]http.HandlerFunc
	formsStore data.FormsStore
}

// NewFormsController returns a new FormsController instance
func NewFormsController(formsData data.FormsStore) *FormsController {
	return (&FormsController{
		formsStore: formsData,
	}).initEndPoints()
}

// initEndPoints initializes all the children end points of /forms/
func (fc *FormsController) initEndPoints() *FormsController {
	fc.endPoints = map[string]http.HandlerFunc{
		"GET /all/":    fc.handleGetForms,
		"GET /single/": fc.handleGetForm,

		"POST /gen/": fc.handleGenerateForm,
	}
	return fc
}

// ServeHTTP the big fella, responds with a valid handler otherwise responses with a 404 status code
//
func (fc *FormsController) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res.Header().Set("Access-Control-Allow-Origin", "*")

	// DEBUG:
	fmt.Println(req.Method + " " + req.URL.Path)

	fullPath := strings.TrimPrefix(req.URL.Path, "/forms")
	path := fullPath
	if strings.Contains(fullPath, "/single") {
		path = strings.TrimSuffix(path, path[len("/single/"):])
	}
	if handler, validEndpoint := fc.endPoints[req.Method+" "+path]; validEndpoint {
		handler(res, req)
		return
	}
	if req.Method != "OPTIONS" {
		res.WriteHeader(404)
	}
}

// handleGetForms handles the request
// GET /forms/all
func (fc *FormsController) handleGetForms(res http.ResponseWriter, req *http.Request) {
	forms := fc.getForms()

	err := json.NewEncoder(res).Encode(map[string]interface{}{
		"forms": forms,
	})
	if err != nil {
		res.WriteHeader(500)
		return
	}
}

// handleGetForm handles the request
// GET /forms/single/{form_name}
func (fc *FormsController) handleGetForm(res http.ResponseWriter, req *http.Request) {
	formName := req.URL.Path[len("/forms/single/"):]
	form, err := fc.getForm(formName)

	if err != nil {
		res.WriteHeader(400)
		_ = json.NewEncoder(res).Encode(map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	err = json.NewEncoder(res).Encode(form)
	if err != nil {
		res.WriteHeader(500)
		return
	}
}

// handleGenerateForm handles the request
// POST /forms/gen/
//
// the request body is just a models.Form instance
//
func (fc *FormsController) handleGenerateForm(res http.ResponseWriter, req *http.Request) {
	form := new(models.Form)

	err := json.NewDecoder(req.Body).Decode(form)
	if err != nil {
		res.WriteHeader(400)
		return
	}

	img, err := fc.generateForm(form)
	if err != nil {
		res.WriteHeader(400)
		_ = json.NewEncoder(res).Encode(map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	err = json.NewEncoder(res).Encode(map[string]interface{}{
		"img": img,
	})
	if err != nil {
		res.WriteHeader(500)
		return
	}
}

// getForms returns available forms from the given forms store
func (fc *FormsController) getForms() []models.Form {
	forms, err := fc.formsStore.GetAll()
	if err != nil {
		return nil
	}

	return forms
}

// getForm returns a form of the given name from the given forms store
func (fc *FormsController) getForm(name string) (models.Form, error) {
	return fc.formsStore.Get(name)
}

// generateForm generates a form with the given data
func (fc *FormsController) generateForm(form *models.Form) (string, error) {
	form1, err := fc.getForm(form.Name)
	if err != nil {
		return "", err
	}

	formImage := formgen.NewFormImageFromB64Image(form1.B64FormImg)

	formGen := formgen.NewFormGenerator(
		form.Name,
		formImage,
		form.Fields,
	)

	finalForm, err := formGen.MakeForm()

	return finalForm.B64FormImg, err
}
