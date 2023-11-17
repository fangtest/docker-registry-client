package main

import "github.com/fangtest/docker-registry-client/registry"

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

	tags, err := hub.Tags("emc")
	if err != nil {

	}
	println(tags)

	manifest, err := hub.Manifest("emc", "latest")
	if err != nil {

	}
	println(manifest)

	manifestv2, err := hub.ManifestV2("emc", "latest")
	if err != nil {

	}
	println(manifestv2)

	digest, err := hub.ManifestDigest("emc", "latest")
	if err != nil {

	}
	println(digest)

	//hub.
	//
	//metadata, err := hub.HasBlob("emc", "659271073e878d41687ca602743bed7733625ff41f009b4bfb62f1f7d5908bf4")
	//if err != nil {
	//
	//}
	//metadata.Descriptor()
}
