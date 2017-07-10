package main

import (
	// "encoding/json"
	// "fmt"
	"testing"
	"reflect"
	// "math/rand"
	// "time"
)
func TestMeli001(t *testing.T){
	if testing.Short() {
        t.Skip("skipping test in short mode.")
    }

}
func Test002(t *testing.T){
	if testing.Short() {
        t.Skip("skipping test in short mode.")
    }
	id :="MLA5725"
	var cosos =[]string {"MLA4711", "MLA6520", "MLA6070", "MLA86360", "MLA3381", "MLA4610", "MLA2227", "MLA86838", "MLA6537","MLA8531", "MLA400928", "MLA1747", "MLA1771", "MLA86080", "MLA377674", "MLA4589", "MLA6177"}
	hijos := Hijos(id)
	if !reflect.DeepEqual(cosos,hijos){
		t.Error(id+" childs have changed or there's an error getting it's childs")
	}

}
