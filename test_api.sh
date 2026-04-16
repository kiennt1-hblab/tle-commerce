#!/bin/bash

API_URL="http://localhost:3000/api"

echo "========================================="
echo "Testing BookStore API Endpoints"
echo "========================================="
echo ""

echo "0. Testing Health Check"
echo "-----------------------------------------"
curl -s "$API_URL/../health" | jq
echo ""

echo "1. Testing GET /api/books"
echo "-----------------------------------------"
curl -s "$API_URL/books" | jq '.[0:2]'
echo ""

echo "2. Testing GET /api/books/:id"
echo "-----------------------------------------"
curl -s "$API_URL/books/1" | jq
echo ""

echo "3. Testing GET /api/books/search?q=harry"
echo "-----------------------------------------"
curl -s "$API_URL/books/search?q=harry" | jq
echo ""

echo "4. Testing POST /api/orders"
echo "-----------------------------------------"
curl -s -X POST "$API_URL/orders" \
  -H "Content-Type: application/json" \
  -d '{
    "customer": {
      "name": "Test User",
      "email": "test@example.com",
      "phone": "0987654321",
      "address": "456 Test Street",
      "note": "API test order"
    },
    "items": [
      {"id": 1, "quantity": 1, "price": 12.99},
      {"id": 2, "quantity": 2, "price": 14.99}
    ],
    "total": 42.97
  }' | jq
echo ""

echo "========================================="
echo "All tests completed!"
echo "========================================="
