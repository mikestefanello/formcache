package handlers

import (
	"net/http"

	"github.com/mikestefanello/formcache/cache"
)

type StepFormContent struct {
	Token string
	Code  string
	Name  string
}

func (h *HTTPHandler) Steps(w http.ResponseWriter, r *http.Request) {
	c, err := h.cache.GetFormCache(r)
	if err != nil {
		c = cache.NewCacheData()
		c.Data["step"] = 1
	}

	switch c.Data["step"] {
	case 1:
		h.stepOne(w, r, c)
	case 2:
		h.stepTwo(w, r, c, true)
	case 3:
		h.stepThree(w, r, c, true)
	default:
		h.stepOne(w, r, c)
	}
}

func (h *HTTPHandler) stepOne(w http.ResponseWriter, r *http.Request, c cache.CacheData) {
	page := Page{Title: "Step one"}

	if r.Method == http.MethodPost {
		v := validateStepOne(&page, r)
		if v {
			c.Data["step"] = 2
			c.Data["code"] = r.FormValue("code")
			h.cache.SetData(&c)
			h.stepTwo(w, r, c, false)
			return
		}
	}

	h.Render(w, "step-one", page)
}

func (h *HTTPHandler) stepTwo(w http.ResponseWriter, r *http.Request, c cache.CacheData, submitted bool) {
	page := Page{
		Title: "Step two",
		Content: StepFormContent{
			Code:  c.Data["code"].(string),
			Token: c.ID,
		},
	}

	if submitted {
		if r.FormValue("name") == "" {
			page.AddMessage("danger", "You forgot to add your name!")
		} else {
			c.Data["step"] = 3
			c.Data["name"] = r.FormValue("name")
			h.cache.SetData(&c)
			h.stepThree(w, r, c, false)
			return
		}
	}

	h.Render(w, "step-two", page)
}

func (h *HTTPHandler) stepThree(w http.ResponseWriter, r *http.Request, c cache.CacheData, submitted bool) {
	page := Page{
		Content: StepFormContent{
			Code:  c.Data["code"].(string),
			Name:  c.Data["name"].(string),
			Token: c.ID,
		},
	}

	if submitted {
		page.AddMessage("success", "Operation successful!")
		h.cache.DeleteData(c.ID)
		h.Render(w, "finished", page)
		return
	}

	page.Title = "Step three"
	h.Render(w, "step-three", page)
}

func validateStepOne(p *Page, r *http.Request) bool {
	valid := true
	if r.FormValue("invalidate") == "on" {
		p.AddMessage("danger", "You invalidated the form!")
		valid = false
	}
	if r.FormValue("code") == "" {
		p.AddMessage("danger", "Code is required")
		valid = false
	}
	return valid
}
