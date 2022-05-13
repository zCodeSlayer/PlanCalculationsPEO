using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace calculations.models
{
    public class MaterialCost
    {
        public int Id { get; set; }
        public string ActualDate { get; set; }
        public float Cost { get; set; }
        public int IdNorm { get; set; }
        public int IdMaterial { get; set; }

    }
}
