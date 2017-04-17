// +build integration

package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/docker/docker/pkg/namesgenerator"
	"github.com/docker/go-plugins-helpers/volume"
)

func TestCapabilities(t *testing.T) {
	// Get the list of capabilities the driver supports.
	t.Logf("POST /VolumeDriver.Capabilities")

	// Create json for request - local variable json masks the global symbol json referring to the JSON module
	t.Logf("json.Marshal(volume.Request{})")
	jsn, err := json.Marshal(volume.Request{})
	if err != nil {
		t.Fatal(err)
	}

	// Create request to server
	body := bytes.NewBuffer(jsn)
	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://plugin:8080/VolumeDriver.Capabilities", body)
	if err != nil {
		t.Fatal(err)
	}

	// Fetch Request
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal("Failed to connect to server! ", err)
	}

	if resp == nil {
		t.Fatal("resp is nil!")
	}
	if resp.Body == nil {
		t.Fatal("resp.Body is nil!")
	}
	defer resp.Body.Close()

	var r volume.Response

	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		t.Fatal(err)
	}

	if r.Capabilities.Scope != "local" {
		t.Fatal("Scope should be local!")
	}
}

func TestList(t *testing.T) {
	// Get the list of volumes registered with the plugin.
	t.Logf("POST /VolumeDriver.List")

	// Create json for request - local variable json masks the global symbol json referring to the JSON module
	t.Logf("json.Marshal(volume.Request{})")
	jsn, err := json.Marshal(volume.Request{})
	if err != nil {
		t.Fatal(err)
	}

	// Create request to server
	body := bytes.NewBuffer(jsn)
	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://plugin:8080/VolumeDriver.List", body)
	if err != nil {
		t.Fatal(err)
	}

	// Fetch request
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal("Failed to connect to server! ", err)
	}

	if resp == nil {
		t.Fatal("resp is nil!")
	}
	if resp.Body == nil {
		t.Fatal("resp.Body is nil!")
	}
	defer resp.Body.Close()

	var r volume.Response

	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		t.Fatal(err)
	}

	if r.Err != "" {
		t.Fatal("Failed to list volumes!", r.Err)
	}

	if len(r.Volumes) != 0 {
		t.Fatal("List of volumes should be 0! Actual:", len(r.Volumes))
	}
}

