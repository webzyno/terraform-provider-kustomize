package main

import (
	"os"
	"path/filepath"

	"sigs.k8s.io/kustomize/kyaml/filesys"
)

const KUSTOMIZATION = "kustomization.yaml"

type OverlayFS struct {
	memoryFS          filesys.FileSystem
	diskFS            filesys.FileSystem
	kustomizationPath string
}

func NewOverlayFS() (filesys.FileSystem, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	return &OverlayFS{
		memoryFS:          filesys.MakeFsInMemory(),
		diskFS:            filesys.MakeFsOnDisk(),
		kustomizationPath: filepath.Join(wd, KUSTOMIZATION),
	}, nil
}

func (f *OverlayFS) Create(path string) (filesys.File, error) {
	return f.diskFS.Create(path)
}

func (f *OverlayFS) Mkdir(path string) error {
	return f.diskFS.Mkdir(path)
}

func (f *OverlayFS) MkdirAll(path string) error {
	return f.diskFS.MkdirAll(path)
}

func (f *OverlayFS) RemoveAll(path string) error {
	return f.diskFS.RemoveAll(path)
}

func (f *OverlayFS) Open(path string) (filesys.File, error) {
	return f.diskFS.Open(path)
}

func (f *OverlayFS) IsDir(path string) bool {
	return f.diskFS.IsDir(path)
}

func (f *OverlayFS) ReadDir(path string) ([]string, error) {
	return f.diskFS.ReadDir(path)
}

func (f *OverlayFS) CleanedAbs(path string) (filesys.ConfirmedDir, string, error) {
	if path == f.kustomizationPath {
		_, _, err := f.memoryFS.CleanedAbs(".")
		if err != nil {
			return "", "", err
		}
		return filesys.ConfirmedDir(filepath.Dir(f.kustomizationPath)), KUSTOMIZATION, err
	}
	return f.diskFS.CleanedAbs(path)
}

func (f *OverlayFS) Exists(path string) bool {
	if path == f.kustomizationPath {
		return f.memoryFS.Exists(KUSTOMIZATION)
	}
	return f.diskFS.Exists(path)
}

func (f *OverlayFS) Glob(pattern string) ([]string, error) {
	return f.diskFS.Glob(pattern)
}

func (f *OverlayFS) ReadFile(path string) ([]byte, error) {
	if path == f.kustomizationPath {
		return f.memoryFS.ReadFile(KUSTOMIZATION)
	}
	return f.diskFS.ReadFile(path)
}

func (f *OverlayFS) WriteFile(path string, data []byte) error {
	if path == KUSTOMIZATION {
		return f.memoryFS.WriteFile(KUSTOMIZATION, data)
	}
	return f.diskFS.WriteFile(path, data)
}

func (f *OverlayFS) Walk(path string, walkFn filepath.WalkFunc) error {
	return f.diskFS.Walk(path, walkFn)
}
