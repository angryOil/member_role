package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"member_role/internal/controller"
	"member_role/internal/controller/req"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
	c controller.Controller
}

func NewHandler(c controller.Controller) http.Handler {
	r := mux.NewRouter()
	h := Handler{c: c}
	r.HandleFunc("/member-roles/{cafeId:[0-9]+}/{memberId:[0-9]+}", h.getList).Methods(http.MethodGet)
	r.HandleFunc("/member-roles/{cafeId:[0-9]+}/{memberId:[0-9]+}", h.create).Methods(http.MethodPost)
	//r.HandleFunc("/member-roles/{cafeId:[0-9]+}/{memberId:[0-9]+}/{id:[0-9]+}", h.patch).Methods(http.MethodPatch)
	//r.HandleFunc("/member-roles/{cafeId:[0-9]+}/{memberId:[0-9]+}/{id:[0-9]+}", h.delete).Methods(http.MethodDelete)
	return r
}

func (h Handler) create(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cafeId, err := strconv.Atoi(vars["cafeId"])
	if err != nil {
		http.Error(w, "invalid cafe id", http.StatusBadRequest)
		return
	}
	memberId, err := strconv.Atoi(vars["memberId"])
	if err != nil {
		http.Error(w, "invalid member id", http.StatusBadRequest)
		return
	}

	var d req.CreateDto
	err = json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.c.Create(r.Context(), cafeId, memberId, d)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h Handler) getList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cafeId, err := strconv.Atoi(vars["cafeId"])
	if err != nil {
		http.Error(w, "invalid cafe id", http.StatusBadRequest)
		return
	}
	memberId, err := strconv.Atoi(vars["memberId"])
	if err != nil {
		http.Error(w, "invalid member id", http.StatusBadRequest)
		return
	}

	list, err := h.c.GetListByMemberId(r.Context(), cafeId, memberId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(list)
	if err != nil {
		log.Println("getList json.Marshal err: ", err)
		http.Error(w, "interanl server error", http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}