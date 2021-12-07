package main
import (
  "github.com/webview/webview"
  "io/ioutil"
)

func main() {
  /* === VARIABLES FOR EASY EDITING === */
  var (
    winWidth      int = 1920
    winHieght     int = 1080
    winTitle      string = "example"
  );
  
  /* === BASIC-WINDOW-CONFIG === */
  w := webview.New(true); defer w.Destroy();
  w.SetTitle(winTitle); w.SetSize(winWidth, winHeight, webview.HintNone);
  
  /* === WINDOW-CONTENT-&-DISPLAY-METHODS === */
  w.Navigate(`data:text/html, <p>example</p>`); // for displaying html (inline)
  w.Navigate("https://example.com"); // for displaying a website
  
  file, _ := ioutil.ReadFile("example.html"); // for displaying html (from local file) (pt.1)
  sF := string(file); // (pt.2)
  w.Navigate(`data:text/html,` + sF); // (pt.3)
  
  w.Run(); // run the application!
}
