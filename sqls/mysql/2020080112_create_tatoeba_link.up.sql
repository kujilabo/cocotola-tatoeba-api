create table `tatoeba_link` (
 `from` int not null
,`to` int not null
,unique(`from`, `to`)
,foreign key(`from`) references `tatoeba_sentence`(`sentence_number`) on delete cascade
,foreign key(`to`) references `tatoeba_sentence`(`sentence_number`) on delete cascade
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
