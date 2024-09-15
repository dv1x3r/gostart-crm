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
(9, 1, 'Power Supply', 1, 0, 9),
(10, 2, 'Generation', 0, 1, 1),
(11, 2, 'CPU Type', 0, 1, 2),
(12, 2, 'CPU Model', 1, 0, 3),
(13, 2, 'RAM Amount', 1, 1, 4),
(14, 2, 'RAM Type', 1, 1, 5),
(16, 2, 'Disk Capacity', 1, 1, 6),
(17, 2, 'Disk Type', 1, 1, 7),
(18, 2, 'Video Card', 1, 0, 8),
(19, 2, 'Screen Size', 1, 1, 9);

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
(50, 5, 'None', 10);

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
(1, 'Hot', '2693FF', 1),
(2, 'New', '8FDB48', 2),
(3, 'Price', 'CC0814', 3);

insert into product ([id], [code], [name], [slug], [description], [quantity], [price], [brand_id], [category_id], [supplier_id], [status_id], [is_published], [created_at], [updated_at]) VALUES
(1, 'D1XPRO', 'Dell Inspiron X1 Pro', 'd1xpro-dell-inspiron-x1-pro', 'High-performance desktop for gaming', 10, 899.99, 1, 2, 1, 2, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(2, 'H2P4GE', 'HP Pavilion 24', 'h2p4ge-hp-pavilion-24', 'All-in-one desktop with 4K display', 15, 1199.99, 2, 2, 1, 1, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(3, 'L3TGE', 'Lenovo ThinkCentre T3', 'l3tge-lenovo-thinkcentre-t3', 'Compact desktop for professionals', 20, 799.99, 3, 2, 2, 3, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(4, 'A4IMX', 'Apple iMac X', 'a4imx-apple-imac-x', 'Sleek design with powerful specs', 12, 1499.99, 4, 2, 1, 2, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(5, 'A5XPS', 'Asus XPS 990', 'a5xps-asus-xps-990', 'Gaming desktop with advanced graphics', 8, 1299.99, 5, 2, 1, 1, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(6, 'M6Z12', 'MSI Z12 Gaming PC', 'm6z12-msi-z12-gaming-pc', 'High-end gaming rig', 5, 1599.99, 7, 2, 1, 3, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(7, 'G7R5T', 'Gigabyte R5 Tower', 'g7r5t-gigabyte-r5-tower', 'Reliable desktop for everyday use', 25, 699.99, 8, 2, 1, 1, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(8, 'R8DAG', 'Razer DeathAdder Gaming PC', 'r8dag-razer-deathadder-gaming-pc', 'Gaming desktop with RGB lighting', 10, 1399.99, 9, 2, 1, 2, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(9, 'S9RAZ', 'Samsung Razer Pro', 's9raz-samsung-razer-pro', 'Professional desktop with SSD', 18, 1099.99, 10, 2, 1, 3, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(10, 'D10WPC', 'Dell Workstation Pro', 'd10wpc-dell-workstation-pro', 'Powerful workstation for heavy tasks', 8, 1399.99, 1, 2, 2, 2, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(11, 'H11OM', 'HP Omen X', 'h11om-hp-omen-x', 'High-performance gaming desktop', 6, 1599.99, 2, 2, 1, 1, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(12, 'L12IM', 'Lenovo IdeaCentre M', 'l12im-lenovo-ideacentre-m', 'Versatile desktop for home and office', 22, 799.99, 3, 2, 1, 3, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(13, 'A13S4', 'Apple Studio 4K', 'a13s4-apple-studio-4k', 'High-resolution desktop for creative work', 9, 1499.99, 4, 2, 1, 2, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(14, 'A14G5', 'Asus Gaming X5', 'a14g5-asus-gaming-x5', 'Gaming desktop with high refresh rate', 7, 1199.99, 5, 2, 1, 3, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(15, 'M15RT', 'MSI Ryzen Tower', 'm15rt-msi-ryzen-tower', 'Desktop with AMD Ryzen processor', 13, 1299.99, 7, 2, 2, 1, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(16, 'G16DE', 'Gigabyte Desk Elite', 'g16de-gigabyte-desk-elite', 'Premium desktop with enhanced performance', 11, 1399.99, 8, 2, 2, 3, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(17, 'R17DA', 'Razer Desktop Advanced', 'r17da-razer-desktop-advanced', 'High-end desktop for gaming enthusiasts', 6, 1499.99, 9, 2, 1, 1, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(18, 'S18P2', 'Samsung Performance 2', 's18p2-samsung-performance-2', 'Desktop with premium performance', 14, 1199.99, 10, 2, 1, 3, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(19, 'D19M', 'Dell Mini PC', 'd19m-dell-mini-pc', 'Compact desktop for small spaces', 20, 699.99, 1, 2, 1, 2, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(20, 'H20LE', 'HP Elite Tower', 'h20le-hp-elite-tower', 'Professional tower for business use', 16, 1199.99, 2, 2, 2, 1, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(21, 'L21U8', 'Lenovo Ultra 8', 'l21u8-lenovo-ultra-8', 'High-performance ultra desktop', 10, 1399.99, 3, 2, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(22, 'A22P5', 'Apple Pro 5', 'a22p5-apple-pro-5', 'Desktop for professional use with high specs', 11, 1499.99, 4, 2, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(23, 'A23G2', 'Asus Gaming 2', 'a23g2-asus-gaming-2', 'Desktop with advanced cooling system', 7, 1299.99, 5, 2, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(24, 'M24M1', 'MSI Mega 1', 'm24m1-msi-mega-1', 'Powerful desktop with high capacity', 9, 1399.99, 7, 2, 2, 1, 0, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(25, 'G25DX', 'Gigabyte DX Series', 'g25dx-gigabyte-dx-series', 'Desktop with powerful graphics', 12, 1199.99, 8, 2, 1, NULL, 0, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(26, 'R26F3', 'Razer Fusion 3', 'r26f3-razer-fusion-3', 'Gaming desktop with latest features', 8, 1599.99, 9, 2, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(27, 'S27H6', 'Samsung High 6', 's27h6-samsung-high-6', 'High-end desktop with sleek design', 14, 1299.99, 10, 2, 1, NULL, 0, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(28, 'D28WX', 'Dell Workstation X', 'd28wx-dell-workstation-x', 'Professional workstation with top specs', 5, 1499.99, 1, 2, 2, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(29, 'H29MS', 'HP Pro Series', 'h29ms-hp-pro-series', 'Series of desktops for professional use', 12, 1199.99, 2, 2, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(30, 'L30G4', 'Lenovo Gaming 4', 'l30g4-lenovo-gaming-4', 'High-performance gaming desktop', 11, 1399.99, 3, 2, 2, 1, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(31, 'A31P7', 'Apple Pro 7', 'a31p7-apple-pro-7', 'Professional desktop with enhanced specs', 6, 1499.99, 4, 2, 1, NULL, 0, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(32, 'A32X1', 'Asus X1 Elite', 'a32x1-asus-x1-elite', 'Elite gaming desktop with high-end specs', 8, 1399.99, 5, 2, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(33, 'M33M9', 'MSI M9 Max', 'm33m9-msi-m9-max', 'Maximized performance desktop', 9, 1499.99, 7, 2, 2, 1, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(34, 'G34Z8', 'Gigabyte Z8 Series', 'g34z8-gigabyte-z8-series', 'Series desktop with high specs', 12, 1199.99, 8, 2, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(35, 'R35V5', 'Razer V5 Elite', 'r35v5-razer-v5-elite', 'Elite gaming desktop with top features', 10, 1599.99, 9, 2, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(36, 'S36X3', 'Samsung X3 Pro', 's36x3-samsung-x3-pro', 'High-performance desktop with pro specs', 15, 1399.99, 10, 2, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(37, 'D37GX', 'Dell Gaming X', 'd37gx-dell-gaming-x', 'Gaming desktop with advanced graphics', 7, 1499.99, 1, 2, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(38, 'H38P3', 'HP Power 3', 'h38p3-hp-power-3', 'Powerful desktop for high-demand tasks', 14, 1299.99, 2, 2, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(39, 'L39G6', 'Lenovo Gaming 6', 'l39g6-lenovo-gaming-6', 'High-end gaming desktop with great specs', 11, 1399.99, 3, 2, 2, 1, 0, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(40, 'A40X9', 'Apple X9 Professional', 'a40x9-apple-x9-professional', 'Top-of-the-line professional desktop', 6, 1599.99, 4, 2, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(41, 'D41E4', 'Dell XPS 15', 'd41e4-dell-xps-15', 'High-performance laptop for professionals', 25, 1499.99, 1, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(42, 'H42S3', 'HP Spectre x360', 'h42s3-hp-spectre-x360', 'Convertible laptop with high specs', 20, 1399.99, 2, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(43, 'L43T8', 'Lenovo ThinkPad T14', 'l43t8-lenovo-thinkpad-t14', 'Reliable business laptop with great performance', 30, 1299.99, 3, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(44, 'A44M2', 'Apple MacBook Pro', 'a44m2-apple-macbook-pro', 'Professional laptop with M2 chip', 18, 1799.99, 4, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(45, 'A45G6', 'Asus ROG Zephyrus', 'a45g6-asus-rog-zephyrus', 'Gaming laptop with high-end graphics', 22, 1499.99, 5, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(46, 'M46B3', 'MSI Bravo 15', 'm46b3-msi-bravo-15', 'Affordable gaming laptop with decent specs', 20, 1199.99, 7, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(47, 'G47A7', 'Gigabyte Aero 15', 'g47a7-gigabyte-aero-15', 'High-performance laptop for creators', 15, 1599.99, 8, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(48, 'R48V6', 'Razer Blade 15', 'r48v6-razer-blade-15', 'Premium gaming laptop with top specs', 10, 1699.99, 9, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(49, 'S49X9', 'Samsung Galaxy Book', 's49x9-samsung-galaxy-book', 'Convertible laptop with AMOLED screen', 12, 1399.99, 10, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(50, 'D50Q7', 'Dell Latitude 7420', 'd50q7-dell-latitude-7420', 'Business laptop with high durability', 24, 1299.99, 1, 3, 2, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(51, 'H51G2', 'HP Envy 14', 'h51g2-hp-envy-14', 'Stylish laptop with high-end features', 18, 1399.99, 2, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(52, 'L52S5', 'Lenovo Yoga 7i', 'l52s5-lenovo-yoga-7i', 'Convertible laptop with high performance', 25, 1299.99, 3, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(53, 'A53M1', 'Apple MacBook Air', 'a53m1-apple-macbook-air', 'Lightweight laptop with impressive specs', 15, 1199.99, 4, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(54, 'A54X8', 'Asus ZenBook 14', 'a54x8-asus-zenbook-14', 'High-performance ultrabook', 20, 1399.99, 5, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(55, 'M55B9', 'MSI Creator 17', 'm55b9-msi-creator-17', 'Laptop for creative professionals', 10, 1699.99, 7, 3, 2, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(56, 'G56T3', 'Gigabyte G5', 'g56t3-gigabyte-g5', 'Affordable gaming laptop with good specs', 22, 1199.99, 8, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(57, 'R57P6', 'Razer Stealth 13', 'r57p6-razer-stealth-13', 'Compact yet powerful laptop', 16, 1499.99, 9, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(58, 'S58M7', 'Samsung Notebook 9', 's58m7-samsung-notebook-9', 'High-performance notebook with sleek design', 18, 1399.99, 10, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(59, 'D59A2', 'Dell Inspiron 14', 'd59a2-dell-inspiron-14', 'Affordable laptop for everyday use', 28, 999.99, 1, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
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
(71, 'A71X9', 'Apple MacBook Air M2', 'a71x9-apple-macbook-air-m2', 'Updated MacBook Air with M2 chip', 14, 1399.99, 4, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(72, 'A72G3', 'Asus Zephyrus G14', 'a72g3-asus-zephyrus-g14', 'Compact gaming laptop with top specs', 12, 1499.99, 5, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(73, 'M73B6', 'MSI Prestige 14', 'm73b6-msi-prestige-14', 'Business laptop with sleek design', 16, 1299.99, 7, 3, 2, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(74, 'G74T9', 'Gigabyte Aorus 15', 'g74t9-gigabyte-aorus-15', 'Gaming laptop with high specs', 14, 1599.99, 8, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(75, 'R75X8', 'Razer Blade Stealth', 'r75x8-razer-blade-stealth', 'High-performance laptop with stealth design', 12, 1499.99, 9, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(76, 'S76P4', 'Samsung Galaxy Book Pro', 's76p4-samsung-galaxy-book-pro', 'Professional laptop with premium features', 18, 1399.99, 10, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(77, 'D77G5', 'Dell G7 17', 'd77g5-dell-g7-17', 'Gaming laptop with large display', 10, 1599.99, 1, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(78, 'H78X6', 'HP ProBook 450', 'h78x6-hp-probook-450', 'Business laptop with robust features', 16, 1299.99, 2, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(79, 'L79B8', 'Lenovo IdeaPad 3', 'l79b8-lenovo-ideapad-3', 'Affordable laptop for everyday use', 25, 999.99, 3, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(80, 'A80X7', 'Apple MacBook 13', 'a80x7-apple-macbook-13', 'Compact laptop with high performance', 14, 1399.99, 4, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(81, 'A81G5', 'Asus TUF Gaming', 'a81g5-asus-tuf-gaming', 'Durable gaming laptop with good specs', 20, 1399.99, 5, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(82, 'M82B4', 'MSI GF63', 'm82b4-msi-gf63', 'Affordable gaming laptop with decent specs', 16, 1199.99, 7, 3, 2, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(83, 'G83T9', 'Gigabyte G5 Gaming', 'g83t9-gigabyte-g5-gaming', 'Gaming laptop with great performance', 12, 1499.99, 8, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(84, 'R84X7', 'Razer Blade 17', 'r84x7-razer-blade-17', 'High-end gaming laptop with large display', 14, 1799.99, 9, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(85, 'S85P2', 'Samsung Notebook 7', 's85p2-samsung-notebook-7', 'High-performance laptop with sleek design', 18, 1299.99, 10, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(86, 'D86G8', 'Dell Inspiron 15', 'd86g8-dell-inspiron-15', 'Versatile laptop for various needs', 22, 1099.99, 1, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(87, 'H87X9', 'HP Pavilion x360', 'h87x9-hp-pavilion-x360', 'Convertible laptop with solid performance', 18, 1299.99, 2, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(88, 'L88B5', 'Lenovo ThinkPad X1', 'l88b5-lenovo-thinkpad-x1', 'Premium business laptop with high performance', 15, 1599.99, 3, 3, 2, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(89, 'A89X3', 'Apple MacBook Pro 16', 'a89x3-apple-macbook-pro-16', 'High-end professional laptop', 13, 1999.99, 4, 3, 2, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(90, 'A90G7', 'Asus ZenBook Pro', 'a90g7-asus-zenbook-pro', 'Professional ultrabook with advanced features', 10, 1699.99, 5, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(91, 'M91B6', 'MSI GE76 Raider', 'm91b6-msi-ge76-raider', 'High-end gaming laptop with exceptional specs', 7, 1799.99, 7, 3, 2, NULL, 0, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(92, 'G92T5', 'Gigabyte G5 Pro', 'g92t5-gigabyte-g5-pro', 'Gaming laptop with premium features', 14, 1599.99, 8, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(93, 'R93X4', 'Razer Blade Stealth 13', 'r93x4-razer-blade-stealth-13', 'Compact high-performance laptop', 12, 1299.99, 9, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200)),
(94, 'S94M6', 'Samsung Odyssey 15', 's94m6-samsung-odyssey-15', 'Gaming laptop with powerful specs', 18, 1499.99, 10, 3, 1, NULL, 1, (SELECT (ABS(RANDOM()) % (1704067199 - 1672531200 + 1)) + 1672531200), (SELECT (ABS(RANDOM()) % (1735689599 - 1704067200 + 1)) + 1704067200));

-- insert into product_attribute([product_id], [attribute_set_id], [attribute_value_id])
-- select id, 1, 1
-- from product;

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
-- +goose StatementEnd
