package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: подходящей записи не найдено")

var FarmsSchema = `
CREATE TABLE IF NOT EXISTS farm
(
    id   INTEGER UNIQUE NOT NULL PRIMARY KEY AUTOINCREMENT,
    name TEXT(100)
);

CREATE TABLE IF NOT EXISTS brand
(
    id   INTEGER UNIQUE NOT NULL PRIMARY KEY AUTOINCREMENT,
    name TEXT(100)
);

	CREATE TABLE IF NOT EXISTS plant
(
    id            INTEGER UNIQUE NOT NULL PRIMARY KEY AUTOINCREMENT,
    brand_id      INT            NOT NULL DEFAULT (1),
    name          text(100)      NOT NULL UNIQUE,
    farm_id       int            not null,
    am            REAL           NOT NULL,
    rrp           REAL           NOT NULL,
    rrp_intensity REAL           NOT NULL,
    s1            REAL           NOT NULL,
    s1_intensity  REAL           NOT NULL,
    s2            REAL           NOT NULL,
    s2_intensity  REAL           NOT NULL,
    m1            REAL           NOT NULL,
    m1_intensity  REAL           NOT NULL,
    m2            REAL           NOT NULL,
    m2_intensity  REAL           NOT NULL,
    m3            REAL           NOT NULL,
    m3_intensity  REAL           NOT NULL,
    dm            REAL           NOT NULL,
    dm_intensity  REAL           NOT NULL,
    FOREIGN KEY (farm_id) REFERENCES farm (id) ON DELETE RESTRICT,
    FOREIGN KEY (brand_id) REFERENCES brand (id) ON DELETE RESTRICT
);
`

type Farm struct {
	Id   int    `db:"id" json:"id,omitempty"`
	Name string `db:"name" json:"name,omitempty"`

	FarmPlants
}
type FarmPlants struct {
	PlBrand map[int][]Plant
}

func (fp *FarmPlants) SliceByBrand(brandId int) []Plant {
	return fp.PlBrand[brandId]
}

func NewFarmPlants(plants []Plant) FarmPlants {
	fp := FarmPlants{PlBrand: make(map[int][]Plant)}

	for _, plant := range plants {
		fp.PlBrand[plant.BrandId] = append(fp.PlBrand[plant.BrandId], plant)
	}
	return fp
}

// Расчеты количества ТО и ремонтов

func (f *Farm) RRPCount(brand int) float64 {
	p := f.SliceByBrand(brand)
	return p[0].AM * float64(len(p)) / p[0].RRP
}

func (f *Farm) S2Count(brand int) float64 {
	p := f.SliceByBrand(brand)
	return p[0].AM*float64(len(p))/p[0].S2 - f.RRPCount(brand)
}

func (f *Farm) S1Count(brand int) float64 {
	p := f.SliceByBrand(brand)
	return float64(len(p))*p[0].AM/p[0].S1 - f.RRPCount(brand) - f.S2Count(brand)
}

func (f *Farm) M1Count(brand int) float64 {
	p := f.SliceByBrand(brand)
	return float64(len(p)) * p[0].AM / p[0].M1
}

func (f *Farm) M2Count(brand int) float64 {
	p := f.SliceByBrand(brand)
	return float64(len(p)) * p[0].AM / p[0].M2
}

func (f *Farm) M3Count(brand int) float64 {
	p := f.SliceByBrand(brand)
	return float64(len(p)) * p[0].AM / p[0].M3
}

func date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func (f *Farm) DMCount() float64 {
	t1 := date(time.Now().Year(), 1, 1)
	t2 := date(time.Now().Year()+1, 1, 1)
	days := t2.Sub(t1).Hours() / 24
	return days
}

// Расчеты годовой трудоемкости ТО и ремонтов

func (f *Farm) RRPAnnualLaborIntensity(brand int) float64 {
	p := f.SliceByBrand(brand)
	return f.RRPCount(brand) * p[0].RRPIntensity
}

func (f *Farm) S2AnnualLaborIntensity(brand int) float64 {
	p := f.SliceByBrand(brand)
	return f.S2Count(brand) * p[0].S2Intensity
}

func (f *Farm) S1AnnualLaborIntensity(brand int) float64 {
	p := f.SliceByBrand(brand)
	return f.S1Count(brand) * p[0].S1Intensity
}

func (f *Farm) M1AnnualLaborIntensity(brand int) float64 {
	p := f.SliceByBrand(brand)
	return f.M1Count(brand) * p[0].M1Intensity
}

func (f *Farm) M2AnnualLaborIntensity(brand int) float64 {
	p := f.SliceByBrand(brand)
	return f.M1Count(brand) * p[0].M2Intensity
}

func (f *Farm) M3AnnualLaborIntensity(brand int) float64 {
	p := f.SliceByBrand(brand)
	return f.M1Count(brand) * p[0].M3Intensity
}

func (f *Farm) DMAnnualLaborIntensity(brand int) float64 {
	p := f.SliceByBrand(brand)
	return f.DMCount() * p[0].DMIntensity
}
