using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace calculations.packed_models
{
    public class Directory
    {
        public string Name { get; set; }
        public List<string> Columns { get; set; }
        public List<List<string>> Data { get; set; }
    }
}