func TestCreateRemove(t *testing.T) {
	/**************************CREATE********************************/
	// Create a new volume with name and options.
	t.Logf("POST /VolumeDriver.Create")

	// Create json for request - local variable json masks the global symbol json referring to the JSON module
	t.Logf("json.Marshal(volume.Request{Name: \"\", Options: map[string]string{}})")
	jsn, err := json.Marshal(volume.Request{
		Name:    "",
		Options: map[string]string{},
	})
	if err != nil {
		t.Fatal(err)
	}

	// Create request to server
	body := bytes.NewBuffer(jsn)
	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://plugin:8080/VolumeDriver.Create", body)
	if err != nil {
		t.Fatal(err)
	}

	// Fetch request
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp == nil {
		t.Fatal("resp is nil!")
	}
	if resp.Body == nil {
		t.Fatal("resp.Body is nil!")
	}

	var r volume.Response

	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		t.Fatal(err)
	}

	if r.Err == "" {
		t.Fatal("Attempting to create a volume with an empty volume_name should cause an error")
	}

	resp.Body.Close()

	/**************************CREATE********************************/
	// Get a Docker approved random name
	volumeName := namesgenerator.GetRandomName(0)

	// Create a new volume with name and options.
	t.Logf("POST /VolumeDriver.Create")

	// Create json for request - local variable json masks the global symbol json referring to the JSON module
	t.Logf("json.Marshal(volume.Request{Name: %s, Options: map[string]string{}})", volumeName)
	jsn, err = json.Marshal(volume.Request{
		Name:    volumeName,
		Options: map[string]string{},
	})

	if err != nil {
		t.Fatal(err)
	}

	// Create request to server
	body = bytes.NewBuffer(jsn)
	client = &http.Client{}
	req, err = http.NewRequest("POST", "http://plugin:8080/VolumeDriver.Create", body)
	if err != nil {
		t.Fatal(err)
	}

	// Fetch request
	req.Header.Add("Content-Type", "application/json")
	resp, err = client.Do(req)
	if err != nil {
		t.Fatal("Failed to connect to server! ", err)
	}

	if resp == nil {
		t.Fatal("resp is nil!")
	}
	if resp.Body == nil {
		t.Fatal("resp.Body is nil!")
	}

	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		t.Fatal(err)
	}

	if r.Err != "" {
		t.Fatal("Failed to create volume! {Name: ", volumeName, "} ", r.Err)
	}

	resp.Body.Close()

	/**************************CREATE********************************/
	// Create a new volume with name and options.
	t.Logf("POST /VolumeDriver.Create")

	// Create json for request - local variable json masks the global symbol json referring to the JSON module
	t.Logf("json.Marshal(volume.Request{Name: %s, Options: map[string]string{}})", volumeName)
	jsn, err = json.Marshal(volume.Request{
		Name:    volumeName,
		Options: map[string]string{},
	})

	if err != nil {
		t.Fatal(err)
	}

	// Create request to server
	body = bytes.NewBuffer(jsn)
	client = &http.Client{}
	req, err = http.NewRequest("POST", "http://plugin:8080/VolumeDriver.Create", body)
	if err != nil {
		t.Fatal(err)
	}

	// Fetch request
	req.Header.Add("Content-Type", "application/json")
	resp, err = client.Do(req)
	if err != nil {
		t.Fatal("Failed to connect to server! ", err)
	}

	if resp == nil {
		t.Fatal("resp is nil!")
	}
	if resp.Body == nil {
		t.Fatal("resp.Body is nil!")
	}

	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		t.Fatal(err)
	}

	if r.Err == "" {
		t.Fatal("Attempting to create a volume without a unique volume_name should cause an error")
	}

	resp.Body.Close()

	/**************************REMOVE********************************/
	// Remove (Delete) a paricular volume.
	t.Logf("POST /VolumeDriver.Remove")

	// Create json for request - local variable json masks the global symbol json referring to the JSON module
	t.Logf("json.Marshal(volume.Request{Name: %s})", volumeName)
	jsn, err = json.Marshal(volume.Request{
		Name: volumeName,
	})

	if err != nil {
		t.Fatal(err)
	}

	// Create request to server
	body = bytes.NewBuffer(jsn)
	client = &http.Client{}
	req, err = http.NewRequest("POST", "http://plugin:8080/VolumeDriver.Remove", body)
	if err != nil {
		t.Fatal(err)
	}

	// Fetch request
	req.Header.Add("Content-Type", "application/json")
	resp, err = client.Do(req)
	if err != nil {
		t.Fatal("Failed to connect to server! ", err)
	}

	if resp == nil {
		t.Fatal("resp is nil!")
	}
	if resp.Body == nil {
		t.Fatal("resp.Body is nil!")
	}

	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		t.Fatal(err)
	}

	if r.Err != "" {
		t.Fatal("Failed to delete volume! {Name: ", volumeName, "} ", r.Err)
	}

	resp.Body.Close()

	/**************************REMOVE********************************/
	t.Logf("POST /VolumeDriver.Remove")

	// Create json for request - local variable json masks the global symbol json referring to the JSON module
	t.Logf("json.Marshal(volume.Request{Name: %s})", volumeName)
	jsn, err = json.Marshal(volume.Request{
		Name: volumeName,
	})
	if err != nil {
		t.Fatal(err)
	}

	// Create request for server
	body = bytes.NewBuffer(jsn)
	client = &http.Client{}
	req, err = http.NewRequest("POST", "http://plugin:8080/VolumeDriver.Remove", body)
	if err != nil {
		t.Fatal(err)
	}

	// Fetch request
	req.Header.Add("Content-Type", "application/json")
	resp, err = client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp == nil {
		t.Fatal("resp is nil!")
	}
	if resp.Body == nil {
		t.Fatal("resp.Body is nil!")
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		t.Fatal(err)
	}

	if r.Err == "" {
		t.Fatal("Error should occur when removing a nonexistant volume")
	}
}

