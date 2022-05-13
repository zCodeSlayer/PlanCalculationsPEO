using calculations.packed_models;
using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;

namespace calculations
{
    public partial class DetailedCalculationForm : Form
    {
        private PackedCalculation _calculation;
        public DetailedCalculationForm(PackedCalculation calculation)
        {
            _calculation = calculation;
            InitializeComponent();
        }

        private void DetailedCalculationForm_Load(object sender, EventArgs e)
        {
            this.Text = $"Подробная информация по калькуляции на продукт {_calculation.Product.Name}";
            StartDateLabel.Text = $"Дата начала: {_calculation.Start_date.Split("T")[0]}";
            FinishDateLabel.Text = $"Дата окончания: {_calculation.End_date.Split("T")[0]}";
            CalculationLabel.Text = $"Калькуляция: {_calculation.Full_cost}";
            dataGridView1.Rows.Clear();
            dataGridView2.Rows.Clear();
            dataGridView3.Rows.Clear();
            dataGridView4.Rows.Clear();

            foreach (MaterialWithActualCost material in _calculation.Product.Materials)
                dataGridView1.Rows.Add(material.Name, material.Cost);
            foreach (ProfessionWithActualCost profession in _calculation.Product.Professions)
                dataGridView2.Rows.Add(profession.Description, profession.Cost);
            foreach(PackedCostItem costItem in _calculation.Cost_items)
                dataGridView3.Rows.Add(costItem.Name, costItem.Cost);
            foreach (CalculatedCostItem calculatedCostItem in _calculation.Calculated_cost_items)
                dataGridView4.Rows.Add(calculatedCostItem.Name, calculatedCostItem.Cost);
        }

        private void label3_Click(object sender, EventArgs e)
        {

        }

        private void dataGridView4_CellContentClick(object sender, DataGridViewCellEventArgs e)
        {

        }
    }
}
