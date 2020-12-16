#!/bin/bash

echo "Build frontend"
cd frontend
npm i
npm run build
docker build . -t firl.us-frontend
cd ..
echo "Build backend"
cd backend
docker build . -t firl.us-backend
cd..
echo "Compose"
docker-compose up -d