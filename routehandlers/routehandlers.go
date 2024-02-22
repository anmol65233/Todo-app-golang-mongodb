package routehandlers

import (
	"anmol/todo/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func GetHomefunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Worlds!"))
}

func GetAllTask(w http.ResponseWriter, r *http.Request) {
	payload := models.GetAllTask()
	json.NewEncoder(w).Encode(payload)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.ToDoList
	json.NewDecoder(r.Body).Decode(&task)
	models.CreateTask(task)
	json.NewEncoder(w).Encode(task)
	fmt.Println(task)
}

func GetOne(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var customMessage = make(map[string]string)
	payload := models.GetById(id)
	length := len(payload)
	if length == 0 {
		customMessage["message"] = "no value found!"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(customMessage)
	} else {
		json.NewEncoder(w).Encode(payload)
	}
}

func TaskUpdateStatus(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	status, _ := strconv.ParseBool(chi.URLParam(r, "status"))
	var customMessage = make(map[string]string)
	payload := models.UpdateStatus(id, status)
	if payload == 0 {
		customMessage["message"] = "no value found!"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(customMessage)
	} else {
		w.WriteHeader(http.StatusOK)
		customMessage["message"] = "data modified!"
	}
	json.NewEncoder(w).Encode(customMessage)
}

func DeleteOne(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var customMessage = make(map[string]string)
	payload := models.DeleteOne(id)
	if payload == 0 {
		customMessage["message"] = "no value found!"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(customMessage)
	} else {
		w.WriteHeader(http.StatusOK)
		customMessage["message"] = "data deleted!"
	}
}
