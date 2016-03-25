package video

import(    
    "log"
    "time"
    "net/http"
    "io/ioutil"
    "encoding/xml"
    
    "github.com/gorilla/websocket"
)

var streamURLs []string

// upgrader var to manage the websocket
var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool { return true }, // allow javascript client to connect
}

// HandlerVideo handles the connection for the video func
func HandlerVideo(w http.ResponseWriter, r *http.Request) {    
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Print("Error Upgrade: ", err)
        return
    }
    defer conn.Close()
    
    cont := 0
    
    mpdconf, err := processMPDConf("file/tears_of_steel_720p_dash.mpd")
    if err != nil {
        log.Fatal(err)
        return
    }
    
    elementsList := mpdconf.GetSegmentList()
    
    streamURLs = elementsList.GetMediaList()
    
    for { // infinite loop
        streamContent, err := ReadStreamURL("file", streamURLs[cont]) //get URL file extracted from the .mpd
        if err != nil {
            log.Fatal("Stream could not be read: ", err)
            break
        }
        
        if err := conn.WriteMessage(websocket.BinaryMessage, streamContent); err != nil { //send new file to the client
            log.Println("write: ", err)
            break
        }
        
        if(cont > len(streamURLs)) { //eof
            break
        }
        cont++ 
                
        time.Sleep(10 * time.Second)
    }
}

func processMPDConf(loc string) (mpd MPD, err error) {
    content, err := ioutil.ReadFile(loc)
    if err != nil {
        return MPD{}, err
    }
    
    var mpdConf MPD
    xml.Unmarshal(content, &mpdConf)
        
    return mpdConf, nil
}

// GetSegmentList : gets elements list from an MPD configuration file
func (mpd *MPD) GetSegmentList() SegmentList {
    return mpd.Period.Adapt.Rep.SegmentL
}

// GetMediaList : returns a string slice with the different media URLs
func (segList SegmentList) GetMediaList() []string {
    var ret []string
    
    ret = append(ret, segList.InitNode.SourceURL)
    
    for _, value := range segList.SegmentNodes {
        ret = append(ret, value.Media)
    }
    
    return ret
}

// ReadStreamURL : read stream url
func ReadStreamURL(folder, url string) ([]byte, error) {
    ret, err := ioutil.ReadFile(folder + "/" + url)
    if err != nil {
        return nil, err
    }
    return ret, nil
}