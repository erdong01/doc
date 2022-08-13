对map 加锁优化  可以用一个 goruntine 通过chan通信 对map进行操作 避免锁的竞争 达到优化
