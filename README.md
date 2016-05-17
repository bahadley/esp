# esp


### How can I run it? ###
esp can be run as a Linux command (for example):     

$ ESP_ADDR=192.168.0.250 ESP_PORT=22221 ESP_SINK_ADDR=192.168.0.251 ESP_SINK_PORT=22220 esp    

with environment variables configured, or just   

$ esp   

to use default values.   


### Configurable environment variables ###

*ESP_ADDR*=127.0.0.1 (IPv4 address for the esp node - listens on for ingress sensor tuples)    
*ESP_PORT*=22221 (Port for the esp node - listens on for ingress sensor tuples)    
*ESP_SINK_ADDR*=127.0.0.1 (IPv4 address for the sink node - destination for result stream)   
*ESP_SINK_PORT*=22220 (Port for the sink node - destination for result stream)   
*ESP_SYNC_PORT*=22219 (Port for the esp node - listens on for ingress sync tuples)    

*ESP_WINDOW_SIZE*=4 (Maximum number of tuples that can be buffered in the window)   
*ESP_AGGREGATE_SIZE*=2 (Number of tuples that will trigger an egress result tuple)
