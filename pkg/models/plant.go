package models

type Plant struct {
	Id      string  `db:"id"  json:"id,omitempty"`
	BrandId int     `db:"brand_id"  json:"brand_id,omitempty"`
	Name    string  `db:"name" json:"name,omitempty"`
	FarmId  string  `db:"farm_id" json:"farm_id,omitempty"`
	AM      float64 `db:"am" json:"am,omitempty"` // AM - годовой пробег

	RRP float64 `db:"ppr" json:"rrp,omitempty"` // Routine repairs planned - Текущий ремонт плановый

	RRPIntensity float64 `db:"rrp_intensity" json:"rrp_intensity,omitempty"` // Трудоемкость текущего ремонта полного
	S1           float64 `db:"s1" json:"s1,omitempty"`                       // Service1 - ТО1

	S1Intensity float64 `db:"s1_intensity" json:"s1_intensity,omitempty"` // IntensityService1 - Трудоемкость ТО1
	S2          float64 `db:"s2" json:"s2,omitempty"`                     // Service2 - ТО2

	S2Intensity float64 `db:"s2_intensity" json:"s2_intensity,omitempty"` // IntensityService2 - ТО2
	M1          float64 `db:"m1" json:"m1,omitempty"`                     // Maintenance1 - Текущий ремонт 1

	M1Intensity float64 `db:"m1_intensity" json:"m1_intensity,omitempty"` // IntensityMaintenance1 - Текущий ремонт 1
	M2          float64 `db:"m2" json:"m2,omitempty"`                     // Maintenance2 - Текущий ремонт 2

	M2Intensity float64 `db:"m2_intensity" json:"m2_intensity,omitempty"` // IntensityMaintenance2 - Текущий ремонт 2
	M3          float64 `db:"m3" json:"m3,omitempty"`                     // Maintenance3 - Текущий ремонт 3

	M3Intensity float64 `db:"m3_intensity" json:"m3_intensity,omitempty"` // IntensityMaintenance3 - Текущий ремонт 3
	DM          float64 `db:"dm" json:"dm,omitempty"`                     // DailyMaintenance - Ежедневный текущий ремонт

	DMIntensity float64 `db:"dm_intensity" json:"dm_intensity,omitempty"` // IntensityDailyMaintenance - Ежедневный текущий ремонт

	BrandName string `json:"brand_named,omitempty"`
}
