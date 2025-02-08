package infrastructure

import (
	imubackend "IMUbackend/gen/imubackend"
	"fmt"
	"io"
	"mime/multipart"
)

func MarkdownDecoder(mr *multipart.Reader, p **imubackend.CreateArticlePayload) error {
	// var v *imubackend.CreateMarkdownPayload
	v := &imubackend.CreateArticlePayload{}
	for {
		part, err := mr.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to load part: %s", err)
		}
		switch part.FormName() {
		case "articleName":
			articleNameBytes, err := io.ReadAll(part)
			if err != nil {
				return fmt.Errorf("failed to read articleName: %s", err)
			}
			v.ArticleName = string(articleNameBytes)

		case "content":
			contentBytes, err := io.ReadAll(part)
			if err != nil {
				return fmt.Errorf("failed to read contentBytes: %s", err)
			}
			v.Content = string(contentBytes)

		case "token":
			tokenBytes, err := io.ReadAll(part)
			if err != nil {
				return fmt.Errorf("failed to read token: %s", err)
			}
			v.Token = string(tokenBytes)
		
		case "image":
			imgBytes, err := io.ReadAll(part)
			if err != nil {
				return fmt.Errorf("failed to read images: %s", err)
			}
			imgFile := &imubackend.File{
				Name: new(string),
				Content: imgBytes,
			}
			*imgFile.Name = part.FileName()
			v.Image = append(v.Image, imgFile)
		}
	}
	*p = v
	return nil
}
