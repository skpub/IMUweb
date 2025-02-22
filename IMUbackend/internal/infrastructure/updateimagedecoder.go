package infrastructure

import (
	imubackend "IMUbackend/gen/imubackend"
	"fmt"
	"io"
	"mime/multipart"
)

func UpdateImgDecoder(mr *multipart.Reader, p **imubackend.UpdateImgPayload) error {
	v := &imubackend.UpdateImgPayload{}
	for {
		part, err := mr.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to load part: %s", err)
		}
		switch part.FormName() {
		case "token":
			tokenBytes, err := io.ReadAll(part)
			if err != nil {
				return fmt.Errorf("failed to read token: %s", err)
			}
			v.Token = string(tokenBytes)
		case "img":
			imgBytes, err := io.ReadAll(part)
			if err != nil {
				return fmt.Errorf("failed to read img: %s", err)
			}
			imgFile := &imubackend.File{
				Name: new(string),
				Content: imgBytes,
			}
			*imgFile.Name = part.FileName()
			v.Img = imgFile
		}
	}
	*p = v
	return nil
}
