--
-- PostgreSQL database dump
--

-- Dumped from database version 15.3
-- Dumped by pg_dump version 15.3

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: balances; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.balances (
    id character(36) NOT NULL,
    account_number character varying(255) NOT NULL,
    account_name character varying(255) NOT NULL,
    bank_name character varying(255) NOT NULL,
    name character varying(255) NOT NULL,
    amount numeric(16,2) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.balances OWNER TO admin;

--
-- Name: cart_items; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.cart_items (
    id character(36) NOT NULL,
    cart_id character(36) NOT NULL,
    product_unit_id character(36) NOT NULL,
    qty numeric(16,2) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.cart_items OWNER TO admin;

--
-- Name: carts; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.carts (
    id character(36) NOT NULL,
    cashier_session_id character(36) NOT NULL,
    name character varying(255),
    is_active boolean NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.carts OWNER TO admin;

--
-- Name: cashier_sessions; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.cashier_sessions (
    id character(36) NOT NULL,
    user_id character(36) NOT NULL,
    status character varying(255) NOT NULL,
    starting_cash numeric(16,2) NOT NULL,
    ending_cash numeric(16,2),
    started_at timestamp without time zone NOT NULL,
    ended_at timestamp without time zone,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.cashier_sessions OWNER TO admin;

--
-- Name: customer_debts; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.customer_debts (
    id character(36) NOT NULL,
    customer_id character(36) NOT NULL,
    debt_source character varying(255) NOT NULL,
    debt_source_id character varying(255) NOT NULL,
    due_date date,
    status character varying(255) NOT NULL,
    amount numeric(16,2) NOT NULL,
    remaining_amount numeric(16,2) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.customer_debts OWNER TO admin;

--
-- Name: customer_payments; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.customer_payments (
    id character(36) NOT NULL,
    user_id character(36) NOT NULL,
    image_file_id character(36) NOT NULL,
    customer_debt_id character(36) NOT NULL,
    amount numeric(16,2) NOT NULL,
    description text,
    paid_at timestamp without time zone NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.customer_payments OWNER TO admin;

--
-- Name: customer_type_discounts; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.customer_type_discounts (
    id character(36) NOT NULL,
    product_id character(36) NOT NULL,
    customer_type_id character(36) NOT NULL,
    is_active boolean NOT NULL,
    discount_percentage numeric(16,2),
    discount_amount numeric(16,2),
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.customer_type_discounts OWNER TO admin;

--
-- Name: customer_types; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.customer_types (
    id character(36) NOT NULL,
    name character varying(255) NOT NULL,
    description text,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.customer_types OWNER TO admin;

--
-- Name: customers; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.customers (
    id character(36) NOT NULL,
    name character varying(255) NOT NULL,
    email character varying(255) NOT NULL,
    address text NOT NULL,
    phone character varying(20) NOT NULL,
    latitude double precision NOT NULL,
    longitude double precision NOT NULL,
    is_active boolean NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    customer_type_id character(36)
);


ALTER TABLE public.customers OWNER TO admin;

--
-- Name: debt_payments; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.debt_payments (
    id character(36) NOT NULL,
    user_id character(36) NOT NULL,
    image_file_id character(36) NOT NULL,
    debt_id character(36) NOT NULL,
    amount numeric(16,2) NOT NULL,
    description text,
    paid_at timestamp without time zone NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.debt_payments OWNER TO admin;

--
-- Name: debts; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.debts (
    id character(36) NOT NULL,
    debt_source character varying(255) NOT NULL,
    debt_source_identifier character varying(255) NOT NULL,
    due_date date,
    status character varying(255) NOT NULL,
    amount numeric(16,2) NOT NULL,
    remaining_amount numeric(16,2) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.debts OWNER TO admin;

--
-- Name: delivery_order_drivers; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.delivery_order_drivers (
    id character(36) NOT NULL,
    delivery_order_id character(36) NOT NULL,
    driver_user_id character(36) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.delivery_order_drivers OWNER TO admin;

--
-- Name: delivery_order_images; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.delivery_order_images (
    id character(36) NOT NULL,
    delivery_order_id character(36) NOT NULL,
    file_id character(36) NOT NULL,
    description character varying(255),
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.delivery_order_images OWNER TO admin;

--
-- Name: delivery_order_item_costs; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.delivery_order_item_costs (
    id character(36) NOT NULL,
    delivery_order_item_id character(36) NOT NULL,
    qty numeric(16,2) NOT NULL,
    base_cost_price numeric(16,2) NOT NULL,
    total_cost_price numeric(16,2) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.delivery_order_item_costs OWNER TO admin;

--
-- Name: delivery_order_items; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.delivery_order_items (
    id character(36) NOT NULL,
    delivery_order_id character(36) NOT NULL,
    product_unit_id character(36) NOT NULL,
    qty numeric(16,2) NOT NULL,
    scale_to_base numeric(16,2) NOT NULL,
    price_per_unit numeric(16,2) NOT NULL,
    discount_per_unit numeric(16,2) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.delivery_order_items OWNER TO admin;

--
-- Name: delivery_order_positions; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.delivery_order_positions (
    id character(36) NOT NULL,
    delivery_order_id character(36) NOT NULL,
    driver_user_id character(36) NOT NULL,
    latitude double precision NOT NULL,
    longitude double precision NOT NULL,
    bearing double precision NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.delivery_order_positions OWNER TO admin;

--
-- Name: delivery_order_return_images; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.delivery_order_return_images (
    id character(36) NOT NULL,
    delivery_order_return_id character(36) NOT NULL,
    file_id character(36) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.delivery_order_return_images OWNER TO admin;

--
-- Name: delivery_order_returns; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.delivery_order_returns (
    id character(36) NOT NULL,
    delivery_order_id character(36) NOT NULL,
    user_id character(36) NOT NULL,
    description text,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.delivery_order_returns OWNER TO admin;

--
-- Name: delivery_order_reviews; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.delivery_order_reviews (
    id character(36) NOT NULL,
    delivery_order_id character(36) NOT NULL,
    star_rating integer NOT NULL,
    description text,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    type character varying(255) DEFAULT ''::character varying NOT NULL
);


ALTER TABLE public.delivery_order_reviews OWNER TO admin;

--
-- Name: delivery_orders; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.delivery_orders (
    id character(36) NOT NULL,
    customer_id character(36) NOT NULL,
    user_id character(36) NOT NULL,
    invoice_number character varying(255) NOT NULL,
    date date NOT NULL,
    status character varying(255) NOT NULL,
    total_price numeric(16,2) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.delivery_orders OWNER TO admin;

--
-- Name: files; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.files (
    id character(36) NOT NULL,
    name character varying(255) NOT NULL,
    type character varying(255) NOT NULL,
    path text NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.files OWNER TO admin;

--
-- Name: permissions; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.permissions (
    id character(36) NOT NULL,
    title character varying(100) NOT NULL,
    description character varying(255) NOT NULL,
    is_active boolean NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.permissions OWNER TO admin;

--
-- Name: product_discounts; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.product_discounts (
    id character(36) NOT NULL,
    product_id character(36) NOT NULL,
    minimum_qty numeric(16,2) NOT NULL,
    is_active boolean NOT NULL,
    discount_percentage numeric(16,2),
    discount_amount numeric(16,2),
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.product_discounts OWNER TO admin;

--
-- Name: product_receive_images; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.product_receive_images (
    id character(36) NOT NULL,
    product_receive_id character(36) NOT NULL,
    file_id character(36) NOT NULL,
    description text,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.product_receive_images OWNER TO admin;

--
-- Name: product_receive_items; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.product_receive_items (
    id character(36) NOT NULL,
    product_receive_id character(36) NOT NULL,
    product_unit_id character(36) NOT NULL,
    user_id character(36) NOT NULL,
    qty_eligible numeric(16,2) NOT NULL,
    qty_received numeric(16,2) NOT NULL,
    scale_to_base numeric(16,2) NOT NULL,
    price_per_unit numeric(16,2) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.product_receive_items OWNER TO admin;

--
-- Name: product_receive_return_images; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.product_receive_return_images (
    id character(36) NOT NULL,
    product_receive_return_id character(36) NOT NULL,
    file_id character(36) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.product_receive_return_images OWNER TO admin;

--
-- Name: product_receive_returns; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.product_receive_returns (
    id character(36) NOT NULL,
    product_receive_id character(36) NOT NULL,
    user_id character(36) NOT NULL,
    description text,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.product_receive_returns OWNER TO admin;

--
-- Name: product_receives; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.product_receives (
    id character(36) NOT NULL,
    purchase_order_id character(36) NOT NULL,
    supplier_id character(36) NOT NULL,
    user_id character(36) NOT NULL,
    invoice_number character varying(255) NOT NULL,
    date date NOT NULL,
    status character varying(255) NOT NULL,
    total_price numeric(16,2) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.product_receives OWNER TO admin;

--
-- Name: product_return_images; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.product_return_images (
    id character(36) NOT NULL,
    product_return_id character(36) NOT NULL,
    file_id character(36) NOT NULL,
    description text,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.product_return_images OWNER TO admin;

--
-- Name: product_return_items; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.product_return_items (
    id character(36) NOT NULL,
    product_return_id character(36) NOT NULL,
    product_unit_id character(36) NOT NULL,
    qty numeric(16,2) NOT NULL,
    scale_to_base numeric(16,2) NOT NULL,
    base_cost_price numeric(16,2) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.product_return_items OWNER TO admin;

--
-- Name: product_returns; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.product_returns (
    id character(36) NOT NULL,
    supplier_id character(36) NOT NULL,
    user_id character(36) NOT NULL,
    invoice_number character varying(255) NOT NULL,
    date date NOT NULL,
    status character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.product_returns OWNER TO admin;

--
-- Name: product_stock_adjustments; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.product_stock_adjustments (
    id character(36) NOT NULL,
    user_id character(36) NOT NULL,
    product_stock_id character(36) NOT NULL,
    previous_qty numeric(16,2) NOT NULL,
    updated_qty numeric(16,2) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.product_stock_adjustments OWNER TO admin;

--
-- Name: product_stock_mutations; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.product_stock_mutations (
    id character(36) NOT NULL,
    product_unit_id character(36) NOT NULL,
    type character varying(255) NOT NULL,
    identifier_id character varying(255) NOT NULL,
    qty numeric(16,2) NOT NULL,
    scale_to_base numeric(16,2) NOT NULL,
    base_qty_left numeric(16,2) NOT NULL,
    base_cost_price numeric(16,2) NOT NULL,
    mutated_at timestamp without time zone NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.product_stock_mutations OWNER TO admin;

--
-- Name: product_stocks; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.product_stocks (
    id character(36) NOT NULL,
    product_id character(36) NOT NULL,
    qty numeric(16,2) NOT NULL,
    base_cost_price numeric(16,2) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.product_stocks OWNER TO admin;

--
-- Name: product_units; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.product_units (
    id character(36) NOT NULL,
    to_unit_id character(36),
    unit_id character(36) NOT NULL,
    product_id character(36) NOT NULL,
    scale numeric(16,2) NOT NULL,
    scale_to_base numeric(16,2) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.product_units OWNER TO admin;

--
-- Name: products; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.products (
    id character(36) NOT NULL,
    image_file_id character(36) NOT NULL,
    name character varying(255) NOT NULL,
    description text,
    price numeric(16,2),
    is_active boolean NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.products OWNER TO admin;

--
-- Name: purchase_order_images; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.purchase_order_images (
    id character(36) NOT NULL,
    purchase_order_id character(36) NOT NULL,
    file_id character(36) NOT NULL,
    description text,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.purchase_order_images OWNER TO admin;

--
-- Name: purchase_order_items; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.purchase_order_items (
    id character(36) NOT NULL,
    purchase_order_id character(36) NOT NULL,
    product_unit_id character(36) NOT NULL,
    user_id character(36) NOT NULL,
    qty numeric(16,2) NOT NULL,
    scale_to_base numeric(16,2) NOT NULL,
    price_per_unit numeric(16,2) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.purchase_order_items OWNER TO admin;

--
-- Name: purchase_orders; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.purchase_orders (
    id character(36) NOT NULL,
    supplier_id character(36) NOT NULL,
    user_id character(36) NOT NULL,
    invoice_number character varying(255) NOT NULL,
    date date NOT NULL,
    status character varying(255) NOT NULL,
    total_estimated_price numeric(16,2) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.purchase_orders OWNER TO admin;

--
-- Name: role_permissions; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.role_permissions (
    role_id character(36) NOT NULL,
    permission_id character(36) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.role_permissions OWNER TO admin;

--
-- Name: roles; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.roles (
    id character(36) NOT NULL,
    name character varying(255) NOT NULL,
    description text,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);


ALTER TABLE public.roles OWNER TO admin;

--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


ALTER TABLE public.schema_migrations OWNER TO admin;

--
-- Name: sequences; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.sequences (
    id character(36) NOT NULL,
    unique_identifier character varying(255) NOT NULL,
    sequence integer NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.sequences OWNER TO admin;

--
-- Name: shop_order_items; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.shop_order_items (
    id character(36) NOT NULL,
    shop_order_id character(36) NOT NULL,
    product_unit_id character(36) NOT NULL,
    platform_product_id character varying(255) NOT NULL,
    image_link character varying(255),
    quantity numeric(16,2) NOT NULL,
    original_price numeric(16,2) NOT NULL,
    sale_price numeric(16,2) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.shop_order_items OWNER TO admin;

--
-- Name: shop_orders; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.shop_orders (
    id character(36) NOT NULL,
    tracking_number character varying(255),
    platform_identifier character varying(255),
    platform_type character varying(255) NOT NULL,
    tracking_status character varying(255) NOT NULL,
    recipient_name character varying(255) NOT NULL,
    recipient_full_address text NOT NULL,
    recipient_phone_number character varying(20) NOT NULL,
    shipping_fee numeric(16,2) NOT NULL,
    service_fee numeric(16,2) NOT NULL,
    total_original_product_price numeric(16,2) NOT NULL,
    subtotal numeric(16,2) NOT NULL,
    tax numeric(16,2) NOT NULL,
    total_amount numeric(16,2) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.shop_orders OWNER TO admin;

--
-- Name: shopee_configs; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.shopee_configs (
    partner_id character varying(255) NOT NULL,
    partner_key character varying(255) NOT NULL,
    access_token character varying(255),
    refresh_token character varying(255),
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.shopee_configs OWNER TO admin;

--
-- Name: supplier_types; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.supplier_types (
    id character(36) NOT NULL,
    name character varying(255) NOT NULL,
    description text,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.supplier_types OWNER TO admin;

--
-- Name: suppliers; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.suppliers (
    id character(36) NOT NULL,
    supplier_type_id character(36) NOT NULL,
    code character varying(255) NOT NULL,
    name character varying(255) NOT NULL,
    is_active boolean NOT NULL,
    address text NOT NULL,
    phone character varying(20) NOT NULL,
    email character varying(255),
    description text,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.suppliers OWNER TO admin;

--
-- Name: tiktok_configs; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.tiktok_configs (
    app_key character varying(255) NOT NULL,
    app_secret character varying(255) NOT NULL,
    warehouse_id character varying(255) NOT NULL,
    shop_id character varying(255) NOT NULL,
    shop_cipher character varying(255) NOT NULL,
    access_token character varying(255),
    refresh_token character varying(255),
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.tiktok_configs OWNER TO admin;

--
-- Name: tiktok_products; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.tiktok_products (
    tiktok_product_id character varying(255) NOT NULL,
    product_id character(36) NOT NULL,
    status character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.tiktok_products OWNER TO admin;

--
-- Name: transaction_item_costs; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.transaction_item_costs (
    id character(36) NOT NULL,
    transaction_item_id character(36) NOT NULL,
    qty numeric(16,2) NOT NULL,
    base_cost_price numeric(16,2) NOT NULL,
    total_cost_price numeric(16,2) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.transaction_item_costs OWNER TO admin;

--
-- Name: transaction_items; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.transaction_items (
    id character(36) NOT NULL,
    transaction_id character(36) NOT NULL,
    product_unit_id character(36) NOT NULL,
    qty numeric(16,2) NOT NULL,
    price_per_unit numeric(16,2) NOT NULL,
    discount_per_unit numeric(16,2),
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.transaction_items OWNER TO admin;

--
-- Name: transaction_payments; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.transaction_payments (
    id character(36) NOT NULL,
    transaction_id character(36) NOT NULL,
    payment_type character varying(255) NOT NULL,
    reference_number character varying(255),
    total numeric(16,2) NOT NULL,
    total_paid numeric(16,2) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.transaction_payments OWNER TO admin;

--
-- Name: transactions; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.transactions (
    id character(36) NOT NULL,
    cashier_session_id character(36) NOT NULL,
    status character varying(255) NOT NULL,
    total numeric(16,2) NOT NULL,
    payment_at timestamp without time zone,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.transactions OWNER TO admin;

--
-- Name: units; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.units (
    id character(36) NOT NULL,
    name character varying(255) NOT NULL,
    description text,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.units OWNER TO admin;

--
-- Name: user_access_tokens; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.user_access_tokens (
    id character(100) NOT NULL,
    user_id character(36) NOT NULL,
    revoked boolean NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    expired_at timestamp without time zone NOT NULL
);


ALTER TABLE public.user_access_tokens OWNER TO admin;

--
-- Name: user_roles; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.user_roles (
    user_id character(36) NOT NULL,
    role_id character(36) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.user_roles OWNER TO admin;

--
-- Name: users; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.users (
    id character(36) NOT NULL,
    username character varying(255) NOT NULL,
    name character varying(255) NOT NULL,
    password character varying(255) NOT NULL,
    is_active boolean NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.users OWNER TO admin;

--
-- Data for Name: balances; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.balances (id, account_number, account_name, bank_name, name, amount, created_at, updated_at) FROM stdin;
20d4ee82-525e-42a2-8c6e-1f5d1339509c	18923091823	NAME BCA	BCA Name	BCA	0.00	2024-06-12 03:53:00.345693	2024-06-12 03:53:00.345693
\.


--
-- Data for Name: cart_items; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.cart_items (id, cart_id, product_unit_id, qty, created_at, updated_at) FROM stdin;
\.


--
-- Data for Name: carts; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.carts (id, cashier_session_id, name, is_active, created_at, updated_at) FROM stdin;
\.


--
-- Data for Name: cashier_sessions; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.cashier_sessions (id, user_id, status, starting_cash, ending_cash, started_at, ended_at, created_at, updated_at) FROM stdin;
b0ee624b-040c-41f9-8d21-945db1adc44f	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	COMPLETED	10000.00	200000.00	2024-06-07 15:51:12.594263	2024-06-12 04:01:48.050689	2024-06-07 15:51:12.595043	2024-06-12 04:01:48.053904
bfd60333-4c6b-48ee-933a-31810f1381c9	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	COMPLETED	1000000.00	20000.00	2024-06-18 06:26:24.251458	2024-06-21 09:56:15.184713	2024-06-18 06:26:24.252315	2024-06-21 09:56:15.185824
d737f625-5234-4008-a9ef-49381c8f5a13	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	ACTIVE	100000.00	\N	2024-06-21 09:56:24.456299	\N	2024-06-21 09:56:24.457354	2024-06-21 09:56:24.457354
553c817c-8066-4218-b826-135610eb86b1	68ed7124-fa08-4720-b741-9fe4fa697c21	COMPLETED	10000.00	20000.00	2024-06-27 07:37:14.546153	2024-06-27 08:46:01.130387	2024-06-27 07:37:14.548319	2024-06-27 08:46:01.13171
3705dfe3-4097-4c0b-b888-895436467d09	68ed7124-fa08-4720-b741-9fe4fa697c21	COMPLETED	100000.00	340000.00	2024-06-27 09:27:45.028876	2024-06-27 09:29:11.367714	2024-06-27 09:27:45.029886	2024-06-27 09:29:11.36908
f24c78ab-23cf-4da7-9350-3bfe1716e3cf	68ed7124-fa08-4720-b741-9fe4fa697c21	COMPLETED	10000.00	10000.00	2024-06-27 09:29:38.289748	2024-06-28 00:32:44.787726	2024-06-27 09:29:38.290331	2024-06-28 00:32:44.789025
fc192a4e-4a31-42c6-bf2e-ba06b5f54509	68ed7124-fa08-4720-b741-9fe4fa697c21	COMPLETED	1000000.00	1300000.00	2024-06-28 00:33:14.653275	2024-06-28 00:34:19.011585	2024-06-28 00:33:14.654195	2024-06-28 00:34:19.01282
e2b2a92b-4684-453e-b224-329185eed415	68ed7124-fa08-4720-b741-9fe4fa697c21	COMPLETED	10000.00	100000.00	2024-06-28 02:34:54.251749	2024-06-29 09:45:50.097295	2024-06-28 02:34:54.252958	2024-06-29 09:45:50.099705
7da6cd43-122c-462c-bf01-3cf4f6dfcede	16638b6a-5ca3-4d1f-ab0f-ac4ce4ce30a0	ACTIVE	1000000.00	\N	2024-06-29 09:47:01.343013	\N	2024-06-29 09:47:01.344302	2024-06-29 09:47:01.344302
db53eb35-1c2c-48a8-84d1-6eca34b80528	68ed7124-fa08-4720-b741-9fe4fa697c21	COMPLETED	1000.00	1000.00	2024-07-01 02:54:12.888895	2024-07-01 03:07:59.810155	2024-07-01 02:54:12.889805	2024-07-01 03:07:59.822605
0a3376e6-9176-4850-b589-9ad812b03985	68ed7124-fa08-4720-b741-9fe4fa697c21	COMPLETED	10000.00	110000.00	2024-07-01 03:08:10.316373	2024-07-01 03:08:21.019852	2024-07-01 03:08:10.317225	2024-07-01 03:08:21.024005
dd5e607e-4c1f-4ca4-8fab-e2f51a8fe6c7	68ed7124-fa08-4720-b741-9fe4fa697c21	COMPLETED	10000.00	10000.00	2024-07-01 03:42:35.898642	2024-07-01 03:42:38.113373	2024-07-01 03:42:35.899557	2024-07-01 03:42:38.11578
f05792d7-1552-451b-8821-92cbe563f82c	33392eb4-0f87-43c6-9893-9c014fe6d561	COMPLETED	100000.00	100000.00	2024-07-08 15:54:27.439563	2024-07-08 15:54:34.060722	2024-07-08 15:54:27.44063	2024-07-08 15:54:34.067781
62ca2919-3d31-45a7-a42a-1dbdfd6f43da	33392eb4-0f87-43c6-9893-9c014fe6d561	COMPLETED	100000.00	690000.00	2024-07-08 15:54:44.461236	2024-07-08 15:55:14.216474	2024-07-08 15:54:44.462033	2024-07-08 15:55:14.21865
ba5bf7c4-8b51-49ad-a3cc-2d2d70d791df	33392eb4-0f87-43c6-9893-9c014fe6d561	COMPLETED	100000.00	110000.00	2024-07-08 16:02:15.412511	2024-07-08 16:02:40.773634	2024-07-08 16:02:15.41296	2024-07-08 16:02:40.777734
4d55ec36-e9fc-4b9e-99a1-6aea4ac3afd3	33392eb4-0f87-43c6-9893-9c014fe6d561	COMPLETED	0.00	10000.00	2024-07-08 16:03:32.571861	2024-07-08 16:03:49.413628	2024-07-08 16:03:32.572593	2024-07-08 16:03:49.415338
afd5d5ab-4b0f-4789-a460-a6c8ee228c98	68ed7124-fa08-4720-b741-9fe4fa697c21	COMPLETED	10000.00	210000.00	2024-07-01 09:20:51.679572	2024-07-10 02:16:30.272147	2024-07-01 09:20:51.680914	2024-07-10 02:16:30.278936
a4ede830-480d-438c-971c-4a0939c8345e	68ed7124-fa08-4720-b741-9fe4fa697c21	COMPLETED	100000.00	112000.00	2024-07-10 02:16:34.0805	2024-07-10 02:16:46.667832	2024-07-10 02:16:34.081298	2024-07-10 02:16:46.671546
bfc6cfb1-7113-4148-adbd-0c397b9adaff	68ed7124-fa08-4720-b741-9fe4fa697c21	COMPLETED	100000.00	112000.00	2024-07-10 02:19:04.276052	2024-07-10 02:19:15.603512	2024-07-10 02:19:04.276784	2024-07-10 02:19:15.605558
\.


--
-- Data for Name: customer_debts; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.customer_debts (id, customer_id, debt_source, debt_source_id, due_date, status, amount, remaining_amount, created_at, updated_at) FROM stdin;
d2e3019c-8100-45d0-82c0-8cbb60669c61	f74d78f8-26ee-4699-89da-134cfbcda903	DELIVERY_ORDER	07f95170-222d-4d10-a904-e5f57285fca9	\N	UNPAID	1500000.00	1000000.00	2024-06-07 16:09:40.2597	2024-06-07 16:10:31.5525
de964676-5ff1-45cd-a7e7-430898798f7d	f74d78f8-26ee-4699-89da-134cfbcda903	DELIVERY_ORDER	85dc9c49-c02b-4305-a5ce-de4010e03520	\N	UNPAID	650000.00	650000.00	2024-06-18 06:50:45.708827	2024-06-18 06:50:45.708827
5371202f-7b92-49ee-b5ca-5a2bc72f1023	f74d78f8-26ee-4699-89da-134cfbcda903	DELIVERY_ORDER	3545b41a-cfc4-4126-9b82-c241acaad4b4	\N	UNPAID	3120000.00	3120000.00	2024-06-18 06:59:49.054657	2024-06-18 06:59:49.054657
1ce11a1e-10bf-449b-9f91-5cd33d6a68b5	f74d78f8-26ee-4699-89da-134cfbcda903	DELIVERY_ORDER	9b3af119-e11e-4d69-ad93-331f0eb1e8bf	\N	UNPAID	1820000.00	1820000.00	2024-06-18 07:00:12.586758	2024-06-18 07:00:12.586758
5740dc8f-138e-4c55-b395-fdab208c4405	f74d78f8-26ee-4699-89da-134cfbcda903	DELIVERY_ORDER	516ff344-743d-4f6f-9e2d-c617ffc54402	\N	UNPAID	1560000.00	1560000.00	2024-06-18 07:18:19.122925	2024-06-18 07:18:19.122925
51e162ad-c3ca-4b11-b74e-c509a5b3eeb1	f74d78f8-26ee-4699-89da-134cfbcda903	DELIVERY_ORDER	726acba1-51d3-4cee-aea0-377c8c43d4b6	\N	CANCELED	1300000.00	1300000.00	2024-06-20 02:22:46.920988	2024-06-20 03:59:35.249155
cb9afa95-41a0-491b-b027-a7b8ec6e404c	f74d78f8-26ee-4699-89da-134cfbcda903	DELIVERY_ORDER	ebd6eba0-ea19-418a-ba48-0704dd984bcd	\N	CANCELED	1300000.00	1300000.00	2024-06-20 09:46:58.230369	2024-06-20 09:49:27.124695
05d984fa-208e-4cf0-8d3a-56ebcf18d91e	f74d78f8-26ee-4699-89da-134cfbcda903	DELIVERY_ORDER	1b2731a7-f231-40f4-97e3-0454abadb848	\N	PAID	260000.00	0.00	2024-06-18 07:18:35.323416	2024-06-27 08:50:48.075019
0146bd08-93f2-430c-8bfc-35c45e353336	77ce2ab4-a3e3-4d58-bd6c-d81a364a938c	DELIVERY_ORDER	4a99a612-d372-40c9-9a75-555b7fef8136	\N	PAID	8000000.00	0.00	2024-06-27 09:23:18.288047	2024-06-28 00:19:10.566048
30fe6bbb-6288-4109-aa79-98896e730b1e	77ce2ab4-a3e3-4d58-bd6c-d81a364a938c	DELIVERY_ORDER	3bcf5adc-492b-432c-ac26-df68f6c44a3b	\N	PAID	950000.00	0.00	2024-06-29 09:33:20.889518	2024-07-01 11:42:03.945069
59c67533-2b9a-4921-a691-bfc259c59449	e26f6d2b-553c-4a4a-8618-e1af14087a89	DELIVERY_ORDER	90bd19ca-d557-4bc1-bd8b-58ae4ccf736a	\N	HALF_PAID	2000000.00	1480000.00	2024-06-28 00:29:08.969151	2024-07-02 07:38:42.042979
a2f42740-fee7-44e2-a096-2304a1747bef	77ce2ab4-a3e3-4d58-bd6c-d81a364a938c	DELIVERY_ORDER	954d09b1-fc60-46f0-af64-fb7a47d461d1	\N	HALF_PAID	4750000.00	4250000.00	2024-07-08 12:29:19.452936	2024-07-10 15:32:08.80792
\.


--
-- Data for Name: customer_payments; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.customer_payments (id, user_id, image_file_id, customer_debt_id, amount, description, paid_at, created_at, updated_at) FROM stdin;
c494b999-334b-445a-a5f8-4989d68b714b	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	cec57216-25dd-412c-bf70-9fe54e33c11a	d2e3019c-8100-45d0-82c0-8cbb60669c61	500000.00	\N	2024-06-07 16:10:31.546211	2024-06-07 16:10:31.550632	2024-06-07 16:10:31.550632
a88b6293-ec2b-4b67-9a8d-96b42a00a9e4	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	3bc20135-def1-44c5-8e6a-d104172bfdf3	05d984fa-208e-4cf0-8d3a-56ebcf18d91e	200000.00	\N	2024-06-27 08:50:39.282437	2024-06-27 08:50:39.290259	2024-06-27 08:50:39.290259
c0da3d58-3297-4c4e-b398-93d6d1c2ed71	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	38a4bd2f-2b31-4067-a355-2c924778f6d3	05d984fa-208e-4cf0-8d3a-56ebcf18d91e	60000.00	\N	2024-06-27 08:50:48.067108	2024-06-27 08:50:48.073986	2024-06-27 08:50:48.073986
42b428e6-ce62-4714-93ea-22b9fd930318	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	0e88c71a-e7c3-4b99-bd57-de8ef5220f3f	0146bd08-93f2-430c-8bfc-35c45e353336	6000000.00	\N	2024-06-28 00:18:58.86771	2024-06-28 00:18:58.874355	2024-06-28 00:18:58.874355
41de6446-f527-4154-9ca8-442499019b7a	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	a770c8cf-71f8-4f4e-b4ce-7e61436e803c	0146bd08-93f2-430c-8bfc-35c45e353336	2000000.00	\N	2024-06-28 00:19:10.562812	2024-06-28 00:19:10.565231	2024-06-28 00:19:10.565231
f9c98431-1477-40f4-91b4-c31b03aadefe	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	26249674-49b6-42c9-8412-4ff1b30e4552	30fe6bbb-6288-4109-aa79-98896e730b1e	200000.00	\N	2024-07-01 11:40:17.406337	2024-07-01 11:40:17.414508	2024-07-01 11:40:17.414508
ccd94971-090e-4c11-b902-2ef107d9f528	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	e199a34b-54a7-4c0c-a1be-8a5d2eb33eb2	30fe6bbb-6288-4109-aa79-98896e730b1e	750000.00	\N	2024-07-01 11:42:03.93657	2024-07-01 11:42:03.944264	2024-07-01 11:42:03.944264
d0acbcf7-9b87-4478-86c8-e4775aaef26f	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	fa26da76-a459-47fc-a405-34477fa491f0	59c67533-2b9a-4921-a691-bfc259c59449	500000.00	\N	2024-07-01 11:42:21.02867	2024-07-01 11:42:21.035128	2024-07-01 11:42:21.035128
f53ca92b-a0d8-41af-9006-1b82deee4368	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	9373fcd0-92d9-48e9-b0c9-d0737dc3341a	59c67533-2b9a-4921-a691-bfc259c59449	20000.00	\N	2024-07-02 07:38:42.035319	2024-07-02 07:38:42.040135	2024-07-02 07:38:42.040135
978ad65b-419f-4166-ad4c-d2cacbfeb0ee	33392eb4-0f87-43c6-9893-9c014fe6d561	7957cf37-7c5a-4b32-a1e7-2d0d4e056c47	a2f42740-fee7-44e2-a096-2304a1747bef	500000.00	Pelunasan Pertama 	2024-07-10 15:32:08.799717	2024-07-10 15:32:08.804561	2024-07-10 15:32:08.804561
\.


--
-- Data for Name: customer_type_discounts; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.customer_type_discounts (id, product_id, customer_type_id, is_active, discount_percentage, discount_amount, created_at, updated_at) FROM stdin;
bf8d17c6-d0f7-4068-8510-8ac052326794	e1bf0592-7850-4602-a740-6aae98dfd281	717244a6-8318-49ca-b61b-f8cb1e58a63d	t	\N	20000.00	2024-06-12 02:26:44.317759	2024-06-12 02:26:44.317759
e17ce2b8-7eb1-446c-bb92-097364ee1481	e1bf0592-7850-4602-a740-6aae98dfd281	58c3fa57-7850-4043-b430-665645995425	t	10.00	\N	2024-06-18 06:23:50.77092	2024-06-18 06:23:50.77092
3f211ef4-7421-422b-8f0a-34e80575d679	e1bf0592-7850-4602-a740-6aae98dfd281	04defde4-03c1-4025-a3aa-1618f454d26f	t	\N	1000.00	2024-06-27 08:48:07.456616	2024-06-27 08:48:07.456616
0051dd5d-75ac-4104-a467-06f7864d6a81	d0813034-19a1-406f-af25-c18d2f301614	04defde4-03c1-4025-a3aa-1618f454d26f	t	\N	5000.00	2024-06-29 09:31:26.247744	2024-06-29 09:31:26.247744
1a00b564-5153-4778-90ad-6169265d9b1e	d0813034-19a1-406f-af25-c18d2f301614	4dc84be2-e85e-4d92-bfe1-4cb8caddb094	t	\N	123.00	2024-07-02 09:38:01.162967	2024-07-02 09:38:01.162967
\.


--
-- Data for Name: customer_types; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.customer_types (id, name, description, created_at, updated_at) FROM stdin;
717244a6-8318-49ca-b61b-f8cb1e58a63d	Langganan Tier I	\N	2024-06-06 09:29:38.594459	2024-06-06 09:29:38.594459
58c3fa57-7850-4043-b430-665645995425	s	d	2024-06-18 06:23:39.476348	2024-06-18 06:23:39.476348
04defde4-03c1-4025-a3aa-1618f454d26f	Tier 1	\N	2024-06-27 08:46:52.87047	2024-06-27 08:46:52.87047
4dc84be2-e85e-4d92-bfe1-4cb8caddb094	Tier II	\N	2024-06-28 00:17:00.292713	2024-06-28 00:17:00.292713
098a6e7e-2652-498c-8221-34f22164146c	1	\N	2024-07-03 03:21:09.981579	2024-07-03 03:21:09.981579
e0ab38ac-b9d4-4206-909e-770b7bf06d0f	2	\N	2024-07-03 03:21:14.764447	2024-07-03 03:21:14.764447
76de837c-3cec-4f4f-9849-acfe7f64259f	3	\N	2024-07-03 03:21:17.452223	2024-07-03 03:21:17.452223
5062a1a9-2346-4b20-80a5-df5f43e859e6	4	\N	2024-07-03 03:21:19.979681	2024-07-03 03:21:19.979681
a2728a02-a768-4b6b-ac80-4776908aa073	5	\N	2024-07-03 03:21:22.420344	2024-07-03 03:21:22.420344
09535714-982e-4856-896d-d645699d0fea	6	\N	2024-07-03 03:21:25.360862	2024-07-03 03:21:25.360862
0a5dc240-f2cc-4a51-bfad-95ce3ba791d4	7	\N	2024-07-03 03:21:28.672377	2024-07-03 03:21:28.672377
12be0913-304a-4507-a1c8-cd6b75728e06	8	\N	2024-07-03 03:21:32.039296	2024-07-03 03:21:32.039296
a6327d82-53e4-40d0-ab27-476f28572e0c	9	\N	2024-07-03 03:21:34.747095	2024-07-03 03:21:34.747095
ec87b147-0e99-44d6-a55b-aa20c047764b	10	\N	2024-07-03 03:21:37.978483	2024-07-03 03:21:37.978483
\.


--
-- Data for Name: customers; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.customers (id, name, email, address, phone, latitude, longitude, is_active, created_at, updated_at, customer_type_id) FROM stdin;
f74d78f8-26ee-4699-89da-134cfbcda903	Bobby Doe	bobby@gmail.com	Jln. Mahkamah	+6283163573103	3.574727	98.688233	f	2024-06-06 09:29:38.598088	2024-07-08 12:15:15.864554	717244a6-8318-49ca-b61b-f8cb1e58a63d
34e44a7e-e864-43bb-a352-e7282f463c27	Test	jitufarhan2@gmail.com	Jl. Test	+62810923801928	-6.256602054794918	106.8706226348877	f	2024-06-28 03:35:25.016176	2024-07-08 12:15:25.293329	4dc84be2-e85e-4d92-bfe1-4cb8caddb094
77ce2ab4-a3e3-4d58-bd6c-d81a364a938c	John Doe (Toko)	johndoe@gmail.com	Jln. Tasbih	+6281264721516	3.5796196005647403	98.68364096117416	t	2024-06-27 08:48:55.335698	2024-07-08 12:15:54.865798	04defde4-03c1-4025-a3aa-1618f454d26f
e26f6d2b-553c-4a4a-8618-e1af14087a89	Jane Doe	jane.doe@gmail.com	Jln. Tasbih	+6285261302277	3.579304	98.683555	t	2024-06-28 00:17:31.082832	2024-07-08 12:26:55.530378	4dc84be2-e85e-4d92-bfe1-4cb8caddb094
\.


--
-- Data for Name: debt_payments; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.debt_payments (id, user_id, image_file_id, debt_id, amount, description, paid_at, created_at, updated_at) FROM stdin;
6613d815-ff68-4a7c-8f5b-a21b045bb1ba	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	c6bd93ac-0bba-4b43-94a9-e96f6184d2e7	8e8dcd18-0ac4-4c39-8b89-dea640c164cb	1000000.00	\N	2024-06-27 08:51:20.370845	2024-06-27 08:51:20.380974	2024-06-27 08:51:20.380974
6407bddd-1794-434e-8d6a-7f83e8644500	33392eb4-0f87-43c6-9893-9c014fe6d561	10f6ef3c-85ab-46e7-8be0-a0b6e0a6379e	1b73fc6b-72fb-4039-aa02-3c615156e54b	250000.00	Pelunasn Pertama	2024-07-10 15:33:13.503636	2024-07-10 15:33:13.512858	2024-07-10 15:33:13.512858
\.


--
-- Data for Name: debts; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.debts (id, debt_source, debt_source_identifier, due_date, status, amount, remaining_amount, created_at, updated_at) FROM stdin;
2decf0e7-76ed-4096-9856-d6b1325a95d2	PRODUCT_RECEIVE	12450cfa-19be-4272-ac6c-1389b98e79da	\N	UNPAID	120000.00	120000.00	2024-06-07 15:45:34.227528	2024-06-07 15:45:34.227528
8e8dcd18-0ac4-4c39-8b89-dea640c164cb	PRODUCT_RECEIVE	ea86af8e-c1a8-41eb-9028-dd87eb92a405	\N	PAID	1000000.00	0.00	2024-06-20 02:22:16.372611	2024-06-27 08:51:20.384338
7cc311d7-2687-43de-8bfb-1af5581e610a	PRODUCT_RECEIVE	d8984d7a-307e-46d7-b8ae-bd429cc23784	\N	UNPAID	860000.00	860000.00	2024-06-27 08:59:44.329297	2024-06-27 08:59:44.329297
e9124890-0b68-436a-881e-3c61fb1859ba	PRODUCT_RECEIVE	c7bc1479-51e8-462c-9b30-0824a574b17a	\N	UNPAID	7800000.00	7800000.00	2024-06-27 09:21:13.520736	2024-06-27 09:21:13.520736
4f626b5c-da44-4463-8c30-8d4773194307	PRODUCT_RECEIVE	204b6093-103c-46da-9bf9-a07359e7c158	\N	UNPAID	7800000.00	7800000.00	2024-06-28 00:26:44.714525	2024-06-28 00:26:44.714525
d8a0fdb9-db51-422e-9017-b29cd835e99e	PRODUCT_RECEIVE	824e5e08-151f-4143-a5b7-70e63d96acae	\N	UNPAID	5000000.00	5000000.00	2024-07-01 11:11:30.170023	2024-07-01 11:11:30.170023
cdc529bd-1d50-4d57-9ec7-e36cf9664197	PRODUCT_RECEIVE	1433e671-a380-4d75-ac1c-7aedda724f93	\N	UNPAID	1200000.00	1200000.00	2024-07-10 02:04:45.294027	2024-07-10 02:04:45.294027
1e3011ce-4d15-47f1-af7e-62e0f75af4be	PRODUCT_RECEIVE	969955f9-5964-4e38-a05c-4033d02f1f12	\N	UNPAID	1200000.00	1200000.00	2024-07-10 02:18:50.204665	2024-07-10 02:18:50.204665
1b73fc6b-72fb-4039-aa02-3c615156e54b	PRODUCT_RECEIVE	7ad0043b-a3dc-46b8-8e9e-14aded64b9f6	\N	HALF_PAID	300000.00	50000.00	2024-07-10 02:55:58.49184	2024-07-10 15:33:13.515932
\.


--
-- Data for Name: delivery_order_drivers; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.delivery_order_drivers (id, delivery_order_id, driver_user_id, created_at, updated_at) FROM stdin;
8a6a218a-487a-48de-9de0-af4c16d7d1e2	07f95170-222d-4d10-a904-e5f57285fca9	fbfcbc34-77b9-4901-82f1-e5fa78d5aa48	2024-06-07 15:45:49.557983	2024-06-07 15:45:49.557983
1fad1195-acef-4c75-abb6-05692d04aa59	85dc9c49-c02b-4305-a5ce-de4010e03520	fbfcbc34-77b9-4901-82f1-e5fa78d5aa48	2024-06-18 06:50:29.691006	2024-06-18 06:50:29.691006
30dcebc9-fcc4-4d44-8d68-092d7b2c8e35	3545b41a-cfc4-4126-9b82-c241acaad4b4	fbfcbc34-77b9-4901-82f1-e5fa78d5aa48	2024-06-18 06:59:39.442164	2024-06-18 06:59:39.442164
59b2d4d4-1d0b-4c36-b367-77865a9de1d0	9b3af119-e11e-4d69-ad93-331f0eb1e8bf	fbfcbc34-77b9-4901-82f1-e5fa78d5aa48	2024-06-18 07:00:00.486686	2024-06-18 07:00:00.486686
54135968-78cf-4b17-9279-0ec3e18db43a	516ff344-743d-4f6f-9e2d-c617ffc54402	fbfcbc34-77b9-4901-82f1-e5fa78d5aa48	2024-06-18 07:18:11.903411	2024-06-18 07:18:11.903411
31bf7b91-bbef-47f9-a856-69ab3d4fb955	1b2731a7-f231-40f4-97e3-0454abadb848	fbfcbc34-77b9-4901-82f1-e5fa78d5aa48	2024-06-18 07:18:28.731673	2024-06-18 07:18:28.731673
361c04b6-2d83-4d04-aa09-248258a7e6bd	726acba1-51d3-4cee-aea0-377c8c43d4b6	fbfcbc34-77b9-4901-82f1-e5fa78d5aa48	2024-06-20 02:22:40.021518	2024-06-20 02:22:40.021518
6098c2c8-9f04-4637-b751-ccc972d3f441	ebd6eba0-ea19-418a-ba48-0704dd984bcd	fbfcbc34-77b9-4901-82f1-e5fa78d5aa48	2024-06-20 09:46:27.155683	2024-06-20 09:46:27.155683
823a25c0-18fd-4fc5-b7fe-573e3a1ae24a	4a99a612-d372-40c9-9a75-555b7fef8136	fbfcbc34-77b9-4901-82f1-e5fa78d5aa48	2024-06-27 09:22:46.26261	2024-06-27 09:22:46.26261
ec3c2d47-0125-4700-b667-51ed7bcdb4f9	90bd19ca-d557-4bc1-bd8b-58ae4ccf736a	fbfcbc34-77b9-4901-82f1-e5fa78d5aa48	2024-06-28 00:28:29.495114	2024-06-28 00:28:29.495114
d2e024ed-49e3-4bf6-a242-2526c0983f49	3bcf5adc-492b-432c-ac26-df68f6c44a3b	fbfcbc34-77b9-4901-82f1-e5fa78d5aa48	2024-06-29 09:32:32.781593	2024-06-29 09:32:32.781593
5a8cd62d-7693-4b4b-b7e2-36a52d6c8a1b	954d09b1-fc60-46f0-af64-fb7a47d461d1	fbfcbc34-77b9-4901-82f1-e5fa78d5aa48	2024-07-08 12:28:48.459348	2024-07-08 12:28:48.459348
\.


--
-- Data for Name: delivery_order_images; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.delivery_order_images (id, delivery_order_id, file_id, description, created_at, updated_at) FROM stdin;
d0957f4b-0e47-4c82-83e4-088ffb72c43f	07f95170-222d-4d10-a904-e5f57285fca9	851a94ea-bb56-473f-b36d-73b3e49c3225	\N	2024-06-07 15:49:50.468844	2024-06-07 15:49:50.468844
6bd1baff-bab6-4d5b-a5b3-a847c6d0007f	4a99a612-d372-40c9-9a75-555b7fef8136	e1b37960-ad21-4978-bc56-0d0d0c0a1ace	\N	2024-06-27 09:23:13.086904	2024-06-27 09:23:13.086904
662ec435-cbf5-4a2e-8b46-065a40b032c6	3bcf5adc-492b-432c-ac26-df68f6c44a3b	c29d71c5-75b8-4300-a6d8-fa1cf938d42f	jhhjhu	2024-06-29 09:33:12.979138	2024-06-29 09:33:12.979138
\.


--
-- Data for Name: delivery_order_item_costs; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.delivery_order_item_costs (id, delivery_order_item_id, qty, base_cost_price, total_cost_price, created_at, updated_at) FROM stdin;
8f4cfd90-f36e-4229-a651-b84d17cbf9b0	171e61e8-49aa-4e98-9f9d-be38458f6153	10.00	77972.56	311890.24	2024-06-07 16:09:40.257419	2024-06-07 16:09:40.257419
6afbf5c4-5d76-4f90-b313-6b096958f69d	171e61e8-49aa-4e98-9f9d-be38458f6153	6.00	77972.56	311890.24	2024-06-07 16:09:40.25742	2024-06-07 16:09:40.25742
234bcee0-fc70-481d-a2ac-68fdc6c08150	171e61e8-49aa-4e98-9f9d-be38458f6153	2.00	77972.56	20000.00	2024-06-07 16:09:40.257421	2024-06-07 16:09:40.257421
5a864618-c342-4902-866a-b829b32d2677	a1f5b59c-5959-4e91-86a0-6cb811438971	5.00	77972.56	155945.12	2024-06-18 06:50:45.705597	2024-06-18 06:50:45.705597
90427e3e-f57e-4b63-9701-f164db4f23ac	a1f5b59c-5959-4e91-86a0-6cb811438971	3.00	77972.56	155945.12	2024-06-18 06:50:45.705598	2024-06-18 06:50:45.705598
5487f81a-2fdb-4488-b219-b7ed8fd4e12e	a1f5b59c-5959-4e91-86a0-6cb811438971	1.00	77972.56	10000.00	2024-06-18 06:50:45.705599	2024-06-18 06:50:45.705599
c14871d6-7e5f-4333-863d-7c078f255257	4167fa9e-8979-4650-8219-b94d271bc226	2.00	77972.56	77972.56	2024-06-18 06:59:49.052868	2024-06-18 06:59:49.052868
db920714-d657-4064-bcc9-9eb38387645a	4167fa9e-8979-4650-8219-b94d271bc226	1.00	77972.56	10000.00	2024-06-18 06:59:49.052869	2024-06-18 06:59:49.052869
31216f84-1b53-4a29-a151-4a8ef8ce5d52	1610d1a2-66ad-43ac-9c76-482ea92a2fb3	1.00	77972.56	10000.00	2024-06-18 07:00:12.585617	2024-06-18 07:00:12.585617
a0239f58-35c6-418b-a87a-cb991fa0abda	eec7419e-97d2-49d7-8790-e60d2345579c	2.00	77972.56	20000.00	2024-06-18 07:00:12.585618	2024-06-18 07:00:12.585618
0b32433e-1225-44af-b970-c756bd7c1675	cce6716e-4c44-4409-b714-ecac87e0eba0	1.00	77972.56	10000.00	2024-06-18 07:18:19.122002	2024-06-18 07:18:19.122002
e3a60ce7-2710-4a05-9167-9d09afb84b72	95956370-9d8a-4914-afa5-18f71c562897	2.00	77972.56	77972.56	2024-06-18 07:18:35.322058	2024-06-18 07:18:35.322058
1695d4eb-0ce9-433b-b45c-9709fbe29d92	95956370-9d8a-4914-afa5-18f71c562897	1.00	77972.56	10000.00	2024-06-18 07:18:35.322059	2024-06-18 07:18:35.322059
14017fb3-9764-462b-bd77-602db7a9def3	c8d78f3a-fb8e-4238-8dd1-dc95f01be003	10.00	1785.67	8928.35	2024-06-20 02:22:46.919026	2024-06-20 02:22:46.919026
3bd98e53-04ec-469d-88d7-cf4ab82bcde3	c8d78f3a-fb8e-4238-8dd1-dc95f01be003	5.00	1785.67	45000.00	2024-06-20 02:22:46.919027	2024-06-20 02:22:46.919027
fab754b0-9484-4e26-bcaa-aa39615eb685	161bb992-a2d8-4824-be50-5cb1e13acff2	10.00	1785.67	8928.35	2024-06-20 09:46:58.226383	2024-06-20 09:46:58.226383
30396a3b-54c8-4545-aa7a-198f8ff361ab	161bb992-a2d8-4824-be50-5cb1e13acff2	5.00	1785.67	45000.00	2024-06-20 09:46:58.226384	2024-06-20 09:46:58.226384
292a5b71-d672-4efc-94d6-90f48ea80ec1	d59eeb5a-7093-465d-9f7f-7cd09d15243b	4.00	36730.77	36730.77	2024-06-27 09:23:18.285801	2024-06-27 09:23:18.285801
52016d3d-94de-43b1-963a-e40b47820ec5	d59eeb5a-7093-465d-9f7f-7cd09d15243b	3.00	36730.77	36730.77	2024-06-27 09:23:18.285802	2024-06-27 09:23:18.285802
e0e10e8b-0cfd-446c-8670-c0f53b3f40bd	d59eeb5a-7093-465d-9f7f-7cd09d15243b	2.00	36730.77	36730.77	2024-06-27 09:23:18.285803	2024-06-27 09:23:18.285803
21244159-1de8-4c71-a5f5-862e2efb435e	d59eeb5a-7093-465d-9f7f-7cd09d15243b	1.00	36730.77	8000.00	2024-06-27 09:23:18.285804	2024-06-27 09:23:18.285804
9b79326f-badc-48c0-afe3-67d77a3df7da	27613a85-d10a-4c7a-a486-d214a4e14e23	2.00	70740.74	160000.00	2024-06-28 00:29:08.967259	2024-06-28 00:29:08.967259
81fae1ce-c8a2-46d9-91f3-a68ccf3a18de	246fcfc9-5978-4403-a9b5-2afe11729e1e	1.00	70740.74	700000.00	2024-06-29 09:33:20.885558	2024-06-29 09:33:20.885558
21552f93-49f0-4073-be3b-9bae2f400e89	e594cc85-8b0e-44a1-90fe-d584b9034ee9	5.00	83902.43	3500000.00	2024-07-08 12:29:19.44958	2024-07-08 12:29:19.44958
\.


--
-- Data for Name: delivery_order_items; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.delivery_order_items (id, delivery_order_id, product_unit_id, qty, scale_to_base, price_per_unit, discount_per_unit, created_at, updated_at) FROM stdin;
171e61e8-49aa-4e98-9f9d-be38458f6153	07f95170-222d-4d10-a904-e5f57285fca9	cec71c3c-5a20-461b-9370-051ab3eeba76	10.00	1.00	150000.00	0.00	2024-06-07 15:48:47.898234	2024-06-07 15:48:47.898234
a1f5b59c-5959-4e91-86a0-6cb811438971	85dc9c49-c02b-4305-a5ce-de4010e03520	cec71c3c-5a20-461b-9370-051ab3eeba76	5.00	1.00	150000.00	20000.00	2024-06-18 06:50:37.048957	2024-06-18 06:50:37.048957
4167fa9e-8979-4650-8219-b94d271bc226	3545b41a-cfc4-4126-9b82-c241acaad4b4	fa437b9d-89cd-47a6-8877-2a8e91336450	2.00	12.00	150000.00	20000.00	2024-06-18 06:59:46.202779	2024-06-18 06:59:46.202779
1610d1a2-66ad-43ac-9c76-482ea92a2fb3	9b3af119-e11e-4d69-ad93-331f0eb1e8bf	fa437b9d-89cd-47a6-8877-2a8e91336450	1.00	12.00	150000.00	20000.00	2024-06-18 07:00:04.85022	2024-06-18 07:00:04.85022
eec7419e-97d2-49d7-8790-e60d2345579c	9b3af119-e11e-4d69-ad93-331f0eb1e8bf	cec71c3c-5a20-461b-9370-051ab3eeba76	2.00	1.00	150000.00	20000.00	2024-06-18 07:00:09.683415	2024-06-18 07:00:09.683415
cce6716e-4c44-4409-b714-ecac87e0eba0	516ff344-743d-4f6f-9e2d-c617ffc54402	fa437b9d-89cd-47a6-8877-2a8e91336450	1.00	12.00	150000.00	20000.00	2024-06-18 07:18:16.635412	2024-06-18 07:18:16.635412
95956370-9d8a-4914-afa5-18f71c562897	1b2731a7-f231-40f4-97e3-0454abadb848	cec71c3c-5a20-461b-9370-051ab3eeba76	2.00	1.00	150000.00	20000.00	2024-06-18 07:18:33.151486	2024-06-18 07:18:33.151486
c8d78f3a-fb8e-4238-8dd1-dc95f01be003	726acba1-51d3-4cee-aea0-377c8c43d4b6	cec71c3c-5a20-461b-9370-051ab3eeba76	10.00	1.00	150000.00	20000.00	2024-06-20 02:22:45.038885	2024-06-20 02:22:45.038885
161bb992-a2d8-4824-be50-5cb1e13acff2	ebd6eba0-ea19-418a-ba48-0704dd984bcd	cec71c3c-5a20-461b-9370-051ab3eeba76	10.00	1.00	150000.00	20000.00	2024-06-20 09:46:51.889128	2024-06-20 09:46:51.889128
d59eeb5a-7093-465d-9f7f-7cd09d15243b	4a99a612-d372-40c9-9a75-555b7fef8136	d3e101c6-9603-4370-a4db-ea52b450f5b1	4.00	20.00	100000.00	0.00	2024-06-27 09:23:07.721523	2024-06-27 09:23:07.721523
27613a85-d10a-4c7a-a486-d214a4e14e23	90bd19ca-d557-4bc1-bd8b-58ae4ccf736a	151c9e14-4403-4ed6-9973-f401a67841f2	2.00	10.00	100000.00	0.00	2024-06-28 00:28:51.833036	2024-06-28 00:28:51.833036
246fcfc9-5978-4403-a9b5-2afe11729e1e	3bcf5adc-492b-432c-ac26-df68f6c44a3b	151c9e14-4403-4ed6-9973-f401a67841f2	1.00	10.00	100000.00	5000.00	2024-06-29 09:32:42.301651	2024-06-29 09:32:42.301651
e594cc85-8b0e-44a1-90fe-d584b9034ee9	954d09b1-fc60-46f0-af64-fb7a47d461d1	151c9e14-4403-4ed6-9973-f401a67841f2	5.00	10.00	100000.00	5000.00	2024-07-08 12:29:14.213022	2024-07-08 12:29:14.213022
\.


--
-- Data for Name: delivery_order_positions; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.delivery_order_positions (id, delivery_order_id, driver_user_id, latitude, longitude, bearing, created_at, updated_at) FROM stdin;
6d62a8ce-e1cd-4905-a2d9-19ac670fafc5	3bcf5adc-492b-432c-ac26-df68f6c44a3b	fbfcbc34-77b9-4901-82f1-e5fa78d5aa48	3.5788255	98.683946	154.4514923095703	2024-06-30 15:45:32.763887	2024-07-01 11:21:47.546428
7ebbb702-f208-43c9-bf94-0910cc947345	90bd19ca-d557-4bc1-bd8b-58ae4ccf736a	fbfcbc34-77b9-4901-82f1-e5fa78d5aa48	3.58335	98.68734	250.30999755859375	2024-06-28 00:29:30.33343	2024-06-28 00:30:20.301586
91502352-cf53-4f49-abd4-3c41e4a5f373	9b3af119-e11e-4d69-ad93-331f0eb1e8bf	fbfcbc34-77b9-4901-82f1-e5fa78d5aa48	3.58318	98.68741	337.6600036621094	2024-06-18 07:09:15.352635	2024-06-18 07:10:45.437671
ed6ef27b-6c51-4822-ad79-b123f66262e8	1b2731a7-f231-40f4-97e3-0454abadb848	fbfcbc34-77b9-4901-82f1-e5fa78d5aa48	3.58335	98.68734	250.30999755859375	2024-06-18 07:20:40.747249	2024-06-18 07:20:49.962086
27f8da63-7fd1-4cfd-bdf4-8d0039681c22	726acba1-51d3-4cee-aea0-377c8c43d4b6	fbfcbc34-77b9-4901-82f1-e5fa78d5aa48	3.58335	98.68734	250.30999755859375	2024-06-20 02:34:13.091474	2024-06-20 03:59:31.728407
9ae1bc9d-ae54-41ae-ace7-47d232153fcb	07f95170-222d-4d10-a904-e5f57285fca9	fbfcbc34-77b9-4901-82f1-e5fa78d5aa48	3.5875166666666667	98.690745	0	2024-06-11 02:32:14.184323	2024-06-18 06:55:52.776417
398c145a-4a4a-421c-a92f-4f6a730357a6	4a99a612-d372-40c9-9a75-555b7fef8136	fbfcbc34-77b9-4901-82f1-e5fa78d5aa48	3.58335	98.68734	250.30999755859375	2024-06-27 09:24:59.172755	2024-06-27 09:25:49.038707
\.


--
-- Data for Name: delivery_order_return_images; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.delivery_order_return_images (id, delivery_order_return_id, file_id, created_at, updated_at) FROM stdin;
\.


--
-- Data for Name: delivery_order_returns; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.delivery_order_returns (id, delivery_order_id, user_id, description, created_at, updated_at) FROM stdin;
\.


--
-- Data for Name: delivery_order_reviews; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.delivery_order_reviews (id, delivery_order_id, star_rating, description, created_at, updated_at, type) FROM stdin;
c7ddfda3-e48f-40f7-a33f-93cd4233a364	4a99a612-d372-40c9-9a75-555b7fef8136	5	\N	2024-06-27 09:26:20.934715	2024-06-27 09:26:20.934715	
cb9aa6df-550f-447d-ab89-8059df5537bf	90bd19ca-d557-4bc1-bd8b-58ae4ccf736a	4	\N	2024-06-28 00:30:56.830006	2024-06-28 00:30:56.830006	
289b68fb-b625-4fc3-9c23-8e58f048eee4	3bcf5adc-492b-432c-ac26-df68f6c44a3b	5	gjignisdgs	2024-07-01 11:24:32.541712	2024-07-01 11:24:32.541712	
f5b1a52a-d7c2-4a6d-9ef1-f2c56c61f202	07f95170-222d-4d10-a904-e5f57285fca9	0	Barang Rusak Semua	2024-07-04 08:57:31.121076	2024-07-04 08:57:31.121076	DELIVERY
\.


--
-- Data for Name: delivery_orders; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.delivery_orders (id, customer_id, user_id, invoice_number, date, status, total_price, created_at, updated_at) FROM stdin;
ebd6eba0-ea19-418a-ba48-0704dd984bcd	f74d78f8-26ee-4699-89da-134cfbcda903	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	DO-2024620-0002	2024-06-20	CANCELED	1300000.00	2024-06-20 09:46:23.73946	2024-06-20 09:49:27.122665
b1492da9-578f-4e3f-a961-943c0ba5a0ce	77ce2ab4-a3e3-4d58-bd6c-d81a364a938c	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	DO-2024627-0001	2024-06-27	PENDING	0.00	2024-06-27 09:01:22.865764	2024-06-27 09:01:22.865764
07f95170-222d-4d10-a904-e5f57285fca9	f74d78f8-26ee-4699-89da-134cfbcda903	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	DO-202466-0001	2024-06-06	COMPLETED	1500000.00	2024-06-06 09:39:05.3483	2024-06-18 06:57:57.855161
85dc9c49-c02b-4305-a5ce-de4010e03520	f74d78f8-26ee-4699-89da-134cfbcda903	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	DO-2024612-0001	2024-06-12	COMPLETED	650000.00	2024-06-12 09:15:39.584167	2024-06-18 06:58:40.706593
4a99a612-d372-40c9-9a75-555b7fef8136	77ce2ab4-a3e3-4d58-bd6c-d81a364a938c	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	DO-2024627-0002	2024-06-27	COMPLETED	8000000.00	2024-06-27 09:22:42.197458	2024-06-27 09:25:50.427218
3545b41a-cfc4-4126-9b82-c241acaad4b4	f74d78f8-26ee-4699-89da-134cfbcda903	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	DO-2024618-0001	2024-06-18	COMPLETED	3120000.00	2024-06-18 06:59:35.500502	2024-06-18 07:08:39.552384
90bd19ca-d557-4bc1-bd8b-58ae4ccf736a	e26f6d2b-553c-4a4a-8618-e1af14087a89	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	DO-2024628-0001	2024-06-28	COMPLETED	2000000.00	2024-06-28 00:28:26.203839	2024-06-28 00:30:24.985986
9b3af119-e11e-4d69-ad93-331f0eb1e8bf	f74d78f8-26ee-4699-89da-134cfbcda903	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	DO-2024618-0002	2024-06-18	COMPLETED	1820000.00	2024-06-18 06:59:57.128592	2024-06-18 07:20:00.64983
3bcf5adc-492b-432c-ac26-df68f6c44a3b	77ce2ab4-a3e3-4d58-bd6c-d81a364a938c	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	DO-2024629-0001	2024-06-29	COMPLETED	950000.00	2024-06-29 09:32:28.490082	2024-07-01 11:21:50.228349
1b2731a7-f231-40f4-97e3-0454abadb848	f74d78f8-26ee-4699-89da-134cfbcda903	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	DO-2024618-0004	2024-06-18	COMPLETED	260000.00	2024-06-18 07:18:24.97815	2024-06-18 07:20:52.020575
516ff344-743d-4f6f-9e2d-c617ffc54402	f74d78f8-26ee-4699-89da-134cfbcda903	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	DO-2024618-0003	2024-06-18	COMPLETED	1560000.00	2024-06-18 07:18:08.596966	2024-06-20 02:21:11.023246
954d09b1-fc60-46f0-af64-fb7a47d461d1	77ce2ab4-a3e3-4d58-bd6c-d81a364a938c	33392eb4-0f87-43c6-9893-9c014fe6d561	DO-202478-0001	2024-07-08	ONGOING	4750000.00	2024-07-08 12:28:42.921562	2024-07-08 12:29:19.447434
726acba1-51d3-4cee-aea0-377c8c43d4b6	f74d78f8-26ee-4699-89da-134cfbcda903	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	DO-2024620-0001	2024-06-20	CANCELED	1300000.00	2024-06-20 02:22:35.506964	2024-06-20 03:59:35.245753
\.


--
-- Data for Name: files; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.files (id, name, type, path, created_at, updated_at) FROM stdin;
35c972f5-2537-490d-90ca-36ba659d6c20	22.png	PRODUCT_IMAGE	product/e1bf0592-7850-4602-a740-6aae98dfd281/35c972f5-2537-490d-90ca-36ba659d6c20.png	2024-06-06 14:37:03.652693	2024-06-06 14:37:03.652693
db1850d1-a9c3-49d7-a3be-46fb88eb6b09	23.png	PURCHASE_ORDER_IMAGE	purchase_order/b45d509c-7440-43ab-97d5-4296321ebe22/db1850d1-a9c3-49d7-a3be-46fb88eb6b09.png	2024-06-07 02:42:10.739926	2024-06-07 02:42:10.739926
95500471-ec6c-40d5-8fb4-0577d6b7c188	25.png	PRODUCT_RETURN_IMAGE	product_return/b977aa16-a0be-42b7-a2f6-0e6a944bfda8/95500471-ec6c-40d5-8fb4-0577d6b7c188.png	2024-06-07 07:14:06.569976	2024-06-07 07:14:06.569976
cc775b17-0e56-42fb-87b9-ddc21591e734	23.png	PRODUCT_IMAGE	product/96773beb-3646-4078-893f-19a278e98a30/cc775b17-0e56-42fb-87b9-ddc21591e734.png	2024-06-07 07:58:48.422942	2024-06-07 07:58:48.422942
851a94ea-bb56-473f-b36d-73b3e49c3225	25.png	DELIVERY_ORDER_IMAGE	delivery_order/07f95170-222d-4d10-a904-e5f57285fca9/851a94ea-bb56-473f-b36d-73b3e49c3225.png	2024-06-07 15:49:50.468002	2024-06-07 15:49:50.468002
cec57216-25dd-412c-bf70-9fe54e33c11a	25.png	CUSTOMER_PAYMENT_IMAGE	customer_payment/c494b999-334b-445a-a5f8-4989d68b714b/cec57216-25dd-412c-bf70-9fe54e33c11a.png	2024-06-07 16:10:31.549511	2024-06-07 16:10:31.549511
3bc20135-def1-44c5-8e6a-d104172bfdf3	26.png	CUSTOMER_PAYMENT_IMAGE	customer_payment/a88b6293-ec2b-4b67-9a8d-96b42a00a9e4/3bc20135-def1-44c5-8e6a-d104172bfdf3.png	2024-06-27 08:50:39.288567	2024-06-27 08:50:39.288567
38a4bd2f-2b31-4067-a355-2c924778f6d3	27.png	CUSTOMER_PAYMENT_IMAGE	customer_payment/c0da3d58-3297-4c4e-b398-93d6d1c2ed71/38a4bd2f-2b31-4067-a355-2c924778f6d3.png	2024-06-27 08:50:48.073039	2024-06-27 08:50:48.073039
c6bd93ac-0bba-4b43-94a9-e96f6184d2e7	28.png	DEBT_PAYMENT_IMAGE	debt_payment/6613d815-ff68-4a7c-8f5b-a21b045bb1ba/c6bd93ac-0bba-4b43-94a9-e96f6184d2e7.png	2024-06-27 08:51:20.380168	2024-06-27 08:51:20.380168
86676299-895f-4dc4-9399-fc30cbb487b7	29.png	PRODUCT_IMAGE	product/69a3f894-1ff2-4f61-97c2-3c957eea7914/86676299-895f-4dc4-9399-fc30cbb487b7.png	2024-06-27 08:52:04.695584	2024-06-27 08:52:04.695584
5d0b0fc8-ed9b-4f76-a956-37a5d9e23c74	31.png	PURCHASE_ORDER_IMAGE	purchase_order/29d46083-0872-4a3a-96a3-510ad7613a10/5d0b0fc8-ed9b-4f76-a956-37a5d9e23c74.png	2024-06-27 08:57:26.170827	2024-06-27 08:57:26.170827
7f274f6c-76f0-4e49-bf54-4805464396a5	32.png	PRODUCT_RECEIVE_IMAGE	product_receive/d8984d7a-307e-46d7-b8ae-bd429cc23784/7f274f6c-76f0-4e49-bf54-4805464396a5.png	2024-06-27 08:59:05.167625	2024-06-27 08:59:05.167625
7e238d0a-d192-4ad8-a13c-3dd150ea5a21	25.png	PRODUCT_RETURN_IMAGE	product_return/19ca44b3-1537-49fe-a68d-d472e909b046/7e238d0a-d192-4ad8-a13c-3dd150ea5a21.png	2024-06-27 09:00:47.583755	2024-06-27 09:00:47.583755
46da0322-efbe-41ec-bf95-b33330cd0e6b	25.png	PURCHASE_ORDER_IMAGE	purchase_order/6c227a2d-c6fd-48b4-8aa8-c4045a8981fb/46da0322-efbe-41ec-bf95-b33330cd0e6b.png	2024-06-27 09:19:34.064006	2024-06-27 09:19:34.064006
2ceca493-79a1-46d8-8e5b-7ef3b2e2e8d0	26.png	PRODUCT_RECEIVE_IMAGE	product_receive/c7bc1479-51e8-462c-9b30-0824a574b17a/2ceca493-79a1-46d8-8e5b-7ef3b2e2e8d0.png	2024-06-27 09:21:00.841617	2024-06-27 09:21:00.841617
8da9c904-5d27-4b58-b18c-744efb23e03f	27.png	PRODUCT_RETURN_IMAGE	product_return/a4ec4f6c-b3f1-4420-8cce-ed7b79d48914/8da9c904-5d27-4b58-b18c-744efb23e03f.png	2024-06-27 09:22:08.322012	2024-06-27 09:22:08.322012
e1b37960-ad21-4978-bc56-0d0d0c0a1ace	28.png	DELIVERY_ORDER_IMAGE	delivery_order/4a99a612-d372-40c9-9a75-555b7fef8136/e1b37960-ad21-4978-bc56-0d0d0c0a1ace.png	2024-06-27 09:23:13.083511	2024-06-27 09:23:13.083511
0e88c71a-e7c3-4b99-bd57-de8ef5220f3f	25.png	CUSTOMER_PAYMENT_IMAGE	customer_payment/42b428e6-ce62-4714-93ea-22b9fd930318/0e88c71a-e7c3-4b99-bd57-de8ef5220f3f.png	2024-06-28 00:18:58.871917	2024-06-28 00:18:58.871917
a770c8cf-71f8-4f4e-b4ce-7e61436e803c	26.png	CUSTOMER_PAYMENT_IMAGE	customer_payment/41de6446-f527-4154-9ca8-442499019b7a/a770c8cf-71f8-4f4e-b4ce-7e61436e803c.png	2024-06-28 00:19:10.564747	2024-06-28 00:19:10.564747
ea593a9a-3f06-46f9-8aa5-c06d431f4f12	27.png	PRODUCT_IMAGE	product/d0813034-19a1-406f-af25-c18d2f301614/ea593a9a-3f06-46f9-8aa5-c06d431f4f12.png	2024-06-28 00:20:03.394823	2024-06-28 00:20:03.394823
d689923d-2b04-4cab-afb3-87cfb46fdc71	29.png	PURCHASE_ORDER_IMAGE	purchase_order/86f2a2b3-eda0-458f-a52c-455a81d91818/d689923d-2b04-4cab-afb3-87cfb46fdc71.png	2024-06-28 00:25:25.31228	2024-06-28 00:25:25.31228
c29d71c5-75b8-4300-a6d8-fa1cf938d42f	25.png	DELIVERY_ORDER_IMAGE	delivery_order/3bcf5adc-492b-432c-ac26-df68f6c44a3b/c29d71c5-75b8-4300-a6d8-fa1cf938d42f.png	2024-06-29 09:33:12.977553	2024-06-29 09:33:12.977553
66e0519f-c0ec-4ca2-85da-b1b1c2283012	25.png	PRODUCT_IMAGE	product/bcfd2b82-8125-4e06-b6d8-2e42c70fdd74/66e0519f-c0ec-4ca2-85da-b1b1c2283012.png	2024-06-29 15:38:14.33951	2024-06-29 15:38:14.33951
26249674-49b6-42c9-8412-4ff1b30e4552	25.png	CUSTOMER_PAYMENT_IMAGE	customer_payment/f9c98431-1477-40f4-91b4-c31b03aadefe/26249674-49b6-42c9-8412-4ff1b30e4552.png	2024-07-01 11:40:17.41247	2024-07-01 11:40:17.41247
e199a34b-54a7-4c0c-a1be-8a5d2eb33eb2	26.png	CUSTOMER_PAYMENT_IMAGE	customer_payment/ccd94971-090e-4c11-b902-2ef107d9f528/e199a34b-54a7-4c0c-a1be-8a5d2eb33eb2.png	2024-07-01 11:42:03.943652	2024-07-01 11:42:03.943652
fa26da76-a459-47fc-a405-34477fa491f0	27.png	CUSTOMER_PAYMENT_IMAGE	customer_payment/d0acbcf7-9b87-4478-86c8-e4775aaef26f/fa26da76-a459-47fc-a405-34477fa491f0.png	2024-07-01 11:42:21.034439	2024-07-01 11:42:21.034439
9373fcd0-92d9-48e9-b0c9-d0737dc3341a	25.png	CUSTOMER_PAYMENT_IMAGE	customer_payment/f53ca92b-a0d8-41af-9006-1b82deee4368/9373fcd0-92d9-48e9-b0c9-d0737dc3341a.png	2024-07-02 07:38:42.038559	2024-07-02 07:38:42.038559
f735c1b5-cae8-4a02-83bb-24a3cc62823c	25.png	PRODUCT_IMAGE	product/2ea93d98-d8d5-4ce2-b13f-00ac30f863fc/f735c1b5-cae8-4a02-83bb-24a3cc62823c.png	2024-07-03 03:06:00.058417	2024-07-03 03:06:00.058417
3b760ad8-1f2e-4cf9-ad25-0eddf9966698	26.png	PRODUCT_IMAGE	product/32ca75e1-7d0f-4fc0-ad09-78d16d820499/3b760ad8-1f2e-4cf9-ad25-0eddf9966698.png	2024-07-03 03:06:43.253781	2024-07-03 03:06:43.253781
b15546e2-d5f3-4603-8014-1e37c2902742	28.png	PRODUCT_IMAGE	product/a525adb6-2358-4bae-920c-1f1249fbd3ff/b15546e2-d5f3-4603-8014-1e37c2902742.png	2024-07-03 03:07:01.973465	2024-07-03 03:07:01.973465
3eb64ea4-cacc-4700-a3c9-934b4c2ea402	29.png	PRODUCT_IMAGE	product/0fefe4a3-0daf-4056-82c3-76bd49bb0bef/3eb64ea4-cacc-4700-a3c9-934b4c2ea402.png	2024-07-03 03:07:29.953732	2024-07-03 03:07:29.953732
13781a6c-1655-4c37-a905-287132abd378	30.png	PRODUCT_IMAGE	product/fe95e0cb-6070-4f93-8c69-4a2b61d3992e/13781a6c-1655-4c37-a905-287132abd378.png	2024-07-03 03:08:08.250362	2024-07-03 03:08:08.250362
c90220c3-440d-4224-9da2-e5112b52f0d9	31.png	PRODUCT_IMAGE	product/2f2173c8-05e1-4328-b6f8-c769927d43bd/c90220c3-440d-4224-9da2-e5112b52f0d9.png	2024-07-03 03:08:31.836826	2024-07-03 03:08:31.836826
20621060-f268-4863-8246-1ec1187789e5	25.png	PRODUCT_RETURN_IMAGE	product_return/c1d72997-80c0-464d-81e4-d6c58b958064/20621060-f268-4863-8246-1ec1187789e5.png	2024-07-08 15:03:37.215956	2024-07-08 15:03:37.215956
552b2ddd-1403-4911-85da-ad97043742e1	26.png	PURCHASE_ORDER_IMAGE	purchase_order/66989bdd-f2e8-404b-87ed-d283180a325e/552b2ddd-1403-4911-85da-ad97043742e1.png	2024-07-08 15:17:08.585967	2024-07-08 15:17:08.585967
b39a3ddc-6b2f-46d5-8ecf-c01abb5e29de	27.png	PRODUCT_IMAGE	product/692e6c93-f159-40b0-bfb8-180ff6c7dfad/b39a3ddc-6b2f-46d5-8ecf-c01abb5e29de.png	2024-07-08 15:56:20.624088	2024-07-08 15:56:20.624088
ab461851-6d15-4442-85aa-cc13758df092	25.png	PRODUCT_IMAGE	product/0b451368-e840-439d-9f07-824beb763eda/ab461851-6d15-4442-85aa-cc13758df092.png	2024-07-08 16:06:01.233632	2024-07-08 16:06:01.233632
6970f822-6ba5-4cfd-8cdd-ca5e9b5502fa	25.png	PRODUCT_IMAGE	product/412d3012-8a9f-44ea-a8a0-273aaa6a9cfc/6970f822-6ba5-4cfd-8cdd-ca5e9b5502fa.png	2024-07-09 00:43:16.894131	2024-07-09 00:43:16.894131
f2532f25-64eb-4cc0-badf-c64992882ffa	25.png	PRODUCT_IMAGE	product/50b0499b-54e2-4406-8da7-ec66af69e932/f2532f25-64eb-4cc0-badf-c64992882ffa.png	2024-07-10 02:00:18.142031	2024-07-10 02:00:18.142031
7957cf37-7c5a-4b32-a1e7-2d0d4e056c47	25.png	CUSTOMER_PAYMENT_IMAGE	customer_payment/978ad65b-419f-4166-ad4c-d2cacbfeb0ee/7957cf37-7c5a-4b32-a1e7-2d0d4e056c47.png	2024-07-10 15:32:08.80292	2024-07-10 15:32:08.80292
10f6ef3c-85ab-46e7-8be0-a0b6e0a6379e	26.png	DEBT_PAYMENT_IMAGE	debt_payment/6407bddd-1794-434e-8d6a-7f83e8644500/10f6ef3c-85ab-46e7-8be0-a0b6e0a6379e.png	2024-07-10 15:33:13.51	2024-07-10 15:33:13.51
\.


--
-- Data for Name: permissions; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.permissions (id, title, description, is_active, created_at, updated_at) FROM stdin;
8827e93d-f16e-4ad3-90f2-501729aefb48	BALANCE_CREATE	BALANCE CREATE	t	2024-07-10 02:53:07.860287	2024-07-10 02:53:07.860287
d2756d0e-a71a-47df-8a74-8caf53213987	BALANCE_FETCH	BALANCE FETCH	t	2024-07-10 02:53:07.860288	2024-07-10 02:53:07.860288
a6dcc882-a0a2-46ef-b0f5-b56d0d08f859	BALANCE_GET	BALANCE GET	t	2024-07-10 02:53:07.860289	2024-07-10 02:53:07.860289
2d02ce31-4c53-4827-933d-479423a997f2	BALANCE_UPDATE	BALANCE UPDATE	t	2024-07-10 02:53:07.86029	2024-07-10 02:53:07.86029
47c542b7-e00c-469b-bbd8-a9c4c4d93c50	BALANCE_DELETE	BALANCE DELETE	t	2024-07-10 02:53:07.860291	2024-07-10 02:53:07.860291
504b4d4b-f5cd-43ca-8d5c-afc0b9c7e072	CART_GET_ACTIVE	CART GET ACTIVE	t	2024-07-10 02:53:07.860292	2024-07-10 02:53:07.860292
ab064625-5901-4ca4-86c6-7d310991c468	CART_ADD_ITEM	CART ADD ITEM	t	2024-07-10 02:53:07.860293	2024-07-10 02:53:07.860293
600fa561-7aa6-4bc9-b0a9-95bcaeb9b598	CART_UPDATE_ITEM	CART UPDATE ITEM	t	2024-07-10 02:53:07.860294	2024-07-10 02:53:07.860294
5b0121e7-c4ea-40f6-ac45-894bee6f52c7	CART_DELETE_ITEM	CART DELETE ITEM	t	2024-07-10 02:53:07.860295	2024-07-10 02:53:07.860295
9b413c7c-847b-4343-b5d1-21263385f949	CART_SET_ACTIVE	CART SET ACTIVE	t	2024-07-10 02:53:07.860296	2024-07-10 02:53:07.860296
6c2d12af-7d30-41ca-815a-df17a11054ce	CART_SET_IN_ACTIVE	CART SET IN ACTIVE	t	2024-07-10 02:53:07.860297	2024-07-10 02:53:07.860297
1665c450-4d60-4607-ba2a-e593a0892b17	CART_DELETE	CART DELETE	t	2024-07-10 02:53:07.860298	2024-07-10 02:53:07.860298
ce029379-9607-460d-858e-ddf66f914821	CASHIER_SESSION_FETCH	CASHIER SESSION FETCH	t	2024-07-10 02:53:07.860299	2024-07-10 02:53:07.860299
d5269684-097e-48fa-b6b0-abd1d3322a3d	CASHIER_SESSION_FETCH_FOR_CURRENT_USER	CASHIER SESSION FETCH FOR CURRENT USER	t	2024-07-10 02:53:07.8603	2024-07-10 02:53:07.8603
0e81bdb3-c418-45b4-bdf4-75e9fbad332b	CASHIER_SESSION_START	CASHIER SESSION START	t	2024-07-10 02:53:07.860301	2024-07-10 02:53:07.860301
1e327b7f-00b0-4f8a-9295-e10c014826de	CASHIER_SESSION_GET	CASHIER SESSION GET	t	2024-07-10 02:53:07.860302	2024-07-10 02:53:07.860302
dd0ac738-9352-4128-90fb-33ea413d1ef9	CASHIER_SESSION_FETCH_TRANSACTION	CASHIER SESSION FETCH TRANSACTION	t	2024-07-10 02:53:07.860303	2024-07-10 02:53:07.860303
8b606434-7590-4114-adcc-3ed18669a19d	CASHIER_SESSION_DOWNLOAD_REPORT	CASHIER SESSION DOWNLOAD REPORT	t	2024-07-10 02:53:07.860304	2024-07-10 02:53:07.860304
c5231c99-cbe3-41bf-bcc6-3f90b463c07d	CASHIER_SESSION_GET_CURRENT	CASHIER SESSION GET CURRENT	t	2024-07-10 02:53:07.860305	2024-07-10 02:53:07.860305
4165e784-1d77-4c59-83ae-bb8097ddaf48	CASHIER_SESSION_END	CASHIER SESSION END	t	2024-07-10 02:53:07.860306	2024-07-10 02:53:07.860306
9aaa4ca6-33b4-41d9-81e3-9f6e71590d0c	CUSTOMER_CREATE	CUSTOMER CREATE	t	2024-07-10 02:53:07.860307	2024-07-10 02:53:07.860307
c6dbeb2d-52e9-411b-8e84-5d594b32b20d	CUSTOMER_FETCH	CUSTOMER FETCH	t	2024-07-10 02:53:07.860308	2024-07-10 02:53:07.860308
90ad8535-3ea7-4f6e-8e58-8f316377cee0	CUSTOMER_GET	CUSTOMER GET	t	2024-07-10 02:53:07.860309	2024-07-10 02:53:07.860309
82e4aae1-1d14-4511-a21b-37cc6ce28391	CUSTOMER_UPDATE	CUSTOMER UPDATE	t	2024-07-10 02:53:07.86031	2024-07-10 02:53:07.86031
f5405977-8f91-4954-bafc-e4c4c5f628e4	CUSTOMER_DELETE	CUSTOMER DELETE	t	2024-07-10 02:53:07.860311	2024-07-10 02:53:07.860311
d67e64c2-b814-48f8-beb7-74d60afe579f	CUSTOMER_OPTION_FOR_CUSTOMER_DEBT_REPORT_FORM	CUSTOMER OPTION FOR CUSTOMER DEBT REPORT FORM	t	2024-07-10 02:53:07.860312	2024-07-10 02:53:07.860312
877cf804-d173-4f9a-8818-3eeba0fe0183	CUSTOMER_OPTION_FOR_WHATSAPP_CUSTOMER_DEBT_BROADCAST_FORM	CUSTOMER OPTION FOR WHATSAPP CUSTOMER DEBT BROADCAST FORM	t	2024-07-10 02:53:07.860313	2024-07-10 02:53:07.860313
91b0dd33-0bbe-4641-844d-ec6294163f4d	CUSTOMER_OPTION_FOR_DELIVERY_ORDER_FORM	CUSTOMER OPTION FOR DELIVERY ORDER FORM	t	2024-07-10 02:53:07.860314	2024-07-10 02:53:07.860314
ec412394-c888-4deb-a87d-7138ae9b12e0	CUSTOMER_OPTION_FOR_DELIVERY_ORDER_FILTER	CUSTOMER OPTION FOR DELIVERY ORDER FILTER	t	2024-07-10 02:53:07.860315	2024-07-10 02:53:07.860315
c7f918e0-8b20-4d2f-80a6-4f861e12faa4	CUSTOMER_DEBT_UPLOAD_IMAGE	CUSTOMER DEBT UPLOAD IMAGE	t	2024-07-10 02:53:07.860316	2024-07-10 02:53:07.860316
8498428c-ab74-4f89-858b-045240445f2e	CUSTOMER_DEBT_DOWNLOAD_REPORT	CUSTOMER DEBT DOWNLOAD REPORT	t	2024-07-10 02:53:07.860317	2024-07-10 02:53:07.860317
33cdc1aa-b351-4fe8-bd22-87752c178712	CUSTOMER_DEBT_FETCH	CUSTOMER DEBT FETCH	t	2024-07-10 02:53:07.860318	2024-07-10 02:53:07.860318
0bbd03cb-4610-4a15-8edd-7cee9e6f8a36	CUSTOMER_DEBT_GET	CUSTOMER DEBT GET	t	2024-07-10 02:53:07.860319	2024-07-10 02:53:07.860319
641a1c56-5d11-4b45-a678-b6c63e89c5ed	CUSTOMER_DEBT_PAYMENT	CUSTOMER DEBT PAYMENT	t	2024-07-10 02:53:07.86032	2024-07-10 02:53:07.86032
b3da3e77-b9d1-49c9-99f3-9e81e8ed32d2	CUSTOMER_TYPE_CREATE	CUSTOMER TYPE CREATE	t	2024-07-10 02:53:07.860321	2024-07-10 02:53:07.860321
d47a4dde-9aea-42d7-8723-9d263a9b2c8e	CUSTOMER_TYPE_FETCH	CUSTOMER TYPE FETCH	t	2024-07-10 02:53:07.860322	2024-07-10 02:53:07.860322
20e9761b-a5a4-4086-b25b-58987ad95cc7	CUSTOMER_TYPE_GET	CUSTOMER TYPE GET	t	2024-07-10 02:53:07.860323	2024-07-10 02:53:07.860323
9e7bdca1-35da-4d3f-9799-00f4e79323a9	CUSTOMER_TYPE_UPDATE	CUSTOMER TYPE UPDATE	t	2024-07-10 02:53:07.860324	2024-07-10 02:53:07.860324
1a426476-0f40-4582-9800-a167de643a4f	CUSTOMER_TYPE_DELETE	CUSTOMER TYPE DELETE	t	2024-07-10 02:53:07.860325	2024-07-10 02:53:07.860325
f2a9d1fb-0a3e-4e8b-ad32-d47d840505a4	CUSTOMER_TYPE_ADD_DISCOUNT	CUSTOMER TYPE ADD DISCOUNT	t	2024-07-10 02:53:07.860326	2024-07-10 02:53:07.860326
8a06286f-0df6-4d21-9fe3-975758125239	CUSTOMER_TYPE_UPDATE_DISCOUNT	CUSTOMER TYPE UPDATE DISCOUNT	t	2024-07-10 02:53:07.860327	2024-07-10 02:53:07.860327
321b797f-3b85-4f4c-8ba2-f91aa24c4161	CUSTOMER_TYPE_DELETE_DISCOUNT	CUSTOMER TYPE DELETE DISCOUNT	t	2024-07-10 02:53:07.860328	2024-07-10 02:53:07.860328
2ced73f0-461e-4bf4-ac89-ad8436c4d59c	CUSTOMER_TYPE_OPTION_FOR_CUSTOMER_FORM	CUSTOMER TYPE OPTION FOR CUSTOMER FORM	t	2024-07-10 02:53:07.860329	2024-07-10 02:53:07.860329
f2c1794b-e513-47f1-a009-0b3329080565	CUSTOMER_TYPE_OPTION_FOR_WHATSAPP_PRODUCT_PRICE_CHANGE_BROADCAST_FORM	CUSTOMER TYPE OPTION FOR WHATSAPP PRODUCT PRICE CHANGE BROADCAST FORM	t	2024-07-10 02:53:07.86033	2024-07-10 02:53:07.86033
6a5d87b1-3a89-48c4-b9b7-e9dc6bed11b8	DASHBOARD_SUMMARIZE_DEBT	DASHBOARD SUMMARIZE DEBT	t	2024-07-10 02:53:07.860331	2024-07-10 02:53:07.860331
c6a9c058-19fb-4eff-b050-517750e87888	DASHBOARD_SUMMARIZE_TRANSACTION	DASHBOARD SUMMARIZE TRANSACTION	t	2024-07-10 02:53:07.860332	2024-07-10 02:53:07.860332
a0c74ca7-8945-4226-bac4-43ffe2f841c6	DEBT_UPLOAD_IMAGE	DEBT UPLOAD IMAGE	t	2024-07-10 02:53:07.860333	2024-07-10 02:53:07.860333
163c477a-e44d-4c4f-96c7-7d24a686e85e	DEBT_FETCH	DEBT FETCH	t	2024-07-10 02:53:07.860334	2024-07-10 02:53:07.860334
8d31041b-440a-4e57-8fbe-5a6e75c0f87f	DEBT_DOWNLOAD_REPORT	DEBT DOWNLOAD REPORT	t	2024-07-10 02:53:07.860335	2024-07-10 02:53:07.860335
66d678b6-b95e-4c19-8d16-226d9dc86e9f	DEBT_GET	DEBT GET	t	2024-07-10 02:53:07.860336	2024-07-10 02:53:07.860336
2c0be1d0-8f75-4df1-95eb-b084a268ba65	DEBT_PAYMENT	DEBT PAYMENT	t	2024-07-10 02:53:07.860337	2024-07-10 02:53:07.860337
9e32f57b-fccb-4041-bb56-6a7a44b8f625	DELIVERY_ORDER_CREATE	DELIVERY ORDER CREATE	t	2024-07-10 02:53:07.860338	2024-07-10 02:53:07.860338
3a2cec95-710e-424a-84ac-5f36a8a6fe1c	DELIVERY_ORDER_DOWNLOAD_REPORT	DELIVERY ORDER DOWNLOAD REPORT	t	2024-07-10 02:53:07.860339	2024-07-10 02:53:07.860339
8a62b161-d14c-4cca-b26a-1cb8c6e03d7a	DELIVERY_ORDER_UPLOAD	DELIVERY ORDER UPLOAD	t	2024-07-10 02:53:07.86034	2024-07-10 02:53:07.86034
539a1f07-7bb0-4a0d-b207-647f7be51b8a	DELIVERY_ORDER_ADD_ITEM	DELIVERY ORDER ADD ITEM	t	2024-07-10 02:53:07.860341	2024-07-10 02:53:07.860341
7bcc4d1a-e10f-48a9-bbf4-687f504520a9	DELIVERY_ORDER_ADD_IMAGE	DELIVERY ORDER ADD IMAGE	t	2024-07-10 02:53:07.860342	2024-07-10 02:53:07.860342
99bd3a4d-b5af-4d80-900d-f2a3a590a730	DELIVERY_ORDER_ADD_DRIVER	DELIVERY ORDER ADD DRIVER	t	2024-07-10 02:53:07.860343	2024-07-10 02:53:07.860343
ab3217e4-eef1-4e88-a651-2aebba7421bd	DELIVERY_ORDER_FETCH	DELIVERY ORDER FETCH	t	2024-07-10 02:53:07.860344	2024-07-10 02:53:07.860344
c76e3196-b313-4d15-845c-fb9a4bec9840	DELIVERY_ORDER_FETCH_DRIVER	DELIVERY ORDER FETCH DRIVER	t	2024-07-10 02:53:07.860345	2024-07-10 02:53:07.860345
ffe81eb0-560c-44cb-b53c-5e7733ac652d	DELIVERY_ORDER_GET	DELIVERY ORDER GET	t	2024-07-10 02:53:07.860346	2024-07-10 02:53:07.860346
f9b46964-fa73-43e8-8f5a-1041ed8ddefb	DELIVERY_ORDER_ACTIVE_FOR_DRIVER	DELIVERY ORDER ACTIVE FOR DRIVER	t	2024-07-10 02:53:07.860347	2024-07-10 02:53:07.860347
b1ed1d28-a4c7-4539-986d-628d501b5959	DELIVERY_ORDER_MARK_ONGOING	DELIVERY ORDER MARK ONGOING	t	2024-07-10 02:53:07.860348	2024-07-10 02:53:07.860348
a49bc736-f731-42bd-ba6c-7064f7d75d5d	DELIVERY_ORDER_DELIVERING	DELIVERY ORDER DELIVERING	t	2024-07-10 02:53:07.860349	2024-07-10 02:53:07.860349
8fa77f3c-9d9f-4f2e-a7f0-8c32571448ec	DELIVERY_ORDER_UPDATE	DELIVERY ORDER UPDATE	t	2024-07-10 02:53:07.86035	2024-07-10 02:53:07.86035
80295769-1e98-4bd6-9c91-29693ad689f1	DELIVERY_ORDER_CANCEL	DELIVERY ORDER CANCEL	t	2024-07-10 02:53:07.860351	2024-07-10 02:53:07.860351
cefb77ac-1f58-4f37-82b0-7fba3540782a	DELIVERY_ORDER_MARK_COMPLETED	DELIVERY ORDER MARK COMPLETED	t	2024-07-10 02:53:07.860352	2024-07-10 02:53:07.860352
013195bc-285e-4a02-bcfc-6e1750442584	DELIVERY_ORDER_RETURNED	DELIVERY ORDER RETURNED	t	2024-07-10 02:53:07.860353	2024-07-10 02:53:07.860353
2d59a87a-c80f-437e-9d3d-fac82859cde3	DELIVERY_ORDER_DELIVERY_LOCATION	DELIVERY ORDER DELIVERY LOCATION	t	2024-07-10 02:53:07.860354	2024-07-10 02:53:07.860354
f8d4581e-ce94-4d49-a2d5-a73348cf1da6	DELIVERY_ORDER_DELETE	DELIVERY ORDER DELETE	t	2024-07-10 02:53:07.860355	2024-07-10 02:53:07.860355
7abd445c-e14e-4153-8580-0f7bfaf615ea	DELIVERY_ORDER_DELETE_ITEM	DELIVERY ORDER DELETE ITEM	t	2024-07-10 02:53:07.860356	2024-07-10 02:53:07.860356
ec53f87b-f420-428d-8f04-e55234ae146d	DELIVERY_ORDER_DELETE_IMAGE	DELIVERY ORDER DELETE IMAGE	t	2024-07-10 02:53:07.860357	2024-07-10 02:53:07.860357
e0515501-9069-4ff0-bdb6-ca0cbe734c76	DELIVERY_ORDER_DELETE_DRIVER	DELIVERY ORDER DELETE DRIVER	t	2024-07-10 02:53:07.860358	2024-07-10 02:53:07.860358
26529ea7-d633-46e7-a080-fefe8924cf0a	DELIVERY_ORDER_REVIEW_FETCH	DELIVERY ORDER REVIEW FETCH	t	2024-07-10 02:53:07.860359	2024-07-10 02:53:07.860359
4e4b3bae-4446-4785-86bb-5148772f2589	DELIVERY_ORDER_REVIEW_GET	DELIVERY ORDER REVIEW GET	t	2024-07-10 02:53:07.86036	2024-07-10 02:53:07.86036
a9195ab7-6e2f-4716-b56d-c5888599ca92	PRODUCT_CREATE	PRODUCT CREATE	t	2024-07-10 02:53:07.860361	2024-07-10 02:53:07.860361
f888ed1c-8daf-46ff-b80a-7b95d5503145	PRODUCT_UPLOAD	PRODUCT UPLOAD	t	2024-07-10 02:53:07.860362	2024-07-10 02:53:07.860362
171bf39a-c8f0-4bd4-8bab-9bb113660ba3	PRODUCT_FETCH	PRODUCT FETCH	t	2024-07-10 02:53:07.860363	2024-07-10 02:53:07.860363
6e3ed14e-426f-46f8-a9c4-b8be122d2c89	PRODUCT_GET	PRODUCT GET	t	2024-07-10 02:53:07.860364	2024-07-10 02:53:07.860364
601d894f-8ec0-42bf-9a30-e2df4ad8aa23	PRODUCT_UPDATE	PRODUCT UPDATE	t	2024-07-10 02:53:07.860365	2024-07-10 02:53:07.860365
fcc24c96-608c-47f7-9b55-128231841c3f	PRODUCT_DELETE	PRODUCT DELETE	t	2024-07-10 02:53:07.860366	2024-07-10 02:53:07.860366
45ce843c-9dee-49f8-b358-581a61e4e641	PRODUCT_OPTION_FOR_PRODUCT_RECEIVE_ITEM_FORM	PRODUCT OPTION FOR PRODUCT RECEIVE ITEM FORM	t	2024-07-10 02:53:07.860367	2024-07-10 02:53:07.860367
71d7b694-8f90-4d01-adf8-bc952e1679bb	PRODUCT_OPTION_FOR_DELIVERY_ORDER_ITEM_FORM	PRODUCT OPTION FOR DELIVERY ORDER ITEM FORM	t	2024-07-10 02:53:07.860368	2024-07-10 02:53:07.860368
efccb3c8-2351-441c-abf8-ad819243fa3c	PRODUCT_OPTION_FOR_CUSTOMER_TYPE_DISCOUNT_FORM	PRODUCT OPTION FOR CUSTOMER TYPE DISCOUNT FORM	t	2024-07-10 02:53:07.860369	2024-07-10 02:53:07.860369
691f7440-1d77-4c66-9fc6-3abfd0593cb8	PRODUCT_OPTION_FOR_CART_ADD_ITEM_FORM	PRODUCT OPTION FOR CART ADD ITEM FORM	t	2024-07-10 02:53:07.86037	2024-07-10 02:53:07.86037
4b1b4078-b751-427d-8f86-11df92d7a1b3	PRODUCT_OPTION_FOR_PRODUCT_DISCOUNT_FORM	PRODUCT OPTION FOR PRODUCT DISCOUNT FORM	t	2024-07-10 02:53:07.860371	2024-07-10 02:53:07.860371
da4f9a7b-4153-4735-a4a9-da943ea52cbb	PRODUCT_DISCOUNT_CREATE	PRODUCT DISCOUNT CREATE	t	2024-07-10 02:53:07.860372	2024-07-10 02:53:07.860372
c8d83756-e082-4332-82da-5d7efad64abc	PRODUCT_DISCOUNT_FETCH	PRODUCT DISCOUNT FETCH	t	2024-07-10 02:53:07.860373	2024-07-10 02:53:07.860373
c0451ce1-5573-4964-b96a-869a813f4bc2	PRODUCT_DISCOUNT_GET	PRODUCT DISCOUNT GET	t	2024-07-10 02:53:07.860374	2024-07-10 02:53:07.860374
1e2eb6ca-e32b-420a-bc07-53a3d6c6227a	PRODUCT_DISCOUNT_UPDATE	PRODUCT DISCOUNT UPDATE	t	2024-07-10 02:53:07.860375	2024-07-10 02:53:07.860375
9d933c71-4073-44e9-a3ce-aeca83eca948	PRODUCT_DISCOUNT_DELETE	PRODUCT DISCOUNT DELETE	t	2024-07-10 02:53:07.860376	2024-07-10 02:53:07.860376
fca47e9f-9a7b-4eef-a312-3c525cc7871c	PRODUCT_RECEIVE_UPLOAD	PRODUCT RECEIVE UPLOAD	t	2024-07-10 02:53:07.860377	2024-07-10 02:53:07.860377
7879169e-1c71-40e5-b077-a647f7581d80	PRODUCT_RECEIVE_ADD_IMAGE	PRODUCT RECEIVE ADD IMAGE	t	2024-07-10 02:53:07.860378	2024-07-10 02:53:07.860378
1d84ac87-15a6-4454-936a-30692a5b7941	PRODUCT_RECEIVE_UPDATE	PRODUCT RECEIVE UPDATE	t	2024-07-10 02:53:07.860379	2024-07-10 02:53:07.860379
df196d96-a49d-48df-b1aa-96e9e8d70df0	PRODUCT_RECEIVE_CANCEL	PRODUCT RECEIVE CANCEL	t	2024-07-10 02:53:07.86038	2024-07-10 02:53:07.86038
a4554516-a1d9-4d6c-b68e-dad15163a569	PRODUCT_RECEIVE_MARK_COMPLETE	PRODUCT RECEIVE MARK COMPLETE	t	2024-07-10 02:53:07.860381	2024-07-10 02:53:07.860381
7766b2e2-b835-43ba-acbc-434c24fd6e5c	PRODUCT_RECEIVE_FETCH	PRODUCT RECEIVE FETCH	t	2024-07-10 02:53:07.860382	2024-07-10 02:53:07.860382
440fff70-d0be-44c2-9a0f-3fe1b561631a	PRODUCT_RECEIVE_GET	PRODUCT RECEIVE GET	t	2024-07-10 02:53:07.860383	2024-07-10 02:53:07.860383
c81e91b1-45e2-446a-9395-59dbf00cfb7f	PRODUCT_RECEIVE_UPDATE_ITEM	PRODUCT RECEIVE UPDATE ITEM	t	2024-07-10 02:53:07.860384	2024-07-10 02:53:07.860384
3b74461a-1789-4601-9416-881c98cf9587	PRODUCT_RECEIVE_DELETE	PRODUCT RECEIVE DELETE	t	2024-07-10 02:53:07.860385	2024-07-10 02:53:07.860385
1230d117-1892-4c15-9179-a517d9d5351d	PRODUCT_RECEIVE_DELETE_IMAGE	PRODUCT RECEIVE DELETE IMAGE	t	2024-07-10 02:53:07.860386	2024-07-10 02:53:07.860386
7e696943-0f22-46f1-90d7-b016b6e5f81d	PRODUCT_RETURN_CREATE	PRODUCT RETURN CREATE	t	2024-07-10 02:53:07.860387	2024-07-10 02:53:07.860387
1a912ef6-8a84-4e17-9c86-e5355f2b9c91	PRODUCT_RETURN_UPLOAD	PRODUCT RETURN UPLOAD	t	2024-07-10 02:53:07.860388	2024-07-10 02:53:07.860388
038712a3-a8d6-4626-a4e5-e42507713e3e	PRODUCT_RETURN_ADD_ITEM	PRODUCT RETURN ADD ITEM	t	2024-07-10 02:53:07.860389	2024-07-10 02:53:07.860389
358ce19f-812a-445d-a686-fe10b339d94f	PRODUCT_RETURN_ADD_IMAGE	PRODUCT RETURN ADD IMAGE	t	2024-07-10 02:53:07.86039	2024-07-10 02:53:07.86039
8a88808b-3c50-44bf-885c-7697ca73b749	PRODUCT_RETURN_UPDATE	PRODUCT RETURN UPDATE	t	2024-07-10 02:53:07.860391	2024-07-10 02:53:07.860391
57ecf435-b67e-4e3a-b597-a5b9b1bca788	PRODUCT_RETURN_CANCEL	PRODUCT RETURN CANCEL	t	2024-07-10 02:53:07.860392	2024-07-10 02:53:07.860392
7a3da445-7f55-4cde-9f83-c47bed4719b9	PRODUCT_RETURN_MARK_COMPLETE	PRODUCT RETURN MARK COMPLETE	t	2024-07-10 02:53:07.860393	2024-07-10 02:53:07.860393
52be45c3-f444-4a7f-bfa4-43b4bd94ad18	PRODUCT_RETURN_FETCH	PRODUCT RETURN FETCH	t	2024-07-10 02:53:07.860394	2024-07-10 02:53:07.860394
3588e557-84ac-4ca5-87ba-96086f41528f	PRODUCT_RETURN_GET	PRODUCT RETURN GET	t	2024-07-10 02:53:07.860395	2024-07-10 02:53:07.860395
5ecfa2fa-da00-4f0f-a4a5-4a75ef16deba	PRODUCT_RETURN_DELETE	PRODUCT RETURN DELETE	t	2024-07-10 02:53:07.860396	2024-07-10 02:53:07.860396
fe5b6be5-599f-414b-8a35-117e94cbf06f	PRODUCT_RETURN_DELETE_ITEM	PRODUCT RETURN DELETE ITEM	t	2024-07-10 02:53:07.860397	2024-07-10 02:53:07.860397
7379cd4a-2c0b-4f7f-bbf2-f2f740746b11	PRODUCT_RETURN_DELETE_IMAGE	PRODUCT RETURN DELETE IMAGE	t	2024-07-10 02:53:07.860398	2024-07-10 02:53:07.860398
fb23c89d-55be-424b-a545-7b5978b73da4	PRODUCT_STOCK_ADJUSTMENT_FETCH	PRODUCT STOCK ADJUSTMENT FETCH	t	2024-07-10 02:53:07.860399	2024-07-10 02:53:07.860399
f2067064-7f15-409c-8348-dd9a7d7f9ea5	PRODUCT_STOCK_FETCH	PRODUCT STOCK FETCH	t	2024-07-10 02:53:07.8604	2024-07-10 02:53:07.8604
319b03c9-593c-4d93-818c-4842227b5390	PRODUCT_STOCK_GET	PRODUCT STOCK GET	t	2024-07-10 02:53:07.860401	2024-07-10 02:53:07.860401
db0e803f-3dd0-421a-ba39-75bdffa8467a	PRODUCT_STOCK_DOWNLOAD_REPORT	PRODUCT STOCK DOWNLOAD REPORT	t	2024-07-10 02:53:07.860402	2024-07-10 02:53:07.860402
ebc849ac-f205-4ee0-be30-aaffe0fd1197	PRODUCT_STOCK_ADJUSTMENT	PRODUCT STOCK ADJUSTMENT	t	2024-07-10 02:53:07.860403	2024-07-10 02:53:07.860403
eff24eed-1013-469d-92c0-2be65b9283bb	PRODUCT_UNIT_CREATE	PRODUCT UNIT CREATE	t	2024-07-10 02:53:07.860404	2024-07-10 02:53:07.860404
0bd6b4f2-1831-415f-a9f5-a429dd13fa47	PRODUCT_UNIT_GET	PRODUCT UNIT GET	t	2024-07-10 02:53:07.860405	2024-07-10 02:53:07.860405
cc39ef50-77d5-4b16-b6bb-b6d946b5cbc7	PRODUCT_UNIT_UPDATE	PRODUCT UNIT UPDATE	t	2024-07-10 02:53:07.860406	2024-07-10 02:53:07.860406
6c46505a-258f-41be-9ee5-1d6d822ee411	PRODUCT_UNIT_DELETE	PRODUCT UNIT DELETE	t	2024-07-10 02:53:07.860407	2024-07-10 02:53:07.860407
dac0c5ef-acd5-4e8a-99a4-eec374e359a8	PRODUCT_UNIT_OPTION_FOR_PRODUCT_RECEIVE_ITEM_FORM	PRODUCT UNIT OPTION FOR PRODUCT RECEIVE ITEM FORM	t	2024-07-10 02:53:07.860408	2024-07-10 02:53:07.860408
63768160-5ecd-4965-9680-b6ccf56c7e3b	PRODUCT_UNIT_OPTION_FOR_DELIVERY_ORDER_ITEM_FORM	PRODUCT UNIT OPTION FOR DELIVERY ORDER ITEM FORM	t	2024-07-10 02:53:07.860409	2024-07-10 02:53:07.860409
08962a85-84b7-4b50-87ea-d57e87d924bb	PURCHASE_ORDER_CREATE	PURCHASE ORDER CREATE	t	2024-07-10 02:53:07.86041	2024-07-10 02:53:07.86041
45b30169-67ab-4e98-94dd-30b4f6e9d3c9	PURCHASE_ORDER_UPLOAD	PURCHASE ORDER UPLOAD	t	2024-07-10 02:53:07.860411	2024-07-10 02:53:07.860411
303fa09c-a035-4c8c-ae0b-feee49bce02d	PURCHASE_ORDER_ADD_ITEM	PURCHASE ORDER ADD ITEM	t	2024-07-10 02:53:07.860412	2024-07-10 02:53:07.860412
0e6b4487-f1f9-4a4c-910a-197f44c4e0ef	PURCHASE_ORDER_ADD_IMAGE	PURCHASE ORDER ADD IMAGE	t	2024-07-10 02:53:07.860413	2024-07-10 02:53:07.860413
e2672b71-d57c-4b54-8fc9-d96f329a2226	PURCHASE_ORDER_UPDATE	PURCHASE ORDER UPDATE	t	2024-07-10 02:53:07.860414	2024-07-10 02:53:07.860414
5559bc7d-cd37-49ce-8125-fff62c614418	PURCHASE_ORDER_CANCEL	PURCHASE ORDER CANCEL	t	2024-07-10 02:53:07.860415	2024-07-10 02:53:07.860415
41f824c1-66d4-43c0-8980-b9a6c439db61	PURCHASE_ORDER_ONGOING	PURCHASE ORDER ONGOING	t	2024-07-10 02:53:07.860416	2024-07-10 02:53:07.860416
19d906a6-3ce7-40c5-b882-322a26dfeb1a	PURCHASE_ORDER_MARK_COMPLETE	PURCHASE ORDER MARK COMPLETE	t	2024-07-10 02:53:07.860417	2024-07-10 02:53:07.860417
65b4ea55-b0bc-4501-b887-3fb445e1e270	PURCHASE_ORDER_FETCH	PURCHASE ORDER FETCH	t	2024-07-10 02:53:07.860418	2024-07-10 02:53:07.860418
dcf28c10-c5c3-4c8c-926f-59e6d01ab72e	PURCHASE_ORDER_GET	PURCHASE ORDER GET	t	2024-07-10 02:53:07.860419	2024-07-10 02:53:07.860419
bd95fa84-462a-4cb6-adb1-7b9e08a70c1d	PURCHASE_ORDER_DELETE	PURCHASE ORDER DELETE	t	2024-07-10 02:53:07.86042	2024-07-10 02:53:07.86042
1ed42cc6-072f-459b-b50c-650b1bb36ed5	PURCHASE_ORDER_DELETE_ITEM	PURCHASE ORDER DELETE ITEM	t	2024-07-10 02:53:07.860421	2024-07-10 02:53:07.860421
bf0ec2fe-0008-4a75-a796-ef89c6e2da31	PURCHASE_ORDER_DELETE_IMAGE	PURCHASE ORDER DELETE IMAGE	t	2024-07-10 02:53:07.860422	2024-07-10 02:53:07.860422
ddfea278-864d-4423-bac0-ec74ffc59752	ROLE_OPTION_FOR_USER_FORM	ROLE OPTION FOR USER FORM	t	2024-07-10 02:53:07.860423	2024-07-10 02:53:07.860423
ae65c638-a388-4d58-987a-3a05235a6da6	SHOP_ORDER_FETCH	SHOP ORDER FETCH	t	2024-07-10 02:53:07.860424	2024-07-10 02:53:07.860424
5cb4530b-317e-429a-be88-5a52e4030aa0	SHOP_ORDER_GET	SHOP ORDER GET	t	2024-07-10 02:53:07.860425	2024-07-10 02:53:07.860425
d633cc03-c384-4840-ab62-c2e9d96370a2	SSR_WHATSAPP_LOGIN	SSR WHATSAPP LOGIN	t	2024-07-10 02:53:07.860426	2024-07-10 02:53:07.860426
0bf21e0d-9d18-44c2-8c0b-506a32b41f6a	SUPPLIER_CREATE	SUPPLIER CREATE	t	2024-07-10 02:53:07.860427	2024-07-10 02:53:07.860427
1f08e582-1628-480a-b188-084e6eb82bed	SUPPLIER_FETCH	SUPPLIER FETCH	t	2024-07-10 02:53:07.860428	2024-07-10 02:53:07.860428
b700bf02-b0dd-49f9-b90e-d279dddbc0d9	SUPPLIER_GET	SUPPLIER GET	t	2024-07-10 02:53:07.860429	2024-07-10 02:53:07.860429
4bca273e-0832-443e-a20b-7f3fdebc57c0	SUPPLIER_UPDATE	SUPPLIER UPDATE	t	2024-07-10 02:53:07.86043	2024-07-10 02:53:07.86043
acae13ed-8b52-4762-a3e2-03113f5f4ca1	SUPPLIER_DELETE	SUPPLIER DELETE	t	2024-07-10 02:53:07.860431	2024-07-10 02:53:07.860431
0b8b3b00-2e1e-45ba-bc1a-c2cfea663b8a	SUPPLIER_OPTION_FOR_PRODUCT_RECEIVE_FORM	SUPPLIER OPTION FOR PRODUCT RECEIVE FORM	t	2024-07-10 02:53:07.860432	2024-07-10 02:53:07.860432
f1eed559-4bd1-46ec-8125-050bbe8172b6	SUPPLIER_OPTION_FOR_PRODUCT_RECEIVE_FILTER	SUPPLIER OPTION FOR PRODUCT RECEIVE FILTER	t	2024-07-10 02:53:07.860433	2024-07-10 02:53:07.860433
f005d6f4-2170-4ef5-9213-9ada244d29d2	SUPPLIER_TYPE_CREATE	SUPPLIER TYPE CREATE	t	2024-07-10 02:53:07.860434	2024-07-10 02:53:07.860434
2be7ef08-e2fe-4598-ad67-ff731f1b8750	SUPPLIER_TYPE_FETCH	SUPPLIER TYPE FETCH	t	2024-07-10 02:53:07.860435	2024-07-10 02:53:07.860435
7bfa7ca8-f298-4ac7-8aaf-5df4356c2929	SUPPLIER_TYPE_GET	SUPPLIER TYPE GET	t	2024-07-10 02:53:07.860436	2024-07-10 02:53:07.860436
64746194-1303-41d5-8177-70deab18c1fd	SUPPLIER_TYPE_UPDATE	SUPPLIER TYPE UPDATE	t	2024-07-10 02:53:07.860437	2024-07-10 02:53:07.860437
550a041a-8322-44e3-8450-d74480e06633	SUPPLIER_TYPE_DELETE	SUPPLIER TYPE DELETE	t	2024-07-10 02:53:07.860438	2024-07-10 02:53:07.860438
f7af494a-c5f4-4b34-907a-be688598eacb	SUPPLIER_TYPE_OPTION_FOR_SUPPLIER_FORM	SUPPLIER TYPE OPTION FOR SUPPLIER FORM	t	2024-07-10 02:53:07.860439	2024-07-10 02:53:07.860439
567b771d-8e06-47bb-91e3-d1b3ec48d82d	TIKTOK_PRODUCT_CREATE	TIKTOK PRODUCT CREATE	t	2024-07-10 02:53:07.86044	2024-07-10 02:53:07.86044
3e351283-36ed-4e20-a7fc-4b266eefcb13	TIKTOK_PRODUCT_UPLOAD_IMAGE	TIKTOK PRODUCT UPLOAD IMAGE	t	2024-07-10 02:53:07.860441	2024-07-10 02:53:07.860441
01d08933-ebf6-4538-83d5-ada5ced8b35b	TIKTOK_PRODUCT_FETCH_BRANDS	TIKTOK PRODUCT FETCH BRANDS	t	2024-07-10 02:53:07.860442	2024-07-10 02:53:07.860442
8a7cdabc-6086-4f3d-a707-299a2ef8aaa0	TIKTOK_PRODUCT_FETCH_CATEGORIES	TIKTOK PRODUCT FETCH CATEGORIES	t	2024-07-10 02:53:07.860443	2024-07-10 02:53:07.860443
2f9f5814-3da4-4635-b674-8e26544a7ebe	TIKTOK_PRODUCT_GET_RULES	TIKTOK PRODUCT GET RULES	t	2024-07-10 02:53:07.860444	2024-07-10 02:53:07.860444
db8a62b3-d44f-4c30-b24c-ea28e0d15324	TIKTOK_PRODUCT_GET_CATEGORY_ATTRIBUTES	TIKTOK PRODUCT GET CATEGORY ATTRIBUTES	t	2024-07-10 02:53:07.860445	2024-07-10 02:53:07.860445
b127ab20-d3e8-47cb-809a-c3843ef251f8	TIKTOK_PRODUCT_GET	TIKTOK PRODUCT GET	t	2024-07-10 02:53:07.860446	2024-07-10 02:53:07.860446
021fe4e3-196f-49c3-92b4-72975cce0979	TIKTOK_PRODUCT_UPDATE	TIKTOK PRODUCT UPDATE	t	2024-07-10 02:53:07.860447	2024-07-10 02:53:07.860447
73e75a98-a317-4034-a21a-b43974f65262	TIKTOK_PRODUCT_RECOMMENDED_CATEGORY	TIKTOK PRODUCT RECOMMENDED CATEGORY	t	2024-07-10 02:53:07.860448	2024-07-10 02:53:07.860448
80e75eba-d970-4461-bcac-94b02f7871a5	TIKTOK_PRODUCT_RECOMMENDED_ACTIVATE	TIKTOK PRODUCT RECOMMENDED ACTIVATE	t	2024-07-10 02:53:07.860449	2024-07-10 02:53:07.860449
83797eac-545d-43a7-865f-8d3053249be9	TIKTOK_PRODUCT_RECOMMENDED_DEACTIVATE	TIKTOK PRODUCT RECOMMENDED DEACTIVATE	t	2024-07-10 02:53:07.86045	2024-07-10 02:53:07.86045
836e79ba-5bdd-4670-a7a0-6f137ac9db89	TRANSACTION_CHECKOUT_CART	TRANSACTION CHECKOUT CART	t	2024-07-10 02:53:07.860451	2024-07-10 02:53:07.860451
85758192-2ecc-45e8-908d-409bb20272d9	TRANSACTION_GET	TRANSACTION GET	t	2024-07-10 02:53:07.860452	2024-07-10 02:53:07.860452
23eded05-4e26-45ce-bb47-ca9df2007d54	USER_CREATE	USER CREATE	t	2024-07-10 02:53:07.860453	2024-07-10 02:53:07.860453
e594d5cd-fd64-45ac-bb09-074b460b6e3f	USER_FETCH	USER FETCH	t	2024-07-10 02:53:07.860454	2024-07-10 02:53:07.860454
a3ba044f-d1f0-48ee-b0ff-df4fb87c2616	USER_GET	USER GET	t	2024-07-10 02:53:07.860455	2024-07-10 02:53:07.860455
47ca9a63-550c-4203-97e7-5225a45b6459	USER_UPDATE	USER UPDATE	t	2024-07-10 02:53:07.860456	2024-07-10 02:53:07.860456
08dcf0d1-a01e-4eef-95ce-2bbc783eed9e	USER_UPDATE_PASSWORD	USER UPDATE PASSWORD	t	2024-07-10 02:53:07.860457	2024-07-10 02:53:07.860457
ec4e44fa-a6f1-4cdf-8baa-85ebac7912a7	USER_UPDATE_ACTIVE	USER UPDATE ACTIVE	t	2024-07-10 02:53:07.860458	2024-07-10 02:53:07.860458
2b496e84-cf57-4a5b-abd4-0ff652fe4138	USER_UPDATE_INACTIVE	USER UPDATE INACTIVE	t	2024-07-10 02:53:07.860459	2024-07-10 02:53:07.860459
6e6e2de2-156c-4af1-8f15-1c9df5d1e074	USER_ADD_ROLE	USER ADD ROLE	t	2024-07-10 02:53:07.86046	2024-07-10 02:53:07.86046
634c7e46-3f5b-4f96-8cbd-4e2f49fc1209	USER_DELETE_ROLE	USER DELETE ROLE	t	2024-07-10 02:53:07.860461	2024-07-10 02:53:07.860461
ca5ea509-e6d8-4dab-85c6-0e15a66cab24	USER_OPTION_FOR_CASHIER_SESSION_FILTER	USER OPTION FOR CASHIER SESSION FILTER	t	2024-07-10 02:53:07.860462	2024-07-10 02:53:07.860462
eda3e8f9-a887-4e53-b4f0-63cc1da481b4	USER_OPTION_FOR_DELIVERY_ORDER_DRIVER_FORM	USER OPTION FOR DELIVERY ORDER DRIVER FORM	t	2024-07-10 02:53:07.860463	2024-07-10 02:53:07.860463
88f3e7b5-3fd4-431c-90e2-cd38c3232297	USER_OPTION_FOR_PRODUCT_STOCK_ADJUSTMENT_FILTER	USER OPTION FOR PRODUCT STOCK ADJUSTMENT FILTER	t	2024-07-10 02:53:07.860464	2024-07-10 02:53:07.860464
e83fbf50-1a29-4133-af74-f32834e21b02	UNIT_CREATE	UNIT CREATE	t	2024-07-10 02:53:07.860465	2024-07-10 02:53:07.860465
4465ae1a-59a2-4814-b92d-660298bc62bc	UNIT_FETCH	UNIT FETCH	t	2024-07-10 02:53:07.860466	2024-07-10 02:53:07.860466
352aecd8-b7d3-46e9-886c-17a6fc0c1176	UNIT_GET	UNIT GET	t	2024-07-10 02:53:07.860467	2024-07-10 02:53:07.860467
5b8e13c3-1a1a-4847-a246-e0cfd27326b6	UNIT_UPDATE	UNIT UPDATE	t	2024-07-10 02:53:07.860468	2024-07-10 02:53:07.860468
2872765f-6393-4de6-a100-290f401e64a1	UNIT_DELETE	UNIT DELETE	t	2024-07-10 02:53:07.860469	2024-07-10 02:53:07.860469
257e2fcd-8962-480b-8723-9cfdd215e6b8	UNIT_OPTION_FOR_PRODUCT_UNIT_FORM	UNIT OPTION FOR PRODUCT UNIT FORM	t	2024-07-10 02:53:07.86047	2024-07-10 02:53:07.86047
7cf3ff8c-315d-450b-aa52-6bc7d03762bf	UNIT_OPTION_FOR_PRODUCT_UNIT_TO_UNIT_FORM	UNIT OPTION FOR PRODUCT UNIT TO UNIT FORM	t	2024-07-10 02:53:07.860471	2024-07-10 02:53:07.860471
27c05dc9-6ef7-4d03-ba09-c3a0e9a4459e	WHATSAPP_IS_LOGGED_IN	WHATSAPP IS LOGGED IN	t	2024-07-10 02:53:07.860472	2024-07-10 02:53:07.860472
bdac9253-ef8c-4b81-95a8-394d7a818eca	WHATSAPP_PRODUCT_PRICE_CHANGE_BROADCAST	WHATSAPP PRODUCT PRICE CHANGE BROADCAST	t	2024-07-10 02:53:07.860473	2024-07-10 02:53:07.860473
30c4fcc4-b630-4afb-b117-2a1547cfd9c3	WHATSAPP_CUSTOMER_DEBT_BROADCAST	WHATSAPP CUSTOMER DEBT BROADCAST	t	2024-07-10 02:53:07.860474	2024-07-10 02:53:07.860474
9bee79e5-88b1-4f69-906a-7227260c5e35	WHATSAPP_CUSTOMER_TYPE_DISCOUNT_BROADCAST	WHATSAPP CUSTOMER TYPE DISCOUNT BROADCAST	t	2024-07-10 02:53:07.860475	2024-07-10 02:53:07.860475
455150f8-a28c-4644-b6c9-bb53639a44ae	WHATSAPP_LOGOUT	WHATSAPP LOGOUT	t	2024-07-10 02:53:07.860476	2024-07-10 02:53:07.860476
\.


--
-- Data for Name: product_discounts; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.product_discounts (id, product_id, minimum_qty, is_active, discount_percentage, discount_amount, created_at, updated_at) FROM stdin;
08bbc40c-01ea-459a-a479-f60c00ec0d75	e1bf0592-7850-4602-a740-6aae98dfd281	1.00	t	\N	10000.00	2024-06-12 08:45:42.946581	2024-06-12 08:45:42.946581
93ee7fa4-7164-4a1b-a3bc-0d9c4af4f84f	69a3f894-1ff2-4f61-97c2-3c957eea7914	10.00	t	\N	999.00	2024-06-27 08:56:09.052568	2024-06-27 08:56:09.052568
b1d760ab-1c40-4e33-b743-9199a2ee3cfc	d0813034-19a1-406f-af25-c18d2f301614	10.00	t	\N	10000.00	2024-06-28 00:24:06.986528	2024-06-28 00:24:06.986528
\.


--
-- Data for Name: product_receive_images; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.product_receive_images (id, product_receive_id, file_id, description, created_at, updated_at) FROM stdin;
127a1801-0d09-4f23-b574-81064bc3ea9b	d8984d7a-307e-46d7-b8ae-bd429cc23784	7f274f6c-76f0-4e49-bf54-4805464396a5	\N	2024-06-27 08:59:05.168307	2024-06-27 08:59:05.168307
b5e9deda-ae4a-4d09-a420-cd1cb05d4d04	c7bc1479-51e8-462c-9b30-0824a574b17a	2ceca493-79a1-46d8-8e5b-7ef3b2e2e8d0	\N	2024-06-27 09:21:00.84309	2024-06-27 09:21:00.84309
\.


--
-- Data for Name: product_receive_items; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.product_receive_items (id, product_receive_id, product_unit_id, user_id, qty_eligible, qty_received, scale_to_base, price_per_unit, created_at, updated_at) FROM stdin;
2a09e059-93c9-42bb-83cc-b0f40855859e	12450cfa-19be-4272-ac6c-1389b98e79da	cec71c3c-5a20-461b-9370-051ab3eeba76	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	12.00	12.00	1.00	10000.00	2024-06-07 15:45:18.211341	2024-06-07 15:45:18.211341
116f7ebe-4dc6-4332-a781-f67edeac8634	ea86af8e-c1a8-41eb-9028-dd87eb92a405	fa437b9d-89cd-47a6-8877-2a8e91336450	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	100.00	100.00	12.00	10000.00	2024-06-20 02:22:05.242937	2024-06-20 02:22:05.242937
02e360c8-9e56-4820-bb5a-2b625de24dfe	d8984d7a-307e-46d7-b8ae-bd429cc23784	d3e101c6-9603-4370-a4db-ea52b450f5b1	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	10.00	10.00	20.00	70000.00	2024-06-27 08:58:21.521037	2024-06-27 08:58:21.521037
014d40fb-cbb3-468a-b1df-64b530c834a7	d8984d7a-307e-46d7-b8ae-bd429cc23784	a28b11ad-6f08-4607-badd-d481c0e6be4c	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	18.00	20.00	1.00	8000.00	2024-06-27 08:58:21.521036	2024-06-27 08:58:54.99384
b03c8982-99c8-4652-8536-6358ccd7c0e8	c7bc1479-51e8-462c-9b30-0824a574b17a	d3e101c6-9603-4370-a4db-ea52b450f5b1	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	10.00	10.00	20.00	700000.00	2024-06-27 09:20:17.808235	2024-06-27 09:20:17.808235
8201a1ed-454a-4c16-9fa9-34ca4875cce7	c7bc1479-51e8-462c-9b30-0824a574b17a	a28b11ad-6f08-4607-badd-d481c0e6be4c	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	8.00	10.00	1.00	80000.00	2024-06-27 09:20:17.808234	2024-06-27 09:20:43.839787
4b45feed-b642-46fc-b2c9-68bd47bfcc6f	204b6093-103c-46da-9bf9-a07359e7c158	151c9e14-4403-4ed6-9973-f401a67841f2	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	10.00	10.00	10.00	700000.00	2024-06-28 00:25:58.186368	2024-06-28 00:25:58.186368
848967a5-3a02-42ba-a3a0-9132b1aebd80	204b6093-103c-46da-9bf9-a07359e7c158	61218f4d-b5ac-41db-8db6-1f3c3f18a10e	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	8.00	10.00	1.00	80000.00	2024-06-28 00:25:58.186367	2024-06-28 00:26:18.426187
15ffd506-c78c-490e-a8c0-70c283002643	824e5e08-151f-4143-a5b7-70e63d96acae	151c9e14-4403-4ed6-9973-f401a67841f2	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	5.00	5.00	10.00	1000000.00	2024-07-01 11:10:59.493052	2024-07-01 11:10:59.493052
74e1634a-dbd8-47f9-8c5c-b167c50527ea	2ccac487-1cfc-4be8-a7b9-aa3081058f4a	cec71c3c-5a20-461b-9370-051ab3eeba76	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	2.00	5.00	1.00	9000.00	2024-06-07 02:54:03.559902	2024-07-08 15:26:10.71275
f912427b-4c5d-45e0-8eda-b0036bb04a7d	1433e671-a380-4d75-ac1c-7aedda724f93	f8fd2e10-1a96-44ae-9898-50fa57d8dcf9	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	100.00	100.00	1.00	10000.00	2024-07-10 02:04:39.129498	2024-07-10 02:04:39.129498
84dc808e-a198-4107-8cde-3c36bb3fa975	1433e671-a380-4d75-ac1c-7aedda724f93	f8fd2e10-1a96-44ae-9898-50fa57d8dcf9	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	10.00	10.00	1.00	20000.00	2024-07-10 02:04:39.129499	2024-07-10 02:04:39.129499
63afadec-5414-40b8-8585-5ad8bee664b7	969955f9-5964-4e38-a05c-4033d02f1f12	f8fd2e10-1a96-44ae-9898-50fa57d8dcf9	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	100.00	100.00	1.00	10000.00	2024-07-10 02:18:44.820968	2024-07-10 02:18:44.820968
4b507e94-fe07-4daf-8b7f-e9b4288a93f1	969955f9-5964-4e38-a05c-4033d02f1f12	f8fd2e10-1a96-44ae-9898-50fa57d8dcf9	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	10.00	10.00	1.00	20000.00	2024-07-10 02:18:44.820969	2024-07-10 02:18:44.820969
59e9deb2-167c-4d2e-a187-18a98d983156	7ad0043b-a3dc-46b8-8e9e-14aded64b9f6	f8fd2e10-1a96-44ae-9898-50fa57d8dcf9	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	10.00	10.00	1.00	30000.00	2024-07-10 02:55:52.936371	2024-07-10 02:55:52.936371
\.


--
-- Data for Name: product_receive_return_images; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.product_receive_return_images (id, product_receive_return_id, file_id, created_at, updated_at) FROM stdin;
\.


--
-- Data for Name: product_receive_returns; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.product_receive_returns (id, product_receive_id, user_id, description, created_at, updated_at) FROM stdin;
\.


--
-- Data for Name: product_receives; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.product_receives (id, purchase_order_id, supplier_id, user_id, invoice_number, date, status, total_price, created_at, updated_at) FROM stdin;
2ccac487-1cfc-4be8-a7b9-aa3081058f4a	b45d509c-7440-43ab-97d5-4296321ebe22	e0225ca5-4094-47e9-8494-43795dfddd97	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	wsadeaf	2024-06-07	PENDING	45000.00	2024-06-07 02:54:03.555402	2024-06-07 02:54:03.555402
12450cfa-19be-4272-ac6c-1389b98e79da	407abd68-9449-4f1a-aed6-1fd38fbd90f3	e0225ca5-4094-47e9-8494-43795dfddd97	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	INV/001/2024	2024-06-07	COMPLETED	120000.00	2024-06-07 15:45:18.209645	2024-06-07 15:45:34.230811
ea86af8e-c1a8-41eb-9028-dd87eb92a405	b41e582b-aff6-400c-8d55-b7342c1a5808	e0225ca5-4094-47e9-8494-43795dfddd97	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	rehrehsre	2024-06-20	COMPLETED	1000000.00	2024-06-20 02:22:05.240331	2024-06-20 02:22:16.380095
d8984d7a-307e-46d7-b8ae-bd429cc23784	29d46083-0872-4a3a-96a3-510ad7613a10	e6815351-8cd2-4471-a81e-ac21eb65e7ea	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	C001-98252/512512/w521	2024-06-27	COMPLETED	860000.00	2024-06-27 08:58:21.519642	2024-06-27 08:59:44.330447
c7bc1479-51e8-462c-9b30-0824a574b17a	6c227a2d-c6fd-48b4-8aa8-c4045a8981fb	e6815351-8cd2-4471-a81e-ac21eb65e7ea	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	C001-598102521/15215125/215	2024-06-27	COMPLETED	7800000.00	2024-06-27 09:20:17.801957	2024-06-27 09:21:13.523091
204b6093-103c-46da-9bf9-a07359e7c158	86f2a2b3-eda0-458f-a52c-455a81d91818	48fba4fe-787a-45fb-8c65-3ff6418b09a6	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	INVC-C001-24215/125215	2024-06-28	COMPLETED	7800000.00	2024-06-28 00:25:58.184829	2024-06-28 00:26:44.715259
824e5e08-151f-4143-a5b7-70e63d96acae	037a8901-ad4a-40e9-9b34-088e78fcad06	e0225ca5-4094-47e9-8494-43795dfddd97	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	a	2024-07-01	COMPLETED	5000000.00	2024-07-01 11:10:59.488472	2024-07-01 11:11:30.173415
1433e671-a380-4d75-ac1c-7aedda724f93	d394c940-2349-4179-8c3a-c2755bb84419	e0225ca5-4094-47e9-8494-43795dfddd97	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	INv-0111	2024-07-10	COMPLETED	1200000.00	2024-07-10 02:04:39.123374	2024-07-10 02:04:45.295367
969955f9-5964-4e38-a05c-4033d02f1f12	66ad8587-155b-4f3c-897c-5cdd93cb3799	e0225ca5-4094-47e9-8494-43795dfddd97	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	INV-001	2024-07-10	COMPLETED	1200000.00	2024-07-10 02:18:44.819851	2024-07-10 02:18:50.205445
7ad0043b-a3dc-46b8-8e9e-14aded64b9f6	f67ca56b-2eba-4835-b50a-9a53b88a021c	48fba4fe-787a-45fb-8c65-3ff6418b09a6	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	INv-002	2024-07-10	COMPLETED	300000.00	2024-07-10 02:55:52.931804	2024-07-10 02:55:58.495074
\.


--
-- Data for Name: product_return_images; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.product_return_images (id, product_return_id, file_id, description, created_at, updated_at) FROM stdin;
14c4de76-4c97-4587-acb5-d305cea5368e	b977aa16-a0be-42b7-a2f6-0e6a944bfda8	95500471-ec6c-40d5-8fb4-0577d6b7c188	\N	2024-06-07 07:14:06.574709	2024-06-07 07:14:06.574709
2606eb5e-dc83-41ad-8db1-812fbf080c55	19ca44b3-1537-49fe-a68d-d472e909b046	7e238d0a-d192-4ad8-a13c-3dd150ea5a21	\N	2024-06-27 09:00:47.585891	2024-06-27 09:00:47.585891
90547f4f-3568-4889-b3e0-bef9a1452b13	a4ec4f6c-b3f1-4420-8cce-ed7b79d48914	8da9c904-5d27-4b58-b18c-744efb23e03f	\N	2024-06-27 09:22:08.323489	2024-06-27 09:22:08.323489
da3c8e8d-69cc-490f-9a96-19b32af45148	c1d72997-80c0-464d-81e4-d6c58b958064	20621060-f268-4863-8246-1ec1187789e5	Gambar Cat Kaleng Biru	2024-07-08 15:03:37.217314	2024-07-08 15:03:37.217314
\.


--
-- Data for Name: product_return_items; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.product_return_items (id, product_return_id, product_unit_id, qty, scale_to_base, base_cost_price, created_at, updated_at) FROM stdin;
f4455f59-f24f-4ec1-bbba-b67757ec9f29	b977aa16-a0be-42b7-a2f6-0e6a944bfda8	cec71c3c-5a20-461b-9370-051ab3eeba76	4.00	1.00	10000.00	2024-06-07 08:03:13.198714	2024-06-07 08:03:13.198714
81e21941-a40e-45bb-9c26-a6baaeb15062	19ca44b3-1537-49fe-a68d-d472e909b046	a28b11ad-6f08-4607-badd-d481c0e6be4c	2.00	1.00	8000.00	2024-06-27 09:00:41.542613	2024-06-27 09:00:41.542613
532a19f0-1a0d-4b4e-9f43-f3a3d471c80b	a4ec4f6c-b3f1-4420-8cce-ed7b79d48914	a28b11ad-6f08-4607-badd-d481c0e6be4c	5.00	1.00	36730.77	2024-06-27 09:21:58.405719	2024-06-27 09:21:58.405719
07c869e1-6be7-45b3-9550-b96ef6d71a94	a6d2658d-8ad4-455a-96e6-e30759ea2cd6	61218f4d-b5ac-41db-8db6-1f3c3f18a10e	3.00	1.00	70740.74	2024-06-28 00:27:27.0193	2024-06-28 00:27:27.0193
8eff1538-96c1-4e85-a115-ea6af1a5bb0a	7ec86114-2327-42b1-95ce-9ff4b664d9fe	61218f4d-b5ac-41db-8db6-1f3c3f18a10e	10.00	1.00	83902.43	2024-07-01 11:13:54.27789	2024-07-01 11:13:54.27789
6095082b-127b-4a4c-bd06-0221cda09a9b	c1d72997-80c0-464d-81e4-d6c58b958064	cec71c3c-5a20-461b-9370-051ab3eeba76	1.00	1.00	1785.67	2024-07-08 14:56:40.628704	2024-07-08 14:56:40.628704
\.


--
-- Data for Name: product_returns; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.product_returns (id, supplier_id, user_id, invoice_number, date, status, created_at, updated_at) FROM stdin;
b977aa16-a0be-42b7-a2f6-0e6a944bfda8	e0225ca5-4094-47e9-8494-43795dfddd97	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	aeghe	2024-06-07	COMPLETED	2024-06-07 06:57:53.909823	2024-06-07 08:03:57.701614
19ca44b3-1537-49fe-a68d-d472e909b046	e6815351-8cd2-4471-a81e-ac21eb65e7ea	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	C001-RETUR-123129/safasfas	2024-06-27	COMPLETED	2024-06-27 09:00:33.668994	2024-06-27 09:00:53.552278
a4ec4f6c-b3f1-4420-8cce-ed7b79d48914	e6815351-8cd2-4471-a81e-ac21eb65e7ea	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	C001-RETUR-125215/215215	2024-06-27	COMPLETED	2024-06-27 09:21:52.890851	2024-06-27 09:22:12.762777
a6d2658d-8ad4-455a-96e6-e30759ea2cd6	48fba4fe-787a-45fb-8c65-3ff6418b09a6	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	RETUR-C002-125125/1251	2024-06-28	COMPLETED	2024-06-28 00:27:18.90876	2024-06-28 00:27:33.638184
7ec86114-2327-42b1-95ce-9ff4b664d9fe	e0225ca5-4094-47e9-8494-43795dfddd97	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	AAA	2024-06-08	COMPLETED	2024-06-07 17:07:27.894766	2024-07-01 11:14:01.936264
c1d72997-80c0-464d-81e4-d6c58b958064	48fba4fe-787a-45fb-8c65-3ff6418b09a6	33392eb4-0f87-43c6-9893-9c014fe6d561	INV/001/2024	2024-07-08	PENDING	2024-07-08 14:56:19.946746	2024-07-08 14:56:40.626863
\.


--
-- Data for Name: product_stock_adjustments; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.product_stock_adjustments (id, user_id, product_stock_id, previous_qty, updated_qty, created_at, updated_at) FROM stdin;
cdd8a8b5-6bdc-4fed-9951-96ca68a927d2	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	6feb27eb-d326-4234-ae0b-db810680fa57	0.00	5.00	2024-06-06 14:25:49.963865	2024-06-06 14:25:49.963865
240d5519-ea57-45fc-b67d-64984c9c04c3	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	6feb27eb-d326-4234-ae0b-db810680fa57	5.00	10.00	2024-06-07 07:44:58.275191	2024-06-07 07:44:58.275191
effb28d9-58f5-4cad-b4d2-56c087334370	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	6feb27eb-d326-4234-ae0b-db810680fa57	10.00	0.00	2024-06-07 07:45:48.498057	2024-06-07 07:45:48.498057
06da1d46-3deb-4244-a687-f3e8a68af4bc	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	6feb27eb-d326-4234-ae0b-db810680fa57	0.00	10.00	2024-06-07 07:57:56.498767	2024-06-07 07:57:56.498767
f61e10f7-0a99-420a-8b8d-f0c18c875158	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	6feb27eb-d326-4234-ae0b-db810680fa57	6.00	10.00	2024-06-07 08:09:47.817874	2024-06-07 08:09:47.817874
67caa1b3-b5c0-491f-bad1-7b2651db0ff7	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	6feb27eb-d326-4234-ae0b-db810680fa57	10.00	15.00	2024-06-07 08:21:35.191597	2024-06-07 08:21:35.191597
4f87ab54-d05a-497a-99cf-a63fdff85a3f	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	6feb27eb-d326-4234-ae0b-db810680fa57	15.00	20.00	2024-06-07 08:21:50.993825	2024-06-07 08:21:50.993825
1a343d0b-b5c6-4931-9c9d-ed3b3d895f0b	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	6feb27eb-d326-4234-ae0b-db810680fa57	20.00	82.00	2024-06-07 15:48:02.327088	2024-06-07 15:48:02.327088
3627053c-c2fc-41e5-a236-924d7be7b19f	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	420a8e9a-4918-4040-8e28-25a6ef23b66d	0.00	5.00	2024-06-27 08:55:20.323518	2024-06-27 08:55:20.323518
ba383cce-79b1-4803-a509-2c86c31eb5b1	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	420a8e9a-4918-4040-8e28-25a6ef23b66d	5.00	0.00	2024-06-27 08:55:29.798025	2024-06-27 08:55:29.798025
0adb79c1-f019-43f6-8f39-bf1a57c7c32f	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	fafe6f0b-afe9-4f2c-b730-0196a773bdb6	105.00	100.00	2024-06-28 00:27:52.783583	2024-06-28 00:27:52.783583
9b57426f-a5d3-4239-8371-bfad037cb0f6	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	fafe6f0b-afe9-4f2c-b730-0196a773bdb6	58.00	60.00	2024-07-01 09:17:33.318075	2024-07-01 09:17:33.318075
e5a0ccfd-f8b5-4f4a-8ad9-b73ddcfbc464	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	fafe6f0b-afe9-4f2c-b730-0196a773bdb6	57.00	55.00	2024-07-01 09:23:55.913955	2024-07-01 09:23:55.913955
99321d10-f8b4-4aaa-adef-2a4aa0070860	33392eb4-0f87-43c6-9893-9c014fe6d561	2930d942-cab3-4da2-be1c-005e586f9b9c	0.00	10.00	2024-07-08 16:01:58.244941	2024-07-08 16:01:58.244941
3223afbb-34dd-4460-af46-eab8cb61b557	33392eb4-0f87-43c6-9893-9c014fe6d561	2930d942-cab3-4da2-be1c-005e586f9b9c	9.00	20.00	2024-07-08 16:03:24.074861	2024-07-08 16:03:24.074861
d5b6b4c0-3014-41b9-8f22-200284169cbf	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	43f1210d-5210-4cae-80ce-ad846341b38f	109.00	0.00	2024-07-10 02:17:51.52743	2024-07-10 02:17:51.52743
\.


--
-- Data for Name: product_stock_mutations; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.product_stock_mutations (id, product_unit_id, type, identifier_id, qty, scale_to_base, base_qty_left, base_cost_price, mutated_at, created_at, updated_at) FROM stdin;
edf109d9-67e5-4fe0-9dd1-c8761fa93ae2	a28b11ad-6f08-4607-badd-d481c0e6be4c	PRODUCT_RECEIVE_ITEM	014d40fb-cbb3-468a-b1df-64b530c834a7	18.00	1.00	13.00	8000.00	2024-06-27 08:59:42.132424	2024-06-27 08:59:44.331751	2024-06-27 09:33:06.747125
af7fa561-8e96-4504-aa0d-c85e55e44636	4ca122af-8969-44d6-ac51-fd3a1a8c504b	PRODUCT_STOCK_ADJUSTMENT	2930d942-cab3-4da2-be1c-005e586f9b9c	10.00	1.00	8.00	10000.00	2024-07-08 16:01:58.237703	2024-07-08 16:01:58.241231	2024-07-08 16:03:43.14582
37f0992c-b818-4246-8a38-11df55f184bf	61218f4d-b5ac-41db-8db6-1f3c3f18a10e	PRODUCT_RECEIVE_ITEM	848967a5-3a02-42ba-a3a0-9132b1aebd80	8.00	1.00	0.00	80000.00	2024-06-28 00:26:42.731646	2024-06-28 00:26:44.716811	2024-06-28 00:33:48.925174
4e89966d-9c71-4d16-965c-e4a4faead0da	f8fd2e10-1a96-44ae-9898-50fa57d8dcf9	PRODUCT_RECEIVE_ITEM	f912427b-4c5d-45e0-8eda-b0036bb04a7d	100.00	1.00	99.00	10000.00	2024-07-10 02:04:45.219672	2024-07-10 02:04:45.299689	2024-07-10 02:16:43.112255
9b5bf858-8963-4559-932a-992d751675fe	d3e101c6-9603-4370-a4db-ea52b450f5b1	PRODUCT_RECEIVE_ITEM	02e360c8-9e56-4820-bb5a-2b625de24dfe	10.00	20.00	200.00	3500.00	2024-06-27 08:59:42.132424	2024-06-27 08:59:44.33175	2024-06-27 08:59:44.33175
34ffa03c-e161-4fd8-9a68-9869a9a1ac7f	cec71c3c-5a20-461b-9370-051ab3eeba76	PRODUCT_STOCK_ADJUSTMENT	6feb27eb-d326-4234-ae0b-db810680fa57	10.00	1.00	0.00	10000.00	2024-06-07 07:57:56.490658	2024-06-07 07:57:56.494926	2024-06-18 06:59:49.05153
48c3f232-a799-4610-b2f9-65eb90c44f81	f8fd2e10-1a96-44ae-9898-50fa57d8dcf9	PRODUCT_RECEIVE_ITEM	63afadec-5414-40b8-8585-5ad8bee664b7	100.00	1.00	100.00	10000.00	2024-07-10 02:18:50.131091	2024-07-10 02:18:50.209425	2024-07-10 02:18:50.209425
24fea276-add8-4501-9657-6ed8505ef7a4	d3e101c6-9603-4370-a4db-ea52b450f5b1	PRODUCT_RECEIVE_ITEM	b03c8982-99c8-4652-8536-6358ccd7c0e8	10.00	20.00	200.00	35000.00	2024-06-27 09:21:11.222336	2024-06-27 09:21:13.525663	2024-06-27 09:21:13.525663
ff592ae5-5e78-4b01-89bf-f83b720f7c85	a28b11ad-6f08-4607-badd-d481c0e6be4c	PRODUCT_RECEIVE_ITEM	8201a1ed-454a-4c16-9fa9-34ca4875cce7	8.00	1.00	8.00	80000.00	2024-06-27 09:21:11.222336	2024-06-27 09:21:13.525664	2024-06-27 09:21:13.525664
feca7350-9cb9-4e02-9927-de3bfb01d31d	f8fd2e10-1a96-44ae-9898-50fa57d8dcf9	PRODUCT_RECEIVE_ITEM	4b507e94-fe07-4daf-8b7f-e9b4288a93f1	10.00	1.00	10.00	20000.00	2024-07-10 02:18:50.131091	2024-07-10 02:18:50.209426	2024-07-10 02:18:50.209426
0d2e810d-f4e7-4409-9a1a-44a547eb0623	cec71c3c-5a20-461b-9370-051ab3eeba76	PRODUCT_STOCK_ADJUSTMENT	6feb27eb-d326-4234-ae0b-db810680fa57	4.00	1.00	0.00	10000.00	2024-06-07 08:09:47.804586	2024-06-07 08:09:47.812727	2024-06-18 07:18:35.320465
1aab4a81-94d1-44fe-a13d-70e3381e6bb1	fa437b9d-89cd-47a6-8877-2a8e91336450	PRODUCT_RECEIVE_ITEM	116f7ebe-4dc6-4332-a781-f67edeac8634	100.00	12.00	1200.00	833.33	2024-06-20 02:22:14.210777	2024-06-20 02:22:16.382825	2024-06-20 02:22:16.382825
ce8da5b4-8e57-4de5-9978-85aec06743ba	61218f4d-b5ac-41db-8db6-1f3c3f18a10e	PRODUCT_STOCK_ADJUSTMENT	fafe6f0b-afe9-4f2c-b730-0196a773bdb6	2.00	1.00	2.00	10000.00	2024-07-01 09:17:31.233285	2024-07-01 09:17:33.313734	2024-07-01 09:17:33.313734
15763865-7a08-4fdd-a7a6-69c0795482e0	cec71c3c-5a20-461b-9370-051ab3eeba76	PRODUCT_STOCK_ADJUSTMENT	6feb27eb-d326-4234-ae0b-db810680fa57	5.00	1.00	0.00	9000.00	2024-06-07 08:21:35.174145	2024-06-07 08:21:35.18844	2024-06-20 02:22:46.917613
a17cf21f-6948-4efa-bf28-9c817289e182	cec71c3c-5a20-461b-9370-051ab3eeba76	DELIVERY_ORDER_ITEM_COST_CANCEL	14017fb3-9764-462b-bd77-602db7a9def3	10.00	1.00	10.00	1785.67	2024-06-20 03:59:35.098769	2024-06-20 03:59:35.246976	2024-06-20 03:59:35.246976
637d70da-ff6c-4c46-bbd5-08362bd634f7	cec71c3c-5a20-461b-9370-051ab3eeba76	DELIVERY_ORDER_ITEM_COST_CANCEL	3bd98e53-04ec-469d-88d7-cf4ab82bcde3	5.00	1.00	5.00	1785.67	2024-06-20 03:59:35.098769	2024-06-20 03:59:35.246977	2024-06-20 03:59:35.246977
371057ec-f05f-4c60-ba66-aa641c30c609	f8fd2e10-1a96-44ae-9898-50fa57d8dcf9	PRODUCT_RECEIVE_ITEM	84dc808e-a198-4107-8cde-3c36bb3fa975	10.00	1.00	9.00	20000.00	2024-07-10 02:04:45.219672	2024-07-10 02:04:45.29969	2024-07-10 02:19:11.448697
a9b6f157-c82f-4dc4-9d54-d02353b85fd0	cec71c3c-5a20-461b-9370-051ab3eeba76	PRODUCT_STOCK_ADJUSTMENT	6feb27eb-d326-4234-ae0b-db810680fa57	5.00	1.00	0.00	9000.00	2024-06-07 08:21:50.988991	2024-06-07 08:21:50.992695	2024-06-20 09:46:58.224581
64081618-29ec-4b5a-b8a0-ed806d11164a	cec71c3c-5a20-461b-9370-051ab3eeba76	DELIVERY_ORDER_ITEM_COST_CANCEL	fab754b0-9484-4e26-bcaa-aa39615eb685	10.00	1.00	10.00	1785.67	2024-06-20 09:49:26.974513	2024-06-20 09:49:27.123449	2024-06-20 09:49:27.123449
4471e1eb-d33b-44b4-b631-82f245700667	cec71c3c-5a20-461b-9370-051ab3eeba76	DELIVERY_ORDER_ITEM_COST_CANCEL	30396a3b-54c8-4545-aa7a-198f8ff361ab	5.00	1.00	5.00	1785.67	2024-06-20 09:49:26.974513	2024-06-20 09:49:27.12345	2024-06-20 09:49:27.12345
9813daa7-ec98-4896-9779-3caf8b1fcca1	151c9e14-4403-4ed6-9973-f401a67841f2	PRODUCT_RECEIVE_ITEM	15ffd506-c78c-490e-a8c0-70c283002643	5.00	10.00	50.00	100000.00	2024-07-01 11:11:27.63402	2024-07-01 11:11:30.17732	2024-07-01 11:11:30.17732
7d59ea9c-b270-4735-89fb-8a97d5688f4c	f8fd2e10-1a96-44ae-9898-50fa57d8dcf9	PRODUCT_RECEIVE_ITEM	59e9deb2-167c-4d2e-a187-18a98d983156	10.00	1.00	10.00	30000.00	2024-07-10 02:55:58.41572	2024-07-10 02:55:58.502371	2024-07-10 02:55:58.502371
b06613b2-d758-4b7b-803d-45d2ac84d905	151c9e14-4403-4ed6-9973-f401a67841f2	PRODUCT_RECEIVE_ITEM	4b45feed-b642-46fc-b2c9-68bd47bfcc6f	10.00	10.00	74.00	70000.00	2024-06-28 00:26:42.731646	2024-06-28 00:26:44.71681	2024-07-08 12:29:19.443678
b633d7c5-9138-4499-bff3-7ae13434fa68	a28b11ad-6f08-4607-badd-d481c0e6be4c	PRODUCT_STOCK_ADJUSTMENT	420a8e9a-4918-4040-8e28-25a6ef23b66d	5.00	1.00	0.00	8000.00	2024-06-27 08:55:20.315849	2024-06-27 08:55:20.320087	2024-06-27 09:23:18.284727
e98c0f7e-468f-408b-b294-7d02e204494f	cec71c3c-5a20-461b-9370-051ab3eeba76	PRODUCT_STOCK_ADJUSTMENT	6feb27eb-d326-4234-ae0b-db810680fa57	62.00	1.00	58.00	100000.00	2024-06-07 15:48:02.322364	2024-06-07 15:48:02.325768	2024-07-08 15:54:59.914019
c49caffd-ef71-4013-8830-f9a6638f681d	cec71c3c-5a20-461b-9370-051ab3eeba76	PRODUCT_RECEIVE_ITEM	2a09e059-93c9-42bb-83cc-b0f40855859e	12.00	1.00	0.00	10000.00	2024-06-07 15:45:31.985179	2024-06-07 15:45:34.234283	2024-06-27 09:33:06.745455
a80e5c00-0e52-4966-bb2c-9c8fe1a7f27d	4ca122af-8969-44d6-ac51-fd3a1a8c504b	PRODUCT_STOCK_ADJUSTMENT	2930d942-cab3-4da2-be1c-005e586f9b9c	11.00	1.00	11.00	20000.00	2024-07-08 16:03:24.070917	2024-07-08 16:03:24.074108	2024-07-08 16:03:24.074108
\.


--
-- Data for Name: product_stocks; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.product_stocks (id, product_id, qty, base_cost_price, created_at, updated_at) FROM stdin;
420a8e9a-4918-4040-8e28-25a6ef23b66d	69a3f894-1ff2-4f61-97c2-3c957eea7914	118.00	36730.77	2024-06-27 08:52:04.699442	2024-06-27 09:33:06.753441
a5178edd-6356-4eef-a90f-2ce0666099f4	96773beb-3646-4078-893f-19a278e98a30	0.00	0.00	2024-06-07 07:58:48.42545	2024-06-07 07:58:48.42545
3c093cf8-90b6-4363-8905-440cab3876b7	bcfd2b82-8125-4e06-b6d8-2e42c70fdd74	0.00	0.00	2024-06-29 15:38:14.342217	2024-06-29 15:38:14.342217
0a6be06f-6eaf-4057-9a34-abb74b657208	2ea93d98-d8d5-4ce2-b13f-00ac30f863fc	0.00	0.00	2024-07-03 03:06:00.062015	2024-07-03 03:06:00.062015
4270b4ae-7f5b-44f2-8085-3fa4b8cc3987	32ca75e1-7d0f-4fc0-ad09-78d16d820499	0.00	0.00	2024-07-03 03:06:43.25717	2024-07-03 03:06:43.25717
ab899915-3e0c-4ab2-9984-2b0200895504	a525adb6-2358-4bae-920c-1f1249fbd3ff	0.00	0.00	2024-07-03 03:07:01.976968	2024-07-03 03:07:01.976968
a467ad40-93dc-45c4-be31-ae8bf115c881	0fefe4a3-0daf-4056-82c3-76bd49bb0bef	0.00	0.00	2024-07-03 03:07:29.955809	2024-07-03 03:07:29.955809
7a3b8b1b-0bce-4207-9019-91c4c80cc475	fe95e0cb-6070-4f93-8c69-4a2b61d3992e	0.00	0.00	2024-07-03 03:08:08.252573	2024-07-03 03:08:08.252573
106e244d-282c-40d8-ba5a-5c95db866d22	2f2173c8-05e1-4328-b6f8-c769927d43bd	0.00	0.00	2024-07-03 03:08:31.840083	2024-07-03 03:08:31.840083
fafe6f0b-afe9-4f2c-b730-0196a773bdb6	d0813034-19a1-406f-af25-c18d2f301614	43.00	83902.43	2024-06-28 00:20:03.404863	2024-07-08 12:29:14.20551
6feb27eb-d326-4234-ae0b-db810680fa57	e1bf0592-7850-4602-a740-6aae98dfd281	1208.00	1785.67	2024-06-06 09:29:38.590544	2024-07-08 15:55:01.756106
2930d942-cab3-4da2-be1c-005e586f9b9c	692e6c93-f159-40b0-bfb8-180ff6c7dfad	19.00	15500.00	2024-07-08 15:56:20.627214	2024-07-08 16:03:43.152401
812cfc24-b578-4ea4-9cc1-41b39d980a03	0b451368-e840-439d-9f07-824beb763eda	0.00	0.00	2024-07-08 16:06:01.23583	2024-07-08 16:06:01.23583
961fe0de-15ef-481b-b23d-1f3735d0e81c	412d3012-8a9f-44ea-a8a0-273aaa6a9cfc	0.00	0.00	2024-07-09 00:43:16.898766	2024-07-09 00:43:16.898766
43f1210d-5210-4cae-80ce-ad846341b38f	50b0499b-54e2-4406-8da7-ec66af69e932	119.00	12513.37	2024-07-10 02:00:18.145115	2024-07-10 02:55:58.499463
\.


--
-- Data for Name: product_units; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.product_units (id, to_unit_id, unit_id, product_id, scale, scale_to_base, created_at, updated_at) FROM stdin;
cec71c3c-5a20-461b-9370-051ab3eeba76	\N	6745548c-ea48-4db8-b7d9-ed2cf1175ade	e1bf0592-7850-4602-a740-6aae98dfd281	1.00	1.00	2024-06-06 09:29:38.584681	2024-06-06 09:29:38.584681
fa437b9d-89cd-47a6-8877-2a8e91336450	6745548c-ea48-4db8-b7d9-ed2cf1175ade	3867eb2b-8905-402c-bce3-c5953262ec03	e1bf0592-7850-4602-a740-6aae98dfd281	12.00	12.00	2024-06-06 09:29:38.584682	2024-06-06 09:29:38.584682
a28b11ad-6f08-4607-badd-d481c0e6be4c	\N	fbc61727-4eae-4f00-9ebf-2ff2c2f8156b	69a3f894-1ff2-4f61-97c2-3c957eea7914	1.00	1.00	2024-06-27 08:52:43.374949	2024-06-27 08:52:43.374949
d3e101c6-9603-4370-a4db-ea52b450f5b1	\N	3867eb2b-8905-402c-bce3-c5953262ec03	69a3f894-1ff2-4f61-97c2-3c957eea7914	20.00	20.00	2024-06-27 08:52:55.300837	2024-06-27 08:52:55.300837
61218f4d-b5ac-41db-8db6-1f3c3f18a10e	\N	6745548c-ea48-4db8-b7d9-ed2cf1175ade	d0813034-19a1-406f-af25-c18d2f301614	1.00	1.00	2024-06-28 00:20:34.451496	2024-06-28 00:20:34.451496
151c9e14-4403-4ed6-9973-f401a67841f2	\N	3867eb2b-8905-402c-bce3-c5953262ec03	d0813034-19a1-406f-af25-c18d2f301614	10.00	10.00	2024-06-28 00:20:42.113833	2024-06-28 00:20:42.113833
2c4c95d8-5e8f-410f-bbf4-02af780dc018	\N	fbc61727-4eae-4f00-9ebf-2ff2c2f8156b	2ea93d98-d8d5-4ce2-b13f-00ac30f863fc	1.00	1.00	2024-07-03 03:06:19.549567	2024-07-03 03:06:19.549567
f3e30180-a5b3-4af9-bae1-a73b497a9367	\N	fbc61727-4eae-4f00-9ebf-2ff2c2f8156b	32ca75e1-7d0f-4fc0-ad09-78d16d820499	1.00	1.00	2024-07-03 03:06:52.70815	2024-07-03 03:06:52.70815
8d880352-39e3-419a-ac91-dcce54978c28	\N	fbc61727-4eae-4f00-9ebf-2ff2c2f8156b	a525adb6-2358-4bae-920c-1f1249fbd3ff	1.00	1.00	2024-07-03 03:07:05.87568	2024-07-03 03:07:05.87568
eece9d6f-8740-42df-9097-12c9f8a29deb	\N	fbc61727-4eae-4f00-9ebf-2ff2c2f8156b	0fefe4a3-0daf-4056-82c3-76bd49bb0bef	1.00	1.00	2024-07-03 03:07:38.617175	2024-07-03 03:07:38.617175
3f60ce7e-5f89-4efe-bce7-d735ce8a9d78	\N	6745548c-ea48-4db8-b7d9-ed2cf1175ade	fe95e0cb-6070-4f93-8c69-4a2b61d3992e	1.00	1.00	2024-07-03 03:08:21.11169	2024-07-03 03:08:21.11169
ddca86b1-161c-4b25-b15f-736a450858c3	\N	6745548c-ea48-4db8-b7d9-ed2cf1175ade	2f2173c8-05e1-4328-b6f8-c769927d43bd	1.00	1.00	2024-07-03 03:08:39.318963	2024-07-03 03:08:39.318963
4ca122af-8969-44d6-ac51-fd3a1a8c504b	\N	fbc61727-4eae-4f00-9ebf-2ff2c2f8156b	692e6c93-f159-40b0-bfb8-180ff6c7dfad	1.00	1.00	2024-07-08 15:58:22.428898	2024-07-08 15:58:22.428898
cf0e07d0-e078-4992-b8ac-7108b78a4db6	fbc61727-4eae-4f00-9ebf-2ff2c2f8156b	3867eb2b-8905-402c-bce3-c5953262ec03	692e6c93-f159-40b0-bfb8-180ff6c7dfad	1.00	1.00	2024-07-08 16:03:10.672062	2024-07-08 16:03:10.672062
2ea5c37e-d581-4653-bb74-c7c355ce4498	3867eb2b-8905-402c-bce3-c5953262ec03	6745548c-ea48-4db8-b7d9-ed2cf1175ade	692e6c93-f159-40b0-bfb8-180ff6c7dfad	1.00	1.00	2024-07-08 16:06:05.530996	2024-07-08 16:06:05.530996
7fbc5a3b-6f6c-497e-b265-cd12a65314d5	\N	6745548c-ea48-4db8-b7d9-ed2cf1175ade	412d3012-8a9f-44ea-a8a0-273aaa6a9cfc	1.00	1.00	2024-07-09 00:43:27.22165	2024-07-09 00:43:27.22165
f8fd2e10-1a96-44ae-9898-50fa57d8dcf9	\N	fbc61727-4eae-4f00-9ebf-2ff2c2f8156b	50b0499b-54e2-4406-8da7-ec66af69e932	1.00	1.00	2024-07-10 02:00:37.068088	2024-07-10 02:00:37.068088
\.


--
-- Data for Name: products; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.products (id, image_file_id, name, description, price, is_active, created_at, updated_at) FROM stdin;
96773beb-3646-4078-893f-19a278e98a30	cc775b17-0e56-42fb-87b9-ddc21591e734	Spidol Snowman Black Color Permanent	\N	\N	f	2024-06-07 07:58:48.424132	2024-06-07 07:58:48.424132
412d3012-8a9f-44ea-a8a0-273aaa6a9cfc	6970f822-6ba5-4cfd-8cdd-ca5e9b5502fa	Cat Kaleng Biru	\N	50000.00	t	2024-07-09 00:43:16.896322	2024-07-09 00:43:29.871198
50b0499b-54e2-4406-8da7-ec66af69e932	f2532f25-64eb-4cc0-badf-c64992882ffa	Produk HPP	\N	12000.00	t	2024-07-10 02:00:18.143535	2024-07-10 02:17:46.892428
e1bf0592-7850-4602-a740-6aae98dfd281	35c972f5-2537-490d-90ca-36ba659d6c20	Kaleng Cat Merah	Test	150000.00	t	2024-06-06 09:29:38.577335	2024-06-21 09:28:13.050501
69a3f894-1ff2-4f61-97c2-3c957eea7914	86676299-895f-4dc4-9399-fc30cbb487b7	Produk Pcs	\N	100000.00	t	2024-06-27 08:52:04.697143	2024-06-27 08:53:02.828006
d0813034-19a1-406f-af25-c18d2f301614	ea593a9a-3f06-46f9-8aa5-c06d431f4f12	Produk Cat Kaleng	\N	100000.00	t	2024-06-28 00:20:03.397408	2024-06-28 00:20:53.365229
bcfd2b82-8125-4e06-b6d8-2e42c70fdd74	66e0519f-c0ec-4ca2-85da-b1b1c2283012	Kuas	Kuas 4"	\N	f	2024-06-29 15:38:14.340844	2024-06-29 15:38:14.340844
2ea93d98-d8d5-4ce2-b13f-00ac30f863fc	f735c1b5-cae8-4a02-83bb-24a3cc62823c	egweg	rehsreh	10000.00	t	2024-07-03 03:06:00.059829	2024-07-03 03:06:26.264277
32ca75e1-7d0f-4fc0-ad09-78d16d820499	3b760ad8-1f2e-4cf9-ad25-0eddf9966698	awegewag	\N	1111.00	t	2024-07-03 03:06:43.255696	2024-07-03 03:06:54.814043
a525adb6-2358-4bae-920c-1f1249fbd3ff	b15546e2-d5f3-4603-8014-1e37c2902742	uykgyuk	\N	12415.00	t	2024-07-03 03:07:01.974509	2024-07-03 03:07:13.938928
0fefe4a3-0daf-4056-82c3-76bd49bb0bef	3eb64ea4-cacc-4700-a3c9-934b4c2ea402	aerheah	\N	236236.00	t	2024-07-03 03:07:29.954756	2024-07-03 03:07:45.484498
fe95e0cb-6070-4f93-8c69-4a2b61d3992e	13781a6c-1655-4c37-a905-287132abd378	rtj	\N	35.00	t	2024-07-03 03:08:08.251222	2024-07-03 03:08:22.612179
692e6c93-f159-40b0-bfb8-180ff6c7dfad	b39a3ddc-6b2f-46d5-8ecf-c01abb5e29de	Lampu	Hannoch	10000.00	t	2024-07-08 15:56:20.625316	2024-07-08 16:04:20.187107
0b451368-e840-439d-9f07-824beb763eda	ab461851-6d15-4442-85aa-cc13758df092	Balon	Balon	10000.00	f	2024-07-08 16:06:01.234851	2024-07-08 16:06:11.517328
2f2173c8-05e1-4328-b6f8-c769927d43bd	c90220c3-440d-4224-9da2-e5112b52f0d9	rsehs	\N	12000.00	t	2024-07-03 03:08:31.838435	2024-07-08 16:06:40.545485
\.


--
-- Data for Name: purchase_order_images; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.purchase_order_images (id, purchase_order_id, file_id, description, created_at, updated_at) FROM stdin;
4a775dce-384b-48b3-98ab-9b94ef55c802	b45d509c-7440-43ab-97d5-4296321ebe22	db1850d1-a9c3-49d7-a3be-46fb88eb6b09	safaf	2024-06-07 02:42:10.741451	2024-06-07 02:42:10.741451
c1c5b0a5-0b7b-42b2-8161-72cc28ffc698	29d46083-0872-4a3a-96a3-510ad7613a10	5d0b0fc8-ed9b-4f76-a956-37a5d9e23c74	\N	2024-06-27 08:57:26.171586	2024-06-27 08:57:26.171586
d9399f67-1748-4e00-a050-872a264e2e36	6c227a2d-c6fd-48b4-8aa8-c4045a8981fb	46da0322-efbe-41ec-bf95-b33330cd0e6b	\N	2024-06-27 09:19:34.065138	2024-06-27 09:19:34.065138
bf1d1d72-9a28-45f1-b751-d391e7bf53c5	86f2a2b3-eda0-458f-a52c-455a81d91818	d689923d-2b04-4cab-afb3-87cfb46fdc71	\N	2024-06-28 00:25:25.313564	2024-06-28 00:25:25.313564
a0ac4c2f-2591-4ea4-ab36-4ba2da36bb0a	66989bdd-f2e8-404b-87ed-d283180a325e	552b2ddd-1403-4911-85da-ad97043742e1	Gambar	2024-07-08 15:17:08.587738	2024-07-08 15:17:08.587738
\.


--
-- Data for Name: purchase_order_items; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.purchase_order_items (id, purchase_order_id, product_unit_id, user_id, qty, scale_to_base, price_per_unit, created_at, updated_at) FROM stdin;
b801836f-c152-4eda-b6f1-0e891987b978	b45d509c-7440-43ab-97d5-4296321ebe22	cec71c3c-5a20-461b-9370-051ab3eeba76	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	5.00	1.00	9000.00	2024-06-07 02:41:59.125508	2024-06-07 02:41:59.125508
2b280938-d7e6-4665-918f-c4f42614abc8	407abd68-9449-4f1a-aed6-1fd38fbd90f3	cec71c3c-5a20-461b-9370-051ab3eeba76	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	12.00	1.00	10000.00	2024-06-07 15:45:04.80518	2024-06-07 15:45:04.80518
af1103fe-14ea-42f6-ab7a-1023c740a39f	2b6b930a-44c4-4ea7-9f4b-82cb7b6296c0	cec71c3c-5a20-461b-9370-051ab3eeba76	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	10.00	1.00	10000.00	2024-06-12 04:15:21.410998	2024-06-12 04:15:21.410998
68ae0f37-8371-45be-b3f5-2bcebc93a6a6	b41e582b-aff6-400c-8d55-b7342c1a5808	fa437b9d-89cd-47a6-8877-2a8e91336450	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	100.00	12.00	10000.00	2024-06-20 02:21:58.593223	2024-06-20 02:21:58.593223
79140d94-6e28-4cc0-9b73-1f46e000a68c	29d46083-0872-4a3a-96a3-510ad7613a10	a28b11ad-6f08-4607-badd-d481c0e6be4c	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	20.00	1.00	8000.00	2024-06-27 08:56:54.385645	2024-06-27 08:56:54.385645
52d78224-08de-45f5-95be-c207e035413b	29d46083-0872-4a3a-96a3-510ad7613a10	d3e101c6-9603-4370-a4db-ea52b450f5b1	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	10.00	20.00	70000.00	2024-06-27 08:57:15.503451	2024-06-27 08:57:15.503451
92e00e16-d758-4870-9d47-ed873350a884	6c227a2d-c6fd-48b4-8aa8-c4045a8981fb	a28b11ad-6f08-4607-badd-d481c0e6be4c	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	10.00	1.00	80000.00	2024-06-27 09:19:00.646833	2024-06-27 09:19:00.646833
92dc60b2-2898-4044-aedc-cde6136575e0	6c227a2d-c6fd-48b4-8aa8-c4045a8981fb	d3e101c6-9603-4370-a4db-ea52b450f5b1	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	10.00	20.00	700000.00	2024-06-27 09:19:18.669465	2024-06-27 09:19:18.669465
b8d03984-e65d-453d-92ac-532017db4a40	86f2a2b3-eda0-458f-a52c-455a81d91818	61218f4d-b5ac-41db-8db6-1f3c3f18a10e	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	10.00	1.00	80000.00	2024-06-28 00:24:59.670638	2024-06-28 00:24:59.670638
de6028e2-0c49-41fa-a9c0-08ed5316727a	86f2a2b3-eda0-458f-a52c-455a81d91818	151c9e14-4403-4ed6-9973-f401a67841f2	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	10.00	10.00	700000.00	2024-06-28 00:25:12.953401	2024-06-28 00:25:12.953401
1d342b59-9b7d-480d-9e97-2ab823d7ebb8	037a8901-ad4a-40e9-9b34-088e78fcad06	151c9e14-4403-4ed6-9973-f401a67841f2	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	5.00	10.00	1000000.00	2024-07-01 11:10:24.067766	2024-07-01 11:10:24.067766
7f03893b-45e0-4762-8cc8-cf1592cd4d59	66989bdd-f2e8-404b-87ed-d283180a325e	61218f4d-b5ac-41db-8db6-1f3c3f18a10e	33392eb4-0f87-43c6-9893-9c014fe6d561	100.00	1.00	10000.00	2024-07-08 15:16:48.634685	2024-07-08 15:16:48.634685
38e71829-9070-45fe-b364-bd772fa662be	d394c940-2349-4179-8c3a-c2755bb84419	f8fd2e10-1a96-44ae-9898-50fa57d8dcf9	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	100.00	1.00	10000.00	2024-07-10 02:04:15.947069	2024-07-10 02:04:15.947069
27d02137-7b2a-43a4-b575-9b1f3e0e23e5	d394c940-2349-4179-8c3a-c2755bb84419	f8fd2e10-1a96-44ae-9898-50fa57d8dcf9	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	10.00	1.00	20000.00	2024-07-10 02:04:30.612365	2024-07-10 02:04:30.612365
934125c0-f352-4a11-82cc-5dbf11bce799	66ad8587-155b-4f3c-897c-5cdd93cb3799	f8fd2e10-1a96-44ae-9898-50fa57d8dcf9	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	100.00	1.00	10000.00	2024-07-10 02:18:27.072743	2024-07-10 02:18:27.072743
91418c2f-fe45-450d-9e38-a6d96bc0ce15	66ad8587-155b-4f3c-897c-5cdd93cb3799	f8fd2e10-1a96-44ae-9898-50fa57d8dcf9	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	10.00	1.00	20000.00	2024-07-10 02:18:38.957442	2024-07-10 02:18:38.957442
f9c6a727-a621-4b43-b889-346c4e2fbf74	f67ca56b-2eba-4835-b50a-9a53b88a021c	f8fd2e10-1a96-44ae-9898-50fa57d8dcf9	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	10.00	1.00	30000.00	2024-07-10 02:55:48.949622	2024-07-10 02:55:48.949622
\.


--
-- Data for Name: purchase_orders; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.purchase_orders (id, supplier_id, user_id, invoice_number, date, status, total_estimated_price, created_at, updated_at) FROM stdin;
b45d509c-7440-43ab-97d5-4296321ebe22	e0225ca5-4094-47e9-8494-43795dfddd97	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	wsadeaf	2024-06-07	COMPLETED	45000.00	2024-06-07 02:40:22.156802	2024-06-07 02:54:03.550034
66989bdd-f2e8-404b-87ed-d283180a325e	48fba4fe-787a-45fb-8c65-3ff6418b09a6	33392eb4-0f87-43c6-9893-9c014fe6d561	INV/002/2024	2024-07-08	ONGOING	1000000.00	2024-07-08 15:15:43.667767	2024-07-08 15:17:16.864682
e37af22f-1ee8-40db-aa33-735ba2332db9	48fba4fe-787a-45fb-8c65-3ff6418b09a6	33392eb4-0f87-43c6-9893-9c014fe6d561	INV/001/2024	2024-07-09	PENDING	0.00	2024-07-08 17:23:13.503303	2024-07-08 17:23:13.503303
407abd68-9449-4f1a-aed6-1fd38fbd90f3	e0225ca5-4094-47e9-8494-43795dfddd97	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	INV/001/2024	2024-06-07	COMPLETED	120000.00	2024-06-07 15:31:05.341628	2024-06-07 15:45:18.208695
2b6b930a-44c4-4ea7-9f4b-82cb7b6296c0	e0225ca5-4094-47e9-8494-43795dfddd97	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	INV/002/2024	2024-06-07	PENDING	100000.00	2024-06-07 16:30:58.219677	2024-06-12 04:15:21.409734
b41e582b-aff6-400c-8d55-b7342c1a5808	e0225ca5-4094-47e9-8494-43795dfddd97	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	rehrehsre	2024-06-20	COMPLETED	1000000.00	2024-06-20 02:21:49.247303	2024-06-20 02:22:05.239544
d394c940-2349-4179-8c3a-c2755bb84419	e0225ca5-4094-47e9-8494-43795dfddd97	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	INv-0111	2024-07-10	COMPLETED	1200000.00	2024-07-10 02:04:03.241697	2024-07-10 02:04:39.122676
29d46083-0872-4a3a-96a3-510ad7613a10	e6815351-8cd2-4471-a81e-ac21eb65e7ea	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	C001-98252/512512/w521	2024-06-27	COMPLETED	860000.00	2024-06-27 08:56:36.814622	2024-06-27 08:58:21.518932
6c227a2d-c6fd-48b4-8aa8-c4045a8981fb	e6815351-8cd2-4471-a81e-ac21eb65e7ea	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	C001-598102521/15215125/215	2024-06-27	COMPLETED	7800000.00	2024-06-27 09:18:41.041405	2024-06-27 09:20:17.801064
66ad8587-155b-4f3c-897c-5cdd93cb3799	e0225ca5-4094-47e9-8494-43795dfddd97	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	INV-001	2024-07-10	COMPLETED	1200000.00	2024-07-10 02:18:15.799903	2024-07-10 02:18:44.819171
86f2a2b3-eda0-458f-a52c-455a81d91818	48fba4fe-787a-45fb-8c65-3ff6418b09a6	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	INVC-C001-24215/125215	2024-06-28	COMPLETED	7800000.00	2024-06-28 00:24:43.528936	2024-06-28 00:25:58.183905
f67ca56b-2eba-4835-b50a-9a53b88a021c	48fba4fe-787a-45fb-8c65-3ff6418b09a6	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	INv-002	2024-07-10	COMPLETED	300000.00	2024-07-10 02:55:38.403724	2024-07-10 02:55:52.930879
037a8901-ad4a-40e9-9b34-088e78fcad06	e0225ca5-4094-47e9-8494-43795dfddd97	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	a	2024-06-18	COMPLETED	5000000.00	2024-06-18 06:32:23.311601	2024-07-01 11:10:59.48741
\.


--
-- Data for Name: role_permissions; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.role_permissions (role_id, permission_id, created_at, updated_at) FROM stdin;
73b6799d-ebd6-491f-9b47-272fb0f22914	d67e64c2-b814-48f8-beb7-74d60afe579f	2024-07-10 02:53:07.911497	2024-07-10 02:53:07.911497
73b6799d-ebd6-491f-9b47-272fb0f22914	877cf804-d173-4f9a-8818-3eeba0fe0183	2024-07-10 02:53:07.911498	2024-07-10 02:53:07.911498
73b6799d-ebd6-491f-9b47-272fb0f22914	91b0dd33-0bbe-4641-844d-ec6294163f4d	2024-07-10 02:53:07.911499	2024-07-10 02:53:07.911499
73b6799d-ebd6-491f-9b47-272fb0f22914	ec412394-c888-4deb-a87d-7138ae9b12e0	2024-07-10 02:53:07.9115	2024-07-10 02:53:07.9115
73b6799d-ebd6-491f-9b47-272fb0f22914	c7f918e0-8b20-4d2f-80a6-4f861e12faa4	2024-07-10 02:53:07.911501	2024-07-10 02:53:07.911501
73b6799d-ebd6-491f-9b47-272fb0f22914	8498428c-ab74-4f89-858b-045240445f2e	2024-07-10 02:53:07.911502	2024-07-10 02:53:07.911502
73b6799d-ebd6-491f-9b47-272fb0f22914	33cdc1aa-b351-4fe8-bd22-87752c178712	2024-07-10 02:53:07.911503	2024-07-10 02:53:07.911503
73b6799d-ebd6-491f-9b47-272fb0f22914	0bbd03cb-4610-4a15-8edd-7cee9e6f8a36	2024-07-10 02:53:07.911504	2024-07-10 02:53:07.911504
73b6799d-ebd6-491f-9b47-272fb0f22914	641a1c56-5d11-4b45-a678-b6c63e89c5ed	2024-07-10 02:53:07.911505	2024-07-10 02:53:07.911505
73b6799d-ebd6-491f-9b47-272fb0f22914	f2c1794b-e513-47f1-a009-0b3329080565	2024-07-10 02:53:07.911506	2024-07-10 02:53:07.911506
73b6799d-ebd6-491f-9b47-272fb0f22914	6a5d87b1-3a89-48c4-b9b7-e9dc6bed11b8	2024-07-10 02:53:07.911507	2024-07-10 02:53:07.911507
73b6799d-ebd6-491f-9b47-272fb0f22914	c6a9c058-19fb-4eff-b050-517750e87888	2024-07-10 02:53:07.911508	2024-07-10 02:53:07.911508
73b6799d-ebd6-491f-9b47-272fb0f22914	a0c74ca7-8945-4226-bac4-43ffe2f841c6	2024-07-10 02:53:07.911509	2024-07-10 02:53:07.911509
73b6799d-ebd6-491f-9b47-272fb0f22914	163c477a-e44d-4c4f-96c7-7d24a686e85e	2024-07-10 02:53:07.91151	2024-07-10 02:53:07.91151
73b6799d-ebd6-491f-9b47-272fb0f22914	8d31041b-440a-4e57-8fbe-5a6e75c0f87f	2024-07-10 02:53:07.911511	2024-07-10 02:53:07.911511
73b6799d-ebd6-491f-9b47-272fb0f22914	66d678b6-b95e-4c19-8d16-226d9dc86e9f	2024-07-10 02:53:07.911512	2024-07-10 02:53:07.911512
73b6799d-ebd6-491f-9b47-272fb0f22914	2c0be1d0-8f75-4df1-95eb-b084a268ba65	2024-07-10 02:53:07.911513	2024-07-10 02:53:07.911513
73b6799d-ebd6-491f-9b47-272fb0f22914	9e32f57b-fccb-4041-bb56-6a7a44b8f625	2024-07-10 02:53:07.911514	2024-07-10 02:53:07.911514
73b6799d-ebd6-491f-9b47-272fb0f22914	3a2cec95-710e-424a-84ac-5f36a8a6fe1c	2024-07-10 02:53:07.911515	2024-07-10 02:53:07.911515
73b6799d-ebd6-491f-9b47-272fb0f22914	8a62b161-d14c-4cca-b26a-1cb8c6e03d7a	2024-07-10 02:53:07.911516	2024-07-10 02:53:07.911516
73b6799d-ebd6-491f-9b47-272fb0f22914	539a1f07-7bb0-4a0d-b207-647f7be51b8a	2024-07-10 02:53:07.911517	2024-07-10 02:53:07.911517
73b6799d-ebd6-491f-9b47-272fb0f22914	7bcc4d1a-e10f-48a9-bbf4-687f504520a9	2024-07-10 02:53:07.911518	2024-07-10 02:53:07.911518
73b6799d-ebd6-491f-9b47-272fb0f22914	99bd3a4d-b5af-4d80-900d-f2a3a590a730	2024-07-10 02:53:07.911519	2024-07-10 02:53:07.911519
73b6799d-ebd6-491f-9b47-272fb0f22914	ab3217e4-eef1-4e88-a651-2aebba7421bd	2024-07-10 02:53:07.91152	2024-07-10 02:53:07.91152
73b6799d-ebd6-491f-9b47-272fb0f22914	ffe81eb0-560c-44cb-b53c-5e7733ac652d	2024-07-10 02:53:07.911521	2024-07-10 02:53:07.911521
73b6799d-ebd6-491f-9b47-272fb0f22914	b1ed1d28-a4c7-4539-986d-628d501b5959	2024-07-10 02:53:07.911522	2024-07-10 02:53:07.911522
73b6799d-ebd6-491f-9b47-272fb0f22914	8fa77f3c-9d9f-4f2e-a7f0-8c32571448ec	2024-07-10 02:53:07.911523	2024-07-10 02:53:07.911523
73b6799d-ebd6-491f-9b47-272fb0f22914	80295769-1e98-4bd6-9c91-29693ad689f1	2024-07-10 02:53:07.911524	2024-07-10 02:53:07.911524
73b6799d-ebd6-491f-9b47-272fb0f22914	013195bc-285e-4a02-bcfc-6e1750442584	2024-07-10 02:53:07.911525	2024-07-10 02:53:07.911525
73b6799d-ebd6-491f-9b47-272fb0f22914	f8d4581e-ce94-4d49-a2d5-a73348cf1da6	2024-07-10 02:53:07.911526	2024-07-10 02:53:07.911526
73b6799d-ebd6-491f-9b47-272fb0f22914	7abd445c-e14e-4153-8580-0f7bfaf615ea	2024-07-10 02:53:07.911527	2024-07-10 02:53:07.911527
73b6799d-ebd6-491f-9b47-272fb0f22914	ec53f87b-f420-428d-8f04-e55234ae146d	2024-07-10 02:53:07.911528	2024-07-10 02:53:07.911528
73b6799d-ebd6-491f-9b47-272fb0f22914	e0515501-9069-4ff0-bdb6-ca0cbe734c76	2024-07-10 02:53:07.911529	2024-07-10 02:53:07.911529
73b6799d-ebd6-491f-9b47-272fb0f22914	26529ea7-d633-46e7-a080-fefe8924cf0a	2024-07-10 02:53:07.91153	2024-07-10 02:53:07.91153
73b6799d-ebd6-491f-9b47-272fb0f22914	4e4b3bae-4446-4785-86bb-5148772f2589	2024-07-10 02:53:07.911531	2024-07-10 02:53:07.911531
73b6799d-ebd6-491f-9b47-272fb0f22914	a9195ab7-6e2f-4716-b56d-c5888599ca92	2024-07-10 02:53:07.911532	2024-07-10 02:53:07.911532
73b6799d-ebd6-491f-9b47-272fb0f22914	f888ed1c-8daf-46ff-b80a-7b95d5503145	2024-07-10 02:53:07.911533	2024-07-10 02:53:07.911533
73b6799d-ebd6-491f-9b47-272fb0f22914	171bf39a-c8f0-4bd4-8bab-9bb113660ba3	2024-07-10 02:53:07.911534	2024-07-10 02:53:07.911534
73b6799d-ebd6-491f-9b47-272fb0f22914	6e3ed14e-426f-46f8-a9c4-b8be122d2c89	2024-07-10 02:53:07.911535	2024-07-10 02:53:07.911535
73b6799d-ebd6-491f-9b47-272fb0f22914	601d894f-8ec0-42bf-9a30-e2df4ad8aa23	2024-07-10 02:53:07.911536	2024-07-10 02:53:07.911536
73b6799d-ebd6-491f-9b47-272fb0f22914	fcc24c96-608c-47f7-9b55-128231841c3f	2024-07-10 02:53:07.911537	2024-07-10 02:53:07.911537
73b6799d-ebd6-491f-9b47-272fb0f22914	45ce843c-9dee-49f8-b358-581a61e4e641	2024-07-10 02:53:07.911538	2024-07-10 02:53:07.911538
73b6799d-ebd6-491f-9b47-272fb0f22914	71d7b694-8f90-4d01-adf8-bc952e1679bb	2024-07-10 02:53:07.911539	2024-07-10 02:53:07.911539
73b6799d-ebd6-491f-9b47-272fb0f22914	efccb3c8-2351-441c-abf8-ad819243fa3c	2024-07-10 02:53:07.91154	2024-07-10 02:53:07.91154
73b6799d-ebd6-491f-9b47-272fb0f22914	691f7440-1d77-4c66-9fc6-3abfd0593cb8	2024-07-10 02:53:07.911541	2024-07-10 02:53:07.911541
73b6799d-ebd6-491f-9b47-272fb0f22914	4b1b4078-b751-427d-8f86-11df92d7a1b3	2024-07-10 02:53:07.911542	2024-07-10 02:53:07.911542
73b6799d-ebd6-491f-9b47-272fb0f22914	da4f9a7b-4153-4735-a4a9-da943ea52cbb	2024-07-10 02:53:07.911543	2024-07-10 02:53:07.911543
73b6799d-ebd6-491f-9b47-272fb0f22914	c8d83756-e082-4332-82da-5d7efad64abc	2024-07-10 02:53:07.911544	2024-07-10 02:53:07.911544
73b6799d-ebd6-491f-9b47-272fb0f22914	c0451ce1-5573-4964-b96a-869a813f4bc2	2024-07-10 02:53:07.911545	2024-07-10 02:53:07.911545
73b6799d-ebd6-491f-9b47-272fb0f22914	1e2eb6ca-e32b-420a-bc07-53a3d6c6227a	2024-07-10 02:53:07.911546	2024-07-10 02:53:07.911546
73b6799d-ebd6-491f-9b47-272fb0f22914	9d933c71-4073-44e9-a3ce-aeca83eca948	2024-07-10 02:53:07.911547	2024-07-10 02:53:07.911547
73b6799d-ebd6-491f-9b47-272fb0f22914	fca47e9f-9a7b-4eef-a312-3c525cc7871c	2024-07-10 02:53:07.911548	2024-07-10 02:53:07.911548
73b6799d-ebd6-491f-9b47-272fb0f22914	7879169e-1c71-40e5-b077-a647f7581d80	2024-07-10 02:53:07.911549	2024-07-10 02:53:07.911549
73b6799d-ebd6-491f-9b47-272fb0f22914	1d84ac87-15a6-4454-936a-30692a5b7941	2024-07-10 02:53:07.91155	2024-07-10 02:53:07.91155
73b6799d-ebd6-491f-9b47-272fb0f22914	df196d96-a49d-48df-b1aa-96e9e8d70df0	2024-07-10 02:53:07.911551	2024-07-10 02:53:07.911551
73b6799d-ebd6-491f-9b47-272fb0f22914	a4554516-a1d9-4d6c-b68e-dad15163a569	2024-07-10 02:53:07.911552	2024-07-10 02:53:07.911552
73b6799d-ebd6-491f-9b47-272fb0f22914	7766b2e2-b835-43ba-acbc-434c24fd6e5c	2024-07-10 02:53:07.911553	2024-07-10 02:53:07.911553
73b6799d-ebd6-491f-9b47-272fb0f22914	440fff70-d0be-44c2-9a0f-3fe1b561631a	2024-07-10 02:53:07.911554	2024-07-10 02:53:07.911554
73b6799d-ebd6-491f-9b47-272fb0f22914	c81e91b1-45e2-446a-9395-59dbf00cfb7f	2024-07-10 02:53:07.911555	2024-07-10 02:53:07.911555
73b6799d-ebd6-491f-9b47-272fb0f22914	3b74461a-1789-4601-9416-881c98cf9587	2024-07-10 02:53:07.911556	2024-07-10 02:53:07.911556
73b6799d-ebd6-491f-9b47-272fb0f22914	1230d117-1892-4c15-9179-a517d9d5351d	2024-07-10 02:53:07.911557	2024-07-10 02:53:07.911557
73b6799d-ebd6-491f-9b47-272fb0f22914	7e696943-0f22-46f1-90d7-b016b6e5f81d	2024-07-10 02:53:07.911558	2024-07-10 02:53:07.911558
73b6799d-ebd6-491f-9b47-272fb0f22914	1a912ef6-8a84-4e17-9c86-e5355f2b9c91	2024-07-10 02:53:07.911559	2024-07-10 02:53:07.911559
73b6799d-ebd6-491f-9b47-272fb0f22914	038712a3-a8d6-4626-a4e5-e42507713e3e	2024-07-10 02:53:07.91156	2024-07-10 02:53:07.91156
73b6799d-ebd6-491f-9b47-272fb0f22914	358ce19f-812a-445d-a686-fe10b339d94f	2024-07-10 02:53:07.911561	2024-07-10 02:53:07.911561
73b6799d-ebd6-491f-9b47-272fb0f22914	8a88808b-3c50-44bf-885c-7697ca73b749	2024-07-10 02:53:07.911562	2024-07-10 02:53:07.911562
73b6799d-ebd6-491f-9b47-272fb0f22914	57ecf435-b67e-4e3a-b597-a5b9b1bca788	2024-07-10 02:53:07.911563	2024-07-10 02:53:07.911563
73b6799d-ebd6-491f-9b47-272fb0f22914	7a3da445-7f55-4cde-9f83-c47bed4719b9	2024-07-10 02:53:07.911564	2024-07-10 02:53:07.911564
73b6799d-ebd6-491f-9b47-272fb0f22914	52be45c3-f444-4a7f-bfa4-43b4bd94ad18	2024-07-10 02:53:07.911565	2024-07-10 02:53:07.911565
73b6799d-ebd6-491f-9b47-272fb0f22914	3588e557-84ac-4ca5-87ba-96086f41528f	2024-07-10 02:53:07.911566	2024-07-10 02:53:07.911566
73b6799d-ebd6-491f-9b47-272fb0f22914	5ecfa2fa-da00-4f0f-a4a5-4a75ef16deba	2024-07-10 02:53:07.911567	2024-07-10 02:53:07.911567
73b6799d-ebd6-491f-9b47-272fb0f22914	fe5b6be5-599f-414b-8a35-117e94cbf06f	2024-07-10 02:53:07.911568	2024-07-10 02:53:07.911568
73b6799d-ebd6-491f-9b47-272fb0f22914	7379cd4a-2c0b-4f7f-bbf2-f2f740746b11	2024-07-10 02:53:07.911569	2024-07-10 02:53:07.911569
73b6799d-ebd6-491f-9b47-272fb0f22914	fb23c89d-55be-424b-a545-7b5978b73da4	2024-07-10 02:53:07.91157	2024-07-10 02:53:07.91157
73b6799d-ebd6-491f-9b47-272fb0f22914	f2067064-7f15-409c-8348-dd9a7d7f9ea5	2024-07-10 02:53:07.911571	2024-07-10 02:53:07.911571
73b6799d-ebd6-491f-9b47-272fb0f22914	319b03c9-593c-4d93-818c-4842227b5390	2024-07-10 02:53:07.911572	2024-07-10 02:53:07.911572
73b6799d-ebd6-491f-9b47-272fb0f22914	db0e803f-3dd0-421a-ba39-75bdffa8467a	2024-07-10 02:53:07.911573	2024-07-10 02:53:07.911573
73b6799d-ebd6-491f-9b47-272fb0f22914	ebc849ac-f205-4ee0-be30-aaffe0fd1197	2024-07-10 02:53:07.911574	2024-07-10 02:53:07.911574
73b6799d-ebd6-491f-9b47-272fb0f22914	eff24eed-1013-469d-92c0-2be65b9283bb	2024-07-10 02:53:07.911575	2024-07-10 02:53:07.911575
73b6799d-ebd6-491f-9b47-272fb0f22914	0bd6b4f2-1831-415f-a9f5-a429dd13fa47	2024-07-10 02:53:07.911576	2024-07-10 02:53:07.911576
73b6799d-ebd6-491f-9b47-272fb0f22914	cc39ef50-77d5-4b16-b6bb-b6d946b5cbc7	2024-07-10 02:53:07.911577	2024-07-10 02:53:07.911577
73b6799d-ebd6-491f-9b47-272fb0f22914	6c46505a-258f-41be-9ee5-1d6d822ee411	2024-07-10 02:53:07.911578	2024-07-10 02:53:07.911578
73b6799d-ebd6-491f-9b47-272fb0f22914	dac0c5ef-acd5-4e8a-99a4-eec374e359a8	2024-07-10 02:53:07.911579	2024-07-10 02:53:07.911579
73b6799d-ebd6-491f-9b47-272fb0f22914	63768160-5ecd-4965-9680-b6ccf56c7e3b	2024-07-10 02:53:07.91158	2024-07-10 02:53:07.91158
73b6799d-ebd6-491f-9b47-272fb0f22914	08962a85-84b7-4b50-87ea-d57e87d924bb	2024-07-10 02:53:07.911581	2024-07-10 02:53:07.911581
73b6799d-ebd6-491f-9b47-272fb0f22914	45b30169-67ab-4e98-94dd-30b4f6e9d3c9	2024-07-10 02:53:07.911582	2024-07-10 02:53:07.911582
73b6799d-ebd6-491f-9b47-272fb0f22914	303fa09c-a035-4c8c-ae0b-feee49bce02d	2024-07-10 02:53:07.911583	2024-07-10 02:53:07.911583
73b6799d-ebd6-491f-9b47-272fb0f22914	0e6b4487-f1f9-4a4c-910a-197f44c4e0ef	2024-07-10 02:53:07.911584	2024-07-10 02:53:07.911584
73b6799d-ebd6-491f-9b47-272fb0f22914	e2672b71-d57c-4b54-8fc9-d96f329a2226	2024-07-10 02:53:07.911585	2024-07-10 02:53:07.911585
73b6799d-ebd6-491f-9b47-272fb0f22914	5559bc7d-cd37-49ce-8125-fff62c614418	2024-07-10 02:53:07.911586	2024-07-10 02:53:07.911586
73b6799d-ebd6-491f-9b47-272fb0f22914	41f824c1-66d4-43c0-8980-b9a6c439db61	2024-07-10 02:53:07.911587	2024-07-10 02:53:07.911587
73b6799d-ebd6-491f-9b47-272fb0f22914	19d906a6-3ce7-40c5-b882-322a26dfeb1a	2024-07-10 02:53:07.911588	2024-07-10 02:53:07.911588
73b6799d-ebd6-491f-9b47-272fb0f22914	65b4ea55-b0bc-4501-b887-3fb445e1e270	2024-07-10 02:53:07.911589	2024-07-10 02:53:07.911589
73b6799d-ebd6-491f-9b47-272fb0f22914	dcf28c10-c5c3-4c8c-926f-59e6d01ab72e	2024-07-10 02:53:07.91159	2024-07-10 02:53:07.91159
73b6799d-ebd6-491f-9b47-272fb0f22914	bd95fa84-462a-4cb6-adb1-7b9e08a70c1d	2024-07-10 02:53:07.911591	2024-07-10 02:53:07.911591
73b6799d-ebd6-491f-9b47-272fb0f22914	1ed42cc6-072f-459b-b50c-650b1bb36ed5	2024-07-10 02:53:07.911592	2024-07-10 02:53:07.911592
73b6799d-ebd6-491f-9b47-272fb0f22914	bf0ec2fe-0008-4a75-a796-ef89c6e2da31	2024-07-10 02:53:07.911593	2024-07-10 02:53:07.911593
73b6799d-ebd6-491f-9b47-272fb0f22914	ae65c638-a388-4d58-987a-3a05235a6da6	2024-07-10 02:53:07.911594	2024-07-10 02:53:07.911594
73b6799d-ebd6-491f-9b47-272fb0f22914	5cb4530b-317e-429a-be88-5a52e4030aa0	2024-07-10 02:53:07.911595	2024-07-10 02:53:07.911595
73b6799d-ebd6-491f-9b47-272fb0f22914	0b8b3b00-2e1e-45ba-bc1a-c2cfea663b8a	2024-07-10 02:53:07.911596	2024-07-10 02:53:07.911596
73b6799d-ebd6-491f-9b47-272fb0f22914	f1eed559-4bd1-46ec-8125-050bbe8172b6	2024-07-10 02:53:07.911597	2024-07-10 02:53:07.911597
73b6799d-ebd6-491f-9b47-272fb0f22914	567b771d-8e06-47bb-91e3-d1b3ec48d82d	2024-07-10 02:53:07.911598	2024-07-10 02:53:07.911598
73b6799d-ebd6-491f-9b47-272fb0f22914	3e351283-36ed-4e20-a7fc-4b266eefcb13	2024-07-10 02:53:07.911599	2024-07-10 02:53:07.911599
73b6799d-ebd6-491f-9b47-272fb0f22914	01d08933-ebf6-4538-83d5-ada5ced8b35b	2024-07-10 02:53:07.9116	2024-07-10 02:53:07.9116
73b6799d-ebd6-491f-9b47-272fb0f22914	8a7cdabc-6086-4f3d-a707-299a2ef8aaa0	2024-07-10 02:53:07.911601	2024-07-10 02:53:07.911601
73b6799d-ebd6-491f-9b47-272fb0f22914	2f9f5814-3da4-4635-b674-8e26544a7ebe	2024-07-10 02:53:07.911602	2024-07-10 02:53:07.911602
73b6799d-ebd6-491f-9b47-272fb0f22914	db8a62b3-d44f-4c30-b24c-ea28e0d15324	2024-07-10 02:53:07.911603	2024-07-10 02:53:07.911603
73b6799d-ebd6-491f-9b47-272fb0f22914	b127ab20-d3e8-47cb-809a-c3843ef251f8	2024-07-10 02:53:07.911604	2024-07-10 02:53:07.911604
73b6799d-ebd6-491f-9b47-272fb0f22914	021fe4e3-196f-49c3-92b4-72975cce0979	2024-07-10 02:53:07.911605	2024-07-10 02:53:07.911605
73b6799d-ebd6-491f-9b47-272fb0f22914	73e75a98-a317-4034-a21a-b43974f65262	2024-07-10 02:53:07.911606	2024-07-10 02:53:07.911606
73b6799d-ebd6-491f-9b47-272fb0f22914	80e75eba-d970-4461-bcac-94b02f7871a5	2024-07-10 02:53:07.911607	2024-07-10 02:53:07.911607
73b6799d-ebd6-491f-9b47-272fb0f22914	83797eac-545d-43a7-865f-8d3053249be9	2024-07-10 02:53:07.911608	2024-07-10 02:53:07.911608
73b6799d-ebd6-491f-9b47-272fb0f22914	ca5ea509-e6d8-4dab-85c6-0e15a66cab24	2024-07-10 02:53:07.911609	2024-07-10 02:53:07.911609
73b6799d-ebd6-491f-9b47-272fb0f22914	eda3e8f9-a887-4e53-b4f0-63cc1da481b4	2024-07-10 02:53:07.91161	2024-07-10 02:53:07.91161
73b6799d-ebd6-491f-9b47-272fb0f22914	88f3e7b5-3fd4-431c-90e2-cd38c3232297	2024-07-10 02:53:07.911611	2024-07-10 02:53:07.911611
73b6799d-ebd6-491f-9b47-272fb0f22914	e83fbf50-1a29-4133-af74-f32834e21b02	2024-07-10 02:53:07.911612	2024-07-10 02:53:07.911612
73b6799d-ebd6-491f-9b47-272fb0f22914	4465ae1a-59a2-4814-b92d-660298bc62bc	2024-07-10 02:53:07.911613	2024-07-10 02:53:07.911613
73b6799d-ebd6-491f-9b47-272fb0f22914	352aecd8-b7d3-46e9-886c-17a6fc0c1176	2024-07-10 02:53:07.911614	2024-07-10 02:53:07.911614
73b6799d-ebd6-491f-9b47-272fb0f22914	5b8e13c3-1a1a-4847-a246-e0cfd27326b6	2024-07-10 02:53:07.911615	2024-07-10 02:53:07.911615
73b6799d-ebd6-491f-9b47-272fb0f22914	2872765f-6393-4de6-a100-290f401e64a1	2024-07-10 02:53:07.911616	2024-07-10 02:53:07.911616
73b6799d-ebd6-491f-9b47-272fb0f22914	257e2fcd-8962-480b-8723-9cfdd215e6b8	2024-07-10 02:53:07.911617	2024-07-10 02:53:07.911617
73b6799d-ebd6-491f-9b47-272fb0f22914	7cf3ff8c-315d-450b-aa52-6bc7d03762bf	2024-07-10 02:53:07.911618	2024-07-10 02:53:07.911618
73b6799d-ebd6-491f-9b47-272fb0f22914	bdac9253-ef8c-4b81-95a8-394d7a818eca	2024-07-10 02:53:07.911619	2024-07-10 02:53:07.911619
73b6799d-ebd6-491f-9b47-272fb0f22914	30c4fcc4-b630-4afb-b117-2a1547cfd9c3	2024-07-10 02:53:07.91162	2024-07-10 02:53:07.91162
73b6799d-ebd6-491f-9b47-272fb0f22914	9bee79e5-88b1-4f69-906a-7227260c5e35	2024-07-10 02:53:07.911621	2024-07-10 02:53:07.911621
00c12bd8-7470-40c5-938e-029b1239650c	c76e3196-b313-4d15-845c-fb9a4bec9840	2024-07-10 02:53:07.933438	2024-07-10 02:53:07.933438
00c12bd8-7470-40c5-938e-029b1239650c	ffe81eb0-560c-44cb-b53c-5e7733ac652d	2024-07-10 02:53:07.933439	2024-07-10 02:53:07.933439
00c12bd8-7470-40c5-938e-029b1239650c	f9b46964-fa73-43e8-8f5a-1041ed8ddefb	2024-07-10 02:53:07.93344	2024-07-10 02:53:07.93344
00c12bd8-7470-40c5-938e-029b1239650c	a49bc736-f731-42bd-ba6c-7064f7d75d5d	2024-07-10 02:53:07.933441	2024-07-10 02:53:07.933441
00c12bd8-7470-40c5-938e-029b1239650c	80295769-1e98-4bd6-9c91-29693ad689f1	2024-07-10 02:53:07.933442	2024-07-10 02:53:07.933442
00c12bd8-7470-40c5-938e-029b1239650c	cefb77ac-1f58-4f37-82b0-7fba3540782a	2024-07-10 02:53:07.933443	2024-07-10 02:53:07.933443
00c12bd8-7470-40c5-938e-029b1239650c	2d59a87a-c80f-437e-9d3d-fac82859cde3	2024-07-10 02:53:07.933444	2024-07-10 02:53:07.933444
bcc5cab0-6ac4-4aa7-b80d-7e7ef5113a10	504b4d4b-f5cd-43ca-8d5c-afc0b9c7e072	2024-07-10 02:53:07.936047	2024-07-10 02:53:07.936047
bcc5cab0-6ac4-4aa7-b80d-7e7ef5113a10	ab064625-5901-4ca4-86c6-7d310991c468	2024-07-10 02:53:07.936048	2024-07-10 02:53:07.936048
bcc5cab0-6ac4-4aa7-b80d-7e7ef5113a10	600fa561-7aa6-4bc9-b0a9-95bcaeb9b598	2024-07-10 02:53:07.936049	2024-07-10 02:53:07.936049
bcc5cab0-6ac4-4aa7-b80d-7e7ef5113a10	5b0121e7-c4ea-40f6-ac45-894bee6f52c7	2024-07-10 02:53:07.93605	2024-07-10 02:53:07.93605
bcc5cab0-6ac4-4aa7-b80d-7e7ef5113a10	9b413c7c-847b-4343-b5d1-21263385f949	2024-07-10 02:53:07.936051	2024-07-10 02:53:07.936051
bcc5cab0-6ac4-4aa7-b80d-7e7ef5113a10	6c2d12af-7d30-41ca-815a-df17a11054ce	2024-07-10 02:53:07.936052	2024-07-10 02:53:07.936052
bcc5cab0-6ac4-4aa7-b80d-7e7ef5113a10	1665c450-4d60-4607-ba2a-e593a0892b17	2024-07-10 02:53:07.936053	2024-07-10 02:53:07.936053
bcc5cab0-6ac4-4aa7-b80d-7e7ef5113a10	d5269684-097e-48fa-b6b0-abd1d3322a3d	2024-07-10 02:53:07.936054	2024-07-10 02:53:07.936054
bcc5cab0-6ac4-4aa7-b80d-7e7ef5113a10	0e81bdb3-c418-45b4-bdf4-75e9fbad332b	2024-07-10 02:53:07.936055	2024-07-10 02:53:07.936055
bcc5cab0-6ac4-4aa7-b80d-7e7ef5113a10	1e327b7f-00b0-4f8a-9295-e10c014826de	2024-07-10 02:53:07.936056	2024-07-10 02:53:07.936056
bcc5cab0-6ac4-4aa7-b80d-7e7ef5113a10	dd0ac738-9352-4128-90fb-33ea413d1ef9	2024-07-10 02:53:07.936057	2024-07-10 02:53:07.936057
bcc5cab0-6ac4-4aa7-b80d-7e7ef5113a10	8b606434-7590-4114-adcc-3ed18669a19d	2024-07-10 02:53:07.936058	2024-07-10 02:53:07.936058
bcc5cab0-6ac4-4aa7-b80d-7e7ef5113a10	c5231c99-cbe3-41bf-bcc6-3f90b463c07d	2024-07-10 02:53:07.936059	2024-07-10 02:53:07.936059
bcc5cab0-6ac4-4aa7-b80d-7e7ef5113a10	4165e784-1d77-4c59-83ae-bb8097ddaf48	2024-07-10 02:53:07.93606	2024-07-10 02:53:07.93606
bcc5cab0-6ac4-4aa7-b80d-7e7ef5113a10	691f7440-1d77-4c66-9fc6-3abfd0593cb8	2024-07-10 02:53:07.936061	2024-07-10 02:53:07.936061
bcc5cab0-6ac4-4aa7-b80d-7e7ef5113a10	836e79ba-5bdd-4670-a7a0-6f137ac9db89	2024-07-10 02:53:07.936062	2024-07-10 02:53:07.936062
bcc5cab0-6ac4-4aa7-b80d-7e7ef5113a10	85758192-2ecc-45e8-908d-409bb20272d9	2024-07-10 02:53:07.936063	2024-07-10 02:53:07.936063
8d01b6df-d26b-4b50-81df-b01167bedf0c	ce029379-9607-460d-858e-ddf66f914821	2024-07-10 02:53:07.942176	2024-07-10 02:53:07.942176
8d01b6df-d26b-4b50-81df-b01167bedf0c	1e327b7f-00b0-4f8a-9295-e10c014826de	2024-07-10 02:53:07.942177	2024-07-10 02:53:07.942177
8d01b6df-d26b-4b50-81df-b01167bedf0c	dd0ac738-9352-4128-90fb-33ea413d1ef9	2024-07-10 02:53:07.942178	2024-07-10 02:53:07.942178
8d01b6df-d26b-4b50-81df-b01167bedf0c	8b606434-7590-4114-adcc-3ed18669a19d	2024-07-10 02:53:07.942179	2024-07-10 02:53:07.942179
8d01b6df-d26b-4b50-81df-b01167bedf0c	9aaa4ca6-33b4-41d9-81e3-9f6e71590d0c	2024-07-10 02:53:07.94218	2024-07-10 02:53:07.94218
8d01b6df-d26b-4b50-81df-b01167bedf0c	c6dbeb2d-52e9-411b-8e84-5d594b32b20d	2024-07-10 02:53:07.942181	2024-07-10 02:53:07.942181
8d01b6df-d26b-4b50-81df-b01167bedf0c	90ad8535-3ea7-4f6e-8e58-8f316377cee0	2024-07-10 02:53:07.942182	2024-07-10 02:53:07.942182
8d01b6df-d26b-4b50-81df-b01167bedf0c	82e4aae1-1d14-4511-a21b-37cc6ce28391	2024-07-10 02:53:07.942183	2024-07-10 02:53:07.942183
8d01b6df-d26b-4b50-81df-b01167bedf0c	f5405977-8f91-4954-bafc-e4c4c5f628e4	2024-07-10 02:53:07.942184	2024-07-10 02:53:07.942184
8d01b6df-d26b-4b50-81df-b01167bedf0c	91b0dd33-0bbe-4641-844d-ec6294163f4d	2024-07-10 02:53:07.942185	2024-07-10 02:53:07.942185
8d01b6df-d26b-4b50-81df-b01167bedf0c	ec412394-c888-4deb-a87d-7138ae9b12e0	2024-07-10 02:53:07.942186	2024-07-10 02:53:07.942186
8d01b6df-d26b-4b50-81df-b01167bedf0c	b3da3e77-b9d1-49c9-99f3-9e81e8ed32d2	2024-07-10 02:53:07.942187	2024-07-10 02:53:07.942187
8d01b6df-d26b-4b50-81df-b01167bedf0c	d47a4dde-9aea-42d7-8723-9d263a9b2c8e	2024-07-10 02:53:07.942188	2024-07-10 02:53:07.942188
8d01b6df-d26b-4b50-81df-b01167bedf0c	20e9761b-a5a4-4086-b25b-58987ad95cc7	2024-07-10 02:53:07.942189	2024-07-10 02:53:07.942189
8d01b6df-d26b-4b50-81df-b01167bedf0c	9e7bdca1-35da-4d3f-9799-00f4e79323a9	2024-07-10 02:53:07.94219	2024-07-10 02:53:07.94219
8d01b6df-d26b-4b50-81df-b01167bedf0c	1a426476-0f40-4582-9800-a167de643a4f	2024-07-10 02:53:07.942191	2024-07-10 02:53:07.942191
8d01b6df-d26b-4b50-81df-b01167bedf0c	f2a9d1fb-0a3e-4e8b-ad32-d47d840505a4	2024-07-10 02:53:07.942192	2024-07-10 02:53:07.942192
8d01b6df-d26b-4b50-81df-b01167bedf0c	8a06286f-0df6-4d21-9fe3-975758125239	2024-07-10 02:53:07.942193	2024-07-10 02:53:07.942193
8d01b6df-d26b-4b50-81df-b01167bedf0c	321b797f-3b85-4f4c-8ba2-f91aa24c4161	2024-07-10 02:53:07.942194	2024-07-10 02:53:07.942194
8d01b6df-d26b-4b50-81df-b01167bedf0c	2ced73f0-461e-4bf4-ac89-ad8436c4d59c	2024-07-10 02:53:07.942195	2024-07-10 02:53:07.942195
8d01b6df-d26b-4b50-81df-b01167bedf0c	f2c1794b-e513-47f1-a009-0b3329080565	2024-07-10 02:53:07.942196	2024-07-10 02:53:07.942196
8d01b6df-d26b-4b50-81df-b01167bedf0c	6a5d87b1-3a89-48c4-b9b7-e9dc6bed11b8	2024-07-10 02:53:07.942197	2024-07-10 02:53:07.942197
8d01b6df-d26b-4b50-81df-b01167bedf0c	c6a9c058-19fb-4eff-b050-517750e87888	2024-07-10 02:53:07.942198	2024-07-10 02:53:07.942198
8d01b6df-d26b-4b50-81df-b01167bedf0c	ddfea278-864d-4423-bac0-ec74ffc59752	2024-07-10 02:53:07.942199	2024-07-10 02:53:07.942199
8d01b6df-d26b-4b50-81df-b01167bedf0c	ae65c638-a388-4d58-987a-3a05235a6da6	2024-07-10 02:53:07.9422	2024-07-10 02:53:07.9422
8d01b6df-d26b-4b50-81df-b01167bedf0c	5cb4530b-317e-429a-be88-5a52e4030aa0	2024-07-10 02:53:07.942201	2024-07-10 02:53:07.942201
8d01b6df-d26b-4b50-81df-b01167bedf0c	d633cc03-c384-4840-ab62-c2e9d96370a2	2024-07-10 02:53:07.942202	2024-07-10 02:53:07.942202
8d01b6df-d26b-4b50-81df-b01167bedf0c	0bf21e0d-9d18-44c2-8c0b-506a32b41f6a	2024-07-10 02:53:07.942203	2024-07-10 02:53:07.942203
8d01b6df-d26b-4b50-81df-b01167bedf0c	1f08e582-1628-480a-b188-084e6eb82bed	2024-07-10 02:53:07.942204	2024-07-10 02:53:07.942204
8d01b6df-d26b-4b50-81df-b01167bedf0c	b700bf02-b0dd-49f9-b90e-d279dddbc0d9	2024-07-10 02:53:07.942205	2024-07-10 02:53:07.942205
8d01b6df-d26b-4b50-81df-b01167bedf0c	4bca273e-0832-443e-a20b-7f3fdebc57c0	2024-07-10 02:53:07.942206	2024-07-10 02:53:07.942206
8d01b6df-d26b-4b50-81df-b01167bedf0c	acae13ed-8b52-4762-a3e2-03113f5f4ca1	2024-07-10 02:53:07.942207	2024-07-10 02:53:07.942207
8d01b6df-d26b-4b50-81df-b01167bedf0c	f005d6f4-2170-4ef5-9213-9ada244d29d2	2024-07-10 02:53:07.942208	2024-07-10 02:53:07.942208
8d01b6df-d26b-4b50-81df-b01167bedf0c	2be7ef08-e2fe-4598-ad67-ff731f1b8750	2024-07-10 02:53:07.942209	2024-07-10 02:53:07.942209
8d01b6df-d26b-4b50-81df-b01167bedf0c	7bfa7ca8-f298-4ac7-8aaf-5df4356c2929	2024-07-10 02:53:07.94221	2024-07-10 02:53:07.94221
8d01b6df-d26b-4b50-81df-b01167bedf0c	64746194-1303-41d5-8177-70deab18c1fd	2024-07-10 02:53:07.942211	2024-07-10 02:53:07.942211
8d01b6df-d26b-4b50-81df-b01167bedf0c	550a041a-8322-44e3-8450-d74480e06633	2024-07-10 02:53:07.942212	2024-07-10 02:53:07.942212
8d01b6df-d26b-4b50-81df-b01167bedf0c	f7af494a-c5f4-4b34-907a-be688598eacb	2024-07-10 02:53:07.942213	2024-07-10 02:53:07.942213
8d01b6df-d26b-4b50-81df-b01167bedf0c	23eded05-4e26-45ce-bb47-ca9df2007d54	2024-07-10 02:53:07.942214	2024-07-10 02:53:07.942214
8d01b6df-d26b-4b50-81df-b01167bedf0c	e594d5cd-fd64-45ac-bb09-074b460b6e3f	2024-07-10 02:53:07.942215	2024-07-10 02:53:07.942215
8d01b6df-d26b-4b50-81df-b01167bedf0c	a3ba044f-d1f0-48ee-b0ff-df4fb87c2616	2024-07-10 02:53:07.942216	2024-07-10 02:53:07.942216
8d01b6df-d26b-4b50-81df-b01167bedf0c	47ca9a63-550c-4203-97e7-5225a45b6459	2024-07-10 02:53:07.942217	2024-07-10 02:53:07.942217
8d01b6df-d26b-4b50-81df-b01167bedf0c	08dcf0d1-a01e-4eef-95ce-2bbc783eed9e	2024-07-10 02:53:07.942218	2024-07-10 02:53:07.942218
8d01b6df-d26b-4b50-81df-b01167bedf0c	ec4e44fa-a6f1-4cdf-8baa-85ebac7912a7	2024-07-10 02:53:07.942219	2024-07-10 02:53:07.942219
8d01b6df-d26b-4b50-81df-b01167bedf0c	2b496e84-cf57-4a5b-abd4-0ff652fe4138	2024-07-10 02:53:07.94222	2024-07-10 02:53:07.94222
8d01b6df-d26b-4b50-81df-b01167bedf0c	6e6e2de2-156c-4af1-8f15-1c9df5d1e074	2024-07-10 02:53:07.942221	2024-07-10 02:53:07.942221
8d01b6df-d26b-4b50-81df-b01167bedf0c	634c7e46-3f5b-4f96-8cbd-4e2f49fc1209	2024-07-10 02:53:07.942222	2024-07-10 02:53:07.942222
8d01b6df-d26b-4b50-81df-b01167bedf0c	ca5ea509-e6d8-4dab-85c6-0e15a66cab24	2024-07-10 02:53:07.942223	2024-07-10 02:53:07.942223
8d01b6df-d26b-4b50-81df-b01167bedf0c	eda3e8f9-a887-4e53-b4f0-63cc1da481b4	2024-07-10 02:53:07.942224	2024-07-10 02:53:07.942224
8d01b6df-d26b-4b50-81df-b01167bedf0c	27c05dc9-6ef7-4d03-ba09-c3a0e9a4459e	2024-07-10 02:53:07.942225	2024-07-10 02:53:07.942225
8d01b6df-d26b-4b50-81df-b01167bedf0c	455150f8-a28c-4644-b6c9-bb53639a44ae	2024-07-10 02:53:07.942226	2024-07-10 02:53:07.942226
\.


--
-- Data for Name: roles; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.roles (id, name, description, created_at, updated_at) FROM stdin;
8d01b6df-d26b-4b50-81df-b01167bedf0c	Super Admin	\N	2024-06-06 09:29:38.549134	2024-06-06 09:29:38.549134
bcc5cab0-6ac4-4aa7-b80d-7e7ef5113a10	Cashier	\N	2024-06-06 09:29:38.549135	2024-06-06 09:29:38.549135
00c12bd8-7470-40c5-938e-029b1239650c	Driver	\N	2024-06-06 09:29:38.549136	2024-06-06 09:29:38.549136
73b6799d-ebd6-491f-9b47-272fb0f22914	Inventory	\N	2024-06-06 09:29:38.549137	2024-06-06 09:29:38.549137
\.


--
-- Data for Name: schema_migrations; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.schema_migrations (version, dirty) FROM stdin;
202406061503	f
\.


--
-- Data for Name: sequences; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.sequences (id, unique_identifier, sequence, created_at, updated_at) FROM stdin;
06e4bca8-91f9-47ee-8d66-67428e10330e	DO-202466-	1	2024-06-06 09:39:05.345173	2024-06-06 09:39:05.345173
b6225252-0bf1-4c84-9513-b70c19f2191f	DO-2024612-	1	2024-06-12 09:15:39.581299	2024-06-12 09:15:39.581299
d0ca945a-6ad8-47ef-bda6-92fb0742ad94	DO-2024618-	1	2024-06-18 06:59:35.498299	2024-06-18 06:59:35.498299
d438eba8-abea-48c4-a3c2-c3eca9ad832c	DO-2024618-	2	2024-06-18 06:59:57.125786	2024-06-18 06:59:57.125786
05dd19c5-2b56-4e72-ae6c-76a8ec3f226f	DO-2024618-	3	2024-06-18 07:18:08.594493	2024-06-18 07:18:08.594493
3ad4ff60-2d56-43f1-8dc6-6be0746e8c6c	DO-2024618-	4	2024-06-18 07:18:24.976781	2024-06-18 07:18:24.976781
8ebb8649-ed4d-40b9-8182-278d1a12180b	DO-2024620-	1	2024-06-20 02:22:35.497725	2024-06-20 02:22:35.497725
3d0a2c30-29ae-4fc0-a396-7eef1fdba0eb	DO-2024620-	2	2024-06-20 09:46:23.735342	2024-06-20 09:46:23.735342
8c988676-bede-46a8-95fc-43382b1cda9a	DO-2024627-	1	2024-06-27 09:01:22.861923	2024-06-27 09:01:22.861923
ae93d4a8-6636-4d41-96b2-ba778fbfe15e	DO-2024627-	2	2024-06-27 09:22:42.195105	2024-06-27 09:22:42.195105
60b20145-eda9-49fa-adbd-fb0b0ef8135b	DO-2024628-	1	2024-06-28 00:28:26.194375	2024-06-28 00:28:26.194375
283110e2-83d2-4bd1-afa2-6eb318c168e9	DO-2024629-	1	2024-06-29 09:32:28.477671	2024-06-29 09:32:28.477671
8cd0623a-aa7f-44b0-9ce7-1a2229a02d2b	DO-202478-	1	2024-07-08 12:28:42.91577	2024-07-08 12:28:42.91577
\.


--
-- Data for Name: shop_order_items; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.shop_order_items (id, shop_order_id, product_unit_id, platform_product_id, image_link, quantity, original_price, sale_price, created_at, updated_at) FROM stdin;
516bf2ce-567a-4a47-ad8f-c0a0c14d1db8	7143f491-819a-4282-8266-023b84f999ee	61218f4d-b5ac-41db-8db6-1f3c3f18a10e	1729700799344314083	https://p16-oec-va.ibyteimg.com/tos-maliva-i-o3syd03w52-us/b1fd4a14759d43a299bca73cf44c3b33~tplv-o3syd03w52-origin-jpeg.jpeg?from=1413970683	2.00	100000.00	100000.00	2024-07-01 08:57:12.720017	2024-07-01 08:57:12.720017
a0bc9abe-b936-456f-a1ab-c523408bb499	7677c35c-048f-4938-850a-e114d8bb2e48	61218f4d-b5ac-41db-8db6-1f3c3f18a10e	1729700799344314083	https://p16-oec-va.ibyteimg.com/tos-maliva-i-o3syd03w52-us/b1fd4a14759d43a299bca73cf44c3b33~tplv-o3syd03w52-origin-jpeg.jpeg?from=1413970683	3.00	100000.00	100000.00	2024-07-01 09:21:27.694613	2024-07-01 09:21:27.694613
6c4ae63b-5654-4a67-84d0-15997f536891	3f253f24-612b-4170-b6db-06ec8b883178	61218f4d-b5ac-41db-8db6-1f3c3f18a10e	1729700799344314083	https://p16-oec-va.ibyteimg.com/tos-maliva-i-o3syd03w52-us/b1fd4a14759d43a299bca73cf44c3b33~tplv-o3syd03w52-origin-jpeg.jpeg?from=1413970683	2.00	100000.00	100000.00	2024-07-01 10:51:27.823219	2024-07-01 10:51:27.823219
\.


--
-- Data for Name: shop_orders; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.shop_orders (id, tracking_number, platform_identifier, platform_type, tracking_status, recipient_name, recipient_full_address, recipient_phone_number, shipping_fee, service_fee, total_original_product_price, subtotal, tax, total_amount, created_at, updated_at) FROM stdin;
7677c35c-048f-4938-850a-e114d8bb2e48	TT1001952433	579010255970076682	TIKTOK_SHOP	CANCEL	a***	Indonesia, Jakarta, South Jakarta, Se*******, Se********,Su******************************************************************************************	(+86)123******43	116122.00	1000.00	300000.00	300000.00	0.00	417122.00	2024-07-01 09:21:27.689555	2024-07-04 17:35:42.967333
3f253f24-612b-4170-b6db-06ec8b883178	TT1003606715	579010518109620234	TIKTOK_SHOP	COMPLETED	a***	Indonesia, Jakarta, South Jakarta, Se*******, Se********,Su******************************************************************************************	(+86)123******43	77414.00	1000.00	200000.00	200000.00	0.00	278414.00	2024-07-01 10:51:27.797019	2024-07-07 17:57:50.73359
7143f491-819a-4282-8266-023b84f999ee	TT1005952133	579010173564061706	TIKTOK_SHOP	COMPLETED	a***	Indonesia, Jakarta, South Jakarta, Se*******, Se********,Su******************************************************************************************	(+86)123******43	77414.00	1000.00	200000.00	200000.00	0.00	278414.00	2024-07-01 08:57:12.709649	2024-07-07 19:22:09.902767
\.


--
-- Data for Name: shopee_configs; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.shopee_configs (partner_id, partner_key, access_token, refresh_token, created_at, updated_at) FROM stdin;
		\N	\N	2024-06-06 09:29:38.564141	2024-06-06 09:29:38.564141
\.


--
-- Data for Name: supplier_types; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.supplier_types (id, name, description, created_at, updated_at) FROM stdin;
8b6d5f25-0d77-4fc8-ba51-ef8aa4e9a515	Supplier A	\N	2024-06-06 09:29:38.603415	2024-06-06 09:29:38.603415
d685a90f-1ce6-41ce-bc50-b8f83d987b5e	Supplier Produk Cat	\N	2024-06-27 08:49:15.159398	2024-06-27 08:49:15.159398
8ef6b4d0-997d-4472-8d13-13a1c7de9e8f	Gunung Sari	Supplier Listrik	2024-07-10 15:26:53.682671	2024-07-10 15:27:31.34975
\.


--
-- Data for Name: suppliers; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.suppliers (id, supplier_type_id, code, name, is_active, address, phone, email, description, created_at, updated_at) FROM stdin;
e0225ca5-4094-47e9-8494-43795dfddd97	8b6d5f25-0d77-4fc8-ba51-ef8aa4e9a515	S-123	Supplier A	t	Jln. Tilak	+6285286869797	\N	\N	2024-06-06 09:29:38.607503	2024-06-06 09:29:38.607503
e6815351-8cd2-4471-a81e-ac21eb65e7ea	d685a90f-1ce6-41ce-bc50-b8f83d987b5e	CAT-001	Supplier Cat Dinding	t	Jln. Tasbih	+6285261302277	supplier.cat@gmail.com	\N	2024-06-27 08:49:45.753438	2024-06-27 08:49:45.753438
48fba4fe-787a-45fb-8c65-3ff6418b09a6	d685a90f-1ce6-41ce-bc50-b8f83d987b5e	C-002	Produsen Cat Dinding	t	Jln. Tasbih	+6285295695825	supplier2@gmail.com	\N	2024-06-28 00:18:18.854815	2024-06-28 00:18:18.854815
0c5938c0-d01c-41d7-ac33-254c04149893	d685a90f-1ce6-41ce-bc50-b8f83d987b5e	gngsr	Sinar Surya	t	Jl. Veteran No12	+6285155118251	dkms@gmail.com	Jual bahan bangunan	2024-07-10 15:28:26.65228	2024-07-10 15:28:26.65228
\.


--
-- Data for Name: tiktok_configs; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.tiktok_configs (app_key, app_secret, warehouse_id, shop_id, shop_cipher, access_token, refresh_token, created_at, updated_at) FROM stdin;
6bgdj91pdnm9v	6cfbf8374b80618ca0d1b5eafd87ca0e23554e57	7333151372009178885	7495591168837323491	ROW_ij-EHgAAAAAFH7_LWApa2DADTZh6ANIA	ROW_LWeg7gAAAAAjuAWBFh7OoD5X1P3Y_MSshmDMZG4rcBi48ay1C4wMwLRWz0ZiBR2yKAPor8EMpLUvnjsWFSTx4fWFn8Me9B2B1RaFihkHEQtnSHt9KemECgnnwEBlRtIP0-vUEB2WUSRnqxipVJ3T8J3MBIGq5NzSv7gfG3P7nE7J0BmBPhZkww	ROW_jJQ6wwAAAACx9YYnlqBwGkIEDGqW7sd7mABV-T1yzbDUOBEP29Tzvo1Y5oV59cCzpz_ukS8BG_k	2024-06-06 09:29:38.562214	2024-07-10 17:00:00.375186
\.


--
-- Data for Name: tiktok_products; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.tiktok_products (tiktok_product_id, product_id, status, created_at, updated_at) FROM stdin;
1729654934777792227	e1bf0592-7850-4602-a740-6aae98dfd281	ACTIVE	2024-06-07 07:23:55.636879	2024-06-07 07:23:55.636879
1729700241018881763	69a3f894-1ff2-4f61-97c2-3c957eea7914	ACTIVE	2024-06-27 08:54:55.224031	2024-06-27 08:54:55.224031
1729700799344314083	d0813034-19a1-406f-af25-c18d2f301614	ACTIVE	2024-06-28 00:22:39.147562	2024-06-28 00:22:39.147562
1729713664706578147	412d3012-8a9f-44ea-a8a0-273aaa6a9cfc	ACTIVE	2024-07-09 00:46:38.693212	2024-07-09 00:46:38.693212
\.


--
-- Data for Name: transaction_item_costs; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.transaction_item_costs (id, transaction_item_id, qty, base_cost_price, total_cost_price, created_at, updated_at) FROM stdin;
c64110bc-a75d-4d52-a0db-996b208bcf31	639f1e70-aaf0-4771-9afd-ca072f8fd247	2.00	77972.56	155945.12	2024-06-07 15:51:32.82861	2024-06-07 15:51:32.82861
1b888732-6ff5-43e0-957d-957b6cbba2e2	4620dc34-6ab9-4424-8b55-6a7befd9f8d7	1.00	1785.67	1785.67	2024-06-26 04:30:11.027517	2024-06-26 04:30:11.027517
4f781bc0-f7d8-4c34-9394-cbb94b292435	05e2f90f-89cd-415a-b89b-8cf474c227b0	1.00	1785.67	1785.67	2024-06-26 04:30:53.729153	2024-06-26 04:30:53.729153
59de502f-91bd-43d7-b8f5-e55a1f14fca0	82339f07-3b83-4ae7-bb1c-233e0758b03e	1.00	1785.67	1785.67	2024-06-26 04:47:50.037399	2024-06-26 04:47:50.037399
7d5b1941-d9ae-4dad-b42f-2b0460a68275	0ede5d68-7d2b-4580-81fd-fef218d19319	1.00	1785.67	1785.67	2024-06-26 04:49:48.491309	2024-06-26 04:49:48.491309
7f17ea12-b1f0-4c5c-9126-f17d40476ba8	8fcae4b9-c220-4e2b-af75-821514460cde	1.00	1785.67	1785.67	2024-06-26 05:01:03.292088	2024-06-26 05:01:03.292088
2664e45b-c47e-4f57-90d1-3ec469427f39	7a446890-fd15-48c0-a44b-a82a3202dc33	1.00	1785.67	1785.67	2024-06-26 05:01:52.886827	2024-06-26 05:01:52.886827
16b3bd84-04f5-40ff-9703-f10434187936	b74e0bc8-8f95-4e12-b21f-cd39d8258610	1.00	1785.67	1785.67	2024-06-26 05:03:22.479286	2024-06-26 05:03:22.479286
4bfd7f98-4e9a-4033-91c4-36b1a2a8f5f0	576016a6-3d74-4fdc-a047-1b51a91afa3c	1.00	1785.67	1785.67	2024-06-26 05:03:37.757247	2024-06-26 05:03:37.757247
280433cd-c4bb-42d5-a179-fa1fa2633d7e	8bcff1b7-e31c-4e9a-9b51-d9a53360ed4c	1.00	1785.67	1785.67	2024-06-26 05:04:06.60931	2024-06-26 05:04:06.60931
cb6074bc-1102-47ef-9e50-72f2cc603b66	70179da8-5756-46ca-9940-38a9bf9ba957	1.00	1785.67	1785.67	2024-06-26 06:54:51.900995	2024-06-26 06:54:51.900995
da9954c1-8efd-4861-9b91-0d56b4d7a4d9	78a77e4e-2bc9-48c6-a42a-d722f963a70a	1.00	36730.77	36730.77	2024-06-27 09:28:28.526633	2024-06-27 09:28:28.526633
a42e3925-1c82-45f8-b9e4-71903ebcd0af	751d23f2-4d2b-4a2d-8618-073ad373ecc5	1.00	1785.67	1785.67	2024-06-27 09:28:28.526634	2024-06-27 09:28:28.526634
39d5184b-c777-498c-94c6-04c0bb80f865	1886c6d6-e53a-447f-b668-26d161725226	1.00	36730.77	36730.77	2024-06-27 09:29:55.841179	2024-06-27 09:29:55.841179
354adaff-465f-49da-a1f7-c04fc9f5639b	ee6e35ba-3611-4a9d-b112-bb981cbb63cc	2.00	36730.77	73461.54	2024-06-27 09:31:20.324646	2024-06-27 09:31:20.324646
3551bc55-aa41-453a-b985-95d813ce14c0	af37dd48-dbc7-48e2-ae77-b004831ee48b	1.00	1785.67	1785.67	2024-06-27 09:33:06.750599	2024-06-27 09:33:06.750599
2376e2ed-2455-4646-bd51-c5e5c685ca94	afbd26da-c464-4236-93bf-1b74466c4d1c	1.00	36730.77	36730.77	2024-06-27 09:33:06.7506	2024-06-27 09:33:06.7506
5069a95f-1999-485e-a04d-3e57a88d2af9	9ff999c2-4dad-46a0-8f1c-d0c83693072f	3.00	70740.74	212222.22	2024-06-28 00:33:48.929252	2024-06-28 00:33:48.929252
846c5dea-50db-417c-89e8-f729b48101cb	c04f2d72-17a9-4242-be43-d183d1118c9e	4.00	70740.74	2829629.60	2024-06-28 02:35:07.355861	2024-06-28 02:35:07.355861
94685802-1535-4637-8c37-d459bb47cff5	c79a009c-4559-4868-9d05-1e34e8cdb60c	3.00	70740.74	2122222.20	2024-06-29 09:58:01.419393	2024-06-29 09:58:01.419393
987f442a-3d49-4670-856c-477f0f199e81	38971ff1-5c85-4fce-aaf4-764401325b74	1.00	70740.74	707407.40	2024-07-01 03:08:16.945027	2024-07-01 03:08:16.945027
c13cb272-c79d-462b-a405-b8e127292de9	f03b2ffc-58c2-4f04-b98f-d6983f788af6	2.00	68716.05	1374321.00	2024-07-01 09:20:59.892134	2024-07-01 09:20:59.892134
e9c26bcf-4e9b-4b9b-8507-1afc9c9b100c	a657503e-be4d-4a9e-89ba-1ad7dea7b29d	4.00	1785.67	7142.68	2024-07-08 15:54:59.919424	2024-07-08 15:54:59.919424
173edbb9-72a1-4b17-a8af-159d9a28c1b2	bcbb6ecc-35d0-417c-bdca-3ce41ad51fb4	1.00	10000.00	10000.00	2024-07-08 16:02:26.855509	2024-07-08 16:02:26.855509
48e3cb0e-800d-4e0e-bb8d-1b2b1f652ec7	9dbf1d48-b01b-4b83-8ab3-e5ba52e31576	1.00	15500.00	15500.00	2024-07-08 16:03:43.149881	2024-07-08 16:03:43.149881
7fe0453b-85f5-4c23-9988-c62e1c904d02	c2fa4501-8bdd-4da3-9d02-856e520e626a	1.00	10909.09	10909.09	2024-07-10 02:16:43.122534	2024-07-10 02:16:43.122534
54e78f3b-fda6-4524-a2a6-5ed849aac065	9142d8a7-4a0f-4bf1-baf5-f8d168bfbd16	1.00	10909.09	10909.09	2024-07-10 02:19:11.455505	2024-07-10 02:19:11.455505
\.


--
-- Data for Name: transaction_items; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.transaction_items (id, transaction_id, product_unit_id, qty, price_per_unit, discount_per_unit, created_at, updated_at) FROM stdin;
639f1e70-aaf0-4771-9afd-ca072f8fd247	85ee56d5-9a8f-42e7-8bdc-3aa1512ec5d0	cec71c3c-5a20-461b-9370-051ab3eeba76	2.00	150000.00	\N	2024-06-07 15:51:32.825419	2024-06-07 15:51:32.825419
4620dc34-6ab9-4424-8b55-6a7befd9f8d7	d7e6f3ac-a4ba-4e14-901e-9e63c67b1458	cec71c3c-5a20-461b-9370-051ab3eeba76	1.00	150000.00	10000.00	2024-06-26 04:30:11.023709	2024-06-26 04:30:11.023709
05e2f90f-89cd-415a-b89b-8cf474c227b0	67acec7e-bb3b-499a-92e7-4d23f6f7eb31	cec71c3c-5a20-461b-9370-051ab3eeba76	1.00	150000.00	10000.00	2024-06-26 04:30:53.727689	2024-06-26 04:30:53.727689
82339f07-3b83-4ae7-bb1c-233e0758b03e	458047e3-4000-487e-964d-4be12d701557	cec71c3c-5a20-461b-9370-051ab3eeba76	1.00	150000.00	10000.00	2024-06-26 04:47:50.034012	2024-06-26 04:47:50.034012
0ede5d68-7d2b-4580-81fd-fef218d19319	26bfac7e-9b6c-4bb4-b59e-e25870e83ee9	cec71c3c-5a20-461b-9370-051ab3eeba76	1.00	150000.00	10000.00	2024-06-26 04:49:48.490382	2024-06-26 04:49:48.490382
8fcae4b9-c220-4e2b-af75-821514460cde	9d4d8b64-22bc-4d67-afa4-d2719ce861bd	cec71c3c-5a20-461b-9370-051ab3eeba76	1.00	150000.00	10000.00	2024-06-26 05:01:03.288086	2024-06-26 05:01:03.288086
7a446890-fd15-48c0-a44b-a82a3202dc33	f5e5036c-0a13-481c-ab93-1ccf2cf51ca4	cec71c3c-5a20-461b-9370-051ab3eeba76	1.00	150000.00	10000.00	2024-06-26 05:01:52.885189	2024-06-26 05:01:52.885189
b74e0bc8-8f95-4e12-b21f-cd39d8258610	7efd1898-66be-45ae-bf09-85b4478078df	cec71c3c-5a20-461b-9370-051ab3eeba76	1.00	150000.00	10000.00	2024-06-26 05:03:22.478006	2024-06-26 05:03:22.478006
576016a6-3d74-4fdc-a047-1b51a91afa3c	434c6873-eca8-41fe-a3ea-f28ef64e03f0	cec71c3c-5a20-461b-9370-051ab3eeba76	1.00	150000.00	10000.00	2024-06-26 05:03:37.756619	2024-06-26 05:03:37.756619
8bcff1b7-e31c-4e9a-9b51-d9a53360ed4c	073c3ef1-72f9-4cbf-bc04-5a19b378ec22	cec71c3c-5a20-461b-9370-051ab3eeba76	1.00	150000.00	10000.00	2024-06-26 05:04:06.608359	2024-06-26 05:04:06.608359
70179da8-5756-46ca-9940-38a9bf9ba957	44a8d583-eb5d-4a32-a8b8-0fb084a8dfde	cec71c3c-5a20-461b-9370-051ab3eeba76	1.00	150000.00	10000.00	2024-06-26 06:54:51.899097	2024-06-26 06:54:51.899097
78a77e4e-2bc9-48c6-a42a-d722f963a70a	b23eb768-6987-4af1-969a-9d396ddc945a	a28b11ad-6f08-4607-badd-d481c0e6be4c	1.00	100000.00	\N	2024-06-27 09:28:28.52435	2024-06-27 09:28:28.52435
751d23f2-4d2b-4a2d-8618-073ad373ecc5	b23eb768-6987-4af1-969a-9d396ddc945a	cec71c3c-5a20-461b-9370-051ab3eeba76	1.00	150000.00	10000.00	2024-06-27 09:28:28.524351	2024-06-27 09:28:28.524351
1886c6d6-e53a-447f-b668-26d161725226	4bec0189-8da7-4f2f-98f0-9ef92ad00a36	a28b11ad-6f08-4607-badd-d481c0e6be4c	1.00	100000.00	\N	2024-06-27 09:29:55.84035	2024-06-27 09:29:55.84035
ee6e35ba-3611-4a9d-b112-bb981cbb63cc	5c8140ea-4462-4f85-8acb-0d1998bcb56e	a28b11ad-6f08-4607-badd-d481c0e6be4c	2.00	100000.00	\N	2024-06-27 09:31:20.323031	2024-06-27 09:31:20.323031
af37dd48-dbc7-48e2-ae77-b004831ee48b	a7364625-a12f-4d50-8fa2-32ec8df06b26	cec71c3c-5a20-461b-9370-051ab3eeba76	1.00	150000.00	10000.00	2024-06-27 09:33:06.749631	2024-06-27 09:33:06.749631
afbd26da-c464-4236-93bf-1b74466c4d1c	a7364625-a12f-4d50-8fa2-32ec8df06b26	a28b11ad-6f08-4607-badd-d481c0e6be4c	1.00	100000.00	\N	2024-06-27 09:33:06.749632	2024-06-27 09:33:06.749632
9ff999c2-4dad-46a0-8f1c-d0c83693072f	73e1140c-dfb6-4d2f-9474-25eedf7fa425	61218f4d-b5ac-41db-8db6-1f3c3f18a10e	3.00	100000.00	\N	2024-06-28 00:33:48.928096	2024-06-28 00:33:48.928096
c04f2d72-17a9-4242-be43-d183d1118c9e	3a70f48c-49eb-4100-a1ff-0781bb03ea12	61218f4d-b5ac-41db-8db6-1f3c3f18a10e	4.00	100000.00	\N	2024-06-28 02:35:07.354113	2024-06-28 02:35:07.354113
c79a009c-4559-4868-9d05-1e34e8cdb60c	d48b7681-531a-4aae-a1b0-6b57d366d187	61218f4d-b5ac-41db-8db6-1f3c3f18a10e	3.00	100000.00	\N	2024-06-29 09:58:01.415361	2024-06-29 09:58:01.415361
38971ff1-5c85-4fce-aaf4-764401325b74	cb683226-9b91-46ad-a13a-c7c38dedc290	61218f4d-b5ac-41db-8db6-1f3c3f18a10e	1.00	100000.00	\N	2024-07-01 03:08:16.941508	2024-07-01 03:08:16.941508
f03b2ffc-58c2-4f04-b98f-d6983f788af6	643d3f68-95a8-461c-ae84-ed51a88c6dfb	61218f4d-b5ac-41db-8db6-1f3c3f18a10e	2.00	100000.00	\N	2024-07-01 09:20:59.883596	2024-07-01 09:20:59.883596
a657503e-be4d-4a9e-89ba-1ad7dea7b29d	b20d6bee-8020-423a-8b92-5c794edaa2eb	cec71c3c-5a20-461b-9370-051ab3eeba76	4.00	150000.00	10000.00	2024-07-08 15:54:59.918047	2024-07-08 15:54:59.918047
bcbb6ecc-35d0-417c-bdca-3ce41ad51fb4	95e5f612-d352-48eb-8338-776e44d26501	4ca122af-8969-44d6-ac51-fd3a1a8c504b	1.00	10000.00	\N	2024-07-08 16:02:26.853958	2024-07-08 16:02:26.853958
9dbf1d48-b01b-4b83-8ab3-e5ba52e31576	382a3371-4fe2-41ae-b02b-091cd67e78c6	4ca122af-8969-44d6-ac51-fd3a1a8c504b	1.00	10000.00	\N	2024-07-08 16:03:43.148307	2024-07-08 16:03:43.148307
c2fa4501-8bdd-4da3-9d02-856e520e626a	87d66a03-3756-4afb-a84b-c7f79ffaa788	f8fd2e10-1a96-44ae-9898-50fa57d8dcf9	1.00	12000.00	\N	2024-07-10 02:16:43.119502	2024-07-10 02:16:43.119502
9142d8a7-4a0f-4bf1-baf5-f8d168bfbd16	fdfff5f2-ca91-49e6-8a26-98fece18f35e	f8fd2e10-1a96-44ae-9898-50fa57d8dcf9	1.00	12000.00	\N	2024-07-10 02:19:11.452848	2024-07-10 02:19:11.452848
\.


--
-- Data for Name: transaction_payments; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.transaction_payments (id, transaction_id, payment_type, reference_number, total, total_paid, created_at, updated_at) FROM stdin;
53259ad8-5778-4a56-b570-3137eaa48b65	85ee56d5-9a8f-42e7-8bdc-3aa1512ec5d0	CASH	\N	300000.00	300000.00	2024-06-07 15:51:32.821042	2024-06-07 15:51:32.821042
eb77b8bb-ff80-409c-b600-c8f8b08e390a	d7e6f3ac-a4ba-4e14-901e-9e63c67b1458	CASH	\N	140000.00	150000.00	2024-06-26 04:30:11.019836	2024-06-26 04:30:11.019836
212aea76-0e67-4c60-ac60-537ea86ab16a	67acec7e-bb3b-499a-92e7-4d23f6f7eb31	CASH	\N	140000.00	150000.00	2024-06-26 04:30:53.726084	2024-06-26 04:30:53.726084
0a125c90-d19f-4585-8206-f1935e8e1dec	458047e3-4000-487e-964d-4be12d701557	CASH	\N	140000.00	150000.00	2024-06-26 04:47:50.03289	2024-06-26 04:47:50.03289
8550d455-0669-443c-b6bc-4b7af6712872	26bfac7e-9b6c-4bb4-b59e-e25870e83ee9	CASH	\N	140000.00	150000.00	2024-06-26 04:49:48.489727	2024-06-26 04:49:48.489727
372f2f2b-ff53-43f6-ba29-54fa196a0a0d	9d4d8b64-22bc-4d67-afa4-d2719ce861bd	CASH	\N	140000.00	150000.00	2024-06-26 05:01:03.284845	2024-06-26 05:01:03.284845
4da368c6-e2a5-4c1d-a2f6-cb79aa6a8506	f5e5036c-0a13-481c-ab93-1ccf2cf51ca4	CASH	\N	140000.00	150000.00	2024-06-26 05:01:52.883679	2024-06-26 05:01:52.883679
0ddc97dc-9c22-4f58-ab80-ec03f8e95a60	7efd1898-66be-45ae-bf09-85b4478078df	CASH	\N	140000.00	150000.00	2024-06-26 05:03:22.477153	2024-06-26 05:03:22.477153
7d9ef5a9-fa7c-42b6-aa6d-d1465b660bf3	434c6873-eca8-41fe-a3ea-f28ef64e03f0	CASH	\N	140000.00	150000.00	2024-06-26 05:03:37.755956	2024-06-26 05:03:37.755956
5e5142ef-beac-45d1-8f78-fa5f8774cb74	073c3ef1-72f9-4cbf-bc04-5a19b378ec22	BCA_TRANSFER	465465465465465465	140000.00	140000.00	2024-06-26 05:04:06.607771	2024-06-26 05:04:06.607771
cd3d4c7b-106e-49c3-b1c3-58f83aa36bdf	44a8d583-eb5d-4a32-a8b8-0fb084a8dfde	CASH	\N	140000.00	150000.00	2024-06-26 06:54:51.896965	2024-06-26 06:54:51.896965
91055dae-4227-4160-91ee-7fbc1ae8a7f9	b23eb768-6987-4af1-969a-9d396ddc945a	CASH	\N	240000.00	250000.00	2024-06-27 09:28:28.521705	2024-06-27 09:28:28.521705
502905df-5a0b-467c-8223-411bc8f083f6	4bec0189-8da7-4f2f-98f0-9ef92ad00a36	CASH	\N	100000.00	110000.00	2024-06-27 09:29:55.839729	2024-06-27 09:29:55.839729
13870414-04a5-4cf7-b71a-bf6e11b3cc8b	5c8140ea-4462-4f85-8acb-0d1998bcb56e	CASH	\N	200000.00	210000.00	2024-06-27 09:31:20.321544	2024-06-27 09:31:20.321544
01179426-7b7d-4721-9c47-41efff4202d8	a7364625-a12f-4d50-8fa2-32ec8df06b26	CASH	\N	240000.00	250000.00	2024-06-27 09:33:06.748677	2024-06-27 09:33:06.748677
ad08814d-281f-4e07-af03-657eac99251d	73e1140c-dfb6-4d2f-9474-25eedf7fa425	CASH	\N	300000.00	300000.00	2024-06-28 00:33:48.927079	2024-06-28 00:33:48.927079
3ffebce9-5b49-4dfe-bf16-ec2d78b5b3da	3a70f48c-49eb-4100-a1ff-0781bb03ea12	CASH	\N	400000.00	400000.00	2024-06-28 02:35:07.353111	2024-06-28 02:35:07.353111
fc1c00f6-387d-4eb5-b8b5-9dc43ed8cebc	d48b7681-531a-4aae-a1b0-6b57d366d187	CASH	\N	300000.00	310000.00	2024-06-29 09:58:01.411772	2024-06-29 09:58:01.411772
e4c6467c-a252-48d9-b0ee-011d825beab0	cb683226-9b91-46ad-a13a-c7c38dedc290	CASH	\N	100000.00	100000.00	2024-07-01 03:08:16.939266	2024-07-01 03:08:16.939266
1f49059a-4bd2-48f0-953f-0f6407a50e84	643d3f68-95a8-461c-ae84-ed51a88c6dfb	CASH	\N	200000.00	200000.00	2024-07-01 09:20:59.878833	2024-07-01 09:20:59.878833
b7ace8ad-61d5-43dc-a4ea-e65180df1f78	b20d6bee-8020-423a-8b92-5c794edaa2eb	CASH	\N	590000.00	600000.00	2024-07-08 15:54:59.916685	2024-07-08 15:54:59.916685
e2a218f3-cc56-415c-a6ab-07b2b0cd9ba2	95e5f612-d352-48eb-8338-776e44d26501	CASH	\N	10000.00	10000.00	2024-07-08 16:02:26.852374	2024-07-08 16:02:26.852374
3c7e9ff0-5fbb-4868-a4ae-f7ddc207ef68	382a3371-4fe2-41ae-b02b-091cd67e78c6	CASH	\N	10000.00	10000.00	2024-07-08 16:03:43.147593	2024-07-08 16:03:43.147593
8dcb9100-5af7-4b47-ac15-bd93e8201745	87d66a03-3756-4afb-a84b-c7f79ffaa788	CASH	\N	12000.00	12000.00	2024-07-10 02:16:43.11707	2024-07-10 02:16:43.11707
f9424c5f-74b9-4ad3-90e6-d9a4dd4ec473	fdfff5f2-ca91-49e6-8a26-98fece18f35e	CASH	\N	12000.00	12000.00	2024-07-10 02:19:11.451406	2024-07-10 02:19:11.451406
\.


--
-- Data for Name: transactions; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.transactions (id, cashier_session_id, status, total, payment_at, created_at, updated_at) FROM stdin;
85ee56d5-9a8f-42e7-8bdc-3aa1512ec5d0	b0ee624b-040c-41f9-8d21-945db1adc44f	PAID	300000.00	2024-06-07 15:51:32.744811	2024-06-07 15:51:32.816892	2024-06-07 15:51:32.816892
d7e6f3ac-a4ba-4e14-901e-9e63c67b1458	d737f625-5234-4008-a9ef-49381c8f5a13	PAID	140000.00	2024-06-26 04:30:10.949099	2024-06-26 04:30:11.015465	2024-06-26 04:30:11.015465
67acec7e-bb3b-499a-92e7-4d23f6f7eb31	d737f625-5234-4008-a9ef-49381c8f5a13	PAID	140000.00	2024-06-26 04:30:53.663981	2024-06-26 04:30:53.724377	2024-06-26 04:30:53.724377
458047e3-4000-487e-964d-4be12d701557	d737f625-5234-4008-a9ef-49381c8f5a13	PAID	140000.00	2024-06-26 04:47:49.97107	2024-06-26 04:47:50.031795	2024-06-26 04:47:50.031795
26bfac7e-9b6c-4bb4-b59e-e25870e83ee9	d737f625-5234-4008-a9ef-49381c8f5a13	PAID	140000.00	2024-06-26 04:49:48.430083	2024-06-26 04:49:48.488987	2024-06-26 04:49:48.488987
9d4d8b64-22bc-4d67-afa4-d2719ce861bd	d737f625-5234-4008-a9ef-49381c8f5a13	PAID	140000.00	2024-06-26 05:01:03.204167	2024-06-26 05:01:03.281439	2024-06-26 05:01:03.281439
f5e5036c-0a13-481c-ab93-1ccf2cf51ca4	d737f625-5234-4008-a9ef-49381c8f5a13	PAID	140000.00	2024-06-26 05:01:52.808353	2024-06-26 05:01:52.881582	2024-06-26 05:01:52.881582
7efd1898-66be-45ae-bf09-85b4478078df	d737f625-5234-4008-a9ef-49381c8f5a13	PAID	140000.00	2024-06-26 05:03:22.414491	2024-06-26 05:03:22.475582	2024-06-26 05:03:22.475582
434c6873-eca8-41fe-a3ea-f28ef64e03f0	d737f625-5234-4008-a9ef-49381c8f5a13	PAID	140000.00	2024-06-26 05:03:37.696439	2024-06-26 05:03:37.755297	2024-06-26 05:03:37.755297
073c3ef1-72f9-4cbf-bc04-5a19b378ec22	d737f625-5234-4008-a9ef-49381c8f5a13	PAID	140000.00	2024-06-26 05:04:06.546923	2024-06-26 05:04:06.607165	2024-06-26 05:04:06.607165
44a8d583-eb5d-4a32-a8b8-0fb084a8dfde	d737f625-5234-4008-a9ef-49381c8f5a13	PAID	140000.00	2024-06-26 06:54:51.828149	2024-06-26 06:54:51.894613	2024-06-26 06:54:51.894613
b23eb768-6987-4af1-969a-9d396ddc945a	3705dfe3-4097-4c0b-b888-895436467d09	PAID	240000.00	2024-06-27 09:28:28.458914	2024-06-27 09:28:28.520076	2024-06-27 09:28:28.520076
4bec0189-8da7-4f2f-98f0-9ef92ad00a36	f24c78ab-23cf-4da7-9350-3bfe1716e3cf	PAID	100000.00	2024-06-27 09:29:55.780084	2024-06-27 09:29:55.838947	2024-06-27 09:29:55.838947
5c8140ea-4462-4f85-8acb-0d1998bcb56e	f24c78ab-23cf-4da7-9350-3bfe1716e3cf	PAID	200000.00	2024-06-27 09:31:20.260162	2024-06-27 09:31:20.320175	2024-06-27 09:31:20.320175
a7364625-a12f-4d50-8fa2-32ec8df06b26	f24c78ab-23cf-4da7-9350-3bfe1716e3cf	PAID	240000.00	2024-06-27 09:33:06.687542	2024-06-27 09:33:06.747795	2024-06-27 09:33:06.747795
73e1140c-dfb6-4d2f-9474-25eedf7fa425	fc192a4e-4a31-42c6-bf2e-ba06b5f54509	PAID	300000.00	2024-06-28 00:33:48.86769	2024-06-28 00:33:48.926037	2024-06-28 00:33:48.926037
3a70f48c-49eb-4100-a1ff-0781bb03ea12	e2b2a92b-4684-453e-b224-329185eed415	PAID	400000.00	2024-06-28 02:35:07.292217	2024-06-28 02:35:07.352198	2024-06-28 02:35:07.352198
d48b7681-531a-4aae-a1b0-6b57d366d187	7da6cd43-122c-462c-bf01-3cf4f6dfcede	PAID	300000.00	2024-06-29 09:58:01.345857	2024-06-29 09:58:01.409069	2024-06-29 09:58:01.409069
cb683226-9b91-46ad-a13a-c7c38dedc290	0a3376e6-9176-4850-b589-9ad812b03985	PAID	100000.00	2024-07-01 03:08:16.872371	2024-07-01 03:08:16.937468	2024-07-01 03:08:16.937468
643d3f68-95a8-461c-ae84-ed51a88c6dfb	afd5d5ab-4b0f-4789-a460-a6c8ee228c98	PAID	200000.00	2024-07-01 09:20:59.807977	2024-07-01 09:20:59.875436	2024-07-01 09:20:59.875436
b20d6bee-8020-423a-8b92-5c794edaa2eb	62ca2919-3d31-45a7-a42a-1dbdfd6f43da	PAID	590000.00	2024-07-08 15:54:59.856131	2024-07-08 15:54:59.915198	2024-07-08 15:54:59.915198
95e5f612-d352-48eb-8338-776e44d26501	ba5bf7c4-8b51-49ad-a3cc-2d2d70d791df	PAID	10000.00	2024-07-08 16:02:26.791709	2024-07-08 16:02:26.851259	2024-07-08 16:02:26.851259
382a3371-4fe2-41ae-b02b-091cd67e78c6	4d55ec36-e9fc-4b9e-99a1-6aea4ac3afd3	PAID	10000.00	2024-07-08 16:03:43.087146	2024-07-08 16:03:43.146749	2024-07-08 16:03:43.146749
87d66a03-3756-4afb-a84b-c7f79ffaa788	a4ede830-480d-438c-971c-4a0939c8345e	PAID	12000.00	2024-07-10 02:16:43.050277	2024-07-10 02:16:43.11437	2024-07-10 02:16:43.11437
fdfff5f2-ca91-49e6-8a26-98fece18f35e	bfc6cfb1-7113-4148-adbd-0c397b9adaff	PAID	12000.00	2024-07-10 02:19:11.389566	2024-07-10 02:19:11.449625	2024-07-10 02:19:11.449625
\.


--
-- Data for Name: units; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.units (id, name, description, created_at, updated_at) FROM stdin;
6745548c-ea48-4db8-b7d9-ed2cf1175ade	Kaleng	\N	2024-06-06 09:29:38.569559	2024-06-06 09:29:38.569559
3867eb2b-8905-402c-bce3-c5953262ec03	Dus	\N	2024-06-06 09:29:38.56956	2024-06-06 09:29:38.56956
fbc61727-4eae-4f00-9ebf-2ff2c2f8156b	Pcs	\N	2024-06-27 08:51:43.007738	2024-06-27 08:51:43.007738
\.


--
-- Data for Name: user_access_tokens; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.user_access_tokens (id, user_id, revoked, created_at, updated_at, expired_at) FROM stdin;
be0952f96ae04d488c457cc5a371b433531b1c9c25821c3b1a3142f886fdf1e848918e9ee1cddc7e                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-06-06 09:29:43.953518	2024-06-06 09:29:43.953518	2024-06-07 09:29:43.95317
c6f46ced853aa20f83e6962881573e59ecd9eae349f28a052369263a874b0ab9f71e9c990b5b6ae0                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-06-06 14:22:04.810147	2024-06-06 14:22:04.810147	2024-06-07 14:22:04.810124
c187a2f56780cd7690b9cdc1c552b09709da7ece8747118532975c19155dcb88120e5350904b55b6                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-06-07 02:16:28.73016	2024-06-07 02:16:28.73016	2024-06-08 02:16:28.730139
ed70ab8d453b82df9e1fc36399d8c2296418df82f18c0fafa70264e85441589dfe829d7c7a917049                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-06-07 03:18:03.694519	2024-06-07 03:18:03.694519	2024-06-08 03:18:03.693975
86d2c9ef8d72d0d3e9b31d3fa2a60b9cdea6daa28d10a5294f9c184623e2bd39719e731ea1288105                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-06-07 03:44:47.411681	2024-06-07 03:44:47.411681	2024-06-08 03:44:47.41166
2950ee88602a6442d0f249677b4de6ad176f09a98d25d13c057881136b1275d9246ec0220cd51d80                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-06-07 09:40:48.088868	2024-06-07 09:40:48.088868	2024-06-08 09:40:48.088846
6ea6836dffe8cef4f09f3f7e36081adb408feab66f8db3eb33d39d62343618852f3deae586f33239                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-06-07 15:29:55.556177	2024-06-07 15:29:55.556177	2024-06-08 15:29:55.556132
589d9790d08a02093f7afc7480fb7a387dfd9a2582e25434aad868ec0a17bb762a9cbb274bd87ab0                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-06-08 05:36:33.770651	2024-06-08 05:36:33.770651	2024-06-09 05:36:33.770627
e0284245a188f657f7f437ff6ab11b167e26fc41b02f1a84ce9057289ba25e4bd06a9643e2f08ea8                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-06-08 10:32:56.428516	2024-06-08 10:32:56.428516	2024-06-09 10:32:56.428495
37c0893aa81edb0188c7bec10eaa16e2ca761eaa4eda924f2a661199d78eb0755bc3b42e4c6da26c                    	fbfcbc34-77b9-4901-82f1-e5fa78d5aa48	f	2024-06-11 02:14:49.968673	2024-06-11 02:14:49.968673	2024-06-12 02:14:49.968661
b2acadb6652bdc5d3d29e726fd6055b94fb24cec6b22ee136bbf75ddd57f5c1d5922512df84b5905                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-06-11 02:15:26.195209	2024-06-11 02:15:26.195209	2024-06-12 02:15:26.195188
0b29d62c845407d3bac0b61a78fd91944168f04261004449c14f2d106b771d602b14140be5a770f4                    	fbfcbc34-77b9-4901-82f1-e5fa78d5aa48	f	2024-06-11 02:52:14.589902	2024-06-11 02:52:14.589902	2024-06-12 02:52:14.58924
ee4a56170b12c78453f8b8e8e4fae96049679ff21b486d3615c9d3667c034a34bf18d04757b6a7e7                    	fbfcbc34-77b9-4901-82f1-e5fa78d5aa48	f	2024-06-11 02:56:07.743356	2024-06-11 02:56:07.743356	2024-06-12 02:56:07.743342
887a35bfc197f44b70e782065cd70d86a9d76241b216cdfc3a9f2eb9a237a2bacf8f6e17f20c64de                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-06-11 03:16:29.100998	2024-06-11 03:16:29.100998	2024-06-12 03:16:29.100978
57e570ca284bae1bb83ce3d964f117907981f1970fa8b7c1b4ae8d4e633bb9305f51ab35ea4df4f0                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-06-12 01:37:59.039966	2024-06-12 01:37:59.039966	2024-06-13 01:37:59.039947
87d9282c8024865f7976d2652afb3b62bc59d0d99d55105a8756b9183152a912342a2bf8fd1f3e2b                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-06-12 03:42:35.512167	2024-06-12 03:42:35.512167	2024-06-13 03:42:35.512146
c919572abe7257f04aa39d8d3d43ccec950289cac96d56d9e9b8066ee0fb540891598091df3d7501                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-06-14 07:09:24.889564	2024-06-14 07:09:24.889564	2024-06-15 07:09:24.889525
e7f582ff8f69efb01a6299ac8765e8e3c004c5df6001cc81feeb55cdea631a5eb137698c36fc570d                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-06-14 07:17:37.174145	2024-06-14 07:17:37.174145	2024-06-15 07:17:37.174125
fb9d87806549bca8081657ae947a98c56001ec503aeef6c6e96553ec8759d3fcfb5a797115453846                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-06-18 06:22:49.839812	2024-06-18 06:22:49.839812	2024-06-19 06:22:49.839786
c60887eca8ace0c50d3fb1c17a4170130d93b465ec13f3aaa75fe82a2db85dddd13a47be90abc245                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-06-18 06:49:10.097889	2024-06-18 06:49:10.097889	2024-06-19 06:49:10.097858
1bdc068c96a8845609961f8c1d8e7817e2ab7d0b6ecf23f6994e94d5a77d1f8a984710f05ca1d0ac                    	fbfcbc34-77b9-4901-82f1-e5fa78d5aa48	f	2024-06-18 06:51:40.155241	2024-06-18 06:51:40.155241	2024-06-19 06:51:40.155231
8460f9e80257f010424f0eff5eef0e5140310ad2c9feb307d0e091e5edddddb57004b7e1509a6e20                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-06-18 06:53:39.342337	2024-06-18 06:53:39.342337	2024-06-19 06:53:39.342311
c43f4bf3918c7df9df45f07ec75f8234e672809dff949542595e523f376d0c4727b69baa3413d2ea                    	fbfcbc34-77b9-4901-82f1-e5fa78d5aa48	f	2024-06-18 06:53:56.14391	2024-06-18 06:53:56.14391	2024-06-19 06:53:56.143878
bd828e18f29034a4fea69d0a9436793375f854d56a9167a31fbeee1592e1977c162ea3c88c8f5e10                    	fbfcbc34-77b9-4901-82f1-e5fa78d5aa48	f	2024-06-18 06:56:56.427965	2024-06-18 06:56:56.427965	2024-06-19 06:56:56.427955
9e5f8abd88ead348d367e9d881834da75fe9edaa29cb33ac9afeac632478c8775344ed9e1c4ca12f                    	fbfcbc34-77b9-4901-82f1-e5fa78d5aa48	f	2024-06-18 07:05:46.484634	2024-06-18 07:05:46.484634	2024-06-19 07:05:46.484604
cbffeef4bb15462a2980dba3b431f591cf231f1b1797cf78c415a317cf15b69b8959489a1531998b                    	fbfcbc34-77b9-4901-82f1-e5fa78d5aa48	f	2024-06-18 07:07:19.360505	2024-06-18 07:07:19.360505	2024-06-19 07:07:19.360495
4a866c71c0030ce865f8ad6d0bc0504bcb67a4545417b22dc637714f25770db5f051ac20ef2e3586                    	fbfcbc34-77b9-4901-82f1-e5fa78d5aa48	f	2024-06-18 07:19:12.114182	2024-06-18 07:19:12.114182	2024-06-19 07:19:12.11417
c546180ddf427699c3f19f57f6832e36856aad6eec69ef40ff725724c5f01e97a750f8517c75269d                    	fbfcbc34-77b9-4901-82f1-e5fa78d5aa48	f	2024-06-19 07:33:47.626837	2024-06-19 07:33:47.626837	2024-06-20 07:33:47.626809
844b69282d469bcc856158fba8ead388b9f595cb9d7c411bc3208e9eaf4d8dc0ddf2d40de03776ba                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-06-20 02:21:41.392488	2024-06-20 02:21:41.392488	2024-06-21 02:21:41.392472
20590f0848d8215e3379afb04d019de91dcbbca59d3961d2ce4bca1b5f1b5dee24508083a1b1a0f7                    	fbfcbc34-77b9-4901-82f1-e5fa78d5aa48	f	2024-06-20 08:31:13.835746	2024-06-20 08:31:13.835746	2024-06-21 08:31:13.834443
b36630657642f0f8c88194d5cbf48df258dd65c5f0e2d08b2f522c2df4157ba9b7706bf0ea453839                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-06-21 06:07:55.43965	2024-06-21 06:07:55.43965	2024-06-22 06:07:55.439632
1023af5b79de3d78d552423e7440fc7512a3b9f9044ae5488a0b819f7d2e3fd45d747a6b01beaf60                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-06-21 08:44:27.654192	2024-06-21 08:44:27.654192	2024-06-22 08:44:27.65417
ed76a60eb2b736806958cf2bb6ef8df71cf94f9d57997369276480d724a7d674c0a700f15ea2c948                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-06-21 09:31:17.076217	2024-06-21 09:31:17.076217	2024-06-22 09:31:17.07619
4fa39c7e76411ac925dfc6ff73b82d6d7f9d32c43868244392fbcaed08a29da142f55c82687347c6                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-06-21 09:31:32.191396	2024-06-21 09:31:32.191396	2024-06-22 09:31:32.191371
dda664d5c240b00ce484c8873c26fe4dc418ccba68fa9508ec269c52e16fd18dafdb0e2a3adfd27b                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-06-21 10:02:27.531443	2024-06-21 10:02:27.531443	2024-06-22 10:02:27.531416
2d2d7250f9a62248f586a0dc6d1483bed95ddbbdcdba7ea09e0d0fdfeec7741fd5ba4caf8bb67403                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-06-21 10:02:56.711092	2024-06-21 10:02:56.711092	2024-06-22 10:02:56.711065
f5107d1ef6080148a5170d7e52be984e45d109fe170ffb1d88b169fd5223533179bfecad9ca2659b                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-06-26 03:06:42.21514	2024-06-26 03:06:42.21514	2024-06-27 03:06:42.215111
4fff4c604af649b2554b8ddc4ba558b99ba1d84eeaf0ac3630e6c6ecef3d53169f8aee326a30bfe7                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-06-26 04:27:17.592481	2024-06-26 04:27:17.592481	2024-06-27 04:27:17.592453
fc9bd9a3483144615869433b852186fa8525f72b7b27d79cc3a337a3153f9c13e6f23debbed19769                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-06-26 04:32:58.889026	2024-06-26 04:32:58.889026	2024-06-27 04:32:58.889008
54e8ebc2f74915a4a6fb3fe62c736a7a0c73f009581a3e1b3b1960edb23483b8a4af200b82ddb104                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-06-27 06:23:47.239989	2024-06-27 06:23:47.239989	2024-06-28 06:23:47.239965
59d434de78dafc553f2c42237969e6479e264fba1db0a90cb9c65abd30fa4d12cd2971f69344eec6                    	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	f	2024-06-27 06:28:31.719321	2024-06-27 06:28:31.719321	2024-06-28 06:28:31.719295
2592445f4384c746594f1348e9734d6801f0adc9cb125a76525f08cecafe6ff40e346a000f4f9388                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-06-27 06:29:01.082063	2024-06-27 06:29:01.082063	2024-06-28 06:29:01.082036
a3874331533b870acccae405aee4942891793721ef9e6a6d8b398d31fb23233d7daa0b7f4cc5a200                    	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	f	2024-06-27 06:29:15.855322	2024-06-27 06:29:15.855322	2024-06-28 06:29:15.855299
63fd78ae1fce3663043e40fbd7f1c2ae39220f62564e6f4447bf5097b1ce762e8023871b60582ed9                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-06-27 06:32:39.440154	2024-06-27 06:32:39.440154	2024-06-28 06:32:39.440136
31d2aa95d5d12fbed96164ce7c2e22d03a1c4c046b77c14d914e1bef600e61dbe6c1e0f41572ceda                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-06-27 06:33:43.598519	2024-06-27 06:33:43.598519	2024-06-28 06:33:43.598496
9e98d7a00227e0d2280c6e94fca4fb0e532aa906bcd5bc3be6c34a94ae47d185dda13c37fc0b32ef                    	68ed7124-fa08-4720-b741-9fe4fa697c21	f	2024-06-27 06:35:26.259562	2024-06-27 06:35:26.259562	2024-06-28 06:35:26.259535
1c5ee9017cd9c4dee611d74a58b94fd77fadfcc462488b15c14a50d036461326c6f1e0b0446b28c9                    	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	f	2024-06-27 06:42:08.76863	2024-06-27 06:42:08.76863	2024-06-28 06:42:08.768604
87549bd12581e3e4dc088f7282641eeee6d62858e2b7172e23e3cb4ae0c495844e1f02171844eaca                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-06-27 07:09:38.470134	2024-06-27 07:09:38.470134	2024-06-28 07:09:38.470115
1f0c5e8123b19e2fcafe10a7663617efaaaefdb49c557c049d10c2b526d5fb30091f94d4b62f8eb0                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-06-27 07:29:54.099567	2024-06-27 07:29:54.099567	2024-06-28 07:29:54.099542
3a95cdc39deb7a6e2010f737a35eb3fcaec1df077121af4b037dc0d220105fa966c97f62b9e8f4de                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-06-27 07:33:38.75008	2024-06-27 07:33:38.75008	2024-06-28 07:33:38.750059
f724cd62cea3ecd885c07ce418b2b0c30fee791a98c34237d91cbc3d3651de66bde987edcc66ca7f                    	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	f	2024-06-27 07:34:06.130352	2024-06-27 07:34:06.130352	2024-06-28 07:34:06.130326
68fdde8cd9ea83124e758fb35f8272f3474492b730da5abf91b1780fe7e3ec556be540099f3e80b3                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-06-27 07:36:22.734544	2024-06-27 07:36:22.734544	2024-06-28 07:36:22.734521
0d9d255419eacdf288f7ccb6917888b5754b96b987e3dc15166a47567e59bc6571443b1d7fcf8142                    	68ed7124-fa08-4720-b741-9fe4fa697c21	f	2024-06-27 07:36:52.17738	2024-06-27 07:36:52.17738	2024-06-28 07:36:52.177344
819f532b7669419b9dbc5e7869a4fe4c1521c5e3212593114040848cd37c72143fada031b2b35ecf                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-06-27 07:40:41.938307	2024-06-27 07:40:41.938307	2024-06-28 07:40:41.937771
eb37ab7ffe327ab05e84a7299deb6abaa6f03e4c7cafdf47ba4c7fab2cf222dc5d59c37ee18f23c4                    	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	f	2024-06-27 07:41:04.684229	2024-06-27 07:41:04.684229	2024-06-28 07:41:04.684124
1f1b3f1cdce47301ca1c4218b02e652e4b368bdda6eb537246c113004fd42f152d6e9ed8c2fbaa79                    	fbfcbc34-77b9-4901-82f1-e5fa78d5aa48	f	2024-06-27 08:41:09.784377	2024-06-27 08:41:09.784377	2024-06-28 08:41:09.784354
02cc4f6103b476d4853c077a87da0dba703e7c6255f5961986aec3420c74ea896ed9a409e9a7e90e                    	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	f	2024-06-27 08:41:17.274739	2024-06-27 08:41:17.274739	2024-06-28 08:41:17.274714
1dac0a3f1f4c7985c4c2f528ca61809d6ec29be5b43847d1c6ef991b5d40b05687996443a7abaf2d                    	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	f	2024-06-27 08:42:37.912116	2024-06-27 08:42:37.912116	2024-06-28 08:42:37.91209
c4423cbfa5e2abde71bdb9850892c63040048348ffe7174757334ddaa4866fb669d1e7b99ae2af5d                    	68ed7124-fa08-4720-b741-9fe4fa697c21	f	2024-06-27 08:43:04.89413	2024-06-27 08:43:04.89413	2024-06-28 08:43:04.894103
c6532d1b01bb7e9afb2a2f421225df3c7a8f50633d41b1f4201ebd3c60933e5f3be01d89e61f9a69                    	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	f	2024-06-27 08:46:24.626337	2024-06-27 08:46:24.626337	2024-06-28 08:46:24.626312
31374db156271b2d50ce610321690478755684fed90237d06f40ec83a13f0079fefd408d9566d15c                    	fbfcbc34-77b9-4901-82f1-e5fa78d5aa48	f	2024-06-27 09:24:30.277066	2024-06-27 09:24:30.277066	2024-06-28 09:24:30.277055
b3d542bbecd596b4cb5da3c6a9580c04237fa29118df4f8a08e5d5c1ad8a1277a8078a4a33dfb8d2                    	68ed7124-fa08-4720-b741-9fe4fa697c21	f	2024-06-27 09:27:17.516705	2024-06-27 09:27:17.516705	2024-06-28 09:27:17.516686
34ef1202177d70322879847352309e8a8db98b6f703b8d800316abeacd576d3be5eadd3ee8d9b443                    	68ed7124-fa08-4720-b741-9fe4fa697c21	f	2024-06-27 09:29:27.740197	2024-06-27 09:29:27.740197	2024-06-28 09:29:27.740173
e00a35f330012a434a517516e6069ba7cf4a49f8c33e0b220ac01dd21bca0a4bfad587932275d058                    	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	f	2024-06-28 00:16:34.582895	2024-06-28 00:16:34.582895	2024-06-29 00:16:34.582871
100de45881d489f63af6b4f6b7bd43c7528a0c3bd877006b628c53d629e07af8ce5e63f2ebea1e8a                    	68ed7124-fa08-4720-b741-9fe4fa697c21	f	2024-06-28 00:32:33.533834	2024-06-28 00:32:33.533834	2024-06-29 00:32:33.533816
596f7719b83ee7c1f81d3f46678959ec9df23916edd0aee1b94a23bf7943a1d96d902fcf4776ca5c                    	68ed7124-fa08-4720-b741-9fe4fa697c21	f	2024-06-28 02:34:44.975152	2024-06-28 02:34:44.975152	2024-06-29 02:34:44.97513
ffc720e8c724bef127751caaa9c4635d18a724fb954ccd720edddaaa2b07df49348439619303fa68                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-06-28 02:43:50.10885	2024-06-28 02:43:50.10885	2024-06-29 02:43:50.108828
20ede2de5b8bb408496dc4b8b5d9b85c1bf8dd913a55ce58fffa2b33c23762efed3e208951d76340                    	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	f	2024-06-28 02:53:08.543876	2024-06-28 02:53:08.543876	2024-06-29 02:53:08.543854
9f24bedc46294bae42112df762f7ebea42b8c7c4f6d0d6429fc7c9705a2348e7adf7a69928fc9bd7                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-06-28 07:37:34.480587	2024-06-28 07:37:34.480587	2024-06-29 07:37:34.480563
febcac9a0d50fd3bc3ee315841bfe6710c831bcc54657db7bbad298b779d62ea734670db480e5748                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-06-29 09:04:48.331706	2024-06-29 09:04:48.331706	2024-06-30 09:04:48.33124
77e829dbc0691624de3484cbd4514b8e4a2835efd067624245fd4c064eb89b8f722b95ef1b8db9ca                    	16638b6a-5ca3-4d1f-ab0f-ac4ce4ce30a0	f	2024-06-29 09:08:12.151652	2024-06-29 09:08:12.151652	2024-06-30 09:08:12.151635
060697f60aa9ccb65d4cf2929c4402db88c4b160f336c3e30a6992b8d706db1166634cb8c60f8a41                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-06-29 09:08:58.314226	2024-06-29 09:08:58.314226	2024-06-30 09:08:58.314196
68f43263d353cbfec74b77e8ead61e324dc04d3d6283b686a92b85467b8d2006e102d5101dee5b7f                    	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	f	2024-06-29 09:26:39.465832	2024-06-29 09:26:39.465832	2024-06-30 09:26:39.465807
aa8105be30d6689118a7f3faf1e1aa5006d7214b0e2b10dbfae1523fcaa941982ec9a3b9529db364                    	68ed7124-fa08-4720-b741-9fe4fa697c21	f	2024-06-29 09:35:57.824354	2024-06-29 09:35:57.824354	2024-06-30 09:35:57.824332
9c31ad8042aaa0bab13b9d622be6a46e394a20f0dc2e3ea7bd09bf8fd56469394b57c5d8c3e11525                    	16638b6a-5ca3-4d1f-ab0f-ac4ce4ce30a0	f	2024-06-29 09:36:19.086343	2024-06-29 09:36:19.086343	2024-06-30 09:36:19.086319
69d2f9cb9b64a7fc8b359a7e3f146986340d94ee8508ea4987705f8cf23e50743085a700daf8ae66                    	68ed7124-fa08-4720-b741-9fe4fa697c21	f	2024-06-29 09:36:38.282694	2024-06-29 09:36:38.282694	2024-06-30 09:36:38.282671
a4c5f1ef1070b753a9a74ecc38e3b98c2529314c436e9b7753255e9dcfd32504cc3020ba6400ed65                    	16638b6a-5ca3-4d1f-ab0f-ac4ce4ce30a0	f	2024-06-29 09:45:58.752226	2024-06-29 09:45:58.752226	2024-06-30 09:45:58.752203
afe652ad65124dfed53fec378c8edeb7d53afb963d791a37edc130e8d57767c2ea98814443898266                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-06-29 10:00:02.077239	2024-06-29 10:00:02.077239	2024-06-30 10:00:02.07721
12575c649de52db5b7932726d08ddf542b731013c1e9bd56368b10f01e2a60c4ef3450012e67f47d                    	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	f	2024-06-29 10:02:26.493889	2024-06-29 10:02:26.493889	2024-06-30 10:02:26.493869
fa7feb24daa521b3c4fc75b91d349e4bcb7f5a4f9ce9850852f61dff4f38567d5d1bc07ac8e64247                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-06-29 15:33:31.33953	2024-06-29 15:33:31.33953	2024-06-30 15:33:31.339502
b1edf23c7183e6b042bfe528b9cb20edd59004a9c88fabb2066b83cc05355829b227e79daaedfafd                    	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	f	2024-06-29 15:34:15.182655	2024-06-29 15:34:15.182655	2024-06-30 15:34:15.182628
d381f7cb00e40a68f40f57f3f62e98e8f8dec2889c2332fce86ca7e787f2e58d74ae668de79f865b                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-06-30 10:53:28.240321	2024-06-30 10:53:28.240321	2024-07-01 10:53:28.240295
656ddfb09b32af075fbbcc660bd11a9a77b787cb95ae2fbdb0eaea49d8462c5cb2ce1fadd7794147                    	33392eb4-0f87-43c6-9893-9c014fe6d561	f	2024-06-30 10:54:04.306638	2024-06-30 10:54:04.306638	2024-07-01 10:54:04.306618
65b4495d493557a2eaa10dfbaa7ad901e565f86dcad661601b5276ea1243850115bb7b46071f6f54                    	fbfcbc34-77b9-4901-82f1-e5fa78d5aa48	f	2024-06-30 10:54:32.020301	2024-06-30 10:54:32.020301	2024-07-01 10:54:32.020282
3c579f0de343b6bbde35639ee0ebca9aa3a9b1a1c4bac96d9b36475d77b9a5e45b292bdf319af1ab                    	33392eb4-0f87-43c6-9893-9c014fe6d561	f	2024-06-30 10:54:39.397963	2024-06-30 10:54:39.397963	2024-07-01 10:54:39.397945
ddc277fe50ab91c3ced49f1081cce3f650a89e83f57875c398cd9c01fb360e3ff8513298332b6d13                    	fbfcbc34-77b9-4901-82f1-e5fa78d5aa48	f	2024-06-30 15:44:53.978385	2024-06-30 15:44:53.978385	2024-07-01 15:44:53.978332
f062f2210e963320b730866c0565a62212fc17306eb8b0bc1a68b64bda1d541b03ecfb92314ee350                    	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	f	2024-06-30 15:45:59.5147	2024-06-30 15:45:59.5147	2024-07-01 15:45:59.514681
6cc51da176ff9159b0328918a25ecae5224d9784f19d0ef2ecb3fb58f58121545014fa6d769eeb77                    	68ed7124-fa08-4720-b741-9fe4fa697c21	f	2024-07-01 02:54:09.501048	2024-07-01 02:54:09.501048	2024-07-02 02:54:09.501028
976521c1fa5a1c194bb2e73ab9701eca1c6f4ce1aaa8cafcf6f7833b360ae1b271a9c36e1f7f03ad                    	68ed7124-fa08-4720-b741-9fe4fa697c21	f	2024-07-01 03:42:32.364053	2024-07-01 03:42:32.364053	2024-07-02 03:42:32.364037
f64ea71f4b356c8804649d06124d3d0f67a8b5343f6b8579b3d6bfce5d9e1f58b87f276a669f6498                    	fbfcbc34-77b9-4901-82f1-e5fa78d5aa48	f	2024-07-01 06:34:44.000292	2024-07-01 06:34:44.000292	2024-07-02 06:34:44.000244
6ca71c10bb2ae2e03c54959ea7f53aad787ba88d91a3ddcbb47b1c820399c331891a84d4ab9a3c20                    	fbfcbc34-77b9-4901-82f1-e5fa78d5aa48	f	2024-07-01 06:41:33.609486	2024-07-01 06:41:33.609486	2024-07-02 06:41:33.609465
ddbd25e5f0dcd5459791053294ad821c86069b8eef4e398cdffd1d48e8658844be9cf28fe9378089                    	fbfcbc34-77b9-4901-82f1-e5fa78d5aa48	f	2024-07-01 07:17:07.289173	2024-07-01 07:17:07.289173	2024-07-02 07:17:07.289156
e01c4334049854f588c4d675da06299632651e4da55abc2858f3f0ba18b50936a2b4fd3673f39dab                    	fbfcbc34-77b9-4901-82f1-e5fa78d5aa48	f	2024-07-01 07:18:21.353173	2024-07-01 07:18:21.353173	2024-07-02 07:18:21.35312
a45ec662871cd0138a05a331d3f0f6a4e297ad6149c0f11e1d78322c78e72c9ecfa5f16f0cee212f                    	fbfcbc34-77b9-4901-82f1-e5fa78d5aa48	f	2024-07-01 07:43:11.338358	2024-07-01 07:43:11.338358	2024-07-02 07:43:11.338324
46ba8c745c17e4df41561bbbbd3732edc988d3461783a6813cf928abc833ad8fcc61a7e66eebc7b9                    	fbfcbc34-77b9-4901-82f1-e5fa78d5aa48	f	2024-07-01 07:47:35.162692	2024-07-01 07:47:35.162692	2024-07-02 07:47:35.162661
96d8ba96d9f4c97a847e5f6fbef1c284264f2280dcbbe48e32c5253287fb8da866bb0ce6993ec235                    	fbfcbc34-77b9-4901-82f1-e5fa78d5aa48	f	2024-07-01 08:46:39.674389	2024-07-01 08:46:39.674389	2024-07-02 08:46:39.674374
fd193e30ac52ff3368ec31e57a9e877dad9339ad9a8880408c44b6c74adc5855ed064753106f55c8                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-07-01 08:54:54.795794	2024-07-01 08:54:54.795794	2024-07-02 08:54:54.795774
00e9487f5474b82e28512f9997d1c98c5af68e9ef334f5982945784a53704603068c455eac0326a7                    	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	f	2024-07-01 08:55:02.125644	2024-07-01 08:55:02.125644	2024-07-02 08:55:02.125615
bf40d12531c0b5794892de761f447713028f4aad9698964d8b1d065313beac62d8bb3b949d155009                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-07-01 09:08:29.429832	2024-07-01 09:08:29.429832	2024-07-02 09:08:29.429806
6ef7f6d60d34ddd5e4ebe57f3dfa80048755af31a5b30b2cbc4c02257276071604ef3d4433085978                    	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	f	2024-07-01 09:17:11.385419	2024-07-01 09:17:11.385419	2024-07-02 09:17:11.3854
2e23b36a4ca708497bc9184bdfcf5b3d80d5ff4df82e615ab9fa89da038eca5756738b6db917aede                    	68ed7124-fa08-4720-b741-9fe4fa697c21	f	2024-07-01 09:17:53.988347	2024-07-01 09:17:53.988347	2024-07-02 09:17:53.988323
f23ffcc553c9d70909343c18f5788aea3932172e98812a29aa8f881b341b804ce36ad1b219fc5cde                    	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	f	2024-07-01 09:21:14.607483	2024-07-01 09:21:14.607483	2024-07-02 09:21:14.60746
22b936b8e7b3b5f3cd173b5a60d47b8fa47dfc2e5a3afda68bcf2c35f2028e32bcb019d4a3ad1f2b                    	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	f	2024-07-01 09:34:35.904401	2024-07-01 09:34:35.904401	2024-07-02 09:34:35.904376
9d3ba7fabde9e662d354262665ff4121044ef94f0ecb5fb4b02aa018d029c12fbda598a7a271015d                    	68ed7124-fa08-4720-b741-9fe4fa697c21	f	2024-07-01 09:34:48.79192	2024-07-01 09:34:48.79192	2024-07-02 09:34:48.791899
33f4801e3ee11fc534a730d9f8f9e7d47294831fb47aed49452afaf784919674d8499fa07f05ab32                    	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	f	2024-07-01 09:54:01.062398	2024-07-01 09:54:01.062398	2024-07-02 09:54:01.06237
89f18c001dfb833cb80481a45aef77c142dc3b79e40be3d912040e3f8eb6cdb08b615114f7c274f3                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-07-01 10:30:40.632032	2024-07-01 10:30:40.632032	2024-07-02 10:30:40.632005
79a453f5be98505aa4ea51d806669723679dd08ab4135603db67eee8699a22b4d3431a8583dda078                    	68ed7124-fa08-4720-b741-9fe4fa697c21	f	2024-07-01 10:30:53.507103	2024-07-01 10:30:53.507103	2024-07-02 10:30:53.507075
88b2a91f192e05a4472573d83c0f8c837e7f262184a778106a555dd6d5edd4d34aaef5c3e417b962                    	fbfcbc34-77b9-4901-82f1-e5fa78d5aa48	f	2024-07-01 10:40:35.434663	2024-07-01 10:40:35.434663	2024-07-02 10:40:35.434651
e13495847a1ca63c21264762a6c1affaf8e8ccb3744614d48e206a2a977664888eb9047f22bb1cef                    	fbfcbc34-77b9-4901-82f1-e5fa78d5aa48	f	2024-07-01 11:17:47.228839	2024-07-01 11:17:47.228839	2024-07-02 11:17:47.228827
886ffe125688cfece4cb03ef8ae404914fd5d6c48620e63eafe15c7ef017d0bf1167bb3d6e742d52                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-07-02 07:14:25.867868	2024-07-02 07:14:25.867868	2024-07-03 07:14:25.867853
460a28240081cb75c176603f5948624ea4c9624c1cc23b98f571f0cede14556605cf6dcb27b54769                    	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	f	2024-07-02 07:20:18.415334	2024-07-02 07:20:18.415334	2024-07-03 07:20:18.414425
9124d20cb389983840e163d9dce31254d9d6604258e2542175265a7f1b83040ed185e80f6532bb5d                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-07-02 07:24:44.57573	2024-07-02 07:24:44.57573	2024-07-03 07:24:44.575714
1ba9fbcba3205bb00461c18d59984791721660a68759b0e8878348c481f5d715c881986a72c25a99                    	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	f	2024-07-02 07:34:51.043372	2024-07-02 07:34:51.043372	2024-07-03 07:34:51.043351
011fd5a9988f5d751dfee26ec8f31d2423d5f61a2d404bdd9b145056cc068f17b6a2f8d86c4c014b                    	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	f	2024-07-02 09:26:15.314374	2024-07-02 09:26:15.314374	2024-07-03 09:26:15.314352
caa9a36666c0597d6cd10a8c29a84d62c4d8ff39f4cf92dcbd12f5baf324fbca2b246c078979b4d0                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-07-02 09:46:06.909392	2024-07-02 09:46:06.909392	2024-07-03 09:46:06.909371
4ad7ddf089517fa2bd3a4e349bda2d75542e6682ca268fece9b5f06149adb688fbb3743dd86f8284                    	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	f	2024-07-02 09:46:54.322635	2024-07-02 09:46:54.322635	2024-07-03 09:46:54.322607
2c38206a74a514216fa6aeec8d48fd269460fb59b29a01ea4749c29b2f8c25186865ef3db1884f96                    	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	f	2024-07-03 02:07:50.787669	2024-07-03 02:07:50.787669	2024-07-04 02:07:50.787648
75133df55b903aef74067e692bc40575cf0f96b2d2a1054f499352a7b5aec2fa527f74493c8c6f82                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-07-03 02:11:09.85259	2024-07-03 02:11:09.85259	2024-07-04 02:11:09.848517
e063255111f6a2a332fead0d95674203d108c4cd4984b7c445517829dc9ad531d52ca252e91dbc2c                    	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	f	2024-07-03 02:11:39.988835	2024-07-03 02:11:39.988835	2024-07-04 02:11:39.988812
924ba227561a8815095db4752cf9afedf3e468ab187603d16c1169fd369c2c81ff7c03e7e3d7ee89                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-07-03 02:13:21.151689	2024-07-03 02:13:21.151689	2024-07-04 02:13:21.151667
c31886e37adfb9207802358dd6b3c677cff8582f8ad5b4363cedcf5a71aab3ea4c440f057eb43718                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-07-03 02:36:54.514929	2024-07-03 02:36:54.514929	2024-07-04 02:36:54.514904
e97d6896e4b3fe54935f498863b40cab8145cfdeadd8765c384329c11db661f607661cd0708a9390                    	68ed7124-fa08-4720-b741-9fe4fa697c21	f	2024-07-03 02:50:29.746823	2024-07-03 02:50:29.746823	2024-07-04 02:50:29.746802
422fc4b9180e35637a23e5670b940739199c121a73e5b3d10594ffa8a1fbf995ecd9d424567744f0                    	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	f	2024-07-03 03:05:45.676353	2024-07-03 03:05:45.676353	2024-07-04 03:05:45.676328
3e2411f1d797544e4fe407204a11daa2904c2b601922cb62052d052e64c8ebac5ddce01df0fb921e                    	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	f	2024-07-03 03:05:54.271234	2024-07-03 03:05:54.271234	2024-07-04 03:05:54.271212
fcdff258640a77436f3e03655860250ce902963eb380c74a7a56c0b8ae07bdae5804c2492f577eb7                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-07-03 03:20:46.432956	2024-07-03 03:20:46.432956	2024-07-04 03:20:46.432928
11769595b634c0ccf965b2c3f36f5c383b62c3f7712bf585cb140b3af5ce3cc944d241f442e16312                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-07-03 03:21:01.90554	2024-07-03 03:21:01.90554	2024-07-04 03:21:01.905518
2bfb51052107179e24a4f45c2e4757779ee334b7c7a97eaaa02173f233f57bd00dfd7552a5fff2d8                    	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	f	2024-07-03 03:49:33.766498	2024-07-03 03:49:33.766498	2024-07-04 03:49:33.766477
dbd756d2658cbbf6c906cba0732873913dbdaf98095bf428ec98fa6646dd12381b44a7abfc236395                    	33392eb4-0f87-43c6-9893-9c014fe6d561	f	2024-07-03 16:23:18.22715	2024-07-03 16:23:18.22715	2024-07-04 16:23:18.227131
684de8e9f341f880d9d22e0e7ae202efcd1b8cbc08743ef59d1acf61df4f04c8322ceaf317fe27e5                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-07-04 00:25:30.932667	2024-07-04 00:25:30.932667	2024-07-05 00:25:30.93263
c3553cf830bdd1afb303486a37b9b1151c4a03be1800eb851dbc2da9c055b80a2450f36e3a23b46e                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-07-04 00:25:31.904021	2024-07-04 00:25:31.904021	2024-07-05 00:25:31.903994
73b4e1f0f1f1165315b327e0ec1fc3c4c4b8a57ddee6ed258e8b04ab4e7fb587bc838f18767f94a9                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-07-04 00:25:48.347582	2024-07-04 00:25:48.347582	2024-07-05 00:25:48.34756
c8529589019e07120e15483f62dbcfc148639d3b0080b77de0b09c7898901ebaa8e78334615e27aa                    	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	f	2024-07-04 02:08:52.662848	2024-07-04 02:08:52.662848	2024-07-05 02:08:52.662829
d6feab588274d895a0d73cf00778a712956c060d75b08700be952c89984f1b2bda696c9eb394a15a                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-07-04 03:22:59.509377	2024-07-04 03:22:59.509377	2024-07-05 03:22:59.509357
fadce968ea3846d1e85dfb27cd5921aeaf28dfd5d65a46ab6072f2741cf8d8b4f1057b2df8854033                    	68ed7124-fa08-4720-b741-9fe4fa697c21	f	2024-07-04 03:23:06.790893	2024-07-04 03:23:06.790893	2024-07-05 03:23:06.790872
379e1135e1663371954777522718286156cdf859dd53e8d439443ed24336ffd16455758c0822e452                    	68ed7124-fa08-4720-b741-9fe4fa697c21	f	2024-07-04 03:26:03.420162	2024-07-04 03:26:03.420162	2024-07-05 03:26:03.42014
f0feb5901c6daff30cdce9d169be8731b024cb9a222f38506f02edab8b22decbf41288c56744e240                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-07-04 03:30:17.774623	2024-07-04 03:30:17.774623	2024-07-05 03:30:17.7746
c260657d72083964356ed39046475994fe572f107eb8f1c5548b1a026a0001a5b7648e47a7c43a0f                    	68ed7124-fa08-4720-b741-9fe4fa697c21	f	2024-07-04 03:43:36.982165	2024-07-04 03:43:36.982165	2024-07-05 03:43:36.982141
264473cedc93459ad2e29ecf0d0ee00a5888cdc63518193cdf2a0515dfa0814db2380a293536b698                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-07-04 04:00:59.017172	2024-07-04 04:00:59.017172	2024-07-05 04:00:59.01715
3dc411262bbf60273d34a71a627af1dd805584c4c6b0f565b60b1753ef032daff36de9118506f8d9                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-07-04 04:14:06.02237	2024-07-04 04:14:06.02237	2024-07-05 04:14:06.022347
c6af77befefd7a706510808dc71db14a6b0afbaa6acdd4eaed9db351b7e2a18f63db66b16151af4c                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-07-04 04:28:48.288362	2024-07-04 04:28:48.288362	2024-07-05 04:28:48.288329
30f08f1d9fbef304be22bab827673fcef0beb7c89c4c8eb364907d87d988e4d82d3bb0a0289d6b09                    	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	f	2024-07-04 04:28:59.836563	2024-07-04 04:28:59.836563	2024-07-05 04:28:59.836504
e4aa1b6c5f6f0a36562535d00c82f94cd5a7ccef88a18839adbcbc138f511dea6eb2a2c4d3576717                    	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	f	2024-07-04 08:45:11.988722	2024-07-04 08:45:11.988722	2024-07-05 08:45:11.988703
8e3b32dc9c02caf32ae9319e8b3edd3cec92750dac8b829d1ec0bce1f3da9d157856c6b813466bae                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-07-04 09:21:12.737401	2024-07-04 09:21:12.737401	2024-07-05 09:21:12.73738
11e4241194008fcc85af4870382e3c07de1566260ee4982607c732d971c81f211c42028d9957003b                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-07-04 09:21:59.10381	2024-07-04 09:21:59.10381	2024-07-05 09:21:59.103786
fb199f78e1d2afe3ccf81275a658f4803c2c7c0d53152c3fdc539230cc0f3de3192b7d6148e7acde                    	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	f	2024-07-04 09:22:34.363532	2024-07-04 09:22:34.363532	2024-07-05 09:22:34.3635
c6461f185178d6ba24f55af549f198a2c4b5b1c66ccf105fdde4ae2eb78711a982c34844bb242cff                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-07-04 09:23:22.411828	2024-07-04 09:23:22.411828	2024-07-05 09:23:22.411805
0178323ad73fab13bc3f8ad1013223f363426dfc397f7949c2babcd89007de0ad0a870e16f278f2e                    	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	f	2024-07-04 09:25:49.419902	2024-07-04 09:25:49.419902	2024-07-05 09:25:49.419871
80de636a0cc746e7d94f80f2a2d541cca8d58990abb545724c6244148bc9f63eef0f00e4a0290958                    	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	f	2024-07-04 09:28:19.982119	2024-07-04 09:28:19.982119	2024-07-05 09:28:19.982097
a90dea94a867e850976dd26e05fa3926968d880e35e978e994090917fbccbd84f38c3229ef1c3de7                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-07-04 09:36:28.797829	2024-07-04 09:36:28.797829	2024-07-05 09:36:28.797784
d754cecc9b9efcbedc98778ad90480ce168b9ca41814288aef99a67304b3f7812274c304c0c4423b                    	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	f	2024-07-04 09:37:04.042602	2024-07-04 09:37:04.042602	2024-07-05 09:37:04.042573
33fdaba9a40dbd43041014a6e376d62968700f0406ce7fd09d639cf1572ecdd217c0de542643d849                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-07-04 09:46:06.196631	2024-07-04 09:46:06.196631	2024-07-05 09:46:06.196599
93c35d354310dbc5f1720dcc2b766259c205ecbeeed88b9203e4bbad4bc9ee45cbf604690820f769                    	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	f	2024-07-04 09:46:46.957696	2024-07-04 09:46:46.957696	2024-07-05 09:46:46.957678
6710124f61b543a32a8c1685a3eee11d31b1b2babce2c6a146bdb1887921d510373edf610cf27d58                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-07-04 10:00:12.201835	2024-07-04 10:00:12.201835	2024-07-05 10:00:12.201812
9e6c932053c4ed6b6c2c2bb401bc1d9198ddd1dcdcbbfafafa92bba6bca8b6222c171d8174704ee8                    	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	f	2024-07-04 10:00:36.585742	2024-07-04 10:00:36.585742	2024-07-05 10:00:36.585727
9361ad8c9063ca6109ec74da44bf1332b14f1dcff8fca0b6717be32c10b5e81428646f27ab4a1c99                    	33392eb4-0f87-43c6-9893-9c014fe6d561	f	2024-07-04 16:46:07.722613	2024-07-04 16:46:07.722613	2024-07-05 16:46:07.722581
c47d3af052f42ee6bacfec2633c66d4e82d76410112f06d14f63d83c88f980ef844e4897f95cecbc                    	33392eb4-0f87-43c6-9893-9c014fe6d561	f	2024-07-08 12:02:53.017324	2024-07-08 12:02:53.017324	2024-07-09 12:02:53.017268
f8ebc96d6e17df8263faaf283c20ab72088e96a4619465224a35c5a79bb6a6b22fa48ac5d7683b92                    	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	f	2024-07-08 15:28:32.721961	2024-07-08 15:28:32.721961	2024-07-09 15:28:32.721939
e5fb0701ec912d97cc253280221f332b1dc03e0870ccc5d632a7311b36aa78702d6338fe005d2ee0                    	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	f	2024-07-08 16:01:04.233432	2024-07-08 16:01:04.233432	2024-07-09 16:01:04.23341
9950efc971410dc2554339d51b02a7d54fc914b353123fc82146ea5fe6520f0f15f2f35b8b596cd1                    	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	f	2024-07-09 00:42:49.907956	2024-07-09 00:42:49.907956	2024-07-10 00:42:49.907933
d58d69d11edc8bc43f7acb56029994f3b81ef4a9e6de7ec8ac4733293401f2278f5348cec4b955b6                    	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	f	2024-07-09 15:47:02.558768	2024-07-09 15:47:02.558768	2024-07-10 15:47:02.558749
927b067e070700315f48a9e9fcb048b9e7cac1878c568143812a30b4c15b38a9c336e2121b85ce40                    	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	f	2024-07-10 01:57:49.861636	2024-07-10 01:57:49.861636	2024-07-11 01:57:49.861615
700b5289351eb43ccc474ae2d72bf67b1df665ec3c3c5be4622451573eb2eb0e3fd39f0a3db78880                    	68ed7124-fa08-4720-b741-9fe4fa697c21	f	2024-07-10 02:04:55.461347	2024-07-10 02:04:55.461347	2024-07-11 02:04:55.461327
0b672f0983baa478e770abd249f19f35cd857b943ea72bc3b78592ff131e5b48974d56802a924cc7                    	68ed7124-fa08-4720-b741-9fe4fa697c21	f	2024-07-10 02:05:19.897908	2024-07-10 02:05:19.897908	2024-07-11 02:05:19.897889
7396d040be67ea049cf4b091a9150fa560924b49c76340a7be371f3a54cf80bb29dccfc7c6095564                    	68ed7124-fa08-4720-b741-9fe4fa697c21	f	2024-07-10 02:06:11.320676	2024-07-10 02:06:11.320676	2024-07-11 02:06:11.320652
26a12f70ad4bdda1a6dc72693dfc9b828965448298c60d1e6c0e23522f7f8eb555164111ce8025df                    	68ed7124-fa08-4720-b741-9fe4fa697c21	f	2024-07-10 02:06:27.47432	2024-07-10 02:06:27.47432	2024-07-11 02:06:27.474291
e5a7f5fb0505cacd6c6b2585d38b501a694c77a76cedc5deefbb2cda21bcaa690aa4ef43f41ae443                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-07-10 02:06:37.484856	2024-07-10 02:06:37.484856	2024-07-11 02:06:37.484833
91922bd6eaebc0776e68ef827b198a50c9cd9274a2373aa5807726fcd6f7138a271c81baf2179a6f                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-07-10 02:12:52.163434	2024-07-10 02:12:52.163434	2024-07-11 02:12:52.163411
9b848be0f86b6df58c6dfd534a28ef3fad2c3d5dbbbac508bf203795888424d86a641e8f54522258                    	68ed7124-fa08-4720-b741-9fe4fa697c21	f	2024-07-10 02:12:59.061489	2024-07-10 02:12:59.061489	2024-07-11 02:12:59.061449
1133ff6179865165c1a596a56b0cedf45cfa9e41a7906954fc73a42300d6a181044aa023efd23568                    	68ed7124-fa08-4720-b741-9fe4fa697c21	f	2024-07-10 02:14:04.815497	2024-07-10 02:14:04.815497	2024-07-11 02:14:04.81547
316193b77fce5d00f06c384e74353c5ddce2abd9c6692564e5c2e45da1e0432d9ee9548b22a58ec1                    	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	f	2024-07-10 02:17:29.489437	2024-07-10 02:17:29.489437	2024-07-11 02:17:29.489403
bb78ec6cb855155b25a848045ca2ba3354173a9a63b0c8947f223eaa9ffdabead5650524896adc24                    	68ed7124-fa08-4720-b741-9fe4fa697c21	f	2024-07-10 02:18:59.327608	2024-07-10 02:18:59.327608	2024-07-11 02:18:59.327589
dab7d6d3b6e5a3f0a9ba2edb1baecfabd3b458e772692c34ab00b51633d3783ba46041ef40de9d77                    	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	f	2024-07-10 02:30:42.693379	2024-07-10 02:30:42.693379	2024-07-11 02:30:42.69335
e65b9d1eb9464a30d3a677dba7b210b4f5a5dafa81339dc65368665ff3e8aa83180e1270a354ec4b                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-07-10 02:40:00.208191	2024-07-10 02:40:00.208191	2024-07-11 02:40:00.208167
66c81a8f48ed300dd855d26c96cae0b170b182c7799e261baf82aaa926291b5bd4faf7483278cd3e                    	1e7a5eb3-c5c4-4fa3-b971-168765bfc413	f	2024-07-10 02:40:16.768717	2024-07-10 02:40:16.768717	2024-07-11 02:40:16.768692
b040a10557cdd0e396fcf15795dc69d6c0376eb6e554aed010ed9c66a4198ae01c1a02985c3f5b11                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-07-10 02:40:43.076826	2024-07-10 02:40:43.076826	2024-07-11 02:40:43.076796
78c5883b70f13fb713457d4bc442f99db0ddc3a5a3bb2153cc3d12926c8f904e40610a25b443485e                    	53e457f9-7d11-4e24-8bb5-a5134be9e2e7	f	2024-07-10 02:59:24.731495	2024-07-10 02:59:24.731495	2024-07-11 02:59:24.731477
dd8459ad24e694f06cb194c599903bec3158929611307288634a9df269b35109c1fd24d1ebf324bc                    	33392eb4-0f87-43c6-9893-9c014fe6d561	f	2024-07-10 14:51:47.796221	2024-07-10 14:51:47.796221	2024-07-11 14:51:47.796197
e90c6362ae4e255539315dbc0f8168dcdd810908985470af69f7b99c98b14cea28175e8f380f9e66                    	33392eb4-0f87-43c6-9893-9c014fe6d561	f	2024-07-10 15:17:43.372707	2024-07-10 15:17:43.372707	2024-07-11 15:17:43.372685
\.


--
-- Data for Name: user_roles; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.user_roles (user_id, role_id, created_at, updated_at) FROM stdin;
53e457f9-7d11-4e24-8bb5-a5134be9e2e7	8d01b6df-d26b-4b50-81df-b01167bedf0c	2024-06-06 09:29:38.558624	2024-06-06 09:29:38.558624
1e7a5eb3-c5c4-4fa3-b971-168765bfc413	73b6799d-ebd6-491f-9b47-272fb0f22914	2024-06-06 09:29:38.558626	2024-06-06 09:29:38.558626
68ed7124-fa08-4720-b741-9fe4fa697c21	bcc5cab0-6ac4-4aa7-b80d-7e7ef5113a10	2024-06-06 09:29:38.558627	2024-06-06 09:29:38.558627
fbfcbc34-77b9-4901-82f1-e5fa78d5aa48	00c12bd8-7470-40c5-938e-029b1239650c	2024-06-06 09:29:38.558628	2024-06-06 09:29:38.558628
eb4df53b-deea-47ee-8396-d6efc7962121	bcc5cab0-6ac4-4aa7-b80d-7e7ef5113a10	2024-06-27 07:10:12.960126	2024-06-27 07:10:12.960126
16638b6a-5ca3-4d1f-ab0f-ac4ce4ce30a0	bcc5cab0-6ac4-4aa7-b80d-7e7ef5113a10	2024-06-29 09:06:59.587851	2024-06-29 09:06:59.587851
33392eb4-0f87-43c6-9893-9c014fe6d561	73b6799d-ebd6-491f-9b47-272fb0f22914	2024-06-30 10:53:48.40983	2024-06-30 10:53:48.40983
33392eb4-0f87-43c6-9893-9c014fe6d561	00c12bd8-7470-40c5-938e-029b1239650c	2024-06-30 10:53:51.214591	2024-06-30 10:53:51.214591
33392eb4-0f87-43c6-9893-9c014fe6d561	bcc5cab0-6ac4-4aa7-b80d-7e7ef5113a10	2024-06-30 10:53:54.021741	2024-06-30 10:53:54.021741
33392eb4-0f87-43c6-9893-9c014fe6d561	8d01b6df-d26b-4b50-81df-b01167bedf0c	2024-06-30 10:53:56.674753	2024-06-30 10:53:56.674753
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.users (id, username, name, password, is_active, created_at, updated_at) FROM stdin;
53e457f9-7d11-4e24-8bb5-a5134be9e2e7	super.admin.one	Super Admin One	$2a$10$5Vw7Is.qZ2.0yLf919fMye.0AFlaXD0gbS3M4k7yQjN0OifCYi3hG	t	2024-06-06 09:29:38.554956	2024-06-06 09:29:38.554956
1e7a5eb3-c5c4-4fa3-b971-168765bfc413	inventory.one	Inventory One	$2a$10$5Vw7Is.qZ2.0yLf919fMye.0AFlaXD0gbS3M4k7yQjN0OifCYi3hG	t	2024-06-06 09:29:38.554957	2024-06-06 09:29:38.554957
68ed7124-fa08-4720-b741-9fe4fa697c21	cashier.one	Cashier One	$2a$10$5Vw7Is.qZ2.0yLf919fMye.0AFlaXD0gbS3M4k7yQjN0OifCYi3hG	t	2024-06-06 09:29:38.554958	2024-06-06 09:29:38.554958
fbfcbc34-77b9-4901-82f1-e5fa78d5aa48	driver.one	Driver One	$2a$10$5Vw7Is.qZ2.0yLf919fMye.0AFlaXD0gbS3M4k7yQjN0OifCYi3hG	t	2024-06-06 09:29:38.554959	2024-06-06 09:29:38.554959
eb4df53b-deea-47ee-8396-d6efc7962121	cashier.two.one	Cashier 2	$2a$10$R5Ewweu.a.Wx0/paHnfbU.kXoOwMm5VtEIvvy0s.WD1qbVkZTGXRi	t	2024-06-27 07:10:07.508235	2024-06-27 07:30:20.899999
16638b6a-5ca3-4d1f-ab0f-ac4ce4ce30a0	cashier.three	cashiern3	$2a$10$0x30GWQESBo/AubrAgY81uxY/uEpl7BAGAMB7kI9lWMikuuQG/jSa	t	2024-06-29 09:06:52.62957	2024-06-29 09:06:52.62957
33392eb4-0f87-43c6-9893-9c014fe6d561	wianto	Wianto WIjaya	$2a$10$LOlcZ8Mxja5IDekqEoN/Ye9wnmbzjxokOf79rQy0czczWgLARpa.G	t	2024-06-30 10:53:42.09971	2024-06-30 10:53:42.09971
0179c835-a049-47f3-a161-902b1ea89840	william	William	$2a$10$wJAK4ghTUnJwJRx.0/jg3O0mvoFIRE9.OWBpiL3izhHX7Ra/37b6u	t	2024-07-10 15:19:31.402689	2024-07-10 15:19:31.402689
\.


--
-- Name: balances balances_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.balances
    ADD CONSTRAINT balances_pk PRIMARY KEY (id);


--
-- Name: cart_items cart_items_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.cart_items
    ADD CONSTRAINT cart_items_pk PRIMARY KEY (id);


--
-- Name: cart_items cart_items_uk_1; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.cart_items
    ADD CONSTRAINT cart_items_uk_1 UNIQUE (cart_id, product_unit_id);


--
-- Name: carts carts_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.carts
    ADD CONSTRAINT carts_pk PRIMARY KEY (id);


--
-- Name: cashier_sessions cashier_sessions_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.cashier_sessions
    ADD CONSTRAINT cashier_sessions_pk PRIMARY KEY (id);


--
-- Name: customer_debts customer_debts_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.customer_debts
    ADD CONSTRAINT customer_debts_pk PRIMARY KEY (id);


--
-- Name: customer_payments customer_payments_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.customer_payments
    ADD CONSTRAINT customer_payments_pk PRIMARY KEY (id);


--
-- Name: customer_type_discounts customer_type_discounts_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.customer_type_discounts
    ADD CONSTRAINT customer_type_discounts_pk PRIMARY KEY (id);


--
-- Name: customer_type_discounts customer_type_discounts_uk_1; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.customer_type_discounts
    ADD CONSTRAINT customer_type_discounts_uk_1 UNIQUE (product_id, customer_type_id);


--
-- Name: customer_types customer_types_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.customer_types
    ADD CONSTRAINT customer_types_pk PRIMARY KEY (id);


--
-- Name: customer_types customer_types_uk_name; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.customer_types
    ADD CONSTRAINT customer_types_uk_name UNIQUE (name);


--
-- Name: customers customers_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.customers
    ADD CONSTRAINT customers_pk PRIMARY KEY (id);


--
-- Name: debt_payments debt_payments_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.debt_payments
    ADD CONSTRAINT debt_payments_pk PRIMARY KEY (id);


--
-- Name: debts debts_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.debts
    ADD CONSTRAINT debts_pk PRIMARY KEY (id);


--
-- Name: delivery_order_reviews delivery_order_delivery_order_id_uk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.delivery_order_reviews
    ADD CONSTRAINT delivery_order_delivery_order_id_uk UNIQUE (delivery_order_id);


--
-- Name: delivery_order_drivers delivery_order_drivers_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.delivery_order_drivers
    ADD CONSTRAINT delivery_order_drivers_pk PRIMARY KEY (id);


--
-- Name: delivery_order_images delivery_order_images_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.delivery_order_images
    ADD CONSTRAINT delivery_order_images_pk PRIMARY KEY (id);


--
-- Name: delivery_order_item_costs delivery_order_item_costs_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.delivery_order_item_costs
    ADD CONSTRAINT delivery_order_item_costs_pk PRIMARY KEY (id);


--
-- Name: delivery_order_items delivery_order_items_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.delivery_order_items
    ADD CONSTRAINT delivery_order_items_pk PRIMARY KEY (id);


--
-- Name: delivery_order_positions delivery_order_positions_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.delivery_order_positions
    ADD CONSTRAINT delivery_order_positions_pk PRIMARY KEY (id);


--
-- Name: delivery_order_positions delivery_order_positions_uk_delivery_order_id; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.delivery_order_positions
    ADD CONSTRAINT delivery_order_positions_uk_delivery_order_id UNIQUE (delivery_order_id);


--
-- Name: delivery_order_return_images delivery_order_return_images_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.delivery_order_return_images
    ADD CONSTRAINT delivery_order_return_images_pk PRIMARY KEY (id);


--
-- Name: delivery_order_returns delivery_order_returns_delivery_order_id_uk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.delivery_order_returns
    ADD CONSTRAINT delivery_order_returns_delivery_order_id_uk UNIQUE (delivery_order_id);


--
-- Name: delivery_order_returns delivery_order_returns_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.delivery_order_returns
    ADD CONSTRAINT delivery_order_returns_pk PRIMARY KEY (id);


--
-- Name: delivery_order_reviews delivery_order_reviews_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.delivery_order_reviews
    ADD CONSTRAINT delivery_order_reviews_pk PRIMARY KEY (id);


--
-- Name: delivery_orders delivery_orders_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.delivery_orders
    ADD CONSTRAINT delivery_orders_pk PRIMARY KEY (id);


--
-- Name: files files_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.files
    ADD CONSTRAINT files_pk PRIMARY KEY (id);


--
-- Name: permissions permission_title_uk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.permissions
    ADD CONSTRAINT permission_title_uk UNIQUE (title);


--
-- Name: permissions permissions_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.permissions
    ADD CONSTRAINT permissions_pk PRIMARY KEY (id);


--
-- Name: product_discounts product_discounts_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.product_discounts
    ADD CONSTRAINT product_discounts_pk PRIMARY KEY (id);


--
-- Name: product_discounts product_discounts_uk_product_id; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.product_discounts
    ADD CONSTRAINT product_discounts_uk_product_id UNIQUE (product_id);


--
-- Name: product_receive_items product_receive_items_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.product_receive_items
    ADD CONSTRAINT product_receive_items_pk PRIMARY KEY (id);


--
-- Name: product_receive_return_images product_receive_return_images_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.product_receive_return_images
    ADD CONSTRAINT product_receive_return_images_pk PRIMARY KEY (id);


--
-- Name: product_receive_returns product_receive_returns_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.product_receive_returns
    ADD CONSTRAINT product_receive_returns_pk PRIMARY KEY (id);


--
-- Name: product_receive_returns product_receive_returns_product_receive_id_uk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.product_receive_returns
    ADD CONSTRAINT product_receive_returns_product_receive_id_uk UNIQUE (product_receive_id);


--
-- Name: product_receive_images product_receives_images_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.product_receive_images
    ADD CONSTRAINT product_receives_images_pk PRIMARY KEY (id);


--
-- Name: product_receives product_receives_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.product_receives
    ADD CONSTRAINT product_receives_pk PRIMARY KEY (id);


--
-- Name: product_return_images product_return_images_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.product_return_images
    ADD CONSTRAINT product_return_images_pk PRIMARY KEY (id);


--
-- Name: product_return_items product_return_items_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.product_return_items
    ADD CONSTRAINT product_return_items_pk PRIMARY KEY (id);


--
-- Name: product_returns product_returns_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.product_returns
    ADD CONSTRAINT product_returns_pk PRIMARY KEY (id);


--
-- Name: product_stock_adjustments product_stock_adjustments_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.product_stock_adjustments
    ADD CONSTRAINT product_stock_adjustments_pk PRIMARY KEY (id);


--
-- Name: product_stock_mutations product_stock_mutations_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.product_stock_mutations
    ADD CONSTRAINT product_stock_mutations_pk PRIMARY KEY (id);


--
-- Name: product_stocks product_stocks_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.product_stocks
    ADD CONSTRAINT product_stocks_pk PRIMARY KEY (id);


--
-- Name: product_stocks product_stocks_uk_1; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.product_stocks
    ADD CONSTRAINT product_stocks_uk_1 UNIQUE (product_id);


--
-- Name: product_units product_units_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.product_units
    ADD CONSTRAINT product_units_pk PRIMARY KEY (id);


--
-- Name: products products_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_pk PRIMARY KEY (id);


--
-- Name: products products_uk_name; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_uk_name UNIQUE (name);


--
-- Name: purchase_order_images purchase_order_images_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.purchase_order_images
    ADD CONSTRAINT purchase_order_images_pk PRIMARY KEY (id);


--
-- Name: purchase_order_items purchase_order_items_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.purchase_order_items
    ADD CONSTRAINT purchase_order_items_pk PRIMARY KEY (id);


--
-- Name: purchase_orders purchase_orders_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.purchase_orders
    ADD CONSTRAINT purchase_orders_pk PRIMARY KEY (id);


--
-- Name: role_permissions role_permission_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.role_permissions
    ADD CONSTRAINT role_permission_pk PRIMARY KEY (role_id, permission_id);


--
-- Name: roles roles_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.roles
    ADD CONSTRAINT roles_pk PRIMARY KEY (id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: sequences sequences_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.sequences
    ADD CONSTRAINT sequences_pk PRIMARY KEY (id);


--
-- Name: sequences sequences_uk_1; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.sequences
    ADD CONSTRAINT sequences_uk_1 UNIQUE (unique_identifier, sequence);


--
-- Name: shop_order_items shop_order_items_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.shop_order_items
    ADD CONSTRAINT shop_order_items_pk PRIMARY KEY (id);


--
-- Name: shop_orders shop_orders_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.shop_orders
    ADD CONSTRAINT shop_orders_pk PRIMARY KEY (id);


--
-- Name: shop_orders shop_orders_uk_platform_identifier; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.shop_orders
    ADD CONSTRAINT shop_orders_uk_platform_identifier UNIQUE (platform_identifier);


--
-- Name: shop_orders shop_orders_uk_tracking_number; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.shop_orders
    ADD CONSTRAINT shop_orders_uk_tracking_number UNIQUE (tracking_number);


--
-- Name: shopee_configs shopee_configs_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.shopee_configs
    ADD CONSTRAINT shopee_configs_pk PRIMARY KEY (partner_id);


--
-- Name: supplier_types supplier_types_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.supplier_types
    ADD CONSTRAINT supplier_types_pk PRIMARY KEY (id);


--
-- Name: supplier_types supplier_types_uk_name; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.supplier_types
    ADD CONSTRAINT supplier_types_uk_name UNIQUE (name);


--
-- Name: suppliers suppliers_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.suppliers
    ADD CONSTRAINT suppliers_pk PRIMARY KEY (id);


--
-- Name: suppliers suppliers_uk_code; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.suppliers
    ADD CONSTRAINT suppliers_uk_code UNIQUE (code);


--
-- Name: tiktok_configs tiktok_configs_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.tiktok_configs
    ADD CONSTRAINT tiktok_configs_pk PRIMARY KEY (app_key);


--
-- Name: tiktok_products tiktok_products_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.tiktok_products
    ADD CONSTRAINT tiktok_products_pk PRIMARY KEY (tiktok_product_id);


--
-- Name: tiktok_products tiktok_products_product_uk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.tiktok_products
    ADD CONSTRAINT tiktok_products_product_uk UNIQUE (product_id);


--
-- Name: transaction_item_costs transaction_item_costs_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.transaction_item_costs
    ADD CONSTRAINT transaction_item_costs_pk PRIMARY KEY (id);


--
-- Name: transaction_items transaction_items_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.transaction_items
    ADD CONSTRAINT transaction_items_pk PRIMARY KEY (id);


--
-- Name: transaction_payments transaction_payments_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.transaction_payments
    ADD CONSTRAINT transaction_payments_pk PRIMARY KEY (id);


--
-- Name: transaction_payments transaction_payments_uk_transaction_id; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.transaction_payments
    ADD CONSTRAINT transaction_payments_uk_transaction_id UNIQUE (transaction_id);


--
-- Name: transactions transactions_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_pk PRIMARY KEY (id);


--
-- Name: units units_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.units
    ADD CONSTRAINT units_pk PRIMARY KEY (id);


--
-- Name: units units_uk_name; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.units
    ADD CONSTRAINT units_uk_name UNIQUE (name);


--
-- Name: user_access_tokens user_access_tokens_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.user_access_tokens
    ADD CONSTRAINT user_access_tokens_pk PRIMARY KEY (id);


--
-- Name: user_roles user_roles_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.user_roles
    ADD CONSTRAINT user_roles_pk PRIMARY KEY (user_id, role_id);


--
-- Name: users users_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pk PRIMARY KEY (id);


--
-- Name: users users_uk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_uk UNIQUE (username);


--
-- Name: cart_items cart_items_carts_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.cart_items
    ADD CONSTRAINT cart_items_carts_fk FOREIGN KEY (cart_id) REFERENCES public.carts(id);


--
-- Name: cart_items cart_items_product_units_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.cart_items
    ADD CONSTRAINT cart_items_product_units_fk FOREIGN KEY (product_unit_id) REFERENCES public.product_units(id);


--
-- Name: carts carts_cashier_sessions_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.carts
    ADD CONSTRAINT carts_cashier_sessions_fk FOREIGN KEY (cashier_session_id) REFERENCES public.cashier_sessions(id);


--
-- Name: cashier_sessions cashier_sessions_users_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.cashier_sessions
    ADD CONSTRAINT cashier_sessions_users_fk FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: customer_debts customer_debts_customers_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.customer_debts
    ADD CONSTRAINT customer_debts_customers_fk FOREIGN KEY (customer_id) REFERENCES public.customers(id);


--
-- Name: customer_payments customer_payment_files_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.customer_payments
    ADD CONSTRAINT customer_payment_files_fk FOREIGN KEY (image_file_id) REFERENCES public.files(id);


--
-- Name: customer_payments customer_payments_customer_debts_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.customer_payments
    ADD CONSTRAINT customer_payments_customer_debts_fk FOREIGN KEY (customer_debt_id) REFERENCES public.customer_debts(id);


--
-- Name: customer_type_discounts customer_type_discounts_customer_types_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.customer_type_discounts
    ADD CONSTRAINT customer_type_discounts_customer_types_fk FOREIGN KEY (customer_type_id) REFERENCES public.customer_types(id);


--
-- Name: customer_type_discounts customer_type_discounts_products_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.customer_type_discounts
    ADD CONSTRAINT customer_type_discounts_products_fk FOREIGN KEY (product_id) REFERENCES public.products(id);


--
-- Name: customers customers_customer_types_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.customers
    ADD CONSTRAINT customers_customer_types_fk FOREIGN KEY (customer_type_id) REFERENCES public.customer_types(id);


--
-- Name: debt_payments debt_payments_debts_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.debt_payments
    ADD CONSTRAINT debt_payments_debts_fk FOREIGN KEY (debt_id) REFERENCES public.debts(id);


--
-- Name: debt_payments debt_payments_files_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.debt_payments
    ADD CONSTRAINT debt_payments_files_fk FOREIGN KEY (image_file_id) REFERENCES public.files(id);


--
-- Name: debt_payments debt_payments_users_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.debt_payments
    ADD CONSTRAINT debt_payments_users_fk FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: delivery_order_drivers delivery_order_drivers_delivery_orders_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.delivery_order_drivers
    ADD CONSTRAINT delivery_order_drivers_delivery_orders_fk FOREIGN KEY (delivery_order_id) REFERENCES public.delivery_orders(id);


--
-- Name: delivery_order_drivers delivery_order_drivers_users_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.delivery_order_drivers
    ADD CONSTRAINT delivery_order_drivers_users_fk FOREIGN KEY (driver_user_id) REFERENCES public.users(id);


--
-- Name: delivery_order_images delivery_order_images_delivery_orders_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.delivery_order_images
    ADD CONSTRAINT delivery_order_images_delivery_orders_fk FOREIGN KEY (delivery_order_id) REFERENCES public.delivery_orders(id);


--
-- Name: delivery_order_images delivery_order_images_files_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.delivery_order_images
    ADD CONSTRAINT delivery_order_images_files_fk FOREIGN KEY (file_id) REFERENCES public.files(id);


--
-- Name: delivery_order_item_costs delivery_order_item_costs_delivery_order_items_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.delivery_order_item_costs
    ADD CONSTRAINT delivery_order_item_costs_delivery_order_items_fk FOREIGN KEY (delivery_order_item_id) REFERENCES public.delivery_order_items(id);


--
-- Name: delivery_order_items delivery_order_items_delivery_orders_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.delivery_order_items
    ADD CONSTRAINT delivery_order_items_delivery_orders_fk FOREIGN KEY (delivery_order_id) REFERENCES public.delivery_orders(id);


--
-- Name: delivery_order_items delivery_order_items_product_units_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.delivery_order_items
    ADD CONSTRAINT delivery_order_items_product_units_fk FOREIGN KEY (product_unit_id) REFERENCES public.product_units(id);


--
-- Name: delivery_order_positions delivery_order_positions_delivery_orders_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.delivery_order_positions
    ADD CONSTRAINT delivery_order_positions_delivery_orders_fk FOREIGN KEY (delivery_order_id) REFERENCES public.delivery_orders(id);


--
-- Name: delivery_order_positions delivery_order_positions_users_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.delivery_order_positions
    ADD CONSTRAINT delivery_order_positions_users_fk FOREIGN KEY (driver_user_id) REFERENCES public.users(id);


--
-- Name: delivery_order_return_images delivery_order_return_images_delivery_order_returns_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.delivery_order_return_images
    ADD CONSTRAINT delivery_order_return_images_delivery_order_returns_fk FOREIGN KEY (delivery_order_return_id) REFERENCES public.delivery_order_returns(id);


--
-- Name: delivery_order_return_images delivery_order_return_images_files_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.delivery_order_return_images
    ADD CONSTRAINT delivery_order_return_images_files_fk FOREIGN KEY (file_id) REFERENCES public.files(id);


--
-- Name: delivery_order_returns delivery_order_returns_delivery_orders_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.delivery_order_returns
    ADD CONSTRAINT delivery_order_returns_delivery_orders_fk FOREIGN KEY (delivery_order_id) REFERENCES public.delivery_orders(id);


--
-- Name: delivery_order_returns delivery_order_returns_users_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.delivery_order_returns
    ADD CONSTRAINT delivery_order_returns_users_fk FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: delivery_order_reviews delivery_order_reviews_delivery_orders_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.delivery_order_reviews
    ADD CONSTRAINT delivery_order_reviews_delivery_orders_fk FOREIGN KEY (delivery_order_id) REFERENCES public.delivery_orders(id);


--
-- Name: delivery_orders delivery_orders_customers_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.delivery_orders
    ADD CONSTRAINT delivery_orders_customers_fk FOREIGN KEY (customer_id) REFERENCES public.customers(id);


--
-- Name: delivery_orders delivery_orders_users_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.delivery_orders
    ADD CONSTRAINT delivery_orders_users_fk FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: product_discounts product_discounts_products_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.product_discounts
    ADD CONSTRAINT product_discounts_products_fk FOREIGN KEY (product_id) REFERENCES public.products(id);


--
-- Name: product_receive_items product_receive_items_product_receives_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.product_receive_items
    ADD CONSTRAINT product_receive_items_product_receives_fk FOREIGN KEY (product_receive_id) REFERENCES public.product_receives(id);


--
-- Name: product_receive_items product_receive_items_product_units_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.product_receive_items
    ADD CONSTRAINT product_receive_items_product_units_fk FOREIGN KEY (product_unit_id) REFERENCES public.product_units(id);


--
-- Name: product_receive_items product_receive_items_users_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.product_receive_items
    ADD CONSTRAINT product_receive_items_users_fk FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: product_receive_return_images product_receive_return_images_files_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.product_receive_return_images
    ADD CONSTRAINT product_receive_return_images_files_fk FOREIGN KEY (file_id) REFERENCES public.files(id);


--
-- Name: product_receive_returns product_receive_returns_product_receives_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.product_receive_returns
    ADD CONSTRAINT product_receive_returns_product_receives_fk FOREIGN KEY (product_receive_id) REFERENCES public.product_receives(id);


--
-- Name: product_receive_returns product_receive_returns_users_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.product_receive_returns
    ADD CONSTRAINT product_receive_returns_users_fk FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: product_receive_images product_receives_images_files_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.product_receive_images
    ADD CONSTRAINT product_receives_images_files_fk FOREIGN KEY (file_id) REFERENCES public.files(id);


--
-- Name: product_receive_images product_receives_images_product_receives_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.product_receive_images
    ADD CONSTRAINT product_receives_images_product_receives_fk FOREIGN KEY (product_receive_id) REFERENCES public.product_receives(id);


--
-- Name: product_receives product_receives_purchase_orders_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.product_receives
    ADD CONSTRAINT product_receives_purchase_orders_fk FOREIGN KEY (purchase_order_id) REFERENCES public.purchase_orders(id);


--
-- Name: product_receives product_receives_suppliers_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.product_receives
    ADD CONSTRAINT product_receives_suppliers_fk FOREIGN KEY (supplier_id) REFERENCES public.suppliers(id);


--
-- Name: product_receives product_receives_users_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.product_receives
    ADD CONSTRAINT product_receives_users_fk FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: product_return_images product_return_images_files_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.product_return_images
    ADD CONSTRAINT product_return_images_files_fk FOREIGN KEY (file_id) REFERENCES public.files(id);


--
-- Name: product_return_images product_return_images_product_returns_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.product_return_images
    ADD CONSTRAINT product_return_images_product_returns_fk FOREIGN KEY (product_return_id) REFERENCES public.product_returns(id);


--
-- Name: product_return_items product_return_items_product_returns_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.product_return_items
    ADD CONSTRAINT product_return_items_product_returns_fk FOREIGN KEY (product_return_id) REFERENCES public.product_returns(id);


--
-- Name: product_return_items product_return_items_product_units_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.product_return_items
    ADD CONSTRAINT product_return_items_product_units_fk FOREIGN KEY (product_unit_id) REFERENCES public.product_units(id);


--
-- Name: product_returns product_returns_suppliers_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.product_returns
    ADD CONSTRAINT product_returns_suppliers_fk FOREIGN KEY (supplier_id) REFERENCES public.suppliers(id);


--
-- Name: product_returns product_returns_users_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.product_returns
    ADD CONSTRAINT product_returns_users_fk FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: product_stock_adjustments product_stock_adjustments_product_stocks_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.product_stock_adjustments
    ADD CONSTRAINT product_stock_adjustments_product_stocks_fk FOREIGN KEY (product_stock_id) REFERENCES public.product_stocks(id);


--
-- Name: product_stock_adjustments product_stock_adjustments_users_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.product_stock_adjustments
    ADD CONSTRAINT product_stock_adjustments_users_fk FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: product_stock_mutations product_stock_mutations_product_units_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.product_stock_mutations
    ADD CONSTRAINT product_stock_mutations_product_units_fk FOREIGN KEY (product_unit_id) REFERENCES public.product_units(id);


--
-- Name: product_stocks product_stocks_products_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.product_stocks
    ADD CONSTRAINT product_stocks_products_fk FOREIGN KEY (product_id) REFERENCES public.products(id);


--
-- Name: product_units product_units_products_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.product_units
    ADD CONSTRAINT product_units_products_fk FOREIGN KEY (product_id) REFERENCES public.products(id);


--
-- Name: product_units product_units_units_fk_1; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.product_units
    ADD CONSTRAINT product_units_units_fk_1 FOREIGN KEY (to_unit_id) REFERENCES public.units(id);


--
-- Name: product_units product_units_units_fk_2; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.product_units
    ADD CONSTRAINT product_units_units_fk_2 FOREIGN KEY (unit_id) REFERENCES public.units(id);


--
-- Name: products products_files_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_files_fk FOREIGN KEY (image_file_id) REFERENCES public.files(id);


--
-- Name: purchase_order_images purchase_order_images_files_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.purchase_order_images
    ADD CONSTRAINT purchase_order_images_files_fk FOREIGN KEY (file_id) REFERENCES public.files(id);


--
-- Name: purchase_order_images purchase_order_images_purchase_orders_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.purchase_order_images
    ADD CONSTRAINT purchase_order_images_purchase_orders_fk FOREIGN KEY (purchase_order_id) REFERENCES public.purchase_orders(id);


--
-- Name: purchase_order_items purchase_order_items_product_units_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.purchase_order_items
    ADD CONSTRAINT purchase_order_items_product_units_fk FOREIGN KEY (product_unit_id) REFERENCES public.product_units(id);


--
-- Name: purchase_order_items purchase_order_items_purchase_orders_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.purchase_order_items
    ADD CONSTRAINT purchase_order_items_purchase_orders_fk FOREIGN KEY (purchase_order_id) REFERENCES public.purchase_orders(id);


--
-- Name: purchase_order_items purchase_order_items_users_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.purchase_order_items
    ADD CONSTRAINT purchase_order_items_users_fk FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: purchase_orders purchase_orders_suppliers_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.purchase_orders
    ADD CONSTRAINT purchase_orders_suppliers_fk FOREIGN KEY (supplier_id) REFERENCES public.suppliers(id);


--
-- Name: purchase_orders purchase_orders_users_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.purchase_orders
    ADD CONSTRAINT purchase_orders_users_fk FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: role_permissions role_permission_permission_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.role_permissions
    ADD CONSTRAINT role_permission_permission_fk FOREIGN KEY (permission_id) REFERENCES public.permissions(id);


--
-- Name: role_permissions role_permission_roles_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.role_permissions
    ADD CONSTRAINT role_permission_roles_fk FOREIGN KEY (role_id) REFERENCES public.roles(id);


--
-- Name: shop_order_items shop_order_items_product_units_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.shop_order_items
    ADD CONSTRAINT shop_order_items_product_units_fk FOREIGN KEY (product_unit_id) REFERENCES public.product_units(id);


--
-- Name: shop_order_items shop_order_items_shop_orders_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.shop_order_items
    ADD CONSTRAINT shop_order_items_shop_orders_fk FOREIGN KEY (shop_order_id) REFERENCES public.shop_orders(id);


--
-- Name: suppliers suppliers_supplier_types_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.suppliers
    ADD CONSTRAINT suppliers_supplier_types_fk FOREIGN KEY (supplier_type_id) REFERENCES public.supplier_types(id);


--
-- Name: tiktok_products tiktok_products_products_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.tiktok_products
    ADD CONSTRAINT tiktok_products_products_fk FOREIGN KEY (product_id) REFERENCES public.products(id);


--
-- Name: transaction_item_costs transaction_item_costs_transaction_items_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.transaction_item_costs
    ADD CONSTRAINT transaction_item_costs_transaction_items_fk FOREIGN KEY (transaction_item_id) REFERENCES public.transaction_items(id);


--
-- Name: transaction_items transaction_items_product_units_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.transaction_items
    ADD CONSTRAINT transaction_items_product_units_fk FOREIGN KEY (product_unit_id) REFERENCES public.product_units(id);


--
-- Name: transaction_items transaction_items_transactions_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.transaction_items
    ADD CONSTRAINT transaction_items_transactions_fk FOREIGN KEY (transaction_id) REFERENCES public.transactions(id);


--
-- Name: transaction_payments transaction_payments_transactions_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.transaction_payments
    ADD CONSTRAINT transaction_payments_transactions_fk FOREIGN KEY (transaction_id) REFERENCES public.transactions(id);


--
-- Name: transactions transactions_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_fk FOREIGN KEY (cashier_session_id) REFERENCES public.cashier_sessions(id);


--
-- Name: user_access_tokens user_access_tokens_users_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.user_access_tokens
    ADD CONSTRAINT user_access_tokens_users_fk FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: user_roles user_roles_roles_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.user_roles
    ADD CONSTRAINT user_roles_roles_fk FOREIGN KEY (role_id) REFERENCES public.roles(id);


--
-- Name: user_roles user_roles_users_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.user_roles
    ADD CONSTRAINT user_roles_users_fk FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- PostgreSQL database dump complete
--

