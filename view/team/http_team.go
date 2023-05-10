package team

import (
	"boleiro/domain/entities"
	"boleiro/domain/usecases/team"
	"boleiro/view"
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"strconv"
)

type newHttpTeamModule struct {
	useCases team.UseCases
}

func NewHttpTeamModule(useCases team.UseCases) view.HttpModulew {
	return &newHttpTeamModule{
		useCases: useCases,
	}
}

func (n newHttpTeamModule) Setup(router *mux.Router) {
	router.HandleFunc("/team", n.create).Methods("POST")
	router.HandleFunc("/team/{id}", n.update).Methods("PUT")
	router.HandleFunc("/team", n.getAll).Methods("GET")
	router.HandleFunc("/team/{id}", n.delete).Methods("DELETE")
	log.Println("listening to /team")
}
func (n newHttpTeamModule) create(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("bytes", b)

	var team entities.Team
	if err = json.Unmarshal(b, &team); err != nil {
		log.Println("[create] Error json.Unmarshal", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println(team)

	err = n.useCases.Create(r.Context(), team)
	if err != nil {
		log.Println("[create] Error Create", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write([]byte("success"))
	if err != nil {
		log.Println("[create] Error Write", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
func (n newHttpTeamModule) update(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("[update] Error ReadAll", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var team entities.Team
	if err = json.Unmarshal(b, &team); err != nil {
		log.Println("[update] Error ReadAll", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	teamId, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		log.Println("[update] Error ParseInt")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = n.useCases.Update(r.Context(), team, teamId)
	if err != nil {
		log.Println("[Update] Error Update", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write([]byte("success"))
	if err != nil {
		log.Println("[Update] Error Write", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func (n newHttpTeamModule) getAll(w http.ResponseWriter, r *http.Request) {

	teamList, err := n.useCases.GetAll(r.Context())
	if err != nil {
		log.Println("[getAll] Error GetAll", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, err := json.Marshal(teamList)
	if err != nil {
		log.Println("[getAll] Error Marshal", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = w.Write(b)
	if err != nil {
		log.Println("[getAll] Error Write", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func (n newHttpTeamModule) delete(w http.ResponseWriter, r *http.Request) {

	teamId, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		log.Println("[delete] Error ParseInt")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = n.useCases.Delete(r.Context(), teamId)
	if err != nil {
		log.Println("[delete] Error Delete", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write([]byte("success"))
	if err != nil {
		log.Println("[delete] Error Write", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
