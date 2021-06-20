# redis-test  
redis版本4.0.9，节点8GB内存，4核8线程  
## redis benchmark  
使用redis-benchmark，测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。  
```shell  
    for bytes in 10 20 50 100 200 1024 5120
    do
    redis-benchmark -t get,set -d $bytes
    done
```  
上面的命令默认是50个并发，一共发送10W条数据，最后发现性能好像没有什么变化，请求普遍都是 <=1 milliseconds完成的。  
## info memory  
写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息  , 分析上述不同 value 大小下，平均每个 key 的占用内存空间。  
* 根据题意是，写入的数据量大约是5MB，key的数量就是(5MB / n B)，n从10 20 50 100 200 1024 5120中取一个值。
* 用的是go-redis实现的  

    | value大小(B) | value数量(万) | value平均内存占用(B) | 总占用内存(MB) |
    | - | - | - | - |
    | 10 | 52.4 | 136 | 68 |
    | 20 | 26.2 | 152 | 38 |
    | 50 | 10.4 | 172 | 17.2 |
    | 100 | 5.2 | 220 | 11 |
    | 200 | 2.6 | 332 | 8.3 |
    | 1024 | 0.512 | 1394 | 6.8 |
    | 5120 | 0.1024 | 8312 | 8.1 |
## 总结  
1. 当value数量越多的时候，由于key也是需要占用内存的，redis里面用的是一个RedisObject保存各种key和value的元数据，所以当value数量多的比数量少的占用的总内存要大，一部分都是由于key的元数据造成的。  
2. 当value越大的时候，其平均内存占用内存的增长率（1+平均占用内存/value大小）也就越小，但是value太大的话或许会导致get的时间过长，导致Redis单线程阻塞。  
