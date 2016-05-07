# esp


### How can I run it? ###
esp can be run as a Linux command (for example):     

$ ESP_ADDR=192.168.0.250 ESP_PORT=22221 ESP_SINK_ADDR=192.168.0.251 ESP_SINK_PORT=22220 esp    

with environment variables configured, or just   

$ esp   

to use default values.   


### Configurable environment variables ###

*ESP_ADDR*=127.0.0.1 (IPv4 address for the esp node)    
*ESP_PORT*=22221 (port for the esp node)
*ESP_SINK_ADDR*=127.0.0.1 (IPv4 address for the sink node)
*ESP_SINK_PORT*=22220 (port for the sink node)
