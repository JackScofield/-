## PromQL

在time-series中的每一个点称为一个样本 sample

样本：

* 指标 metric，metric name和描述当前样本特征的labelsets
* 时间戳 timestamp 一个精准到ms的时间戳
* 样本值 value， 一个float64的浮点型数据表示样本值

指标metric格式

```
<metric name>{<label name>=<label value>,...}
eg:
api_http_requests_total{method="POST", handler="/messages"}
```

### metric 的四种类型

```
# HELP 注释
# TYPE <metric name> type
<metric name>{<label name>=<label value>,...} value
```

#### Counter 只增不减的计数器

一般定义counter类型的metric名称时推荐使用_total作为后缀



#### Gauge：可增可减的仪表盘



#### Histogram和Summary

主要记录用于统计和分析样本的分布情况

一般这两类数据都很长

区别在于Summary是直接在客户端计算了数据分布的分位数情况。而Histogram的分位数计算需要通过histogram_quantile(φ float, b instant-vector)函数进行计算。其中φ（0<φ<1）表示需要计算的分位数，如果需要计算中位数φ取值为0.5，以此类推即可。



### 初识PromQL

prometheus通过metrics name 和对应的一组label set来唯一定义一条时间序列。可以通过对这个label set进行过滤，聚合，统计然后得到新的时间序列



![image-20200825102639665](C:\Users\Mr Scofield\AppData\Roaming\Typora\typora-user-images\image-20200825102639665.png)







## Exporter 自定义

Exporter返回的样本数据由三个部分组成：样本的一般注释信息（HELP)，样本的类型注释信息（TYPE）和样本本身。

```
# HELP <metrics_name> <doc_string>
# TYPE <metrics_name> <metrics_type>
metrics_name {<label_name="label_value">,...} metric_value
```

