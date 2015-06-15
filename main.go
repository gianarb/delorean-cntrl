package main

import (
	"encoding/json"
	MQTT "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
	"html/template"
	"io/ioutil"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	t, _ := template.ParseFiles("tpl/home.html")
	t.Execute(w, map[string]string{})
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "POST" {
		data, _ := ioutil.ReadAll(r.Body)
		var b byte

		var opts = MQTT.NewClientOptions()
		opts.AddBroker("tcp://test.mosquitto.org:1883")

		c := MQTT.NewClient(opts)
		t := c.Connect()
		t.Wait()

		c.Publish("go-iot", b, false, data)

		content := []string{"OK"}

		js, _ := json.Marshal(content)
		w.Write(js)
	}

}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/api", apiHandler)
	http.ListenAndServe(":8080", nil)
}
