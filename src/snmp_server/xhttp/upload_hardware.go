package xhttp

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"snmp_server/globalvars"

	"github.com/cihub/seelog"

	"github.com/gin-gonic/gin"
)

func uploadHardware(c *gin.Context) {
	devType := c.PostForm("dev_type")
	seelog.Info("dev type is :", devType)
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusOK, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	devTpeFtpDir := globalvars.FTPDir + devType
	dirFile, err := os.Open(devTpeFtpDir)

	seelog.Infof("open %s result %v,%v", devTpeFtpDir, dirFile, err)
	if err == nil {
		dirFile.Close()
	} else {
		//create dir for devType
		seelog.Warnf("Open %s err:%s", devTpeFtpDir, err)
		if os.IsNotExist(err) {
			err = os.MkdirAll(devTpeFtpDir, os.ModeDir|0777)
			if err != nil {
				c.String(http.StatusOK, fmt.Sprintf("create FTP DIR error: %s", err.Error()))
				return
			}
			cmd := exec.Command("chown", globalvars.GetFTPChown(), globalvars.FTPDir+devType)
			err = cmd.Run()
			if err != nil {
				c.String(http.StatusOK, fmt.Sprintf("chown ftpdir devtype err: %s", err.Error()))
				return
			}
		} else {
			if err != nil {
				c.String(http.StatusOK, fmt.Sprintf("Open FTP DIR error: %s", err.Error()))
				return
			}
		}
	}

	filename := devTpeFtpDir + "/" + file.Filename
	os.Remove(filename)
	if err := c.SaveUploadedFile(file, filename); err != nil {
		c.String(http.StatusOK, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}
	cmd := exec.Command("chown", globalvars.GetFTPChown(), filename)
	err = cmd.Run()
	if err != nil {
		c.String(http.StatusOK, fmt.Sprintf("chown file err: %s", err.Error()))
		return
	}

	c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully with fields .", file.Filename))
}

func uploadHardwareWeb(c *gin.Context) {

	template := `<!doctype html>
	<html lang="en">
	<head>
		<meta charset="utf-8">
		<title>Single file upload</title>
	</head>
	<body>
	<h1>Upload single file with fields</h1>
	
	<form action="/v1/upload_hardware" method="post" enctype="multipart/form-data">
		Files: <input type="file" name="file"><br><br>
		device type: 
		<select name="dev_type">
		<option value="KN518" selected>KN518</option>
		<option value="KN519">KN519</option>
		</select><br><br>
		<input type="submit" value="Submit">
	</form>
	</body>`
	c.Data(http.StatusOK, c.ContentType(), []byte(template))
}
