package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"member_role/internal/controller"
	"member_role/internal/controller/req"
	"member_role/internal/controller/res"
	"member_role/internal/page"
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
	r.HandleFunc("/member-roles/{cafeId:[0-9]+}", h.getList).Methods(http.MethodGet)
	r.HandleFunc("/member-roles/{cafeId:[0-9]+}/{memberId:[0-9]+}", h.getByMemberId).Methods(http.MethodGet)
	r.HandleFunc("/member-roles/{cafeId:[0-9]+}/{memberId:[0-9]+}", h.create).Methods(http.MethodPost)
	r.HandleFunc("/member-roles/{cafeId:[0-9]+}/{memberId:[0-9]+}/{id:[0-9]+}", h.upsert).Methods(http.MethodPut)
	r.HandleFunc("/member-roles/{cafeId:[0-9]+}/{memberId:[0-9]+}/{id:[0-9]+}", h.delete).Methods(http.MethodDelete)
	return r
}

const (
	InvalidCafeId       = "invalid cafe id"
	InvalidMemberId     = "invalid member id"
	invalidRoleId       = "invalid role id"
	InternalServerError = "internal server error"
)

func (h Handler) getByMemberId(w http.ResponseWriter, r *http.Request) {
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
		log.Println("getByMemberId json.Marshal err: ", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}

func (h Handler) getList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cafeId, err := strconv.Atoi(vars["cafeId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	reqPage := page.GetPageReqByRequest(r)

	list, total, err := h.c.GetList(r.Context(), cafeId, reqPage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(res.NewListTotalDto(list, total))
	if err != nil {
		log.Println("getList json.Marshal err: ", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}

func (h Handler) upsert(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, invalidRoleId, http.StatusBadRequest)
		return
	}
	cafeId, err := strconv.Atoi(vars["cafeId"])
	if err != nil {
		http.Error(w, InvalidCafeId, http.StatusBadRequest)
		return
	}
	memberId, err := strconv.Atoi(vars["memberId"])
	if err != nil {
		http.Error(w, InvalidMemberId, http.StatusBadRequest)
		return
	}
	var d req.PutDto
	err = json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.c.Upsert(r.Context(), id, cafeId, memberId, d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h Handler) delete(w http.ResponseWriter, r *http.Request) {
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
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "invalid member role id", http.StatusBadRequest)
		return
	}
	err = h.c.Delete(r.Context(), cafeId, memberId, id)
	if err != nil {
		if strings.Contains(err.Error(), "invalid") {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h Handler) create(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cafeId, err := strconv.Atoi(vars["cafeId"])
	if err != nil {
		http.Error(w, InvalidCafeId, http.StatusBadRequest)
		return
	}
	memberId, err := strconv.Atoi(vars["memberId"])
	if err != nil {
		http.Error(w, InvalidMemberId, http.StatusBadRequest)
		return
	}

	var cD req.CreateDto
	err = json.NewDecoder(r.Body).Decode(&cD)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.c.CreateRole(r.Context(), memberId, cafeId, cD)
	if err != nil {
		if strings.Contains(err.Error(), InternalServerError) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
