package goseaweedfs

import (
	"io"
	"mime"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type FilePart struct {
	Reader     io.ReadCloser
	FileName   string
	FileSize   int64
	MimeType   string
	ModTime    int64
	Collection string

	// TTL Time to live.
	// 3m: 3 minutes
	// 4h: 4 hours
	// 5d: 5 days
	// 6w: 6 weeks
	// 7M: 7 months
	// 8y: 8 years
	TTL    string
	Server string
	FileID string
}

func (f *FilePart) Close() (err error) {
	err = f.Reader.Close()
	return
}

func NewFilePartFromReader(reader io.ReadCloser, fileName string, fileSize int64) *FilePart {
	ret := &FilePart{
		Reader:   reader,
		FileName: fileName,
		FileSize: fileSize,
	}

	ext := strings.ToLower(path.Ext(fileName))
	if ext != "" {
		ret.MimeType = mime.TypeByExtension(ext)
	}

	return ret
}

func NewFilePart(fullPathFileName string) (*FilePart, error) {
	fh, openErr := os.Open(fullPathFileName)
	if openErr != nil {
		return nil, openErr
	}

	ret := &FilePart{
		Reader:   fh,
		FileName: filepath.Base(fullPathFileName),
	}
	if fi, fiErr := fh.Stat(); fiErr == nil {
		ret.ModTime = fi.ModTime().UTC().Unix()
		ret.FileSize = fi.Size()
	} else {
		return nil, fiErr
	}

	ext := strings.ToLower(path.Ext(ret.FileName))
	if ext != "" {
		ret.MimeType = mime.TypeByExtension(ext)
	}
	return ret, nil
}

func NewFileParts(fullPathFileNames []string) (ret []*FilePart, err error) {
	ret = make([]*FilePart, len(fullPathFileNames))
	for _, file := range fullPathFileNames {
		if fp, err := NewFilePart(file); err == nil {
			ret = append(ret, fp)
		} else {
			closeFileParts(ret)
			return nil, err
		}
	}

	return
}

func closeFileParts(fps []*FilePart) {
	for i := range fps {
		fps[i].Close()
	}
}
