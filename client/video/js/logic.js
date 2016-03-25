var conn = new WebSocket("ws://localhost:8080/streamwsvideo");

var video = document.getElementById('video');

window.MediaSource = window.MediaSource || window.WebKitMediaSource;
if (!!!window.MediaSource) {
  alert('MediaSource API is not available');
}

var mediaSource = new MediaSource()
video.src = window.URL.createObjectURL(mediaSource);

var sourceBuffer = mediaSource.addSourceBuffer('video/mp4; codecs="avc1.42E01E, mp4a.40.2"')
    
conn.onopen = function(event) {
    console.log("connection opened!");
}
    
conn.onmessage = function(event) {
    console.log(event.data)

    var fileReader = new FileReader();
    fileReader.onload = function() {
        sourceBuffer.appendBuffer(this.result);
    };
    fileReader.readAsArrayBuffer(event.data);
}