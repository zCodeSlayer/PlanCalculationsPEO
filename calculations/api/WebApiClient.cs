using calculations.models;
using calculations.packed_models;
using Newtonsoft.Json;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Net;
using System.Text;
using System.Text.Json.Serialization;
using System.Threading.Tasks;


namespace calculations.api
{
    public class WebApiClient
    {
        public User Authorization(User user, string HttpAdress="http://127.0.0.1:8080")
        {
            WebRequest request = WebRequest.Create($"{HttpAdress}/api/auth"); //  HttpAdress = http://127.0.0.1:8080
            request.Method = "POST";
            request.ContentType = "application/json";

            using (var streamWriter = new StreamWriter(request.GetRequestStream()))
            {
                string json = JsonConvert.SerializeObject(new { login=user.Login, password=user.Password});
                streamWriter.Write(json);
            }

            var httpResponse = (HttpWebResponse)request.GetResponse();
            using (var streamReader = new StreamReader(httpResponse.GetResponseStream()))
            {
                var result = streamReader.ReadToEnd();
                User? res = JsonConvert.DeserializeObject<User>(result);
                user.Id = res.Id;
                user.Token = res.Token;
            }

            return user;
        }
        
        public void CheckPingServerOrRulesAuthUser(User user = null, string groupName = null, string HttpAdress = "http://127.0.0.1:8080")
        {
            string massage = "server ping massage: ";
            WebRequest request = WebRequest.Create($"{HttpAdress}/api/ping");
            if ((user != null) && (groupName != null))
            {
                massage = "user ping massage: ";
                request = WebRequest.Create($"{HttpAdress}/api/{groupName}/ping");
                request.Method = "GET";
                request.Headers.Add("access", user.Token);
            }
            
            HttpWebResponse response = (HttpWebResponse)request.GetResponse();

            using (Stream stream = response.GetResponseStream())
            {
                using (StreamReader reader = new StreamReader(stream))
                {
                    string line = "";
                    while ((line = reader.ReadLine()) != null)
                    {
                        massage += line;
                    }
                }
            }
            response.Close();
            Console.WriteLine(massage);
        }

        public void CreateUser(User user, string HttpAdress = "http://127.0.0.1:8080")
        {
            WebRequest request = WebRequest.Create($"{HttpAdress}/api/roles/create_user");
            request.Method = "POST";
            request.Headers.Add("access", user.Token);
            request.ContentType = "application/json";

            using (var streamWriter = new StreamWriter(request.GetRequestStream()))
            {
                string json = JsonConvert.SerializeObject(new {login = user.Login, password = user.Password, id_role = "5"});
                streamWriter.Write(json);
            }

            //todo дописать метод
        }


        public List<PackedProduct> GetAllProducts(string HttpAdress = "http://127.0.0.1:8080")
        {
            WebRequest request = WebRequest.Create($"{HttpAdress}/api/products/get"); //  HttpAdress = http://127.0.0.1:8080
            request.Method = "GET";
            request.ContentType = "application/json";
            List<PackedProduct> res = new List<PackedProduct>();

            var httpResponse = (HttpWebResponse)request.GetResponse();
            using (var streamReader = new StreamReader(httpResponse.GetResponseStream()))
            {
                var result = streamReader.ReadToEnd();
                res = JsonConvert.DeserializeObject<List<PackedProduct>>(result);
            }

            foreach (var item in res)
                Console.WriteLine(item.Name);

            return res;

        }
        public List<CostItem> GetAllCostitems(string HttpAdress = "http://127.0.0.1:8080")
        {
            WebRequest request = WebRequest.Create($"{HttpAdress}/api/cost_items/get"); //  HttpAdress = http://127.0.0.1:8080
            request.Method = "GET";
            request.ContentType = "application/json";
            List<CostItem> res = new List<CostItem>();

            var httpResponse = (HttpWebResponse)request.GetResponse();
            using (var streamReader = new StreamReader(httpResponse.GetResponseStream()))
            {
                var result = streamReader.ReadToEnd();
                res = JsonConvert.DeserializeObject<List<CostItem>>(result);
            }

            foreach (var item in res) 
                Console.WriteLine(item.Name);

            return res;
        }

