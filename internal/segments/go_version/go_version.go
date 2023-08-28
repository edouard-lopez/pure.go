package go_version

import "runtime"

var GetGoVersion = runtime.Version // so we can mock it

func Get() string {
	return GetGoVersion()
}
