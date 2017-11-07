package JSONResponse

import (
	"fmt"
	"net/http"
)

func ResponseWhithData(w http.ResponseWriter, json []byte, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
<<<<<<< HEAD
<<<<<<< HEAD
	fmt.Fprintf(w, "{ %q: \"%v\"}", "data", json)
	//w.Write(json)
=======
	w.Write(json)
>>>>>>> c039ae87024e99b2581efc1c8627aa042410cded
=======
	w.Write(json)
>>>>>>> c039ae87024e99b2581efc1c8627aa042410cded
}

func ResponseWhithMessage(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	fmt.Fprintf(w, "{ %q: %q}", "message", message)
}
