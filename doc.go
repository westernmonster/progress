// Package progress provides methods for monitoring the progress
// of file downloads and uploads.
//
// Usage:
/*
   package main

   import (
     "log"
     "os"
     "net/url"

     "github.com/Bowery/progress"
   )

   var (
      progressClient = progress.New()
   )

   func main() {
     output, _ := os.Create("output.mp4")
     defer output.Close()

     url, _ := url.Parse("http://YOUR_URL")
     progChan, errChan := progressClient.Get(url, nil, output)

     isDownloaded := false
     for !isDownloaded {
       select {
       case status := <-progChan:
          if status.IsFinished() {
            isDownloaded = true
            break
          }
       case err := <-errChan:
         log.Println(err)
       }
     }
   }
*/
package progress
