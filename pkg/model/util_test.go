package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUtil_Test_GetImageRepoAndVersion_With_Valid_Keycloak_Image(t *testing.T) {
	// given
	image := "quay.io/keycloak/keycloak:7.0.1"

	// when
	imageRepo, imageMajor, imageMinor, imagePatch := GetImageRepoAndVersion(image)

	// then
	assert.Equal(t, imageRepo, "quay.io/keycloak/keycloak")
	assert.Equal(t, imageMajor, "7")
	assert.Equal(t, imageMinor, "0")
	assert.Equal(t, imagePatch, "1")
}

func TestUtil_Test_GetImageRepoAndVersion_With_Valid_RHSSO_Image(t *testing.T) {
	// given
	image := "registry.access.redhat.com/redhat-sso-7/sso73-openshift:1.0-15"

	// when
	imageRepo, imageMajor, imageMinor, imagePatch := GetImageRepoAndVersion(image)

	// then
	assert.Equal(t, imageRepo, "registry.access.redhat.com/redhat-sso-7/sso73-openshift")
	assert.Equal(t, imageMajor, "1")
	assert.Equal(t, imageMinor, "0-15")
	assert.Equal(t, imagePatch, "")
}

func TestUtil_Test_GetImageRepoAndVersion_With_Valid_RHSSO_CVE_Image(t *testing.T) {
	// given
	image := "registry.access.redhat.com/redhat-sso-7/sso73-openshift:1.0-15.1567588155"

	// when
	imageRepo, imageMajor, imageMinor, imagePatch := GetImageRepoAndVersion(image)

	// then
	assert.Equal(t, imageRepo, "registry.access.redhat.com/redhat-sso-7/sso73-openshift")
	assert.Equal(t, imageMajor, "1")
	assert.Equal(t, imageMinor, "0-15")
	assert.Equal(t, imagePatch, "1567588155")
}

func TestUtil_Test_GetImageRepoAndVersion_With_No_Image(t *testing.T) {
	// given
	image := ""

	// when
	imageRepo, imageMajor, imageMinor, imagePatch := GetImageRepoAndVersion(image)

	// then
	assert.Equal(t, imageRepo, "")
	assert.Equal(t, imageMajor, "")
	assert.Equal(t, imageMinor, "")
	assert.Equal(t, imagePatch, "")
}
