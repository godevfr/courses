package main

import (
	"encoding/json"
	"fmt"

	"gioui.org/app"
	"github.com/godevfr/courses/internal/gui"
	fb "github.com/huandu/facebook/v2"
)

var fbparams = fb.Params{
	"fields":       "first_name,albums{photos{images}}",
	"access_token": "EAAN7ZBnr7RI8BAPkLZBdZCuClkcx1hguag9os5jOqRps9m5cIWZBshon7wsXVAALg3kEXwTv0B3ZABo6se4ZCL8iQ7iyUYCUWkyZAAKpggn1eV4VYY07tItyk7pwa0RhC6oMJ9POCVNmyFtGi1fwzT8NwQ8ZAkQ6AJyAtF0Faox39QkcybQMme25kJzbNZB8QekyQ94D3JDMuK6psVByZCKSLYz8TGMiQjrojHLg7VeubZCuwZDZD",
}

func main() {
	res, err := fb.Get("/me", fbparams)
	if err != nil {
		panic(err)
	}
	b, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}
	var album Albums

	err = json.Unmarshal(b, &album)
	if err != nil {
		panic(err)
	}
	var data string
	for _, alb := range album.Albums.Data {
		for _, photo := range alb.Photos.Data {
			data += fmt.Sprintln(photo.Images[0].Source)
		}
	}

	go gui.StartGUI("FB pictures", "cheval")
	app.Main()

}

type Albums struct {
	FirstName string `json:"first_name"`
	Albums    struct {
		Data []struct {
			Photos struct {
				Data []struct {
					Images []struct {
						Height int    `json:"height"`
						Source string `json:"source"`
						Width  int    `json:"width"`
					} `json:"images"`
					ID string `json:"id"`
				} `json:"data"`
				Paging struct {
					Cursors struct {
						Before string `json:"before"`
						After  string `json:"after"`
					} `json:"cursors"`
					Next string `json:"next"`
				} `json:"paging"`
			} `json:"photos"`
			ID string `json:"id"`
		} `json:"data"`
		Paging struct {
			Cursors struct {
				Before string `json:"before"`
				After  string `json:"after"`
			} `json:"cursors"`
			Next string `json:"next"`
		} `json:"paging"`
	} `json:"albums"`
	ID string `json:"id"`
}
