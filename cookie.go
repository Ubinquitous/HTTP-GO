func handler(w http.ResponseWriter, r *http.Request){
	w.Header().Add("Set-Cookie", "VISIT=TRUE")

	if _, ok := r.Header["Cookie"]; ok {
		fmt.Fprintf(w, "<html><body>You have some cookie.</body></html>\n")
	} else {
		fmt.Fprintf(w, "<html><body>Welcome.</body></html>\n")
	}
}

/*
> console.log(document.cookie);
"_ga=GA1.2....; c_user=100002291...; csm=2; p = 02; act=147--2358...;..."
*/
/*
$ curl --http1.0 -c cookie.txt -b cookie.txt -b "name=value"
http://example.com/helloworld
*/