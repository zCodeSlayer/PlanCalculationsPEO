using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace calculations.packed_models
{
    public class PackedCalculation
    {
        public int Id { get; set; }
        public string Start_date { get; set; }
        public string End_date { get; set; }
        public PackedProductWithActualCosts Product { get; set; }
        public List<PackedCostItem> Cost_items { get; set; }
        public List<CalculatedCostItem> Calculated_cost_items { get; set;  }
        public float Full_cost { get; set; }
    }
}
