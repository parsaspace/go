package parsaspace

import (
	"errors"
	"log"
	"time"

	"github.com/levigross/grequests"
)

// Parsaspace : parsaspace struct
type Parsaspace struct {
	APIKey  string
	APIBase string
}

// parsaspace routes
const (
	FilesRoute         string = "/v1/files/list"
	RemoveRoute        string = "/v1/files/remove"
	RenameRoute        string = "/v1/files/rename"
	MoveRoute          string = "/v1/files/move"
	CopyRoute          string = "/v1/files/Copy"
	CreateDirRoute     string = "/v1/files/Createfolder"
	UploadRoute        string = "/v1/files/upload"
	RemoteUploadRoute  string = "/v1/remote/new"
	RemoteUploadStatus string = "/v1/remote/status"
)

// Response : response model
type Response struct {
	Result  string `json:"result"`
	List    []List `json:"list"`
	Message string
}

// List : response list item
type List struct {
	Name         string    `json:"Name"`
	IsFolder     string    `json:"IsFolder"`
	Size         int       `json:"Size"`
	LastModified time.Time `json:"LastModified"`
}

// http client

// NewClient : create new parsaspace client
func NewClient(key string) Parsaspace {
	return Parsaspace{APIKey: "Bearer " + key, APIBase: "https://api.parsaspace.com/"}
}

// Files : List of your files
func (p Parsaspace) Files(domain, path string) ([]List, error) {
	var ls Response

	// This will upload the file as a multipart mime request
	resp, err := grequests.Post(p.APIBase+FilesRoute,
		&grequests.RequestOptions{
			Headers: map[string]string{"Authorization": p.APIKey},
			Data:    map[string]string{"domain": domain, "path": path},
		})

	if err != nil {
		log.Println("Unable to make request", resp.Error)
		return ls.List, resp.Error
	}

	if resp.Ok != true {
		log.Println("Request did not return OK")
		return ls.List, errors.New(("Request did not return OK"))
	}

	resp.JSON(&ls)

	return ls.List, nil
}

// Remove : remove file from your website
func (p Parsaspace) Remove(domain, path string) error {
	var dt Response

	// This will upload the file as a multipart mime request
	resp, err := grequests.Post(p.APIBase+FilesRoute,
		&grequests.RequestOptions{
			Headers: map[string]string{"Authorization": p.APIKey},
			Data:    map[string]string{"domain": domain, "path": path},
		})

	if err != nil {
		log.Println("Unable to make request", resp.Error)
		return resp.Error
	}

	resp.JSON(&dt)

	if dt.Result != "success" {
		log.Println("Request did not return OK")
		return errors.New(("Request did not return OK"))
	}

	return nil
}

// Rename : rename file in your website
func (p Parsaspace) Rename(domain, source, destination string) error {
	var dt Response

	// This will upload the file as a multipart mime request
	resp, err := grequests.Post(p.APIBase+FilesRoute,
		&grequests.RequestOptions{
			Headers: map[string]string{"Authorization": p.APIKey},
			Data:    map[string]string{"domain": domain, "source": source, "destination": destination},
		})

	if err != nil {
		log.Println("Unable to make request", resp.Error)
		return resp.Error
	}

	resp.JSON(&dt)

	if dt.Result != "success" {
		log.Println("Request did not return OK")
		return errors.New(("Request did not return OK"))
	}

	return nil
}

// Move : move file
func (p Parsaspace) Move(domain, source, destination string) error {
	var dt Response

	// This will upload the file as a multipart mime request
	resp, err := grequests.Post(p.APIBase+FilesRoute,
		&grequests.RequestOptions{
			Headers: map[string]string{"Authorization": p.APIKey},
			Data:    map[string]string{"domain": domain, "source": source, "destination": destination},
		})

	if err != nil {
		log.Println("Unable to make request", resp.Error)
		return resp.Error
	}

	resp.JSON(&dt)

	if dt.Result != "success" {
		log.Println("Request did not return OK")
		return errors.New(("Request did not return OK"))
	}

	return nil
}

// Copy : copy file
func (p Parsaspace) Copy(domain, source, destination string) error {
	var dt Response

	// This will upload the file as a multipart mime request
	resp, err := grequests.Post(p.APIBase+FilesRoute,
		&grequests.RequestOptions{
			Headers: map[string]string{"Authorization": p.APIKey},
			Data:    map[string]string{"domain": domain, "source": source, "destination": destination},
		})

	if err != nil {
		log.Println("Unable to make request", resp.Error)
		return resp.Error
	}

	resp.JSON(&dt)

	if dt.Result != "success" {
		log.Println("Request did not return OK")
		return errors.New(("Request did not return OK"))
	}

	return nil
}

// NewDir : create new directory
func (p Parsaspace) NewDir(domain, path string) error {
	var dt Response

	// This will upload the file as a multipart mime request
	resp, err := grequests.Post(p.APIBase+FilesRoute,
		&grequests.RequestOptions{
			Headers: map[string]string{"Authorization": p.APIKey},
			Data:    map[string]string{"domain": domain, "path": path},
		})

	if err != nil {
		log.Println("Unable to make request", resp.Error)
		return resp.Error
	}

	resp.JSON(&dt)

	if dt.Result != "success" {
		log.Println("Request did not return OK")
		return errors.New(("Request did not return OK"))
	}

	return nil
}

// Upload : upload file
func (p Parsaspace) Upload(domain, path, file string) error {
	fd, err := grequests.FileUploadFromDisk(file)

	if err != nil {
		log.Println("Unable to open file: ", err)
		return err
	}

	// This will upload the file as a multipart mime request
	resp, err := grequests.Post(p.APIBase+UploadRoute,
		&grequests.RequestOptions{
			Files:   fd,
			Headers: map[string]string{"Authorization": p.APIKey},
			Data:    map[string]string{"domain": domain, "path": path},
		})

	if err != nil {
		log.Println("Unable to make request", resp.Error)
		return resp.Error
	}

	if resp.Ok != true {
		log.Println("Request did not return OK")
		return errors.New(("Request did not return OK"))
	}

	return nil
}

// RemoteUpload : remote upload
func (p Parsaspace) RemoteUpload(domain, path, url, filename, checkid string) error {
	var dt Response

	// This will upload the file as a multipart mime request
	resp, err := grequests.Post(p.APIBase+FilesRoute,
		&grequests.RequestOptions{
			Headers: map[string]string{"Authorization": p.APIKey},
			Data:    map[string]string{"domain": domain, "path": path, "url": url, "filename": filename, "checkid": checkid},
		})

	if err != nil {
		log.Println("Unable to make request", resp.Error)
		return resp.Error
	}

	resp.JSON(&dt)

	if dt.Result != "success" {
		log.Println("Request did not return OK")
		return errors.New(("Request did not return OK"))
	}

	return nil
}

// RemoteUploadStatus : remoteupload status
func (p Parsaspace) RemoteUploadStatus(checkid string) error {
	var dt Response

	// This will upload the file as a multipart mime request
	resp, err := grequests.Post(p.APIBase+FilesRoute,
		&grequests.RequestOptions{
			Headers: map[string]string{"Authorization": p.APIKey},
			Data:    map[string]string{"checkid": checkid},
		})

	if err != nil {
		log.Println("Unable to make request", resp.Error)
		return resp.Error
	}

	resp.JSON(&dt)

	if dt.Result != "success" {
		log.Println("Request did not return OK")
		return errors.New(("Request did not return OK"))
	}

	return nil
}
