run:
	.bin/air

preset:
	mkdir -p .bin
	curl -fLo .bin/air https://raw.githubusercontent.com/cosmtrek/air/master/bin/darwin/air 
	chmod +x .bin/air