using Main;
using System;
using System.Diagnostics;
using System.IO;
using System.Net.Http;
using System.Threading;
using System.Xml.Serialization;

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

            Console.WriteLine("\n=== XML ===\n");
            timer.Start();
            SendRequest(@"http://localhost:8080/sendXml");
            timer.Stop();
            requestTime = timer.Elapsed.TotalMilliseconds;

            timer.Reset();

            Thread.Sleep(200);

            timer.Start();
            XmlSerializer serializer = new XmlSerializer(typeof(Person));
            using (StringReader reader = new StringReader(response))
            {
                Person person = (Person)(serializer.Deserialize(reader));
            }
            timer.Stop();
            parseTime = timer.Elapsed.TotalMilliseconds;

            Console.WriteLine("+ The request took: " + requestTime + "ms to complete");
            
            Console.WriteLine("+ The length of the response: " + response.Length + " Bytes");

            Console.WriteLine("+ It took " + parseTime + "ms, to parse the response into an object of type Person (with class XmlSerializer)\n\n");

            Console.WriteLine("+ The response in plain-text: \n\n" + response + "\n");
        }

        private static async void SendRequest(string url)
        {
            response = await hc.GetStringAsync(new Uri(url));
        }
    }
}
