//Package to render views
package renderer

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type Renderer struct {
	RenderingEngine string
	RootPath        string
	Secure          bool
	Port            string
	ServerName      string
}

type TemplateData struct {
	IsAuthenticated bool
	IntMap          map[string]int
	StringMap       map[string]string
	FloatMap        map[string]float32
	Data            map[string]interface{}
	CSRFToken       string
	Port            string
	ServerName      string
	Secure          bool
}

func (renderer *Renderer) RenderPage(w http.ResponseWriter, r *http.Request, view string, variables interface{}, data interface{}) error {
	switch strings.ToLower(renderer.RenderingEngine) {
	case "go-templates":
		return renderer.RenderPageFromGoTemplate(w, r, view, data)
	case "jet":
	}
	return nil
}

//Renders html pages created as Go templates
func (renderer *Renderer) RenderPageFromGoTemplate(w http.ResponseWriter, r *http.Request, view string, data interface{}) error {
	template, err := template.ParseFiles(fmt.Sprintf("%s/views/%s.page.html", renderer.RootPath, view))

	if err != nil {
		return err
	}

	templateData := &TemplateData{}

	if data != nil {
		templateData = data.(*TemplateData)
	}

	err = template.Execute(w, &templateData)

	if err != nil {
		return err
	}

	return nil
}
