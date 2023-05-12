package relation

import (
	"boleiro/domain/usecases/relation"
	"boleiro/view"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type newHttpRelationModule struct {
	useCases relation.UseCases
}

func NewHttpRelationModule(useCases relation.UseCases) view.HttpModule {
	return &newHttpRelationModule{
		useCases: useCases,
	}
}
func (n newHttpRelationModule) Setup(router *mux.Router) {
	router.HandleFunc("/players/{idPlayer}/sponsors/{idSponsor}", n.create).Methods("POST")
	router.HandleFunc("/players/{idPlayer}/sponsors", n.getById).Methods("GET")
	router.HandleFunc("/sponsors/{idSponsor}/players", n.getByIdS).Methods("GET")
	router.HandleFunc("/players/{idPlayer}/sponsors/{idSponsor}", n.delete).Methods("DELETE")
}
func (n newHttpRelationModule) create(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idPlayerString := vars["idPlayer"]
	idPlayer, err := strconv.ParseInt(idPlayerString, 10, 64)
	if err != nil {
		log.Println("[create] Error ParseInt", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	idSponsorString := vars["idSponsor"]
	idSponsor, err := strconv.ParseInt(idSponsorString, 10, 64)
	if err != nil {
		log.Println("[create] Error ParseInt", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = n.useCases.Create(r.Context(), idPlayer, idSponsor)
	if err != nil {
		log.Println("[Create] Error", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = w.Write([]byte("success"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println("[setCurrentPresentation] Error Write", err)
		return
	}

}
func (n newHttpRelationModule) delete(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idPlayerString := vars["idPlayer"]
	idPlayer, err := strconv.ParseInt(idPlayerString, 10, 64)
	if err != nil {
		log.Println("[delete] Error ParseInt", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	idSponsorString := vars["idSponsor"]
	idSponsor, err := strconv.ParseInt(idSponsorString, 10, 64)
	if err != nil {
		log.Println("[delete] Error ParseInt", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = n.useCases.Delete(r.Context(), idPlayer, idSponsor)
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
func (n newHttpRelationModule) getById(w http.ResponseWriter, r *http.Request) {

	idPlayer, err := strconv.ParseInt(mux.Vars(r)["idPlayer"], 10, 64)
	sponsors, err := n.useCases.GetById(r.Context(), idPlayer)
	if err != nil {
		log.Println("[getById] Error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(sponsors)
	if err != nil {
		log.Println("[getById] Error Marshal", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(b)
	if err != nil {
		log.Println("[getById] Error Write", err)
		return
	}
}
func (n newHttpRelationModule) getByIdS(w http.ResponseWriter, r *http.Request) {

	idSponsor, err := strconv.ParseInt(mux.Vars(r)["idSponsor"], 10, 64)
	players, err := n.useCases.GetByIdS(r.Context(), idSponsor)
	if err != nil {
		log.Println("[getByIdS] Error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(players)
	if err != nil {
		log.Println("[getByIdS] Error Marshal", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(b)
	if err != nil {
		log.Println("[getByIdS] Error Write", err)
		return
	}
}
