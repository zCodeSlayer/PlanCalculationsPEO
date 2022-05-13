using calculations.models;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace calculations.packed_models
{
    public class PackedProduct
    {
        public int Id { get; set; }
        public string Name { get; set; }
        public List<Material> Materials { get; set; }
        public List<Profession> Professions { get; set; }
    }
}
