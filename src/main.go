package main

import (
	"github.com/webview/webview"
	"io/ioutil"
)

func main() {
	// variables for easy editing
	var (
		winWidth      int = 1920
    	winHieght     int = 1080
    	winTitle      string = "example"
  	);
  
  	// basic window config
  	w := webview.New(true); defer w.Destroy();
  	w.SetTitle(winTitle); w.SetSize(winWidth, winHeight, webview.HintNone);
  
  	// methods fot content rendering
  	w.Navigate(`data:text/html, <p>example</p>`); // from inline
  	w.Navigate("https://example.com"); // from URL

	file, _ := ioutil.ReadFile("example.html"); // from local file (pt.1)
 	w.Navigate(`data:text/html,` + string(file)); // (pt.2)
  
  	w.Run(); // run the application!
}
