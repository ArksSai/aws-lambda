# aws-lambda
Learning how to use aws lambda

![image](https://user-images.githubusercontent.com/51477633/130739769-b53d2e29-f0d3-4c6f-a786-5296d171b26e.png)


* connect_to_rds
* 使用 lambda
  *  引用 github.com/aws/aws-lambda-go/lambda
  *  在main 中 呼叫 lambda()

* rds proxy
  *  通常為連線到 rds proxy，再連線到 rds
  *  secret manager 中申請rds secret
  *  新增一個role、policy使其有權限能連線至rds

* 連線至 rds proxy
  * rds 的連線參數通常另寫在一個 .env 檔，再使用 os.Getenv() 的方式取得

* 上傳至lmabda
  * 打包成 zip 檔上傳。若有 .env 檔，也要記得加到解縮檔中

* 利用 api gateway
  *  通常使用 restful api 的形式來觸發
  *  在 serverless.yml 檔中的 functions 中規劃事件觸發方法

ref：https://qiita.com/maika_kamada/items/6eb6a40c17b4b8995acb
