"strict";

const { ShoppingCart } = require("./shopping-cart.cjs");

const shoppingCart = new ShoppingCart();

(() => {
  for (let i = 0; i < 10; i++) {
    shoppingCart.addItem({
      name: `Random item ${i + 1}`,
      price: Number(Math.random(Math.floor() * 100).toFixed(2)),
    });
  }
})();

console.log(shoppingCart.items[2]);
console.log(shoppingCart.items);
console.log(shoppingCart.total());
console.log(shoppingCart.orderStatus);

shoppingCart.checkout();

console.log(shoppingCart.orderStatus);
