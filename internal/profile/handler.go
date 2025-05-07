package profile

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/Furrity/web-resume/pkg/app"
	"github.com/Furrity/web-resume/pkg/l10n"
	"github.com/Furrity/web-resume/pkg/utils"
	"github.com/go-chi/chi/v5"
)

type ProfileReturn struct {
	Name  string  `json:"name"`
	Title string  `json:"title"`
	About string  `json:"about"`
	Image *string `json:"image,omitempty"`
}

func GetProfileHandler(w http.ResponseWriter, r *http.Request, a *app.App) {
	// everyone can see Profile data
	langCode := chi.URLParam(r, "lang")
	if !l10n.IsSupported(langCode) {
		utils.RespondWithError(
			w,
			http.StatusBadRequest,
			fmt.Sprintf("unsupported language code: %s", langCode),
			nil,
		)
		return
	}

	data, err := a.Queries.GetProfileInfo(r.Context(), langCode)
	if errors.Is(err, sql.ErrNoRows) {
		utils.RespondWithError(
			w,
			http.StatusNotFound,
			"no profile found for that language",
			nil,
		)
		return
	}
	var image *string
	if data.ImageUrl.Valid {
		image = &data.ImageUrl.String
	}

	utils.RespondWithJson(
		w,
		http.StatusOK,
		ProfileReturn{
			Name:  data.Name,
			Title: data.Title,
			About: data.About,
			Image: image,
		},
	)

}
