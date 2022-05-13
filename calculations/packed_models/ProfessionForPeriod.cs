using calculations.models;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace calculations.packed_models
{
    public class ProfessionForPeriod
    {
        public Profession Profession { get; set; }
        public Product Product { get; set; }
        public string StartDate { get; set; }
        public string EndDate { get; set; }

    }
}
