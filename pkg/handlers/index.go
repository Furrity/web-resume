package handlers

import "net/http"

func ResumeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/resume.html")
}
