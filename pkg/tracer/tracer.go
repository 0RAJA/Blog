package tracer

import (
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"time"
)

//链路追踪系统
//一个Trace(跟踪)代表了一个事务或者流程在系统中的执行过程
//一个Span(跨度)代表每个事务中每个工作单元,通常多个span将会组成一个完成的Trace
//一个SpanContext(跨度上下文)代表一个事务的相关追踪信息,不同的Span会根据OpenTracing规范封装不同的属性,包含操作名称,
//	--开始时间和结束时间,标签信息,日志信息,上下文信息等

//Jaeger 分布式链路追踪系统

func NewJaegerTracer(serviceName, agentHostPort string) (opentracing.Tracer, io.Closer, error) {
	//client 的配置项,设置应用的基本信息,
	cfg := &config.Configuration{
		ServiceName: serviceName,
		Sampler: &config.SamplerConfig{ //固定采样,对所有数据都采样
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{ //是否启用LoggingReproter,刷新缓冲区的频率,上报的Agent地址,
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort:  agentHostPort,
		},
	}
	tracer, closer, err := cfg.NewTracer() //根据配置初始化Tracer对象
	if err != nil {
		return nil, nil, err
	}
	opentracing.SetGlobalTracer(tracer) //设置全局的Tracer对象
	return tracer, closer, nil
}
