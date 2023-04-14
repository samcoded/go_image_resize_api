package utils

import (
	"errors"
	"image/jpeg"
	"net/http"
	"os"
	"path"
	"strconv"

	"github.com/nfnt/resize"
)

func Resize(url, width, height string) (string, error) {

	// get name of file from url
	filename := path.Base(url)
	// remove extension from filename
	filename = filename[0 : len(filename)-len(path.Ext(filename))]
	newFilename := "thumb/" + filename + "_" + height + "_" + width + ".jpg"

	// check if file already exists
	// if _, err := os.Stat(newFilename); err == nil {
	// 	return newFilename, nil
	// }

	// get image from url
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return "", errors.New("status code error: " + response.Status)
	}

	img, err := jpeg.Decode(response.Body)
	if err != nil {
		return "", err
	}

	// convert width and height to int
	widthInt, _ := strconv.Atoi(width)

	heightInt, _ := strconv.Atoi(height)

	resizedImg := resize.Resize(uint(widthInt), uint(heightInt), img, resize.Lanczos3)

	out, err := os.Create(newFilename)
	if err != nil {
		return "", err
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, resizedImg, nil)

	// filename := "thumb/test_resizeds.jpg"

	return newFilename, nil

}
