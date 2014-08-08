package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"

	"code.google.com/p/go-uuid/uuid"
)

func controller(w http.ResponseWriter, r *http.Request) {
	action := r.URL.Query()["action"][0]
	fmt.Println(r.Method, action)
	if r.Method == "GET" {
		if action == "config" {
			config(w, r)
		}
	} else if r.Method == "POST" {
		if action == "uploadimage" {
			uploadImage(w, r)
		}
	}
}

func config(w http.ResponseWriter, r *http.Request) {
	w.Write(configJson)
}

func uploadImage(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("upfile")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	filename := strings.Replace(uuid.NewUUID().String(), "-", "", -1) + path.Ext(header.Filename)
	outFile, err := os.Create(path.Join("static", "upload", filename))
	if err != nil {
		panic(err)
	}

	defer outFile.Close()

	io.Copy(outFile, file)

	b, err := json.Marshal(map[string]string{
		"url":      fmt.Sprintf("/static/upload/%s", filename), //保存后的文件路径
		"title":    "",                                         //文件描述，对图片来说在前端会添加到title属性上
		"original": header.Filename,                            //原始文件名
		"state":    "SUCCESS",                                  //上传状态，成功时返回SUCCESS,其他任何值将原样返回至图片上传框中
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
	w.Write(b)
}
