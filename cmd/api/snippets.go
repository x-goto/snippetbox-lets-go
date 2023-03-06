package main

import (
	"errors"
	"fmt"
	"goto/snippetbox/internal/data"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) homeHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		app.notFoundResponse(w, r)
		return
	}

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
		"./ui/html/header.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.logger.PrintError(err, nil)
		app.serverErrorResponse(w, r, err)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		app.logger.PrintError(err, nil)
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) createSnippetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.methodNotAllowedResponse(w, r)
		return
	}

	app.logger.PrintInfo(r.Header.Get("Authorization"), nil)

	title := "itt"
	content := "whatever happens happens"
	expires := "7"

	id, err := app.models.Snippets.Insert(title, content, expires)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
}

func (app *application) showSnippetHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id <= 0 {
		app.logger.PrintError(err, nil)
		app.badRequestResponse(w, r)
		return
	}

	snippet, err := app.models.Snippets.GetByID(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
			return
		default:
			app.serverErrorResponse(w, r, err)
			return
		}
	}

	fmt.Fprintf(w, "%v", snippet)
}
