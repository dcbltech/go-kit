package local

import (
	"context"
	"fmt"
	"io"
	"os"
)

type Storage struct {
	out string
}

func Must(out string) *Storage {
	return &Storage{out: out}
}

func (s *Storage) Exists(_ context.Context, path, name string) (exists bool, err error) {
	_, err = os.Stat(s.pathname(path, name))
	if os.IsNotExist(err) {
		return false, nil
	}

	return true, err
}

func (s *Storage) Store(_ context.Context, path, name string, data []byte) error {
	return s.storeWithGeneration(context.Background(), path, name, data, 0, 0)
}

func (s *Storage) Load(_ context.Context, path, name string) (data []byte, err error) {
	data, _, _, err = s.loadWithGeneration(context.Background(), path, name)

	return
}

func (s *Storage) Delete(_ context.Context, path, name string) error {
	return os.Remove(s.pathname(path, name))
}

func (s *Storage) DeleteAll(_ context.Context, path string) error {
	return os.RemoveAll(s.pathname(path, ""))
}

func (s *Storage) Copy(_ context.Context, srcPath, srcName, dstPath, dstName string) error {
	src, err := os.Open(s.pathname(srcPath, srcName))
	if err != nil {
		return err
	}

	defer func() { _ = src.Close() }()

	dst, err := os.Create(s.pathname(dstPath, dstName))
	if err != nil {
		return err
	}

	defer func() { _ = dst.Close() }()

	if _, err := io.Copy(dst, src); err != nil {
		return err
	}

	return nil
}

func (s *Storage) GenerateSignedURL(path, name string) (url string, err error) {
	return "", nil
}

func (s *Storage) pathname(path string, name string) string {
	o := ""

	if len(path) > 0 {
		o = fmt.Sprintf("%s/%s", path, name)
	} else {
		o = name
	}

	return fmt.Sprintf("%s/%s", s.out, o)
}

func (s *Storage) storeWithGeneration(_ context.Context, path, name string, data []byte, _, _ int64) error {
	if _, err := os.Stat(fmt.Sprintf("%s/%s", s.out, path)); os.IsNotExist(err) {
		err := os.MkdirAll(fmt.Sprintf("%s/%s", s.out, path), 0700)
		if err != nil {
			return err
		}
	}

	return os.WriteFile(s.pathname(path, name), data, 0644)
}

func (s *Storage) loadWithGeneration(_ context.Context, path, name string) (data []byte, generation, metaGeneration int64, err error) {
	reader, err := os.Open(s.pathname(path, name))
	if err != nil {
		return nil, 0, 0, err
	}

	defer func() { _ = reader.Close() }()

	data, err = io.ReadAll(reader)

	return
}
