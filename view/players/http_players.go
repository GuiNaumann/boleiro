package players

import (
	"boleiro/domain/entities"
	"boleiro/domain/usecases/players"
	"boleiro/view"
	"boleiro/view/http_error"
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"strconv"
)

type newHttpPlayerModule struct {
	useCases players.UseCases
}

func NewHttpPlayerModule(useCases players.UseCases) view.HttpModule {
	return &newHttpPlayerModule{
		useCases: useCases,
	}
}

func (n newHttpPlayerModule) Setup(router *mux.Router) {
	router.HandleFunc("/players", n.create).Methods("POST")
	router.HandleFunc("/players/{id}", n.update).Methods("PUT")
	router.HandleFunc("/players", n.getAll).Methods("GET")
	router.HandleFunc("/players/{id}", n.GetById).Methods("GET")
	router.HandleFunc("/players/{id}", n.delete).Methods("DELETE")
}
func (n newHttpPlayerModule) create(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		http_error.HandleError(w, http_error.NewBadRequestError("Ocorreu um erro."))
		return
	}

	log.Println("bytes", b)

	var player entities.Players
	if err = json.Unmarshal(b, &player); err != nil {
		log.Println("[create] Error json.Unmarshal", err)
		http_error.HandleError(w, http_error.NewBadRequestError("jogador não é valido."))
		return
	}
	log.Println(player)

	err = n.useCases.Create(r.Context(), player)
	if err != nil {
		log.Println("[create] Error Create", err)
		http_error.HandleError(w, http_error.NewBadRequestError("ocorreu um erro inesperado."))
		return
	}

	_, err = w.Write([]byte("success"))
	if err != nil {
		log.Println("[create] Error Write", err)
		http_error.HandleError(w, http_error.NewInternalServerError("Ocorreu um erro inesperado"))
		return
	}

}
func (n newHttpPlayerModule) update(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("[update] Error ReadAll", err)
		http_error.HandleError(w, http_error.NewBadRequestError("jogador não é válido."))
		return
	}

	var player entities.Players
	if err = json.Unmarshal(b, &player); err != nil {
		log.Println("[update] Error ReadAll", err)
		http_error.HandleError(w, http_error.NewBadRequestError("jogador não é válido."))
		return
	}

	playersId, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		log.Println("[update] Error ParseInt")
		http_error.HandleError(w, http_error.NewBadRequestError("id do jogador é invalido."))
		return
	}

	err = n.useCases.Update(r.Context(), player, playersId)
	if err != nil {
		log.Println("[Update] Error Update", err)
		http_error.HandleError(w, http_error.NewBadRequestError("Ocorreu um erro ao atualizar."))
		return
	}

	_, err = w.Write([]byte("success"))
	if err != nil {
		log.Println("[Update] Error Write", err)
		http_error.HandleError(w, http_error.NewInternalServerError("Ocorreu um erro inesperado"))
		return
	}
}
func (n newHttpPlayerModule) getAll(w http.ResponseWriter, r *http.Request) {

	orderBy := r.URL.Query().Get("orderBy")
	orderType := r.URL.Query().Get("orderType")
	name := r.URL.Query().Get("name")

	filter := entities.ListFilter{OrderBy: orderBy, OrderType: orderType, Name: name}

	playersList, err := n.useCases.GetAll(r.Context(), filter)
	if err != nil {
		log.Println("[getAll] Error GetAll", err)
		http_error.HandleError(w, http_error.NewBadRequestError("Ocorreu um erro inesperado."))
		return
	}

	b, err := json.Marshal(playersList)
	if err != nil {
		log.Println("[getAll] Error Marshal", err)
		http_error.HandleError(w, http_error.NewInternalServerError("jogador é invalido"))
		return
	}
	_, err = w.Write(b)
	if err != nil {
		log.Println("[getAll] Error Write", err)
		http_error.HandleError(w, http_error.NewInternalServerError("Ocorreu um erro inesperado"))
		return
	}
}
func (n newHttpPlayerModule) delete(w http.ResponseWriter, r *http.Request) {

	playersId, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		log.Println("[delete] Error ParseInt")
		http_error.HandleError(w, http_error.NewBadRequestError("id do jogador é invalido."))
		return
	}

	err = n.useCases.Delete(r.Context(), playersId)
	if err != nil {
		log.Println("[delete] Error Delete", err)
		http_error.HandleError(w, http_error.NewBadRequestError("Ocorreu um erro ao deletar."))
		return
	}

	_, err = w.Write([]byte("success"))
	if err != nil {
		log.Println("[delete] Error Write", err)
		http_error.HandleError(w, http_error.NewInternalServerError("Ocorreu um erro inesperado"))
		return
	}
}
func (n newHttpPlayerModule) GetById(w http.ResponseWriter, r *http.Request) {

	playersId, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	players, err := n.useCases.GetById(r.Context(), playersId)
	if err != nil {
		log.Println("[getById] Error", err)
		http_error.HandleError(w, http_error.NewBadRequestError("id do jogador é invalido."))
		return
	}

	b, err := json.Marshal(players)
	if err != nil {
		log.Println("[getById] Error Marshal", err)
		http_error.HandleError(w, http_error.NewInternalServerError("patrocinador é invalido"))
		return
	}

	_, err = w.Write(b)
	if err != nil {
		log.Println("[getById] Error Write", err)
		http_error.HandleError(w, http_error.NewInternalServerError("Ocorreu um erro inesperado"))
		return
	}
}
