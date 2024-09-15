-- +goose Up
-- +goose StatementBegin
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

-- For PC Desktops (attribute_group_id = 1)
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

-- Insert statements for unique top 100 PC and component brands
-- Insert statements for unique top 100 PC and component brands
-- Insert statements for unique top 50 PC and component brands
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
(1, 'stk', 'stk', 'Stock', 'Available', 1, 1),
(2, 'ord', 'ord', 'Order', 'On request', 1, 2);
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
