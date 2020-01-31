# finleap

## Prerequsite
- Please install docker 

## Git
- sudo git clone https://github.com/AbishSowrirajan/finleap.git

## Commad to proceed 
- cd finleap 

## steps to spin up the server 
- sudo docker-compose up 
  
## Testing is covered automatically when spin up the server using docker-compose up 
 - you can check the test results on terminal when docker image is built (to manually check the testing give command (go test ./... -v --cover)
  
## Please install Curl in linux machine or test from Postman 

## sample Curl for Creating City
- curl -XPOST http://localhost:8080/cities -d  name=berlin  -d longitude=12.34 -d latitude=14.55

## sample Curl for Updating City
- curl -XPATCH http://localhost:8080/cities/1 -d name=fraNCE -d longitude=12.34 -d latitude=14.55

## sample curl for Deleting City 
- curl -XDELETE http://localhost:8080/cities/1

## sample curl to insert the temperature for city 
- curl -XPOST http://localhost:8080/temperatures -d city_id=1  -d max=67 -d min=89

## sample curl to get the forcast for the city  
- curl  GET http://localhost:8080/forcasts/1

## sample curl to insert the webhooks for the city  
- curl -XPOST http://localhost:8080/webhooks -d city_id=3 -d callback_url=https://localhost:8080/temperature7

## sample curl to Delete  the webhooks for the city  
- curl -XDELETE  http://localhost:8080/webhooks/7


## to stop the server 
- sudo docker-compose down 










 
