-- +goose Up
-- +goose StatementBegin
INSERT INTO item (uuid, item, list_id) VALUES
-- Dairy & Eggs
('fce3fcc0-0113-467d-ab6e-4b72073318a4', 'Milk', 0),
('a947c24b-73ba-4c00-869b-584d78b17225', 'Eggs', 0),
('751db320-5751-4f31-b5b9-47e190dc6b25', 'Butter', 0),
('fe1025ea-a89c-4579-a27d-462f54a5615c', 'Cheese', 0),
('95e6319c-01ea-4914-b02e-a43610a285ff', 'Yogurt', 0),
('6b454fdf-ca9d-40c0-8f8c-5f0af27c1052', 'Cream Cheese', 0),
('911a2d1d-5a66-44e5-9b44-7f382c15f0d2', 'Sour Cream', 0),

-- Meat & Poultry
('4bb35e53-50e7-4288-a1b3-4f0fff097df5', 'Chicken Breast', 0),
('c16b295b-4a23-41be-9b57-6c3dcaa92ebc', 'Ground Beef', 0),
('3667f0a9-23f7-4238-bdb1-4dfc40e84f7c', 'Bacon', 0),
('19011e94-e29f-4abf-9386-75108fe6f68c', 'Ham', 0),
('265edf65-5d14-4318-a4e7-0c101abd9ea5', 'Turkey', 0),
('736c7116-613e-4f9b-b932-ddeafb53d006', 'Salmon', 0),
('d78c3f14-84df-480f-a093-a09cd3cd5711', 'Tuna', 0),

-- Fruits
('2b826799-29d7-4377-b5fc-3fc46d49a87f', 'Bananas', 0),
('797d54c4-38ff-4165-87f9-6b44cb7fbd16', 'Apples', 0),
('be962abb-1776-4814-a08d-648179029077', 'Oranges', 0),
('7acd6a80-4086-4677-8a77-7ac47444e916', 'Grapes', 0),
('dd14ffef-dd3e-4e23-a472-db610fdce044', 'Strawberries', 0),
('27e4555d-5d40-4916-baee-1bdad960d09d', 'Blueberries', 0),
('ed51081d-a5b5-4e09-8edd-0a922b649e01', 'Lemons', 0),
('2b2e4120-df2c-4977-8f6b-98bedd953f0d', 'Limes', 0),

-- Vegetables
('319553e8-857e-4ddb-a4c8-e177b7310ab3', 'Carrots', 0),
('a327e5c6-c338-482d-8d5c-40220bf6cd29', 'Broccoli', 0),
('26bf4274-68b1-4965-b9d7-f684fbce0d55', 'Spinach', 0),
('7422b139-3248-4a85-8e11-242cace988cd', 'Lettuce', 0),
('14adae10-72d9-404e-957c-88f8885b9fc1', 'Tomatoes', 0),
('32f215bb-8265-4bdc-9380-426ccf077584', 'Onions', 0),
('1e95f762-3d83-4b11-b65f-6e86af109059', 'Potatoes', 0),
('cbf3786f-2589-408a-b620-7af5511a5011', 'Bell Peppers', 0),
('fc09e35f-1480-42a0-8e49-1fcf9d59efa4', 'Cucumbers', 0),
('0fcf8b02-39ef-4311-a803-9a69098342db', 'Celery', 0),

-- Bread & Bakery
('efaea518-759e-4fac-b4e2-0f011aa7bc92', 'Bread', 0),
('7805a450-9b82-41a4-acae-97bde0a23c3b', 'Bagels', 0),
('94c3bae7-0b76-4ec7-9d50-56bb31cb4875', 'Tortillas', 0),
('1b8d6794-d8f7-4f77-adcf-998918abbabc', 'English Muffins', 0),
('1c2f9a29-8633-4471-91b1-6df54983cd6f', 'Rolls', 0),

-- Pantry Staples
('37405943-4490-4295-80e5-85c75724a5ec', 'Rice', 0),
('ae7b501e-3bcb-44f8-add1-27e336203a8b', 'Pasta', 0),
('4994a120-d1d3-4d41-ba8d-8b43ee324304', 'Flour', 0),
('a03ab72f-526a-4d24-b269-c9aa5aa04167', 'Sugar', 0),
('efd7d377-0b78-4f60-ada1-d9e3ff95adde', 'Salt', 0),
('c81af38c-128b-44f2-a6f1-6a6e6be3a753', 'Black Pepper', 0),
('71908baf-8f82-4d35-8b37-afc4744db9cb', 'Olive Oil', 0),
('868c99ce-c882-4495-9fd9-3a8bd02da4dd', 'Vegetable Oil', 0),
('8a9bd974-cd6b-4833-ada8-dbb489f4ec14', 'Vinegar', 0),

