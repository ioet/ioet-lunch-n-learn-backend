# Lunch & Learn App

## 📋 Brief
The **Lunch & Learn App** is designed to facilitate informal learning sessions within organizations. It allows team members to organize, participate in, and track lunch & learn events, fostering a culture of continuous learning.

## 🛠️ Backend Tech Stack
- **Languaje:** Go
- **Framework:** Gin
- **Database:** Firebase Firestore
- **API:** RESTful services using Gin
- **Version Control:** Git

## 📂 Project Structure

```bash
/
│
├── api/
│   ├── config/
│   ├── dtos/
│   ├── events/
│   ├── exception_handlers/
│   └── routers/
│       ├── v1/
│       └── v2/
│
├── core/
│   ├── src/
│   │   ├── exceptions/
│   │   │   ├── business/
│   │   │   └── repository/
│   │   ├── orchestrators/
│   │   ├── repositories/
│   │   ├── use_cases/
│   │   └── models/
│   └── tests/
│
├── adapters/
│   ├── src/
│   │   ├── repositories/
│   └── tests/
│       └── use_cases/
├── factories
│   ├── config/
│   ├── orchestrators/
│   ├── repositories/
│   └── use_cases/
└── infrastructure/
```
Based on the [IOET Backend Standard](https://www.notion.so/ioet/Backend-Directory-Structure-4629fa856ff0400d9665a9c8cc8c1685)

## 🚀 Endpoints
### **Health Check**
- **GET** `/v1/health_check/` - To ping the API and verify the status.

### **User**
- **GET** `/v1/user/` - To get all users.
- **GET** `/v1/user/:id ` - To get a user by id.
- **POST** `/v1/user/ ` - To create a user.
- **PUT** `/v1/user/house  ` - To modify the user's house.
- **PUT** `/v1/user/points  ` - To modify the user's points.

### **House**
- **GET** `/v1/house/` - To get all houses.
- **GET** `/v1/house/:id` - To get a house by id.
- **POST** `/v1/house/` - To create a house.

### **Lunch n Learn**
- **GET** `/v1/lunch_n_learn/` - To get all lunch n learns.
- **GET** `/v1/lunch_n_learn/:id` - To get a lunch n learn by id.
- **POST** `/v1/lunch_n_learn/` - To create a lunch n learn.
- **PUT** `/v1/lunch_n_learn/` - To update a lunch n learn info.
- **PUT** `/v1/lunch_n_learn/assistant` - To add an assistant to a lunch n learn.
- **PUT** `/v1/lunch_n_learn/presenter` - To add an presenter to a lunch n learn.


## 🛠️ Setup and Installation
1. **Clone the repository:**
  ```bash 
  git clone https://github.com/ioet/ioet-lunch-n-learn-backend.git

  cd ioet-lunch-n-learn-backend
  ```

2. **Install Go dependencies:**
  ```bash
  go mod download
  ```

3. **Run the server:**
  ```bash
  go run main.go
  ```

## 🛡️ Security
The internal IOET Auth Service is used for authentication and authorization.

> **__IMPORTANT:__** API routes are protected with middleware to ensure secure access.

## 🤝 How to Contribute
- Fork the repository.
- Create a new feature branch (git checkout -b feature/YourFeatureName).
- Commit your changes (git commit -m 'Add some feature').
- Push to the branch (git push origin feature/YourFeatureName).
- Open a Pull Request.

