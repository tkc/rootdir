package rootdir

import (
	"io/ioutil"
	"path/filepath"
	"regexp"
	"runtime"
)

var (
	dirStr = "./"
	r      = regexp.MustCompile(`.mod`)
)

func Root() string {
	root := dirRoop()
	return root
}

func dirRoop() string {
	for i := 0; i < 100; i++ {
		exist := dirSearch("./")
		if exist {
			root := findRootDir(dirStr)
			return root
		}
		dirStr += "../"
	}
	return ""
}

func findRootDir(dirStr string) string {
	_, b, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(b), dirStr)
}

func dirSearch(dir string) bool {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if r.MatchString(file.Name()) {
			return true
		}
	}
	return false
}
