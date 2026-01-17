export function capitalize(string) {
   string = string.toLowerCase();
   var arr = string.split(' ');
   for (var i = 0; i < arr.length; i++) {
      arr[i] = arr[i].charAt(0).toUpperCase() + arr[i].slice(1);
   }
   return arr.join(' ');
}
