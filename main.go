/*
This code is meant to be a starting point for the backend of my Platform-as-a-Service thing
Sort of like REPLIT, its by no means final and very much still untested, literally just wrote it 
so I have something and i dont keep putting it off like all my other projects (we'll see how that goes)

This comment will be removed once it's at least somewhat functioning"
*/

package main

import (
    "fmt"
    "net/http"
    "os"
    "encoding/json"

    "github.com/containers/podman/v5/pkg/bindings"
	"github.com/containers/podman/v5/pkg/bindings/containers"
	"github.com/containers/podman/v5/pkg/bindings/images"
	"github.com/containers/podman/v5/pkg/specgen"
)

// not final
type Json struct {
    ID string
    Lang int 
    Runner int 
    token string
}

func main(){
    
    // maybe look into ways to run it without root for safety reasons later, too tired for now
    conn, err := bindings.NewConnection(context.Background(), "unix:///run/podman/podman.sock")
    if err != nil {
		fmt.Println(err)
	}
    
    // handling the different paths & endpoints 
    http.HandleFunc("/", Root)
    http.HandleFunc("/new", func(w http.ResponseWriter, r *http.Request){New(w, r, conn)})
    http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request){Delete(w, r, conn)})
}

func Root(w http.ResponseWriter, r *http.Request){
    if (r.Method == "GET"){
        http.ServeFile(w, r, MainHtml)
    } else {
        w.WriteHeader(http.StatusMethodNotAllowed)
    }
}

func New(w http.ResponseWriter, r *http.Request){
    if (r.Method == "POST"){
        // Image will be user chosen and matched against array of valid images 
        _, err = images.Pull(conn, "quay.io/libpod/alpine_nginx", nil)
        
        if err != nil {
            fmt.Println(err)
            w.WriteHeader(http.StatusInternalServerError)
        }
        
        s := specgen.NewSpecGenerator("quay.io/libpod/alpine_nginx", false)
        s.Name = "foobar" // randomly generate later
        
        createResponse, err := containers.CreateWithSpec(conn, s, nil)
        if err != nil {
            fmt.Println(err)
            w.WriteHeader(http.StatusInternalServerError)
        }

        if err := containers.Start(conn, createResponse.ID, nil); err != nil {
            fmt.Println(err)
            w.WriteHeader(http.StatusInternalServerError)
        }

        inspectData, err := containers.Inspect(conn, s.Name, new(containers.InspectOptions).WithSize(true))
        if err != nil {
            fmt.Println(err)
            w.WriteHeader(http.StatusInternalServerError)
        }
        // the container ID
        ID := inspectData.ID

        w.Write({"status":"started", "id": ID})
    } else {
        w.WriteHeader(http.StatusMethodNotAllowed)
    }
}

func Delete(w http.ResponseWriter, r *http.Request, conn *bindings.Connection){
    if (r.Method == "POST"){
        var Data Json
        // Container ID or name
        err := json.NewDecoder(r.Body).Decode(&Data)
        if (err != nil) {
            fmt.Println(err)
            w.WriteHeader(http.StatusInternalServerError)
        }
        ContainerID := Data.ID 

        // Stop the container
        timeout := 10 * time.Second
        err = containers.Stop(conn, ContainerID, &timeout)
        if err != nil {
            fmt.Println(err)
            w.WriteHeader(http.StatusInternalServerError)
        } else {
            err = containers.Remove(conn, containerID, &containers.RemoveOptions{Force: true})
            if err != nil {
                fmt.Println(err)
                w.WriteHeader(http.StatusInternalServerError)
            } else {
                w.Write({"status":"deleted"})
            }
        }
    
    } else {
        w.WriteHeader(http.StatusMethodNotAllowed)
    }
}
