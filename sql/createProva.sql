/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40101 SET NAMES utf8 */;

use AprendaMais;

create table if not exists prova(
    id_prova integer not null auto_increment primary key,
    id_faculdade integer,
    id_questao integer not null,
    id_aluno integer not null,
    nota decimal(4,2) not null,
    data datetime not null,
    constraint fk_id_faculdade foreign key (id_faculdade) references faculdade (id_faculdade),
    constraint fk_id_questao foreign key (id_questao) references questao (id_questao),
    constraint fk_id_aluno foreign key (id_aluno) references aluno (id_aluno)
)ENGINE =InnoDB DEFAULT CHARSET=latin1;