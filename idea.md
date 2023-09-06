Creating a backend for an Amazon shopping clone is a complex task, and database design is a critical component of this process. Your approach to decrementing product counts on both the user and seller sides is a reasonable one, but you'll need to carefully design your database schema to support this functionality efficiently. Here's a basic outline of how you could structure your PostgreSQL database for this purpose:

1. **Users Table:**
   - Create a table to store user information, including details like username, email, password (hashed and salted), shipping addresses, and payment methods.
   - Include a unique user ID (primary key) for each user.

```sql
CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    -- other user-related fields
);
```

2. **Products Table:**
   - Create a table to store product information, such as product name, description, price, seller information (seller ID), and available quantity.

```sql
CREATE TABLE products (
    product_id SERIAL PRIMARY KEY,
    product_name VARCHAR(255) NOT NULL,
    description TEXT,
    price DECIMAL(10, 2) NOT NULL,
    seller_id INT NOT NULL,
    available_quantity INT NOT NULL,
    -- other product-related fields
);
```

3. **Orders Table:**
   - Create a table to represent orders. Each row in this table corresponds to a single order made by a user.
   - Include order details like user ID, order date, and status (e.g., "placed," "shipped," "delivered").

```sql
CREATE TABLE orders (
    order_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    order_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    status VARCHAR(255) NOT NULL,
    -- other order-related fields
);
```

4. **Order Items Table:**
   - Create a table to represent items within each order. This table will link products to specific orders and include information like quantity purchased and total price.

```sql
CREATE TABLE order_items (
    order_item_id SERIAL PRIMARY KEY,
    order_id INT NOT NULL,
    product_id INT NOT NULL,
    quantity INT NOT NULL,
    total_price DECIMAL(10, 2) NOT NULL,
    -- other order item-related fields
);
```

With this basic database structure, when a user makes a purchase:

- You can decrement the product's available quantity in the `products` table.
- Create a new record in the `orders` table with the order details and status.
- For each item in the user's cart, create a record in the `order_items` table linking it to the respective order.

Remember that this is a simplified outline, and a real-world application would likely require additional tables and more complex relationships to handle things like shipping, payments, and user reviews. Additionally, consider implementing database transactions to ensure data consistency when updating multiple tables during an order placement.