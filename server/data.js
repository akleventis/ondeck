/* 
 GET /person/{id}
 curl --location --request GET 'http://localhost:8080/person/1'

 POST /person
 curl --location --request POST 'http://localhost:8080/person' \
  --header 'Content-Type: application/json' \
  --data-raw '{
      "name": "Tanner",
      "phone": 1111111111
  }'
 
  PATCH /person/{id} =>  only send values to be updated (true PATCH) 
  curl --location --request PATCH 'http://localhost:8080/person/3' \
  --header 'Content-Type: application/json' \
  --data-raw '{
      "name": "Tanner",
      "phone": 1234567890
  }'
*/
const personResponse = {
  "id": 3,
  "name": "Tanner",
  "phone": 1111111111,
};

/* 
 GET /persons
 curl --location --request GET 'http://localhost:8080/persons'
*/
const personsResponse = [
  {
    "id": 4,
    "name": "Shelbe",
    "phone": 1112223333,
  },
  {
    "id": 3,
    "name": "cheeze-dawg",
    "phone": 2223334444,
  },
  {
    "id": 1,
    "name": "Toop",
    "phone": 1111111111,
  },
];

/*
  GET /drink/{id}
  curl --location --request GET 'http://localhost:8080/drink/1'

  POST /drink
  curl --location --request POST 'http://localhost:8080/drink' \
  --header 'Content-Type: application/json' \
  --data-raw '{
      "name": "espresso martini",
      "price": 10
  }'

  PATCH /drink/{id}  =>  only send values to be updated (true PATCH) 
  curl --location --request PATCH 'http://localhost:8080/drink/1' \
  --header 'Content-Type: application/json' \
  --data-raw '{
    "price": 5
  }'

  DELETE /drink/{id} => returns status gone
*/
const drinkResponse = {
  "id": 1,
  "name": "espresso martini",
  "price": 10,
};

/* 
 GET /drinks
 curl --location --request GET 'http://localhost:8080/drinks'
*/
const drinksResponse = [
  {
      "id": 1,
      "name": "espresso martini",
      "price": 10
  },
  {
      "id": 4,
      "name": "old fashion",
      "price": 10
  },
  {
      "id": 3,
      "name": "taquila",
      "price": 5
  }
]

/* 
 POST /order/{person_id}
 curl --location --request POST 'http://localhost:8080/order/1' \
  --header 'Content-Type: application/json' \
  --data-raw '{
      "order": [
          {
              "drink_id": 1,
              "comment": ""
          },
          {
              "drink_id": 3,
              "comment": "double please"
          }
      ]
  }'
*/
const orderResponse = {
  "person": {
      "id": 1,
      "name": "Toop",
      "phone": 1111111111
  },
  "drinks": [
      {
          "drink_id": 1,
          "name": "espresso martini",
          "price": 10,
          "comment": ""
      },
      {
          "drink_id": 3,
          "name": "taquila",
          "price": 5,
          "comment": "double please"
      }
  ],
  "order_number": 8
};

/* 
 GET /orders/{person_id}
 curl --location --request GET 'http://localhost:8080/orders/1'
*/
const personOrderResponse = [
  {
      "person": {
          "id": 1,
          "name": "Toop",
          "phone": 1111111111
      },
      "drinks": [
          {
              "drink_id": 1,
              "name": "espresso martini",
              "price": 10,
              "comment": "hold the gin"
          },
          {
              "drink_id": 3,
              "name": "taquila",
              "price": 10,
              "comment": "swag me out"
          }
      ],
      "order_number": 5
  },
  {
      "person": {
          "id": 1,
          "name": "Toop",
          "phone": 1111111111
      },
      "drinks": [
          {
              "drink_id": 1,
              "name": "espresso martini",
              "price": 10,
              "comment": ""
          },
          {
              "drink_id": 3,
              "name": "taquila",
              "price": 5,
              "comment": "double please"
          }
      ],
      "order_number": 8
  }
];

/* 
 GET /queue
 curl --location --request GET 'http://localhost:8080/queue'
*/
const queueResponse = {
  queue: {
    5: {
      person: {
        id: 1,
        name: "Toop",
        phone: "1111111111",
      },
      drinks: [
        {
          drink_id: 1,
          name: "espresso martini",
          price: 10,
          comment: "hold the gin",
        },
        {
          drink_id: 3,
          name: "taquila",
          price: 10,
          comment: "swag me out",
        },
      ],
      order_number: 5,
    },
    6: {
      person: {
        id: 4,
        name: "Shelbe",
        phone: 1112223333,
      },
      drinks: [
        {
          drink_id: 1,
          name: "espresso martini",
          price: 10,
          comment: "",
        },
        {
          drink_id: 3,
          name: "taquila",
          price: 5,
          comment: "",
        },
      ],
      order_number: 6,
    },
  },
};
