create table el_example (
    id serial primary key not null ,
    name varchar(64) not null ,
    summary text not null ,
    price float not null,
    create_time timestamp not null default now(),
    update_time timestamp
);

comment on table el_example is '演示表';
comment on column el_example.name is '名称';
comment on column el_example.summary is '简介';
comment on column el_example.price is '价格';
comment on column el_example.create_time is '创建时间';
comment on column el_example.update_time is '更新时间';