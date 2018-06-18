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

##
##	
## Development commands:
## =====================
##

##  start-dev:	Start for development
.PHONY : start-dev
start-dev:
	docker-compose -f docker-compose.dev.yaml up -d
	
##  logs-dev:	Logs for development
.PHONY : logs-dev
logs-dev:
	docker-compose -f docker-compose.dev.yaml logs

##  stop-dev:	Stop for development
.PHONY : stop-dev
stop-dev:
	docker-compose -f docker-compose.dev.yaml stop

##  down-dev:	Down for development
.PHONY : down-dev
down-dev:
	docker-compose -f docker-compose.dev.yaml down

##
##      
## Production commands:
## =====================
##

##  start:	Start for production
.PHONY : start
start:
	docker-compose -f docker-compose.prod.yaml up -d --build
	
##  logs:		Logs for production
.PHONY : logs
logs:
	docker-compose -f docker-compose.prod.yaml logs

##  stop:		Stop for production
.PHONY : stop
stop:
	docker-compose -f docker-compose.prod.yaml stop

##  down:		Down for production
.PHONY : down
down:
	docker-compose -f docker-compose.prod.yaml down
    
##
##