func TestCreateListRemoveList(t *testing.T) {
	// Get a Docker approved random name
	volumeName := namesgenerator.GetRandomName(0)

	/**************************CREATE********************************/
	// Create a new volume with name and options.
	t.Logf("POST /VolumeDriver.Create")

	// Create json for request - local variable json masks the global symbol json referring to the JSON module
	t.Logf("json.Marshal(volume.Request{Name: %s, Options: map[string]string{}})", volumeName)
	jsn, err := json.Marshal(volume.Request{
		Name:    volumeName,
		Options: map[string]string{},
	})
	if err != nil {
		t.Fatal(err)
	}

	// Create request to server
	body := bytes.NewBuffer(jsn)
	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://plugin:8080/VolumeDriver.Create", body)
	if err != nil {
		t.Fatal(err)
	}

	// Fetch request
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal("Failed to connect to server! ", err)
	}

	if resp == nil {
		t.Fatal("resp is nil!")
	}
	if resp.Body == nil {
		t.Fatal("resp.Body is nil!")
	}

	var r volume.Response

	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		t.Fatal(err)
	}

	if r.Err != "" {
		t.Fatal("Failed to create volume! {Name: ", volumeName, "} ", r.Err)
	}

	resp.Body.Close()

	/**************************LIST********************************/
	// Get the list of volumes registered with the plugin.
	t.Logf("POST /VolumeDriver.List")

	// Create json for request - local variable json masks the global symbol json referring to the JSON module
	t.Logf("json.Marshal(volume.Request{})")
	jsn, err = json.Marshal(volume.Request{})
	if err != nil {
		t.Fatal(err)
	}

	// Create request to server
	body = bytes.NewBuffer(jsn)
	client = &http.Client{}
	req, err = http.NewRequest("POST", "http://plugin:8080/VolumeDriver.List", body)
	if err != nil {
		t.Fatal(err)
	}

	// Fetch request
	req.Header.Add("Content-Type", "application/json")
	resp, err = client.Do(req)
	if err != nil {
		t.Fatal("Failed to connect to server! ", err)
	}

	if resp == nil {
		t.Fatal("resp is nil!")
	}
	if resp.Body == nil {
		t.Fatal("resp.Body is nil!")
	}

	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		t.Fatal(err)
	}

	if r.Err != "" {
		t.Fatal("Failed to list volumes!", r.Err)
	}

	if len(r.Volumes) != 1 {
		t.Fatal("List of volumes should be 1! Actual:", len(r.Volumes))
	}

	resp.Body.Close()

	/**************************REMOVE********************************/
	// Remove (Delete) a paricular volume.
	t.Logf("POST /VolumeDriver.Remove")

	// Create json for request - local variable json masks the global symbol json referring to the JSON module
	t.Logf("json.Marshal(volume.Request{Name: %s})", volumeName)
	jsn, err = json.Marshal(volume.Request{
		Name: volumeName,
	})

	if err != nil {
		t.Fatal(err)
	}

	// Create request to server
	body = bytes.NewBuffer(jsn)
	client = &http.Client{}
	req, err = http.NewRequest("POST", "http://plugin:8080/VolumeDriver.Remove", body)
	if err != nil {
		t.Fatal(err)
	}

	// Fetch request
	req.Header.Add("Content-Type", "application/json")
	resp, err = client.Do(req)
	if err != nil {
		t.Fatal("Failed to connect to server! ", err)
	}

	if resp == nil {
		t.Fatal("resp is nil!")
	}
	if resp.Body == nil {
		t.Fatal("resp.Body is nil!")
	}

	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		t.Fatal(err)
	}

	if r.Err != "" {
		t.Fatal("Failed to delete volume! {Name: ", volumeName, "} ", r.Err)
	}

	resp.Body.Close()

	/**************************LIST********************************/
	// Get the list of volumes registered with the plugin.
	t.Logf("POST /VolumeDriver.List")

	// Create json for request - local variable json masks the global symbol json referring to the JSON module
	t.Logf("json.Marshal(volume.Request{})")
	jsn, err = json.Marshal(volume.Request{})
	if err != nil {
		t.Fatal(err)
	}

	// Create request to server
	body = bytes.NewBuffer(jsn)
	client = &http.Client{}
	req, err = http.NewRequest("POST", "http://plugin:8080/VolumeDriver.List", body)
	if err != nil {
		t.Fatal(err)
	}

	// Fetch request
	req.Header.Add("Content-Type", "application/json")
	resp, err = client.Do(req)
	if err != nil {
		t.Fatal("Failed to connect to server! ", err)
	}

	if resp == nil {
		t.Fatal("resp is nil!")
	}
	if resp.Body == nil {
		t.Fatal("resp.Body is nil!")
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		t.Fatal(err)
	}

	if r.Err != "" {
		t.Fatal("Failed to list volumes!", r.Err)
	}

	if len(r.Volumes) != 0 {
		t.Fatal("List of volumes should be 0! Actual:", len(r.Volumes))
	}
}

