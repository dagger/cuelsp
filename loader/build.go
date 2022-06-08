package loader

import (
	"fmt"
	"io/fs"
	"path"
	"path/filepath"
	"sync"

	cueload "cuelang.org/go/cue/load"
)

// Dir load a cue instance from a directory
func Dir(src, file string) (*Instance, error) {
	return Build(src, nil, "./"+filepath.Dir(file))
}

// File load a cue instance from a single file
func File(src, file string) (*Instance, error) {
	return Build(src, nil, file)
}

// Build a cue instance from the files in fs.
func Build(src string, overlays map[string]fs.FS, file string) (*Instance, error) {
	var muCfg sync.RWMutex
	buildConfig := &cueload.Config{
		Dir:     src,
		Overlay: map[string]cueload.Source{},
	}

	// Map the source files into the overlay
	for mnt, f := range overlays {
		f := f
		mnt := mnt
		err := fs.WalkDir(f, ".", func(p string, entry fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if !entry.Type().IsRegular() {
				return nil
			}

			if filepath.Ext(entry.Name()) != ".cue" {
				return nil
			}

			contents, err := fs.ReadFile(f, p)
			if err != nil {
				return fmt.Errorf("%s: %w", p, err)
			}

			overlayPath := path.Join(buildConfig.Dir, mnt, p)
			muCfg.Lock()
			buildConfig.Overlay[overlayPath] = cueload.FromBytes(contents)
			muCfg.Unlock()
			return nil
		})
		if err != nil {
			return nil, err
		}
	}
	instances := cueload.Instances([]string{file}, buildConfig)

	instance := instances[0]
	if err := instance.Err; err != nil {
		return nil, err
	}

	i := NewInstance(instance)
	if err := i.Validate(); err != nil {
		return nil, err
	}

	return i, nil
}
