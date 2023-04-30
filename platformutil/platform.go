package platformutil

import (
	"fmt"
	"runtime"
)

/*
Returns the local OS, Architecture and Variant that is
compatible with Docker supported architectures
*/
func GetLocalOsArchitectureVariant() (string, string, string) {
	architecture := runtime.GOARCH
	fmt.Println(architecture)
	operatingSystem := runtime.GOOS
	switch operatingSystem {
	case "windows":
		fmt.Println("Windows")
	case "darwin":
		fmt.Println("MAC operating system")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s.\n", operatingSystem)
	}

	return "", "", ""
}
