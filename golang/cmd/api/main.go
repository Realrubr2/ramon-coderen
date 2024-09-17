package main

import (
	"fmt"
	"net/http"
	"github.com/Realrubr2/ramon-coderen/golang/internal/handlers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main(){
	log.SetReportCaller(true)
	var router *chi.Mux = chi.NewRouter()
	handlers.Handler(router)

	fmt.Println("starting go API service .....");

	fmt.Println(`
  rrrrr   eeeee    aaaa   l        rrrrr   bbbbb    rrrrr    -  aaaa  pppp   iii
  r    r  e       a    a  l        r    r  b    b   r    r   - a    a p   p   i
  rrrrr   eee     aaaaaa  l        rrrrr   bbbbb    rrrrr    - aaaaaa pppp    i
  r  r    e       a    a  l        r  r    b    b   r  r     - a    a p       i
  r   r   eeeee   a    a  llllll   r   r   bbbbb    r   r    - a    a p       iii
	`)
	err:= http.ListenAndServe("localhost:8080", router)
	if err!= nil {
		log.Error(err)
	}
}