        public List<PackedCalculation> GetAllCalculations(string HttpAdress = "http://127.0.0.1:8080")
        {
            WebRequest request = WebRequest.Create($"{HttpAdress}/api/calculations/get");
            request.Method = "GET";
            request.ContentType = "application/json";
            List<PackedCalculation> res = new List<PackedCalculation>();

            var httpResponse = (HttpWebResponse)request.GetResponse();
            using (var streamReader = new StreamReader(httpResponse.GetResponseStream()))
            {
                var result = streamReader.ReadToEnd();
                Console.WriteLine(result);
                res = JsonConvert.DeserializeObject<List<PackedCalculation>>(result);
            }

            //foreach (var item in res)
                //Console.WriteLine(item.Id);

            return res;
        }

        private record rCostItem(string id, string cost);
        public void CreateCalculation(PackedCalculation calculation, string HttpAdress = "http://127.0.0.1:8080")
        {
            WebRequest request = WebRequest.Create($"{HttpAdress}/api/calculations/create"); //  HttpAdress = http://127.0.0.1:8080
            request.Method = "POST";
            request.ContentType = "application/json";
            request.Headers["access"] = @"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwZXJtaXNzaW9ucyI6ImdvZCBwZXJtaXNzaW9ucyJ9.i-MW_1RhmPQ-6sJErFwjUvw1NhYpCvR6HQ9TaskJAIc";

            using (var streamWriter = new StreamWriter(request.GetRequestStream()))
            {
                List <rCostItem> costItems = new List<rCostItem>();
                foreach (var item in calculation.Cost_items)
                    costItems.Add(new rCostItem(Convert.ToString(item.Id), Convert.ToString(item.Cost)));

                string json = JsonConvert.SerializeObject(new { 
                    start_date=calculation.Start_date,
                    end_date=calculation.End_date, 
                    product= new { id=Convert.ToString(calculation.Product.Id) }, 
                    cost_items = costItems
                });
                streamWriter.Write(json);
                Console.WriteLine(json);
            }

            PackedCalculation res = new PackedCalculation();
            var httpResponse = (HttpWebResponse)request.GetResponse();
            using (var streamReader = new StreamReader(httpResponse.GetResponseStream()))
            {
                var result = streamReader.ReadToEnd();
                Console.WriteLine(result);
                res = JsonConvert.DeserializeObject<PackedCalculation>(result);
            }
            Console.WriteLine();
        }

        public void DeleteCalculation(User user, int calculationId, string HttpAdress = "http://127.0.0.1:8080")
        {
            WebRequest request = WebRequest.Create($"{HttpAdress}/api/calculations/delete/{ calculationId }");
            request.Method = "DELETE";
            request.ContentType = "application/json";
            request.Headers["access"] = user.Token;

            var httpResponse = (HttpWebResponse)request.GetResponse();
            using (var streamReader = new StreamReader(httpResponse.GetResponseStream()))
            {
                var result = streamReader.ReadToEnd();
                Console.WriteLine(result);
            }

        }
        public PackedCalculation GetCalculation(int calculationId, string HttpAdress = "http://127.0.0.1:8080")
        {
            WebRequest request = WebRequest.Create($"{HttpAdress}/api/calculations/get/{ calculationId }");
            request.Method = "GET";
            request.ContentType = "application/json";
            PackedCalculation res = new PackedCalculation();

            var httpResponse = (HttpWebResponse)request.GetResponse();
            using (var streamReader = new StreamReader(httpResponse.GetResponseStream()))
            {
                var result = streamReader.ReadToEnd();
                Console.WriteLine(result);
                res = JsonConvert.DeserializeObject<PackedCalculation>(result);
            }
            Console.WriteLine();
            return res;
        }

