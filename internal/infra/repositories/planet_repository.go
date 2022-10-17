package repositories

import (
	"database/sql"

	"github.com/jailtonjunior94/challenge/internal/domain/entities"
)

type planetRepository struct {
	DB *sql.DB
}

func NewPlanetRepository(db *sql.DB) *planetRepository {
	return &planetRepository{DB: db}
}

func (r *planetRepository) FindAll(name string, page int, limit int) ([]entities.Planet, error) {
	return nil, nil
}

func (r *planetRepository) AddPlanet(p *entities.Planet) (*entities.Planet, error) {
	query := `INSERT INTO Planets VALUES (@planetID, @name, @climate, @terrain);`

	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(
		sql.Named("planetID", p.ID),
		sql.Named("name", p.Name),
		sql.Named("climate", p.Climate),
		sql.Named("terrain", p.Terrain))

	if err != nil {
		return nil, err
	}

	rows, err := result.RowsAffected()
	if rows == 0 {
		return nil, err
	}

	return p, nil
}

func (r *planetRepository) AddFilm(f *entities.Film) error {
	query := `INSERT INTO Films VALUES (@filmID, @planetID, @title, @director, @releaseDate);`

	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(
		sql.Named("filmID", f.ID),
		sql.Named("planetID", f.PlanetID),
		sql.Named("title", f.Title),
		sql.Named("director", f.Director),
		sql.Named("releaseDate", f.ReleaseDate))

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if rows == 0 {
		return err
	}

	return nil
}

func (r *planetRepository) Remove(id string) error {
	planet, err := r.FindByID(id)
	if err != nil {
		return err
	}

	query := `DELETE FROM Planets WHERE Id = @id`
	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(sql.Named("id", planet.ID))
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if rows == 0 {
		return err
	}

	return nil
}

func (r *planetRepository) FindByID(id string) (*entities.Planet, error) {
	query := `SELECT
				CAST(p.Id AS CHAR(36)) Id,
				p.Name,
				p.Climate,
				p.Terrain,
				CAST(f.Id AS CHAR(36)) Id,
				f.PlanetId,
				f.Title,
				f.Director,
				f.ReleaseDate
			FROM
				Planets p 
				LEFT JOIN Films f ON p.Id = f.PlanetId
			WHERE
				p.Id = @id`

	rows, err := r.DB.Query(query, sql.Named("id", id))
	if err != nil {
		return nil, err
	}

	var p entities.Planet
	var f entities.Film
	var films = make(map[string][]entities.Film)

	for rows.Next() {
		if err := rows.Scan(&p.ID, &p.Name, &p.Climate, &p.Terrain, &f.ID, &f.PlanetID, &f.Title, &f.Director, &f.ReleaseDate); err != nil {
			return nil, err
		}

		item := entities.Film{ID: f.ID, PlanetID: f.PlanetID, Title: f.Title, Director: f.Director, ReleaseDate: f.ReleaseDate}
		if items, ok := films[p.ID.String()]; ok {
			films[p.ID.String()] = append(items, item)
		} else {
			films[p.ID.String()] = []entities.Film{item}
		}
	}

	p.AddFilms(films[id])
	return &p, nil
}
