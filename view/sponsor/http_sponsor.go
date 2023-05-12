package sponsor

import (
	"boleiro/domain/entities"
	"boleiro/domain/usecases/sponsor"
	"boleiro/view"
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"strconv"
)

type newHttpSponsorModule struct {
	useCases sponsor.UseCases
}

func NewHttpSponsorModule(useCases sponsor.UseCases) view.HttpModule {
	return &newHttpSponsorModule{
		useCases: useCases,
	}
}

func (n newHttpSponsorModule) Setup(router *mux.Router) {
	router.HandleFunc("/sponsor", n.create).Methods("POST")
	router.HandleFunc("/sponsor/{id}", n.update).Methods("PUT")
	router.HandleFunc("/sponsor", n.getAll).Methods("GET")
	router.HandleFunc("/sponsor/{id}", n.GetById).Methods("GET")
	router.HandleFunc("/sponsor/{id}", n.delete).Methods("DELETE")
	log.Println("listening to /sponsor")
}
func (n newHttpSponsorModule) create(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("bytes", b)

	var sponsor entities.Sponsor
	if err = json.Unmarshal(b, &sponsor); err != nil {
		log.Println("[create] Error json.Unmarshal", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println(sponsor)

	err = n.useCases.Create(r.Context(), sponsor)
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
func (n newHttpSponsorModule) update(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("[update] Error ReadAll", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var sponsor entities.Sponsor
	if err = json.Unmarshal(b, &sponsor); err != nil {
		log.Println("[update] Error ReadAll", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sponsorId, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		log.Println("[update] Error ParseInt")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = n.useCases.Update(r.Context(), sponsor, sponsorId)
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
func (n newHttpSponsorModule) getAll(w http.ResponseWriter, r *http.Request) {

	sponsorList, err := n.useCases.GetAll(r.Context())
	if err != nil {
		log.Println("[getAll] Error GetAll", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, err := json.Marshal(sponsorList)
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
func (n newHttpSponsorModule) delete(w http.ResponseWriter, r *http.Request) {

	sponsorId, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		log.Println("[delete] Error ParseInt")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = n.useCases.Delete(r.Context(), sponsorId)
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
func (n newHttpSponsorModule) GetById(w http.ResponseWriter, r *http.Request) {

	sponsorId, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	sponsor, err := n.useCases.GetById(r.Context(), sponsorId)
	if err != nil {
		log.Println("[getById] Error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(sponsor)
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
