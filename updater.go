package main

import (
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	UUID "github.com/tecnologer/uuid"
)

const (
	versionP                   = "_VERSION_"
	uuidP                      = "_PGUID_"
	versionPattern             = "<?define AppVersion    = \"_VERSION_\" ?>"
	guidPattern                = "<?define ProductGUID   = \"_PGUID_\" ?>"
	assemblyVersionPattern     = "AssemblyVersion(\"_VERSION_\")"
	assemblyFileVersionPattern = "AssemblyFileVersion(\"_VERSION_\")"
)

var versionRegex *regexp.Regexp
var fileVersionRegex *regexp.Regexp

//UpdateProduct updates the information in the file Product
func UpdateProduct(productPath, version string) error {
	var fileData []byte
	var err error

	versionRegex, err = regexp.Compile(`<\?define AppVersion\s*= ".*"\s*\?>`)
	if err != nil {
		return err
	}

	fileVersionRegex, err = regexp.Compile(`<\?define ProductGUID\s*= ".*"\s*\?>`)
	if err != nil {
		return err
	}

	fileData, err = readFile(productPath)
	if err != nil {
		return err
	}

	fileData = replaceDataProduct(fileData, version)

	err = writeFile(productPath, fileData)
	if err != nil {
		return err
	}

	return nil
}

//UpdateAssembly updates the information in the file AssemblyInfo
func UpdateAssembly(assemblyPath, version string) error {
	var fileData []byte
	var err error

	versionRegex, err = regexp.Compile(`AssemblyVersion\(.*\)`)
	if err != nil {
		return err
	}

	fileVersionRegex, err = regexp.Compile(`AssemblyFileVersion\(.*\)`)
	if err != nil {
		return err
	}

	fileData, err = readFile(assemblyPath)
	if err != nil {
		return err
	}

	fileData = replaceDataAssembly(fileData, version)

	err = writeFile(assemblyPath, fileData)
	if err != nil {
		return err
	}

	return nil
}

func readFile(filePath string) ([]byte, error) {
	return ioutil.ReadFile(filePath)
}

func writeFile(filePath string, data []byte) error {
	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)

	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteAt(data, 0) // Write at 0 beginning
	if err != nil {
		return err
	}

	return nil
}

func replaceDataProduct(data []byte, version string) []byte {
	installerVersion := strings.Replace(versionPattern, versionP, version, 1)
	fileVersion := strings.Replace(guidPattern, uuidP, UUID.GetUUID(), 1)

	data = versionRegex.ReplaceAll(data, []byte(installerVersion))
	data = fileVersionRegex.ReplaceAll(data, []byte(fileVersion))

	return data
}

func replaceDataAssembly(data []byte, version string) []byte {
	assemblyVersion := strings.Replace(assemblyVersionPattern, versionP, version, 1)
	fileVersion := strings.Replace(assemblyFileVersionPattern, versionP, version, 1)

	data = versionRegex.ReplaceAll(data, []byte(assemblyVersion))
	data = fileVersionRegex.ReplaceAll(data, []byte(fileVersion))

	return data
}
