package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"

	"github.com/alx-t/go-playground/common"
	"github.com/alx-t/go-playground/data/project"
)

func GetProjects(w http.ResponseWriter, r *http.Request) {
	ps, err := project.Read(common.Db, "")
	j, err := json.Marshal(ProjectsResource{Data: ps})

	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func GetProjectById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		common.DisplayAppError(w, err, "Error id converting", 500)
		return
	}

	p, err := project.ReadOne(common.Db, id)
	j, err := json.Marshal(ProjectResource{Data: p})

	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func CreateProject(w http.ResponseWriter, r *http.Request) {
	var dataResource ProjectResource

	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid Project data", 500)
		return
	}

	prj := &dataResource.Data

	if _, err := project.Insert(common.Db, prj); err != nil {
		common.DisplayAppError(w, err, "Invalid Project data", 500)
		return
	}

	j, err := json.Marshal(ProjectResource{Data: *prj})
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

func UpdateProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		common.DisplayAppError(w, err, "Error id converting", 500)
		return
	}

	var dataResource ProjectResource
	err = json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid Project data", 500)
		return
	}

	prj := &dataResource.Data

	fmt.Println(id)

	if _, err := project.Update(common.Db, id, prj); err != nil {
		common.DisplayAppError(w, err, "Invalid Project data", 500)
		return
	}

	j, err := json.Marshal(ProjectResource{Data: *prj})
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func DeleteProject(w http.ResponseWriter, r *http.Request) {
	j, _ := json.Marshal(map[string]string{"message": "Delete Project"})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	w.Write(j)
}
