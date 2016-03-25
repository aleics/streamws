var conn = new WebSocket("ws://localhost:8080/streamwsimage");

var urlCreator = window.URL || window.webkitURL;  
var img = document.getElementById("img")
    
conn.onopen = function(event) {
    console.log("connection opened!");
}
    
conn.onmessage = function(event) {
    var imageUrl = urlCreator.createObjectURL(event.data);
    img.src = imageUrl
}

function httpGetFreq() {
    var xmlHttp = new XMLHttpRequest();
    xmlHttp.onreadystatechange = function() { 
        if (xmlHttp.readyState == 4 && xmlHttp.status == 200)
            processFreqVal(xmlHttp.responseText);
    }
    xmlHttp.open("GET", "http://localhost:8080/streamwsimage/freq", true); // true for asynchronous 
    xmlHttp.send(null);
}

function processFreqVal(val) {
    document.getElementById("stats-freq").innerHTML = val
}

httpGetFreq();