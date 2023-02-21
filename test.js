// const { exec } = require('child_process');
// exec(`ls -c ${__dirname}`, (error, stdout, stderr) => {
//   if (error) {
//           console.error(`exec error: ${error}`);
//     return;
//   }
//   console.log(`stdout: ${stdout}`);
//   console.error(`stderr: ${stderr}`);
// });

// const reversedArray = (s) => {
//     try {
//         console.log(s.split("").reverse().join(""))
//     } catch (err){
//         console.error(err)
//         console.log(s)
//     }
// }
// reversedArray(1234)

const days = ["Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"]

const dateDay = (dateString) => {
    const newDate = new Date(dateString)
    const day = newDate.getDay()
    return days[day]
}
console.log(dateDay("10/11/2009"))
console.log(dateDay("11/10/2010"))