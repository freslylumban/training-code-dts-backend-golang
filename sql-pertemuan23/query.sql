begin;
insert into employees(name, age, address) values('Salma', 21, 'Jakarta');
end;

start transaction;
insert into employees(name, age, address) values('Salma', 21, 'Jakarta');
commit;
rollback;


start transaction;
insert into customers(name) values('Salma');
commit;
rollback;

start transaction;
insert into orders(customer_id) values(1);
commit;
rollback;

start transaction;
insert into shipments(order_id) values(1);
commit;
rollback;