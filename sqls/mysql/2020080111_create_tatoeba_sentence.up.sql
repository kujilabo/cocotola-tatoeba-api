create table `tatoeba_sentence` (
 `sentence_number` int not null
,`lang3` varchar(3) character set ascii not null
,`text` varchar(500) not null
,`author` varchar(20) character set ascii not null
,`updated_at` datetime not null
,primary key(`sentence_number`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
