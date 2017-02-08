#!/usr/bin/env bash
sout="'#transcode{acodec=mp3,vcodec=dummy}:standard{access=file,mux=raw,dst="$2"}'"

eval 'vlc -I dummy "$1" --sout=$sout vlc://quit'