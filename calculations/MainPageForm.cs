using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;
using System.Net;
using calculations.api;
using calculations.models;
using calculations.packed_models;

namespace calculations
{
    public partial class MainPageForm : Form
    {
        public MainPageForm()
        {
            InitializeComponent();
        }

        private void tabPage1_Click(object sender, EventArgs e)
        {

        }

        private void Pr(object sender, EventArgs e)
        {

        }

        private void DirectoryChenge(object sender, EventArgs e)
        {
            if(tabControl2.SelectedTab == ProfessionsTab)
            {
                WebApiClient apiClient = new WebApiClient();
                List<Profession> professions = apiClient.GetAllProfessions();
                dataGridView1.Rows.Clear();
                foreach (Profession profession in professions)
                    dataGridView1.Rows.Add(profession.Description, profession.Cipher);
            } 
            else if(tabControl2.SelectedTab == ExpendituresTab)
            {
                WebApiClient apiClient = new WebApiClient();
                List<CostItem> сostItems = apiClient.GetAllCostitems();
                dataGridView2.Rows.Clear();
                foreach(CostItem costItem in сostItems)
                    dataGridView2.Rows.Add(costItem.Name, costItem.Cipher);
            } 
            else if(tabControl2.SelectedTab == MaterialsTab)
            {
                WebApiClient apiClient = new WebApiClient();
                List<Material> materials = apiClient.GetAllMaterials();
                dataGridView3.Rows.Clear();
                foreach (Material material in materials)
                    dataGridView3.Rows.Add(material.Name, material.Cipher);
            } 
            else if(tabControl2.SelectedTab == ProductsTab)
            {
                WebApiClient apiClient = new WebApiClient();
                List<PackedProduct> products = apiClient.GetAllProducts();
                dataGridView4.Rows.Clear(); 
                foreach (PackedProduct product in products)
                    dataGridView4.Rows.Add(product.Name);
            }

        }

        private void dataGridView2_CellContentClick(object sender, DataGridViewCellEventArgs e)
        {

        }

        private void CalculationChenge(object sender, EventArgs e)
        {
            if (tabControl1.SelectedTab == CalculationsTab)
            {
                WebApiClient apiClient = new WebApiClient();
                List<PackedCalculation> calculations = apiClient.GetAllCalculations();
                dataGridView5.Rows.Clear();
                if (calculations != null)
                {
                    foreach (PackedCalculation calculation in calculations)
                    {
                        calculation.Start_date = calculation.Start_date.Split("T")[0];
                        calculation.End_date = calculation.End_date.Split("T")[0];
                        dataGridView5.Rows.Add(calculation.Product.Name, calculation.Start_date, calculation.End_date, calculation.Full_cost, calculation.Id);
                    }
                }
            }
        }

        int idCalculation = 0;
        private void button1_Click(object sender, EventArgs e)
        {
            if(idCalculation == 0)
            {
                MessageBox.Show("Выберите калькуляцию");
                return;
            }
            WebApiClient apiClient = new WebApiClient();
            PackedCalculation calculation = apiClient.GetCalculation(idCalculation);
            DetailedCalculationForm detailedCalculationForm = new DetailedCalculationForm(calculation);
            detailedCalculationForm.Show();
        }

        private void AddCalculationButton_Click(object sender, EventArgs e)
        {
            CreateCalculationForm createCalculationForm = new CreateCalculationForm();
            createCalculationForm.Show();
        }

        private void dataGridView4_CellContentClick(object sender, DataGridViewCellEventArgs e)
        {
            
        }

        private void dataGridView5_CellClick(object sender, DataGridViewCellEventArgs e)
        {
            idCalculation = Convert.ToInt32(dataGridView5.Rows[e.RowIndex].Cells[4].Value);
        }

        private void DeleteCalculationButton_Click(object sender, EventArgs e)
        {
            if (idCalculation == 0)
            {
                MessageBox.Show("Выберите калькуляцию для удаления");
                return;
            }
            WebApiClient apiClient = new WebApiClient();
            User user = new User();
            user.Token = @"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwZXJtaXNzaW9ucyI6ImdvZCBwZXJtaXNzaW9ucyJ9.i-MW_1RhmPQ-6sJErFwjUvw1NhYpCvR6HQ9TaskJAIc";
            apiClient.DeleteCalculation(user, idCalculation);
        }

        private void button2_Click(object sender, EventArgs e)
        {
            WebApiClient apiClient = new WebApiClient();
            List<PackedCalculation> calculations = apiClient.GetAllCalculations();
            dataGridView5.Rows.Clear();
            if (calculations != null)
            {
                foreach (PackedCalculation calculation in calculations)
                {
                    calculation.Start_date = calculation.Start_date.Split("T")[0];
                    DateTime startDate = new DateTime(
                        Convert.ToInt32(calculation.Start_date.Split("-")[0]),
                        Convert.ToInt32(calculation.Start_date.Split("-")[1]),
                        Convert.ToInt32(calculation.Start_date.Split("-")[2])
                        );
                    calculation.End_date = calculation.End_date.Split("T")[0];
                    DateTime endtDate = new DateTime(
                        Convert.ToInt32(calculation.End_date.Split("-")[0]),
                        Convert.ToInt32(calculation.End_date.Split("-")[1]),
                        Convert.ToInt32(calculation.End_date.Split("-")[2])
                        );
                    if(startDate >= dateTimePicker1.Value && endtDate <= dateTimePicker2.Value)
                        dataGridView5.Rows.Add(calculation.Product.Name, calculation.Start_date, calculation.End_date, calculation.Full_cost, calculation.Id);
                }
            }
        }
    }
}
