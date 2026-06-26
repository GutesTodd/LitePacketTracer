// Package engines
package engines

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

type JSONRepositoryEngine struct {
	fpath string
}

func (e *JSONRepositoryEngine) Read(target any) error {
	if target == nil {
		return errors.New("target must not be nil")
	}
	finfo, err := os.Stat(e.fpath)
	if err != nil {
		if os.IsNotExist(err) {
			return errors.New("file does not exist")
		}
		if os.IsPermission(err) {
			return errors.New("file dont have permission")
		}
		return err
	}
	if finfo.IsDir() {
		return errors.New("filepath is directory")
	}
	if mode := finfo.Mode(); mode%100 < 4 {
		return errors.New("file does have permission for read")
	}
	data, err := os.ReadFile(e.fpath)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, target)
}

func (e *JSONRepositoryEngine) Save(target any) error {
	if target == nil {
		return errors.New("target must not be nil")
	}

	finfo, err := os.Stat(e.fpath)
	if err == nil && finfo.IsDir() {
		return errors.New("filepath is a directory")
	}
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	dir := filepath.Dir(e.fpath)
	err = os.MkdirAll(dir, 0o755)
	if err != nil {
		return err
	}
	bdata, err := json.MarshalIndent(target, "", "   ")
	if err != nil {
		return err
	}
	return os.WriteFile(e.fpath, bdata, 0o644)
}
