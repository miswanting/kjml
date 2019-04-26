@echo off
title KJML Formatter
set GOPATH=%cd%
go run .\go\src\kjml-fmt
pause