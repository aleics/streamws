package video

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