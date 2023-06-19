package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type EmailInfo struct {
	Address, Domain string
}

type Artist struct {
	Name      string
	EmailInfo *EmailInfo
}

func (a *Artist) DeepCopy() *Artist {
	buffer := bytes.Buffer{}
	encoder := gob.NewEncoder(&buffer)
	_ = encoder.Encode(&a)
	result := Artist{}
	decoder := gob.NewDecoder(&buffer)
	_ = decoder.Decode(&result)

	return &result
}

var gmailArtist = Artist{
	Name: "",
	EmailInfo: &EmailInfo{
		Address: "",
		Domain:  "gmail.com",
	},
}

var hotmailArtist = Artist{
	Name: "",
	EmailInfo: &EmailInfo{
		Address: "",
		Domain:  "hotmail.com",
	},
}

func NewArtist(proto *Artist, name, emailAdress string) *Artist {
	newArtist := proto.DeepCopy()
	newArtist.Name = name
	newArtist.EmailInfo.Address = emailAdress

	return newArtist
}

func NewGmailArtist(name, emailAddress string) *Artist {
	return NewArtist(&gmailArtist, name, emailAddress)
}

func NewHotmailArtist(name, emailAddress string) *Artist {
	return NewArtist(&hotmailArtist, name, emailAddress)
}

func main() {
	max := Artist{"Max", &EmailInfo{"izardui00", "gmail.com"}}
	joaquin := max.DeepCopy()
	joaquin.EmailInfo.Domain = "hotmail.com"
	joaquin.EmailInfo.Address = "joaco123"

	fmt.Println(max, max.EmailInfo)
	fmt.Println(joaquin, joaquin.EmailInfo)

	gmail := NewGmailArtist("GMAIL", "gmail123")
	hotmail := NewHotmailArtist("HOTMAIL", "123hotmail")

	fmt.Println(gmail, gmail.EmailInfo)
	fmt.Println(hotmail, hotmail.EmailInfo)
}
