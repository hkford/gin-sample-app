#!/bin/bash
if pgrep app
then
echo "app running"
killall app
echo "app  got killed"
else
echo "app is not running/stopped"
fi