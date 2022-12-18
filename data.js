// POST /person
const personResponse = {
  name: "Tanner",
  phone: "1111111111",
};

// GET /orders
const ordersResponse = [
  {
    person: {
      name: "Tanner",
      phone: "1111111111",
    },
    drink: {
      name: "Gin and Tonic",
      price: 10,
    },
    order_number: 1,
  },
  {
    person: {
      name: "Toop",
      phone: "8888888888",
    },
    drink: {
      name: "Taquila Soda",
      price: 10,
    },
    order_number: 5,
  },
  {
    person: {
      name: "Ryan",
      phone: "9999999999",
    },
    drink: {
      name: "Taquila Soda",
      price: 10,
    },
    order_number: 3,
  },
];

// GET /queue
const queueResponse = {
  2: {
    person: {
      name: "Toop",
      phone: "8888888888",
    },
    drink: {
      name: "Taquila Soda",
      price: 10,
    },
    order_number: 2,
  },
  1: {
    person: {
      name: "Tanner",
      phone: "1111111111",
    },
    drink: {
      name: "Gin and Tonic",
      price: 10,
    },
    order_number: 1,
  },
  3: {
    person: {
      name: "Ryan",
      phone: "7777777777",
    },
    drink: {
      name: "Taquila Soda",
      price: 10,
    },
    order_number: 3,
  },
};

// POST /order
const orderResponse = [
  {
    person: {
      name: "Tanner",
      phone: "1111111111",
    },
    drink: {
      name: "Gin and Tonic",
      price: 10,
    },
    order_number: 1,
  },
  {
    person: {
      name: "Tanner",
      phone: "1111111111",
    },
    drink: {
      name: "Taquila Soda",
      price: 10,
    },
    order_number: 2,
  },
  {
    person: {
      name: "Tanner",
      phone: "1111111111",
    },
    drink: {
      name: "Michelob",
      price: 10,
    },
    order_number: 3,
  },
];
