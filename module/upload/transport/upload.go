package uploadtransport

import (
	"bytes"
	"edx_go/common"
	app_ctx "edx_go/component"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"path/filepath"
)

func Upload(appCtx app_ctx.AppContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		file, err := ctx.FormFile("file")

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		fileOpen, err := file.Open()

		//if err != nil {
		//	panic(common.ErrInvalidRequest(err))
		//}

		//defer fileOpen.Close()

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		part, _ := writer.CreateFormFile("file", filepath.Base(file.Filename))
		io.Copy(part, fileOpen)
		writer.Close()

		r, _ := http.NewRequest("POST", "https://my-edx-nest.herokuapp.com/api/v1/upload", body)
		r.Header.Add("Content-Type", writer.FormDataContentType())
		client := &http.Client{}
		do, err := client.Do(r)
		if err != nil {
			panic(err)
		}

		read, err := ioutil.ReadAll(do.Body)

		if err != nil {
			panic(err)
		}

		var target map[string]string

		json.Unmarshal(read, &target)

		ctx.JSON(http.StatusOK, target)

	}
}
