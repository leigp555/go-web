// window.dom = {
//     find(){
//         你来补全
//     },
//     style(){
//         你来补全
//     },
//     each(){
//         你来补全
//     }
// }

// const div = dom.find('#test>.red')[0] // 获取对应的元素
// dom.style(div, 'color', 'red') // 设置 div.style.color

// const divList = dom.find('.red') // 获取多个 div.red 元素
// dom.each(divList, (n)=> console.log(n)) // 遍历 divList 里的所有元素

window.dom = {
  find(selector, node) {
    return (node || document).querySelectorAll(selector);
  },
  style(node, name, value) {
    //node color red
    if (arguments.length === 3) {
      node.style[name] = value;
    }

    if (arguments.length === 2) {
      if (typeof name === "string") {
        //node color
        return node.style[name];
      }
      if (name instanceof Object) {
        let object = name;
        for (key in object) {
          node.style[key] = object[key];
        }
      }
    }
  },
  children(node) {
    return node.children;
  },
  each(node, name, value) {
    let a = dom.children(dom.find(node)[0]);

    if (arguments.length === 3) {
      for (let i = 0; i < a.length; i++) {
        dom.style(a[i], name, value);
      }
    } else if (arguments.length === 1) {
      let xx = [];
      for (let i = 0; i < a.length; i++) {
        xx.push(a[i]);
      }
      return xx;
    }
  },
};
//第一题
const div = dom.find("#test>.red")[0]; // 获取对应的元素
console.log(div);
//第二题
dom.style(div, "color", "red"); // 设置 div.style.color
//第三题
const divList = dom.find(".red"); // 获取多个 div.red 元素
console.log(divList);
//第四题
let a = dom.each("#test"); // 遍历 divList 里的所有元素
console.log(a);
