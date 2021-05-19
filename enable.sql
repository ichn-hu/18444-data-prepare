set @@tidb_enable_pipelined_window_function = 1;
select min(t.cuno) as start_key, max(t.cuno) as end_key, count(*) as page_size from (select *, row_number() over(order by cuno) as row_num from account) t group by floor((t.row_num-1)/100000) order by start_key;
