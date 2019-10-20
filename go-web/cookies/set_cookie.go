package main

import (
	"fmt"
	"net/http"
)

func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "first_cookie",
		Value:    "Go Web",
		HttpOnly: true,
	}

	c2 := http.Cookie{
		Name:     "second_cookie",
		Value:    "Manning Books",
		HttpOnly: true,
	}
	//* we use the Set method to add the first cookie
	// w.Header().Set("Set-Cookie", c1.String())
	//* we use the Add method to add the second cookie
	// w.Header().Add("Set-Cookie", c2.String())

	//* we can also use setCookie to send cookies to the browser but, they need to be passed in as reference
	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)

}

func getCookie(w http.ResponseWriter, r *http.Request) {
	h := r.Header["Cookie"]
	fmt.Fprintln(w, h)
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/set_cookie", setCookie)
	http.HandleFunc("/get_cookie", getCookie)
	server.ListenAndServe()
}
