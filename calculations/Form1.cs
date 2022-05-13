using calculations.api;
using calculations.models;
using calculations.packed_models;

namespace calculations
{
    public partial class Form1 : Form
    {
        public Form1()
        {
            InitializeComponent();
        }

        private void button1_Click(object sender, EventArgs e)
        {
            /*HttpClient client = new HttpClient();
            client.BaseAddress = new Uri("127.0.0.1");
            HttpResponseMessage respponse = client.Send(new HttpRequestMessage(HttpMethod.Get, "127.0.0.1:8080/api/ping"));
            MessageBox.Show(respponse.ToString());*/

            WebApiClient apiClient = new WebApiClient();
            //apiClient.Authorization(new models.User { Login="admin", Password="admin" });
            //apiClient.CheckPingServerOrRulesAuthUser(new models.User { Token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwZXJtaXNzaW9ucyI6ImdvZCBwZXJtaXNzaW9ucyJ9.i-MW_1RhmPQ-6sJErFwjUvw1NhYpCvR6HQ9TaskJAIc" }, "super_user");
            // apiClient.CheckPingServerOrRulesAuthUser();
            //apiClient.GetAllCostitems();
            //apiClient.GetAllCalculations();


            /*PackedCalculation calculation = new PackedCalculation();
            calculation.Start_date = "2019-04-13";
            calculation.End_date = "2019-05-13";

            calculation.Id = 0;

            calculation.Product = new PackedProductWithActualCosts();
            calculation.Product.Id = 1;

            PackedCostItem ci1 = new PackedCostItem();
            ci1.Id = 2;
            ci1.Cost = 98;

            PackedCostItem ci2 = new PackedCostItem();
            ci2.Id = 3;
            ci2.Cost = 99;
            calculation.Cost_items = new List<PackedCostItem>() {
                ci1, ci2
            };

            apiClient.CreateCalculation(calculation);
            User user = new User();
            user.Token = @"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwZXJtaXNzaW9ucyI6ImdvZCBwZXJtaXNzaW9ucyJ9.i-MW_1RhmPQ-6sJErFwjUvw1NhYpCvR6HQ9TaskJAIc";*/
            //apiClient.UpdateCalculation(calculation, user);

            //apiClient.GetCalculation(1);

            //apiClient.GetAllMaterials();

            //apiClient.GetAllProfessions();

            //apiClient.GetProfessionCostsForPeriod(1, "", "");

            //apiClient.GetAllProducts();

            User user = new User();
            user.Token = @"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwZXJtaXNzaW9ucyI6ImdvZCBwZXJtaXNzaW9ucyJ9.i-MW_1RhmPQ-6sJErFwjUvw1NhYpCvR6HQ9TaskJAIc";

            //apiClient.DeleteCalculation(user, 1);

            //apiClient.GetAllDirectories();
            //apiClient.GetAllMaterialsRange("1999-01-02", "2025-01-02", 1);

        }
    }
}