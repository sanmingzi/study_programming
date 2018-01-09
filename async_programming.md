# Async programming

[Reference](http://www.cnblogs.com/teroy/p/4015461.html)

http://www.cnblogs.com/jesse2013/p/async-and-await.html

https://msdn.microsoft.com/en-us/library/hh191443(v=vs.110).aspx

## Demo

```
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Web;
using System.Threading;
using System.Threading.Tasks;

namespace TestApp
{
    class Program
    {
        static void Main(string[] args)
        {
            Console.WriteLine("-------主线程启动-------");
            Task<int> task = GetLengthAsync();
            Console.WriteLine("Main方法做其他事情");
            Console.WriteLine("Task返回的值" + task.Result);
            Console.WriteLine("-------主线程结束-------");
        }

        static async Task<int> GetLengthAsync()
        {
            Console.WriteLine("GetLengthAsync Start");
            string str = await GetStringAsync();
            Console.WriteLine("GetLengthAsync End");
            return str.Length;
        }

        static Task<string> GetStringAsync()
        {
            return Task<string>.Run(() => { Thread.Sleep(2000); return "finished"; });
        }
    }
}
```

- "-------主线程启动-------"
- "GetLengthAsync Start"
- "Main方法做其他事情"
- "GetLengthAsync End"
- "Task返回的值" + task.Result
- "-------主线程结束-------"

```
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Web;
using System.Threading;
using System.Threading.Tasks;
using System.Net;

namespace TestApp
{
    class Program
    {
        static void Main(string[] args)
        {
            Console.WriteLine("开始获取博客园首页字符数量");
            Task<int> task1 = CountCharsAsync("http://www.cnblogs.com");
            Console.WriteLine("开始获取百度首页字符数量");
            Task<int> task2 = CountCharsAsync("http://www.baidu.com");

            Console.WriteLine("Main方法中做其他事情");

            Console.WriteLine("博客园:" + task1.Result);
            Console.WriteLine("百度:" + task2.Result);
        }

        static async Task<int> CountCharsAsync(string url)
        {
            WebClient wc = new WebClient();
            string result = await wc.DownloadStringTaskAsync(new Uri(url));
            return result.Length;
        }
    }
}
```

## Task

## async/await

## 锁

## Semaphore