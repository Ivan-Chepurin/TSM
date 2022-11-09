package sqlite

import (
	"database/sql"
	"errors"
	"main/pkg/models"
)

type DBModel struct {
	DB *sql.DB
}

func (db *DBModel) InsertFarm(o models.Farm) (int, error) {
	stmt := `INSERT INTO farm (name) VALUES (?)`
	result, err := db.DB.Exec(stmt, o.Name)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (db *DBModel) GetFarm(farmId int) (models.Farm, error) {
	stmt := `SELECT * FROM farm WHERE id=?`
	row := db.DB.QueryRow(stmt, farmId)
	var farm models.Farm
	err := row.Scan(&farm.Id, &farm.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return farm, models.ErrNoRecord
		}
		return models.Farm{}, err
	}
	plants, err := db.GetAllFarmPlants(farmId)
	if err != nil {
		return models.Farm{}, err
	}
	farm.FarmPlants = models.NewFarmPlants(plants)
	return farm, nil
}

func (db *DBModel) GetAllFarms() ([]models.Farm, error) {
	stmt := `SELECT * FROM farm`
	rows, err := db.DB.Query(stmt)
	var farms []models.Farm
	if err != nil {
		return farms, err
	}
	for rows.Next() {
		farm := models.Farm{}
		err := rows.Scan(&farm.Id, &farm.Name)
		if err != nil {
			return []models.Farm{}, err
		}
		farms = append(farms, farm)
	}
	return farms, err
}

func (db *DBModel) InsertBrand(b models.Brand) (int, error) {
	stmt := `INSERT INTO brand (name) VALUES (?)`
	result, err := db.DB.Exec(stmt, b.Name)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (db *DBModel) GetBrand(brandId int) (models.Brand, error) {
	stmt := `SELECT id, name FROM brand WHERE id=?`
	row := db.DB.QueryRow(stmt, brandId)
	var brand models.Brand
	err := row.Scan(&brand.Id, &brand.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return brand, models.ErrNoRecord
		}
		return models.Brand{}, err
	}
	return brand, nil
}

func (db *DBModel) InsertPlant(o models.Plant, farmId int) (int, error) {
	stmt := `
		INSERT INTO plant (
		brand, name, farm_id, am, rrp, rrp_intensity, s1, s1_intensity, s2, 
		s2_intensity, m1, m1_intensity, m2, m2_intensity, m3, m3_intensity, dm, dm_intensity) 
		VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);
	`
	result, err := db.DB.Exec(stmt,
		o.BrandId, o.Name, o.FarmId, o.AM, o.RRP, o.RRPIntensity,
		o.S1, o.S1Intensity, o.S2, o.S2Intensity, o.M1, o.M1Intensity,
		o.M2, o.M2Intensity, o.M3, o.M3Intensity, o.DM, o.DMIntensity,
	)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (db *DBModel) GetPlant(plantId, farmId int) (models.Plant, error) {
	stmt := `SELECT * FROM plant WHERE id=? AND farm_id=?`
	row := db.DB.QueryRow(stmt, plantId, farmId)
	p := models.Plant{}
	err := row.Scan(AllPlantParams(&p)...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Plant{}, models.ErrNoRecord
		}
		return models.Plant{}, err
	}

	return p, nil
}

func (db *DBModel) GetFarmPlantsByBrand(brand int, farmId int) ([]models.Plant, error) {
	stmt := `SELECT * FROM plant WHERE farm_id=? AND brand=?`
	rows, err := db.DB.Query(stmt, farmId, brand)
	if err != nil {
		return []models.Plant{}, err
	}
	defer rows.Close()

	var plants []models.Plant
	for rows.Next() {
		p := models.Plant{}
		err := rows.Scan(AllPlantParams(&p)...)
		if err != nil {
			return []models.Plant{}, err
		}
		plants = append(plants, p)
	}
	return plants, nil
}

func AllPlantParams(p *models.Plant) []any {
	parameters := []any{
		&p.Id, &p.BrandId, &p.BrandName, &p.Name, &p.FarmId, &p.AM, &p.RRP, &p.RRPIntensity, &p.S1, &p.S1Intensity,
		&p.S2, &p.S2Intensity, &p.M1, &p.M1Intensity, &p.M2, &p.M2Intensity, &p.M3, &p.M3Intensity,
		&p.DM, &p.DMIntensity,
	}
	return parameters
}

func (db *DBModel) GetAllFarmPlants(farmId int) ([]models.Plant, error) {
	stmt := `
		SELECT plant.id, brand_id, brand.name, plant.name, farm_id, am, rrp, rrp_intensity,
			   s2, s2_intensity, s1, s1_intensity, m1, m1_intensity,
			   m2, m2_intensity, m3, m3_intensity, dm, dm_intensity
		FROM plant
				 JOIN brand ON plant.brand_id = brand.id
		WHERE farm_id = ?;
		`
	rows, err := db.DB.Query(stmt, farmId)
	var plants []models.Plant
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return plants, models.ErrNoRecord
		}
		return []models.Plant{}, err
	}
	defer rows.Close()

	for rows.Next() {
		p := models.Plant{}
		err := rows.Scan(AllPlantParams(&p)...)
		if err != nil {
			return []models.Plant{}, err
		}
		plants = append(plants, p)
	}
	return plants, nil
}
