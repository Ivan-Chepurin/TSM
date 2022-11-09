package main

import (
	"errors"
	"fmt"
	"main/pkg/models"
	"net/http"
	"strconv"
)

var (
	unknownErr = errors.New("внутренняя ошибка сервера")
)

func (app *App) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
	}

	farms, err := app.DBModel.GetAllFarms()
	if err != nil {
		app.serverError(w, err)
	}
	for _, farm := range farms {
		w.Write([]byte(fmt.Sprintf("%v\n", farm)))
	}
}

func (app *App) showFarm(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	farm, err := app.DBModel.GetFarm(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
			return
		}
		app.serverError(w, err)
	}
	w.Write([]byte(fmt.Sprintf("%v, %v:\n", farm.Id, farm.Name)))
	for brandId, plants := range farm.FarmPlants.PlBrand {
		w.Write([]byte(fmt.Sprintf("\n")))
		w.Write([]byte(fmt.Sprintf("    %v", brandId)))
		brand, err := app.DBModel.GetBrand(brandId)
		if err != nil {
			if errors.Is(err, models.ErrNoRecord) {
				app.notFound(w)
				return
			}
			app.serverError(w, err)
		}
		w.Write([]byte(fmt.Sprintf(" - %v:\n", brand.Name)))
		for _, plant := range plants {
			w.Write([]byte(fmt.Sprintf("        %#v\n", plant)))
		}
		w.Write([]byte(fmt.Sprintf("\n")))
	}
}

func (app *App) showCalcByBrand(w http.ResponseWriter, r *http.Request) {
	brandId, err := strconv.Atoi(r.URL.Query().Get("brand_id"))
	farmId, err := strconv.Atoi(r.URL.Query().Get("farm_id"))
	//app.DBModel.GetFarm()
	if err != nil || brandId < 1 || farmId < 1 {
		app.notFound(w)
		return
	}
}
