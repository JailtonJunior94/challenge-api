package repositories

import (
	"database/sql"
	"testing"

	"github.com/jailtonjunior94/challenge/internal/domain/entities"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func initDB() *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}

	sqlStatement := `
		CREATE TABLE IF NOT EXISTS Planets 
		(
			Id TEXT NOT NULL PRIMARY KEY,
			Name TEXT,
			Climate TEXT,
			Terrain TEXT
		);

		CREATE TABLE IF NOT EXISTS Films 
		(
			Id TEXT NOT NULL PRIMARY KEY,
			PlanetId TEXT,
			Title TEXT,
			Director TEXT,
			ReleaseDate TEXT
		);
	`
	_, err = db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
	return db
}

func downDB(db *sql.DB) {
	db.Exec("DROP TABLE Planets")
	db.Exec("DROP TABLE Films")

	db.Close()
}

func TestAddPlanet(t *testing.T) {
	db := initDB()
	defer downDB(db)

	planetRepository := NewPlanetRepository(db)
	planet, _ := entities.NewPlanet("Tatooine", "arid", "desert")

	err := planetRepository.AddPlanet(planet)
	assert.Nil(t, err)
}

func TestAddFilm(t *testing.T) {
	db := initDB()
	defer downDB(db)

	planetRepository := NewPlanetRepository(db)
	planet, _ := entities.NewPlanet("Tatooine", "arid", "desert")
	film, _ := entities.NewFilm(planet.ID.String(), "A New Hope", "George Lucas", "1977-05-25")

	err := planetRepository.AddFilm(film)
	assert.Nil(t, err)
}

func TestFindByID(t *testing.T) {
	db := initDB()
	defer downDB(db)

	planetRepository := NewPlanetRepository(db)
	planet, _ := entities.NewPlanet("Tatooine", "arid", "desert")
	err := planetRepository.AddPlanet(planet)
	assert.Nil(t, err)

	filmOne, _ := entities.NewFilm(planet.ID.String(), "A New Hope", "George Lucas", "1977-05-25")
	filmTwo, _ := entities.NewFilm(planet.ID.String(), "Return of the Jedi", "Richard Marquand", "1983-05-25")

	var films []entities.Film
	films = append(films, *filmOne)
	films = append(films, *filmTwo)

	for _, f := range films {
		err := planetRepository.AddFilm(&f)
		assert.Nil(t, err)
	}

	p, err := planetRepository.FindByID(planet.ID.String())
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Equal(t, "Tatooine", p.Name)
	assert.Len(t, p.Films, 2)
}

func TestRemovePlanet(t *testing.T) {
	db := initDB()
	defer downDB(db)

	planetRepository := NewPlanetRepository(db)
	planet, _ := entities.NewPlanet("Tatooine", "arid", "desert")
	err := planetRepository.AddPlanet(planet)
	assert.Nil(t, err)

	filmOne, _ := entities.NewFilm(planet.ID.String(), "A New Hope", "George Lucas", "1977-05-25")
	filmTwo, _ := entities.NewFilm(planet.ID.String(), "Return of the Jedi", "Richard Marquand", "1983-05-25")

	var films []entities.Film
	films = append(films, *filmOne)
	films = append(films, *filmTwo)

	for _, f := range films {
		err := planetRepository.AddFilm(&f)
		assert.Nil(t, err)
	}

	err = planetRepository.Remove(planet.ID.String())
	assert.Nil(t, err)
}
