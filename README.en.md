# udpgo

#### Description
unified develpment platform (golang)

1: integrate log,string function 
2: integrate kafka procuder function
3: integrate influxdb line protocol function
4: string and other helper functions 

#### Software Architecture

#### Installation

go get github.com/chaolihf/udpgo


#### Instructions



#### Contribution


#### Feature

#### Change Log
v0.0.5 add string function: LeftString ,get part string  
v0.0.7 base on gojson(https://github.com/ChengjinWu/gojson.git), add modify json and serial to json string 
v0.0.10 fix json array , GetString(); add future: Keys method
v0.0.11 add log to file and add date formatter
v0.0.12 JsonObject class add newJsonObject and newJsonArray static function ;fix getInt,getLong return error
         kafkaClient's SendMessage function add partition and offset to return result