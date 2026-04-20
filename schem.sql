
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email TEXT UNIQUE NOT NULL,
    role VARCHAR(50),
    password TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS shops (
    id SERIAL PRIMARY KEY,
    name TEXT UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS user_shops (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    shop_id INT NOT NULL REFERENCES shops(id) ON DELETE CASCADE,
    UNIQUE(user_id, shop_id)
);

CREATE TABLE IF NOT EXISTS employees (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    shop_id INT NOT NULL REFERENCES shops(id) ON DELETE CASCADE,
    role TEXT NOT NULL CHECK (role IN ('admin', 'supervisor', 'employee')),
    created_at TIMESTAMP DEFAULT NOW(),
    UNIQUE(user_id, shop_id)
);

-- ENGINE DEFINITIONS & CONFIGURATION
-- -----------------------------------------------------

CREATE TABLE IF NOT EXISTS blockflows (
    id INT PRIMARY KEY,
    block_id INT,
    name VARCHAR(150),
    start_process_id INT NOT NULL REFERENCES processes(id),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS flow_shops (
    flow_id INT NOT NULL REFERENCES blockflows(id) ON DELETE CASCADE,
    shop_id INT NOT NULL REFERENCES shops(id) ON DELETE CASCADE,
    PRIMARY KEY (flow_id, shop_id)
);

-----------Examples
INSERT INTO shops (name) 
VALUES ('My Coffee Shop'), ('Electronics Store'), ('Book Haven');




-- -------------------------------------------------------------------


    CREATE TABLE IF NOT EXISTS processes (
        id SERIAL PRIMARY KEY,
        form_name VARCHAR(100),
        config_json JSONB NOT NULL,
        next_process_id INT REFERENCES processes(id),
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

    CREATE TABLE IF NOT EXISTS executions (
        id SERIAL PRIMARY KEY,
        process_id INT REFERENCES processes(id),
        user_id VARCHAR(50) ,
        shop_id VARCHAR(50) ,
        form_name VARCHAR(100) ,
        data JSONB NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        -- CONSTRAINT unique_exec UNIQUE (process_id, user_id, shop_id, form_name)
    );


INSERT INTO processes (id, form_name, config_json) VALUES
-- BLOCK 1: STORE MANAGEMENT
(101, 'categories', '{"process_id": 101, "form_name": "categories", "table_name": "categories", "use_table": true, "display_type": "table", "pagination_mode": "server", "is_multi_row": false, "is_one_object": true, "next_process_id": 102, "form": {"name": "Category", "variables": [{"name": "name", "type": "text", "label": "Category Name"}]}}'),
(102, 'products', '{"process_id": 102, "form_name": "products", "table_name": "products", "use_table": true, "display_type": "grid", "pagination_mode": "server", "is_multi_row": false, "is_one_object": true, "form": {"name": "Product", "variables": [{"name": "name", "type": "text", "label": "Product Name"}, {"name": "category_id", "type": "select", "label": "Category", "data_source": {"type": "table", "table": "categories", "label_field": "name", "value_field": "id"}}]}}'),

-- BLOCK 2: SUPPLY CHAIN
(201, 'suppliers', '{"process_id": 201, "form_name": "suppliers", "table_name": "suppliers", "use_table": true, "display_type": "table", "pagination_mode": "server", "is_multi_row": false, "is_one_object": true, "next_process_id": 202, "form": {"name": "Supplier", "variables": [{"name": "name", "type": "text", "label": "Supplier Name"}]}}'),
(202, 'purchases', '{"process_id": 202, "form_name": "purchases", "table_name": "purchases", "use_table": true, "display_type": "table", "pagination_mode": "server", "is_multi_row": false, "is_one_object": true, "form": {"name": "Purchase", "variables": [{"name": "supplier_id", "type": "select", "label": "Supplier", "data_source": {"type": "table", "table": "suppliers", "label_field": "name", "value_field": "id"}}, {"name": "amount", "type": "number", "label": "Amount"}]}}'),

-- BLOCK 3: SALES OPERATION
(301, 'customers', '{"process_id": 301, "form_name": "customers", "table_name": "customers", "use_table": true, "display_type": "table", "pagination_mode": "server", "is_multi_row": false, "is_one_object": true, "next_process_id": 302, "form": {"name": "Customer", "variables": [{"name": "full_name", "type": "text", "label": "Customer Name"}]}}'),
(302, 'orders', '{"process_id": 302, "form_name": "orders", "table_name": "orders", "use_table": true, "display_type": "table", "pagination_mode": "server", "is_multi_row": false, "is_one_object": true, "form": {"name": "Order", "variables": [{"name": "customer_id", "type": "select", "label": "Customer", "data_source": {"type": "table", "table": "customers", "label_field": "full_name", "value_field": "id"}}, {"name": "total", "type": "number", "label": "Total Price"}]}}');

INSERT INTO blockflows (id, name, start_process_id) VALUES
(1, 'Store Setup', 101),
(2, 'Procurement', 201),
(3, 'Sales Point', 301);

INSERT INTO flow_shops (flow_id, shop_id) VALUES 
(1, 1),
(2, 1),
(3, 1);



-- Categories
CREATE TABLE IF NOT EXISTS categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    shop_id TEXT NOT NULL,
    user_id TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);

-- Products
CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(150) NOT NULL,
    category_id INT REFERENCES categories(id),
    shop_id TEXT NOT NULL,
    user_id TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);

-- Suppliers
CREATE TABLE IF NOT EXISTS suppliers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(150) NOT NULL,
    shop_id TEXT NOT NULL,
    user_id TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);

-- Purchases
CREATE TABLE IF NOT EXISTS purchases (
    id SERIAL PRIMARY KEY,
    supplier_id INT REFERENCES suppliers(id),
    amount DECIMAL(10, 2) NOT NULL,
    shop_id TEXT NOT NULL,
    user_id TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);

-- Customers
CREATE TABLE IF NOT EXISTS customers (
    id SERIAL PRIMARY KEY,
    full_name VARCHAR(150) NOT NULL,
    shop_id TEXT NOT NULL,
    user_id TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);

-- Orders
CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY,
    customer_id INT REFERENCES customers(id),
    total DECIMAL(10, 2) NOT NULL,
    shop_id TEXT NOT NULL,
    user_id TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);