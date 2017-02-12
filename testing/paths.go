// Package testing is designed to give tests a way to locate any assets without
// messing around with "../../" or other potentially fragile search path
// techniques.
//
// The logic in this package relies on reflection to get the current package path.
// (Current package path in the form: github.com/username/project/directory/packagename)
// Then we rely on the Golang and go-toolchain convention that the sources will exist
// on disk at '${GOPATH}/src/' + 'github.com/username/project/directory/packagename'
//
// Caveats:
// - When building in an IDE, perhaps the GOPATH env var isn't set or propagated correctly.
// - If you do advanced tricks with GOPATH on your system, then you may break the logic here.
// ("advanced tricks" could be something such as:
//      https://web.archive.org/web/20170212200040/https://dmitri.shuralyov.com/blog/18 )
//
// On the bright side, if your GOPATH situation is breaking any logic here, then
// this file gives you one centralized location to just hardcode something that
// works on your system.
package testing

import (
	"os"
	"path"
	"reflect"
	"strings"
)

var envGopath = os.Getenv("GOPATH")

const assetFolder = "sampledata"

// DataAssetFullPath requires the caller to pass in something of some type
// 'typeFromTargetPkg' so that the PkgPath computed here will be the CALLER'S
// package, not the current 'package testing'
func DataAssetFullPath(filebasename string, typeFromTargetPkg interface{}) string {
	// As of Feb 2017, using: go version go1.7.5 darwin/amd64
	// retrieves a string of the form: github.com/username/project/directory/packagename
	packagePath := reflect.TypeOf(typeFromTargetPkg).PkgPath()
	// Using PkgPath trick from: http://stackoverflow.com/a/25263604/10278

	packagePath = "src/" + packagePath

	subparts := strings.Split(packagePath, "/")[0:4]
	// at this point, subparts looks like: [src github.com username project]
	subparts = append(subparts, assetFolder)
	subparts = append(subparts, filebasename)
	// subparts: [src github.com username project assetFolder assetbasename.txt]

	pathback := path.Join(subparts...)

	datapath := path.Join(envGopath, pathback)
	return datapath
}
