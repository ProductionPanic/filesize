#!/bin/bash
# Install the filesize cli tool

# variables
THISDIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PARENTDIR="$(dirname "$THISDIR")"

# Build the app
COMMAND="go build -o filesize $PARENTDIR"
echo "Building the app with: $COMMAND"
OUTPUT=$($COMMAND)
echo $OUTPUT
echo "Done building the app"

# Move the app to /usr/local/bin
echo "Moving the app to /usr/local/bin"
cmd="mv $PARENTDIR/filesize /usr/local/bin"
output=$($cmd)
echo "DOne moving the app to /usr/local/bin"

# prompt the user if they want to remove the source code
echo "Do you want to remove the source code? [y/n]"
for (( ; ; )); do
	read -n 1 -s -r -p "" input
	if [[ $input = "y" ]]; then
		echo "Removing the source code"
		cmd="rm -rf $PARENTDIR"
		output=$($cmd)
		echo "Done removing the source code"
		break
	elif [[ $input = "n" ]]; then
		echo "Not removing the source code"
		break
	fi
done
