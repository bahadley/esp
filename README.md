# esp


### How can I run it? ###
esp can be run as a Linux command (for example):     

``` $ ESP_MASTER=yes esp ```    

with environment variables configured, or just   

``` $ esp ```   

to use default values.   


### Configurable environment variables ###

*ESP_NODE_TYPE*=N (Node type [M=master, S=sink, N=non-master])   
*ESP_ADDR*=127.0.0.1 (IPv4 address used by this esp node)    

*ESP_PORT*=22221 (Port for the esp node - listens on for ingress sensor tuples)    

*ESP_MASTER_ADDR*=127.0.0.1 (IPv4 address of the master esp node.  Only used by non-master nodes as a destination for sync tuples.)    

*ESP_SYNC_PORT*=22219 (Port to listen on for ingress sync tuples.  Only used by master nodes as a destination for sync tuples.)    

*ESP_SINK_ADDR*=127.0.0.1 (IPv4 address for the sink node.  Only used by the master node as a destination for aggregation tuples.)   
*ESP_SINK_PORT*=22220 (Port for the sink node - destination for result stream)   

*ESP_WINDOW_SIZE*=4 (Maximum number of tuples that can be buffered in the window)   
*ESP_AGGREGATE_SIZE*=2 (Number of tuples that will trigger an egress result tuple)    

*ESP_TRACE*=no (Enable tracing [yes, no])
