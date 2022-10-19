package main

import (
	log "github.com/sirupsen/logrus"
)

var a int

func init() {
	//log.JSONFormatter.DisableTimestamp
	log.SetFormatter(&log.JSONFormatter{})
}
func main() {
	for i := 0; i <= 10; i++ {
		a += i * i
		//log.Infof("i: %v a: %v", i, a)

		log.WithFields(log.Fields{
			"a": a,
			"i": i,
		}).Info("вывод переменных i и a")
	}
}
