package handlers

import (
	"net/http"

	easygo "github.com/deepto98/easy-go"
)

//Type to use handlers from app
type Handlers struct {
	App *easygo.EasyGo
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	err := h.App.Renderer.RenderPage(w, r, "home", nil, nil)

	if err != nil {
		h.App.ErrorLog.Println("error render", err)
	}
}
