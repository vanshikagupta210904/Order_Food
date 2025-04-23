Create database Order_Food;
use Order_Food;

create table users(
user_id int auto_increment,
name varchar(255),
phone int8,
registration_date datetime,
last_login datetime,
preferred_address_id int,
user_role int,
primary key (user_id)
);

create table user_roles(
user_roles_id int auto_increment,
role varchar(63),
primary key (user_roles_id)    
);

create table address(
address_id int auto_increment,
user_id int,
latitude float8,
longitude float8,
house_no varchar(63),
area varchar(63),
pincode int,
contact_no int8,
label varchar(63),
primary key (address_id)
);

create table orders(
order_id int auto_increment,
user_id int,
store_id int,
delivery_agent_id int,
order_time datetime,
delivery_time datetime,
total_amount float4,
status varchar(63),
payment_id int,
primary key(order_id)
);

create table order_item(
order_item_id int auto_increment,
order_id int,
product_id int,
quantity int,
primary key(order_item_id)
);

create table category(
category_id int auto_increment,
name varchar(63),
parent_category_id int,
primary key(category_id)
);

create table product(
product_id int auto_increment,
category_id int,
name varchar(63),
price float4,
description text,
primary key(product_id)
);

create table delivery_agent(
delivery_agent_id int auto_increment,
name varchar(255),
phone int8,
status varchar(63),
primary key(delivery_agent_id)
);

create table payments(
payment_id int auto_increment,
order_id int,
payment_type int,
status varchar(63),
transaction_id varchar(255),
amount float4,
primary key (payment_id)
);

create table payment_method(
payment_method_id int auto_increment,
payment_type varchar(63),
primary key(payment_method_id)
);

create table store(
store_id int auto_increment,
operation_starts time,
operation_end time,
area varchar(63),
pincode int,
primary key (store_id)
);

create table service_area(
service_area_id int auto_increment,
store_id int,
area varchar(63),
pincode int,
primary key (service_area_id)
);

create table inventory(
inventory_id int auto_increment,
store_id int,
product_id int,
quantity int,
max_quantity int,
primary key (inventory_id)
);

ALTER TABLE address
ADD CONSTRAINT user_address
FOREIGN KEY (user_id)
REFERENCES users(user_id)
ON DELETE CASCADE;

ALTER TABLE users
ADD CONSTRAINT preferred_address
FOREIGN KEY (preferred_address_id)
REFERENCES address(address_id);

ALTER TABLE users
ADD CONSTRAINT user_rol
FOREIGN KEY (user_role)
REFERENCES user_roles(user_roles_id);

ALTER TABLE orders
ADD CONSTRAINT user_order
FOREIGN KEY (user_id)
REFERENCES users(user_id)
on delete cascade;

ALTER TABLE orders
ADD CONSTRAINT order_store
FOREIGN KEY (store_id)
REFERENCES store(store_id);

ALTER TABLE orders
ADD CONSTRAINT order_delivery_agent
FOREIGN KEY (delivery_agent_id)
REFERENCES delivery_agent(delivery_agent_id);

ALTER TABLE orders
ADD CONSTRAINT order_payment
FOREIGN KEY (payment_id)
REFERENCES payments(payment_id);

ALTER TABLE order_item
ADD CONSTRAINT order_item_id
FOREIGN KEY (order_id)
REFERENCES orders(order_id);

ALTER TABLE order_item
ADD CONSTRAINT orders_product
FOREIGN KEY (product_id)
REFERENCES product(product_id);

ALTER TABLE category
ADD CONSTRAINT parent_category
FOREIGN KEY (parent_category_id)
REFERENCES category(category_id);

ALTER TABLE product
ADD CONSTRAINT product_category
FOREIGN KEY (category_id)
REFERENCES category(category_id);

ALTER TABLE payments
ADD CONSTRAINT payment_order_link
FOREIGN KEY (order_id)
REFERENCES orders(order_id);

ALTER TABLE payments
ADD CONSTRAINT payment_method
FOREIGN KEY (payment_type)
REFERENCES payment_method(payment_method_id);

ALTER TABLE service_area
ADD CONSTRAINT store_service_area
FOREIGN KEY (store_id)
REFERENCES store(store_id);

ALTER TABLE inventory
ADD CONSTRAINT store_inventory
FOREIGN KEY (store_id)
REFERENCES store(store_id);

ALTER TABLE inventory
ADD CONSTRAINT inventory_products
FOREIGN KEY (product_id)
REFERENCES product(product_id);