package htmx

import "net/http"

func SetHeaderForceSwapEntries(h http.Header, el string) {
	h.Set("HX-Force-Swap", "true")
	h.Set("HX-Reselect", el)
	h.Set("HX-Retarget", el)
}
