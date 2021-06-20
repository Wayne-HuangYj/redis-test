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
