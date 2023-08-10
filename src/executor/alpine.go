package executor

import (
	"stubber/helpers"
	"stubber/templates"
)

func stubAlpine(softwarename string) error {
	var err error
	placeholders := map[string]string{
		"{{ MAINTAINER }}":      helpers.Maintainer,
		"{{ PACKAGER }}":        helpers.Packager,
		"{{ SOFTWARE NAME }}":   softwarename,
		"{{ PACKAGE VERSION }}": helpers.VersionNumber,
		"{{ PACKAGE RELEASE }}": helpers.ReleaseNumber,
		"{{ DESCRIPTION }}":     helpers.Description,
		"{{ ARCHITECTURE }}":    helpers.Arch,
		"{{ BINARY NAME }}":     helpers.BinaryName,
	}
	if err = templates.ProcessEmbeddedAsset("apk/APKBUILD", "__alpine/APKBUILD", placeholders); err == nil {
		err = templates.ProcessEmbeddedAsset("apk/Makefile", "__alpine/Makefile", placeholders)
	}
	return err
}

//func apkbuild(softwarename string) error {
//	placeholders := map[string]string{
//		"{{ MAINTAINER }}":      helpers.Maintainer,
//		"{{ PACKAGER }}":        helpers.Packager,
//		"{{ PACKAGE NAME }}":    softwarename,
//		"{{ PACKAGE VERSION }}": helpers.VersionNumber,
//		"{{ PACKAGE RELEASE }}": helpers.ReleaseNumber,
//		"{{ DESCRIPTION }}":     helpers.Description,
//		"{{ ARCHITECTURE }}":    helpers.Arch,
//		"{{ BINARY NAME }}":     helpers.BinaryName,
//		"{{ GO VERSION }}":      helpers.GoVersion,
//	}
//	//if err := templates.ProcessEmbeddedAsset("apk/APKBUILD", "__alpine/APKBUILD", placeholders); err != nil {
//	//	return err
//	//}
//	return templates.ProcessEmbeddedAsset("apk/APKBUILD", "__alpine/APKBUILD", placeholders)
//}
//
//func makefile(softwarename string) error {
//	placeholders := map[string]string{
//		"{{ GO VERSION }}":   helpers.GoVersion,
//		"{{ DESCRIPTION }}":  helpers.Description,
//		"{{ ARCHITECTURE }}": helpers.Arch,
//		"{{ BINARY NAME }}":  helpers.BinaryName,
//	}
//	//if err := templates.ProcessEmbeddedAsset("apk/APKBUILD", "__alpine/APKBUILD", placeholders); err != nil {
//	//	return err
//	//}
//	return templates.ProcessEmbeddedAsset("apk/Makefile", "__alpine/Makefile", placeholders)
//}
