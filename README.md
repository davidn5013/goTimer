# GoTimer

Two step chronograph for the terminal that does not need to spawn a shell

Why? Well in cmd.exe it does not always work spawning a new shell for timing and get the right output and does not give me ms, this is my solution

This tool can probably be used on linux but I do not see why not just use the time command

## Usage
        
First run of GoTimer starts the timer next run of GoTimer will stop the timer. One argument -name of the timer so you can start multi timer


## Install

```
go install github.com/davidn5013/goTimer@latest
``` 

## Todo

- [X] Less delay by storing in env? No the env does not stay. GoTimer spawn it own shell. Failed test saved in stash.

/ David
