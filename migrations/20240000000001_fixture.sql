-- +goose Up
-- +goose StatementBegin

-- WARNING: mock data has been partially created by AI

insert into attribute_group([id], [name]) values
(1, 'Computer Attributes');

insert into attribute_set([id], [attribute_group_id], [name], [in_box], [in_filter], [position]) values
(1, 1, 'Generation', 0, 1, 1),
(2, 1, 'CPU Type', 0, 1, 2),
(3, 1, 'CPU Model', 1, 0, 3),
(4, 1, 'RAM Amount', 1, 1, 4),
(5, 1, 'RAM Type', 1, 1, 5),
(6, 1, 'Disk Capacity', 1, 1, 6),
(7, 1, 'Disk Type', 1, 1, 7),
(8, 1, 'Video Card', 1, 0, 8),
(9, 1, 'Power Supply', 1, 0, 9);

-- Generation attribute_set_id = 1
insert into attribute_value ([id], [attribute_set_id], [name], [position]) values
(1, 1, '1st Gen', 1),
(2, 1, '2nd Gen', 2),
(3, 1, '3rd Gen', 3),
(4, 1, '4th Gen', 4),
(5, 1, '5th Gen', 5),
(6, 1, '6th Gen', 6),
(7, 1, '7th Gen', 7),
(8, 1, '8th Gen', 8),
(9, 1, '9th Gen', 9),
(10, 1, '10th Gen', 10);

-- CPU Type attribute_set_id = 2
insert into attribute_value ([id], [attribute_set_id], [name], [position]) values
(11, 2, 'Intel i3', 1),
(12, 2, 'Intel i5', 2),
(13, 2, 'Intel i7', 3),
(14, 2, 'Intel i9', 4),
(15, 2, 'AMD Ryzen 3', 5),
(16, 2, 'AMD Ryzen 5', 6),
(17, 2, 'AMD Ryzen 7', 7),
(18, 2, 'AMD Ryzen 9', 8),
(19, 2, 'Apple Silicon', 9),
(20, 2, 'Qualcomm', 10);

-- CPU Model attribute_set_id = 3
insert into attribute_value ([id], [attribute_set_id], [name], [position]) values
(21, 3, 'Intel Core i3-14100', 1),
(22, 3, 'Intel Core i5-14400', 2),
(23, 3, 'Intel Core i7-14700', 3),
(24, 3, 'Intel Core i9-11900', 4),
(25, 3, 'AMD Ryzen 3 2200G', 5),
(26, 3, 'AMD Ryzen 5 5600x', 6),
(27, 3, 'AMD Ryzen 7 R7-3700X', 7),
(28, 3, 'AMD Ryzen 9 5900x', 8),
(29, 3, 'Apple M1 Pro 16-core', 9),
(30, 3, 'Snapdragon 8cx Gen 3', 10);

-- RAM Amount attribute_set_id = 4
insert into attribute_value ([id], [attribute_set_id], [name], [position]) values
(31, 4, '256 MB', 1),
(32, 4, '512 MB', 2),
(33, 4, '1 GB', 3),
(34, 4, '2 GB', 4),
(35, 4, '4 GB', 5),
(36, 4, '8 GB', 6),
(37, 4, '16 GB', 7),
(38, 4, '32 GB', 8),
(39, 4, '64 GB', 9),
(40, 4, '128 GB', 10);

-- RAM Type attribute_set_id = 5
insert into attribute_value ([id], [attribute_set_id], [name], [position]) values
(41, 5, 'DDR3', 1),
(42, 5, 'DDR3 ECC Reg.', 2),
(43, 5, 'DDR3L', 3),
(44, 5, 'DDR3L ECC Reg.', 4),
(45, 5, 'DDR4', 5),
(46, 5, 'DDR4 ECC Reg.', 6),
(47, 5, 'DDR5', 7),
(48, 5, 'Integrated', 8),
(49, 5, 'Buffered', 9),
(50, 5, 'External', 10);

-- Disk Capacity attribute_set_id = 6
insert into attribute_value ([id], [attribute_set_id], [name], [position]) values
(51, 6, '120 GB', 1),
(52, 6, '240 GB', 2),
(53, 6, '480 GB', 3),
(54, 6, '960 GB', 4),
(55, 6, '1 TB', 5),
(56, 6, '2 TB', 6),
(57, 6, '4 TB', 7),
(58, 6, '8 TB', 8),
(59, 6, '16 TB', 9),
(60, 6, '32 TB', 10);

-- Disk Type attribute_set_id = 7
insert into attribute_value ([id], [attribute_set_id], [name], [position]) values
(61, 7, 'HDD', 1),
(62, 7, 'SSD', 2),
(63, 7, 'Hybrid', 3),
(64, 7, 'NVMe', 4),
(65, 7, 'SATA', 5),
(66, 7, 'PCIe', 6),
(67, 7, 'M.2', 7),
(68, 7, 'eMMC', 8),
(69, 7, 'U.2', 9),
(70, 7, 'RAID', 10);

-- Video Card attribute_set_id = 8
insert into attribute_value ([id], [attribute_set_id], [name], [position]) values
(71, 8, 'Integrated', 1),
(72, 8, 'GeForce GTX 1050', 2),
(73, 8, 'GeForce GTX 1660', 3),
(74, 8, 'GeForce RTX 3060', 4),
(75, 8, 'GeForce RTX 3080', 5),
(76, 8, 'Radeon RX 580', 6),
(77, 8, 'Radeon RX 6700', 7),
(78, 8, 'Radeon RX 7900', 8),
(79, 8, 'Quadro RTX 4000', 9),
(80, 8, 'Titan V', 10);

-- Power Supply attribute_set_id = 9
insert into attribute_value ([id], [attribute_set_id], [name], [position]) values
(81, 9, '300W', 1),
(82, 9, '400W', 2),
(83, 9, '500W', 3),
(84, 9, '600W', 4),
(85, 9, '700W', 5),
(86, 9, '800W', 6),
(87, 9, '900W', 7),
(88, 9, '1000W', 8),
(89, 9, '1200W', 9),
(90, 9, '1500W', 10);

insert into brand ([id], [name]) values
(1, 'Dell'),
(2, 'HP'),
(3, 'Lenovo'),
(4, 'Apple'),
(5, 'Asus'),
(6, 'Acer'),
(7, 'MSI'),
(8, 'Gigabyte'),
(9, 'Razer'),
(10, 'Samsung'),
(11, 'Microsoft'),
(12, 'Intel'),
(13, 'AMD'),
(14, 'NVIDIA'),
(15, 'Corsair'),
(16, 'Western Digital'),
(17, 'Seagate'),
(18, 'Kingston'),
(19, 'Crucial'),
(20, 'EVGA'),
(21, 'ASRock'),
(22, 'Thermaltake'),
(23, 'Cooler Master'),
(24, 'NZXT'),
(25, 'Alienware'),
(26, 'Logitech'),
(27, 'SteelSeries'),
(28, 'BenQ'),
(29, 'ViewSonic'),
(30, 'ZOTAC'),
(31, 'Palit'),
(32, 'Sapphire'),
(33, 'ADATA'),
(34, 'Patriot'),
(35, 'PNY'),
(36, 'Lian Li'),
(37, 'Fractal Design'),
(38, 'Antec'),
(39, 'Sharkoon'),
(40, 'Sabrent'),
(41, 'Noctua'),
(42, 'Be Quiet!'),
(43, 'DeepCool'),
(44, 'Scythe'),
(45, 'ARCTIC'),
(46, 'Havit'),
(47, 'Toshiba'),
(48, 'LG'),
(49, 'Philips'),
(50, 'Sony');

insert into category([id], [slug], [name], [icon], [is_published], [attribute_group_id], [parent_id], [mp_path], [mp_level], [mp_position]) values
(1, 'computers', 'Computers', 'fa fa-computer', 1, null, null, '1.', 0, 1),
(2, 'desktops', 'Desktops', null, 1, 1, 1, '1.2.', 1, 1),
(3, 'laptops', 'Laptops', null, 1, 1, 1, '1.3', 1, 2),
(4, 'components', 'Components', 'fa fa-memory', 1, null, null, '4.', 0, 2),
(5, 'processors', 'Processors', null, 1, null, 4, '4.5.', 1, 1),
(6, 'motherboards', 'Motherboards', null, 1, null, 4, '4.6.', 1, 2),
(7, 'memory', 'Memory', null, 1, null, 4, '4.7.', 1, 3),
(8, 'storage', 'Storage', null, 1, null, 4, '4.8.', 1, 4),
(9, 'video', 'Video', null, 1, null, 4, '4.9.', 1, 5),
(10, 'power-supplies', 'Power Supplies', null, 1, null, 4, '4.10.', 1, 6);

