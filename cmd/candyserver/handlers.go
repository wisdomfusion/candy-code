package main

import (
	"errors"
	"fmt"
	"github.com/wisdomfusion/candy-code-box/pkg/models"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.NotFound(w)
		return
	}

	candies, err := app.candy.Latest()
	if err != nil {
		app.ServerError(w, err)
		return
	}

	for _, candy := range candies {
		fmt.Fprintf(w, "%v", candy)
	}

	//files := []string{
	//	"./web/home.page.html",
	//	"./web/base.layout.html",
	//	"./web/footer.partial.html",
	//}
	//
	//ts, err := template.ParseFiles(files...)
	//if err != nil {
	//	log.Println(err.Error())
	//	app.ServerError(w, err)
	//	return
	//}
	//
	//err = ts.Execute(w, nil)
	//if err != nil {
	//	log.Println(err.Error())
	//	app.ServerError(w, err)
	//}

	//w.Write([]byte("hello from Candy Code Box."))
}

func (app *application) showCandy(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.NotFound(w)
		return
	}

	c, err := app.candy.Show(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.NotFound(w)
		} else {
			app.ServerError(w, err)
		}
		return
	}

	fmt.Fprintf(w, "%v", c)
}

func (app *application) storeCandy(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.ClientError(w, http.StatusMethodNotAllowed)
		return
	}

	title := "test"
	candy := "test code"
	expireDays := "7"

	id, err := app.candy.Store(title, candy, expireDays)
	if err != nil {
		app.ServerError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/candy?id=%d", id), http.StatusSeeOther)
}

