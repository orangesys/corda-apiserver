package corda

import (
	"bytes"
	"log"
	"net/http"
	"os"
	"os/exec"
	"text/template"

	"github.com/gin-gonic/gin"
	"orangesys.io/corda-apiserver/internal/pkg/filesystem"
)

var baseDir = "/Users/xucheng/go/src/corda-apiserver"

//Service ...
type Service interface {
	GetNodeConf(*gin.Context)
	GetCerts(*gin.Context)
}

type service struct{}

//New ...
func New() Service {
	return &service{}
}

func (s *service) GetNodeConf(c *gin.Context) {
	var req NodeConf
	if err := c.BindQuery(&req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	bytes, _, err := genNodeConf(&req)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.String(http.StatusOK, string(bytes))
}

func (s *service) GetCerts(c *gin.Context) {
	var req NodeConf
	if err := c.BindQuery(&req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	_, _, err := genNodeConf(&req)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	tmpDir := baseDir + "/tmp/" + req.UniqueName()

	if err := os.MkdirAll(tmpDir, os.ModePerm); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	go func() {
		cmd := exec.Command(
			"/bin/bash",
			"-c",
			"/Users/xucheng/go/src/corda-apiserver/initial-registration.sh "+baseDir,
		)
		_, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatalf("cmd.Run() failed with %s\n", err)
		}
		os.Rename("./aa/bb/c1/file.go", "./aa/bb/c2/file.go")

	}()

	c.String(http.StatusOK, "test")
}

func genNodeConf(nodeConf *NodeConf) (content []byte, filepath string, err error) {
	t, err := template.ParseFiles(baseDir + "/config/node.tmpl")
	if err != nil {
		return nil, "", err
	}
	buf := bytes.NewBufferString("")
	if err := t.Execute(buf, nodeConf); err != nil {
		return nil, "", err
	}
	//write to file
	path := baseDir + "/data/" + nodeConf.UniqueName() + "/node.conf"
	if err := filesystem.WriteToFile(path, buf.Bytes()); err != nil {
		return nil, "", err
	}
	return buf.Bytes(), path, nil
}
