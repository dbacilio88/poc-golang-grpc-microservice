package utils

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dbacilio88/poc-golang-grpc-microservice/internal/service"
	"github.com/dbacilio88/poc-golang-grpc-microservice/pkg/env"
	"github.com/dbacilio88/poc-golang-grpc-microservice/pkg/shared/store"
	"go.uber.org/zap"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

/**
 * files
 * <p>
 * This file contains core data structures and logic used throughout the application.
 *
 * <p><strong>Copyright © 2025 – All rights reserved.</strong></p>
 *
 * <p>This source code is distributed under a collaborative license.</p>
 *
 * <p>
 * Contributions, suggestions, and improvements are welcome!
 * You are free to fork, modify, and submit pull requests under the terms of the repository's license.
 * Please ensure proper attribution to the original author(s) and preserve this notice in derivative works.
 * </p>
 *
 * @author Christian Bacilio De La Cruz
 * @email dbacilio88@outlook.es
 * @since 5/7/2025
 */

type IFiles interface {
	ScanDir(dir string, msg service.IMessaging) error
	GenerateFile() error
}
type Files struct {
	*zap.Logger
	*store.MemoryMap
}

func NewFiles(log *zap.Logger, mm *store.MemoryMap) *Files {
	return &Files{
		Logger:    log,
		MemoryMap: mm,
	}
}

func (f *Files) OpenFile(path string) (*os.File, error) {
	f.Info("Init open file with path ", zap.String("path", path))
	file, err := os.Open(path)
	if err != nil {
		f.Error("Error opening file with path ", zap.String("path", path), zap.Error(err))
	}
	return file, err
}
func (f *Files) closeFile(file *os.File) {
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			f.Error("Error closing file with path ", zap.String("path", file.Name()), zap.Error(err))
		}
	}(file)
}

func (f *Files) ReadFile(path string) ([]byte, error) {

	file, err := f.OpenFile(path)
	f.closeFile(file)

	if err != nil {
		return nil, err
	}

	read, err := io.ReadAll(file)

	if err != nil {
		return nil, err
	}

	return read, err
}

func (f *Files) readFileContent(path string) ([]string, error) {
	start := time.Now()
	var lines []string
	f.Info("Read file content with path ", zap.String("path", path))
	file, err := f.OpenFile(path)
	if err != nil {
		return nil, err
	}
	f.closeFile(file)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	end := time.Since(start)
	f.Info("Read file content with path ", zap.String("path", path), zap.Duration("time", end))
	return lines, nil
}

func (f *Files) WriteFile(path string, data []byte) error {
	return nil
}

func (f *Files) CreateFile(path string) (*os.File, error) {
	return os.Create(path)
}

func (f *Files) ReadDir(path string) ([]os.FileInfo, error) {
	return nil, nil
}

func (f *Files) ScanDir(dir string, msg service.IMessaging) error {

	files, err := f.GetFiles(dir)
	if err != nil {
		return err
	}
	for i, file := range files {
		fmt.Println("File ", i, ": ", file.Name())
		if file.IsDir() {
			continue
		}

		ext := filepath.Ext(file.Name())
		if !f.allowedExtensions(ext) {
			fmt.Println("Extension not allowed: ", ext)
			continue
		}
		fmt.Println("Ext: ", ext)

		filePath, _ := filepath.Abs(filepath.Join(dir, file.Name()))
		fmt.Println("file-path ", filePath)

		if v := f.Get(filePath); v != "" {
			fmt.Printf(" value [%s] \n", v)
			continue
		}
		f.Put(filePath, "ok")

		tmp := f.getPathTemp(filePath)

		fmt.Println("tmp-", tmp)

		data := make(map[string]interface{})
		data["path"] = filePath
		jm, _ := json.Marshal(data)

		if err := msg.Send(jm); err != nil {
			return err
		}
	}

	return nil
}

func (f *Files) GetFiles(path string) ([]os.DirEntry, error) {
	f.Info("Get files from dir ", zap.String("path", filepath.Dir(path)))
	files, err := os.ReadDir(path)
	if err != nil {
		msg := fmt.Sprintf("Error reading dir [%s]", err.Error())
		return nil, errors.New(msg)
	}
	return files, nil
}

func (f *Files) GenerateFile() error {
	name := f.generateNameTxt()
	if err := os.WriteFile(name, []byte("hello world"), 0666); err != nil {
		return err
	}
	f.Info("Generated file ", zap.String("name", name))
	return nil
}

func (f *Files) generateNameTxt() string {
	path := env.YAML.Workspace.Files.Path
	format := time.Now().Format("20060102150405")
	name := strings.Join([]string{path, fmt.Sprintf("%s%s", format, ".txt")}, "/")
	return name
}

func (f *Files) allowedExtensions(ext string) bool {
	for _, s := range env.YAML.Workspace.Files.Allowed {
		if s == ext {
			return true
		}
	}
	return false
}

func (f *Files) getPathTemp(path string) string {
	tmp, _ := filepath.Abs(env.YAML.Workspace.Files.Tmp)
	fmt.Println("tmp: ", tmp)
	dir, _ := filepath.Abs(path)
	fmt.Println("dir: ", dir)
	if tmp == filepath.Dir(dir) {
		return dir
	}
	return tmp
}
