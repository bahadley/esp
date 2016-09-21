# esp (edge stream processor)

### What is esp?

It is a basic stream processing application uses for academic experiments.  It implements a
simple aggregation operator with a sliding window.

### How can I run it? ###
esp can be run as a Linux command (for example):     

``` $ ESP_NODE_TYPE=M esp ```    

with environment variables configured, or just   

``` $ esp ```   

to use default values.   


### Configurable environment variables ###

*ESP_NODE_TYPE*=N (Node type [M=master, S=sink, N=non-master])   
*ESP_ADDR*=127.0.0.1 (IPv4 address used by this esp node)    

*ESP_PORT*=22221 (Port to listen on for ingress sensor tuples.  Only used by master and non-master nodes.)    

*ESP_MASTER_ADDR*=127.0.0.1 (IPv4 address of the master node.  Only used by non-master nodes as a destination for sync tuples.)    

*ESP_SYNC_PORT*=22219 (Port to listen on for ingress sync tuples.)    

*ESP_SINK_ADDR*=127.0.0.1 (IPv4 address for the sink node.  Only used by the master node as a destination for aggregation tuples.)   
*ESP_SINK_PORT*=22220 (Port to listen on for aggregation tuples.)   

*ESP_WINDOW_SIZE*=4 (Maximum number of tuples that can be buffered in the window)   
*ESP_AGGREGATE_SIZE*=2 (Number of tuples that will trigger an egress result tuple)    

*ESP_TRACE*=no (Enable tracing [yes, no])
