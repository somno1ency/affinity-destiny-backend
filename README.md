注意:
1.看目录结构就知道,repo下的目录结构与hander/logic的目录结构并不一致,我们想当然地认为应该要是一致的,因为hander最终就是操作db嘛,但实际确实要根据情况来,不能强要求一致,有以下几个原因:
    1.1.handler是与前端/别的服务进行交互的,注意的是别的服务需要这个服务提供的逻辑接口,这些逻辑接口要操作的表可能是多个相关的,并不会每个逻辑刚好操作一张表,所以handler这里的每个逻辑使用的表的数量不是固定的,不应该跟表强绑定
    1.2.handler可能还会有一些逻辑是完全不需要操作表的

2.handler中使用了xhttp "github.com/zeromicro/x/http"扩展插件实现正常数据与错误数据的code&msg格式的json返回,所以最好更改一下handler的生成模板,免得每次生成的模板代码都要改,使用了这种方式后,在user.go中的全局ErrorHandler处理函数就不用了,这种方式只能使错误保持code&msg格式,而正常数据还是裸的,既然用code&msg,就都用code&msg吧