-- Canned Goods
('cefa4d30-c56f-4b6e-a6f8-b77869de5aa0', 'Canned Tomatoes', 0),
('924da387-d8b6-4b9d-bc1c-5938567294b6', 'Canned Beans', 0),
('6754191d-001e-4365-9ddc-f28ae914af06', 'Canned Corn', 0),
('70dd2f5f-833a-4be1-9875-6b8f1d16dba5', 'Chicken Broth', 0),
('13767ac9-08a6-47f6-9358-fa7274742aae', 'Tomato Sauce', 0),

-- Frozen Foods
('645b36a5-c758-466d-9cac-14fef972b4b0', 'Frozen Vegetables', 0),
('775d437f-1589-4e2e-9857-11237ad5ac79', 'Frozen Berries', 0),
('860d3162-6d98-4069-87db-832085ca4419', 'Ice Cream', 0),
('4fc4795e-ca97-4e32-9804-dcb092e7f05f', 'Frozen Pizza', 0),

-- Beverages
('df0af560-b6d8-4891-8677-49be7d029b7d', 'Orange Juice', 0),
('79ec6ad8-65d1-4b33-a7a2-081e5170de0b', 'Coffee', 0),
('1204a710-85ce-4dc4-8b50-cf4f76cb47a8', 'Tea', 0),
('a02386bf-e326-4647-89d1-a7c0dc07e44b', 'Water', 0),
('d765de17-2ee5-4ef8-9da5-97ddcd0fdd11', 'Soda', 0),

-- Snacks
('240e09a9-2824-430e-874c-43ff2a861af0', 'Chips', 0),
('68b2d5d4-35ba-47b3-9f3d-05f89152949a', 'Crackers', 0),
('243dc48d-da03-4797-8ab9-aceeb0700747', 'Nuts', 0),
('c243d582-97ed-4d87-bea3-29281e3530aa', 'Granola Bars', 0),

-- Condiments & Sauces
('c8cbf0da-c198-445e-afe7-7402a006defc', 'Ketchup', 0),
('344a14ce-5b76-4a5c-8b4a-c4814483ec8b', 'Mustard', 0),
('abcb58fb-6f9f-48d6-8355-cf8152e349f3', 'Mayonnaise', 0),
('1f0a9a14-0d0c-4506-9adc-6f11f4fb2488', 'Hot Sauce', 0),
('0e8c2ba3-4197-4469-85ea-5d8144ed2199', 'Soy Sauce', 0),

-- Cleaning & Household
('c4cb020b-a653-4bd9-aa25-fd4f632e6398', 'Dish Soap', 0),
('b6513140-2e85-49b3-8683-dfcc6f98a07e', 'Laundry Detergent', 0),
('8163b0b5-4760-4dc8-afd4-ca8f0e9a0faa', 'Paper Towels', 0),
('943bc80f-8318-40bb-91f8-d2d1cea2faa4', 'Toilet Paper', 0),
('f3d029cc-47d8-4180-87e8-5625b54d0412', 'Trash Bags', 0),

