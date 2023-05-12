package infrastructure

import (
	players_usecase "boleiro/domain/usecases/players"
	relation_usecase "boleiro/domain/usecases/relation"
	sponsor_usecase "boleiro/domain/usecases/sponsor"
	team_usecase "boleiro/domain/usecases/team"
	player_repository "boleiro/infrastructure/repositories/players"
	relation_repository "boleiro/infrastructure/repositories/relation"
	sponsor_repository "boleiro/infrastructure/repositories/sponsor"
	team_repository "boleiro/infrastructure/repositories/team"
	"boleiro/settings"
	"boleiro/view/players"
	"boleiro/view/relation"
	"boleiro/view/sponsor"
	"boleiro/view/team"
	"database/sql"
	"github.com/MadAppGang/httplog"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Setup(settings settings.Settings, router *mux.Router) error {
	db, err := setupDataBase(settings.DataBase)
	if err != nil {
		log.Println("[Setup] Error setupDataBase", err)
		return err
	}

	setupModules(db, router)

	return nil
}

func setupDataBase(databaseSettings settings.DataBase) (*sql.DB, error) {
	db, err := sql.Open("mysql", databaseSettings.GetDBSource())
	if err != nil {
		log.Println("[setupDataBase] Error Open", err)
		return nil, err
	}

	return db, nil
}

func setupModules(db *sql.DB, router *mux.Router) {
	router.Use(ContentTypeMiddleware)
	router.Use(httplog.Logger)

	playerRepository := player_repository.NewRepository(db)
	playersUseCases := players_usecase.NewUseCases(playerRepository)
	players.NewHttpPlayerModule(playersUseCases).Setup(router)

	teamRepository := team_repository.NewRepository(db)
	teamUseCases := team_usecase.NewUseCases(teamRepository)
	team.NewHttpTeamModule(teamUseCases).Setup(router)

	sponsorRepository := sponsor_repository.NewRepository(db)
	sponsorUseCases := sponsor_usecase.NewUseCases(sponsorRepository)
	sponsor.NewHttpSponsorModule(sponsorUseCases).Setup(router)

	relationRepository := relation_repository.NewRepository(db)
	relationUseCases := relation_usecase.NewUseCases(relationRepository)
	relation.NewHttpRelationModule(relationUseCases).Setup(router)
}

// retornar json para todas as funções
func ContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
