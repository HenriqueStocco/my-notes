class Order {
  /** @type {"open" | "closed"} orderStatus */
  #_orderStatus = "open";

  /** @param {readonly { name: string, price: number }[]} shoppingCart */
  constructor(shoppingCart) {}

  /**
   * Get the order status
   * @returns {"open" | "closed"}
   */
  get orderStatus() {
    return this.#_orderStatus;
  }

  /**
   * @returns {void}
   */
  checkout() {
    if (this.shoppingCart.isEmpty()) {
      console.log("Empty shopping cart");

      return;
    }

    this.#_orderStatus = "closed";
    this.sendMessage(`Your request with the total ${this.total()} was already received`);
    this.saveOrder();
    this.shoppingCart.clear();
  }

  /**
   * @param {string} msg
   * @returns {void}
   */
  sendMessage(msg) {
    console.log(`Message sended: ${msg}`);
  }

  /** @returns {void} */
  saveOrder() {
    console.log("Order successfully saved...");
  }
}

module.exports = { Order };
