all:
	@cat makefile

serve_local:
	ruby -run -ehttpd public -p8008

