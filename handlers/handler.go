package handlers

import (
	"log"
	"net/http"
	"runtime/debug"
)

func PanicOnError(err error){
	if err != nil{
		log.Panicln("ERROR: ", err)
	}
}

////Must always be called deferred, i.e.
// defer handlers.HandlePanic()
func HandlePanic(){
	if r := recover(); r != nil{
		debug.PrintStack()
		log.Println("ERROR: ", r)
	}
}

func NetHandlePanic(w http.ResponseWriter){
	if r := recover(); r != nil{
		debug.PrintStack()
		log.Println("ERROR: ", r)
		http.Error(w, "Oops, something went wrong", 500)
	}
}
