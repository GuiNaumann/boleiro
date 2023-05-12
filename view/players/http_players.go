package players

import (
	"boleiro/domain/entities"
	"boleiro/domain/usecases/players"
	"boleiro/view"
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("bytes", b)

	var player entities.Players
	if err = json.Unmarshal(b, &player); err != nil {
		log.Println("[create] Error json.Unmarshal", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println(player)

	err = n.useCases.Create(r.Context(), player)
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
func (n newHttpPlayerModule) update(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("[update] Error ReadAll", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var player entities.Players
	if err = json.Unmarshal(b, &player); err != nil {
		log.Println("[update] Error ReadAll", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	playersId, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		log.Println("[update] Error ParseInt")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = n.useCases.Update(r.Context(), player, playersId)
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
func (n newHttpPlayerModule) getAll(w http.ResponseWriter, r *http.Request) {

	orderBy := r.URL.Query().Get("orderBy")
	orderType := r.URL.Query().Get("orderType")

	filter := entities.ListFilter{OrderBy: orderBy, OrderType: orderType}

	playersList, err := n.useCases.GetAll(r.Context(), filter)
	if err != nil {
		log.Println("[getAll] Error GetAll", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(playersList)
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
func (n newHttpPlayerModule) delete(w http.ResponseWriter, r *http.Request) {

	playersId, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		log.Println("[delete] Error ParseInt")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = n.useCases.Delete(r.Context(), playersId)
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
func (n newHttpPlayerModule) GetById(w http.ResponseWriter, r *http.Request) {

	playersId, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	players, err := n.useCases.GetById(r.Context(), playersId)
	if err != nil {
		log.Println("[getById] Error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(players)
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
