package main

import (
	"encoding/json"
	"fmt"

	"github.com/tsirysndr/go-youtube"
)

func main() {
	// Replace with your Access Token & API KEY
	token := "ya29.ImO2B1rOgUnG-DYL8DmW3nZ42uQ-7lqzHYwFYw1pVpkSKhX_nzGo84Vp7b2OPZFNDWD5RasimPdTXbm5oY4zNix_dPr743ocJKftpc6M2VomOs657ImhFXrW3Q9l-XUUiedwPGI"
	key := "AIzaSyAa8yy0GdcGPHdtD083HiGGx_S0vMPScDM"
	client := youtube.NewClient(token, key)
	result, _ := client.Search.Search("doja cat")
	r, _ := json.Marshal(result)
	fmt.Println(string(r))
}
