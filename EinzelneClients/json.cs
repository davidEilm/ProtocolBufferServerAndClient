using Main;
using System;
using System.Diagnostics;
using System.Net.Http;
using System.Threading;
using Newtonsoft.Json;

namespace ProtoCSharpClient
{
    class Program
    {
        private static HttpClient hc;
        private static string response;

        private static double requestTime;
        private static double parseTime;

        static void Main(string[] args)
        {
            hc = new HttpClient();
            var timer = new Stopwatch();

            Console.WriteLine("\n=== JSON ===\n");
            timer.Start();
            SendRequest(@"http://localhost:8080/sendJson");
            timer.Stop();
            requestTime = timer.Elapsed.TotalMilliseconds;

            timer.Reset();

            Thread.Sleep(200);

            timer.Start();
            Person deserializedProduct = JsonConvert.DeserializeObject<Person>(response);
            timer.Stop();
            parseTime = timer.Elapsed.TotalMilliseconds;

            Console.WriteLine("+ The request took: " + requestTime + "ms to complete");
            
            Console.WriteLine("+ The length of the response: " + response.Length + " Bytes");

            Console.WriteLine("+ It took " + parseTime + "ms, to parse the response into an object of type Person (with class JsonConvert)\n\n");

            Console.WriteLine("+ The response in plain-text: \n\n" + response + "\n");
        }

        private static async void SendRequest(string url)
        {
            response = await hc.GetStringAsync(new Uri(url));
        }
    }
}
