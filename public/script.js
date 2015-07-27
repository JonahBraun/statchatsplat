var chatSocket = new WebSocket("ws://localhost:8008/ws");

chatSocket.onopen = function (event) {
	chatSocket.send("Example Test sent to the server");
};

chatSocket.onmessage = function (event) {
	console.log(event.data);
};

