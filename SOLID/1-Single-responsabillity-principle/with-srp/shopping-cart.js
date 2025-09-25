"strict";

/**
 * Create a new Shopping cart
 * @class ShoppingCart
 **/
class ShoppingCart {
  /** @type {readonly { name: string, price: number }[]} _items */
  #_items = [];

  /**
   * Add an item into the _items object array
   * @param {{ name: string, price: number }} item
   * @returns {void}
   */
  addItem(item) {
    this.#_items.push(item);
  }

  /**
   * Remove an item from the _items object array
   * @param {number} index
   * @returns {void}
   */
  removeItem(index) {
    this.#_items.splice(index, 1);
  }

  /**
   * Get all addded items
   * @returns {readonly {name:string, price: number}[]} items
   */
  get items() {
    return this.#_items;
  }

  /**
   * Sum all prices and return the total price
   * @returns {number} total */
  total() {
    return +this.#_items.reduce((total, next) => total + next.price, 0).toFixed(2);
  }

  /**
   * Return true if the shopping cart is empty
   * @returns {boolean}
   */
  isEmpty() {
    return this.#_items.length === 0;
  }

  /**
   * Clear shopping cart
   * @returns {void} */
  clear() {
    console.log("Shopping cart was cleaned");

    this.#_items.length = 0;
  }
}

module.exports = { ShoppingCart };