-- Personal Care
('3efbdd01-46bf-4a55-886c-9f7d1f8b81d1', 'Toothpaste', 0),
('5c372bc4-1b57-4cbc-9148-6a58c6b0304e', 'Shampoo', 0),
('048a9972-1421-445e-9ee9-d0dc59c7bb26', 'Soap', 0),
('072abbf3-88e4-4794-8975-c36da52e24af', 'Deodorant', 0);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM item where uuid in (
'fce3fcc0-0113-467d-ab6e-4b72073318a4', 'a947c24b-73ba-4c00-869b-584d78b17225', '751db320-5751-4f31-b5b9-47e190dc6b25',
'fe1025ea-a89c-4579-a27d-462f54a5615c', '95e6319c-01ea-4914-b02e-a43610a285ff', '6b454fdf-ca9d-40c0-8f8c-5f0af27c1052',
'911a2d1d-5a66-44e5-9b44-7f382c15f0d2', '4bb35e53-50e7-4288-a1b3-4f0fff097df5', 'c16b295b-4a23-41be-9b57-6c3dcaa92ebc',
'3667f0a9-23f7-4238-bdb1-4dfc40e84f7c', '19011e94-e29f-4abf-9386-75108fe6f68c', '265edf65-5d14-4318-a4e7-0c101abd9ea5',
'736c7116-613e-4f9b-b932-ddeafb53d006', 'd78c3f14-84df-480f-a093-a09cd3cd5711', '2b826799-29d7-4377-b5fc-3fc46d49a87f',
'797d54c4-38ff-4165-87f9-6b44cb7fbd16', 'be962abb-1776-4814-a08d-648179029077', '7acd6a80-4086-4677-8a77-7ac47444e916',
'dd14ffef-dd3e-4e23-a472-db610fdce044', '27e4555d-5d40-4916-baee-1bdad960d09d', 'ed51081d-a5b5-4e09-8edd-0a922b649e01',
'2b2e4120-df2c-4977-8f6b-98bedd953f0d', '319553e8-857e-4ddb-a4c8-e177b7310ab3', 'a327e5c6-c338-482d-8d5c-40220bf6cd29',
'26bf4274-68b1-4965-b9d7-f684fbce0d55', '7422b139-3248-4a85-8e11-242cace988cd', '14adae10-72d9-404e-957c-88f8885b9fc1',
'32f215bb-8265-4bdc-9380-426ccf077584', '1e95f762-3d83-4b11-b65f-6e86af109059', 'cbf3786f-2589-408a-b620-7af5511a5011',
'fc09e35f-1480-42a0-8e49-1fcf9d59efa4', '0fcf8b02-39ef-4311-a803-9a69098342db', 'efaea518-759e-4fac-b4e2-0f011aa7bc92',
'7805a450-9b82-41a4-acae-97bde0a23c3b', '94c3bae7-0b76-4ec7-9d50-56bb31cb4875', '1b8d6794-d8f7-4f77-adcf-998918abbabc',
'1c2f9a29-8633-4471-91b1-6df54983cd6f', '37405943-4490-4295-80e5-85c75724a5ec', 'ae7b501e-3bcb-44f8-add1-27e336203a8b',
'4994a120-d1d3-4d41-ba8d-8b43ee324304', 'a03ab72f-526a-4d24-b269-c9aa5aa04167', 'efd7d377-0b78-4f60-ada1-d9e3ff95adde',
'c81af38c-128b-44f2-a6f1-6a6e6be3a753', '71908baf-8f82-4d35-8b37-afc4744db9cb', '868c99ce-c882-4495-9fd9-3a8bd02da4dd',
'8a9bd974-cd6b-4833-ada8-dbb489f4ec14', 'cefa4d30-c56f-4b6e-a6f8-b77869de5aa0', '924da387-d8b6-4b9d-bc1c-5938567294b6',
'6754191d-001e-4365-9ddc-f28ae914af06', '70dd2f5f-833a-4be1-9875-6b8f1d16dba5', '13767ac9-08a6-47f6-9358-fa7274742aae',
'645b36a5-c758-466d-9cac-14fef972b4b0', '775d437f-1589-4e2e-9857-11237ad5ac79', '860d3162-6d98-4069-87db-832085ca4419',
'4fc4795e-ca97-4e32-9804-dcb092e7f05f', 'df0af560-b6d8-4891-8677-49be7d029b7d', '79ec6ad8-65d1-4b33-a7a2-081e5170de0b',
'1204a710-85ce-4dc4-8b50-cf4f76cb47a8', 'a02386bf-e326-4647-89d1-a7c0dc07e44b', 'd765de17-2ee5-4ef8-9da5-97ddcd0fdd11',
'240e09a9-2824-430e-874c-43ff2a861af0', '68b2d5d4-35ba-47b3-9f3d-05f89152949a', '243dc48d-da03-4797-8ab9-aceeb0700747',
'c243d582-97ed-4d87-bea3-29281e3530aa', 'c8cbf0da-c198-445e-afe7-7402a006defc', '344a14ce-5b76-4a5c-8b4a-c4814483ec8b',
'abcb58fb-6f9f-48d6-8355-cf8152e349f3', '1f0a9a14-0d0c-4506-9adc-6f11f4fb2488', '0e8c2ba3-4197-4469-85ea-5d8144ed2199',
'c4cb020b-a653-4bd9-aa25-fd4f632e6398', 'b6513140-2e85-49b3-8683-dfcc6f98a07e', '8163b0b5-4760-4dc8-afd4-ca8f0e9a0faa',
'943bc80f-8318-40bb-91f8-d2d1cea2faa4', 'f3d029cc-47d8-4180-87e8-5625b54d0412', '3efbdd01-46bf-4a55-886c-9f7d1f8b81d1',
'5c372bc4-1b57-4cbc-9148-6a58c6b0304e', '048a9972-1421-445e-9ee9-d0dc59c7bb26', '072abbf3-88e4-4794-8975-c36da52e24af')

-- +goose StatementEnd
