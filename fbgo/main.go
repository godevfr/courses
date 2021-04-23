package main

import (
	"encoding/json"
	"fmt"
	"image"
	_ "image/jpeg"
	"net/http"

	"gioui.org/app"
	"github.com/godevfr/courses/gui"
	fb "github.com/huandu/facebook/v2"
)

var fbparams = fb.Params{
	"fields":       "first_name,albums{photos{images}}",
	"access_token": "XXX",
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
	var imgURLs []string
	for _, alb := range album.Albums.Data {
		for _, photo := range alb.Photos.Data {
			imgURLs = append(imgURLs, photo.Images[0].Source)
		}
	}

	var imgs []image.Image
	for _, imgURL := range imgURLs {
		img, err := fetchImage(imgURL)
		if err != nil {
			panic(err)
		}
		imgs = append(imgs, img)
	}

	go gui.StartGUI("FB pictures", imgs)
	app.Main()
}

func fetchImage(url string) (image.Image, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("fetchImage: http.Get(%q): %v", url, err)
	}
	defer resp.Body.Close()
	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("fetchImage: image decode failed: %v", err)
	}
	return img, nil
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
