package handlers

import (
	"log"
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
