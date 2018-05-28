all: help

##          ,_---~~~~~----._         
##   _,,_,*^____      _____``*g*\"*, 
##  / __/ /'     ^.  /      \ ^@q   f 
## [  @f | @))    |  | @))   l  0 _/  
##  \`/   \~____ / __ \_____/    \   
##   |           _l__l_           I   
##   }          [______]           I  
##   ]            | | |            |  
##   ]             ~ ~             |  
##   |                            |   
##    |                           |
##
## Gohpers Social Network
## ======================
##
## @author Joan López <joanjan14@gmail.com>
## @description Gophers network example application made with Golang
##
## Available commands are:
##

##  help:		Help
.PHONE : help
help : Makefile
	@sed -n 's/^##//p' $<

##  start-dev:		Start for development
.PHONY : start-dev
start-dev:
	docker-compose -f docker-compose.dev.yaml up -d

##  stop-dev:		Stop for development
.PHONY : stop-dev
stop-dev:
	docker-compose -f docker-compose.dev.yaml stop

##  down-dev:		Down for development
.PHONY : down-dev
down-dev:
	docker-compose -f docker-compose.dev.yaml down
