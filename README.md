# 介绍

不受内存限制，像管道流水一样处理大数据。**Big Data Filter Streamed.**

![](module.png)

原数据不方便直接展示，这里用随机数代替。

读入数据后，数据一条条的流入 `Flow Line` 中，因此不受内存约束，满足所有 `Restriction` 的才会留下来并写出到文件中。我用 `struct` 实现约束，因此可以自定义添加和删除各种条件，如包含特殊字符删除，大于某一阈值删除等，不必重写处理逻辑。当然还可以补充更多的处理逻辑，如缺失填充等。

## demo

在这个文件夹中是一个示例，供用户理解。

- 先调用 `gene_data.py` 生成随机数
- `mpirun --hostfile hostfile -np 20 python python_filter.py` 会调用 `MPI` 进行并行处理
- `python single.py` 是单线程处理
- `go run flow.go` 是流水处理大数据，远远快于 `MPI`

## main

这个文件夹中是完整的示例。

- `ParseLimitation` 手动添加限制或处理条件即可
- `go run flow.go` 约束较小的时候执行，一闪而过
- `go run multi-flow.go` 约束较多时，多个限制并发执行