func TestGetCreateGetRemove(t *testing.T) {
	// Get a Docker approved random name
	volumeName := namesgenerator.GetRandomName(0)

	/**************************GET********************************/
	// Get info relating to a paricular volume.
	t.Logf("POST /VolumeDriver.Get")

	// Create json for request - local variable json masks the global symbol json referring to the JSON module
	t.Logf("json.Marshal(volume.Request{Name: %s})", volumeName)
	jsn, err := json.Marshal(volume.Request{
		Name: volumeName,
	})

	if err != nil {
		t.Fatal(err)
	}

	// Create request to server
	body := bytes.NewBuffer(jsn)
	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://plugin:8080/VolumeDriver.Get", body)
	if err != nil {
		t.Fatal(err)
	}

	// Fetch request
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal("Failed to connect to server! ", err)
	}

	if resp == nil {
		t.Fatal("resp is nil!")
	}
	if resp.Body == nil {
		t.Fatal("resp.Body is nil!")
	}

	var r volume.Response

	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		t.Fatal(err)
	}

	if r.Err == "" {
		t.Fatal("Volume does not exist so there should be an error, but there is not.")
	}

	resp.Body.Close()

	/**************************CREATE********************************/
	// Create a new volume with name and options.
	t.Logf("POST /VolumeDriver.Create")

	// Create json for request - local variable json masks the global symbol json referring to the JSON module
	t.Logf("json.Marshal(volume.Request{Name: %s, Options: map[string]string{}})", volumeName)
	jsn, err = json.Marshal(volume.Request{
		Name:    volumeName,
		Options: map[string]string{},
	})

	if err != nil {
		t.Fatal(err)
	}

	// Create request to server
	body = bytes.NewBuffer(jsn)
	client = &http.Client{}
	req, err = http.NewRequest("POST", "http://plugin:8080/VolumeDriver.Create", body)
	if err != nil {
		t.Fatal(err)
	}

	// Fetch request
	req.Header.Add("Content-Type", "application/json")
	resp, err = client.Do(req)
	if err != nil {
		t.Fatal("Failed to connect to server! ", err)
	}

	if resp == nil {
		t.Fatal("resp is nil!")
	}
	if resp.Body == nil {
		t.Fatal("resp.Body is nil!")
	}

	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		t.Fatal(err)
	}

	if r.Err != "" {
		t.Fatal("Failed to create volume! {Name: ", volumeName, "} ", r.Err)
	}

	resp.Body.Close()

	/**************************GET********************************/
	// Get info relating to a paricular volume.
	t.Logf("POST /VolumeDriver.Get")

	// Create json for request - local variable json masks the global symbol json referring to the JSON module
	t.Logf("json.Marshal(volume.Request{Name: %s})", volumeName)
	jsn, err = json.Marshal(volume.Request{
		Name: volumeName,
	})

	if err != nil {
		t.Fatal(err)
	}

	// Create request to server
	body = bytes.NewBuffer(jsn)
	client = &http.Client{}
	req, err = http.NewRequest("POST", "http://plugin:8080/VolumeDriver.Get", body)
	if err != nil {
		t.Fatal(err)
	}

	// Fetch request
	req.Header.Add("Content-Type", "application/json")
	resp, err = client.Do(req)
	if err != nil {
		t.Fatal("Failed to connect to server! ", err)
	}

	if resp == nil {
		t.Fatal("resp is nil!")
	}
	if resp.Body == nil {
		t.Fatal("resp.Body is nil!")
	}

	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		t.Fatal(err)
	}

	if r.Err != "" {
		t.Fatal("Failed to get info for volume: ", volumeName, r.Err)
	}

	if r.Volume.Name != volumeName {
		t.Fatal("Expected: ", volumeName, "Actual:", r.Volume.Name, r.Err)
	}

	resp.Body.Close()

	/**************************REMOVE********************************/
	// Remove (Delete) a paricular volume.
	t.Logf("POST /VolumeDriver.Remove")

	// Create json for request - local variable json masks the global symbol json referring to the JSON module
	t.Logf("json.Marshal(volume.Request{Name: %s})", volumeName)
	jsn, err = json.Marshal(volume.Request{
		Name: volumeName,
	})

	if err != nil {
		t.Fatal(err)
	}

	// Create request to server
	body = bytes.NewBuffer(jsn)
	client = &http.Client{}
	req, err = http.NewRequest("POST", "http://plugin:8080/VolumeDriver.Remove", body)
	if err != nil {
		t.Fatal(err)
	}

	// Fetch request
	req.Header.Add("Content-Type", "application/json")
	resp, err = client.Do(req)
	if err != nil {
		t.Fatal("Failed to connect to server! ", err)
	}

	if resp == nil {
		t.Fatal("resp is nil!")
	}
	if resp.Body == nil {
		t.Fatal("resp.Body is nil!")
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		t.Fatal(err)
	}

	if r.Err != "" {
		t.Fatal("Failed to delete volume! {Name: ", volumeName, "} ", r.Err)
	}
}
