function tes(kata) {
  for (let i = 0; i < kata.length; i++) {
    let d = kata[i];
    let c = kata[kata.length - i - 1];    
    // console.log(d + " " + c);
    if (c != d) {
      return false;
    }
  }
  return true;
}
console.log(tes("katak"));
