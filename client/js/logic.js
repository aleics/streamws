var conn = new WebSocket("ws://localhost:8080/streamws");
    
conn.onopen = function(event) {
    console.log("connection opened!");
}
    
conn.onmessage = function(event) {
    console.log(event.data)
}