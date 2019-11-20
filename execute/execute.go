package execute

import (
	"fmt"
	"github.com/LazyMechanic/sortman/cli/dialog"
	"github.com/LazyMechanic/sortman/types"
	"github.com/otiai10/copy"
	"github.com/yargevad/filepathx"
	"os"
	"path/filepath"
	"strings"
)

func filesToExecute(request *types.Request) ([]string, error) {
	var result []string

	var files []string
	var exclusions []string

	for _, pattern := range request.Patterns {
		filesTmp, err := filepathx.Glob(pattern)
		if err != nil {
			return nil, err
		}
		files = append(files, filesTmp...)
	}

	for _, exclusion := range request.Exclusions {
		exclusionsTmp, err := filepathx.Glob(exclusion)
		if err != nil {
			return nil, err
		}
		exclusions = append(exclusions, exclusionsTmp...)
	}

	for _, file := range files {
		var isExclusion = false
		for _, exclusion := range exclusions {
			isExclusion = strings.Contains(file, exclusion)
			if isExclusion {
				break
			}
		}

		if !isExclusion {
			result = append(result, file)
		}
	}

	return result, nil
}

func isExist(file string) bool {
	_, err := os.Stat(file)
	return !os.IsNotExist(err)
}

func doAction(request *types.Request, action func(from, to string) error) error {
	for _, inFile := range request.Files {
		if err := os.MkdirAll(request.OutDirectory, os.ModePerm); err != nil {
			return err
		}

		var fileFullName = filepath.Base(inFile)
		var fileExtension = filepath.Ext(inFile)
		var fileName = fileFullName[0 : len(fileFullName)-len(fileExtension)]

		var outFile = filepath.Join(request.OutDirectory, fileFullName)

		var postfix = 1
		// If file already exists
		for isExist(outFile) {
			// Prepend out directory
			outFile = filepath.Join(request.OutDirectory, fmt.Sprintf("%s_%d%s", fileName, postfix, fileExtension))
			postfix += 1
		}

		var err = action(inFile, outFile)

		if err != nil {
			return err
		}
	}
	return nil
}

func Execute(config *types.Config) error {
	var err error
	for index, _ := range config.Requests {
		var request = &config.Requests[index]

		// Prepend .InDirectory to patterns
		for patternIndex, _ := range request.Patterns {
			var pattern = &request.Patterns[patternIndex]
			*pattern = filepath.Join(request.InDirectory, *pattern)
		}

		// Prepend .InDirectory to exclusions
		for exclusionIndex, _ := range request.Exclusions {
			var exclusion = &request.Exclusions[exclusionIndex]
			*exclusion = filepath.Join(request.InDirectory, *exclusion)
		}

		request.Files, err = filesToExecute(request)
		if err != nil {
			return err
		}

		switch config.Action {
		case dialog.CopyAction:
			err = doAction(request, copy.Copy)
		case dialog.MoveAction:
			err = doAction(request, os.Rename)
		default:
			return fmt.Errorf("Wrong action")
		}

		if err != nil {
			return err
		}
	}

	return nil
}
