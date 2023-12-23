package main

import (
   "fmt"
   "strings"
   "net/http"
   "io/ioutil"
)

func main() {

   url := "https://server.zk2016.com/outside/api/extend/v8transform.do"
   method := "POST"

   payload := strings.NewReader(`{"AppNumber":"100001","AppPassword":"5ZBA3iQqEdeT7EvoTx5IYw==","Url":"api/SignIn/SignIn","Data":{"UserId":"749f43ab-b31f-47b4-ac63-062571f308ff","BranchId":"b8beb0be-1f75-42de-bee6-43c0866926bd"},"BranchId":"b8beb0be-1f75-42de-bee6-43c0866926bd"}`)

   client := &http.Client {
   }
   req, err := http.NewRequest(method, url, payload)

   if err != nil {
      fmt.Println(err)
      return
   }
   req.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 14; M2012K11AC Build/UKQ1.230804.001; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/111.0.5563.116 Mobile Safari/537.36 XWEB/5235 MMWEBSDK/20230504 MMWEBID/2827 MicroMessenger/8.0.37.2368(0x28002548) WeChat/arm64 Weixin GPVersion/1 NetType/WIFI Language/zh_CN ABI/arm64 MiniProgramEnv/android")
   req.Header.Add("Content-Type", "application/json")
   req.Header.Add("Accept", "*/*")
   req.Header.Add("Host", "server.zk2016.com")
   req.Header.Add("Connection", "keep-alive")
   req.Header.Add("Cookie", "acw_tc=0bdd26da17033497451263906ef626900d43ae318b0e06fdf91fadd551d575; SERVERID=be70f122d3e66c1054c8d131bd6a8738|1703349745|1703349745")

   res, err := client.Do(req)
   if err != nil {
      fmt.Println(err)
      return
   }
   defer res.Body.Close()

   body, err := ioutil.ReadAll(res.Body)
   if err != nil {
      fmt.Println(err)
      return
   }
   fmt.Println(string(body))
}
