package core

import "net/http"

func Serve(dir string, asset string, port string) {

	fs := http.FileServer(http.Dir(dir))
	http.Handle(asset, http.StripPrefix(asset, fs))

	 err :=http.ListenAndServe(port, nil)

	if err != nil{
		panic(err)
	}
	
}
