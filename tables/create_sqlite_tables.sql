create table hubid(id INTEGER PRIMARY KEY);

create table locations(cod_res_c2net INTEGER,
		cod_res_node TEXT,
                name_location TEXT,
		PRIMARY KEY (cod_res_c2net)
);

create table collected_values(cod_res_node TEXT,
		cod_sensor TEXT,
		byte5 TEXT,
		byte6 TEXT,
		byte7 TEXT,
		timestamp datetime DEFAULT CURRENT_TIMESTAMP
                );

create table collected_values_barcode(cod_res_node TEXT,
		cod_sensor TEXT,
		collected TEXT,
		timestamp datetime DEFAULT CURRENT_TIMESTAMP
);

create table alive_sensors(cod_sensor TEXT,
		cod_res_node TEXT,
                name_sensor TEXT
);

create table sensor(id TEXT,
                name TEXT,
		tipo TEXT,
		PRIMARY KEY (id,name)
);


insert into hubid(id) values('16');
insert into locations(cod_res_c2net,cod_res_node,name_location) values(20,'C1','Cutting Machine');
insert into locations(cod_res_c2net,cod_res_node,name_location) values(19,'A1','Stamping Machine');
insert into locations(cod_res_c2net,cod_res_node,name_location) values(18,'B1','Painting Station');
insert into locations(cod_res_c2net,cod_res_node,name_location) values(17,'FF','Operations Control');
insert into alive_sensors(cod_sensor, cod_res_node, name_sensor) values('13','C1','Ultrasonic parts counter');
insert into alive_sensors(cod_sensor, cod_res_node, name_sensor) values('13','A1','Ultrasonic parts counter');
insert into alive_sensors(cod_sensor, cod_res_node, name_sensor) values('10','B1','Coating Thickness Gauge');
insert into alive_sensors(cod_sensor, cod_res_node, name_sensor) values('01','FF','Reader');




