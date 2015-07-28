"use strict";

$(document).ready(function(){
	var chatSocket = new WebSocket("ws://localhost:8008/ws");

	// TODO: when socket closes, disable chatsend and maybe try to reopen? etc, better error handling
	chatSocket.onopen = function (event) {
		$('#chatsend').prop('disabled', false);		
	};

	chatSocket.onmessage = function (event) {
		$('#chat').prepend(event.data+"\n");
	};

	$('#chatsend').keyup(function(event){
		if(event.keyCode == 13){
			chatSocket.send($('#chatsend').val());
			$('#chatsend').val('');
		}
	});

	// TODO: this should be working but it's not
	$('#chatsend').focus();
});

