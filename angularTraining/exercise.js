  
// Import the path module
const path = require('path');
const process = require('process');
const os =require('os')
console.log("Current directory:", __dirname);

path1 = path.resolve("go");
console.log(path1)

path2 = path.resolve("go", "mod");
console.log(path2)

path3=path.resolve("go/mod")
console.log(path3)
// // Treating of the first segment
// // as root, ignoring the current directory
path3 = path.resolve("/go/mod", "cache");
console.log(path3)

//////parse////
//posix
path1 = path.parse("/angular/app1.js");
console.log(path1);
   
path2 = path.parse("/angular/test/test.js");
console.log(path2);
//windows
path3=path.parse("C:\\Users\\nehab\\angular\\app1.js")
console.log(path3)

/////relative//

//from: It is the file path that would be used as base path.
//to: It is the file path that would be used to find the relative path.
console.log("relative")
path1 = path.relative("angular", "test/test.js");
console.log(path1)
   
path2 = path.relative("angular/test", "test/test.js");
console.log(path2)
   
// When both the paths are same
// It returns blank string
path3 = path.relative("angular/test", "angular/test");
console.log(path3)
// path4=path.relative("datta/test")  //invalid path  gives err 
// console.log(path4)


////process arguments while taking input from cmd line
console.log("input from user")

// This property returns an array containing the arguments passed to the process when run it in the command line.
console.log(process.argv); //argv is argumrnt values
var args=process.argv.slice(2);
console.log(args)
console.log(args[0])

////total and free memory

//This method returns an integer value that specifies the amount of free system memory in bytes.
var free_memory = os.freemem();
console.log(free_memory)
var total_memory=os.totalmem();
console.log(total_memory)