/*
Copyright 2024 API Testing Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language 24 permissions and
limitations under the License.
*/
package render

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/flopp/go-findfont"
	"github.com/linuxsuren/api-testing/pkg/util"
	"github.com/signintech/gopdf"
)

func generateRandomPdf(content string) (data string, err error) {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})

	var fontPath string
	if fontPath, err = findfont.Find(""); err == nil {
		if err = pdf.AddTTFFont("wts11", fontPath); err != nil {
			return
		}

		if err = pdf.SetFont("wts11", "", 14); err != nil {
			return
		}
	} else {
		return
	}

	pdf.AddHeader(func() {})
	pdf.AddFooter(func() {})
	pdf.AddPage()
	if err = pdf.Text(fmt.Sprintf("%s\nGenerated by atest at %v", content, time.Now())); err != nil {
		err = fmt.Errorf("failed to generate pdf content: %v", err)
		return
	}

	buf := new(bytes.Buffer)
	if _, err = pdf.WriteTo(buf); err == nil {
		data = fmt.Sprintf("%s%s", util.PDFBase64Prefix, base64.StdEncoding.EncodeToString(buf.Bytes()))
	}
	return
}
