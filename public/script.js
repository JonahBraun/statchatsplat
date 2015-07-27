"use strict";

$(document).ready(function(){
	var chatSocket = new WebSocket("ws://localhost:8008/ws");

	chatSocket.onopen = function (event) {
		chatSocket.send("Example Test sent to the server");
	};

	chatSocket.onmessage = function (event) {
		$('#chat').append(event.data+"\n");
	};

	$('#chatsend').keyup(function(event){
		if(event.keyCode == 13){
			chatSocket.send($('#chatsend').value);
		}
	});
});

