package main

import (
	"errors"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrSameFile              = errors.New("same file")
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	if fromPath == "" {
		return ErrUnsupportedFile
	}
	if toPath == "" {
		return ErrUnsupportedFile
	}
	if fromPath == toPath {
		return ErrSameFile
	}

	fromFile, err := os.OpenFile(fromPath, os.O_RDONLY, 0o755)
	if err != nil {
		return err
	}
	defer fromFile.Close()

	toFile, err := os.OpenFile(toPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o755)
	if err != nil {
		return err
	}
	defer toFile.Close()

	fi, err := fromFile.Stat()
	if err != nil {
		return err
	}

	if offset > fi.Size() {
		return ErrOffsetExceedsFileSize
	}

	if limit == 0 || limit > fi.Size() {
		limit = fi.Size()
	}

	if offset+limit > fi.Size() {
		limit = fi.Size() - offset
	}

	//

	buff := make([]byte, limit)
	reader := io.ReaderAt(fromFile)

	if _, err := reader.ReadAt(buff, offset); err != nil {
		return err
	}

	writer := io.Writer(toFile)
	if _, err := writer.Write(buff); err != nil {
		return err
	}

	count := int(limit)
	bar := pb.StartNew(count)
	for i := 0; i < count; i++ {
		bar.Increment()
	}
	bar.Finish()

	return nil
}
