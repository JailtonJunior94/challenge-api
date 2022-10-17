package repositories

import (
	"bytes"
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
	var b bytes.Buffer

	b.WriteString(`SELECT CAST(p.Id AS CHAR(36)) Id, p.Name, p.Climate, p.Terrain FROM Planets p`)
	if name != "" {
		b.WriteString(" WHERE p.Name = @name")
	}
	b.WriteString("	ORDER BY p.Name OFFSET @page ROWS FETCH NEXT @limit ROWS ONLY;")

	rows, err := r.DB.Query(b.String(), sql.Named("name", name), sql.Named("page", page), sql.Named("limit", limit))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var planets []entities.Planet
	for rows.Next() {
		var p entities.Planet
		if err := rows.Scan(&p.ID, &p.Name, &p.Climate, &p.Terrain); err != nil {
			return nil, err
		}
		planets = append(planets, p)
	}

	if len(planets) <= 0 {
		return nil, sql.ErrNoRows
	}

	return planets, nil
}

func (r *planetRepository) FindByID(id string) (*entities.Planet, error) {
	query := `SELECT
				CAST(p.Id AS CHAR(36)) Id,
				p.Name,
				p.Climate,
				p.Terrain,
				COALESCE(CAST(f.Id AS CHAR(36)), '') Id,
				COALESCE(CAST(f.PlanetId AS CHAR(36)), '') PlanetId,
				COALESCE(f.Title, '') Title,
				COALESCE(f.Director, '') Director,
				COALESCE(f.ReleaseDate, '') ReleaseDate
			FROM
				Planets p 
				LEFT JOIN Films f ON p.Id = f.PlanetId
			WHERE
				p.Id = @id`

	rows, err := r.DB.Query(query, sql.Named("id", id))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var p entities.Planet
	var f entities.Film
	var films = make(map[string][]entities.Film)

	for rows.Next() {
		if err := rows.Scan(&p.ID, &p.Name, &p.Climate, &p.Terrain, &f.ID, &f.PlanetID, &f.Title, &f.Director, &f.ReleaseDate); err != nil {
			return nil, err
		}

		item := entities.Film{ID: f.ID, PlanetID: f.PlanetID, Title: f.Title, Director: f.Director, ReleaseDate: f.ReleaseDate}
		if items, ok := films[p.ID]; ok {
			films[p.ID] = append(items, item)
		} else {
			films[p.ID] = []entities.Film{item}
		}
	}

	if p.ID == "" {
		return nil, sql.ErrNoRows
	}

	p.AddFilms(films[id])
	return &p, nil
}

func (r *planetRepository) AddPlanet(p *entities.Planet) error {
	query := `INSERT INTO Planets VALUES (@planetID, @name, @climate, @terrain);`

	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(
		sql.Named("planetID", p.ID),
		sql.Named("name", p.Name),
		sql.Named("climate", p.Climate),
		sql.Named("terrain", p.Terrain))

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if rows == 0 {
		return err
	}

	return nil
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
