using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;
using calculations.api;
using calculations.models;
using calculations.packed_models;

namespace calculations
{
    public partial class CreateCalculationForm : Form
    {
        public CreateCalculationForm()
        {
            InitializeComponent();
        }

        private void dataGridView1_CellContentClick(object sender, DataGridViewCellEventArgs e)
        {

        }

        private void CreateCalculationForm_Load(object sender, EventArgs e)
        {
            WebApiClient apiClient = new WebApiClient();
            List<PackedProduct> products = apiClient.GetAllProducts();
            List<CostItem> сostItems = apiClient.GetAllCostitems();
            dataGridView1.Rows.Clear();
            dataGridView2.Rows.Clear();
            foreach (PackedProduct product in products)
                dataGridView1.Rows.Add(product.Name, product.Id);
            foreach (CostItem costItem in сostItems)
                dataGridView2.Rows.Add(costItem.Name, 0, costItem.Id);
        }

        private void dataGridView2_CellContentClick(object sender, DataGridViewCellEventArgs e)
        {

        }

        private void button1_Click(object sender, EventArgs e)
        {
            if(idCalculationsObject == 0)
            {
                MessageBox.Show("Выберете продукт, для которого составляется калькуляция");
                return;
            }
            WebApiClient apiClient = new WebApiClient();
            PackedCalculation calculation = new PackedCalculation();
            string startDate = $"{ dateTimePicker1.Value.Year }-{ dateTimePicker1.Value.Month }-{ dateTimePicker1.Value.Day }";
            string endDate = $"{ dateTimePicker2.Value.Year }-{ dateTimePicker2.Value.Month }-{ dateTimePicker2.Value.Day }";
            calculation.Start_date = startDate;
            calculation.End_date = endDate;
            calculation.Product = new PackedProductWithActualCosts();
            calculation.Cost_items = new List<PackedCostItem>();
            calculation.Product.Id = idCalculationsObject;
            if(dataGridView2.SelectedRows.Count == 0)
            {
                MessageBox.Show("Не выбраны статьи затрат");
                return;
            }
            for (int i = 0; i < dataGridView2.SelectedRows.Count; ++i)
            {
                DataGridViewRow row = dataGridView2.SelectedRows[i];
                PackedCostItem costItem = new PackedCostItem();
                costItem.Id = Convert.ToInt32((row.Cells)[2].Value);
                costItem.Cost = Convert.ToSingle((row.Cells)[1].Value);
                costItem.Name = Convert.ToString((row.Cells)[0].Value);
                
                calculation.Cost_items.Add(costItem);
            }
            apiClient.CreateCalculation(calculation);
        }

        int idCalculationsObject = 0;
        private void dataGridView1_CellClick(object sender, DataGridViewCellEventArgs e)
        {
            idCalculationsObject = Convert.ToInt32(dataGridView1.Rows[e.RowIndex].Cells[1].Value);
        }
    }
}
