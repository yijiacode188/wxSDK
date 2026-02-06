package notifyMessage

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/yijiacode188/wxSDK/subscription/handler/notifyMessage/vo"
	"github.com/yijiacode188/wxSDK/utils"
	"io"
	"mime/multipart"
	"net/url"
	"os"
	"path/filepath"
)

// UploadImage 上传图文消息图片
// 本接口用于上传图文消息内所需的图片
// https://developers.weixin.qq.com/doc/service/api/notify/message/api_uploadimage.html
func (wx *NotifyMessage) UploadImage(fileHeader *multipart.FileHeader) (string, *utils.Response, error) {
	body, contentType, err := CreateMultipartFormData("media", fileHeader)
	if err != nil {
		return "", nil, err
	}
	token, _, err := wx.GetStableAccessToken(false)
	if err != nil {
		return "", nil, err
	}
	params := make(url.Values)
	params.Add("access_token", token)

	result, response, err := utils.HttpPost[vo.UploadImageResponse](&utils.RequestParams{
		Url:         "https://api.weixin.qq.com/cgi-bin/media/uploadimg",
		ContentType: contentType,
		Params:      params,
		Body:        body,
	})
	if err != nil {
		return "", nil, err
	}
	if result.ErrCode != 0 {
		return "", response, errors.New(result.ErrMsg)
	}
	return result.Url, response, nil
}

// CreateMultipartFormData *multipart.FileHeader 构建field 的formData请求
func CreateMultipartFormData(field string, fileHeader *multipart.FileHeader) ([]byte, string, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	// 打开文件
	file, err := fileHeader.Open()

	defer file.Close()
	fieldName := fileHeader.Filename
	if fieldName == "" {
		fieldName = fmt.Sprintf("file")
	}
	if err != nil {
		return nil, "", fmt.Errorf("打开文件 %s 失败: %v", fieldName, err)
	}

	// 创建表单文件部分
	part, err := writer.CreateFormFile(field, fileHeader.Filename)
	if err != nil {
		return nil, "", fmt.Errorf("创建文件部分失败: %v", err)
	}

	// 复制文件内容
	if _, err = io.Copy(part, file); err != nil {
		return nil, "", fmt.Errorf("复制文件内容失败: %v", err)
	}

	if err = writer.Close(); err != nil {
		return nil, "", fmt.Errorf("关闭写入器失败: %v", err)
	}
	return body.Bytes(), writer.FormDataContentType(), nil
}

// GetFileHeader 本地文件生成 *multipart.FileHeader
func GetFileHeader(filePath string) (*multipart.FileHeader, error) {
	// 1. Open the local file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// 2. Create a buffer to hold the multipart form data
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	// 3. Create a form file part
	// "file" is the form field name, filepath.Base(filePath) is the filename
	part, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		return nil, fmt.Errorf("failed to create form file: %w", err)
	}

	// 4. Copy the file content into the part
	_, err = io.Copy(part, file)
	if err != nil {
		return nil, fmt.Errorf("failed to copy file content: %w", err)
	}

	// 5. Close the writer to write the boundary
	err = writer.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to close writer: %w", err)
	}

	// 6. Create a multipart reader to read back the form data
	// This constructs the FileHeader structure properly
	reader := multipart.NewReader(&body, writer.Boundary())

	// 7. Parse the form to get the FileHeader
	// 10 MB max memory, but since we are reading from memory buffer, it should be fine
	form, err := reader.ReadForm(10 << 20)
	if err != nil {
		return nil, fmt.Errorf("failed to read form: %w", err)
	}
	defer form.RemoveAll()

	// 8. Extract the FileHeader
	files := form.File["file"]
	if len(files) == 0 {
		return nil, fmt.Errorf("no file found in form")
	}

	// The first file is what we want
	// Note: The file content is now in memory or temp file managed by the form.
	// If the caller needs to read it, they can use Open() on the FileHeader.
	return files[0], nil
}
