package goseaweedfs

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Filer struct {
	base    *url.URL
	client  *httpClient
	authKey string
}

type FilerUploadResult struct {
	Name    string `json:"name"`
	FileURL string `json:"fileURL"`
	FileID  string `json:"fileID"`
	Size    int64  `json:"size"`
	Error   string `json:"error"`
}

func NewFiler(u string, client *http.Client, opts ...FilerOption) (f *Filer, err error) {
	return newFiler(u, client, opts...)
}

func newFiler(u string, client *http.Client, opts ...FilerOption) (f *Filer, err error) {
	base, err := parseURI(u)
	if err != nil {
		return
	}

	f = &Filer{
		base: base,
	}

	for _, opt := range opts {
		opt(f)
	}

	var clientOpts []HttpClientOption
	if f.authKey != "" {
		clientOpts = append(clientOpts, WithHttpClientAuthKey(f.authKey))
	}
	f.client = newHttpClient(client, clientOpts...)

	return f, nil
}

func (f *Filer) Close() (err error) {
	if f.client != nil {
		f.client.Close()
	}
	return
}

func (f *Filer) UploadFile(localFilePath, newPath, collection, ttl string) (result *FilerUploadResult, err error) {
	fp, err := NewFilePart(localFilePath)
	if err != nil {
		return result, err
	}
	defer fp.Close()

	var fileReader io.Reader
	if fp.FileSize == 0 {
		fileReader = bytes.NewBuffer(EmptyMark.Bytes())
	} else {
		fileReader = fp.Reader
	}

	var data []byte
	data, status, err := f.client.upload(encodeURI(*f.base, newPath, normalize(nil, collection, ttl)), localFilePath, fileReader, fp.MimeType)
	if err != nil {
		if !strings.Contains(err.Error(), "The process cannot access the file because another process has locked a portion of the file.") {
			return result, err
		}
	}

	var res FilerUploadResult
	if err = json.Unmarshal(data, &res); err != nil {
		if status == 404 {
			return nil, errors.New("404 not found")
		}
		return result, err
	}
	result = &res

	if status >= 400 {
		return result, errors.New(res.Error)
	}

	return result, nil
}

func (f *Filer) UploadDir(localDirPath, newPath, collection, ttl string) (results []*FilerUploadResult, err error) {
	// normalizer path
	localDirPath = normalizePath(localDirPath)
	localDirPath = strings.TrimSuffix(localDirPath, "/")

	if !strings.HasPrefix(newPath, "/") {
		newPath = "/" + newPath
	}
	files, err := ListFilesRecursive(localDirPath)
	if err != nil {
		return results, err
	}
	for _, info := range files {
		filePath := normalizePath(info.Path)
		newFilePath := newPath + strings.ReplaceAll(filePath, localDirPath, "")
		result, err := f.UploadFile(filePath, newFilePath, collection, ttl)
		if err != nil {
			return results, err
		}
		results = append(results, result)
	}
	return
}

func (f *Filer) Upload(content io.Reader, fileSize int64, newPath, collection, ttl string) (result *FilerUploadResult, err error) {
	fp := NewFilePartFromReader(io.NopCloser(content), newPath, fileSize)

	var data []byte
	data, _, err = f.client.upload(encodeURI(*f.base, newPath, normalize(nil, collection, ttl)), newPath, io.NopCloser(content), "")
	if err == nil {
		result = &FilerUploadResult{}
		err = json.Unmarshal(data, result)
	}

	_ = fp.Close()
	return
}

func (f *Filer) ListDir(path string) (files []FilerFileInfo, err error) {
	data, _, err := f.GetJson(path, nil)
	if err != nil {
		return files, err
	}
	if len(data) == 0 {
		return
	}
	var res FilerListDirResponse
	if err := json.Unmarshal(data, &res); err != nil {
		return files, err
	}

	for _, file := range res.Entries {
		file = GetFileWithExtendedFields(file)
		files = append(files, file)
	}
	return
}

func (f *Filer) ListDirRecursive(path string) (files []FilerFileInfo, err error) {
	entries, err := f.ListDir(path)
	if err != nil {
		return files, err
	}
	for _, file := range entries {
		file = GetFileWithExtendedFields(file)
		if file.IsDir {
			file.Children, err = f.ListDirRecursive(file.FullPath)
			if err != nil {
				return files, err
			}
		}
		files = append(files, file)
	}
	return
}

func (f *Filer) Get(path string, args url.Values, header map[string]string) (data []byte, statusCode int, err error) {
	data, statusCode, err = f.client.get(encodeURI(*f.base, path, args), header)
	return
}

func (f *Filer) GetJson(path string, args url.Values) (data []byte, statusCode int, err error) {
	header := map[string]string{
		"Accept": "application/json",
	}
	data, statusCode, err = f.client.get(encodeURI(*f.base, path, args), header)
	return
}

func (f *Filer) Download(path string, args url.Values, callback downloadCB) (err error) {
	_, err = f.client.download(encodeURI(*f.base, path, args), callback)
	return
}

func (f *Filer) Delete(path string, args url.Values) (err error) {
	_, err = f.client.delete(encodeURI(*f.base, path, args))
	return
}

func (f *Filer) DeleteDir(path string) (err error) {
	args := map[string][]string{"recursive": {"true"}}
	_, err = f.client.delete(encodeURI(*f.base, path, args))
	return
}

func (f *Filer) DeleteFile(path string) (err error) {
	_, err = f.client.delete(encodeURI(*f.base, path, nil))
	return
}
