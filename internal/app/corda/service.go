package corda

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"text/template"

	"github.com/gin-gonic/gin"
	"orangesys.io/corda-apiserver/internal/pkg/filesystem"
)

var baseDir = "/go/src/corda-apiservice"

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
	filePath := NewNode(&req).GetHomeDir() + "/certificates.zip"
	if !filesystem.Exist(filePath) {
		if err := genCerts(&req); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

	}
	c.File(filePath)
	return
}

func genNodeConf(nodeConf *NodeConf) (content []byte, filepath string, err error) {
	t, err := template.ParseFiles(os.Getenv("CONFIG_PATH") + "/node.tmpl")
	if err != nil {
		return nil, "", err
	}
	buf := bytes.NewBufferString("")
	if err := t.Execute(buf, nodeConf); err != nil {
		return nil, "", err
	}
	//write to file
	path := NewNode(nodeConf).GetHomeDir() + "/node.conf"
	if err := filesystem.WriteToFile(path, buf.Bytes()); err != nil {
		return nil, "", err
	}
	//validate
	tmpDir := NewNode(nodeConf).GetTmpDir()
	if err := filesystem.CreateDir(tmpDir); err != nil {
		return nil, "", err
	}
	cmd := exec.Command(
		"/bin/bash",
		"-c",
		fmt.Sprintf("validate-configuration %s %s", path, tmpDir),
	)
	out, err := cmd.CombinedOutput()
	log.Printf("%s", out)
	if err != nil {
		return nil, "", errors.New("Validate configuration failed")
	}

	return buf.Bytes(), path, nil
}

func genCerts(nodeConf *NodeConf) error {
	_, path, err := genNodeConf(nodeConf)
	if err != nil {
		return err
	}
	node := NewNode(nodeConf)
	if err := filesystem.CreateDir(node.GetTmpDir()); err != nil {
		return err
	}

	//generate certs
	cmd := exec.Command(
		"/bin/bash",
		"-c",
		fmt.Sprintf("initial-registration %s %s %s", path, "trustpass", node.GetTmpDir()),
	)
	if out, err := cmd.CombinedOutput(); err != nil {
		log.Printf("%s", out)
		return err
	}

	//compressed files
	if err := filesystem.Zip(node.GetHomeDir()+"/certificates.zip", node.GetTmpDir()+"/certificates"); err != nil {
		return err
	}

	//clear tmp dir
	go os.RemoveAll(node.GetTmpDir())

	return nil
}
