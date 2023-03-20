#!/bin/sh

# Check parameter
if [ $# -eq 0 ]; then
    echo "Usage: $0 <name> <facebook> <linkedin>"
    exit 1
fi

# Concatenate the parameters to form a string of "$name at Wed Jul  1 14:00:00 CEST 2020"
parentFolder="$1 at $(date)"

# Create a folder
mkdir -p "$parentFolder"/about_me/personal
mkdir "$parentFolder"/about_me/professional
mkdir "$parentFolder"/my_friends
mkdir "$parentFolder"/my_system_info

# Concatenate facebook username with url
facebookUrl="https://www.facebook.com/$2"

# Create file with facebook url
echo "$facebookUrl" > "$parentFolder"/about_me/personal/facebook.txt

# Concatenate linkedin username with url
linkedinUrl="https://www.linkedin.com/in/$3"

# Create file with linkedin url
echo "$linkedinUrl" > "$parentFolder"/about_me/professional/linkedin.txt

# Create file from the output of the command curl
curl -s https://gist.githubusercontent.com/tegarimansyah/e91f335753ab2c7fb12815779677e914/raw/94864388379fecee450fde26e3e73bfb2bcda194/list%2520of%2520my%2520friends.txt > "$parentFolder"/my_friends/list_of_my_friends.txt

# Create file from the username and output of the command uname
echo "My username: $USER" >> "$parentFolder"/my_system_info/about_this_laptop.txt
uname -a >> "$parentFolder"/my_system_info/about_this_laptop.txt

# Create file from the output of the command ping to google.com
ping google.com -c 3 > "$parentFolder"/my_system_info/internet_connection.txt