        public List<packed_models.Directory> GetAllDirectories(string HttpAdress = "http://127.0.0.1:8080")
        {
            WebRequest request = WebRequest.Create($"{HttpAdress}/api/directories/get");
            request.Method = "GET";
            request.ContentType = "application/json";
            List<packed_models.Directory> res = new List<packed_models.Directory>();

            var httpResponse = (HttpWebResponse)request.GetResponse();
            using (var streamReader = new StreamReader(httpResponse.GetResponseStream()))
            {
                var result = streamReader.ReadToEnd();
                Console.WriteLine(result);
                res = JsonConvert.DeserializeObject<List<packed_models.Directory>>(result);
            }
            Console.WriteLine();
            return res;
        }

        public List<Material> GetAllMaterials(string HttpAdress = "http://127.0.0.1:8080")
        {
            WebRequest request = WebRequest.Create($"{HttpAdress}/api/materials/get");
            request.Method = "GET";
            request.ContentType = "application/json";
            List<Material> res = new List<Material>();

            var httpResponse = (HttpWebResponse)request.GetResponse();
            using (var streamReader = new StreamReader(httpResponse.GetResponseStream()))
            {
                var result = streamReader.ReadToEnd();
                Console.WriteLine(result);
                res = JsonConvert.DeserializeObject<List<Material>>(result);
            }
            Console.WriteLine();
            return res;
        }

        public List<PackedMaterialCost> GetAllMaterialsRange(string _start_date, string _end_date, int idMaterial, string HttpAdress = "http://127.0.0.1:8080")
        {
            WebRequest request = WebRequest.Create($"{HttpAdress}/api/materials/get/range");
            request.Method = "POST";
            request.ContentType = "application/json";
            List<PackedMaterialCost> res = new List<PackedMaterialCost>();

            using (var streamWriter = new StreamWriter(request.GetRequestStream()))
            {
                string json = JsonConvert.SerializeObject(
                    new { material = new { id = Convert.ToString(idMaterial) },
                        start_date = _start_date,
                        end_date = _end_date 
                    }
                    );
                streamWriter.Write(json);
                Console.WriteLine(json);
            }

            var httpResponse = (HttpWebResponse)request.GetResponse();
            using (var streamReader = new StreamReader(httpResponse.GetResponseStream()))
            {
                var result = streamReader.ReadToEnd();
                res = JsonConvert.DeserializeObject<List<PackedMaterialCost>>(result);
            }
            Console.WriteLine();
            return res;
        }

        public List<Profession> GetAllProfessions(string HttpAdress = "http://127.0.0.1:8080")
        {
            WebRequest request = WebRequest.Create($"{HttpAdress}/api/professions/get");
            request.Method = "GET";
            request.ContentType = "application/json";
            List<Profession> res = new List<Profession>();

            var httpResponse = (HttpWebResponse)request.GetResponse();
            using (var streamReader = new StreamReader(httpResponse.GetResponseStream()))
            {
                var result = streamReader.ReadToEnd();
                Console.WriteLine(result);
                res = JsonConvert.DeserializeObject<List<Profession>>(result);
            }
            Console.WriteLine();
            return res;
        }

        public void GetProfessionCostsForPeriod(int idProfession, string _start_date, string _end_date, string HttpAdress = "http://127.0.0.1:8080")
        {
            WebRequest request = WebRequest.Create($"{HttpAdress}/api/professions/get/range");
            request.Method = "POST";
            request.ContentType = "application/json";
            List<Profession> res = new List<Profession>();

            using (var streamWriter = new StreamWriter(request.GetRequestStream()))
            {
                string json = JsonConvert.SerializeObject(new { profession = new { id = idProfession }, start_date = _start_date, end_date = _end_date });
                streamWriter.Write(json);
            }

            var httpResponse = (HttpWebResponse)request.GetResponse();
            using (var streamReader = new StreamReader(httpResponse.GetResponseStream()))
            {
                var result = streamReader.ReadToEnd();
                res = JsonConvert.DeserializeObject<List<Profession>>(result);
            }
            Console.WriteLine();
        }
    }


}
