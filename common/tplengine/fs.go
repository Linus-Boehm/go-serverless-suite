/*
This is a file with just the FileSystem, as linters do not like this usage, so we have just to disable linting for this file
*/
package tplengine

import "embed"

//go:embed manifests/*
var DefaultManifests embed.FS
