package team

import (
	"boleiro/domain/entities"
	team_usecase "boleiro/domain/usecases/team"
	"boleiro/view"
	"boleiro/view/http_error"
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"strconv"
)

type newHttpTeamModule struct {
	useCases team_usecase.UseCases
}

func NewHttpTeamModule(useCases team_usecase.UseCases) view.HttpModule {
	return &newHttpTeamModule{
		useCases: useCases,
	}
}

func (n newHttpTeamModule) Setup(router *mux.Router) {
	router.HandleFunc("/teams", n.create).Methods("POST")
	router.HandleFunc("/teams/{id}", n.update).Methods("PUT")
	router.HandleFunc("/teams", n.getAll).Methods("GET")
	router.HandleFunc("/teams/{id}", n.getById).Methods("GET")
	router.HandleFunc("/teams/{id}", n.delete).Methods("DELETE")
}
func (n newHttpTeamModule) create(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("[create] Error ReadAll", err)
		http_error.HandleError(w, http_error.NewInternalServerError(http_error.UnexpectedError))
		return
	}

	var team entities.Team
	if err = json.Unmarshal(b, &team); err != nil {
		log.Println("[create] Error json.Unmarshal", err)
		http_error.HandleError(w, http_error.NewBadRequestError("Time não é válido."))
		return
	}

	err = n.useCases.Create(r.Context(), team)
	if err != nil {
		log.Println("[create] Error Create", err)
		http_error.HandleError(w, err)
		return
	}

	_, err = w.Write([]byte("success"))
	if err != nil {
		log.Println("[create] Error Write", err)
		http_error.HandleError(w, http_error.NewInternalServerError(http_error.UnexpectedError))
		return
	}

}
func (n newHttpTeamModule) update(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("[update] Error ReadAll", err)
		http_error.HandleError(w, http_error.NewBadRequestError("Time não é válido."))
		return
	}

	var team entities.Team
	if err = json.Unmarshal(b, &team); err != nil {
		log.Println("[update] Error ReadAll", err)
		http_error.HandleError(w, http_error.NewBadRequestError("time não é válido."))
		return
	}

	teamId, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		log.Println("[update] Error ParseInt", err)
		http_error.HandleError(w, http_error.NewBadRequestError("id do time é invalido."))
		return
	}

	err = n.useCases.Update(r.Context(), team, teamId)
	if err != nil {
		log.Println("[Update] Error Update", err)
		http_error.HandleError(w, err)
		return
	}

	_, err = w.Write([]byte("success"))
	if err != nil {
		log.Println("[Update] Error Write", err)
		http_error.HandleError(w, http_error.NewInternalServerError(http_error.UnexpectedError))
		return
	}
}
func (n newHttpTeamModule) getAll(w http.ResponseWriter, r *http.Request) {

	teamList, err := n.useCases.GetAll(r.Context())
	if err != nil {
		log.Println("[getAll] Error GetAll", err)
		http_error.HandleError(w, err)
		return
	}
	b, err := json.Marshal(teamList)
	if err != nil {
		log.Println("[getAll] Error Marshal", err)
		http_error.HandleError(w, http_error.NewInternalServerError(http_error.UnexpectedError))
		return
	}
	_, err = w.Write(b)
	if err != nil {
		log.Println("[getAll] Error Write", err)
		http_error.HandleError(w, http_error.NewInternalServerError(http_error.UnexpectedError))
		return
	}
}
func (n newHttpTeamModule) delete(w http.ResponseWriter, r *http.Request) {

	teamId, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		log.Println("[delete] Error ParseInt", err)
		http_error.HandleError(w, http_error.NewBadRequestError("id do time é invalido."))
		return
	}

	err = n.useCases.Delete(r.Context(), teamId)
	if err != nil {
		log.Println("[delete] Error Delete", err)
		http_error.HandleError(w, err)
		return
	}

	_, err = w.Write([]byte("success"))
	if err != nil {
		log.Println("[delete] Error Write", err)
		http_error.HandleError(w, http_error.NewInternalServerError(http_error.UnexpectedError))
		return
	}
}
func (n newHttpTeamModule) getById(w http.ResponseWriter, r *http.Request) {

	teamId, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		log.Println("[getById] Error ParseInt", err)
		http_error.HandleError(w, http_error.NewBadRequestError("Id do time não invalido."))
	}

	players, err := n.useCases.GetById(r.Context(), teamId)
	if err != nil {
		log.Println("[getById] Error GetById", err)
		http_error.HandleError(w, err)
		return
	}

	b, err := json.Marshal(players)
	if err != nil {
		log.Println("[getById] Error Marshal", err)
		http_error.HandleError(w, http_error.NewInternalServerError(http_error.UnexpectedError))
		return
	}

	_, err = w.Write(b)
	if err != nil {
		log.Println("[getById] Error Write", err)
		http_error.HandleError(w, http_error.NewInternalServerError(http_error.UnexpectedError))
		return
	}
}
