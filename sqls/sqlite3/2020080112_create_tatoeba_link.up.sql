create table `tatoeba_link` (
 `from` int not null
,`to` int not null
,unique(`from`, `to`)
,foreign key(`from`) references `tatoeba_sentence`(`sentence_number`)
,foreign key(`to`) references `tatoeba_sentence`(`sentence_number`)
);
