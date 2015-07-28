all:
	@cat makefile

old_serve_local:
	ruby -run -ehttpd public -p8008

server: 
	go install && statchatsplat
