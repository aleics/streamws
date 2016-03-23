package main

import (
    "flag"
    "log"
    "time"
    "net/http"
    "io/ioutil"
    "encoding/xml"
    
    "github.com/gorilla/websocket"
)

// address flag
var addr = flag.String("addr", "localhost:8080", "http service address")

var streamURLs []string

// upgrader var to manage the websocket
var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool { return true }, // allow javascript client to connect
}

// handler handles the connection
func handler(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Print("Error Upgrade: ", err)
        return
    }
    defer conn.Close()
    
    cont := 0
    
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
        
        cont++        
        if(cont > len(streamURLs)) { //eof
            break
        }
        
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

// main function
func main() {
    
    mpdconf, err := processMPDConf("file/tears_of_steel_720p_dash.mpd")
    if err != nil {
        log.Fatal(err)
    }
    
    elementsList := mpdconf.GetSegmentList()
    
    streamURLs = elementsList.GetMediaList()
    
    flag.Parse()
    log.SetFlags(0)
    http.HandleFunc("/streamws", handler)
    log.Fatal(http.ListenAndServe(*addr, nil))
}


// MPD struct
type MPD struct {
    ProgramInfo ProgramInformation `xml:"ProgramInformation"`
    Period Period `xml:"Period"`
}
// ProgramInformation struct
type ProgramInformation struct {
    MoreInfoURL string `xml:"moreInformationURL,attr"`
    Title string `xml:"Title"`
}
// Period struct
type Period struct {
    Duration string `xml:"duration,attr"`
    Adapt AdaptationSet `xml:"AdaptationSet"`
}
// AdaptationSet struct
type AdaptationSet struct {
    SegmentAlignment string `xml:"segmentAlignment,attr"`
    MaxWidth int `xml:"maxWidth,attr"`
    MaxHeight int `xml:"maxHeight,attr"`
    MaxFrameRate int `xml:"maxFrameRate,attr"`
    Par string `xml:"par,attr"`
    Lang string `xml:"lang,attr"`
    
    ContComp []ContentComponent `xml:"ContentComponent"`
    Rep Representation `xml:"Representation"`
}
// ContentComponent struct
type ContentComponent struct {
    ID int `xml:"id,attr"`
    ContentType string `xml:",contentType"`    
}
// Representation struct
type Representation struct {
    ID int `xml:"id,attr"`
    MimeType string `xml:"mimeType,attr"`
    Codecs string `xml:"codecs,attr"`
    Width int `xml:"width,attr"`
    Height int `xml:"height,attr"`
    FrameRate int `xml:"frameRate,attr"`
    Sar string `xml:"sar,attr"`
    AudioSamplingRate int `xml:"audioSamplingRate,attr"`
    StartWithSAP int `xml:"startWithSAP,attr"`
    Bandwidth int `xml:"bandwidth,attr"`
    
    AudioChanConf AudioChannelConf `xml:"AudioChannelConfiguration"`
    SegmentL SegmentList `xml:"SegmentList"`    
}
// AudioChannelConf struct
type AudioChannelConf struct {
    SchemeURI string `xml:"schemeIdUri,attr"`
    Value int `xml:"value,attr"`
}
// SegmentList struct
type SegmentList struct {
    InitNode Initialization `xml:"Initialization"`
    SegmentNodes []SegmentURL `xml:"SegmentURL"`
    Timescale int `xml:"timescale,attr"`
    Duration int `xml:"duration,attr"`
}
// Initialization struct
type Initialization struct {
    SourceURL string `xml:"sourceURL,attr"`
}
// SegmentURL struct
type SegmentURL struct {
    Media string `xml:"media,attr"`
}