# About
This code is written on the GO(golang) language and could be used to receive Zacks Rank https://www.zacks.com/stocks/zacks-rank

# Current functions
 - store history of Zacks Rank changes
 - support in-memory and DB level caches
 - get current Zacks Rank of a stock
 - get the history of Zacks Rank of a stock
 - watch for changes

# Example
 - Use docker container to start a service: docker run -p 8080:8080 docker.io/iadolgov/zacks:latest
 - To receive Zacks Rank for the ticker "M": open the page http://localhost:8080/M
 - To receive Zacks Rank History for the ticker "M": open the page http://localhost:8080/M/history
 
# RoadMap
 - add DB cache implementation(mongodb)
 - notify about Zack Rank changes
 
 # Pull Requests and Questions
  - Pull Requests: https://github.com/IAD/zacks/pulls
  - Questions: https://github.com/IAD/zacks/issues