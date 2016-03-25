package image

import (
    "log"
    "time"
    "strconv"
    "net/http"
    "io/ioutil" 
    "encoding/json"

    "github.com/gorilla/websocket"
)

// upgrader var to manage the websocket
var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool { return true }, // allow javascript client to connect
}

var waitTime = 42 //41.66667 Milliseconds ->  1/24 fps

type freqReq struct {
    Val int `json:"freq"`
}

// Freq handler
func Freq(w http.ResponseWriter, r *http.Request) {
    if origin := r.Header.Get("Origin"); origin != "" {
        w.Header().Set("Access-Control-Allow-Origin", origin)
        w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
        w.Header().Set("Access-Control-Allow-Headers",
            "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
    }
     
    if r.Method == http.MethodPost {
        decoder := json.NewDecoder(r.Body)
        
        var newFreq freqReq   
        err := decoder.Decode(&newFreq)
        if err != nil {
            w.Write([]byte("Body format invalid"))
        }
        
        
        if newFreq.Val > 0 {
            waitTime = newFreq.Val
        }
        
        return
    } else if r.Method == http.MethodGet {
        w.Write([]byte(strconv.Itoa(waitTime)))
        return
    }
    w.Write([]byte("Method '" + r.Method + "' not allowed."))
}

// HandlerImage handles the connection for the image func
func HandlerImage(w http.ResponseWriter, r *http.Request) {    
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Print("Error Upgrade: ", err)
        return
    }
    defer conn.Close()
    
    cont := 1
    
    for {
        currentImgFile := "frames" + strconv.Itoa(cont) + ".jpg"        
        imgContent, err := ReadImage("frames", currentImgFile)
        if err != nil {
            log.Println("EOF")
            cont = 1
        }
        
        if err := conn.WriteMessage(websocket.BinaryMessage, imgContent); err != nil { //send new file to the client
            log.Println("write: ", err)
            break
        }
        cont++
        time.Sleep(time.Duration(waitTime) * time.Millisecond)
    }
}

// ReadImage func
func ReadImage(folder, image string) ([]byte, error) {
    ret, err := ioutil.ReadFile(folder + "/" + image)
    if err != nil {
        return nil, err
    }
    return ret, nil
}