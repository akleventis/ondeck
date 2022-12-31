// Concern: Sending phone number over http may be privacy risk
// Solution: Assign an ID to each person && use as primary key throughout 
// POST /person && GET /person/{id}
const personResponse = {
  id: 123,
  name: "Tanner",
  phone: "1111111111",
};

// GET /queue
const queueResponse = {
  2: {
    person: {
      id: 123,
      name: "Tanner",
      phone: "2222222222",
    },
    drinks: [
      {
        name: "Michelob",
        price: 4,
      },
      {
        name: "Taquila Soda",
        price: 10,
      },
    ],
    order_number: 2,
  },
  1: {
    person: {
      id: 124,
      name: "Toop",
      phone: "8888888888",
    },
    drinks: [
      {
        name: "Gin and Tonic",
        price: 10,
      },
      {
        name: "Taquila Soda",
        price: 10,
      },
    ],
    order_number: 1,
  },
};

// POST /order
const orderResponse = [
  {
    person: {
      id: 123,
      name: "Tanner",
      phone: "1111111111",
    },
    drinks: [
      {
        name: "Gin and Tonic",
        price: 10,
      },
      {
        name: "Taquila Soda",
        price: 10,
      },
    ],
    order_number: 2,
  },
];

// GET /orders/{person_id}
const personOrderResponse = [
  {
    person: {
      id: 123,
      name: "Tanner",
      phone: "1111111111",
    },
    drinks: [
      {
        name: "Gin and Tonic",
        price: 10,
      },
      {
        name: "Taquila Soda",
        price: 10,
      },
    ],
    order_number: 2,
  },
  {
    person: {
      id: 123,
      name: "Tanner",
      phone: "1111111111",
    },
    drinks: [
      {
        name: "Michelob",
        price: 4,
      },
      {
        name: "Gin and Tonic",
        price: 10,
      },
    ],
    order_number: 5,
  },
]



// --------------------------- //


// Since we really only care about the queue, this call seems unnecessary
// GET /orders
// const ordersResponse = [
//   {
//     person: {
//       name: "Tanner",
//       phone: "1111111111",
//     },
//     drink: {
//       name: "Gin and Tonic",
//       price: 10,
//     },
//     order_number: 1,
//   },
//   {
//     person: {
//       name: "Toop",
//       phone: "8888888888",
//     },
//     drink: {
//       name: "Taquila Soda",
//       price: 10,
//     },
//     order_number: 5,
//   },
//   {
//     person: {
//       name: "Ryan",
//       phone: "9999999999",
//     },
//     drink: {
//       name: "Taquila Soda",
//       price: 10,
//     },
//     order_number: 3,
//   },
// ];
