
```yml

name: Go-App CI/CD

on:
  push:
    branches:
      - main
jobs:
  build:
    runs-on: ubuntu-latest

    steps:
    - name: Checkout Repository
      uses: actions/checkout@v4

    - name: Set up Go
      uses: actions/setup-go@v4
      with:
        go-version: 1.21

    - name: Install Dependencies
      run: go mod download

    - name: Build
      run: GOOS=linux go build -o build/invoiceapp -v
 
    - name: Create .env file
      run: |
        echo "SECRETKEY=${{ secrets.SECRETKEY }}" > .env
        echo "DBUSER=${{ secrets.DBUSER }}" >> .env
        echo "DBPASSWORD=${{ secrets.DBPASSWORD }}" >> .env
        echo "DBHOST=${{ secrets.DBHOST }}" >> .env
        echo "DBPORT=${{ secrets.DBPORT }}" >> .env
        echo "DBNAME=${{ secrets.DBNAME }}" >> .env
         
    - name: Copy output via scp
      uses: appleboy/scp-action@master
      with:
        host: ${{ secrets.VM_HOST }}
        username: ${{ secrets.VM_USERNAME }}
        key: ${{ secrets.VM_SSH_KEY }}
        port: ${{ secrets.VM_SSH_PORT }}
        source: "build/"
        target: "/var/www/invoice-go-app"

    - name: Deploy to Ubuntu VM
      uses: appleboy/ssh-action@master
      with:
        host: ${{ secrets.VM_HOST }}
        username: ${{ secrets.VM_USERNAME }}
        key: ${{ secrets.VM_SSH_KEY }}
        port: ${{ secrets.VM_SSH_PORT }}
        script: ls
        
```
    
 

## Introduction
This YAML code defines a GitHub Actions workflow for Continuous Integration and Continuous Deployment (CI/CD) of a Go application.

## Steps

### 1. Checkout Repository
   - **Name**: Checkout Repository
   - **Action**: [actions/checkout@v4]
   - **Purpose**: Fetches the latest commit from the repository.

### 2. Set up Go
   - **Name**: Set up Go
   - **Action**: [actions/setup-go@v4]
   - **Purpose**: Sets up the Go environment with the specified Go version.

### 3. Install Dependencies
   - **Name**: Install Dependencies
   - **Purpose**: Downloads Go module dependencies using `go mod download`.

### 4. Build
   - **Name**: Build
   - **Purpose**: Compiles the Go application into an exe file (`invoiceapp.exe`) and places it in the 'build/' folder.

### 5. Create .env file
   - **Name**: Create .env file
   - **Purpose**: Creates a `.env` file with sensitive information (e.g., database credentials) sourced from GitHub Secrets.

### 6. Copy output using scp
   - **Name**: Copy output via scp
   - **Action**: [appleboy/scp-action@master]
   - **Purpose**: Copies the files of the 'build/' directory to a remote server using SCP. The target directory on the server is "/var/www/invoice-go-app".

### 7. Deploy to Ubuntu VM
   - **Name**: Deploy to Ubuntu VM
   - **Action**: [appleboy/ssh-action@master]
   - **Purpose**: Establish an SSH connection to an Ubuntu VM and run the 'ls' command.

## Secrets Used
List of the secrets stored in GitHub Actions Secrets:
- `SECRETKEY`
- `DBUSER`
- `DBPASSWORD`
- `DBHOST`
- `DBPORT`
- `DBNAME`
- `VM_HOST`
- `VM_USERNAME`
- `VM_SSH_KEY`
- `VM_SSH_PORT`

## Deployment Flow
1. The Go application is built.
2. A `.env` file is created with sensitive information.
3. The files of the 'build/' folder are copied to a remote server using SCP.
4. The Go app is deployed to an Ubuntu VM by initiating an SSH connection and executing deployment scripts.

 ## Backend database diagram


![Diagramdbinvoiceapp](https://github.com/rebiiin/invoiceapp/assets/3890058/e54f8155-9fe1-4978-80c1-8beddc9a21d8)


  
## MySQL database ddl script of the backend

```sql

CREATE DATABASE `dbinvoiceapp` 

CREATE TABLE `customers` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `first_name` varchar(50) DEFAULT NULL,
  `last_name` varchar(50) DEFAULT NULL,
  `address` varchar(150) DEFAULT NULL,
  `phone` varchar(20) DEFAULT NULL,
  `balance` double DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_customers_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `invoice_lines` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `quantity` bigint DEFAULT NULL,
  `price` double DEFAULT NULL,
  `product_id` bigint DEFAULT NULL,
  `invoice_id` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_invoice_lines_deleted_at` (`deleted_at`),
  KEY `fk_invoices_invoice_lines` (`invoice_id`),
  KEY `fk_products_invoice_lines` (`product_id`),
  CONSTRAINT `fk_invoices_invoice_lines` FOREIGN KEY (`invoice_id`) REFERENCES `invoices` (`id`),
  CONSTRAINT `fk_products_invoice_lines` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `invoices` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `invoice_number` varchar(100) DEFAULT NULL,
  `date` datetime(3) DEFAULT NULL,
  `total` double DEFAULT NULL,
  `customer_id` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_invoices_deleted_at` (`deleted_at`),
  KEY `fk_customers_invoices` (`customer_id`),
  CONSTRAINT `fk_customers_invoices` FOREIGN KEY (`customer_id`) REFERENCES `customers` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `products` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `barcode` varchar(50) DEFAULT NULL,
  `quantity` bigint DEFAULT NULL,
  `price` double DEFAULT NULL,
  `supplier_id` bigint DEFAULT NULL,
  `name` varchar(50) DEFAULT NULL,
  `image` varchar(250) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_products_deleted_at` (`deleted_at`),
  KEY `fk_suppliers_products` (`supplier_id`),
  CONSTRAINT `fk_suppliers_products` FOREIGN KEY (`supplier_id`) REFERENCES `suppliers` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `suppliers` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(50) DEFAULT NULL,
  `phone` varchar(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_suppliers_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `users` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `email` varchar(150) DEFAULT NULL,
  `password` varchar(500) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`),
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

```
