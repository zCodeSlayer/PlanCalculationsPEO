using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace calculations.packed_models
{
    public class PackedProductWithActualCosts
    {
        public int Id { get; set; }
        public string Name { get; set; }
        public List<MaterialWithActualCost> Materials { get; set; }
        public List<ProfessionWithActualCost> Professions { get; set; }
    }
}
