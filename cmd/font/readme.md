###目录结构
      font
        -proto　　grpc　协议及服务端代码
            -bought 用户购买字体
            -auth   用户授权字体
            -lease  用户租用字体
           
        
### user_font　redis　结构
    userFontMap
    key : uid 
    value : {"style" : [fontsArray]}