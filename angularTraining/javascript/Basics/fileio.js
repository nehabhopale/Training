const fs=require('fs');

filenames = fs.readdirSync(__dirname);
  
console.log(filenames);
  
let data = "hii";
  
fs.writeFileSync("ex.txt", data);
const data1 = fs.readFileSync('./time.js',{encoding:'utf8', flag:'r'});
 
// Display the file data
console.log(data1);

