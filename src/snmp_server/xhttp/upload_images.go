package xhttp

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var imageSuffixes = map[string]interface{}{
	"jpg":  nil,
	"jpeg": nil,
	"gif":  nil,
	"png":  nil,
	"bmp":  nil,
}

func uploadImage(c *gin.Context) {

	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusOK, fmt.Sprintf("err: %s", err.Error()))
		return
	}
	fields := strings.Split(file.Filename, ".")
	suffix := strings.ToLower(fields[len(fields)-1])
	if _, ok := imageSuffixes[suffix]; !ok {
		c.String(http.StatusOK, fmt.Sprintf("err:not support image file ext name(%s), only suport [*.jpg, *.jpeg, *.gif, *.png, *.bmp]", suffix))
		return
	}
	filename := "./images/" + time.Now().Format("20060102150402") + "." + suffix
	if err := c.SaveUploadedFile(file, filename); err != nil {
		c.String(http.StatusOK, fmt.Sprintf("err: %s", err.Error()))
		return
	}
	c.String(http.StatusOK, strings.TrimPrefix(filename, "./"))
}

func uploadImagesWeb(c *gin.Context) {

	template := `<!doctype html>
	<html lang="en">
	<head>
		<meta charset="utf-8">
		<title>Single image file upload</title>
	</head>
	<body>
	<h1>Upload single file with fields</h1>
	
	<form action="/v1/upload_image" method="post" enctype="multipart/form-data">
		Files: <input type="file" name="file"><br><br>
		<input type="submit" value="Submit">
	</form>
	</body>`
	c.Data(http.StatusOK, c.ContentType(), []byte(template))
}
