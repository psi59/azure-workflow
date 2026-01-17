package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"azure-workflow/internal/alfred"
	"azure-workflow/internal/azure"
	"azure-workflow/internal/search"
)

func ProcessQuery(servicesFile, query string) (string, error) {
	services, err := azure.LoadServices(servicesFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "DEBUG: LoadServices error: %v\n", err)
		return "", err
	}
	fmt.Fprintf(os.Stderr, "DEBUG: loaded %d services\n", len(services))
	if len(services) > 0 {
		fmt.Fprintf(os.Stderr, "DEBUG: first service: %+v\n", services[0])
	}

	fmt.Fprintf(os.Stderr, "DEBUG: query='%s', len=%d\n", query, len(query))
	results := search.Search(services, query)
	fmt.Fprintf(os.Stderr, "DEBUG: search results: %d\n", len(results))

	items := alfred.NewItemsFromServices(results)
	fmt.Fprintf(os.Stderr, "DEBUG: items: %d\n", len(items))

	output := alfred.Output{Items: items}

	jsonBytes, err := json.Marshal(output)
	if err != nil {
		return "", err
	}

	return string(jsonBytes), nil
}

func main() {
	fmt.Fprintf(os.Stderr, "DEBUG: os.Args=%v, len=%d\n", os.Args, len(os.Args))
	query := ""
	if len(os.Args) > 1 {
		query = os.Args[1]
	}

	// 디버그: 현재 작업 디렉토리
	cwd, _ := os.Getwd()
	fmt.Fprintf(os.Stderr, "DEBUG: cwd=%s\n", cwd)

	// 디버그: 실행 파일 경로
	execPath, _ := os.Executable()
	fmt.Fprintf(os.Stderr, "DEBUG: execPath=%s\n", execPath)

	// 디버그: symlink 해결
	realPath, _ := filepath.EvalSymlinks(execPath)
	fmt.Fprintf(os.Stderr, "DEBUG: realPath=%s\n", realPath)

	// services.yaml 파일 경로 찾기 (Alfred는 workflow 폴더를 작업 디렉토리로 사용)
	servicesFile := "services.yaml"

	// 현재 디렉토리에 없으면 실행 파일 경로 기준으로 찾기
	if _, err := os.Stat(servicesFile); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "DEBUG: services.yaml not found in cwd, trying exec path\n")
		execPath, err := os.Executable()
		if err == nil {
			// symlink를 해결하여 실제 경로 찾기
			realPath, err := filepath.EvalSymlinks(execPath)
			if err == nil {
				servicesFile = filepath.Join(filepath.Dir(realPath), "services.yaml")
			} else {
				servicesFile = filepath.Join(filepath.Dir(execPath), "services.yaml")
			}
		}
	}

	fmt.Fprintf(os.Stderr, "DEBUG: servicesFile=%s\n", servicesFile)

	// 파일 존재 여부 확인
	if _, err := os.Stat(servicesFile); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "DEBUG: servicesFile does not exist!\n")
	} else {
		fmt.Fprintf(os.Stderr, "DEBUG: servicesFile exists\n")
	}

	output, err := ProcessQuery(servicesFile, query)
	if err != nil {
		fmt.Fprintf(os.Stderr, "DEBUG: ProcessQuery error: %s\n", err.Error())
		fmt.Printf(`{"items":[{"title":"Error","subtitle":"%s"}]}`, err.Error())
		return
	}

	fmt.Fprintf(os.Stderr, "DEBUG: output length=%d\n", len(output))
	fmt.Println(output)
}
