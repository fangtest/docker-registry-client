package main

import (
	"encoding/json"
	"github.com/docker/distribution/manifest/schema1"
	"github.com/fangtest/docker-registry-client/registry"
)

func main() {
	url := "http://192.168.12.144:5000"
	username := "" // anonymous
	password := "" // anonymous
	hub, err := registry.New(url, username, password)
	if err != nil {

	}
	repositories, err := hub.Repositories()
	if err != nil {

	}
	println(repositories)

	tags, err := hub.Tags("app-center-ws")
	if err != nil {

	}
	println(tags)

	manifest, err := hub.Manifest("app-center-ws", "latest-origin-master-5a105bafbc-20233007153043-26")
	if err != nil {

	}
	println(manifest)

	time := buildCreateTime(manifest)
	println(time)

	manifestv2, sha256, err := hub.ManifestV2("app-center-ws", "latest-origin-master-5a105bafbc-20233007153043-26")
	if err != nil {

	}
	println(manifestv2)
	println(sha256)

	digest, err := hub.ManifestDigest("app-center-ws", "latest-origin-master-5a105bafbc-20233007153043-26")
	if err != nil {

	}
	println(digest)

	metadata, err := hub.BlobMetadata("app-center-ws", "latest-origin-master-5a105bafbc-20233007153043-26")
	if err != nil {

	}
	println(metadata.Size)

}

type ImageManifest struct {
	Created string `json:"created"`
}

func buildCreateTime(manifest *schema1.SignedManifest) string {
	compatibility := manifest.History[0].V1Compatibility
	imageManifest := ImageManifest{}

	if err := json.Unmarshal([]byte(compatibility), &imageManifest); err != nil {
		return ""
	}
	return imageManifest.Created
}
