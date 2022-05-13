using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Net.Http;
using System.Net;

namespace calculations
{
    internal class Http
    {
        void TakeHttp(Uri requestUri,  Dictionary<string, string> headers, string methodType = "GET")
        {
            WebRequest request = WebRequest.Create(requestUri); // "http://127.0.0.1:8080/api/super_user/ping"
            foreach(var header in headers)
            {
                request.Headers.Add(header.Key, header.Value);
            }
            //request.Headers.Add("access", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwZXJtaXNzaW9ucyI6ImdvZCBwZXJtaXNzaW9ucyJ9.i-MW_1RhmPQ-6sJErFwjUvw1NhYpCvR6HQ9TaskJAIc");
            request.Method = methodType;
            //
            WebResponse response = request.GetResponse();
            using (Stream stream = response.GetResponseStream())
            {
                using (StreamReader reader = new StreamReader(stream))
                {
                    string line = "";
                    while ((line = reader.ReadLine()) != null)
                    {
                        Console.WriteLine(line);
                    }
                }
            }
            response.Close();
            //Console.WriteLine("Запрос выполнен");
            //Console.Read();
        }
    }
}
