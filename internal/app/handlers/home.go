package handlers

import (
	"net/http"

	"github.com/mgm103/personal-portfolio-v2/internal/app/components"
)

func Home(w http.ResponseWriter, r *http.Request) {
	base := components.Base("Marcus Marinelli")
	base.Render(r.Context(), w)
}
