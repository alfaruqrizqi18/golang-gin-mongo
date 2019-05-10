package config

import mgo "gopkg.in/mgo.v2"

var Host string = "localhost"
var DatabaseName string = "golang-gin-mongo"

func Connect() (*mgo.Session, error) {
	var session, err = mgo.Dial(Host)
	if err != nil {
		return nil, err
	}
	return session, nil
}
