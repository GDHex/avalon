install:
	@echo "Welcome to Avalon setup"
	@echo "The installation will create 3 folders keys/ signatures/ and data/"
	@echo "Keys folder will contain the keypair files, please store them somewhere safe"
	@echo "Signatures folder will keep the signatures that where produced"
	@echo "Data folder must contain the files that need to be signed or verified"
	mkdir keys
	mkdir signatures
	mkdir data
	go build