insert into supplier([id], [slug], [code], [name], [description], [is_published], [position]) values
(1, 'stk', 'stk', 'Stock', 'Available in stock', 1, 1),
(2, 'ord', 'ord', 'Order', 'Available upon request', 1, 2);

insert into product_status([id], [name], [color], [position]) values
(1, 'Hot', '0ea5e9', 1),
(2, 'New', '84cc16', 2),
(3, 'Price', 'dc2626', 3);

insert into product ([id], [code], [name], [slug], [description], [quantity], [price], [brand_id], [category_id], [supplier_id], [status_id], [is_published], [created_at], [updated_at]) VALUES
(1, 'D1XPRO', 'Dell Inspiron X1 Pro', 'd1xpro-dell-inspiron-x1-pro', 'High-performance desktop for gaming', 10, 899.99, 1, 2, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(2, 'H2P4GE', 'HP Pavilion 24', 'h2p4ge-hp-pavilion-24', 'All-in-one desktop with 4K display', 15, 1199.99, 2, 2, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(3, 'L3TGE', 'Lenovo ThinkCentre T3', 'l3tge-lenovo-thinkcentre-t3', 'Compact desktop for professionals', 20, 799.99, 3, 2, 2, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(4, 'A4IMX', 'Apple iMac X', 'a4imx-apple-imac-x', 'Sleek design with powerful specs', 12, 1499.99, 4, 2, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(5, 'A5XPS', 'Asus XPS 990', 'a5xps-asus-xps-990', 'Gaming desktop with advanced graphics', 8, 1299.99, 5, 2, 1, 1, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(6, 'M6Z12', 'MSI Z12 Gaming PC', 'm6z12-msi-z12-gaming-pc', 'High-end gaming rig', 5, 1599.99, 7, 2, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(7, 'G7R5T', 'Gigabyte R5 Tower', 'g7r5t-gigabyte-r5-tower', 'Reliable desktop for everyday use', 25, 699.99, 8, 2, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(8, 'R8DAG', 'Razer DeathAdder Gaming PC', 'r8dag-razer-deathadder-gaming-pc', 'Gaming desktop with RGB lighting', 10, 1399.99, 9, 2, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(9, 'S9RAZ', 'Samsung Razer Pro', 's9raz-samsung-razer-pro', 'Professional desktop with SSD', 18, 1099.99, 10, 2, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(10, 'D10WPC', 'Dell Workstation Pro', 'd10wpc-dell-workstation-pro', 'Powerful workstation for heavy tasks', 8, 1399.99, 1, 2, 2, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(11, 'H11OM', 'HP Omen X', 'h11om-hp-omen-x', 'High-performance gaming desktop', 6, 1599.99, 2, 2, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(12, 'L12IM', 'Lenovo IdeaCentre M', 'l12im-lenovo-ideacentre-m', 'Versatile desktop for home and office', 22, 799.99, 3, 2, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(13, 'A13S4', 'Apple Studio 4K', 'a13s4-apple-studio-4k', 'High-resolution desktop for creative work', 9, 1499.99, 4, 2, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(14, 'A14G5', 'Asus Gaming X5', 'a14g5-asus-gaming-x5', 'Gaming desktop with high refresh rate', 7, 1199.99, 5, 2, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(15, 'M15RT', 'MSI Ryzen Tower', 'm15rt-msi-ryzen-tower', 'Desktop with AMD Ryzen processor', 13, 1299.99, 7, 2, 2, 2, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(16, 'G16DE', 'Gigabyte Desk Elite', 'g16de-gigabyte-desk-elite', 'Premium desktop with enhanced performance', 11, 1399.99, 8, 2, 2, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(17, 'R17DA', 'Razer Desktop Advanced', 'r17da-razer-desktop-advanced', 'High-end desktop for gaming enthusiasts', 6, 1499.99, 9, 2, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(18, 'S18P2', 'Samsung Performance 2', 's18p2-samsung-performance-2', 'Desktop with premium performance', 14, 1199.99, 10, 2, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(19, 'D19M', 'Dell Mini PC', 'd19m-dell-mini-pc', 'Compact desktop for small spaces', 20, 699.99, 1, 2, 1, 3, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(20, 'H20LE', 'HP Elite Tower', 'h20le-hp-elite-tower', 'Professional tower for business use', 16, 1199.99, 2, 2, 2, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(21, 'L21U8', 'Lenovo Ultra 8', 'l21u8-lenovo-ultra-8', 'High-performance ultra desktop', 10, 1399.99, 3, 2, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(22, 'A22P5', 'Apple Pro 5', 'a22p5-apple-pro-5', 'Desktop for professional use with high specs', 11, 1499.99, 4, 2, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(23, 'A23G2', 'Asus Gaming 2', 'a23g2-asus-gaming-2', 'Desktop with advanced cooling system', 7, 1299.99, 5, 2, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(24, 'M24M1', 'MSI Mega 1', 'm24m1-msi-mega-1', 'Powerful desktop with high capacity', 9, 1399.99, 7, 2, 2, NULL, 0, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(25, 'G25DX', 'Gigabyte DX Series', 'g25dx-gigabyte-dx-series', 'Desktop with powerful graphics', 12, 1199.99, 8, 2, 1, NULL, 0, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(26, 'R26F3', 'Razer Fusion 3', 'r26f3-razer-fusion-3', 'Gaming desktop with latest features', 8, 1599.99, 9, 2, 1, 3, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(27, 'S27H6', 'Samsung High 6', 's27h6-samsung-high-6', 'High-end desktop with sleek design', 14, 1299.99, 10, 2, 1, NULL, 0, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(28, 'D28WX', 'Dell Workstation X', 'd28wx-dell-workstation-x', 'Professional workstation with top specs', 5, 1499.99, 1, 2, 2, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(29, 'H29MS', 'HP Pro Series', 'h29ms-hp-pro-series', 'Series of desktops for professional use', 12, 1199.99, 2, 2, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(30, 'L30G4', 'Lenovo Gaming 4', 'l30g4-lenovo-gaming-4', 'High-performance gaming desktop', 11, 1399.99, 3, 2, 2, 1, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(31, 'A31P7', 'Apple Pro 7', 'a31p7-apple-pro-7', 'Professional desktop with enhanced specs', 6, 1499.99, 4, 2, 1, NULL, 0, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(32, 'A32X1', 'Asus X1 Elite', 'a32x1-asus-x1-elite', 'Elite gaming desktop with high-end specs', 8, 1399.99, 5, 2, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(33, 'M33M9', 'MSI M9 Max', 'm33m9-msi-m9-max', 'Maximized performance desktop', 9, 1499.99, 7, 2, 2, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(34, 'G34Z8', 'Gigabyte Z8 Series', 'g34z8-gigabyte-z8-series', 'Series desktop with high specs', 12, 1199.99, 8, 2, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(35, 'R35V5', 'Razer V5 Elite', 'r35v5-razer-v5-elite', 'Elite gaming desktop with top features', 10, 1599.99, 9, 2, 1, 1, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(36, 'S36X3', 'Samsung X3 Pro', 's36x3-samsung-x3-pro', 'High-performance desktop with pro specs', 15, 1399.99, 10, 2, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(37, 'D37GX', 'Dell Gaming X', 'd37gx-dell-gaming-x', 'Gaming desktop with advanced graphics', 7, 1499.99, 1, 2, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(38, 'H38P3', 'HP Power 3', 'h38p3-hp-power-3', 'Powerful desktop for high-demand tasks', 14, 1299.99, 2, 2, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(39, 'L39G6', 'Lenovo Gaming 6', 'l39g6-lenovo-gaming-6', 'High-end gaming desktop with great specs', 11, 1399.99, 3, 2, 2, NULL, 0, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(40, 'A40X9', 'Apple X9 Professional', 'a40x9-apple-x9-professional', 'Top-of-the-line professional desktop', 6, 1599.99, 4, 2, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(41, 'D41E4', 'Dell XPS 15', 'd41e4-dell-xps-15', 'High-performance laptop for professionals', 25, 1499.99, 1, 3, 1, 1, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(42, 'H42S3', 'HP Spectre x360', 'h42s3-hp-spectre-x360', 'Convertible laptop with high specs', 20, 1399.99, 2, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(43, 'L43T8', 'Lenovo ThinkPad T14', 'l43t8-lenovo-thinkpad-t14', 'Reliable business laptop with great performance', 30, 1299.99, 3, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(44, 'A44M2', 'Apple MacBook Pro', 'a44m2-apple-macbook-pro', 'Professional laptop with M2 chip', 18, 1799.99, 4, 3, 1, 2, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(45, 'A45G6', 'Asus ROG Zephyrus', 'a45g6-asus-rog-zephyrus', 'Gaming laptop with high-end graphics', 22, 1499.99, 5, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(46, 'M46B3', 'MSI Bravo 15', 'm46b3-msi-bravo-15', 'Affordable gaming laptop with decent specs', 20, 1199.99, 7, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(47, 'G47A7', 'Gigabyte Aero 15', 'g47a7-gigabyte-aero-15', 'High-performance laptop for creators', 15, 1599.99, 8, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(48, 'R48V6', 'Razer Blade 15', 'r48v6-razer-blade-15', 'Premium gaming laptop with top specs', 10, 1699.99, 9, 3, 1, 2, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(49, 'S49X9', 'Samsung Galaxy Book', 's49x9-samsung-galaxy-book', 'Convertible laptop with AMOLED screen', 12, 1399.99, 10, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(50, 'D50Q7', 'Dell Latitude 7420', 'd50q7-dell-latitude-7420', 'Business laptop with high durability', 24, 1299.99, 1, 3, 2, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(51, 'H51G2', 'HP Envy 14', 'h51g2-hp-envy-14', 'Stylish laptop with high-end features', 18, 1399.99, 2, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(52, 'L52S5', 'Lenovo Yoga 7i', 'l52s5-lenovo-yoga-7i', 'Convertible laptop with high performance', 25, 1299.99, 3, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(53, 'A53M1', 'Apple MacBook Air', 'a53m1-apple-macbook-air', 'Lightweight laptop with impressive specs', 15, 1199.99, 4, 3, 1, 3, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(54, 'A54X8', 'Asus ZenBook 14', 'a54x8-asus-zenbook-14', 'High-performance ultrabook', 20, 1399.99, 5, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(55, 'M55B9', 'MSI Creator 17', 'm55b9-msi-creator-17', 'Laptop for creative professionals', 10, 1699.99, 7, 3, 2, 3, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(56, 'G56T3', 'Gigabyte G5', 'g56t3-gigabyte-g5', 'Affordable gaming laptop with good specs', 22, 1199.99, 8, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(57, 'R57P6', 'Razer Stealth 13', 'r57p6-razer-stealth-13', 'Compact yet powerful laptop', 16, 1499.99, 9, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(58, 'S58M7', 'Samsung Notebook 9', 's58m7-samsung-notebook-9', 'High-performance notebook with sleek design', 18, 1399.99, 10, 3, 1, NULL, 0, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(59, 'D59A2', 'Dell Inspiron 14', 'd59a2-dell-inspiron-14', 'Affordable laptop for everyday use', 28, 999.99, 1, 3, 1, 3, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(60, 'H60Q4', 'HP Pavilion 15', 'h60q4-hp-pavilion-15', 'Laptop with balanced performance and price', 19, 1199.99, 2, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(61, 'L61B7', 'Lenovo Legion 5', 'l61b7-lenovo-legion-5', 'Gaming laptop with powerful graphics', 17, 1499.99, 3, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(62, 'A62X4', 'Apple MacBook 16', 'a62x4-apple-macbook-16', 'High-end laptop with large display', 13, 1799.99, 4, 3, 2, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(63, 'A63G9', 'Asus TUF Dash F15', 'a63g9-asus-tuf-dash-f15', 'Durable gaming laptop', 14, 1399.99, 5, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(64, 'M64M3', 'MSI GS66 Stealth', 'm64m3-msi-gs66-stealth', 'High-performance laptop with advanced features', 8, 1599.99, 7, 3, 2, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(65, 'G65X8', 'Gigabyte G7', 'g65x8-gigabyte-g7', 'Gaming laptop with high refresh rate', 20, 1499.99, 8, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(66, 'R66P5', 'Razer Book 13', 'r66p5-razer-book-13', 'Compact high-performance laptop', 12, 1299.99, 9, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(67, 'S67M2', 'Samsung Odyssey', 's67m2-samsung-odyssey', 'Gaming laptop with immersive experience', 15, 1599.99, 10, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(68, 'D68E4', 'Dell Latitude 9410', 'd68e4-dell-latitude-9410', 'Business laptop with enterprise features', 23, 1399.99, 1, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(69, 'H69G2', 'HP Elite Dragonfly', 'h69g2-hp-elite-dragonfly', 'Premium business laptop with high-end features', 10, 1799.99, 2, 3, 2, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(70, 'L70S7', 'Lenovo Flex 5', 'l70s7-lenovo-flex-5', 'Convertible laptop with good performance', 18, 1199.99, 3, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(71, 'A71X9', 'Apple MacBook Air M2', 'a71x9-apple-macbook-air-m2', 'Updated MacBook Air with M2 chip', 14, 1399.99, 4, 3, 1, 2, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(72, 'A72G3', 'Asus Zephyrus G14', 'a72g3-asus-zephyrus-g14', 'Compact gaming laptop with top specs', 12, 1499.99, 5, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(73, 'M73B6', 'MSI Prestige 14', 'm73b6-msi-prestige-14', 'Business laptop with sleek design', 16, 1299.99, 7, 3, 2, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(74, 'G74T9', 'Gigabyte Aorus 15', 'g74t9-gigabyte-aorus-15', 'Gaming laptop with high specs', 14, 1599.99, 8, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(75, 'R75X8', 'Razer Blade Stealth', 'r75x8-razer-blade-stealth', 'High-performance laptop with stealth design', 12, 1499.99, 9, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(76, 'S76P4', 'Samsung Galaxy Book Pro', 's76p4-samsung-galaxy-book-pro', 'Professional laptop with premium features', 18, 1399.99, 10, 3, 1, NULL, 0, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(77, 'D77G5', 'Dell G7 17', 'd77g5-dell-g7-17', 'Gaming laptop with large display', 10, 1599.99, 1, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(78, 'H78X6', 'HP ProBook 450', 'h78x6-hp-probook-450', 'Business laptop with robust features', 16, 1299.99, 2, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(79, 'L79B8', 'Lenovo IdeaPad 3', 'l79b8-lenovo-ideapad-3', 'Affordable laptop for everyday use', 25, 999.99, 3, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(80, 'A80X7', 'Apple MacBook 13', 'a80x7-apple-macbook-13', 'Compact laptop with high performance', 14, 1399.99, 4, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(81, 'A81G5', 'Asus TUF Gaming', 'a81g5-asus-tuf-gaming', 'Durable gaming laptop with good specs', 20, 1399.99, 5, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(82, 'M82B4', 'MSI GF63', 'm82b4-msi-gf63', 'Affordable gaming laptop with decent specs', 16, 1199.99, 7, 3, 2, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(83, 'G83T9', 'Gigabyte G5 Gaming', 'g83t9-gigabyte-g5-gaming', 'Gaming laptop with great performance', 12, 1499.99, 8, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(84, 'R84X7', 'Razer Blade 17', 'r84x7-razer-blade-17', 'High-end gaming laptop with large display', 14, 1799.99, 9, 3, 1, 2, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(85, 'S85P2', 'Samsung Notebook 7', 's85p2-samsung-notebook-7', 'High-performance laptop with sleek design', 18, 1299.99, 10, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(86, 'D86G8', 'Dell Inspiron 15', 'd86g8-dell-inspiron-15', 'Versatile laptop for various needs', 22, 1099.99, 1, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(87, 'H87X9', 'HP Pavilion x360', 'h87x9-hp-pavilion-x360', 'Convertible laptop with solid performance', 18, 1299.99, 2, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(88, 'L88B5', 'Lenovo ThinkPad X1', 'l88b5-lenovo-thinkpad-x1', 'Premium business laptop with high performance', 15, 1599.99, 3, 3, 2, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(89, 'A89X3', 'Apple MacBook Pro 16', 'a89x3-apple-macbook-pro-16', 'High-end professional laptop', 13, 1999.99, 4, 3, 2, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(90, 'A90G7', 'Asus ZenBook Pro', 'a90g7-asus-zenbook-pro', 'Professional ultrabook with advanced features', 10, 1699.99, 5, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(91, 'M91B6', 'MSI GE76 Raider', 'm91b6-msi-ge76-raider', 'High-end gaming laptop with exceptional specs', 7, 1799.99, 7, 3, 2, NULL, 0, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(92, 'G92T5', 'Gigabyte G5 Pro', 'g92t5-gigabyte-g5-pro', 'Gaming laptop with premium features', 14, 1599.99, 8, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(93, 'R93X4', 'Razer Blade Stealth 13', 'r93x4-razer-blade-stealth-13', 'Compact high-performance laptop', 12, 1299.99, 9, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(94, 'S94M6', 'Samsung Odyssey 15', 's94m6-samsung-odyssey-15', 'Gaming laptop with powerful specs', 18, 1499.99, 10, 3, 1, NULL, 0, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200));

insert into product_attribute([product_id], [attribute_set_id], [attribute_value_id])
select p.id, ats.id, (select id from attribute_value where attribute_set_id = ats.id order by random() limit 1)
from product as p
join category as c on c.id = p.category_id
join attribute_group as atg on atg.id = c.attribute_group_id
join attribute_set as ats on ats.attribute_group_id = atg.id;

-- makr sure, that cpu model matches cpu type
update product_attribute
set attribute_value_id = (
    select pa.attribute_value_id + 10
    from product_attribute as pa
    where pa.product_id = product_attribute.product_id
    and pa.attribute_set_id = 2
)
where attribute_set_id = 3;

INSERT INTO product ([code], [name], [slug], [description], [quantity], [price], [brand_id], [category_id], [supplier_id], [status_id], [is_published], [created_at], [updated_at]) VALUES
('A1B2C', 'Intel Core i9-11900K', 'a1b2c-intel-core-i9-11900k', 'High-performance desktop processor with 8 cores.', 100, 499.99, 12, 5, 1, 1, 1, 1672528800, 1698150000),
('D3E4F', 'AMD Ryzen 9 5900X', 'd3e4f-amd-ryzen-9-5900x', '12-core processor with multi-threading capabilities.', 80, 549.99, 13, 5, 2, 2, 1, 1675120800, 1697235600),
('G5H6J', 'Intel Core i7-11700K', 'g5h6j-intel-core-i7-11700k', '8-core, 16-thread desktop processor for gaming.', 120, 399.99, 12, 5, 1, 3, 1, 1680120000, 1698150000),
('K7L8M', 'AMD Ryzen 7 5800X', 'k7l8m-amd-ryzen-7-5800x', '8-core high-performance desktop processor.', 90, 449.99, 13, 5, 2, NULL, 1, 1677612000, 1698150000),
('N9P0Q', 'Intel Core i5-11600K', 'n9p0q-intel-core-i5-11600k', '6-core processor with Intel Turbo Boost Technology.', 140, 299.99, 12, 5, 1, 1, 1, 1679239200, 1695118800),
('R1S2T', 'AMD Ryzen 5 5600X', 'r1s2t-amd-ryzen-5-5600x', '6-core processor, optimized for gaming and content creation.', 150, 299.99, 13, 5, 1, NULL, 1, 1680811200, 1696414800),
('U3V4W', 'Intel Core i3-10100F', 'u3v4w-intel-core-i3-10100f', '4-core desktop processor for budget builds.', 160, 99.99, 12, 5, 2, NULL, 1, 1678135200, 1693611600),
('X5Y6Z', 'AMD Ryzen 3 3100', 'x5y6z-amd-ryzen-3-3100', '4-core processor, affordable and efficient.', 180, 89.99, 13, 5, 2, NULL, 1, 1676227200, 1692327600),
('A7B8C', 'Intel Pentium Gold G6400', 'a7b8c-intel-pentium-gold-g6400', 'Budget dual-core processor for basic tasks.', 200, 59.99, 12, 5, 1, 2, 1, 1673904000, 1696414800),
('D9E0F', 'AMD Athlon 3000G', 'd9e0f-amd-athlon-3000g', 'Dual-core processor with integrated graphics.', 220, 49.99, 13, 5, 2, NULL, 1, 1673126400, 1693791600),
('G1H2J', 'Intel Xeon W-1290P', 'g1h2j-intel-xeon-w-1290p', '10-core processor designed for workstation builds.', 70, 799.99, 12, 5, 1, NULL, 1, 1678658400, 1694509200),
('K3L4M', 'AMD Ryzen Threadripper 3990X', 'k3l4m-amd-ryzen-threadripper-3990x', '64-core processor designed for massive workloads.', 50, 3999.99, 13, 5, 2, NULL, 1, 1674691200, 1692068400),
('N5P6Q', 'Intel Core i9-10900X', 'n5p6q-intel-core-i9-10900x', '10-core processor with support for overclocking.', 110, 749.99, 12, 5, 1, NULL, 1, 1675545600, 1692942000),
('R7S8T', 'AMD Ryzen 9 3950X', 'r7s8t-amd-ryzen-9-3950x', '16-core processor optimized for high-end gaming.', 95, 699.99, 13, 5, 2, NULL, 1, 1681329600, 1693201200),
('U9V0W', 'Intel Celeron G5905', 'u9v0w-intel-celeron-g5905', 'Dual-core processor for low-budget desktops.', 210, 49.99, 12, 5, 2, NULL, 1, 1679868000, 1692327600),
('X1Y2Z', 'AMD Ryzen 5 3400G', 'x1y2z-amd-ryzen-5-3400g', '4-core APU with integrated Vega graphics.', 170, 159.99, 13, 5, 1, 2, 1, 1682433600, 1694859600),
('A3B4C', 'Intel Core i7-10700K', 'a3b4c-intel-core-i7-10700k', '8-core, 16-thread processor for gaming.', 130, 379.99, 12, 5, 2, 3, 1, 1683756000, 1697593200),
('D5E6F', 'AMD Ryzen 3 3300X', 'd5e6f-amd-ryzen-3-3300x', '4-core budget processor for gaming.', 180, 119.99, 13, 5, 1, 1, 1, 1685203200, 1697593200),
('G7H8J', 'Intel Core i9-10980XE', 'g7h8j-intel-core-i9-10980xe', '18-core processor for high-end workstations.', 60, 1099.99, 12, 5, 2, NULL, 1, 1686436800, 1694067600),
('K9L0M', 'AMD Ryzen Threadripper 3970X', 'k9l0m-amd-ryzen-threadripper-3970x', '32-core processor for massive workloads.', 55, 1999.99, 13, 5, 1, NULL, 1, 1688476800, 1692327600);

INSERT INTO product ([code], [name], [slug], [description], [quantity], [price], [brand_id], [category_id], [supplier_id], [status_id], [is_published], [created_at], [updated_at]) VALUES
('M1B2C', 'Asus ROG Strix Z490-E Gaming', 'm1b2c-asus-rog-strix-z490-e-gaming', 'High-end motherboard for Intel 10th gen processors.', 100, 299.99, 5, 6, 1, NULL, 1, 1672528800, 1697616000),
('M3D4E', 'MSI MPG B550 Gaming Edge', 'm3d4e-msi-mpg-b550-gaming-edge', 'Mid-tier motherboard for AMD Ryzen processors.', 80, 189.99, 7, 6, 2, NULL, 1, 1675221600, 1694115600),
('M5F6G', 'Gigabyte Aorus X570 Master', 'm5f6g-gigabyte-aorus-x570-master', 'Premium motherboard for AMD Ryzen with PCIe 4.0.', 120, 329.99, 8, 6, 1, 3, 1, 1678120000, 1696335600),
('M7H8I', 'ASRock B450M Steel Legend', 'm7h8i-asrock-b450m-steel-legend', 'Micro ATX motherboard for budget AMD builds.', 140, 89.99, 21, 6, 2, NULL, 1, 1679832000, 1697616000),
('M9J0K', 'Asus TUF Gaming B550-PLUS', 'm9j0k-asus-tuf-gaming-b550-plus', 'Durable motherboard for mid-tier AMD systems.', 150, 159.99, 5, 6, 1, 2, 1, 1680324000, 1697510400),
('M1L2M', 'MSI B450 Tomahawk Max', 'm1l2m-msi-b450-tomahawk-max', 'Popular motherboard for budget AMD builds.', 130, 119.99, 7, 6, 2, NULL, 1, 1681428000, 1693804800),
('M3N4O', 'Gigabyte Z490 Aorus Elite', 'm3n4o-gigabyte-z490-aorus-elite', 'ATX motherboard with robust power delivery.', 110, 249.99, 8, 6, 1, NULL, 1, 1679616000, 1695019200),
('M5P6Q', 'ASRock X570 Phantom Gaming 4', 'm5p6q-asrock-x570-phantom-gaming-4', 'Affordable motherboard with PCIe 4.0 for AMD Ryzen.', 160, 139.99, 21, 6, 2, 1, 1, 1679220000, 1693791600),
('M7R8S', 'Asus Prime Z590-P', 'm7r8s-asus-prime-z590-p', 'Motherboard for Intel 11th gen CPUs with PCIe 4.0.', 90, 179.99, 5, 6, 1, NULL, 1, 1677132000, 1694074800),
('M9T0U', 'MSI MAG B560 Tomahawk', 'm9t0u-msi-mag-b560-tomahawk', 'Intel motherboard with next-gen connectivity.', 115, 149.99, 7, 6, 2, 2, 1, 1680129600, 1696335600),
('M1V2W', 'Gigabyte Z590 Aorus Pro AX', 'm1v2w-gigabyte-z590-aorus-pro-ax', 'Premium motherboard with Wi-Fi and PCIe 4.0.', 70, 299.99, 8, 6, 1, NULL, 1, 1676632800, 1694517600),
('M3X4Y', 'ASRock B550 Phantom Gaming-ITX/AX', 'm3x4y-asrock-b550-phantom-gaming-itx-ax', 'Compact ITX motherboard for AMD builds.', 80, 209.99, 21, 6, 2, 2, 1, 1675320000, 1697164800),
('M5Z6A', 'Asus ROG Crosshair VIII Hero', 'm5z6a-asus-rog-crosshair-viii-hero', 'High-end motherboard for AMD Ryzen 5000 series.', 60, 379.99, 5, 6, 1, 1, 1, 1684219200, 1693201200),
('M7B8C', 'MSI MPG Z590 Gaming Carbon WiFi', 'm7b8c-msi-mpg-z590-gaming-carbon-wifi', 'Intel Z590 motherboard with robust overclocking support.', 95, 259.99, 7, 6, 2, NULL, 1, 1682402400, 1693804800),
('M9D0E', 'Gigabyte B450 Aorus M', 'm9d0e-gigabyte-b450-aorus-m', 'Budget-friendly micro ATX motherboard for AMD.', 130, 89.99, 8, 6, 1, NULL, 1, 1676239200, 1696414800),
('M1F2G', 'ASRock Z490 Phantom Gaming 4', 'm1f2g-asrock-z490-phantom-gaming-4', 'Intel Z490 motherboard with high-speed memory support.', 150, 119.99, 21, 6, 2, 1, 1, 1685313600, 1697017200),
('M3H4I', 'Asus Prime B450-PLUS', 'm3h4i-asus-prime-b450-plus', 'Entry-level motherboard for budget AMD systems.', 160, 69.99, 5, 6, 1, NULL, 1, 1685721600, 1697593200),
('M5J6K', 'MSI MEG X570 Unify', 'm5j6k-msi-meg-x570-unify', 'Enthusiast-grade motherboard for AMD with PCIe 4.0.', 85, 319.99, 7, 6, 2, NULL, 1, 1687627200, 1697143200),
('M7L8M', 'Gigabyte X570 Aorus Pro', 'm7l8m-gigabyte-x570-aorus-pro', 'High-performance motherboard with PCIe 4.0 for AMD.', 125, 239.99, 8, 6, 1, 3, 1, 1689628800, 1693190400),
('M9N0O', 'ASRock B550M Pro4', 'm9n0o-asrock-b550m-pro4', 'Micro ATX motherboard for mid-range AMD builds.', 140, 99.99, 21, 6, 2, NULL, 1, 1681939200, 1693201200);

INSERT INTO product ([code], [name], [slug], [description], [quantity], [price], [brand_id], [category_id], [supplier_id], [status_id], [is_published], [created_at], [updated_at]) VALUES
('R1B2C', 'Corsair Vengeance LPX 16GB DDR4', 'r1b2c-corsair-vengeance-lpx-16gb-ddr4', 'High-performance DDR4 memory for gaming.', 200, 79.99, 15, 7, 1, 1, 1, 1672528800, 1697616000),
('R3D4E', 'G.Skill Ripjaws V 16GB DDR4', 'r3d4e-gskill-ripjaws-v-16gb-ddr4', 'Dual-channel memory for gaming and multitasking.', 180, 74.99, 33, 7, 2, NULL, 1, 1675221600, 1694115600),
('R5F6G', 'Kingston HyperX Fury 16GB DDR4', 'r5f6g-kingston-hyperx-fury-16gb-ddr4', 'Reliable DDR4 memory with heat spreader.', 220, 89.99, 18, 7, 1, 3, 1, 1678120000, 1696335600),
('R7H8I', 'Crucial Ballistix 32GB DDR4', 'r7h8i-crucial-ballistix-32gb-ddr4', 'High-density RAM for high-performance tasks.', 140, 159.99, 19, 7, 2, 1, 1, 1679832000, 1697616000),
('R9J0K', 'Corsair Dominator Platinum RGB 16GB DDR4', 'r9j0k-corsair-dominator-platinum-rgb-16gb-ddr4', 'Premium RGB memory with extreme performance.', 95, 199.99, 15, 7, 1, 2, 1, 1680324000, 1697510400),
('R1L2M', 'Patriot Viper 16GB DDR4', 'r1l2m-patriot-viper-16gb-ddr4', 'Affordable memory for entry-level gaming setups.', 210, 69.99, 34, 7, 2, NULL, 1, 1681428000, 1693804800),
('R3N4O', 'ADATA XPG Spectrix D60G 16GB DDR4', 'r3n4o-adata-xpg-spectrix-d60g-16gb-ddr4', 'RGB memory with vibrant lighting effects.', 170, 89.99, 33, 7, 1, NULL, 1, 1679616000, 1695019200),
('R5P6Q', 'TeamGroup T-Force Vulcan Z 16GB DDR4', 'r5p6q-teamgroup-t-force-vulcan-z-16gb-ddr4', 'Budget-friendly memory for gamers.', 230, 59.99, 33, 7, 2, 1, 1, 1679220000, 1693791600),
('R7R8S', 'G.Skill Trident Z RGB 16GB DDR4', 'r7r8s-gskill-trident-z-rgb-16gb-ddr4', 'RGB memory with high performance and style.', 125, 139.99, 33, 7, 1, 3, 1, 1677132000, 1694074800),
('R9T0U', 'Kingston Fury Beast 16GB DDR5', 'r9t0u-kingston-fury-beast-16gb-ddr5', 'Next-gen DDR5 memory for extreme performance.', 110, 229.99, 18, 7, 2, NULL, 1, 1680129600, 1696335600),
('R1V2W', 'Corsair Vengeance RGB Pro 32GB DDR4', 'r1v2w-corsair-vengeance-rgb-pro-32gb-ddr4', 'High-performance RGB memory with overclocking potential.', 100, 179.99, 15, 7, 1, 1, 1, 1676632800, 1694517600),
('R3X4Y', 'Crucial 16GB DDR4', 'r3x4y-crucial-16gb-ddr4', 'Standard DDR4 memory for everyday computing.', 190, 64.99, 19, 7, 2, NULL, 1, 1675320000, 1697164800),
('R5Z6A', 'Patriot Signature Line 16GB DDR4', 'r5z6a-patriot-signature-line-16gb-ddr4', 'Reliable memory for office and home use.', 180, 54.99, 34, 7, 1, 1, 1, 1684219200, 1693201200),
('R7B8C', 'ADATA XPG Gammix D30 16GB DDR4', 'r7b8c-adata-xpg-gammix-d30-16gb-ddr4', 'Memory optimized for gaming and overclocking.', 150, 79.99, 33, 7, 2, NULL, 1, 1682402400, 1693804800),
('R9D0E', 'Corsair Dominator Platinum 32GB DDR5', 'r9d0e-corsair-dominator-platinum-32gb-ddr5', 'Top-tier DDR5 memory for extreme systems.', 70, 299.99, 15, 7, 1, 3, 1, 1676239200, 1696414800),
('R1F2G', 'TeamGroup T-Force Delta RGB 16GB DDR5', 'r1f2g-teamgroup-t-force-delta-rgb-16gb-ddr5', 'DDR5 memory with advanced performance.', 130, 239.99, 33, 7, 2, 1, 1, 1685313600, 1697017200),
('R3H4I', 'Crucial Ballistix 16GB DDR5', 'r3h4i-crucial-ballistix-16gb-ddr5', 'Fast DDR5 memory for gaming and multitasking.', 90, 189.99, 19, 7, 1, NULL, 1, 1685721600, 1697593200),
('R5J6K', 'G.Skill Ripjaws S5 16GB DDR5', 'r5j6k-gskill-ripjaws-s5-16gb-ddr5', 'Next-gen DDR5 memory for enthusiasts.', 85, 199.99, 33, 7, 2, 2, 1, 1687627200, 1697143200),
('R7L8M', 'Patriot Viper Elite II 16GB DDR4', 'r7l8m-patriot-viper-elite-ii-16gb-ddr4', 'Affordable high-performance memory.', 175, 74.99, 34, 7, 1, NULL, 1, 1689628800, 1693190400),
('R9N0O', 'Corsair Vengeance 16GB DDR5', 'r9n0o-corsair-vengeance-16gb-ddr5', 'Cutting-edge DDR5 memory for demanding tasks.', 115, 249.99, 15, 7, 2, 1, 1, 1681939200, 1693201200);

INSERT INTO product ([code], [name], [slug], [description], [quantity], [price], [brand_id], [category_id], [supplier_id], [status_id], [is_published], [created_at], [updated_at]) VALUES
('S1B2C', 'Samsung 970 EVO Plus 1TB NVMe SSD', 's1b2c-samsung-970-evo-plus-1tb-nvme-ssd', 'High-performance NVMe SSD with read speeds up to 3500 MB/s.', 180, 129.99, 10, 8, 1, 1, 1, 1672528800, 1697616000),
('S3D4E', 'Western Digital Blue 1TB SATA SSD', 's3d4e-western-digital-blue-1tb-sata-ssd', 'Reliable SATA SSD with 560 MB/s read speed.', 150, 99.99, 16, 8, 2, 3, 1, 1675221600, 1694115600),
('S5F6G', 'Seagate Barracuda 2TB HDD', 's5f6g-seagate-barracuda-2tb-hdd', 'High-capacity HDD for storing large amounts of data.', 220, 54.99, 17, 8, 1, NULL, 1, 1678120000, 1696335600),
('S7H8I', 'Crucial MX500 1TB SATA SSD', 's7h8i-crucial-mx500-1tb-sata-ssd', 'SATA SSD with advanced data protection and fast speeds.', 170, 89.99, 19, 8, 2, 2, 1, 1679832000, 1697616000),
('S9J0K', 'Samsung 980 PRO 1TB NVMe SSD', 's9j0k-samsung-980-pro-1tb-nvme-ssd', 'NVMe SSD with PCIe Gen 4.0 and blazing speeds.', 100, 199.99, 10, 8, 1, 1, 1, 1680324000, 1697510400),
('S1L2M', 'ADATA XPG SX8200 Pro 1TB NVMe SSD', 's1l2m-adata-xpg-sx8200-pro-1tb-nvme-ssd', 'High-performance NVMe SSD for gaming and content creation.', 140, 129.99, 33, 8, 2, 3, 1, 1681428000, 1693804800),
('S3N4O', 'Western Digital Black SN850 2TB NVMe SSD', 's3n4o-western-digital-black-sn850-2tb-nvme-ssd', 'Premium NVMe SSD with high-end read/write speeds.', 90, 299.99, 16, 8, 1, NULL, 1, 1679616000, 1695019200),
('S5P6Q', 'Seagate FireCuda 510 1TB NVMe SSD', 's5p6q-seagate-firecuda-510-1tb-nvme-ssd', 'High-speed NVMe SSD designed for gaming.', 130, 179.99, 17, 8, 2, NULL, 1, 1679220000, 1693791600),
('S7R8S', 'Crucial P5 Plus 1TB NVMe SSD', 's7r8s-crucial-p5-plus-1tb-nvme-ssd', 'NVMe SSD with blazing speeds and superior endurance.', 160, 149.99, 19, 8, 1, 1, 1, 1677132000, 1694074800),
('S9T0U', 'Samsung 860 EVO 1TB SATA SSD', 's9t0u-samsung-860-evo-1tb-sata-ssd', 'SATA SSD with optimized performance and reliability.', 200, 109.99, 10, 8, 2, NULL, 1, 1680129600, 1696335600),
('S1V2W', 'Western Digital Blue 4TB HDD', 's1v2w-western-digital-blue-4tb-hdd', 'Large-capacity HDD ideal for mass data storage.', 220, 79.99, 16, 8, 1, 1, 1, 1676632800, 1694517600),
('S3X4Y', 'Seagate IronWolf 4TB NAS HDD', 's3x4y-seagate-ironwolf-4tb-nas-hdd', 'HDD designed for NAS environments with enhanced reliability.', 130, 109.99, 17, 8, 2, NULL, 1, 1675320000, 1697164800),
('S5Z6A', 'Crucial BX500 1TB SATA SSD', 's5z6a-crucial-bx500-1tb-sata-ssd', 'Affordable SSD with fast read/write speeds.', 170, 79.99, 19, 8, 1, 1, 1, 1684219200, 1693201200),
('S7B8C', 'Samsung 870 QVO 2TB SATA SSD', 's7b8c-samsung-870-qvo-2tb-sata-ssd', 'High-capacity SATA SSD for reliable storage.', 150, 189.99, 10, 8, 2, NULL, 1, 1682402400, 1693804800),
('S9D0E', 'Western Digital Black P10 2TB Portable HDD', 's9d0e-western-digital-black-p10-2tb-portable-hdd', 'High-capacity portable HDD for on-the-go storage.', 120, 119.99, 16, 8, 1, 2, 1, 1676239200, 1696414800),
('S1F2G', 'Seagate Barracuda Fast SSD 1TB', 's1f2g-seagate-barracuda-fast-ssd-1tb', 'Portable SSD with high-speed data transfer.', 100, 139.99, 17, 8, 2, NULL, 1, 1685313600, 1697017200),
('S3H4I', 'Crucial X6 2TB Portable SSD', 's3h4i-crucial-x6-2tb-portable-ssd', 'Compact portable SSD for quick and easy storage.', 180, 169.99, 19, 8, 1, NULL, 1, 1685721600, 1697593200),
('S5J6K', 'ADATA SE800 1TB Portable SSD', 's5j6k-adata-se800-1tb-portable-ssd', 'Portable SSD with rugged design and fast transfer speeds.', 130, 159.99, 33, 8, 2, 3, 1, 1687627200, 1697143200),
('S7L8M', 'Samsung T7 2TB Portable SSD', 's7l8m-samsung-t7-2tb-portable-ssd', 'Portable SSD with advanced security features.', 110, 249.99, 10, 8, 1, 1, 1, 1689628800, 1693190400),
('S9N0O', 'Western Digital My Passport 2TB Portable HDD', 's9n0o-western-digital-my-passport-2tb-portable-hdd', 'Portable HDD with backup and encryption.', 90, 89.99, 16, 8, 2, NULL, 1, 1681939200, 1693201200);

INSERT INTO product ([code], [name], [slug], [description], [quantity], [price], [brand_id], [category_id], [supplier_id], [status_id], [is_published], [created_at], [updated_at]) VALUES
('V1B2C', 'NVIDIA GeForce RTX 3080 10GB', 'v1b2c-nvidia-geforce-rtx-3080-10gb', 'High-performance graphics card for gaming and content creation.', 120, 699.99, 14, 9, 1, 1, 1, 1672528800, 1697616000),
('V3D4E', 'AMD Radeon RX 6800 XT 16GB', 'v3d4e-amd-radeon-rx-6800-xt-16gb', 'Powerful GPU for smooth 4K gaming and ray tracing.', 110, 649.99, 13, 9, 2, 3, 1, 1675221600, 1694115600),
('V5F6G', 'NVIDIA GeForce RTX 3070 8GB', 'v5f6g-nvidia-geforce-rtx-3070-8gb', 'High-end GPU designed for 1440p gaming.', 130, 499.99, 14, 9, 1, NULL, 1, 1678120000, 1696335600),
('V7H8I', 'AMD Radeon RX 6700 XT 12GB', 'v7h8i-amd-radeon-rx-6700-xt-12gb', 'Efficient GPU for high-resolution gaming.', 100, 479.99, 13, 9, 2, 1, 1, 1679832000, 1697616000),
('V9J0K', 'NVIDIA GeForce RTX 3090 24GB', 'v9j0k-nvidia-geforce-rtx-3090-24gb', 'Top-tier GPU for gaming, AI, and deep learning.', 70, 1499.99, 14, 9, 1, 2, 1, 1680324000, 1697510400),
('V1L2M', 'AMD Radeon RX 6900 XT 16GB', 'v1l2m-amd-radeon-rx-6900-xt-16gb', 'High-end GPU with advanced ray tracing and compute units.', 80, 899.99, 13, 9, 2, 1, 1, 1681428000, 1693804800),
('V3N4O', 'ZOTAC GeForce RTX 3060 Ti 8GB', 'v3n4o-zotac-geforce-rtx-3060-ti-8gb', 'Efficient GPU for 1080p and 1440p gaming.', 150, 399.99, 30, 9, 1, NULL, 1, 1679616000, 1695019200),
('V5P6Q', 'Sapphire Radeon RX 6600 XT 8GB', 'v5p6q-sapphire-radeon-rx-6600-xt-8gb', 'Affordable GPU for entry-level gaming.', 170, 329.99, 32, 9, 2, 1, 1, 1679220000, 1693791600),
('V7R8S', 'EVGA GeForce GTX 1660 Super 6GB', 'v7r8s-evga-geforce-gtx-1660-super-6gb', 'Mid-range GPU for smooth 1080p gaming.', 190, 229.99, 20, 9, 1, 3, 1, 1677132000, 1694074800),
('V9T0U', 'MSI GeForce RTX 3060 12GB', 'v9t0u-msi-geforce-rtx-3060-12gb', 'Solid GPU for budget-friendly 1080p and 1440p gaming.', 180, 379.99, 7, 9, 2, NULL, 1, 1680129600, 1696335600),
('V1V2W', 'Gigabyte AORUS Radeon RX 6800 XT 16GB', 'v1v2w-gigabyte-aorus-radeon-rx-6800-xt-16gb', 'Overclocked GPU for high-performance gaming.', 90, 699.99, 8, 9, 1, 1, 1, 1676632800, 1694517600),
('V3X4Y', 'ASUS ROG Strix RTX 3080 10GB', 'v3x4y-asus-rog-strix-rtx-3080-10gb', 'Overclocked RTX 3080 with advanced cooling.', 80, 749.99, 5, 9, 2, NULL, 1, 1675320000, 1697164800),
('V5Z6A', 'NVIDIA Quadro P2200 5GB', 'v5z6a-nvidia-quadro-p2200-5gb', 'Professional GPU for CAD and 3D rendering.', 50, 429.99, 14, 9, 1, 2, 1, 1684219200, 1693201200),
('V7B8C', 'AMD Radeon Pro W5700 8GB', 'v7b8c-amd-radeon-pro-w5700-8gb', 'Workstation GPU optimized for content creation and 3D design.', 60, 799.99, 13, 9, 2, NULL, 1, 1682402400, 1693804800),
('V9D0E', 'ZOTAC Gaming GeForce RTX 3090 Trinity 24GB', 'v9d0e-zotac-gaming-geforce-rtx-3090-trinity-24gb', 'High-end GPU for 4K gaming and content creation.', 40, 1449.99, 30, 9, 1, 3, 1, 1676239200, 1696414800),
('V1F2G', 'EVGA GeForce RTX 2080 Super 8GB', 'v1f2g-evga-geforce-rtx-2080-super-8gb', 'Powerful GPU for high-end gaming and VR.', 110, 599.99, 20, 9, 2, NULL, 1, 1685313600, 1697017200),
('V3H4I', 'Gigabyte GeForce RTX 2060 6GB', 'v3h4i-gigabyte-geforce-rtx-2060-6gb', 'Mid-range GPU with ray tracing support.', 170, 349.99, 8, 9, 1, 1, 1, 1685721600, 1697593200),
('V5J6K', 'ASUS TUF Gaming Radeon RX 6700 XT 12GB', 'v5j6k-asus-tuf-gaming-radeon-rx-6700-xt-12gb', 'Durable GPU with solid cooling and performance.', 90, 479.99, 5, 9, 2, 1, 1, 1687627200, 1697143200),
('V7L8M', 'NVIDIA GeForce GTX 1650 4GB', 'v7l8m-nvidia-geforce-gtx-1650-4gb', 'Entry-level GPU for basic gaming and light workloads.', 200, 149.99, 14, 9, 1, NULL, 1, 1689628800, 1693190400),
('V9N0O', 'Sapphire Radeon RX 6500 XT 4GB', 'v9n0o-sapphire-radeon-rx-6500-xt-4gb', 'Budget-friendly GPU for casual gaming.', 210, 199.99, 32, 9, 2, NULL, 1, 1681939200, 1693201200);

INSERT INTO product ([code], [name], [slug], [description], [quantity], [price], [brand_id], [category_id], [supplier_id], [status_id], [is_published], [created_at], [updated_at]) VALUES
('P1B2C', 'Corsair RM850x 850W 80+ Gold', 'p1b2c-corsair-rm850x-850w-80-plus-gold', 'High-performance 850W power supply with 80+ Gold certification.', 150, 129.99, 15, 10, 1, 1, 1, 1672528800, 1697616000),
('P3D4E', 'EVGA SuperNOVA 750 G5 750W 80+ Gold', 'p3d4e-evga-supernova-750-g5-750w-80-plus-gold', 'Reliable 750W power supply with full modular design.', 120, 119.99, 20, 10, 2, 2, 1, 1675221600, 1694115600),
('P5F6G', 'Cooler Master MWE Gold 650W 80+ Gold', 'p5f6g-cooler-master-mwe-gold-650w-80-plus-gold', 'Budget-friendly 650W power supply with 80+ Gold certification.', 180, 89.99, 23, 10, 1, NULL, 1, 1678120000, 1696335600),
('P7H8I', 'Thermaltake Toughpower GF1 850W 80+ Gold', 'p7h8i-thermaltake-toughpower-gf1-850w-80-plus-gold', 'High-efficiency 850W power supply with RGB lighting.', 130, 139.99, 22, 10, 2, 3, 1, 1679832000, 1697616000),
('P9J0K', 'Corsair CX750M 750W 80+ Bronze', 'p9j0k-corsair-cx750m-750w-80-plus-bronze', 'Semi-modular 750W power supply with 80+ Bronze certification.', 140, 99.99, 15, 10, 1, 1, 1, 1680324000, 1697510400),
('P1L2M', 'EVGA 600 W1 600W 80+ White', 'p1l2m-evga-600-w1-600w-80-plus-white', 'Affordable 600W power supply with essential protections.', 170, 49.99, 20, 10, 2, NULL, 1, 1681428000, 1693804800),
('P3N4O', 'Be Quiet! Straight Power 11 750W 80+ Platinum', 'p3n4o-be-quiet-straight-power-11-750w-80-plus-platinum', 'Ultra-quiet 750W power supply with high efficiency.', 100, 169.99, 42, 10, 1, NULL, 1, 1679616000, 1695019200),
('P5P6Q', 'Seasonic Focus GX-750 750W 80+ Gold', 'p5p6q-seasonic-focus-gx-750-750w-80-plus-gold', 'Compact 750W power supply with 80+ Gold certification.', 160, 109.99, 16, 10, 2, 1, 1, 1679220000, 1693791600),
('P7R8S', 'Corsair RM650x 650W 80+ Gold', 'p7r8s-corsair-rm650x-650w-80-plus-gold', '650W power supply with fully modular cables and high efficiency.', 150, 109.99, 15, 10, 1, 3, 1, 1677132000, 1694074800),
('P9T0U', 'Cooler Master V750 750W 80+ Gold', 'p9t0u-cooler-master-v750-750w-80-plus-gold', '750W fully modular power supply with 80+ Gold certification.', 130, 119.99, 23, 10, 2, NULL, 1, 1680129600, 1696335600),
('P1V2W', 'Thermaltake Smart 500W 80+ White', 'p1v2w-thermaltake-smart-500w-80-plus-white', 'Entry-level 500W power supply with essential protections.', 190, 39.99, 22, 10, 1, 1, 1, 1676632800, 1694517600),
('P3X4Y', 'EVGA SuperNOVA 850 G6 850W 80+ Gold', 'p3x4y-evga-supernova-850-g6-850w-80-plus-gold', 'Reliable 850W power supply with compact design.', 100, 139.99, 20, 10, 2, NULL, 1, 1675320000, 1697164800),
('P5Z6A', 'Be Quiet! Pure Power 11 600W 80+ Gold', 'p5z6a-be-quiet-pure-power-11-600w-80-plus-gold', 'Quiet and efficient 600W power supply.', 130, 89.99, 42, 10, 1, 2, 1, 1684219200, 1693201200),
('P7B8C', 'Corsair SF600 600W 80+ Platinum', 'p7b8c-corsair-sf600-600w-80-plus-platinum', 'Compact 600W power supply with Platinum efficiency.', 140, 149.99, 15, 10, 2, NULL, 1, 1682402400, 1693804800),
('P9D0E', 'Cooler Master MasterWatt 550W 80+ Bronze', 'p9d0e-cooler-master-masterwatt-550w-80-plus-bronze', '550W power supply with semi-fanless operation.', 180, 69.99, 23, 10, 1, NULL, 1, 1676239200, 1696414800),
('P1F2G', 'Thermaltake Toughpower GX1 600W 80+ Gold', 'p1f2g-thermaltake-toughpower-gx1-600w-80-plus-gold', '600W power supply with superior efficiency.', 160, 99.99, 22, 10, 2, 1, 1, 1685313600, 1697017200),
('P3H4I', 'Seasonic Prime TX-850 850W 80+ Titanium', 'p3h4i-seasonic-prime-tx-850-850w-80-plus-titanium', 'Ultra-high efficiency 850W power supply with 80+ Titanium.', 90, 229.99, 16, 10, 1, 1, 1, 1685721600, 1697593200),
('P5J6K', 'Be Quiet! Dark Power Pro 12 1200W 80+ Titanium', 'p5j6k-be-quiet-dark-power-pro-12-1200w-80-plus-titanium', '1200W power supply with ultra-quiet performance and high efficiency.', 50, 399.99, 42, 10, 2, 1, 1, 1687627200, 1697143200),
('P7L8M', 'Corsair HX1000 1000W 80+ Platinum', 'p7l8m-corsair-hx1000-1000w-80-plus-platinum', 'Premium 1000W power supply with full modular cables.', 110, 209.99, 15, 10, 1, NULL, 1, 1689628800, 1693190400),
('P9N0O', 'Cooler Master MWE Bronze 500W 80+ Bronze', 'p9n0o-cooler-master-mwe-bronze-500w-80-plus-bronze', 'Affordable 500W power supply with 80+ Bronze efficiency.', 200, 49.99, 23, 10, 2, NULL, 1, 1681939200, 1693201200);

insert into payment_method([id], [name], [position]) values
(1, 'Cash', 1),
(2, 'Bank', 2);

insert into order_status([id], [name], [color], [in_counter], [position]) values
(1, 'Created', '0ea5e9', 1, 1),
(2, 'In progress', '0ea5e9', 1, 2),
(3, 'Canceled', 'a9a9a9', 0, 3),
(4, 'Successful', '000000', 0, 4);

INSERT INTO order_header (order_status_id, payment_method_id, email, first_name, last_name, phone_number, delivery_address, comment, created_at, updated_at) VALUES
(4, 1, 'john.doe@example.com', 'John', 'Doe', '555-1234', '123 Main St', 'Please deliver in the morning.', 1680000000, 1680003600),
(4, 2, 'jane.smith@example.com', 'Jane', 'Smith', '555-5678', '456 Oak St', 'Leave package at the door.', 1683000000, 1683003600),
(3, 1, 'bob.jones@example.com', 'Bob', 'Jones', '555-8765', '789 Pine St', 'Call upon arrival.', 1685000000, 1685007200),
(4, 1, 'alice.brown@example.com', 'Alice', 'Brown', '555-4321', '321 Cedar St', 'No comments.', 1687000000, 1687007200),
(4, 2, 'mike.johnson@example.com', 'Mike', 'Johnson', '555-9999', '101 Maple Ave', 'None.', 1689000000, 1689003600),
(3, 2, 'emily.davis@example.com', 'Emily', 'Davis', '555-1239', '202 Willow St', 'Call first.', 1690000000, 1690003600),
(4, 1, 'chris.white@example.com', 'Chris', 'White', '555-9876', '103 Birch St', '', 1691000000, 1691003600),
(4, 1, 'karen.green@example.com', 'Karen', 'Green', '555-5643', '234 Oakwood Ln', 'Ring the bell twice.', 1692000000, 1692003600),
(4, 2, 'paul.miller@example.com', 'Paul', 'Miller', '555-6543', '98 Spruce St', '', 1693000000, 1693003600),
(3, 2, 'nancy.moore@example.com', 'Nancy', 'Moore', '555-3210', '67 Aspen Ave', '', 1694000000, 1694003600),
(4, 1, 'tony.hall@example.com', 'Tony', 'Hall', '555-9123', '12 Cedarwood Dr', 'No calls please.', 1695000000, 1695003600),
(4, 1, 'laura.taylor@example.com', 'Laura', 'Taylor', '555-7890', '66 Redwood Dr', '', 1696000000, 1696003600),
(4, 1, 'eric.thomas@example.com', 'Eric', 'Thomas', '555-3476', '77 Palm St', 'Knock twice.', 1697000000, 1697003600),
(3, 2, 'diana.lee@example.com', 'Diana', 'Lee', '555-7654', '88 Willow Ct', '', 1698000000, 1698003600),
(4, 2, 'james.adams@example.com', 'James', 'Adams', '555-6574', '99 Cypress Ave', '', 1699000000, 1699003600),
(4, 1, 'rachel.baker@example.com', 'Rachel', 'Baker', '555-2345', '44 Maplewood St', '', 1700000000, 1700003600),
(4, 2, 'matt.clark@example.com', 'Matt', 'Clark', '555-9874', '34 Oak Ln', '', 1701000000, 1701003600),
(3, 1, 'sarah.johnson@example.com', 'Sarah', 'Johnson', '555-2348', '45 Birch Rd', '', 1702000000, 1702003600),
(4, 2, 'daniel.evans@example.com', 'Daniel', 'Evans', '555-4765', '22 Pine Ln', '', 1703000000, 1703003600),
(2, 1, 'olivia.wilson@example.com', 'Olivia', 'Wilson', '555-3947', '78 Redwood St', 'Call if needed.', 1704000000, 1704003600);

INSERT INTO order_line (order_id, product_code, product_snapshot, quantity, price) VALUES
(1, 'COM10001', 'Laptop 15" Intel i5, 8GB RAM', 1, 799.99),
(1, 'COM20001', 'Wireless Mouse', 2, 25.99),
(2, 'LAP30001', 'Gaming Laptop Ryzen 7, 16GB RAM', 1, 1199.99),
(3, 'MON40001', '27" 4K Monitor', 1, 299.99),
(3, 'ACC50001', 'USB-C Hub', 1, 49.99),
(4, 'COM10002', 'Laptop 13" Intel i7, 16GB RAM', 1, 999.99),
(5, 'LAP30002', 'Gaming Desktop Intel i9, 32GB RAM', 1, 1599.99),
(6, 'MON40002', '32" Curved Monitor', 1, 499.99),
(7, 'COM20002', 'Mechanical Keyboard', 1, 89.99),
(8, 'LAP30003', 'Laptop 14" AMD Ryzen 5, 8GB RAM', 1, 699.99),
(9, 'COM10003', 'Ultrabook Intel i5, 8GB RAM', 1, 899.99),
(9, 'ACC50002', 'Laptop Stand', 1, 29.99),
(10, 'LAP30004', 'Gaming Laptop Intel i7, 16GB RAM', 1, 1299.99),
(11, 'MON40003', '24" Full HD Monitor', 1, 199.99),
(12, 'COM10004', 'Laptop 15" AMD Ryzen 7, 16GB RAM', 1, 1099.99),
(13, 'ACC50003', 'Wireless Keyboard', 1, 49.99),
(14, 'COM20003', 'All-in-One Desktop Intel i5, 8GB RAM', 1, 1199.99),
(15, 'MON40004', 'Dual 24" Monitors', 1, 399.99),
(16, 'COM10005', 'Gaming Laptop Intel i9, 32GB RAM', 1, 1899.99),
(17, 'ACC50004', 'External SSD 1TB', 1, 149.99),
(18, 'COM20004', 'Custom Desktop Ryzen 9, 64GB RAM', 1, 2499.99),
(19, 'MON40005', 'Portable Monitor 15.6"', 1, 199.99),
(20, 'COM10006', 'Laptop 13" Intel i5, 8GB RAM', 1, 799.99);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
delete from order_line;
delete from order_header;
delete from order_status;
delete from payment_method;
delete from product_attribute;
delete from product;
delete from product_status;
delete from supplier;
delete from category;
delete from brand;
delete from attribute_value;
delete from attribute_set;
delete from attribute_group;
update sqlite_sequence set seq = 0 where name <> 'goose_db_version';
-- +goose StatementEnd
