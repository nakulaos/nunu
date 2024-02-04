# nunu介绍
本应用暂时采取单体应用开发，后续会针对进行微服务，分布式的相关开发

# 技术栈
## 后端
+ gin
+ gorm
+ mysql
+ redis
+ 监控
  + sentry
  + prometheus
+ 日志
  + filebeat+logstash+kafka+elasticSearch+go-stash+kibana(配置失败)
  + zincSearch
  + meiliSearch
  + file


**待使用**
暂时考虑在支付方面使用以下技术栈，待补充
+ kafka
+ go-zero
+ grpc
+ dtm 
+ kakfa


## 前端/桌面端/移动端
+ vue
+ electron
+ ionic
+ h5

## 功能



## 设计问题

1. 错误处理和日志处理如何解决？国际化msg如何处理？ 

    Code 前端条件码显示

    Msg 前端信息显示，得考虑国际化
    
    Code 具体错误代码
    
    log 后端日志显示，不考虑国际化，采取中文解决，日志在业务逻辑层记录，信息尽可能更全面
    
    数据逻辑层不处理任何错误，直接原地抛出
    
    业务逻辑层处理并记录错误在日志
    如果微服务，则在rpc层处理日志
    
    api/rpc层返回错误,并处理国际化，这里记录错误码

2. 三层架构设计？
    
    数据逻辑层 接口+实现
    
    业务逻辑层 接口+实现，这种写法？
    每一层service都有对应的repo
    ```
   type SettingService struct{}

    type ISettingService interface {
          GetSettingInfo() (*dto.SettingInfo, error)
     }
    func NewISettingService() ISettingService {
         return &SettingService{}
    }
      
   // 然后自定义生成变量
   ```
    api 层 如何考虑？

    
3. 微服务和单体应用如何兼容
    
    
    
    
    


