install:
	@echo "---------------------------------------------------------------------------------"
	@echo "                           Welcome to Avalon setup"
	@echo "---------------------------------------------------------------------------------"
	@echo "The installation will create 3 folders keys/ signatures/ and data/"
	@echo "Keys folder will contain the keypair files, please store them somewhere safe"
	@echo "Signatures folder will keep the signatures that where produced"
	@echo "Data folder must contain the files that need to be signed or verified"
	@echo "Installing..."
	@- if [ ! -d "keys" ]; then mkdir keys; fi
	@- if [ ! -d "signatures" ]; then mkdir signatures; fi
	@- if [ ! -d "data" ]; then mkdir data; fi
	@- go build
	@echo "---------------------------------------------------------------------------------"
	@echo "Done!"
	@echo "---------------------------------------------------------------------------------"
	@echo "                  Installation completed, welcome to Avalon!"
	@echo "---------------------------------------------------------------------------------"