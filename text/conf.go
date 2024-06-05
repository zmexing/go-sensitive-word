package text

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func Files() ([]string, error) {
	filePath, err := CallerDir()
	if err != nil {
		return nil, err
	}
	fmt.Println(filePath)
	return []string{
		filePath + "/COVID-19词库.txt",
		filePath + "/其他词库.txt",
		filePath + "/反动词库.txt",
		filePath + "/暴恐词库.txt",
		filePath + "/民生词库.txt",
		filePath + "/色情词库.txt",
		filePath + "/补充词库.txt",
		filePath + "/贪腐词库.txt",
		filePath + "/零时-Tencent.txt",
	}, nil
}

func CallerDir() (string, error) {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		return "", fmt.Errorf("not found sensitive word file path")
	}
	dir := filepath.Dir(file)

	return dir, nil
}
