.SILENT:

.PHONY: nop
nop:
	echo "make what?"

.PHONY: count
count: 
	countula -gitignore -extensions "go,html,css,js,json,yaml,md,gohtml" > lines