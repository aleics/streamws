var conn = new WebSocket("ws://localhost:8080/streamwsimage");

var urlCreator = window.URL || window.webkitURL;  
var img = document.getElementById("img");

var start_rend, end_rend;
var oldstart_rend;

var renderTime;
var onMsgTime;

var imageUrl;
    
conn.onopen = function(event) {
    console.log("connection opened!");
}
    
conn.onmessage = function(event) {    
    start_rend = new Date();
    if(oldstart_rend != null) {
        onMsgTime = start_rend - oldstart_rend;
    }
    
    oldstart_rend = start_rend;    
    
    imageUrl = urlCreator.createObjectURL(event.data);
    img.src = imageUrl
    
    end_rend = new Date();
    
    renderTime = end_rend.getMilliseconds() - start_rend.getMilliseconds();
    document.getElementById("stats-render-time").innerHTML = renderTime.toString();
    document.getElementById("stats-pr-period-time").innerHTML = onMsgTime;
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

function httpSetFreq(value) {
    var data = '{"freq": ' + value + ' }'
    
    var xmlHttp = new XMLHttpRequest();
    xmlHttp.onreadystatechange = function() { 
        if (xmlHttp.readyState == 4 && xmlHttp.status == 200)
            processFreqVal(value);
    }
    xmlHttp.open("POST", "http://localhost:8080/streamwsimage/freq", true); // true for asynchronous
    xmlHttp.setRequestHeader('Content-Type', 'application/json');
        
    xmlHttp.send(data);
}

function processFreqVal(val) {
    document.getElementById("stats-freq").innerHTML = val;
    document.getElementById("set-freq-input").value = val;
}

httpGetFreq();


var p = document.getElementById("set-freq-input");
p.addEventListener("input", function() {
    httpSetFreq(p.value);
}, false); 